package openxml

// Common Office Open XML namespace URIs.
const (
	// NamespaceWordprocessingML is the main WordprocessingML namespace.
	NamespaceWordprocessingML = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"

	// NamespaceRelationships is the OPC relationships namespace.
	NamespaceRelationships = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"

	// NamespaceDrawingML is the main DrawingML namespace.
	NamespaceDrawingML = "http://schemas.openxmlformats.org/drawingml/2006/main"

	// NamespaceDrawingMLPicture is the DrawingML picture namespace.
	NamespaceDrawingMLPicture = "http://schemas.openxmlformats.org/drawingml/2006/picture"

	// NamespaceDrawingMLWordprocessing is the DrawingML for WordprocessingML namespace.
	NamespaceDrawingMLWordprocessing = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"

	// NamespaceContentTypes is the content types namespace.
	NamespaceContentTypes = "http://schemas.openxmlformats.org/package/2006/content-types"

	// NamespaceMarkupCompatibility is the markup compatibility namespace.
	NamespaceMarkupCompatibility = "http://schemas.openxmlformats.org/markup-compatibility/2006"

	// NamespaceOfficeDocument is the Office document namespace.
	NamespaceOfficeDocument = "urn:schemas-microsoft-com:office:office"

	// NamespaceVML is the VML namespace.
	NamespaceVML = "urn:schemas-microsoft-com:vml"

	// NamespaceSpreadsheetML is the main SpreadsheetML namespace.
	NamespaceSpreadsheetML = "http://schemas.openxmlformats.org/spreadsheetml/2006/main"

	// NamespacePresentationML is the main PresentationML namespace.
	NamespacePresentationML = "http://schemas.openxmlformats.org/presentationml/2006/main"

	// NamespaceDublinCore is the Dublin Core namespace for core properties.
	NamespaceDublinCore = "http://purl.org/dc/elements/1.1/"

	// NamespaceDublinCoreTerms is the Dublin Core Terms namespace.
	NamespaceDublinCoreTerms = "http://purl.org/dc/terms/"

	// NamespaceCoreProperties is the core properties namespace.
	NamespaceCoreProperties = "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"

	// NamespaceExtendedProperties is the extended properties namespace.
	NamespaceExtendedProperties = "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"

	// NamespaceCustomProperties is the custom properties namespace.
	NamespaceCustomProperties = "http://schemas.openxmlformats.org/officeDocument/2006/custom-properties"

	// NamespaceXMLSchema is the XML Schema namespace.
	NamespaceXMLSchema = "http://www.w3.org/2001/XMLSchema"

	// NamespaceXMLSchemaInstance is the XML Schema Instance namespace.
	NamespaceXMLSchemaInstance = "http://www.w3.org/2001/XMLSchema-instance"
)

// Common relationship types.
const (
	// RelationshipTypeOfficeDocument is the main document relationship type.
	RelationshipTypeOfficeDocument = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"

	// RelationshipTypeCoreProperties is the core properties relationship type.
	RelationshipTypeCoreProperties = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"

	// RelationshipTypeExtendedProperties is the extended properties relationship type.
	RelationshipTypeExtendedProperties = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"

	// RelationshipTypeStyles is the styles relationship type.
	RelationshipTypeStyles = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"

	// RelationshipTypeNumbering is the numbering relationship type.
	RelationshipTypeNumbering = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"

	// RelationshipTypeSettings is the settings relationship type.
	RelationshipTypeSettings = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"

	// RelationshipTypeFontTable is the font table relationship type.
	RelationshipTypeFontTable = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"

	// RelationshipTypeWebSettings is the web settings relationship type.
	RelationshipTypeWebSettings = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings"

	// RelationshipTypeTheme is the theme relationship type.
	RelationshipTypeTheme = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"

	// RelationshipTypeImage is the image relationship type.
	RelationshipTypeImage = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"

	// RelationshipTypeHyperlink is the hyperlink relationship type.
	RelationshipTypeHyperlink = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"

	// RelationshipTypeHeader is the header relationship type.
	RelationshipTypeHeader = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"

	// RelationshipTypeFooter is the footer relationship type.
	RelationshipTypeFooter = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"

	// RelationshipTypeFootnotes is the footnotes relationship type.
	RelationshipTypeFootnotes = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footnotes"

	// RelationshipTypeEndnotes is the endnotes relationship type.
	RelationshipTypeEndnotes = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/endnotes"

	// RelationshipTypeComments is the comments relationship type.
	RelationshipTypeComments = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
)

// Content types for common parts.
const (
	// ContentTypeWordprocessingMLDocument is the content type for a Word document.
	ContentTypeWordprocessingMLDocument = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"

	// ContentTypeWordprocessingMLTemplate is the content type for a Word template.
	ContentTypeWordprocessingMLTemplate = "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"

	// ContentTypeWordprocessingMLMacroEnabledDocument is the content type for a macro-enabled Word document.
	ContentTypeWordprocessingMLMacroEnabledDocument = "application/vnd.ms-word.document.macroEnabled.main+xml"

	// ContentTypeStyles is the content type for styles.
	ContentTypeStyles = "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"

	// ContentTypeNumbering is the content type for numbering.
	ContentTypeNumbering = "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"

	// ContentTypeSettings is the content type for settings.
	ContentTypeSettings = "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"

	// ContentTypeFontTable is the content type for font table.
	ContentTypeFontTable = "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"

	// ContentTypeWebSettings is the content type for web settings.
	ContentTypeWebSettings = "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"

	// ContentTypeTheme is the content type for theme.
	ContentTypeTheme = "application/vnd.openxmlformats-officedocument.theme+xml"

	// ContentTypeHeader is the content type for header.
	ContentTypeHeader = "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml"

	// ContentTypeFooter is the content type for footer.
	ContentTypeFooter = "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml"

	// ContentTypeFootnotes is the content type for footnotes.
	ContentTypeFootnotes = "application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml"

	// ContentTypeEndnotes is the content type for endnotes.
	ContentTypeEndnotes = "application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml"

	// ContentTypeComments is the content type for comments.
	ContentTypeComments = "application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml"

	// ContentTypeCoreProperties is the content type for core properties.
	ContentTypeCoreProperties = "application/vnd.openxmlformats-package.core-properties+xml"

	// ContentTypeExtendedProperties is the content type for extended properties.
	ContentTypeExtendedProperties = "application/vnd.openxmlformats-officedocument.extended-properties+xml"
)

// NamespacePrefixes maps namespace URIs to their conventional prefixes.
var NamespacePrefixes = map[string]string{
	NamespaceWordprocessingML:        "w",
	NamespaceRelationships:           "r",
	NamespaceDrawingML:               "a",
	NamespaceDrawingMLPicture:        "pic",
	NamespaceDrawingMLWordprocessing: "wp",
	NamespaceMarkupCompatibility:     "mc",
	NamespaceOfficeDocument:          "o",
	NamespaceVML:                     "v",
	NamespaceSpreadsheetML:           "x",
	NamespacePresentationML:          "p",
	NamespaceDublinCore:              "dc",
	NamespaceDublinCoreTerms:         "dcterms",
	NamespaceCoreProperties:          "cp",
	NamespaceExtendedProperties:      "ep",
	NamespaceXMLSchemaInstance:       "xsi",
}

// GetPrefixForNamespace returns the conventional prefix for a namespace URI.
func GetPrefixForNamespace(
	namespaceURI string,
) string {
	if prefix, ok := NamespacePrefixes[namespaceURI]; ok {
		return prefix
	}

	return ""
}
