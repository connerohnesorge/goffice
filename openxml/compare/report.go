package compare

import "fmt"

// ComparisonResult holds the result of a comparison or merge operation.
type ComparisonResult struct {
	Diffs     []Diff
	Conflicts []Conflict
	Stats     Statistics
}

// Statistics holds summary statistics of changes and conflicts.
type Statistics struct {
	Added      int
	Deleted    int
	Modified   int
	Conflicts  int
	ByType     map[ConflictType]int
	ByDiffType map[DiffType]int
}

// GenerateReport generates a comparison result from diffs and conflicts.
func GenerateReport(diffs []Diff, conflicts []Conflict) *ComparisonResult {
	stats := calculateStats(diffs, conflicts)
	return &ComparisonResult{
		Diffs:     diffs,
		Conflicts: conflicts,
		Stats:     stats,
	}
}

func calculateStats(diffs []Diff, conflicts []Conflict) Statistics {
	stats := Statistics{
		ByType:     make(map[ConflictType]int),
		ByDiffType: make(map[DiffType]int),
	}

	for _, d := range diffs {
		stats.ByDiffType[d.Type]++
		countDiffStats(&d, &stats)
	}

	for _, c := range conflicts {
		stats.Conflicts++
		stats.ByType[c.Type]++
	}

	// Flatten top-level counts for easy access
	stats.Added = stats.ByDiffType[Added]
	stats.Deleted = stats.ByDiffType[Deleted]
	stats.Modified = stats.ByDiffType[Modified]

	return stats
}

func countDiffStats(d *Diff, stats *Statistics) {
	// Recursive counting if needed, but usually we just count top-level or flattened?
	// If we want total changes, we should recurse.
	// But Diff structure is recursive.
	// Let's recurse.
	for _, child := range d.ChildDiffs {
		stats.ByDiffType[child.Type]++
		countDiffStats(&child, stats)
	}
}

// String returns a simple summary string.
func (s Statistics) String() string {
	return fmt.Sprintf("Added: %d, Deleted: %d, Modified: %d, Conflicts: %d", s.Added, s.Deleted, s.Modified, s.Conflicts)
}
