//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// FootnotesPart represents the footnotes part (word/footnotes.xml).
type FootnotesPart struct {
	*openxml.OpenXmlPartData
}

// EndnotesPart represents the endnotes part (word/endnotes.xml).
type EndnotesPart struct {
	*openxml.OpenXmlPartData
}

// Content types and relationship types for footnotes and endnotes.
const (
	ContentTypeFootnotes      = "application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml"
	ContentTypeEndnotes       = "application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml"
	RelationshipTypeFootnotes = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footnotes"
	RelationshipTypeEndnotes  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/endnotes"
)

// newFootnotesPart creates a new footnotes part.
func newFootnotesPart(
	mainPart *MainPart,
) (*FootnotesPart, error) {
	uri := "/word/footnotes.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeFootnotes,
		RelationshipTypeFootnotes,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeFootnotes,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	fp := &FootnotesPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal footnotes content
	fp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(fp, relID); err != nil {
		return nil, err
	}

	return fp, nil
}

// initializeContent sets up minimal footnotes content.
func (fp *FootnotesPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:footnotes xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:footnote w:type="separator" w:id="-1">
    <w:p>
      <w:r>
        <w:separator/>
      </w:r>
    </w:p>
  </w:footnote>
  <w:footnote w:type="continuationSeparator" w:id="0">
    <w:p>
      <w:r>
        <w:continuationSeparator/>
      </w:r>
    </w:p>
  </w:footnote>
</w:footnotes>`
	fp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*FootnotesPart) FixedContentType() string {
	return ContentTypeFootnotes
}

// Footnotes returns the root Footnotes element.
func (fp *FootnotesPart) Footnotes() *elements.Footnotes {
	root := fp.RootElement()
	if root == nil {
		return nil
	}
	if fn, ok := root.(*elements.Footnotes); ok {
		return fn
	}
	// Wrap the root element as Footnotes
	if pre, ok := root.(*openxml.PartRootElementBase); ok {
		return &elements.Footnotes{
			PartRootElementBase: pre,
		}
	}

	return nil
}

// GetOrCreateFootnotes returns the Footnotes element, creating if necessary.
func (fp *FootnotesPart) GetOrCreateFootnotes() *elements.Footnotes {
	fn := fp.Footnotes()
	if fn != nil {
		return fn
	}
	fn = elements.NewFootnotes()
	fp.SetRootElement(fn)

	return fn
}

// AddFootnote adds a new footnote and returns it.
func (fp *FootnotesPart) AddFootnote(
	text string,
) *elements.Footnote {
	fn := fp.GetOrCreateFootnotes()

	return fn.AddFootnote(text)
}

// GetFootnote returns the footnote with the specified ID.
func (fp *FootnotesPart) GetFootnote(
	id int,
) *elements.Footnote {
	fn := fp.Footnotes()
	if fn == nil {
		return nil
	}

	return fn.GetFootnote(id)
}

// GetStream returns a reader for the part content.
func (fp *FootnotesPart) GetStream() io.Reader {
	return fp.OpenXmlPartData.GetStream()
}

// Ensure FootnotesPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*FootnotesPart)(nil)

// newEndnotesPart creates a new endnotes part.
func newEndnotesPart(
	mainPart *MainPart,
) (*EndnotesPart, error) {
	uri := "/word/endnotes.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeEndnotes,
		RelationshipTypeEndnotes,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeEndnotes,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	ep := &EndnotesPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal endnotes content
	ep.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(ep, relID); err != nil {
		return nil, err
	}

	return ep, nil
}

// initializeContent sets up minimal endnotes content.
func (ep *EndnotesPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:endnotes xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:endnote w:type="separator" w:id="-1">
    <w:p>
      <w:r>
        <w:separator/>
      </w:r>
    </w:p>
  </w:endnote>
  <w:endnote w:type="continuationSeparator" w:id="0">
    <w:p>
      <w:r>
        <w:continuationSeparator/>
      </w:r>
    </w:p>
  </w:endnote>
</w:endnotes>`
	ep.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*EndnotesPart) FixedContentType() string {
	return ContentTypeEndnotes
}

// Endnotes returns the root Endnotes element.
func (ep *EndnotesPart) Endnotes() *elements.Endnotes {
	root := ep.RootElement()
	if root == nil {
		return nil
	}
	if en, ok := root.(*elements.Endnotes); ok {
		return en
	}
	// Wrap the root element as Endnotes
	if pre, ok := root.(*openxml.PartRootElementBase); ok {
		return &elements.Endnotes{
			PartRootElementBase: pre,
		}
	}

	return nil
}

// GetOrCreateEndnotes returns the Endnotes element, creating if necessary.
func (ep *EndnotesPart) GetOrCreateEndnotes() *elements.Endnotes {
	en := ep.Endnotes()
	if en != nil {
		return en
	}
	en = elements.NewEndnotes()
	ep.SetRootElement(en)

	return en
}

// AddEndnote adds a new endnote and returns it.
func (ep *EndnotesPart) AddEndnote(
	text string,
) *elements.Endnote {
	en := ep.GetOrCreateEndnotes()

	return en.AddEndnote(text)
}

// GetEndnote returns the endnote with the specified ID.
func (ep *EndnotesPart) GetEndnote(
	id int,
) *elements.Endnote {
	en := ep.Endnotes()
	if en == nil {
		return nil
	}

	return en.GetEndnote(id)
}

// GetStream returns a reader for the part content.
func (ep *EndnotesPart) GetStream() io.Reader {
	return ep.OpenXmlPartData.GetStream()
}

// Ensure EndnotesPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*EndnotesPart)(nil)

// FootnotesPartFactory creates a FootnotesPart from a URI and container.
func FootnotesPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart instead of Package() to avoid locking issues
	// during initialization when the package lock is already held
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeFootnotes,
		packPart,
		container,
	)

	return &FootnotesPart{
		OpenXmlPartData: partData,
	}
}

// EndnotesPartFactory creates an EndnotesPart from a URI and container.
func EndnotesPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart instead of Package() to avoid locking issues
	// during initialization when the package lock is already held
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeEndnotes,
		packPart,
		container,
	)

	return &EndnotesPart{
		OpenXmlPartData: partData,
	}
}

// Register the FootnotesPart and EndnotesPart types.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeFootnotes,
			RelationshipType:   RelationshipTypeFootnotes,
			Factory:            FootnotesPartFactory,
			DefaultURI:         "/word/footnotes.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeEndnotes,
			RelationshipType:   RelationshipTypeEndnotes,
			Factory:            EndnotesPartFactory,
			DefaultURI:         "/word/endnotes.xml",
			IsFixedContentType: true,
		},
	)
}
