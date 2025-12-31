//nolint:revive // Long namespace URIs cannot be broken across lines
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

	// NamespaceDrawingMLChart is the DrawingML chart namespace.
	NamespaceDrawingMLChart = "http://schemas.openxmlformats.org/drawingml/2006/chart"

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

	// NamespaceDiagram is the main Diagram namespace for SmartArt.
	NamespaceDiagram = "http://schemas.openxmlformats.org/drawingml/2006/diagram"

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

// Word extension namespaces (Office 2010-2024).
const (
	// NamespaceWord2010 is the Word 2010 extension namespace (w14).
	NamespaceWord2010 = "http://schemas.microsoft.com/office/word/2010/wordml"

	// NamespaceWord2010WordprocessingDrawing is the Word 2010 WordprocessingDrawing extension namespace (wp14).
	NamespaceWord2010WordprocessingDrawing = "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"

	// NamespaceWord2010WordprocessingCanvas is the Word 2010 WordprocessingCanvas extension namespace (wpc).
	NamespaceWord2010WordprocessingCanvas = "http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas"

	// NamespaceWord2010WordprocessingGroup is the Word 2010 WordprocessingGroup extension namespace (wpg).
	NamespaceWord2010WordprocessingGroup = "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"

	// NamespaceWord2010WordprocessingShape is the Word 2010 WordprocessingShape extension namespace (wps).
	NamespaceWord2010WordprocessingShape = "http://schemas.microsoft.com/office/word/2010/wordprocessingShape"

	// NamespaceWord2013 is the Word 2013 extension namespace (w15).
	NamespaceWord2013 = "http://schemas.microsoft.com/office/word/2012/wordml"

	// NamespaceWord2016 is the Word 2016 extension namespace (w16).
	NamespaceWord2016 = "http://schemas.microsoft.com/office/word/2015/wordml"

	// NamespaceWord2019 is the Word 2019 extension namespace (w19).
	NamespaceWord2019 = "http://schemas.microsoft.com/office/word/2018/wordml"
)

// Excel extension namespaces (Office 2010-2024).
const (
	// NamespaceExcel2009 is the Excel 2010 extension namespace (x14).
	// Note: Uses 2009 URL for historical reasons, corresponds to Office 2010.
	NamespaceExcel2009 = "http://schemas.microsoft.com/office/spreadsheetml/2009/9/main"

	// NamespaceExcel2009AC is the Excel 2010 AlternateContent extension namespace (x14ac).
	NamespaceExcel2009AC = "http://schemas.microsoft.com/office/spreadsheetml/2009/9/ac"

	// NamespaceExcel2013 is the Excel 2013 extension namespace (x15).
	NamespaceExcel2013 = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"

	// NamespaceExcel2013AC is the Excel 2013 AlternateContent extension namespace (x15ac).
	NamespaceExcel2013AC = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/ac"

	// NamespaceExcel2016 is the Excel 2016 extension namespace (x16).
	NamespaceExcel2016 = "http://schemas.microsoft.com/office/spreadsheetml/2014/11/main"

	// NamespaceExcel2016Revision is the Excel 2016 Revision extension namespace (x16r2).
	NamespaceExcel2016Revision = "http://schemas.microsoft.com/office/spreadsheetml/2015/02/main"

	// NamespaceExcel2019 is the Excel 2019 extension namespace (x19).
	NamespaceExcel2019 = "http://schemas.microsoft.com/office/spreadsheetml/2018/9/main"

	// NamespaceExcel2021 is the Excel 2021 extension namespace (x21).
	NamespaceExcel2021 = "http://schemas.microsoft.com/office/spreadsheetml/2020/10/main"

	// NamespaceExcel2024 is the Excel 2024 extension namespace (x24).
	NamespaceExcel2024 = "http://schemas.microsoft.com/office/spreadsheetml/2023/7/main"

	// NamespaceExcel2025 is the Excel 2025 extension namespace (x25).
	NamespaceExcel2025 = "http://schemas.microsoft.com/office/spreadsheetml/2024/8/main"
)

// PowerPoint extension namespaces (Office 2010-2024).
const (
	// NamespacePowerPoint2010 is the PowerPoint 2010 extension namespace (p14).
	NamespacePowerPoint2010 = "http://schemas.microsoft.com/office/powerpoint/2010/main"

	// NamespacePowerPoint2012 is the PowerPoint 2012 extension namespace (p15).
	NamespacePowerPoint2012 = "http://schemas.microsoft.com/office/powerpoint/2012/main"

	// NamespacePowerPoint2016 is the PowerPoint 2016 extension namespace (p16).
	NamespacePowerPoint2016 = "http://schemas.microsoft.com/office/powerpoint/2015/main"

	// NamespacePowerPoint2021 is the PowerPoint 2021 extension namespace (p21).
	NamespacePowerPoint2021 = "http://schemas.microsoft.com/office/powerpoint/2020/main"
)

// DrawingML extension namespaces (shared across Word, Excel, PowerPoint).
const (
	// NamespaceDrawing2010 is the DrawingML 2010 extension namespace (a14).
	NamespaceDrawing2010 = "http://schemas.microsoft.com/office/drawing/2010/main"

	// NamespaceDrawing2012 is the DrawingML 2012 extension namespace (a15).
	NamespaceDrawing2012 = "http://schemas.microsoft.com/office/drawing/2012/main"

	// NamespaceDrawing2014 is the DrawingML 2014 extension namespace (a16).
	NamespaceDrawing2014 = "http://schemas.microsoft.com/office/drawing/2014/main"

	// NamespaceDrawing2016SVG is the DrawingML 2016 SVG extension namespace (asvg).
	NamespaceDrawing2016SVG = "http://schemas.microsoft.com/office/drawing/2016/SVG/main"

	// NamespaceDrawing2016Ink is the DrawingML 2016 Ink extension namespace (aink).
	NamespaceDrawing2016Ink = "http://schemas.microsoft.com/office/drawing/2016/ink"

	// NamespaceChart2014 is the Chart 2014 extension namespace (c15).
	NamespaceChart2014 = "http://schemas.microsoft.com/office/drawing/2012/chart"

	// NamespaceChart2016 is the Chart 2016 extension namespace (c16).
	NamespaceChart2016 = "http://schemas.microsoft.com/office/drawing/2014/chart"

	// NamespaceChart2016r3 is the Chart 2016 R3 extension namespace (c16r3).
	NamespaceChart2016r3 = "http://schemas.microsoft.com/office/drawing/2017/03/chart"
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

	// RelationshipTypeAttachedTemplate is the attached template relationship type.
	RelationshipTypeAttachedTemplate = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/attachedTemplate"

	// PURL namespace relationship types (alternative namespace used by some tools)
	// These are functionally equivalent to the Microsoft namespace versions above.

	// RelationshipTypePURLOfficeDocument is the PURL namespace variant for main document.
	RelationshipTypePURLOfficeDocument = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"

	// RelationshipTypePURLExtendedProperties is the PURL namespace variant for extended properties.
	RelationshipTypePURLExtendedProperties = "http://purl.oclc.org/ooxml/officeDocument/relationships/extendedProperties"

	// RelationshipTypePURLStyles is the PURL namespace variant for styles.
	RelationshipTypePURLStyles = "http://purl.oclc.org/ooxml/officeDocument/relationships/styles"

	// RelationshipTypePURLNumbering is the PURL namespace variant for numbering.
	RelationshipTypePURLNumbering = "http://purl.oclc.org/ooxml/officeDocument/relationships/numbering"

	// RelationshipTypePURLSettings is the PURL namespace variant for settings.
	RelationshipTypePURLSettings = "http://purl.oclc.org/ooxml/officeDocument/relationships/settings"

	// RelationshipTypePURLFontTable is the PURL namespace variant for font table.
	RelationshipTypePURLFontTable = "http://purl.oclc.org/ooxml/officeDocument/relationships/fontTable"

	// RelationshipTypePURLWebSettings is the PURL namespace variant for web settings.
	RelationshipTypePURLWebSettings = "http://purl.oclc.org/ooxml/officeDocument/relationships/webSettings"

	// RelationshipTypePURLTheme is the PURL namespace variant for theme.
	RelationshipTypePURLTheme = "http://purl.oclc.org/ooxml/officeDocument/relationships/theme"

	// RelationshipTypePURLImage is the PURL namespace variant for image.
	RelationshipTypePURLImage = "http://purl.oclc.org/ooxml/officeDocument/relationships/image"

	// RelationshipTypePURLHyperlink is the PURL namespace variant for hyperlink.
	RelationshipTypePURLHyperlink = "http://purl.oclc.org/ooxml/officeDocument/relationships/hyperlink"

	// RelationshipTypePURLHeader is the PURL namespace variant for header.
	RelationshipTypePURLHeader = "http://purl.oclc.org/ooxml/officeDocument/relationships/header"

	// RelationshipTypePURLFooter is the PURL namespace variant for footer.
	RelationshipTypePURLFooter = "http://purl.oclc.org/ooxml/officeDocument/relationships/footer"

	// RelationshipTypePURLFootnotes is the PURL namespace variant for footnotes.
	RelationshipTypePURLFootnotes = "http://purl.oclc.org/ooxml/officeDocument/relationships/footnotes"

	// RelationshipTypePURLEndnotes is the PURL namespace variant for endnotes.
	RelationshipTypePURLEndnotes = "http://purl.oclc.org/ooxml/officeDocument/relationships/endnotes"

	// RelationshipTypePURLComments is the PURL namespace variant for comments.
	RelationshipTypePURLComments = "http://purl.oclc.org/ooxml/officeDocument/relationships/comments"

	// RelationshipTypePURLAttachedTemplate is the PURL namespace variant for attached template.
	RelationshipTypePURLAttachedTemplate = "http://purl.oclc.org/ooxml/officeDocument/relationships/attachedTemplate"
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
	// Main namespaces
	NamespaceWordprocessingML:        "w",
	NamespaceRelationships:           "r",
	NamespaceDrawingML:               "a",
	NamespaceDrawingMLPicture:        "pic",
	NamespaceDrawingMLWordprocessing: "wp",
	NamespaceDrawingMLChart:          "c",
	NamespaceMarkupCompatibility:     "mc",
	NamespaceOfficeDocument:          "o",
	NamespaceVML:                     "v",
	NamespaceSpreadsheetML:           "x",
	NamespacePresentationML:          "p",
	NamespaceDiagram:                 "dgm",
	NamespaceDublinCore:              "dc",
	NamespaceDublinCoreTerms:         "dcterms",
	NamespaceCoreProperties:          "cp",
	NamespaceExtendedProperties:      "ep",
	NamespaceXMLSchemaInstance:       "xsi",

	// Word extension namespaces
	NamespaceWord2010:                      "w14",
	NamespaceWord2010WordprocessingDrawing: "wp14",
	NamespaceWord2010WordprocessingCanvas:  "wpc",
	NamespaceWord2010WordprocessingGroup:   "wpg",
	NamespaceWord2010WordprocessingShape:   "wps",
	NamespaceWord2013:                      "w15",
	NamespaceWord2016:                      "w16",
	NamespaceWord2019:                      "w19",

	// Excel extension namespaces
	NamespaceExcel2009:         "x14",
	NamespaceExcel2009AC:       "x14ac",
	NamespaceExcel2013:         "x15",
	NamespaceExcel2013AC:       "x15ac",
	NamespaceExcel2016:         "x16",
	NamespaceExcel2016Revision: "x16r2",
	NamespaceExcel2019:         "x19",
	NamespaceExcel2021:         "x21",
	NamespaceExcel2024:         "x24",
	NamespaceExcel2025:         "x25",

	// PowerPoint extension namespaces
	NamespacePowerPoint2010: "p14",
	NamespacePowerPoint2012: "p15",
	NamespacePowerPoint2016: "p16",
	NamespacePowerPoint2021: "p21",

	// DrawingML extension namespaces
	NamespaceDrawing2010:    "a14",
	NamespaceDrawing2012:    "a15",
	NamespaceDrawing2014:    "a16",
	NamespaceDrawing2016SVG: "asvg",
	NamespaceDrawing2016Ink: "aink",
	NamespaceChart2014:      "c15",
	NamespaceChart2016:      "c16",
	NamespaceChart2016r3:    "c16r3",
}

// PrefixNamespaces maps conventional prefixes to their namespace URIs.
// This is the inverse of NamespacePrefixes for quick prefix lookups.
var PrefixNamespaces map[string]string

func init() {
	// Build reverse map from NamespacePrefixes
	PrefixNamespaces = make(
		map[string]string,
		len(NamespacePrefixes),
	)
	for ns, prefix := range NamespacePrefixes {
		// For duplicate prefixes (like x14), the last one wins
		// This is acceptable as they should be compatible
		PrefixNamespaces[prefix] = ns
	}
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

// GetNamespaceForPrefix returns the namespace URI for a given prefix.
// Returns an empty string if the prefix is not recognized.
func GetNamespaceForPrefix(prefix string) string {
	if ns, ok := PrefixNamespaces[prefix]; ok {
		return ns
	}

	return ""
}

// RegisterNamespace adds or updates a namespace URI to prefix mapping.
// This can be used to register custom or future extension namespaces at runtime.
func RegisterNamespace(
	namespaceURI, prefix string,
) {
	NamespacePrefixes[namespaceURI] = prefix
	PrefixNamespaces[prefix] = namespaceURI
}
