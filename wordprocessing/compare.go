package wordprocessing

import (
	"errors"
	"fmt"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/compare"
)

// DocumentComparator compares two Word documents.
type DocumentComparator struct {
	comparator compare.Comparator
}

// NewDocumentComparator creates a new DocumentComparator.
func NewDocumentComparator() *DocumentComparator {
	return &DocumentComparator{
		comparator: compare.NewElementComparator(),
	}
}

// Compare compares two documents and returns the differences.
func (dc *DocumentComparator) Compare(doc1, doc2 *Document) ([]compare.Diff, error) {
	if doc1 == nil || doc2 == nil {
		return nil, errors.New("cannot compare nil documents")
	}

	root1 := doc1.MainPart().Document()
	root2 := doc2.MainPart().Document()

	if root1 == nil || root2 == nil {
		return nil, errors.New("document main part is missing or invalid")
	}

	return dc.comparator.Compare(root1, root2), nil
}

// DocumentMerger merges Word documents.
type DocumentMerger struct {
	merger compare.Merger
}

// NewDocumentMerger creates a new DocumentMerger.
func NewDocumentMerger() *DocumentMerger {
	return &DocumentMerger{
		merger: compare.NewElementMerger(),
	}
}

// Merge merges 'other' into 'base'. The 'base' document is modified in place.
func (dm *DocumentMerger) Merge(base, other *Document) error {
	if base == nil || other == nil {
		return errors.New("cannot merge nil documents")
	}

	rootBase := base.MainPart().Document()
	rootOther := other.MainPart().Document()

	if rootBase == nil || rootOther == nil {
		return errors.New("document main part is missing or invalid")
	}

	return dm.merger.Merge(rootBase, rootOther)
}

// ThreeWayMerge merges 'ours' and 'theirs' into 'base'.
// Currently, this modifies 'base' in place to reflect the merge result.
// It returns the list of conflicts found.
func (dm *DocumentMerger) ThreeWayMerge(base, ours, theirs *Document) ([]compare.Conflict, error) {
	if base == nil || ours == nil || theirs == nil {
		return nil, errors.New("cannot merge nil documents")
	}

	rootBase := base.MainPart().Document()
	rootOurs := ours.MainPart().Document()
	rootTheirs := theirs.MainPart().Document()

	if rootBase == nil || rootOurs == nil || rootTheirs == nil {
		return nil, errors.New("document main part is missing or invalid")
	}

	mergedElement, conflicts, err := dm.merger.ThreeWayMerge(rootBase, rootOurs, rootTheirs)
	if err != nil {
		return nil, err
	}

	if root, ok := mergedElement.(openxml.PartRootElement); ok {
		base.MainPart().SetRootElement(root)
	} else {
		return nil, errors.New("merged element is not a PartRootElement")
	}

	return conflicts, nil
}

// SetStrategy sets the merge strategy.
func (dm *DocumentMerger) SetStrategy(strategy compare.MergeStrategy) {
	dm.merger.SetStrategy(strategy)
}

// SetOptions sets the merge options.
func (dm *DocumentMerger) SetOptions(options compare.MergeOptions) {
	dm.merger.SetOptions(options)
}
