package compare

import (
	"fmt"
	"html"
	"strings"
)

// ToHTML generates a simple HTML report of the comparison result.
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
	for _, d := range r.Diffs {
		writeDiffHTML(&sb, d)
	}
	sb.WriteString("</ul>")

	sb.WriteString("</body></html>")
	return sb.String()
}

func writeDiffHTML(sb *strings.Builder, d Diff) {
	class := ""
	switch d.Type {
	case Added:
		class = "added"
	case Deleted:
		class = "deleted"
	case Modified:
		class = "modified"
	}

	sb.WriteString(fmt.Sprintf("<li class='%s'>%s", class, html.EscapeString(d.Message)))
	if d.Key != "" {
		sb.WriteString(fmt.Sprintf(" (Key: %s)", html.EscapeString(d.Key)))
	}
	if len(d.ChildDiffs) > 0 {
		sb.WriteString("<ul>")
		for _, child := range d.ChildDiffs {
			writeDiffHTML(sb, child)
		}
		sb.WriteString("</ul>")
	}
	sb.WriteString("</li>")
}
