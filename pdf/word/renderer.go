// renderer.go provides the entry point for WordprocessingML to PDF conversion.
// It coordinates document parsing, layout, and rendering to produce a PDF file.

package word

import (
	"fmt"
	"strings"
	"time"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice-pdf/drawing"
	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/parts"
)

// DrawingInfo contains information about an inline drawing to be rendered.
type DrawingInfo struct {
	Drawing *elements.Drawing
	X, Y    float64
}

// WordRenderer handles the conversion of a WordprocessingML document to PDF.
type WordRenderer struct {
	// doc is the source Word document.
	doc *wordprocessing.Document

	// engine is the text layout engine.
	engine *layout.TextLayoutEngine

	// pdf is the target PDF document.
	pdf *core.Document

	// options contains rendering configuration.
	options RenderOptions

	// numberingState tracks list counters and formatting.
	numberingState *NumberingState

	// currentPageNumber tracks the current page number (1-based).
	currentPageNumber int

	// totalPages tracks the total number of pages (used for NUMPAGES field).
	// This is set after first pass through the document.
	totalPages int

	// bookmarks tracks bookmark positions for hyperlink navigation.
	bookmarks map[string]BookmarkPosition

	// footnotes collects footnotes to render at page bottom.
	pendingFootnotes []FootnoteInfo

	// endnotes collects endnotes to render at document end.
	collectedEndnotes []EndnoteInfo

	// hyperlinks tracks hyperlinks for PDF annotation creation.
	pendingHyperlinks []HyperlinkInfo

	// fontRegistry maps font family+style to PDF resource names
	// Key format: "FontFamily-StyleName" (e.g., "Arial-Regular", "Arial-Bold")
	fontRegistry map[string]string

	// fontCounter generates unique font resource names
	fontCounter int

	// pendingDrawings collects inline drawings to render in current paragraph
	pendingDrawings []DrawingInfo

	// currentPage tracks the current page being rendered
	currentPage *core.Page
}

// BookmarkPosition tracks the position of a bookmark in the document.
type BookmarkPosition struct {
	PageNumber int
	X, Y       float64
}

// FootnoteInfo contains information about a footnote to be rendered.
type FootnoteInfo struct {
	ID      int
	Content *elements.Footnote
}

// EndnoteInfo contains information about an endnote to be rendered.
type EndnoteInfo struct {
	ID      int
	Content *elements.Endnote
}

// HyperlinkInfo contains information about a hyperlink for PDF annotation.
type HyperlinkInfo struct {
	URL    string
	Anchor string
	Rect   core.Rectangle
	Page   *core.Page
}

// Section represents a document section with its properties and elements.
type Section struct {
	Elements []openxml.Element
	Props    *elements.SectionProperties
}

// RenderOptions configures Word document rendering behavior.
type RenderOptions struct {
	// DefaultPageSize is used if the document doesn't specify one.
	DefaultPageSize core.PageSize

	// EmbedFonts controls whether fonts are embedded in the PDF.
	EmbedFonts bool
}

// DefaultRenderOptions returns the default rendering configuration.
func DefaultRenderOptions() RenderOptions {
	return RenderOptions{
		DefaultPageSize: core.PageSizeLetter,
		EmbedFonts:      true,
	}
}

// NewWordRenderer creates a new WordRenderer for the specified document.
func NewWordRenderer(
	doc *wordprocessing.Document,
	engine *layout.TextLayoutEngine,
) (*WordRenderer, error) {
	pdf, err := core.NewDocument()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create PDF document: %w",
			err,
		)
	}

	// Get the numbering part if it exists
	var numberingPart *parts.NumberingPart
	mainPart := doc.MainPart()
	if mainPart != nil {
		numberingPart = mainPart.NumberingPart()
	}

	return &WordRenderer{
			doc:     doc,
			engine:  engine,
			pdf:     pdf,
			options: DefaultRenderOptions(),
			numberingState: NewNumberingState(
				numberingPart,
			),
			currentPageNumber: 0,
			totalPages:        0,
			bookmarks: make(
				map[string]BookmarkPosition,
			),
			pendingFootnotes:  []FootnoteInfo{},
			collectedEndnotes: []EndnoteInfo{},
			pendingHyperlinks: []HyperlinkInfo{},
			fontRegistry: make(
				map[string]string,
			),
			fontCounter:     0,
			pendingDrawings: []DrawingInfo{},
			currentPage:     nil,
		},
		nil
}

// registerStandardFonts registers the standard PDF fonts used by the renderer.
func (r *WordRenderer) registerStandardFonts(
	page *core.Page,
) {
	// Register the 4 standard fonts we use in the renderer
	// These are the built-in PDF Type1 fonts that don't need embedding
	page.RegisterFont("/F1", "Helvetica")
	page.RegisterFont("/F2", "Helvetica-Bold")
	page.RegisterFont("/F3", "Helvetica-Oblique")
	page.RegisterFont(
		"/F4",
		"Helvetica-BoldOblique",
	)
}

// getFontResourceName returns the PDF resource name for a font, registering it if needed.
// This method ensures each unique font (family + style) gets a unique resource name.
func (r *WordRenderer) getFontResourceName(
	page *core.Page,
	f *font.Font,
) string {
	if f == nil {
		// Fallback to default font if nil
		return "/F1"
	}

	// Create a unique key for this font based on family and style
	styleStr := "Regular"
	baseFontName := f.Family

	switch f.Style {
	case font.StyleBold:
		styleStr = "Bold"
		baseFontName = f.Family + "-Bold"
	case font.StyleItalic:
		styleStr = "Italic"
		baseFontName = f.Family + "-Oblique"
	case font.StyleBoldItalic:
		styleStr = "BoldItalic"
		baseFontName = f.Family + "-BoldOblique"
	}

	fontKey := f.Family + "-" + styleStr

	// Check if we've already registered this font
	if resourceName, exists := r.fontRegistry[fontKey]; exists {
		return resourceName
	}

	// Generate a new resource name
	r.fontCounter++
	resourceName := fmt.Sprintf(
		"/F%d",
		r.fontCounter,
	)

	// Register the font in the page
	page.RegisterFont(resourceName, baseFontName)

	// Store in our registry
	r.fontRegistry[fontKey] = resourceName

	return resourceName
}

// Render converts the Word document to a PDF and writes it to the specified path.
func (r *WordRenderer) Render(
	outputPath string,
) error {
	mainPart := r.doc.MainPart()
	if mainPart == nil {
		return fmt.Errorf("no main document part")
	}

	sections := r.collectSections()
	if len(sections) == 0 {
		// Ensure at least one page if document is empty
		_, _ = r.pdf.AddPage(
			r.options.DefaultPageSize,
		)
		r.currentPageNumber++
	}

	for sectionIdx, section := range sections {
		pageSize := r.options.DefaultPageSize
		if section.Props != nil {
			if pgSz := section.Props.PageSize(); pgSz != nil {
				pageSize = core.PageSize{
					Width: float64(
						pgSz.Width(),
					) / 20.0,
					Height: float64(
						pgSz.Height(),
					) / 20.0,
				}
			}
		}

		page, err := r.pdf.AddPage(pageSize)
		if err != nil {
			return err
		}

		// Register standard fonts for this page
		r.registerStandardFonts(page)

		// Increment page number
		r.currentPageNumber++

		margins := core.Margins{
			Top:    72.0,
			Bottom: 72.0,
			Left:   72.0,
			Right:  72.0,
		}

		if section.Props != nil {
			if pgMar := section.Props.PageMargins(); pgMar != nil {
				margins = core.Margins{
					Top: float64(
						pgMar.Top(),
					) / 20.0,
					Bottom: float64(
						pgMar.Bottom(),
					) / 20.0,
					Left: float64(
						pgMar.Left(),
					) / 20.0,
					Right: float64(
						pgMar.Right(),
					) / 20.0,
				}
			}
		}
		page.SetMargins(margins)

		// Render header for this page
		isFirstPageOfSection := (sectionIdx == 0 && r.currentPageNumber == 1)
		if err := r.renderHeader(page, &section, r.currentPageNumber, isFirstPageOfSection); err != nil {
			return err
		}

		numCols := 1
		colSpace := 0.0
		if section.Props != nil {
			if cols := section.Props.Columns(); cols != nil {
				numCols = cols.Num()
				colSpace = float64(
					cols.Space(),
				) / 20.0
			}
		}
		if numCols < 1 {
			numCols = 1
		}

		contentWidth := page.Width() - margins.Left - margins.Right
		columnWidth := (contentWidth - float64(numCols-1)*colSpace) / float64(
			numCols,
		)

		currentY := page.Height() - margins.Top
		currentCol := 0
		var prevParagraph *elements.Paragraph

		for i, elem := range section.Elements {
			if p, ok := elem.(*elements.Paragraph); ok {
				// Check for hard page break
				if props := p.Properties(); props != nil {
					if props.PageBreakBefore() {
						// Create new page
						page, err = r.pdf.AddPage(
							pageSize,
						)
						if err != nil {
							return err
						}
						r.registerStandardFonts(
							page,
						)
						page.SetMargins(margins)
						r.currentPageNumber++
						currentY = page.Height() - margins.Top
						currentCol = 0
					}
				}

				colX := margins.Left + float64(
					currentCol,
				)*(columnWidth+colSpace)

				// Calculate paragraph height to check if it fits
				paraHeight, lineCount := r.calculateParagraphHeight(
					p,
					columnWidth,
				)
				availableHeight := currentY - margins.Bottom

				// Check pagination constraints
				needsNewPage := false

				// Check if paragraph doesn't fit at all
				if paraHeight > availableHeight {
					needsNewPage = true
				}

				// Check keep-lines-together
				if props := p.Properties(); props != nil {
					if props.KeepLines() &&
						paraHeight > availableHeight {
						needsNewPage = true
					}

					// Check widow/orphan control (default enabled in Word)
					widowControl := true
					if elem := props.GetElement("widowControl", elements.NamespaceWML); elem != nil {
						if attr, found := elem.GetAttribute("val", elements.NamespaceWML); found {
							val := attr.Value()
							widowControl = val != "false" &&
								val != "0"
						}
					}

					if widowControl &&
						lineCount > 1 {
						// Need at least 2 lines on each page
						// If we can't fit at least 2 lines, move to next page
						if availableHeight > 0 &&
							paraHeight > availableHeight {
							// Estimate line height
							lineHeight := paraHeight / float64(
								lineCount,
							)
							if availableHeight < lineHeight*2 {
								needsNewPage = true
							}
						}
					}
				}

				// Check keep-with-next on previous paragraph
				if prevParagraph != nil &&
					!needsNewPage {
					if prevProps := prevParagraph.Properties(); prevProps != nil {
						if prevProps.KeepNext() {
							// If there's not enough space for current paragraph,
							// we should have moved previous paragraph to new page
							// For now, move current paragraph to new page
							if paraHeight > availableHeight {
								needsNewPage = true
							}
						}
					}
				}

				// Create new page if needed
				if needsNewPage {
					currentCol++
					if currentCol >= numCols {
						page, err = r.pdf.AddPage(
							pageSize,
						)
						if err != nil {
							return err
						}
						r.registerStandardFonts(
							page,
						)
						page.SetMargins(margins)
						r.currentPageNumber++
						currentCol = 0
					}
					currentY = page.Height() - margins.Top
					colX = margins.Left + float64(
						currentCol,
					)*(columnWidth+colSpace)
				}

				newY, err := r.renderParagraph(
					page,
					p,
					colX,
					currentY,
					columnWidth,
				)
				if err != nil {
					return err
				}
				currentY = newY
				prevParagraph = p

				// Check if we need a new column/page after rendering
				if currentY < margins.Bottom {
					currentCol++
					if currentCol >= numCols {
						page, err = r.pdf.AddPage(
							pageSize,
						)
						if err != nil {
							return err
						}
						r.registerStandardFonts(
							page,
						)
						page.SetMargins(margins)
						r.currentPageNumber++
						currentCol = 0
					}
					currentY = page.Height() - margins.Top
				}

				// Check for section break in paragraph properties
				if props := p.Properties(); props != nil {
					if sectPr := props.SectionProperties(); sectPr != nil {
						// Section break triggers new section
						// This is handled by collectSections, but we can break here
						if i < len(
							section.Elements,
						)-1 {
							// Force new page for next section
							page, err = r.pdf.AddPage(
								pageSize,
							)
							if err != nil {
								return err
							}
							r.registerStandardFonts(
								page,
							)
							r.currentPageNumber++
							currentY = page.Height() - margins.Top
							currentCol = 0
						}
					}
				}
			} else if elem.LocalName() == "tbl" {
				// Handle table element
				var tbl *elements.Table
				if t, ok := elem.(*elements.Table); ok {
					tbl = t
				} else if comp, ok := elem.(*openxml.CompositeElementBase); ok {
					tbl = &elements.Table{CompositeElementBase: comp}
				}

				if tbl != nil {
					colX := margins.Left + float64(
						currentCol,
					)*(columnWidth+colSpace)

					newY, err := r.renderTable(
						&page,
						tbl,
						colX,
						currentY,
						columnWidth,
						margins,
						pageSize,
					)
					if err != nil {
						return err
					}
					currentY = newY
					prevParagraph = nil

					if currentY < margins.Bottom {
						currentCol++
						if currentCol >= numCols {
							page, err = r.pdf.AddPage(
								pageSize,
							)
							if err != nil {
								return err
							}
							r.registerStandardFonts(page)
							page.SetMargins(margins)
							r.currentPageNumber++
							currentCol = 0
						}
						currentY = page.Height() - margins.Top
					}
				}
			}
		}

		// Render footer for the last page of this section
		if err := r.renderFooter(page, &section, r.currentPageNumber, isFirstPageOfSection); err != nil {
			return err
		}
	}

	// Set total pages for NUMPAGES field (for future rendering passes if needed)
	r.totalPages = r.currentPageNumber

	// Render collected endnotes at document end
	if len(r.collectedEndnotes) > 0 {
		if err := r.renderEndnotes(); err != nil {
			return err
		}
	}

	// Process hyperlinks (add PDF annotations)
	if err := r.processHyperlinks(); err != nil {
		return err
	}

	return r.pdf.WriteToFile(outputPath)
}

// collectSections groups document elements into sections based on sectPr elements.
func (r *WordRenderer) collectSections() []Section {
	var sections []Section
	var currentElements []openxml.Element

	mainPart := r.doc.MainPart()
	root := mainPart.RootElement()
	if root == nil {
		return nil
	}

	var body *elements.Body
	// The root is usually <w:document>, child is <w:body>
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}

			break
		}
	}

	if body == nil {
		return nil
	}

	for child := range body.Children() {
		if child.LocalName() == "p" {
			var p *elements.Paragraph
			if para, ok := child.(*elements.Paragraph); ok {
				p = para
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				p = &elements.Paragraph{CompositeElementBase: comp}
			}

			if p != nil {
				currentElements = append(
					currentElements,
					p,
				)
				if props := p.Properties(); props != nil {
					if sectPr := props.SectionProperties(); sectPr != nil {
						sections = append(
							sections,
							Section{
								Elements: currentElements,
								Props:    sectPr,
							},
						)
						currentElements = nil
					}
				}
			}
		} else if child.LocalName() == "tbl" {
			currentElements = append(currentElements, child)
		}
	}

	if bodySectPr := body.SectionProperties(); bodySectPr != nil {
		sections = append(sections, Section{
			Elements: currentElements,
			Props:    bodySectPr,
		})
	} else if len(currentElements) > 0 {
		sections = append(sections, Section{
			Elements: currentElements,
			Props:    nil,
		})
	} else if len(sections) == 0 {
		sections = append(sections, Section{
			Elements: nil,
			Props:    nil,
		})
	}

	return sections
}

// calculateParagraphHeight calculates the height and line count of a paragraph
// without actually rendering it.
func (r *WordRenderer) calculateParagraphHeight(
	p *elements.Paragraph,
	maxWidth float64,
) (height float64, lineCount int) {
	props := p.Properties()

	spacingBefore := 0.0
	spacingAfter := 0.0

	if props != nil {
		if sp := props.SpacingBetweenLines(); sp != nil {
			spacingBefore = float64(
				sp.Before(),
			) / 20.0
			spacingAfter = float64(
				sp.After(),
			) / 20.0
		}
	}

	// Convert to layout paragraph
	lp := r.convertToLayoutParagraph(p)

	// Perform layout
	opts := layout.DefaultLayoutOptions()
	opts.MaxWidth = maxWidth
	lines := r.engine.LayoutParagraph(lp, opts)

	// Calculate height
	paraHeight := 0.0
	for _, line := range lines {
		paraHeight += line.Height
	}

	return spacingBefore + paraHeight + spacingAfter, len(
		lines,
	)
}

// renderInlineDrawing renders an inline drawing (image) at the specified position.
func (r *WordRenderer) renderInlineDrawing(
	page *core.Page,
	drw *elements.Drawing,
	x, y float64,
) error {
	// Get the inline drawing element
	inline := drw.Inline()
	if inline == nil {
		// Try anchor drawing
		// For now, we'll skip anchor drawings (floating images)
		return nil
	}

	// Get dimensions in EMUs
	width := inline.Width()
	height := inline.Height()

	// Convert EMUs to PDF points (1 EMU = 1/914400 inch, 1 point = 1/72 inch)
	widthPt := float64(width) / 914400.0 * 72.0
	heightPt := float64(height) / 914400.0 * 72.0

	// Get the graphic element
	graphic := inline.Graphic()
	if graphic == nil {
		return nil
	}

	// Get the graphic data
	graphicData := graphic.GraphicData()
	if graphicData == nil {
		return nil
	}

	if graphicData.URI() == drawingml.NamespaceChart {
		chartRelID := ""
		for child := range graphicData.Children() {
			if child.LocalName() != "chart" ||
				child.NamespaceURI() != drawingml.NamespaceChart {
				continue
			}
			attr, found := child.GetAttribute(
				"id",
				openxml.NamespaceRelationships,
			)
			if !found {
				continue
			}
			chartRelID = attr.Value()
			break
		}

		if chartRelID != "" {
			pageSize := core.PageSize{
				Width:  page.Width(),
				Height: page.Height(),
			}
			ctx := core.NewRenderingContextFromPageSize(
				pageSize,
			)

			pageImpl := core.NewPageImpl(
				page.Width(),
				page.Height(),
			)
			pageImpl.AttachResources(r.pdf, page)
			ctx.SetPage(pageImpl)

			mainPart := r.doc.MainPart()
			if mainPart != nil {
				part, err := mainPart.GetPartById(
					chartRelID,
				)
				if err != nil {
					return nil
				}
				chartPart, ok := part.(*parts.ChartPart)
				if !ok {
					return nil
				}

				chartKind, chartData, horizontal, err := drawing.ExtractChartData(
					chartPart.ChartSpace(),
				)
				if err != nil {
					return nil
				}

				chartRenderer := drawing.NewChartRenderer(ctx)
				boundsY := y - heightPt
				switch chartKind {
				case drawing.ChartKindBar:
					_ = chartRenderer.RenderBarChart(
						x,
						boundsY,
						widthPt,
						heightPt,
						chartData,
						horizontal,
					)
				case drawing.ChartKindLine:
					_ = chartRenderer.RenderLineChart(
						x,
						boundsY,
						widthPt,
						heightPt,
						chartData,
					)
				case drawing.ChartKindPie:
					_ = chartRenderer.RenderPieChart(
						x,
						boundsY,
						widthPt,
						heightPt,
						chartData,
					)
				}

				page.WriteContent(
					[]byte(pageImpl.GetContent()),
				)

				return nil
			}
		}
	}

	// Get the picture
	picture := graphicData.Picture()
	if picture == nil {
		return nil
	}

	// Get the blip fill
	blipFill := picture.BlipFill()
	if blipFill == nil {
		return nil
	}

	// Create a rendering context for the image
	pageSize := core.PageSize{
		Width:  page.Width(),
		Height: page.Height(),
	}
	ctx := core.NewRenderingContextFromPageSize(
		pageSize,
	)
	mainPart := r.doc.MainPart()
	if mainPart != nil {
		ctx.SetImageResolver(func(id string) ([]byte, error) {
			part, err := mainPart.GetPartById(id)
			if err != nil {
				return nil, err
			}
			imagePart, ok := part.(*parts.ImagePart)
			if !ok {
				return nil, fmt.Errorf(
					"part %s is not an image",
					id,
				)
			}

			return imagePart.GetData()
		})
	}

	// Create a PageImpl adapter for the page
	pageImpl := core.NewPageImpl(
		page.Width(),
		page.Height(),
	)
	pageImpl.AttachResources(r.pdf, page)
	ctx.SetPage(pageImpl)

	// Create an image renderer
	imageRenderer := drawing.NewImageRenderer(ctx)

	// Create render bounds
	bounds := drawing.RenderBounds{
		X:      x,
		Y:      y - heightPt, // Adjust Y since we're working with PDF coordinates
		Width:  widthPt,
		Height: heightPt,
	}

	// Get the blip element
	blip := blipFill.Blip()
	if blip == nil {
		return nil
	}

	// Get the embed relationship ID
	embedID := blip.Embed()
	if embedID == "" {
		return nil
	}

	// Create a DrawingML BlipFill from the embedded ID
	// The wordprocessing BlipFill and drawingml.BlipFill have the same structure,
	// but we need to create a drawingml instance for the renderer
	drawingMLBlipFill := drawingml.NewBlipFillWithEmbed(
		embedID,
	)

	// Render the picture using the image renderer
	err := imageRenderer.RenderPicture(
		drawingMLBlipFill,
		bounds,
	)
	if err != nil {
		// If rendering fails, draw a placeholder rectangle using PageImpl
		pageImpl.SetFillColor(0.9, 0.9, 1.0)
		pageImpl.DrawRectangle(
			x,
			y-heightPt,
			widthPt,
			heightPt,
			true,
			false,
		)
	}

	// Write the accumulated content to the actual page
	page.WriteContent(
		[]byte(pageImpl.GetContent()),
	)

	return nil
}

// renderParagraph layouts and draws a single paragraph.

func (r *WordRenderer) renderParagraph(
	page *core.Page,

	p *elements.Paragraph,

	startX, startY float64,

	maxWidth float64,
) (float64, error) {
	props := p.Properties()

	spacingBefore := 0.0

	spacingAfter := 0.0

	if props != nil {
		if sp := props.SpacingBetweenLines(); sp != nil {

			spacingBefore = float64(
				sp.Before(),
			) / 20.0

			spacingAfter = float64(
				sp.After(),
			) / 20.0

		}
	}

	currentY := startY - spacingBefore

	// 1. Convert elements.Paragraph to layout.Paragraph

	lp := r.convertToLayoutParagraph(p)

	// 2. Perform layout

	opts := layout.DefaultLayoutOptions()

	opts.MaxWidth = maxWidth

	lines := r.engine.LayoutParagraph(lp, opts)

	// Calculate paragraph height

	paraHeight := 0.0

	for _, line := range lines {
		paraHeight += line.Height
	}

	// 3. Draw paragraph box (shading and borders)

	r.renderParagraphBox(
		page,
		p,
		startX,
		currentY,
		maxWidth,
		paraHeight,
	)

	// 4. Draw lines

	for _, line := range lines {

		currentY -= line.Height

		r.drawLayoutLine(
			page,
			line,
			startX+line.Indent,
			currentY,
		)

	}

	// 5. Render any pending inline drawings (images)
	if len(r.pendingDrawings) > 0 {
		// Render each drawing at the current position
		// For now, we'll render them after the paragraph text
		drawingY := currentY - spacingAfter
		for _, drawingInfo := range r.pendingDrawings {
			// Render the drawing at the paragraph position
			err := r.renderInlineDrawing(
				page,
				drawingInfo.Drawing,
				startX,
				drawingY,
			)
			if err != nil {
				return currentY - spacingAfter, err
			}

			// Get the height of the drawing to adjust position for next one
			if inline := drawingInfo.Drawing.Inline(); inline != nil {
				heightEMU := inline.Height()
				heightPt := float64(
					heightEMU,
				) / 914400.0 * 72.0
				drawingY -= heightPt
			}
		}

		// Clear pending drawings after rendering
		r.pendingDrawings = nil
	}

	return currentY - spacingAfter, nil
}

// renderParagraphBox draws the background and borders for a paragraph.
func (r *WordRenderer) renderParagraphBox(
	page *core.Page,
	p *elements.Paragraph,
	x, y, width, height float64,
) {
	props := p.Properties()
	if props == nil {
		return
	}

	if shd := props.Shading(); shd != nil {
		if fill := shd.Fill(); fill != "" &&
			fill != "auto" {
			color := drawing.ParseHex(fill)
			rect := core.Rectangle{
				LLX: x,
				LLY: y - height,
				URX: x + width,
				URY: y,
			}
			page.WriteContent(
				[]byte(
					fmt.Sprintf("%s rg %s re f\n",
						color.SetFillRGB(),
						rect.String(),
					),
				),
			)
		}
	}

	// Borders
	if borders := props.ParagraphBorders(); borders != nil {
		rect := core.Rectangle{
			LLX: x,
			LLY: y - height,
			URX: x + width,
			URY: y,
		}

		if b := borders.Top(); b != nil &&
			b.Value() != elements.BorderNone {
			r.drawBorder(
				page,
				b,
				rect.LLX,
				rect.URY,
				rect.URX,
				rect.URY,
			)
		}
		if b := borders.Bottom(); b != nil &&
			b.Value() != elements.BorderNone {
			r.drawBorder(
				page,
				b,
				rect.LLX,
				rect.LLY,
				rect.URX,
				rect.LLY,
			)
		}
		if b := borders.Left(); b != nil &&
			b.Value() != elements.BorderNone {
			r.drawBorder(
				page,
				b,
				rect.LLX,
				rect.LLY,
				rect.LLX,
				rect.URY,
			)
		}
		if b := borders.Right(); b != nil &&
			b.Value() != elements.BorderNone {
			r.drawBorder(
				page,
				b,
				rect.URX,
				rect.LLY,
				rect.URX,
				rect.URY,
			)
		}
	}
}

// drawBorder draws a single border line.
func (r *WordRenderer) drawBorder(
	page *core.Page,
	b *elements.Border,
	x1, y1, x2, y2 float64,
) {
	width := float64(
		b.Size(),
	) / 8.0 // 1/8th of a point
	if width == 0 {
		width = 0.5
	}

	color := drawing.Black
	if c := b.Color(); c != "" && c != "auto" {
		color = drawing.ParseHex(c)
	}

	page.WriteContent(
		[]byte(fmt.Sprintf(
			"q %s RG %s w %s %s m %s %s l S Q\n",
			color.SetStrokeRGB(),
			core.FormatFloat(width),
			core.FormatFloat(
				x1,
			),
			core.FormatFloat(y1),
			core.FormatFloat(
				x2,
			),
			core.FormatFloat(y2),
		)),
	)
}

// convertToLayoutParagraph maps a Word paragraph element to our layout model.
func (r *WordRenderer) convertToLayoutParagraph(
	p *elements.Paragraph,
) *layout.Paragraph {
	lp := layout.NewParagraph()

	// Check for numbering first
	var numId, numLevel int
	var hasNumbering bool

	if props := p.Properties(); props != nil {
		if numPr := props.NumberingProperties(); numPr != nil {
			numId = numPr.NumberingId()
			numLevel = numPr.NumberingLevelReference()
			if numId > 0 && numLevel >= 0 {
				hasNumbering = true
			}
		}

		if indent := props.Indentation(); indent != nil {

			lp.Properties.Indentation.Left = float64(
				indent.Left(),
			) / 20.0

			lp.Properties.Indentation.Right = float64(
				indent.Right(),
			) / 20.0

			lp.Properties.Indentation.FirstLine = float64(
				indent.FirstLine(),
			) / 20.0

			lp.Properties.Indentation.Hanging = float64(
				indent.Hanging(),
			) / 20.0
		}

		if j := props.Justification(); j != "" {
			switch j {
			case elements.JustificationCenter:

				lp.Properties.Alignment = layout.AlignCenter
			case elements.JustificationRight:

				lp.Properties.Alignment = layout.AlignRight
			case elements.JustificationBoth:

				lp.Properties.Alignment = layout.AlignJustify
			case elements.JustificationDistribute:

				lp.Properties.Alignment = layout.AlignDistribute
			default:

				lp.Properties.Alignment = layout.AlignLeft
			}
		}
	}

	// Add numbering as the first run if present
	if hasNumbering {
		numberingText := r.numberingState.GetNumberingText(
			numId,
			numLevel,
		)
		if numberingText != "" {
			// Get numbering font
			numberingFont := r.numberingState.GetNumberingFont(
				numId,
				numLevel,
			)
			if numberingFont == "" {
				numberingFont = "Arial"
			}

			// Get font for numbering
			fontSize := 12.0
			f, _ := r.engine.FontCache().
				Get(numberingFont, font.StyleRegular)
			if f == nil {
				f, _ = r.engine.FontCache().
					Get("Arial", font.StyleRegular)
				if f == nil {
					keys := r.engine.FontCache().
						Keys()
					if len(keys) > 0 {
						f, _ = r.engine.FontCache().
							Get(strings.Split(keys[0], ":")[0], font.StyleRegular)
					}
				}
			}

			// Add a tab or space after the numbering
			numberingText += "\t"

			gm := layout.NewGlyphMetrics(f)
			kt := layout.NewKerningTable(f)
			numRun := layout.NewTextRunWithKerning(
				numberingText,
				gm,
				kt,
				fontSize,
			)
			lp.AddRun(numRun)

			// Adjust indentation for numbering
			numLeft, numHanging := r.numberingState.GetNumberingIndent(
				numId,
				numLevel,
			)
			if numLeft > 0 {
				lp.Properties.Indentation.Left = numLeft
			}
			if numHanging > 0 {
				lp.Properties.Indentation.Hanging = numHanging
			}
		}
	}

	// Process paragraph children (runs, hyperlinks, bookmarks, etc.)
	for child := range p.Children() {
		switch child.LocalName() {
		case "r": // Run
			if run, ok := child.(*elements.Run); ok {
				r.processRun(run, lp)
			}

		case "hyperlink": // Hyperlink
			if hyperlink, ok := child.(*elements.Hyperlink); ok {
				r.processHyperlink(hyperlink, lp)
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				h := &elements.Hyperlink{CompositeElementBase: comp}
				r.processHyperlink(h, lp)
			}

		case "bookmarkStart": // Bookmark start
			if bs, ok := child.(*elements.BookmarkStart); ok {
				r.processBookmarkStart(bs)
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				bs := &elements.BookmarkStart{CompositeElementBase: comp}
				r.processBookmarkStart(bs)
			}

		case "bookmarkEnd": // Bookmark end (nothing to do)
			// Bookmark end markers are invisible, no action needed

		}
	}

	return lp
}

// processRun processes a run element and adds it to the layout paragraph.
func (r *WordRenderer) processRun(
	run *elements.Run,
	lp *layout.Paragraph,
) {
	// Check for special content in run
	hasSpecialContent := false
	var specialText string

	for child := range run.Children() {
		switch child.LocalName() {
		case "drawing":
			// Handle inline drawings (images)
			// Cast to Drawing element
			var drw *elements.Drawing
			if d, ok := child.(*elements.Drawing); ok {
				drw = d
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				drw = &elements.Drawing{CompositeElementBase: comp}
			}

			if drw != nil {
				// Store the drawing to be rendered later
				// We'll render it when we render the paragraph
				// For now, collect it with X,Y coordinates to be determined during rendering
				r.pendingDrawings = append(
					r.pendingDrawings,
					DrawingInfo{
						Drawing: drw,
						X:       0, // Will be set during paragraph rendering
						Y:       0, // Will be set during paragraph rendering
					},
				)
			}

			// For now, we skip adding text for images
			// TODO: Implement proper image placeholder in text flow
			return

		case "footnoteReference":
			// Footnote reference
			if fr, ok := child.(*elements.FootnoteReference); ok {
				specialText = r.processFootnoteReference(
					fr,
				)
				hasSpecialContent = true
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				fr := &elements.FootnoteReference{CompositeElementBase: comp}
				specialText = r.processFootnoteReference(fr)
				hasSpecialContent = true
			}

		case "endnoteReference":
			// Endnote reference
			if er, ok := child.(*elements.EndnoteReference); ok {
				specialText = r.processEndnoteReference(
					er,
				)
				hasSpecialContent = true
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				er := &elements.EndnoteReference{CompositeElementBase: comp}
				specialText = r.processEndnoteReference(er)
				hasSpecialContent = true
			}

		case "fldChar": // Field character (start, separate, end)
			// Handle field codes (PAGE, NUMPAGES, etc.)
			// For now, we'll skip complex field handling and rely on InnerText
			// which should contain the cached field value

		case "instrText": // Field instruction text
			// This contains the field code, we evaluate it
			// Try to get text content from the element
			var instrText string
			if textElem, ok := child.(*elements.Text); ok {
				instrText = textElem.InnerText()
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				instrText = leaf.InnerText()
			}
			if instrText != "" {
				specialText = r.evaluateFieldCode(
					instrText,
				)
				hasSpecialContent = true
			}
		}
	}

	// Determine text to render
	text := run.InnerText()
	if hasSpecialContent && specialText != "" {
		text = specialText
	}

	if text == "" {
		return
	}

	fontSize := 12.0
	if rPr := run.Properties(); rPr != nil {
		if sz := rPr.FontSize(); sz > 0 {
			fontSize = float64(sz) / 2.0
		}
	}

	f, _ := r.engine.FontCache().
		Get("Arial", font.StyleRegular)
	if f == nil {
		keys := r.engine.FontCache().Keys()
		if len(keys) > 0 {
			f, _ = r.engine.FontCache().
				Get(strings.Split(keys[0], ":")[0], font.StyleRegular)
		}
	}

	gm := layout.NewGlyphMetrics(f)
	kt := layout.NewKerningTable(f)

	lr := layout.NewTextRunWithKerning(
		text,
		gm,
		kt,
		fontSize,
	)

	if rPr := run.Properties(); rPr != nil {

		lr.Bold = rPr.Bold()

		lr.Italic = rPr.Italic()

		lr.Color = rPr.Color()
		if u := rPr.Underline(); u != elements.UnderlineNone {
			lr.Underline = string(u)
		}

		lr.Strike = rPr.Strike()
		if va := rPr.VerticalTextAlignment(); va != elements.VerticalAlignValue("baseline") {
			switch va {
			case elements.VerticalAlignValue("subscript"):
				lr.Subscript = true
			case elements.VerticalAlignValue("superscript"):
				lr.Superscript = true
			}
		}
	}

	lp.AddRun(lr)
}

// processHyperlink processes a hyperlink element and adds its runs to the layout paragraph.
func (r *WordRenderer) processHyperlink(
	h *elements.Hyperlink,
	lp *layout.Paragraph,
) {
	// Process all runs inside the hyperlink
	for child := range h.Children() {
		if child.LocalName() == "r" {
			if run, ok := child.(*elements.Run); ok {
				r.processRun(run, lp)
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				run := &elements.Run{CompositeElementBase: comp}
				r.processRun(run, lp)
			}
		}
	}

	// Store hyperlink info for later PDF annotation creation
	// We'll need to calculate the bounding box after layout
	var url, anchor string
	if h.IsExternal() {
		// Get URL from relationship
		relID := h.RelationshipId()
		if relID != "" {
			// Get the relationship target
			mainPart := r.doc.MainPart()
			if mainPart != nil {
				// Get relationship by ID
				// Note: This is simplified - in real implementation we'd need
				// to properly resolve the relationship to get the URL
				url = relID // Placeholder - would need proper relationship resolution
			}
		}
	} else if h.IsInternal() {
		anchor = h.Anchor()
	}

	if url != "" || anchor != "" {
		// We'll process hyperlinks after rendering to get proper positions
		// For now, just note that we have hyperlinks to process
	}
}

// processBookmarkStart processes a bookmark start element.
func (r *WordRenderer) processBookmarkStart(
	bs *elements.BookmarkStart,
) {
	// Store bookmark position for future reference
	// We'll update this with actual page position during rendering
	name := bs.Name()
	if name != "" {
		r.bookmarks[name] = BookmarkPosition{
			PageNumber: r.currentPageNumber,
			X:          0, // Will be updated during actual rendering
			Y:          0,
		}
	}
}

// processFootnoteReference processes a footnote reference and returns the superscript number.
func (r *WordRenderer) processFootnoteReference(
	fr *elements.FootnoteReference,
) string {
	footnoteID := fr.Id()

	// Get the footnote content
	mainPart := r.doc.MainPart()
	if mainPart != nil {
		if footnotesPart := mainPart.FootnotesPart(); footnotesPart != nil {
			if footnotes := footnotesPart.Footnotes(); footnotes != nil {
				if footnote := footnotes.GetFootnote(footnoteID); footnote != nil {
					// Collect footnote for rendering at page bottom
					r.pendingFootnotes = append(
						r.pendingFootnotes,
						FootnoteInfo{
							ID:      footnoteID,
							Content: footnote,
						},
					)
				}
			}
		}
	}

	// Return superscript number
	return fmt.Sprintf("%d", footnoteID)
}

// processEndnoteReference processes an endnote reference and returns the superscript number.
func (r *WordRenderer) processEndnoteReference(
	er *elements.EndnoteReference,
) string {
	endnoteID := er.Id()

	// Get the endnote content
	mainPart := r.doc.MainPart()
	if mainPart != nil {
		if endnotesPart := mainPart.EndnotesPart(); endnotesPart != nil {
			if endnotes := endnotesPart.Endnotes(); endnotes != nil {
				if endnote := endnotes.GetEndnote(endnoteID); endnote != nil {
					// Collect endnote for rendering at document end
					r.collectedEndnotes = append(
						r.collectedEndnotes,
						EndnoteInfo{
							ID:      endnoteID,
							Content: endnote,
						},
					)
				}
			}
		}
	}

	// Return superscript number (with Roman numeral by default in Word)
	return fmt.Sprintf("%d", endnoteID)
}

// renderEndnotes renders all collected endnotes at the end of the document.
func (r *WordRenderer) renderEndnotes() error {
	if len(r.collectedEndnotes) == 0 {
		return nil
	}

	// Add a new page for endnotes
	page, err := r.pdf.AddPage(
		r.options.DefaultPageSize,
	)
	if err != nil {
		return err
	}
	r.registerStandardFonts(page)
	r.currentPageNumber++

	margins := core.Margins{
		Top:    72.0,
		Bottom: 72.0,
		Left:   72.0,
		Right:  72.0,
	}
	page.SetMargins(margins)

	currentY := page.Height() - margins.Top
	maxWidth := page.Width() - margins.Left - margins.Right

	// Render "Endnotes" heading (could be rendered as actual text in future)
	_ = "Endnotes" // heading text
	headingSize := 14.0
	currentY -= headingSize + 12.0

	// Render each endnote
	for _, endnoteInfo := range r.collectedEndnotes {
		// Render endnote number (could be rendered as actual text in future)
		_ = fmt.Sprintf(
			"%d. ",
			endnoteInfo.ID,
		) // numText
		currentY -= 12.0

		// Render endnote content
		for para := range endnoteInfo.Content.Paragraphs() {
			newY, err := r.renderParagraph(
				page,
				para,
				margins.Left+18.0,
				currentY,
				maxWidth-18.0,
			)
			if err != nil {
				return err
			}
			currentY = newY

			// Check if we need a new page
			if currentY < margins.Bottom {
				page, err = r.pdf.AddPage(
					r.options.DefaultPageSize,
				)
				if err != nil {
					return err
				}
				r.registerStandardFonts(page)
				page.SetMargins(margins)
				r.currentPageNumber++
				currentY = page.Height() - margins.Top
			}
		}
	}

	return nil
}

// processHyperlinks creates PDF link annotations for all hyperlinks.
func (r *WordRenderer) processHyperlinks() error {
	// This is a placeholder - in a full implementation, we would:
	// 1. Track hyperlink positions during rendering
	// 2. Create PDF link annotations with proper bounding boxes
	// 3. Add GoTo actions for internal links (bookmarks)
	// 4. Add URI actions for external links
	return nil
}

// drawLayoutLine draws a positioned layout line to the PDF page.

func (r *WordRenderer) drawLayoutLine(
	page *core.Page,

	line layout.LayoutLine,

	x, y float64,
) {
	for _, pr := range line.Runs {

		tb := drawing.NewTextBuilder()

		tb.BeginText()

		// 1. Set Font

		// Resolve the font based on the run's style
		resolvedFont := r.resolveFont(pr.Run)

		// Get or register the font resource name
		fontName := r.getFontResourceName(
			page,
			resolvedFont,
		)

		tb.SetFont(fontName, pr.Run.FontSize)

		// 2. Set Color

		if pr.Run.Color != "" &&
			pr.Run.Color != "auto" {

			color := drawing.ParseHex(
				pr.Run.Color,
			)

			page.WriteContent(
				[]byte(color.SetFillRGB() + "\n"),
			)

		} else {
			page.WriteContent([]byte(drawing.Black.SetFillRGB() + "\n"))
		}

		// 3. Position and show text

		tb.SetTextPosition(x+pr.X, y)

		tb.ShowText(pr.Run.Text)

		tb.EndText()

		page.WriteContent([]byte(tb.String()))

		// 4. Handle underline and strike (drawn as lines)

		if pr.Run.Underline != "" {
			r.drawUnderline(page, pr, x, y)
		}

		if pr.Run.Strike {
			r.drawStrike(page, pr, x, y)
		}

	}
}

// resolveFont finds the appropriate font for a run based on style.

func (r *WordRenderer) resolveFont(
	run *layout.TextRun,
) *font.Font {
	style := font.StyleRegular

	if run.Bold && run.Italic {
		style = font.StyleBoldItalic
	} else if run.Bold {
		style = font.StyleBold
	} else if run.Italic {
		style = font.StyleItalic
	}

	f, _ := r.engine.FontCache().
		Get("Arial", style)

	if f == nil {
		// Try to fallback

		f, _ = r.engine.FontCache().
			Get("Arial", font.StyleRegular)
	}

	return f
}

// drawUnderline draws an underline for a run.

func (r *WordRenderer) drawUnderline(
	page *core.Page,
	pr *layout.PositionedRun,
	x, y float64,
) {
	thickness := pr.Run.FontSize / 20.0

	if thickness < 0.5 {
		thickness = 0.5
	}

	offset := pr.Run.FontSize / 10.0

	color := drawing.Black
	if pr.Run.Color != "" &&
		pr.Run.Color != "auto" {
		color = drawing.ParseHex(pr.Run.Color)
	}

	page.WriteContent(
		[]byte(fmt.Sprintf(
			"q %s RG %s w %s %s m %s %s l S Q\n",

			color.SetStrokeRGB(),

			core.FormatFloat(thickness),

			core.FormatFloat(
				x+pr.X,
			),
			core.FormatFloat(y-offset),

			core.FormatFloat(
				x+pr.X+pr.Width,
			),
			core.FormatFloat(y-offset),
		)),
	)
}

// drawStrike draws a strikethrough for a run.

func (r *WordRenderer) drawStrike(
	page *core.Page,
	pr *layout.PositionedRun,
	x, y float64,
) {
	thickness := pr.Run.FontSize / 20.0

	if thickness < 0.5 {
		thickness = 0.5
	}

	offset := pr.Run.FontSize * 0.3 // roughly middle of x-height

	color := drawing.Black
	if pr.Run.Color != "" &&
		pr.Run.Color != "auto" {
		color = drawing.ParseHex(pr.Run.Color)
	}

	page.WriteContent(
		[]byte(fmt.Sprintf(
			"q %s RG %s w %s %s m %s %s l S Q\n",

			color.SetStrokeRGB(),

			core.FormatFloat(thickness),

			core.FormatFloat(
				x+pr.X,
			),
			core.FormatFloat(y+offset),

			core.FormatFloat(
				x+pr.X+pr.Width,
			),
			core.FormatFloat(y+offset),
		)),
	)
}

// renderTable renders a complete table with all cells, borders, and content.
func (r *WordRenderer) renderTable(
	page **core.Page,
	table *elements.Table,
	x, y, maxWidth float64,
	margins core.Margins,
	pageSize core.PageSize,
) (float64, error) {
	// 1. Layout the table
	tl, err := r.LayoutTable(
		table,
		x,
		y,
		maxWidth,
	)
	if err != nil {
		return y, err
	}

	// 2. Check if we need to split across pages
	availableHeight := y - margins.Bottom
	splits := r.SplitTableAcrossPages(
		tl,
		availableHeight,
	)

	currentY := y
	for splitIdx, split := range splits {
		if splitIdx > 0 {
			// Add new page for continuation
			newPage, err := r.pdf.AddPage(
				pageSize,
			)
			if err != nil {
				return currentY, err
			}
			r.registerStandardFonts(newPage)
			newPage.SetMargins(margins)
			r.currentPageNumber++
			*page = newPage
			currentY = (*page).Height() - margins.Top
			split.Y = currentY
		}

		// Render this split
		err := r.renderTableSplit(*page, split)
		if err != nil {
			return currentY, err
		}

		currentY = split.Y - split.TotalHeight
	}

	return currentY, nil
}

// renderHeader renders the header for a page.
func (r *WordRenderer) renderHeader(
	page *core.Page,
	section *Section,
	pageNum int,
	isFirstPage bool,
) error {
	if section.Props == nil {
		return nil
	}

	// Determine which header to use
	var headerRef *elements.HeaderReference

	// Check for first page different
	if isFirstPage && section.Props.TitlePage() {
		headerRef = section.Props.GetHeaderReference(
			elements.HeaderFooterValuesFirst,
		)
	}

	// Check for odd/even different
	if headerRef == nil && pageNum%2 == 0 {
		// Even page - check if we have evenAndOddHeaders setting
		// For now, we'll just use the even header if it exists
		headerRef = section.Props.GetHeaderReference(
			elements.HeaderFooterValuesEven,
		)
	}

	// Fall back to default header
	if headerRef == nil {
		headerRef = section.Props.GetHeaderReference(
			elements.HeaderFooterDefault,
		)
	}

	if headerRef == nil {
		return nil
	}

	// Get the header part
	mainPart := r.doc.MainPart()
	if mainPart == nil {
		return nil
	}

	relID := headerRef.RelationshipId()
	if relID == "" {
		return nil
	}

	// Get the header part from the relationship
	headerPart, err := mainPart.GetPartById(relID)
	if err != nil || headerPart == nil {
		return nil
	}

	// Cast to HeaderPart
	hp, ok := headerPart.(*parts.HeaderPart)
	if !ok {
		return nil
	}

	header := hp.Header()
	if header == nil {
		return nil
	}

	// Render header content at top of page
	margins := page.Margins()
	headerY := page.Height() - margins.Top + 36.0 // Start above the top margin
	headerX := margins.Left
	headerWidth := page.Width() - margins.Left - margins.Right

	// Render each paragraph in the header
	for para := range header.Paragraphs() {
		newY, err := r.renderParagraph(
			page,
			para,
			headerX,
			headerY,
			headerWidth,
		)
		if err != nil {
			return err
		}
		headerY = newY
	}

	return nil
}

// renderFooter renders the footer for a page.
func (r *WordRenderer) renderFooter(
	page *core.Page,
	section *Section,
	pageNum int,
	isFirstPage bool,
) error {
	if section.Props == nil {
		return nil
	}

	// Determine which footer to use
	var footerRef *elements.FooterReference

	// Check for first page different
	if isFirstPage && section.Props.TitlePage() {
		footerRef = section.Props.GetFooterReference(
			elements.HeaderFooterValuesFirst,
		)
	}

	// Check for odd/even different
	if footerRef == nil && pageNum%2 == 0 {
		footerRef = section.Props.GetFooterReference(
			elements.HeaderFooterValuesEven,
		)
	}

	// Fall back to default footer
	if footerRef == nil {
		footerRef = section.Props.GetFooterReference(
			elements.HeaderFooterDefault,
		)
	}

	if footerRef == nil {
		return nil
	}

	// Get the footer part
	mainPart := r.doc.MainPart()
	if mainPart == nil {
		return nil
	}

	relID := footerRef.RelationshipId()
	if relID == "" {
		return nil
	}

	// Get the footer part from the relationship
	footerPart, err := mainPart.GetPartById(relID)
	if err != nil || footerPart == nil {
		return nil
	}

	// Cast to FooterPart
	fp, ok := footerPart.(*parts.FooterPart)
	if !ok {
		return nil
	}

	footer := fp.Footer()
	if footer == nil {
		return nil
	}

	// Render footer content at bottom of page
	margins := page.Margins()
	footerY := margins.Bottom - 36.0 // Start below the bottom margin
	footerX := margins.Left
	footerWidth := page.Width() - margins.Left - margins.Right

	// Render each paragraph in the footer
	for para := range footer.Paragraphs() {
		newY, err := r.renderParagraph(
			page,
			para,
			footerX,
			footerY,
			footerWidth,
		)
		if err != nil {
			return err
		}
		footerY = newY
	}

	return nil
}

// evaluateFieldCode evaluates a field code and returns the result text.
func (r *WordRenderer) evaluateFieldCode(
	fieldCode string,
) string {
	// Trim whitespace
	fieldCode = strings.TrimSpace(fieldCode)

	// Handle PAGE field
	if strings.HasPrefix(
		strings.ToUpper(fieldCode),
		"PAGE",
	) {
		return fmt.Sprintf(
			"%d",
			r.currentPageNumber,
		)
	}

	// Handle NUMPAGES field
	if strings.HasPrefix(
		strings.ToUpper(fieldCode),
		"NUMPAGES",
	) {
		if r.totalPages > 0 {
			return fmt.Sprintf("%d", r.totalPages)
		}

		return "0"
	}

	// Handle DATE field
	if strings.HasPrefix(
		strings.ToUpper(fieldCode),
		"DATE",
	) {
		// For simplicity, return current date
		// In a full implementation, we would parse the format from the field code
		now := time.Now()

		return now.Format("1/2/2006")
	}

	// Handle TIME field
	if strings.HasPrefix(
		strings.ToUpper(fieldCode),
		"TIME",
	) {
		// For simplicity, return current time
		now := time.Now()

		return now.Format("3:04 PM")
	}

	// Unknown field - return empty string
	return ""
}

// renderTableSplit renders a single page's worth of table rows.
func (r *WordRenderer) renderTableSplit(
	page *core.Page,
	tl *TableLayout,
) error {
	// 1. Draw table borders (if any)
	r.renderTableBorders(page, tl)

	// 2. Render each row
	for _, rowLayout := range tl.Rows {
		if err := r.renderTableRow(page, tl, rowLayout); err != nil {
			return err
		}
	}

	return nil
}

// renderTableBorders draws the outer borders of the table.
func (r *WordRenderer) renderTableBorders(
	page *core.Page,
	tl *TableLayout,
) {
	props := tl.Table.TableProperties()
	if props == nil {
		return
	}

	borders := props.TableBorders()
	if borders == nil {
		return
	}

	// Calculate table rectangle
	rect := core.Rectangle{
		LLX: tl.X,
		LLY: tl.Y - tl.TotalHeight,
		URX: tl.X + tl.TotalWidth,
		URY: tl.Y,
	}

	// Draw borders
	if b := borders.GetElement("top", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.LLX,
			rect.URY,
			rect.URX,
			rect.URY,
		)
	}
	if b := borders.GetElement("bottom", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.LLX,
			rect.LLY,
			rect.URX,
			rect.LLY,
		)
	}
	if b := borders.GetElement("left", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.LLX,
			rect.LLY,
			rect.LLX,
			rect.URY,
		)
	}
	if b := borders.GetElement("right", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.URX,
			rect.LLY,
			rect.URX,
			rect.URY,
		)
	}
}

// renderTableRow renders a single table row with all its cells.
func (r *WordRenderer) renderTableRow(
	page *core.Page,
	tl *TableLayout,
	rl RowLayout,
) error {
	for _, cellLayout := range rl.Cells {
		if err := r.renderTableCell(page, cellLayout); err != nil {
			return err
		}
	}

	return nil
}

// renderTableCell renders a single table cell with content, borders, and shading.
func (r *WordRenderer) renderTableCell(
	page *core.Page,
	cl CellLayout,
) error {
	// Skip cells that continue a vertical merge
	if cl.IsVMergeContinue {
		return nil
	}

	// 1. Draw cell shading (background)
	r.renderCellShading(page, cl)

	// 2. Draw cell borders
	r.renderCellBorders(page, cl)

	// 3. Draw cell content
	if err := r.renderCellContent(page, cl); err != nil {
		return err
	}

	return nil
}

// renderCellShading draws the background color of a cell.
func (r *WordRenderer) renderCellShading(
	page *core.Page,
	cl CellLayout,
) {
	props := cl.Cell.TableCellProperties()
	if props == nil {
		return
	}

	shd := props.Shading()
	if shd == nil {
		return
	}

	fill := shd.Fill()
	if fill == "" || fill == "auto" {
		return
	}

	color := drawing.ParseHex(fill)
	rect := core.Rectangle{
		LLX: cl.X,
		LLY: cl.Y - cl.Height,
		URX: cl.X + cl.Width,
		URY: cl.Y,
	}

	page.WriteContent(
		[]byte(
			fmt.Sprintf("%s rg %s re f\n",
				color.SetFillRGB(),
				rect.String(),
			),
		),
	)
}

// renderCellBorders draws the borders of a cell.
func (r *WordRenderer) renderCellBorders(
	page *core.Page,
	cl CellLayout,
) {
	props := cl.Cell.TableCellProperties()
	if props == nil {
		return
	}

	borders := props.TableCellBorders()
	if borders == nil {
		return
	}

	rect := core.Rectangle{
		LLX: cl.X,
		LLY: cl.Y - cl.Height,
		URX: cl.X + cl.Width,
		URY: cl.Y,
	}

	// Draw each border
	if b := borders.GetElement("top", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.LLX,
			rect.URY,
			rect.URX,
			rect.URY,
		)
	}
	if b := borders.GetElement("bottom", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.LLX,
			rect.LLY,
			rect.URX,
			rect.LLY,
		)
	}
	if b := borders.GetElement("left", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.LLX,
			rect.LLY,
			rect.LLX,
			rect.URY,
		)
	}
	if b := borders.GetElement("right", elements.NamespaceWML); b != nil {
		r.drawTableBorderLine(
			page,
			b,
			rect.URX,
			rect.LLY,
			rect.URX,
			rect.URY,
		)
	}
}

// drawTableBorderLine draws a table/cell border line from an XML element.
func (r *WordRenderer) drawTableBorderLine(
	page *core.Page,
	borderElem openxml.Element,
	x1, y1, x2, y2 float64,
) {
	// Get border attributes
	valAttr, hasVal := borderElem.GetAttribute(
		"val",
		elements.NamespaceWML,
	)
	if !hasVal || valAttr.Value() == "none" ||
		valAttr.Value() == "nil" {
		return
	}

	// Get size (in eighths of a point)
	width := 0.5
	if szAttr, found := borderElem.GetAttribute("sz", elements.NamespaceWML); found {
		var sz int
		fmt.Sscanf(szAttr.Value(), "%d", &sz)
		width = float64(sz) / 8.0
	}

	// Get color
	color := drawing.Black
	if colorAttr, found := borderElem.GetAttribute("color", elements.NamespaceWML); found {
		colorVal := colorAttr.Value()
		if colorVal != "" && colorVal != "auto" {
			color = drawing.ParseHex(colorVal)
		}
	}

	// Draw the line
	page.WriteContent(
		[]byte(fmt.Sprintf(
			"q %s RG %s w %s %s m %s %s l S Q\n",
			color.SetStrokeRGB(),
			core.FormatFloat(width),
			core.FormatFloat(
				x1,
			),
			core.FormatFloat(y1),
			core.FormatFloat(
				x2,
			),
			core.FormatFloat(y2),
		)),
	)
}

// renderCellContent renders the paragraph content within a cell.
func (r *WordRenderer) renderCellContent(
	page *core.Page,
	cl CellLayout,
) error {
	// Add padding
	padding := 3.6
	currentY := cl.Y - padding

	for _, lines := range cl.ContentLines {
		for _, line := range lines {
			currentY -= line.Height
			r.drawLayoutLine(
				page,
				line,
				cl.X+padding+line.Indent,
				currentY,
			)
		}
	}

	return nil
}
