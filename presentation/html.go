package presentation

import (
	"errors"
	"strings"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"golang.org/x/net/html"
)

// ImportHTMLTable imports an HTML table into the specified slide.
func (d *Document) ImportHTMLTable(slideIndex int, htmlStr string) (*Table, error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	tableNode := findTableNode(doc)
	if tableNode == nil {
		return nil, errors.New("no table found in HTML")
	}

	rows := parseHTMLRows(tableNode)
	if len(rows) == 0 {
		return nil, errors.New("table has no rows")
	}

	numRows := len(rows)
	numCols := 0
	for _, row := range rows {
		if len(row) > numCols {
			numCols = len(row)
		}
	}

	slide, err := d.GetSlide(slideIndex)
	if err != nil {
		return nil, err
	}

	table := NewTable(slide.Slide(), numRows, numCols)
	relFn := createRelFn(slide)

	for rIdx, row := range rows {
		for cIdx, cellNode := range row {
			cell, _ := table.GetCell(rIdx, cIdx)
			if cell != nil {
				applyHTMLToTextBody(cell.TextBody(), cellNode, relFn)
			}
		}
	}

	return table, nil
}

func findTableNode(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "table" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if res := findTableNode(c); res != nil {
			return res
		}
	}

	return nil
}

func parseHTMLRows(tableNode *html.Node) [][]*html.Node {
	var rows [][]*html.Node
	for r := tableNode.FirstChild; r != nil; r = r.NextSibling {
		if r.Type != html.ElementNode {
			continue
		}
		if r.Data == "tr" {
			rows = append(rows, parseRowNodes(r))

			continue
		}
		if r.Data == "thead" || r.Data == "tbody" || r.Data == "tfoot" {
			for c := r.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "tr" {
					rows = append(rows, parseRowNodes(c))
				}
			}
		}
	}

	return rows
}

func createRelFn(slide *Slide) func(string) string {
	return func(url string) string {
		rel, _ := slide.PackagingPart().CreateRelationship(url, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink", "External")
		if rel != nil {
			return rel.ID()
		}

		return ""
	}
}

// HTMLRect defines a rectangle for HTML text import.
type HTMLRect struct {
	X, Y, W, H float64
}

// ImportHTMLText imports an HTML string into a text shape on the specified slide.
func (d *Document) ImportHTMLText(slideIndex int, rect HTMLRect, htmlStr string) (*elements.Shape, error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	slide, err := d.GetSlide(slideIndex)
	if err != nil {
		return nil, err
	}

	shape := slide.Slide().GetOrCreateCommonSlideData().GetOrCreateShapeTree().AddShape()
	shape.SetPosition(int(drawingml.PointsToEmu(rect.X)), int(drawingml.PointsToEmu(rect.Y)))
	shape.SetSize(int(drawingml.PointsToEmu(rect.W)), int(drawingml.PointsToEmu(rect.H)))

	applyHTMLToTextBody(shape.GetOrCreateTextBody(), doc, createRelFn(slide))

	return shape, nil
}

func parseRowNodes(n *html.Node) []*html.Node {
	var cells []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "td" || c.Data == "th") {
			cells = append(cells, c)
		}
	}

	return cells
}

func applyHTMLToTextBody(tb *drawingml.TextBody, n *html.Node, relFn func(string) string) {
	tb.ClearParagraphs()
	p := tb.AddParagraph("")

	var f func(*html.Node, bool, bool, bool, int, string, string)
	f = func(n *html.Node, bold, italic, underline bool, level int, listType string, linkId string) {
		switch n.Type {
		case html.TextNode:
			handleTextNode(p, n, bold, italic, underline, linkId)
		case html.ElementNode:
			handleElementNode(tb, &p, n, bold, italic, underline, level, listType, linkId, relFn, f)
		default:
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c, false, false, false, 0, "", "")
	}
}

func handleTextNode(p *drawingml.TextParagraph, n *html.Node, bold, italic, underline bool, linkId string) {
	text := strings.TrimSpace(n.Data)
	if text == "" {
		return
	}
	run := p.AddRun(text)
	if bold {
		run.SetBold(true)
	}
	if italic {
		run.SetItalic(true)
	}
	if underline || linkId != "" {
		run.SetUnderline(drawingml.UnderlineSingle)
	}
	if linkId != "" {
		run.SetHyperlink(linkId)
	}
}

//nolint:revive // argument-limit: recursive HTML parsing requires state
func handleElementNode(tb *drawingml.TextBody, p **drawingml.TextParagraph, n *html.Node, bold, italic, underline bool, level int, listType string, linkId string, relFn func(string) string, f func(*html.Node, bool, bool, bool, int, string, string)) {
	newBold, newItalic, newUnderline := bold, italic, underline
	newLevel, newListType, newLinkId := level, listType, linkId

	switch n.Data {
	case "b", "strong":
		newBold = true
	case "i", "em":
		newItalic = true
	case "u":
		newUnderline = true
	case "a":
		newLinkId = getHref(n, relFn)
	case "br":
		*p = tb.AddParagraph("")
		applyListFormatting(*p, newLevel, newListType)

		return
	case "ul", "ol":
		newListType = n.Data
		newLevel++
	case "li":
		*p = tb.AddParagraph("")
		(*p).SetLevel(newLevel - 1)
		applyListFormatting(*p, 1, newListType)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c, newBold, newItalic, newUnderline, newLevel, newListType, newLinkId)
	}
}

func getHref(n *html.Node, relFn func(string) string) string {
	if relFn == nil {
		return ""
	}
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return relFn(attr.Val)
		}
	}

	return ""
}

func applyListFormatting(p *drawingml.TextParagraph, level int, listType string) {
	if listType == "" {
		return
	}
	p.SetLevel(level)
	switch listType {
	case "ul":
		p.SetCharacterBullet("•")
	case "ol":
		p.SetAutoNumberedBullet(drawingml.AutoNumArabicPeriod, 1)
	}
}
