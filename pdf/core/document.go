package core

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// Version represents a PDF version.
type Version int

const (
	// V14 represents PDF version 1.4 (Acrobat 5).
	V14 Version = iota
	// V15 represents PDF version 1.5 (Acrobat 6).
	V15
	// V16 represents PDF version 1.6 (Acrobat 7).
	V16
	// V17 represents PDF version 1.7 (ISO 32000-1:2008).
	V17
	// V20 represents PDF version 2.0 (ISO 32000-2:2017).
	V20
)

// String returns the string representation of the PDF version.
func (v Version) String() string {
	switch v {
	case V14:
		return "1.4"
	case V15:
		return "1.5"
	case V16:
		return "1.6"
	case V17:
		return "1.7"
	case V20:
		return "2.0"
	default:
		return "1.7"
	}
}

// toPdfcpuVersion converts our Version to pdfcpu's Version type.
func (v Version) toPdfcpuVersion() model.Version {
	switch v {
	case V14:
		return model.V14
	case V15:
		return model.V15
	case V16:
		return model.V16
	case V17:
		return model.V17
	case V20:
		return model.V20
	default:
		return model.V17
	}
}

// Metadata contains document metadata fields.
type Metadata struct {
	// Title is the document title.
	Title string
	// Author is the document author.
	Author string
	// Subject is the document subject.
	Subject string
	// Keywords is a comma-separated list of keywords.
	Keywords string
	// Creator is the application that created the original document.
	Creator string
	// Producer is the application that produced the PDF.
	Producer string
	// CreationDate is when the document was created.
	CreationDate time.Time
	// ModDate is when the document was last modified.
	ModDate time.Time
}

// PageSize represents the dimensions of a page in PDF points.
type PageSize struct {
	Width  float64 // Width in points (1 point = 1/72 inch).
	Height float64 // Height in points.
}

// Common page sizes in PDF points (72 points = 1 inch).
var (
	// PageSizeA4 is the ISO A4 page size (210mm x 297mm).
	PageSizeA4 = PageSize{
		Width:  595.28,
		Height: 841.89,
	}
	// PageSizeA3 is the ISO A3 page size (297mm x 420mm).
	PageSizeA3 = PageSize{
		Width:  841.89,
		Height: 1190.55,
	}
	// PageSizeLetter is the US Letter page size (8.5in x 11in).
	PageSizeLetter = PageSize{
		Width:  612,
		Height: 792,
	}
	// PageSizeLegal is the US Legal page size (8.5in x 14in).
	PageSizeLegal = PageSize{
		Width:  612,
		Height: 1008,
	}
	// PageSizeTabloid is the US Tabloid page size (11in x 17in).
	PageSizeTabloid = PageSize{
		Width:  792,
		Height: 1224,
	}
)

// Landscape returns the page size rotated to landscape orientation.
func (ps PageSize) Landscape() PageSize {
	return PageSize{
		Width:  ps.Height,
		Height: ps.Width,
	}
}

// IsLandscape returns true if the page size is in landscape orientation.
func (ps PageSize) IsLandscape() bool {
	return ps.Width > ps.Height
}

// toRect converts the PageSize to a pdfcpu Rectangle.
//
//nolint:unused // exported for future use by rendering implementations
func (ps PageSize) toRect() *types.Rectangle {
	return types.RectForDim(ps.Width, ps.Height)
}

// DocumentOptions contains options for creating a new PDF document.
type DocumentOptions struct {
	// Version is the PDF version to use. Default is V17.
	Version Version
	// Metadata contains document metadata.
	Metadata Metadata
	// CompressContent enables content stream compression. Default is true.
	CompressContent bool
}

// DefaultDocumentOptions returns the default document options.
func DefaultDocumentOptions() *DocumentOptions {
	return &DocumentOptions{
		Version:         V17,
		CompressContent: true,
		Metadata: Metadata{
			Producer:     "goffice-pdf",
			CreationDate: time.Now(),
		},
	}
}

// Document represents a PDF document that can be created and written.
// It provides a clean abstraction over pdfcpu for PDF generation.
type Document struct {
	mu       sync.RWMutex
	ctx      *model.Context
	conf     *model.Configuration
	pages    []*Page
	metadata Metadata
	version  Version
	closed   bool
}

// NewDocument creates a new empty PDF document with default options.
func NewDocument() (*Document, error) {
	return NewDocumentWithOptions(
		DefaultDocumentOptions(),
	)
}

// NewDocumentWithOptions creates a new empty PDF document with the specified options.
func NewDocumentWithOptions(
	opts *DocumentOptions,
) (*Document, error) {
	var effectiveOpts *DocumentOptions
	if opts == nil {
		effectiveOpts = DefaultDocumentOptions()
	} else {
		effectiveOpts = opts
	}

	// Create pdfcpu configuration
	conf := model.NewDefaultConfiguration()
	conf.WriteObjectStream = true
	conf.WriteXRefStream = effectiveOpts.Version >= V15

	// Create context
	dim := types.Dim{
		Width:  PageSizeLetter.Width,
		Height: PageSizeLetter.Height,
	}
	ctx, err := pdfcpu.CreateContextWithXRefTable(
		conf,
		&dim,
	)
	if err != nil {
		return nil, NewDocumentError(
			"create context",
			err,
		)
	}
	ctx.Write = model.NewWriteContext(conf.Eol)

	// Set version
	v := effectiveOpts.Version.toPdfcpuVersion()
	ctx.HeaderVersion = &v
	ctx.HeaderVersion = &v

	doc := &Document{
		ctx:      ctx,
		conf:     conf,
		pages:    make([]*Page, 0),
		metadata: effectiveOpts.Metadata,
		version:  effectiveOpts.Version,
	}

	// Set metadata in the XRefTable
	doc.applyMetadata()

	return doc, nil
}

// applyMetadata sets the document metadata in the XRefTable.
func (d *Document) applyMetadata() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.ctx.Title = d.metadata.Title
	d.ctx.Author = d.metadata.Author
	d.ctx.Subject = d.metadata.Subject
	d.ctx.Keywords = d.metadata.Keywords
	d.ctx.Creator = d.metadata.Creator
	d.ctx.Producer = d.metadata.Producer

	if !d.metadata.CreationDate.IsZero() {
		d.ctx.XRefTable.CreationDate = formatPDFDate(
			d.metadata.CreationDate,
		)
	}
	if !d.metadata.ModDate.IsZero() {
		d.ctx.ModDate = formatPDFDate(
			d.metadata.ModDate,
		)
	}
}

// formatPDFDate formats a time.Time as a PDF date string.
func formatPDFDate(t time.Time) string {
	// PDF date format: D:YYYYMMDDHHmmSSOHH'mm
	return t.Format("D:20060102150405-07'00")
}

// SetMetadata updates the document metadata.
func (d *Document) SetMetadata(
	metadata Metadata,
) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata = metadata
	d.mu.Unlock()

	d.applyMetadata()

	return nil
}

// Metadata returns the current document metadata.
func (d *Document) Metadata() Metadata {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.metadata
}

// SetTitle sets the document title.
func (d *Document) SetTitle(title string) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata.Title = title
	d.ctx.Title = title
	d.mu.Unlock()

	return nil
}

// SetAuthor sets the document author.
func (d *Document) SetAuthor(
	author string,
) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata.Author = author
	d.ctx.Author = author
	d.mu.Unlock()

	return nil
}

// SetSubject sets the document subject.
func (d *Document) SetSubject(
	subject string,
) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata.Subject = subject
	d.ctx.Subject = subject
	d.mu.Unlock()

	return nil
}

// SetKeywords sets the document keywords.
func (d *Document) SetKeywords(
	keywords string,
) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata.Keywords = keywords
	d.ctx.Keywords = keywords
	d.mu.Unlock()

	return nil
}

// SetCreator sets the document creator (the application that created the original).
func (d *Document) SetCreator(
	creator string,
) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata.Creator = creator
	d.ctx.Creator = creator
	d.mu.Unlock()

	return nil
}

// SetProducer sets the document producer (the application that produced the PDF).
func (d *Document) SetProducer(
	producer string,
) error {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()

		return ErrDocumentClosed
	}
	d.metadata.Producer = producer
	d.ctx.Producer = producer
	d.mu.Unlock()

	return nil
}

// Version returns the PDF version of the document.
func (d *Document) Version() Version {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.version
}

// PageCount returns the number of pages in the document.
func (d *Document) PageCount() int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return len(d.pages)
}

// AddPage adds a new page with the specified size and returns a Page for drawing content.
func (d *Document) AddPage(
	size PageSize,
) (*Page, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.closed {
		return nil, ErrDocumentClosed
	}

	if size.Width <= 0 || size.Height <= 0 {
		return nil, ErrInvalidPageSize
	}

	pageNum := len(d.pages) + 1
	page := newPage(pageNum, size)
	d.pages = append(d.pages, page)

	return page, nil
}

// AddPageA4 adds a new A4-sized page in portrait orientation.
func (d *Document) AddPageA4() (*Page, error) {
	return d.AddPage(PageSizeA4)
}

// AddPageLetter adds a new Letter-sized page in portrait orientation.
func (d *Document) AddPageLetter() (*Page, error) {
	return d.AddPage(PageSizeLetter)
}

// AddPageWithOptions adds a new page with the specified options.
func (d *Document) AddPageWithOptions(
	opts *PageOptions,
) (*Page, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.closed {
		return nil, ErrDocumentClosed
	}

	var effectiveOpts *PageOptions
	if opts == nil {
		effectiveOpts = DefaultPageOptions()
	} else {
		effectiveOpts = opts
	}

	if effectiveOpts.Size.Width <= 0 ||
		effectiveOpts.Size.Height <= 0 {
		return nil, ErrInvalidPageSize
	}

	pageNum := len(d.pages) + 1
	page := newPageWithOptions(
		pageNum,
		effectiveOpts,
	)
	d.pages = append(d.pages, page)

	return page, nil
}

// AddPageA3 adds a new A3-sized page in portrait orientation.
func (d *Document) AddPageA3() (*Page, error) {
	return d.AddPage(PageSizeA3)
}

// AddPageA5 adds a new A5-sized page in portrait orientation.
func (d *Document) AddPageA5() (*Page, error) {
	return d.AddPage(PageSizeA5)
}

// AddPageLegal adds a new Legal-sized page in portrait orientation.
func (d *Document) AddPageLegal() (*Page, error) {
	return d.AddPage(PageSizeLegal)
}

// AddPageTabloid adds a new Tabloid-sized page in portrait orientation.
func (d *Document) AddPageTabloid() (*Page, error) {
	return d.AddPage(PageSizeTabloid)
}

// AddPageExecutive adds a new Executive-sized page in portrait orientation.
func (d *Document) AddPageExecutive() (*Page, error) {
	return d.AddPage(PageSizeExecutive)
}

// GetPage returns the page at the specified index (0-based).
func (d *Document) GetPage(
	index int,
) (*Page, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.closed {
		return nil, ErrDocumentClosed
	}

	if index < 0 || index >= len(d.pages) {
		return nil, ErrInvalidPageIndex
	}

	return d.pages[index], nil
}

// Pages returns a copy of all pages in the document.
func (d *Document) Pages() []*Page {
	d.mu.RLock()
	defer d.mu.RUnlock()

	result := make([]*Page, len(d.pages))
	copy(result, d.pages)

	return result
}

// Write writes the PDF document to the provided writer.
func (d *Document) Write(w io.Writer) error {
	if w == nil {
		return ErrNilWriter
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	if d.closed {
		return ErrDocumentClosed
	}

	if len(d.pages) == 0 {
		return ErrNoPages
	}

	// Build the page tree
	if err := d.buildPageTree(); err != nil {
		return NewWriteError(
			"build page tree",
			err,
		)
	}

	// Update modification date
	d.ctx.ModDate = formatPDFDate(
		time.Now(),
	)

	// Write the PDF
	if err := api.WriteContext(d.ctx, w); err != nil {
		return NewWriteError("write context", err)
	}

	return nil
}

// WriteToFile writes the PDF document to the specified file path.
func (d *Document) WriteToFile(
	path string,
) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.closed {
		return ErrDocumentClosed
	}

	if len(d.pages) == 0 {
		return ErrNoPages
	}

	// Build the page tree
	if err := d.buildPageTree(); err != nil {
		return NewWriteError(
			"build page tree",
			err,
		)
	}

	// Update modification date
	d.ctx.ModDate = formatPDFDate(
		time.Now(),
	)

	// Write to file
	if err := api.WriteContextFile(d.ctx, path); err != nil {
		return NewWriteError("write file", err)
	}

	return nil
}

// buildPageTree constructs the PDF page tree from the document's pages.
func (d *Document) buildPageTree() error {
	if len(d.pages) == 0 {
		return ErrNoPages
	}

	xRefTable := d.ctx.XRefTable

	// Get the existing Pages indirect reference from the root catalog
	// CreateContextWithXRefTable already created a Pages object for us
	pagesRefObj := xRefTable.RootDict["Pages"]
	pagesIndRef, ok := pagesRefObj.(types.IndirectRef)
	if !ok {
		return NewDocumentError(
			"get pages reference",
			fmt.Errorf(
				"Pages entry in catalog is not an indirect reference",
			),
		)
	}

	// Dereference to get the Pages dictionary
	pagesObj, err := xRefTable.Dereference(
		pagesIndRef,
	)
	if err != nil {
		return NewDocumentError(
			"dereference pages",
			err,
		)
	}

	pagesDict, ok := pagesObj.(types.Dict)
	if !ok {
		return NewDocumentError(
			"get pages dictionary",
			fmt.Errorf(
				"Pages object is not a dictionary",
			),
		)
	}

	// Create each page and add to the tree
	kids := types.Array{}
	for _, page := range d.pages {
		pageIndRef, err := d.createPDFPage(
			page,
			&pagesIndRef,
		)
		if err != nil {
			return NewPageError(
				page.Number(),
				"create",
				err,
			)
		}
		kids = append(kids, *pageIndRef)
	}

	// Update the Pages dictionary with the actual page count and kids
	pagesDict.Update("Type", types.Name("Pages"))
	pagesDict.Update(
		"Count",
		types.Integer(len(d.pages)),
	)
	pagesDict.Update("Kids", kids)

	// Remove the MediaBox if it exists (Pages nodes shouldn't have MediaBox)
	pagesDict.Delete("MediaBox")

	xRefTable.PageCount = len(d.pages)

	return nil
}

// createPDFPage creates a PDF page object from our Page abstraction.
func (d *Document) createPDFPage(
	page *Page,
	parentRef *types.IndirectRef,
) (*types.IndirectRef, error) {
	xRefTable := d.ctx.XRefTable

	// Get page content
	content := page.Content()

	// Create the page dictionary with MediaBox
	mediaBox := page.MediaBox()
	pageDict := types.Dict(
		map[string]types.Object{
			"Type":     types.Name("Page"),
			"Parent":   *parentRef,
			"MediaBox": mediaBox.Array(),
		},
	)

	// Add optional page boxes if set
	if page.boxes != nil {
		if page.boxes.CropBox != nil {
			pageDict.Insert(
				"CropBox",
				page.boxes.CropBox.Array(),
			)
		}
		if page.boxes.BleedBox != nil {
			pageDict.Insert(
				"BleedBox",
				page.boxes.BleedBox.Array(),
			)
		}
		if page.boxes.TrimBox != nil {
			pageDict.Insert(
				"TrimBox",
				page.boxes.TrimBox.Array(),
			)
		}
		if page.boxes.ArtBox != nil {
			pageDict.Insert(
				"ArtBox",
				page.boxes.ArtBox.Array(),
			)
		}
	}

	// Add rotation if not zero
	if page.rotation != 0 {
		pageDict.Insert(
			"Rotate",
			types.Integer(page.rotation),
		)
	}

	// Add UserUnit if not default (PDF 1.6+)
	if page.userUnit != 1.0 && d.version >= V16 {
		pageDict.Insert(
			"UserUnit",
			types.Float(page.userUnit),
		)
	}

	// Add content stream if there is content
	if len(content) > 0 {
		streamDict, err := xRefTable.NewStreamDictForBuf(
			content,
		)
		if err != nil {
			return nil, err
		}

		// Encode the stream to populate the Raw field
		if err := streamDict.Encode(); err != nil {
			return nil, fmt.Errorf(
				"encode content stream: %w",
				err,
			)
		}

		streamIndRef, err := xRefTable.IndRefForNewObject(
			*streamDict,
		)
		if err != nil {
			return nil, err
		}

		pageDict.Insert("Contents", *streamIndRef)
	}

	// Add resources if the page has any
	resources := page.Resources()
	if len(resources) > 0 {
		pageDict.Insert("Resources", resources)
	}

	return xRefTable.IndRefForNewObject(pageDict)
}

// Close closes the document and releases resources.
// After Close is called, the document cannot be used.
func (d *Document) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.closed {
		return nil // Already closed
	}

	d.closed = true
	d.pages = nil
	d.ctx = nil
	d.conf = nil

	return nil
}

// IsClosed returns true if the document has been closed.
func (d *Document) IsClosed() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.closed
}

// Context returns the underlying pdfcpu Context.
// This is provided for advanced use cases that need direct access to pdfcpu.
// Use with caution as modifications may affect document integrity.
func (d *Document) Context() *model.Context {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.ctx
}

// Page represents a single page in a PDF document.
type Page struct {
	mu        sync.RWMutex
	number    int             // 1-based page number
	size      PageSize        // Page dimensions
	margins   Margins         // Page margins
	boxes     *PageBoxes      // Page boundary boxes (MediaBox, CropBox, etc.)
	rotation  int             // Page rotation (0, 90, 180, 270)
	userUnit  float64         // User space unit size (PDF 1.6+)
	content   *bytes.Buffer   // Content stream buffer
	resources types.Dict      // Resource dictionary
	fonts     map[string]bool // Track registered fonts
	images    map[string]bool // Track registered images
}

// newPage creates a new page with the given number and size.
func newPage(number int, size PageSize) *Page {
	return &Page{
		number:    number,
		size:      size,
		margins:   NoMargins(),
		userUnit:  1.0,
		content:   new(bytes.Buffer),
		resources: types.NewDict(),
		fonts:     make(map[string]bool),
		images:    make(map[string]bool),
	}
}

// newPageWithOptions creates a new page with the given options.
func newPageWithOptions(
	number int,
	opts *PageOptions,
) *Page {
	var effectiveOpts *PageOptions
	if opts == nil {
		effectiveOpts = DefaultPageOptions()
	} else {
		effectiveOpts = opts
	}
	p := &Page{
		number:    number,
		size:      effectiveOpts.Size,
		margins:   effectiveOpts.Margins,
		rotation:  effectiveOpts.Rotation,
		userUnit:  effectiveOpts.UserUnit,
		content:   new(bytes.Buffer),
		resources: types.NewDict(),
		fonts:     make(map[string]bool),
		images:    make(map[string]bool),
	}
	if effectiveOpts.Boxes != nil {
		p.boxes = effectiveOpts.Boxes
	}

	return p
}

// Number returns the 1-based page number.
func (p *Page) Number() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.number
}

// Size returns the page dimensions.
func (p *Page) Size() PageSize {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size
}

// Width returns the page width in points.
func (p *Page) Width() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.Width
}

// Height returns the page height in points.
func (p *Page) Height() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.Height
}

// Content returns the raw content stream bytes.
func (p *Page) Content() []byte {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.content.Bytes()
}

// ContentBuffer returns the content buffer for direct writing.
// This is provided for the drawing package to write PDF operators.
func (p *Page) ContentBuffer() *bytes.Buffer {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.content
}

// Resources returns the page resource dictionary.
func (p *Page) Resources() types.Dict {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.resources
}

// SetResources sets the page resource dictionary.
func (p *Page) SetResources(
	resources types.Dict,
) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.resources = resources
}

// AddResource adds a resource to the page's resource dictionary.
func (p *Page) AddResource(
	category, name string,
	resource types.Object,
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Get or create the category dictionary
	var categoryDict types.Dict
	if cat, ok := p.resources[category]; ok {
		if d, ok := cat.(types.Dict); ok {
			categoryDict = d
		} else {
			categoryDict = types.NewDict()
		}
	} else {
		categoryDict = types.NewDict()
	}

	categoryDict.Insert(name, resource)
	p.resources.Insert(category, categoryDict)
}

// WriteContent writes raw content stream data to the page.
// This is used by the drawing package to add PDF operators.
func (p *Page) WriteContent(
	data []byte,
) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.content.Write(data)
}

// WriteContentString writes a string to the content stream.
func (p *Page) WriteContentString(
	s string,
) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.content.WriteString(s)
}

// Margins returns the page margins.
func (p *Page) Margins() Margins {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.margins
}

// SetMargins sets the page margins.
func (p *Page) SetMargins(margins Margins) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.margins = margins
}

// Boxes returns the page boundary boxes, or nil if not set.
func (p *Page) Boxes() *PageBoxes {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.boxes
}

// SetBoxes sets the page boundary boxes.
func (p *Page) SetBoxes(boxes *PageBoxes) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.boxes = boxes
}

// MediaBox returns the media box for the page.
// If custom boxes are set, returns that MediaBox, otherwise creates one from page size.
func (p *Page) MediaBox() Rectangle {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.boxes != nil {
		return p.boxes.MediaBox
	}

	return NewRectangleForPageSize(p.size)
}

// CropBox returns the effective crop box for the page.
func (p *Page) CropBox() Rectangle {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.boxes != nil {
		return p.boxes.EffectiveCropBox()
	}

	return NewRectangleForPageSize(p.size)
}

// Rotation returns the page rotation in degrees (0, 90, 180, or 270).
func (p *Page) Rotation() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.rotation
}

// SetRotation sets the page rotation in degrees.
// Valid values are 0, 90, 180, 270. Other values are normalized.
func (p *Page) SetRotation(degrees int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	// Normalize to valid values
	d := degrees % 360
	if d < 0 {
		d += 360
	}
	// Round to nearest 90 degrees
	p.rotation = ((d + 45) / 90) * 90
	if p.rotation == 360 {
		p.rotation = 0
	}
}

// RegisterFont registers a font in the page's resource dictionary.
// This adds the font to the /Font resource dictionary so it can be referenced
// in the content stream using the specified name (e.g., "/F1").
// For built-in PDF fonts (Type1), use font names like "Helvetica", "Times-Roman", etc.
func (p *Page) RegisterFont(
	name, baseFont string,
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Ensure /Font dictionary exists in resources
	if p.resources == nil {
		p.resources = types.NewDict()
	}

	fontDict, hasFonts := p.resources["Font"]
	if !hasFonts {
		fontDict = types.NewDict()
		p.resources["Font"] = fontDict
	}

	fontDictTyped, ok := fontDict.(types.Dict)
	if !ok {
		fontDictTyped = types.NewDict()
		p.resources["Font"] = fontDictTyped
	}

	// Add the font to the Font dictionary
	// For Type1 fonts, we create a simple font dictionary
	fontEntry := types.Dict{
		"Type":     types.Name("Font"),
		"Subtype":  types.Name("Type1"),
		"BaseFont": types.Name(baseFont),
	}

	fontDictTyped[name] = fontEntry
	p.fonts[name] = true
}

// RegisterTrueTypeFont registers a TrueType font with embedding in the page's resource dictionary.
// This creates a Type0 composite font with CIDFont and embeds the font data.
// The fontData parameter contains the parsed font information, and glyphsUsed maps the runes that are used.
// This method is thread-safe and can be called concurrently for different pages.
func (d *Document) RegisterTrueTypeFont(
	page *Page,
	name string,
	fontData []byte,
	fontFamily string,
	glyphsUsed map[rune]bool,
) error {
	if page == nil {
		return fmt.Errorf("page is nil")
	}
	if len(fontData) == 0 {
		return fmt.Errorf("font data is empty")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	page.mu.Lock()
	defer page.mu.Unlock()

	xRefTable := d.ctx.XRefTable

	// Ensure /Font dictionary exists in page resources
	if page.resources == nil {
		page.resources = types.NewDict()
	}

	fontDict, hasFonts := page.resources["Font"]
	if !hasFonts {
		fontDict = types.NewDict()
		page.resources["Font"] = fontDict
	}

	fontDictTyped, ok := fontDict.(types.Dict)
	if !ok {
		fontDictTyped = types.NewDict()
		page.resources["Font"] = fontDictTyped
	}

	// Create the font stream object with the TrueType data
	fontStreamDict := types.NewDict()
	fontStreamDict.Insert(
		"Length",
		types.Integer(len(fontData)),
	)
	fontStreamDict.Insert(
		"Length1",
		types.Integer(len(fontData)),
	)

	fontStream := &types.StreamDict{
		Dict:    fontStreamDict,
		Content: fontData,
		Raw:     fontData,
	}

	// Encode the stream
	if err := fontStream.Encode(); err != nil {
		return fmt.Errorf(
			"failed to encode font stream: %w",
			err,
		)
	}

	// Create indirect reference for the font stream
	fontStreamRef, err := xRefTable.IndRefForNewObject(
		*fontStream,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create font stream reference: %w",
			err,
		)
	}

	// Create the FontDescriptor dictionary
	fontDescriptor := types.NewDict()
	fontDescriptor.Insert(
		"Type",
		types.Name("FontDescriptor"),
	)
	fontDescriptor.Insert(
		"FontName",
		types.Name(fontFamily),
	)
	fontDescriptor.Insert(
		"Flags",
		types.Integer(32),
	) // Symbolic font
	fontDescriptor.Insert(
		"FontBBox",
		types.NewIntegerArray(0, 0, 1000, 1000),
	)
	fontDescriptor.Insert(
		"ItalicAngle",
		types.Integer(0),
	)
	fontDescriptor.Insert(
		"Ascent",
		types.Integer(750),
	)
	fontDescriptor.Insert(
		"Descent",
		types.Integer(-250),
	)
	fontDescriptor.Insert(
		"CapHeight",
		types.Integer(700),
	)
	fontDescriptor.Insert(
		"StemV",
		types.Integer(80),
	)
	fontDescriptor.Insert(
		"FontFile2",
		*fontStreamRef,
	)

	// Create indirect reference for the font descriptor
	fontDescriptorRef, err := xRefTable.IndRefForNewObject(
		fontDescriptor,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create font descriptor reference: %w",
			err,
		)
	}

	// Build CIDToGIDMap (Identity mapping for simplicity)
	// For a proper subset, we would map CID to actual glyph IDs
	// For now, we use Identity-H encoding which assumes CID = GID

	// Create the CIDFont dictionary (Type 2 for TrueType)
	cidFont := types.NewDict()
	cidFont.Insert("Type", types.Name("Font"))
	cidFont.Insert(
		"Subtype",
		types.Name("CIDFontType2"),
	)
	cidFont.Insert(
		"BaseFont",
		types.Name(fontFamily),
	)

	// Create CIDSystemInfo
	cidSystemInfo := types.NewDict()
	cidSystemInfo.Insert(
		"Registry",
		types.StringLiteral("Adobe"),
	)
	cidSystemInfo.Insert(
		"Ordering",
		types.StringLiteral("Identity"),
	)
	cidSystemInfo.Insert(
		"Supplement",
		types.Integer(0),
	)
	cidFont.Insert("CIDSystemInfo", cidSystemInfo)

	cidFont.Insert(
		"FontDescriptor",
		*fontDescriptorRef,
	)
	cidFont.Insert(
		"DW",
		types.Integer(1000),
	) // Default width

	// Create indirect reference for the CIDFont
	cidFontRef, err := xRefTable.IndRefForNewObject(
		cidFont,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create CIDFont reference: %w",
			err,
		)
	}

	// Create the Type0 composite font dictionary
	type0Font := types.NewDict()
	type0Font.Insert("Type", types.Name("Font"))
	type0Font.Insert(
		"Subtype",
		types.Name("Type0"),
	)
	type0Font.Insert(
		"BaseFont",
		types.Name(fontFamily),
	)
	type0Font.Insert(
		"Encoding",
		types.Name("Identity-H"),
	)

	// DescendantFonts array
	descendantFonts := types.Array{*cidFontRef}
	type0Font.Insert(
		"DescendantFonts",
		descendantFonts,
	)

	// Add the font to the page's font dictionary
	fontDictTyped[name] = type0Font
	page.fonts[name] = true

	return nil
}

// UserUnit returns the user space unit size in points (PDF 1.6+).
// Default is 1.0 (1 user unit = 1 point).
func (p *Page) UserUnit() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.userUnit
}

// SetUserUnit sets the user space unit size (PDF 1.6+).
// Valid range is 1.0 to 75000.0.
func (p *Page) SetUserUnit(unit float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	u := unit
	if u < 1.0 {
		u = 1.0
	}
	if u > 75000 {
		u = 75000
	}
	p.userUnit = u
}

// ContentArea returns the rectangle representing the area inside the margins.
func (p *Page) ContentArea() Rectangle {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return Rectangle{
		LLX: p.margins.Left,
		LLY: p.margins.Bottom,
		URX: p.size.Width - p.margins.Right,
		URY: p.size.Height - p.margins.Top,
	}
}

// ContentWidth returns the width of the content area (page width minus margins).
func (p *Page) ContentWidth() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.Width - p.margins.HorizontalTotal()
}

// ContentHeight returns the height of the content area (page height minus margins).
func (p *Page) ContentHeight() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.Height - p.margins.VerticalTotal()
}

// CoordinateSystem returns a coordinate system for this page.
func (p *Page) CoordinateSystem() *CoordinateSystem {
	p.mu.RLock()
	defer p.mu.RUnlock()
	cs := NewCoordinateSystem(p.size.Height)
	cs.UserUnit = p.userUnit

	return cs
}

// IsLandscape returns true if the page is in landscape orientation.
func (p *Page) IsLandscape() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.IsLandscape()
}

// IsPortrait returns true if the page is in portrait orientation.
func (p *Page) IsPortrait() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.IsPortrait()
}

// Orientation returns the page orientation.
func (p *Page) Orientation() Orientation {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.size.Orientation()
}
