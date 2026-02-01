// Package compare provides functionality for comparing and merging OpenXML elements.
//
//nolint:revive // max-public-structs: This package intentionally exposes multiple strategy types for flexibility.
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

// MessageTextContentChanged is the message used when text content differs between elements.
const MessageTextContentChanged = "Text content changed"

// MergeStrategy defines how to resolve conflicts during a merge.
type MergeStrategy interface {
	Resolve(conflict *Conflict) (any, error)
}

// StrategyOursWins resolves conflicts by choosing the 'ours' value.
type StrategyOursWins struct{}

func (s *StrategyOursWins) Resolve(conflict *Conflict) (any, error) {
	return conflict.OurValue, nil
}

// StrategyTheirsWins resolves conflicts by choosing the 'theirs' value.
type StrategyTheirsWins struct{}

func (s *StrategyTheirsWins) Resolve(conflict *Conflict) (any, error) {
	return conflict.TheirValue, nil
}

// StrategyCustom resolves conflicts using a custom function.
type StrategyCustom struct {
	Resolver func(conflict *Conflict) (any, error)
}

func (s *StrategyCustom) Resolve(conflict *Conflict) (any, error) {
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

func (s *StrategyCombined) Resolve(conflict *Conflict) (any, error) {
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

func (s *StrategyConflictMarkers) Resolve(conflict *Conflict) (any, error) {
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
	for i := range conflicts {
		val, err := m.strategy.Resolve(&conflicts[i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve conflict: %w", err)
		}

		if err := m.applyResolution(merged, &conflicts[i], val); err != nil {
			return nil, nil, fmt.Errorf("failed to apply resolution: %w", err)
		}
	}

	return merged, conflicts, nil
}

func (m *ElementMerger) applyDiffs(target openxml.Element, diffs []Diff) error {
	for i := range diffs {
		if err := m.applyDiff(target, &diffs[i]); err != nil {
			return err
		}
	}

	return nil
}

func (m *ElementMerger) applyDiff(target openxml.Element, diff *Diff) error {
	// Filter based on options
	if m.options.IgnoreAttributes && (diff.Key != "" || diff.Type == Modified && (diff.Message == "Attribute modified" || diff.Message == "Attribute added")) {
		return nil
	}
	if m.options.IgnoreAttributes && (diff.Type == Added || diff.Type == Deleted) && diff.Key != "" {
		return nil
	}

	if m.options.IgnoreText && diff.Message == MessageTextContentChanged {
		return nil
	}

	if m.options.IgnoreChildren && len(diff.ChildDiffs) > 0 {
		return nil
	}

	switch diff.Type {
	case NoDiff:
		// No action needed when there is no difference
		return nil
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

		if diff.Message == MessageTextContentChanged {
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

func (m *ElementMerger) applyResolution(target openxml.Element, conflict *Conflict, value any) error {
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
	case ConflictStructure:
		// For structural conflicts (child element changes), we need to handle at a higher level
		// during the merge process. The resolution value should be an openxml.Element
		if elem, ok := value.(openxml.Element); ok {
			// Replace or update the child element
			if comp, ok := target.(openxml.CompositeElement); ok {
				// Find and replace the child at the conflict index
				children := make([]openxml.Element, 0)
				i := 0
				for child := range comp.Children() {
					if i == conflict.Index {
						children = append(children, elem)
					} else {
						children = append(children, child)
					}
					i++
				}
				// Note: This is a simplified approach. Full implementation would require
				// more sophisticated child element replacement in the CompositeElement interface
			}
		}
	default:
		return fmt.Errorf("unknown conflict type: %v", conflict.Type)
	}

	return nil
}

func (m *ElementMerger) detectConflicts(ours, theirs []Diff) []Conflict {
	var conflicts []Conflict
	for i := range ours {
		for j := range theirs {
			if m.isConflict(&ours[i], &theirs[j]) {
				conflictType := ConflictAttribute
				if ours[i].Message == MessageTextContentChanged {
					conflictType = ConflictText
				}

				conflicts = append(conflicts, Conflict{
					Type:        conflictType,
					Key:         ours[i].Key,
					Index:       ours[i].Index,
					Description: fmt.Sprintf("Conflict at %s", ours[i].Key),
					OurValue:    ours[i].NewValue,
					TheirValue:  theirs[j].NewValue,
					BaseValue:   ours[i].OldValue,
				})
			}
		}
	}

	return conflicts
}

func (m *ElementMerger) isConflict(d1, d2 *Diff) bool {
	if d1.Type == Modified && d2.Type == Modified {
		if d1.Key != "" && d1.Key == d2.Key {
			return d1.NewValue != d2.NewValue
		}
		if d1.Message == MessageTextContentChanged && d2.Message == MessageTextContentChanged {
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
			// Check for index-based conflict for child elements (non-attributes)
			// Attributes have Key set, children use Index for positioning
			// We need to check both the index match and ensure it's a valid modification conflict
			if d.Key == "" && d.Index == c.Index && c.Index != 0 && d.Type == Modified {
				// This is a child element modification at the same index
				isConflicting = true

				break
			}
			if d.Key == "" && d.Index == c.Index {
				// Potential child conflict
				isConflicting = true

				break
			}
			// Check for text conflict
			if d.Message == MessageTextContentChanged && c.Type == ConflictText {
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
