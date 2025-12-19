package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// HeaderPart represents a header part (word/header1.xml, etc.).
type HeaderPart struct {
	*openxml.OpenXmlPartData
}

// FooterPart represents a footer part (word/footer1.xml, etc.).
type FooterPart struct {
	*openxml.OpenXmlPartData
}

// Content types and relationship types for headers and footers.
const (
	ContentTypeHeader      = "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml"
	ContentTypeFooter      = "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml"
	RelationshipTypeHeader = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"
	RelationshipTypeFooter = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"
)

// Counters for generating unique header/footer filenames
var (
	headerCounter uint64
	footerCounter uint64
)

// newHeaderPart creates a new header part.
func newHeaderPart(
	mainPart *MainPart,
) (*HeaderPart, error) {
	num := atomic.AddUint64(&headerCounter, 1)
	uri := fmt.Sprintf("/word/header%d.xml", num)

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeHeader,
		RelationshipTypeHeader,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeHeader,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	hp := &HeaderPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal header content
	hp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(hp, relID); err != nil {
		return nil, err
	}

	return hp, nil
}

// initializeContent sets up minimal header content.
func (hp *HeaderPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:hdr xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:p>
    <w:r>
      <w:t></w:t>
    </w:r>
  </w:p>
</w:hdr>`
	hp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (hp *HeaderPart) FixedContentType() string {
	return ContentTypeHeader
}

// Header returns the root Header element.
func (hp *HeaderPart) Header() *elements.Header {
	root := hp.RootElement()
	if root == nil {
		return nil
	}
	if h, ok := root.(*elements.Header); ok {
		return h
	}
	// Wrap the root element as a Header
	if pre, ok := root.(*openxml.PartRootElementBase); ok {
		return &elements.Header{
			PartRootElementBase: pre,
		}
	}
	return nil
}

// GetOrCreateHeader returns the Header element, creating if necessary.
func (hp *HeaderPart) GetOrCreateHeader() *elements.Header {
	h := hp.Header()
	if h != nil {
		return h
	}
	h = elements.NewHeader()
	hp.SetRootElement(h)
	return h
}

// GetStream returns a reader for the part content.
func (hp *HeaderPart) GetStream() io.Reader {
	return hp.OpenXmlPartData.GetStream()
}

// Ensure HeaderPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*HeaderPart)(nil)

// newFooterPart creates a new footer part.
func newFooterPart(
	mainPart *MainPart,
) (*FooterPart, error) {
	num := atomic.AddUint64(&footerCounter, 1)
	uri := fmt.Sprintf("/word/footer%d.xml", num)

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeFooter,
		RelationshipTypeFooter,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeFooter,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	fp := &FooterPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal footer content
	fp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(fp, relID); err != nil {
		return nil, err
	}

	return fp, nil
}

// initializeContent sets up minimal footer content.
func (fp *FooterPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:ftr xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:p>
    <w:r>
      <w:t></w:t>
    </w:r>
  </w:p>
</w:ftr>`
	fp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (fp *FooterPart) FixedContentType() string {
	return ContentTypeFooter
}

// Footer returns the root Footer element.
func (fp *FooterPart) Footer() *elements.Footer {
	root := fp.RootElement()
	if root == nil {
		return nil
	}
	if f, ok := root.(*elements.Footer); ok {
		return f
	}
	// Wrap the root element as a Footer
	if pre, ok := root.(*openxml.PartRootElementBase); ok {
		return &elements.Footer{
			PartRootElementBase: pre,
		}
	}
	return nil
}

// GetOrCreateFooter returns the Footer element, creating if necessary.
func (fp *FooterPart) GetOrCreateFooter() *elements.Footer {
	f := fp.Footer()
	if f != nil {
		return f
	}
	f = elements.NewFooter()
	fp.SetRootElement(f)
	return f
}

// GetStream returns a reader for the part content.
func (fp *FooterPart) GetStream() io.Reader {
	return fp.OpenXmlPartData.GetStream()
}

// Ensure FooterPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*FooterPart)(nil)

// HeaderPartFactory creates a HeaderPart from a URI and container.
func HeaderPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeHeader,
		packPart,
		container,
	)
	return &HeaderPart{
		OpenXmlPartData: partData,
	}
}

// FooterPartFactory creates a FooterPart from a URI and container.
func FooterPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeFooter,
		packPart,
		container,
	)
	return &FooterPart{
		OpenXmlPartData: partData,
	}
}

// Register the HeaderPart and FooterPart types.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeHeader,
			RelationshipType:   RelationshipTypeHeader,
			Factory:            HeaderPartFactory,
			DefaultURI:         "/word/header1.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeFooter,
			RelationshipType:   RelationshipTypeFooter,
			Factory:            FooterPartFactory,
			DefaultURI:         "/word/footer1.xml",
			IsFixedContentType: true,
		},
	)
}
