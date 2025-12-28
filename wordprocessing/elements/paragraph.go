//nolint:revive // file-length-limit: paragraph element with comprehensive functionality
package elements

import (
	"iter"
	"strings"

	"github.com/connerohnesorge/goffice/openxml"
)

// Paragraph represents a paragraph element (w:p).
type Paragraph struct {
	*openxml.CompositeElementBase
}

// NewParagraph creates a new Paragraph element with optional initial text.
func NewParagraph(text ...string) *Paragraph {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"p",
		PrefixW,
	)
	p := &Paragraph{CompositeElementBase: elem}
	for _, t := range text {
		if t != "" {
			p.AppendChild(NewRun(t))
		}
	}

	return p
}

// Properties returns the paragraph properties element, or nil if not present.
func (p *Paragraph) Properties() *ParagraphProperties {
	elem := p.GetElement("pPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pp, ok := elem.(*ParagraphProperties); ok {
		return pp
	}
	// Wrap existing element
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ParagraphProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateProperties returns the paragraph properties, creating if needed.
func (p *Paragraph) GetOrCreateProperties() *ParagraphProperties {
	props := p.Properties()
	if props != nil {
		return props
	}
	props = NewParagraphProperties()
	// Properties should be first child
	if first := p.FirstChild(); first != nil {
		p.InsertBefore(props, first)
	} else {
		p.AppendChild(props)
	}

	return props
}

// Runs returns an iterator over all Run elements in this paragraph.
func (p *Paragraph) Runs() iter.Seq[*Run] {
	return func(yield func(*Run) bool) {
		for child := range p.Children() {
			if child.LocalName() != "r" ||
				child.NamespaceURI() != NamespaceWML {
				continue
			}
			var r *Run
			switch v := child.(type) {
			case *Run:
				r = v
			case *openxml.CompositeElementBase:
				r = &Run{CompositeElementBase: v}
			}
			if r != nil && !yield(r) {
				return
			}
		}
	}
}

// InnerText returns the concatenated text content of all runs.
func (p *Paragraph) InnerText() string {
	var sb strings.Builder
	for r := range p.Runs() {
		sb.WriteString(r.InnerText())
	}

	return sb.String()
}

// AppendRun appends a new Run element with the given text.
func (p *Paragraph) AppendRun(text string) *Run {
	r := NewRun(text)
	p.AppendChild(r)

	return r
}

// PrependRun prepends a new Run element with the given text.
func (p *Paragraph) PrependRun(text string) *Run {
	r := NewRun(text)
	// Find first run or end of properties
	var insertBefore openxml.Element
	for child := range p.Children() {
		if child.LocalName() != "pPr" ||
			child.NamespaceURI() != NamespaceWML {
			insertBefore = child

			break
		}
	}
	if insertBefore != nil {
		p.InsertBefore(r, insertBefore)
	} else {
		p.AppendChild(r)
	}

	return r
}

// SetText replaces all text content with the given text.
func (p *Paragraph) SetText(text string) {
	// Remove all existing runs
	var runsToRemove []openxml.Element
	for child := range p.Children() {
		if child.LocalName() == "r" &&
			child.NamespaceURI() == NamespaceWML {
			runsToRemove = append(
				runsToRemove,
				child,
			)
		}
	}
	for _, r := range runsToRemove {
		p.RemoveChild(r)
	}
	// Add new run if text is not empty
	if text != "" {
		p.AppendRun(text)
	}
}

// ClearRuns removes all run elements from the paragraph.
func (p *Paragraph) ClearRuns() {
	var runsToRemove []openxml.Element
	for child := range p.Children() {
		if child.LocalName() == "r" &&
			child.NamespaceURI() == NamespaceWML {
			runsToRemove = append(
				runsToRemove,
				child,
			)
		}
	}
	for _, r := range runsToRemove {
		p.RemoveChild(r)
	}
}

// AppendBreak appends a break element to a new run.
func (p *Paragraph) AppendBreak(
	breakType BreakType,
) *Break {
	r := NewRun("")
	br := r.AppendBreak(breakType)
	p.AppendChild(r)

	return br
}

// Clone creates a deep copy of this Paragraph element.
func (p *Paragraph) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &Paragraph{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Paragraph element.
func (p *Paragraph) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &Paragraph{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Convenience methods for paragraph properties

// SetStyle sets the paragraph style.
func (p *Paragraph) SetStyle(
	styleId string,
) *Paragraph {
	p.GetOrCreateProperties().
		SetParagraphStyleId(styleId)

	return p
}

// SetJustification sets the paragraph alignment.
func (p *Paragraph) SetJustification(
	j JustificationValue,
) *Paragraph {
	p.GetOrCreateProperties().SetJustification(j)

	return p
}

// SetKeepNext sets whether to keep with the next paragraph.
func (p *Paragraph) SetKeepNext(
	b bool,
) *Paragraph {
	p.GetOrCreateProperties().SetKeepNext(b)

	return p
}

// SetKeepLines sets whether to keep lines together.
func (p *Paragraph) SetKeepLines(
	b bool,
) *Paragraph {
	p.GetOrCreateProperties().SetKeepLines(b)

	return p
}

// SetPageBreakBefore sets whether to insert a page break before.
func (p *Paragraph) SetPageBreakBefore(
	b bool,
) *Paragraph {
	p.GetOrCreateProperties().
		SetPageBreakBefore(b)

	return p
}

// SetSpacingBefore sets the space before the paragraph in twips.
func (p *Paragraph) SetSpacingBefore(
	twips int,
) *Paragraph {
	p.GetOrCreateProperties().
		GetOrCreateSpacingBetweenLines().
		SetBefore(twips)

	return p
}

// SetSpacingAfter sets the space after the paragraph in twips.
func (p *Paragraph) SetSpacingAfter(
	twips int,
) *Paragraph {
	p.GetOrCreateProperties().
		GetOrCreateSpacingBetweenLines().
		SetAfter(twips)

	return p
}

// SetLeftIndent sets the left indentation in twips.
func (p *Paragraph) SetLeftIndent(
	twips int,
) *Paragraph {
	p.GetOrCreateProperties().
		GetOrCreateIndentation().
		SetLeft(twips)

	return p
}

// SetRightIndent sets the right indentation in twips.
func (p *Paragraph) SetRightIndent(
	twips int,
) *Paragraph {
	p.GetOrCreateProperties().
		GetOrCreateIndentation().
		SetRight(twips)

	return p
}

// SetFirstLineIndent sets the first line indentation in twips.
func (p *Paragraph) SetFirstLineIndent(
	twips int,
) *Paragraph {
	p.GetOrCreateProperties().
		GetOrCreateIndentation().
		SetFirstLine(twips)

	return p
}

// SetHangingIndent sets the hanging indentation in twips.
func (p *Paragraph) SetHangingIndent(
	twips int,
) *Paragraph {
	p.GetOrCreateProperties().
		GetOrCreateIndentation().
		SetHanging(twips)

	return p
}

// SetNumbering sets the paragraph numbering.
func (p *Paragraph) SetNumbering(
	numId, level int,
) *Paragraph {
	np := p.GetOrCreateProperties().
		GetOrCreateNumberingProperties()
	np.SetNumberingId(numId)
	np.SetNumberingLevel(level)

	return p
}

// ApplyNumbering applies numbering to the paragraph at the specified level.
// This is an alias for SetNumbering that matches the spec naming.
func (p *Paragraph) ApplyNumbering(
	numId, level int,
) *Paragraph {
	return p.SetNumbering(numId, level)
}

// RemoveNumbering removes numbering from the paragraph.
func (p *Paragraph) RemoveNumbering() *Paragraph {
	props := p.Properties()
	if props == nil {
		return p
	}
	np := props.NumberingProperties()
	if np != nil {
		props.RemoveChild(np)
	}

	return p
}

// SetNumberingLevel sets the numbering level (0-8)
// for an already numbered paragraph.
func (p *Paragraph) SetNumberingLevel(
	level int,
) *Paragraph {
	props := p.Properties()
	if props == nil {
		return p
	}
	np := props.NumberingProperties()
	if np != nil {
		np.SetNumberingLevel(level)
	}

	return p
}

// NumberingId returns the numbering ID for this paragraph,
// or 0 if not numbered.
func (p *Paragraph) NumberingId() int {
	props := p.Properties()
	if props == nil {
		return 0
	}
	np := props.NumberingProperties()
	if np == nil {
		return 0
	}

	return np.NumberingId()
}

// NumberingLevel returns the numbering level for this paragraph,
// or -1 if not numbered.
func (p *Paragraph) NumberingLevel() int {
	props := p.Properties()
	if props == nil {
		return -1
	}
	np := props.NumberingProperties()
	if np == nil {
		return -1
	}

	return np.NumberingLevelReference()
}

// IsNumbered returns whether this paragraph has numbering applied.
func (p *Paragraph) IsNumbered() bool {
	props := p.Properties()
	if props == nil {
		return false
	}

	return props.NumberingProperties() != nil
}
