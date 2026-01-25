package presentation

import (
	"fmt"
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

	// Find the first <table> element
	var tableNode *html.Node
	var findTable func(*html.Node)
	findTable = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			tableNode = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTable(c)
			if tableNode != nil {
				return
			}
		}
	}
	findTable(doc)

	if tableNode == nil {
		return nil, fmt.Errorf("no table found in HTML")
	}

	// Parse rows and cells
	var rows [][]*html.Node
	for r := tableNode.FirstChild; r != nil; r = r.NextSibling {
		if r.Type == html.ElementNode && (r.Data == "tr" || r.Data == "thead" || r.Data == "tbody") {
			if r.Data == "tr" {
				rows = append(rows, parseRowNodes(r))
			} else {
				// Handle thead/tbody by looking for tr inside them
				for c := r.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "tr" {
						rows = append(rows, parseRowNodes(c))
					}
				}
			}
		}
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("table has no rows")
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

	relFn := func(url string) string {
		rel, _ := slide.PackagingPart().CreateRelationship(url, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink", "External")
		if rel != nil {
			return rel.ID()
		}
		return ""
	}

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

// ImportHTMLText imports an HTML string into a text shape on the specified slide.
func (d *Document) ImportHTMLText(slideIndex int, x, y, w, h float64, htmlStr string) (*elements.Shape, error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	slide, err := d.GetSlide(slideIndex)
	if err != nil {
		return nil, err
	}

	shape := slide.Slide().GetOrCreateCommonSlideData().GetOrCreateShapeTree().AddShape()
	shape.SetPosition(int(drawingml.PointsToEmu(x)), int(drawingml.PointsToEmu(y)))
	shape.SetSize(int(drawingml.PointsToEmu(w)), int(drawingml.PointsToEmu(h)))

	relFn := func(url string) string {
		rel, _ := slide.PackagingPart().CreateRelationship(url, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink", "External")
		if rel != nil {
			return rel.ID()
		}
		return ""
	}

	applyHTMLToTextBody(shape.GetOrCreateTextBody(), doc, relFn)

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
			text := strings.TrimSpace(n.Data)
			if text != "" {
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
		case html.ElementNode:
			newBold, newItalic, newUnderline := bold, italic, underline
			newLevel := level
			newListType := listType
			newLinkId := linkId

			switch n.Data {
			case "b", "strong":
				newBold = true
			case "i", "em":
				newItalic = true
			case "u":
				newUnderline = true
			case "a":
				if relFn != nil {
					for _, attr := range n.Attr {
						if attr.Key == "href" {
							newLinkId = relFn(attr.Val)
							break
						}
					}
				}
			case "br":
				p = tb.AddParagraph("")
				if newListType != "" {
					p.SetLevel(newLevel)
					if newListType == "ul" {
						p.SetCharacterBullet("•")
					} else if newListType == "ol" {
						p.SetAutoNumberedBullet(drawingml.AutoNumArabicPeriod, 1)
					}
				}
				return
			case "ul", "ol":
				newListType = n.Data
				newLevel++
			case "li":
				p = tb.AddParagraph("")
				p.SetLevel(newLevel - 1)
				if listType == "ul" {
					p.SetCharacterBullet("•")
				} else if listType == "ol" {
					// Simplified: always start at 1 or should we track it?
					// For now let's just use arabic period.
					p.SetAutoNumberedBullet(drawingml.AutoNumArabicPeriod, 1)
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, newBold, newItalic, newUnderline, newLevel, newListType, newLinkId)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c, false, false, false, 0, "", "")
	}
}

func getText(n *html.Node) string {
	var b strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return strings.TrimSpace(b.String())
}
