package compare

import (
	"fmt"

	"github.com/connerohnesorge/goffice/openxml"
)

// Conflict represents a conflict detected during a three-way merge.
type Conflict struct {
	Path        string
	Key         string // Key of the conflicting item
	Index       int    // Index of the conflicting item
	Type        ConflictType
	OurValue    interface{}
	TheirValue  interface{}
	BaseValue   interface{}
	Description string
}

// ConflictType represents the type of conflict.
type ConflictType int

const (
	// ConflictAttribute indicates conflicting attribute changes.
	ConflictAttribute ConflictType = iota
	// ConflictText indicates conflicting text content changes.
	ConflictText
	// ConflictStructure indicates conflicting structural changes (children).
	ConflictStructure
)

// MergeStrategy defines how to resolve conflicts during a merge.
type MergeStrategy interface {
	Resolve(conflict Conflict) (interface{}, error)
}

// StrategyOursWins resolves conflicts by choosing the 'ours' value.
type StrategyOursWins struct{}

func (s *StrategyOursWins) Resolve(conflict Conflict) (interface{}, error) {
	return conflict.OurValue, nil
}

// StrategyTheirsWins resolves conflicts by choosing the 'theirs' value.
type StrategyTheirsWins struct{}

func (s *StrategyTheirsWins) Resolve(conflict Conflict) (interface{}, error) {
	return conflict.TheirValue, nil
}

// StrategyCustom resolves conflicts using a custom function.
type StrategyCustom struct {
	Resolver func(conflict Conflict) (interface{}, error)
}

func (s *StrategyCustom) Resolve(conflict Conflict) (interface{}, error) {
	if s.Resolver == nil {
		return nil, fmt.Errorf("no custom resolver defined")
	}
	return s.Resolver(conflict)
}

// StrategyCombined resolves conflicts by combining values where possible.
// For text, it appends 'theirs' to 'ours'. For others, it defaults to 'ours'.
type StrategyCombined struct {
	Separator string
}

func (s *StrategyCombined) Resolve(conflict Conflict) (interface{}, error) {
	if conflict.Type == ConflictText {
		s1, ok1 := conflict.OurValue.(string)
		s2, ok2 := conflict.TheirValue.(string)
		if ok1 && ok2 {
			sep := s.Separator
			if sep == "" {
				sep = " "
			}
			return s1 + sep + s2, nil
		}
	}
	// Default to ours for non-combinable types
	return conflict.OurValue, nil
}

// StrategyConflictMarkers resolves text conflicts by inserting conflict markers.
type StrategyConflictMarkers struct {
	OursLabel   string
	TheirsLabel string
}

func (s *StrategyConflictMarkers) Resolve(conflict Conflict) (interface{}, error) {
	if conflict.Type == ConflictText {
		s1, ok1 := conflict.OurValue.(string)
		s2, ok2 := conflict.TheirValue.(string)
		if ok1 && ok2 {
			ours := s.OursLabel
			if ours == "" {
				ours = "OURS"
			}
			theirs := s.TheirsLabel
			if theirs == "" {
				theirs = "THEIRS"
			}

			return fmt.Sprintf("<<<<<<< %s\n%s\n=======\n%s\n>>>>>>> %s", ours, s1, s2, theirs), nil
		}
	}
	return conflict.OurValue, nil
}

// MergeOptions configures the behavior of the merge operation.
type MergeOptions struct {
	IgnoreAttributes bool
	IgnoreText       bool
	IgnoreChildren   bool
}

// Merger is the interface for merging OpenXML elements.
type Merger interface {
	Merge(base, other openxml.Element) error
	ThreeWayMerge(base, ours, theirs openxml.Element) (openxml.Element, []Conflict, error)
	SetStrategy(strategy MergeStrategy)
	SetOptions(options MergeOptions)
}

// ElementMerger is the default implementation of Merger.
type ElementMerger struct {
	comparator Comparator
	strategy   MergeStrategy
	options    MergeOptions
}

// NewElementMerger creates a new ElementMerger.
func NewElementMerger() *ElementMerger {
	return &ElementMerger{
		comparator: NewElementComparator(),
		strategy:   &StrategyOursWins{}, // Default strategy
	}
}

func (m *ElementMerger) SetStrategy(strategy MergeStrategy) {
	m.strategy = strategy
}

func (m *ElementMerger) SetOptions(options MergeOptions) {
	m.options = options
}

// Merge merges the changes from 'other' into 'base'.
// In a two-way merge, this effectively makes 'base' identical to 'other'
// by applying all differences.
func (m *ElementMerger) Merge(base, other openxml.Element) error {
	if base == nil || other == nil {
		return fmt.Errorf("cannot merge nil elements")
	}

	diffs := m.comparator.Compare(base, other)
	return m.applyDiffs(base, diffs)
}

// ThreeWayMerge merges changes from 'ours' and 'theirs' relative to 'base'.
// It returns a new merged element and a list of conflicts.
func (m *ElementMerger) ThreeWayMerge(base, ours, theirs openxml.Element) (openxml.Element, []Conflict, error) {
	if base == nil || ours == nil || theirs == nil {
		return nil, nil, fmt.Errorf("cannot merge nil elements")
	}

	merged := base.Clone()

	diffsOurs := m.comparator.Compare(base, ours)
	diffsTheirs := m.comparator.Compare(base, theirs)

	conflicts := m.detectConflicts(diffsOurs, diffsTheirs)

	// Apply non-conflicting changes from ours
	if err := m.applyDiffs(merged, diffsOurs); err != nil {
		return nil, nil, fmt.Errorf("failed to apply ours changes: %w", err)
	}

	// Filter theirs diffs to remove conflicts
	safeTheirsDiffs := m.filterConflictingDiffs(diffsTheirs, conflicts)
	if err := m.applyDiffs(merged, safeTheirsDiffs); err != nil {
		return nil, nil, fmt.Errorf("failed to apply theirs changes: %w", err)
	}

	// Resolve conflicts
	for _, conflict := range conflicts {
		val, err := m.strategy.Resolve(conflict)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve conflict: %w", err)
		}

		if err := m.applyResolution(merged, conflict, val); err != nil {
			return nil, nil, fmt.Errorf("failed to apply resolution: %w", err)
		}
	}

	return merged, conflicts, nil
}

func (m *ElementMerger) applyDiffs(target openxml.Element, diffs []Diff) error {
	for _, diff := range diffs {
		if err := m.applyDiff(target, diff); err != nil {
			return err
		}
	}
	return nil
}

func (m *ElementMerger) applyDiff(target openxml.Element, diff Diff) error {
	// Filter based on options
	if m.options.IgnoreAttributes && (diff.Key != "" || diff.Type == Modified && (diff.Message == "Attribute modified" || diff.Message == "Attribute added")) {
		return nil
	}
	if m.options.IgnoreAttributes && (diff.Type == Added || diff.Type == Deleted) && diff.Key != "" {
		return nil
	}

	if m.options.IgnoreText && diff.Message == "Text content changed" {
		return nil
	}

	if m.options.IgnoreChildren && len(diff.ChildDiffs) > 0 {
		return nil
	}

	// fmt.Printf("Applying diff: Type=%v Key=%s Index=%d\n", diff.Type, diff.Key, diff.Index)
	switch diff.Type {

	case Modified:
		if diff.Key != "" {
			if strVal, ok := diff.NewValue.(string); ok {
				attr := openxml.NewAttribute("", diff.Key, "", strVal)
				target.SetAttribute(attr)
			}
			return nil
		}

		if len(diff.ChildDiffs) > 0 {
			comp, ok := target.(openxml.CompositeElement)
			if !ok {
				return fmt.Errorf("cannot apply child diffs to non-composite element")
			}

			children := openxml.ToSlice(comp.Children())
			if diff.Index >= 0 && diff.Index < len(children) {
				child := children[diff.Index]
				if err := m.applyDiffs(child, diff.ChildDiffs); err != nil {
					return err
				}
			} else if diff.Index >= len(children) {
				if diff.Index == len(children) {
					if len(diff.ChildDiffs) == 1 && diff.ChildDiffs[0].Type == Added {
						newChild, ok := diff.ChildDiffs[0].NewValue.(openxml.Element)
						if ok {
							comp.AppendChild(newChild.Clone())
							return nil
						}
					}
				}
			}
			return nil
		}

		if diff.Message == "Text content changed" {
			leaf, ok := target.(openxml.LeafElement)
			if !ok {
				return fmt.Errorf("cannot set text on non-leaf element")
			}
			if strVal, ok := diff.NewValue.(string); ok {
				leaf.SetInnerText(strVal)
			}
		}

	case Added:
		if diff.Key != "" {
			if strVal, ok := diff.NewValue.(string); ok {
				attr := openxml.NewAttribute("", diff.Key, "", strVal)
				target.SetAttribute(attr)
			}
		}
	case Deleted:
		if diff.Key != "" {
			target.RemoveAttribute(diff.Key, "")
		}
	}
	return nil
}

func (m *ElementMerger) applyResolution(target openxml.Element, conflict Conflict, value interface{}) error {
	switch conflict.Type {
	case ConflictAttribute:
		// Handle attribute modification
		if strVal, ok := value.(string); ok {
			// We need to set the attribute
			attr := openxml.NewAttribute("", conflict.Key, "", strVal)
			target.SetAttribute(attr)
		}
	case ConflictText:
		if strVal, ok := value.(string); ok {
			if leaf, ok := target.(openxml.LeafElement); ok {
				leaf.SetInnerText(strVal)
			}
		}
		// TODO: Handle other conflict types
	}
	return nil
}

func (m *ElementMerger) detectConflicts(ours, theirs []Diff) []Conflict {
	var conflicts []Conflict
	// fmt.Printf("Detecting conflicts: %d ours, %d theirs\n", len(ours), len(theirs))
	for _, d1 := range ours {
		for _, d2 := range theirs {
			// fmt.Printf("Checking conflict: d1(Type=%v, Key=%s) vs d2(Type=%v, Key=%s)\n", d1.Type, d1.Key, d2.Type, d2.Key)
			if m.isConflict(d1, d2) {
				conflictType := ConflictAttribute
				if d1.Message == "Text content changed" {
					conflictType = ConflictText
				}

				conflicts = append(conflicts, Conflict{
					Type:        conflictType,
					Key:         d1.Key,
					Index:       d1.Index,
					Description: fmt.Sprintf("Conflict at %s", d1.Key),
					OurValue:    d1.NewValue,
					TheirValue:  d2.NewValue,
					BaseValue:   d1.OldValue,
				})
			}
		}
	}
	return conflicts
}

func (m *ElementMerger) isConflict(d1, d2 Diff) bool {
	if d1.Type == Modified && d2.Type == Modified {
		if d1.Key != "" && d1.Key == d2.Key {
			return d1.NewValue != d2.NewValue
		}
		if d1.Message == "Text content changed" && d2.Message == "Text content changed" {
			return d1.NewValue != d2.NewValue
		}
		if d1.Index == d2.Index {
			return true
		}
	}
	return false
}

func (m *ElementMerger) filterConflictingDiffs(diffs []Diff, conflicts []Conflict) []Diff {
	var safe []Diff
	for _, d := range diffs {
		isConflicting := false
		for _, c := range conflicts {
			if d.Key != "" && d.Key == c.Key {
				isConflicting = true
				break
			}
			if d.Index == c.Index && (d.Type == Modified && c.Index != 0) { // Index 0 is default, need to be careful?
				// Actually Index 0 is valid index. Default int is 0.
				// But attributes have Index 0.
				// ConflictType might help.
				// For now, if Key matches, it's attribute conflict.
				// If Index matches and it's a child diff?
				// My Conflict detection logic sets Index=d1.Index.
				// If d1 is attribute modified, Index is 0.
				// If d1 is child modified, Index is i.
				// We need to distinguish.
				// If Key is set, use Key. If Key is empty, use Index.
			}
			if d.Key == "" && d.Index == c.Index {
				// Potential child conflict
				isConflicting = true
				break
			}
			// Check for text conflict
			if d.Message == "Text content changed" && c.Type == ConflictText {
				isConflicting = true
				break
			}
		}
		if !isConflicting {
			safe = append(safe, d)
		}
	}
	return safe
}
