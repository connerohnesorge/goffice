//nolint:revive // line-length-limit and file-length-limit: OOXML content types are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// MainPart represents the main document part (word/document.xml).
// This is the primary part containing the document's content.
type MainPart struct {
	*openxml.OpenXmlPartData

	// contentType is the specific content type for this document type
	contentType string
}

// NewMainPart creates a new main document part.
func NewMainPart(
	uri, contentType string,
	container openxml.OpenXmlPartContainer,
) (*MainPart, error) {
	// Create the underlying packaging part
	pkg := container.Package()
	if pkg == nil {
		return nil, ErrNilPackage
	}

	packPart, err := pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewDocument()
		},
	)

	mp := &MainPart{
		OpenXmlPartData: partData,
		contentType:     contentType,
	}

	// Create package-level relationship
	_, err = pkg.CreateRelationship(
		uri,
		openxml.RelationshipTypeOfficeDocument,
		"",
	)
	if err != nil {
		return nil, err
	}

	return mp, nil
}

// NewMainPartFromData wraps an existing OpenXmlPartData as a MainPart.
func NewMainPartFromData(
	data *openxml.OpenXmlPartData,
	contentType string,
) *MainPart {
	return &MainPart{
		OpenXmlPartData: data,
		contentType:     contentType,
	}
}

// FixedContentType returns the content type for this part.
func (mp *MainPart) FixedContentType() string {
	return mp.contentType
}

// InitializeContent sets up minimal document content for a new document.
func (mp *MainPart) InitializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>
    <w:p>
      <w:r>
        <w:t></w:t>
      </w:r>
    </w:p>
  </w:body>
</w:document>`
	mp.SetData([]byte(content))
}

// Document returns the root Document element.
func (mp *MainPart) Document() *elements.Document {
	root := mp.RootElement()
	if root == nil {
		return nil
	}
	if d, ok := root.(*elements.Document); ok {
		return d
	}

	return nil
}

// AddStylesPart adds a styles definitions part to this main part.
func (mp *MainPart) AddStylesPart() (*StylesPart, error) {
	return newStylesPart(mp)
}

// StylesPart returns the styles definitions part if present.
func (mp *MainPart) StylesPart() *StylesPart {
	for part := range mp.Parts() {
		if sp, ok := part.(*StylesPart); ok {
			return sp
		}
	}

	return nil
}

// AddNumberingPart adds a numbering definitions part to this main part.
func (mp *MainPart) AddNumberingPart() (*NumberingPart, error) {
	return newNumberingPart(mp)
}

// NumberingPart returns the numbering definitions part if present.
func (mp *MainPart) NumberingPart() *NumberingPart {
	for part := range mp.Parts() {
		if np, ok := part.(*NumberingPart); ok {
			return np
		}
	}

	return nil
}

// AddSettingsPart adds a document settings part to this main part.
func (mp *MainPart) AddSettingsPart() (*SettingsPart, error) {
	return newSettingsPart(mp)
}

// SettingsPart returns the document settings part if present.
func (mp *MainPart) SettingsPart() *SettingsPart {
	for part := range mp.Parts() {
		if sp, ok := part.(*SettingsPart); ok {
			return sp
		}
	}

	return nil
}

// AddWebSettingsPart adds a web settings part to this main part.
func (mp *MainPart) AddWebSettingsPart() (*WebSettingsPart, error) {
	return newWebSettingsPart(mp)
}

// WebSettingsPart returns the web settings part if present.
func (mp *MainPart) WebSettingsPart() *WebSettingsPart {
	for part := range mp.Parts() {
		if wsp, ok := part.(*WebSettingsPart); ok {
			return wsp
		}
	}

	return nil
}

// AddFontsPart adds a font table part to this main part.
func (mp *MainPart) AddFontsPart() (*FontsPart, error) {
	return newFontsPart(mp)
}

// FontsPart returns the font table part if present.
func (mp *MainPart) FontsPart() *FontsPart {
	for part := range mp.Parts() {
		if fp, ok := part.(*FontsPart); ok {
			return fp
		}
	}

	return nil
}

// AddHeaderPart adds a header part to this main part.
func (mp *MainPart) AddHeaderPart() (*HeaderPart, error) {
	return newHeaderPart(mp)
}

// HeaderParts returns all header parts.
func (mp *MainPart) HeaderParts() []*HeaderPart {
	var headers []*HeaderPart
	for part := range mp.Parts() {
		if hp, ok := part.(*HeaderPart); ok {
			headers = append(headers, hp)
		}
	}

	return headers
}

// AddFooterPart adds a footer part to this main part.
func (mp *MainPart) AddFooterPart() (*FooterPart, error) {
	return newFooterPart(mp)
}

// FooterParts returns all footer parts.
func (mp *MainPart) FooterParts() []*FooterPart {
	var footers []*FooterPart
	for part := range mp.Parts() {
		if fp, ok := part.(*FooterPart); ok {
			footers = append(footers, fp)
		}
	}

	return footers
}

// AddImagePart adds an image part with the specified type.
func (mp *MainPart) AddImagePart(
	imageType ImageType,
) (*ImagePart, error) {
	return newImagePart(mp, imageType)
}

// ImageParts returns all image parts.
func (mp *MainPart) ImageParts() []*ImagePart {
	var images []*ImagePart
	for part := range mp.Parts() {
		if ip, ok := part.(*ImagePart); ok {
			images = append(images, ip)
		}
	}

	return images
}

// AddFootnotesPart adds a footnotes part to this main part.
func (mp *MainPart) AddFootnotesPart() (*FootnotesPart, error) {
	return newFootnotesPart(mp)
}

// FootnotesPart returns the footnotes part if present.
func (mp *MainPart) FootnotesPart() *FootnotesPart {
	for part := range mp.Parts() {
		if fp, ok := part.(*FootnotesPart); ok {
			return fp
		}
	}

	return nil
}

// AddEndnotesPart adds an endnotes part to this main part.
func (mp *MainPart) AddEndnotesPart() (*EndnotesPart, error) {
	return newEndnotesPart(mp)
}

// EndnotesPart returns the endnotes part if present.
func (mp *MainPart) EndnotesPart() *EndnotesPart {
	for part := range mp.Parts() {
		if ep, ok := part.(*EndnotesPart); ok {
			return ep
		}
	}

	return nil
}

// AddCommentsPart adds a comments part to this main part.
func (mp *MainPart) AddCommentsPart() (*CommentsPart, error) {
	return newCommentsPart(mp)
}

// CommentsPart returns the comments part if present.
func (mp *MainPart) CommentsPart() *CommentsPart {
	for part := range mp.Parts() {
		if cp, ok := part.(*CommentsPart); ok {
			return cp
		}
	}

	return nil
}

// AddThemePart adds a theme part to this main part.
func (mp *MainPart) AddThemePart() (*ThemePart, error) {
	return newThemePart(mp)
}

// ThemePart returns the theme part if present.
func (mp *MainPart) ThemePart() *ThemePart {
	for part := range mp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (mp *MainPart) GetStream() io.Reader {
	return mp.OpenXmlPartData.GetStream()
}

// Ensure MainPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*MainPart)(nil)

// MainPartFactory creates a MainPart from a URI and container.
func MainPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Get the packaging part from the container directly without lock
	// This is called during part loading when the container is already locked
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		packPart.ContentType(),
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewDocument()
		},
	)

	return NewMainPartFromData(
		partData,
		packPart.ContentType(),
	)
}

// Register the MainPart type.
func init() {
	// Register for all Word document content types
	contentTypes := []string{
		ContentTypeDocument,
		ContentTypeTemplate,
		ContentTypeMacroEnabled,
		ContentTypeMacroTemplate,
	}

	for _, ct := range contentTypes {
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:        ct,
				RelationshipType:   openxml.RelationshipTypeOfficeDocument,
				Factory:            MainPartFactory,
				DefaultURI:         "/word/document.xml",
				IsFixedContentType: false, // Content type varies by document type
			},
		)
	}
}

// ErrNilPackage is returned when the package is nil.
var ErrNilPackage = partError("package is nil")

type partError string

func (e partError) Error() string {
	return string(e)
}

// Content types for Word document parts.
const (
	ContentTypeDocument      = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"
	ContentTypeTemplate      = "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"
	ContentTypeMacroEnabled  = "application/vnd.ms-word.document.macroEnabled.main+xml"
	ContentTypeMacroTemplate = "application/vnd.ms-word.template.macroEnabled.main+xml"
)

// AddCustomXmlPart adds a custom XML part to this main part.
func (mp *MainPart) AddCustomXmlPart() (*CustomXmlPart, error) {
	return newCustomXmlPart(mp)
}

// CustomXmlParts returns all custom XML parts.
func (mp *MainPart) CustomXmlParts() []*CustomXmlPart {
	var parts []*CustomXmlPart
	for part := range mp.Parts() {
		if cp, ok := part.(*CustomXmlPart); ok {
			parts = append(parts, cp)
		}
	}

	return parts
}

// AddVbaProjectPart adds a VBA project part to this main part.
func (mp *MainPart) AddVbaProjectPart() (*VbaProjectPart, error) {
	return newVbaProjectPart(mp)
}

// VbaProjectPart returns the VBA project part if present.
func (mp *MainPart) VbaProjectPart() *VbaProjectPart {
	for part := range mp.Parts() {
		if vp, ok := part.(*VbaProjectPart); ok {
			return vp
		}
	}

	return nil
}

// AddGlossaryPart adds a glossary document part to this main part.
func (mp *MainPart) AddGlossaryPart() (*GlossaryPart, error) {
	return newGlossaryPart(mp)
}

// GlossaryPart returns the glossary document part if present.
func (mp *MainPart) GlossaryPart() *GlossaryPart {
	for part := range mp.Parts() {
		if gp, ok := part.(*GlossaryPart); ok {
			return gp
		}
	}

	return nil
}

// Helper to add a child part with the appropriate relationship.
func (mp *MainPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	pkg := mp.Package()
	if pkg == nil {
		return nil, "", ErrNilPackage
	}

	// Create the underlying packaging part
	packPart, err := pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, "", err
	}

	// Create relationship from main document to this part
	rel, err := pkg.CreatePartRelationship(
		mp.URI(),
		uri,
		relType,
		"",
	)
	if err != nil {
		return nil, "", err
	}

	return packPart, rel.ID(), nil
}
