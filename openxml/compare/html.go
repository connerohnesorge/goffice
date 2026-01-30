// Package compare provides utilities for comparing and merging OpenXML elements.
// It supports diff generation, three-way merging with conflict resolution, and
// HTML/JSON reporting of comparison results.
package compare

import (
	"fmt"
	"html"
	"strings"
)

// ToHTML generates a simple HTML report of the comparison result.
// The output includes styling for different diff types (added, deleted, modified, conflict)
// and shows a summary of statistics along with detailed differences.
func (r *ComparisonResult) ToHTML() string {
	var sb strings.Builder
	sb.WriteString("<html><head><style>")
	sb.WriteString(".added { background-color: #e6ffec; }")
	sb.WriteString(".deleted { background-color: #ffebe9; }")
	sb.WriteString(".modified { background-color: #fff8c5; }")
	sb.WriteString(".conflict { background-color: #ffcecb; font-weight: bold; }")
	sb.WriteString("</style></head><body>")

	sb.WriteString("<h1>Comparison Report</h1>")

	sb.WriteString("<h2>Summary</h2>")
	sb.WriteString(fmt.Sprintf("<p>%s</p>", html.EscapeString(r.Stats.String())))

	if len(r.Conflicts) > 0 {
		sb.WriteString("<h2>Conflicts</h2><ul>")
		for _, c := range r.Conflicts {
			sb.WriteString(fmt.Sprintf("<li class='conflict'>%s: %s</li>", html.EscapeString(c.Path), html.EscapeString(c.Description)))
		}
		sb.WriteString("</ul>")
	}

	sb.WriteString("<h2>Differences</h2><ul>")
	for i := range r.Diffs {
		writeDiffHTML(&sb, &r.Diffs[i])
	}
	sb.WriteString("</ul>")

	sb.WriteString("</body></html>")

	return sb.String()
}

// writeDiffHTML writes a single diff entry as HTML to the string builder.
// It applies appropriate CSS classes based on the diff type and recursively
// handles child diffs for nested changes.
func writeDiffHTML(sb *strings.Builder, d *Diff) {
	class := ""
	switch d.Type {
	case NoDiff:
		// No styling needed for no difference
	case Added:
		class = "added"
	case Deleted:
		class = "deleted"
	case Modified:
		class = "modified"
	}

	fmt.Fprintf(sb, "<li class='%s'>%s", class, html.EscapeString(d.Message))
	if d.Key != "" {
		fmt.Fprintf(sb, " (Key: %s)", html.EscapeString(d.Key))
	}
	if len(d.ChildDiffs) > 0 {
		sb.WriteString("<ul>")
		for i := range d.ChildDiffs {
			writeDiffHTML(sb, &d.ChildDiffs[i])
		}
		sb.WriteString("</ul>")
	}
	sb.WriteString("</li>")
}
