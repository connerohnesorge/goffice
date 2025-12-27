//nolint:revive // Long namespace URIs cannot be broken across lines
package drawingml

// DrawingML namespace URIs for the main DrawingML schemas.
const (
	// NamespaceMain is the main DrawingML namespace (a:).
	// Used for core drawing elements like shapes, text, and effects.
	NamespaceMain = "http://schemas.openxmlformats.org/drawingml/2006/main"

	// NamespacePicture is the DrawingML Picture namespace (pic:).
	// Used for image content within drawings.
	NamespacePicture = "http://schemas.openxmlformats.org/drawingml/2006/picture"

	// NamespaceChart is the DrawingML Chart namespace (c:).
	// Used for chart elements and data visualization.
	NamespaceChart = "http://schemas.openxmlformats.org/drawingml/2006/chart"

	// NamespaceChartDrawing is the DrawingML ChartDrawing namespace (cdr:).
	// Used for drawings embedded within charts.
	NamespaceChartDrawing = "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing"

	// NamespaceDiagram is the DrawingML Diagram namespace (dgm:).
	// Used for SmartArt diagrams and graphical layouts.
	NamespaceDiagram = "http://schemas.openxmlformats.org/drawingml/2006/diagram"

	// NamespaceSpreadsheetDrawing is the SpreadsheetDrawing namespace (xdr:).
	// Used for positioning drawings within Excel worksheets.
	NamespaceSpreadsheetDrawing = "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"

	// NamespaceWordprocessingDrawing is the WordprocessingDrawing namespace (wp:).
	// Used for anchoring and positioning drawings in Word documents.
	NamespaceWordprocessingDrawing = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"

	// NamespaceLockedCanvas is the Locked Canvas namespace.
	// Used for locked drawing canvases.
	NamespaceLockedCanvas = "http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas"
)

// Microsoft Office extension namespaces for DrawingML.
const (
	// NamespaceDrawingML2010Main is the Office 2010 DrawingML extensions (a14:).
	NamespaceDrawingML2010Main = "http://schemas.microsoft.com/office/drawing/2010/main"

	// NamespaceDrawingML2012Main is the Office 2012 DrawingML extensions (a15:).
	NamespaceDrawingML2012Main = "http://schemas.microsoft.com/office/drawing/2012/main"

	// NamespaceDrawingML2014Main is the Office 2014 DrawingML extensions (a16:).
	NamespaceDrawingML2014Main = "http://schemas.microsoft.com/office/drawing/2014/main"

	// NamespaceDrawingML2016SVG is the Office 2016 SVG extensions (asvg:).
	NamespaceDrawingML2016SVG = "http://schemas.microsoft.com/office/drawing/2016/SVG/main"

	// NamespaceDrawingML3D is the 3D scene namespace.
	NamespaceDrawingML3D = "http://schemas.microsoft.com/office/drawing/2010/slicer"

	// NamespaceChart2014 is the Office 2014 Chart extensions (c15:).
	NamespaceChart2014 = "http://schemas.microsoft.com/office/drawing/2012/chart"

	// NamespaceChart2016 is the Office 2016 Chart extensions (c16:).
	NamespaceChart2016 = "http://schemas.microsoft.com/office/drawing/2014/chart"

	// NamespaceWordprocessingDrawing2010 is the Office 2010 WordprocessingDrawing extensions (wp14:).
	NamespaceWordprocessingDrawing2010 = "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"

	// NamespaceSpreadsheetDrawing2010 is the Office 2010 SpreadsheetDrawing extensions (xdr14:).
	NamespaceSpreadsheetDrawing2010 = "http://schemas.microsoft.com/office/excel/2010/spreadsheetDrawing"
)

// Prefix constants for DrawingML namespaces.
const (
	// PrefixMain is the conventional prefix for main DrawingML elements.
	PrefixMain = "a"

	// PrefixPicture is the conventional prefix for Picture elements.
	PrefixPicture = "pic"

	// PrefixChart is the conventional prefix for Chart elements.
	PrefixChart = "c"

	// PrefixChartDrawing is the conventional prefix for ChartDrawing elements.
	PrefixChartDrawing = "cdr"

	// PrefixDiagram is the conventional prefix for Diagram elements.
	PrefixDiagram = "dgm"

	// PrefixSpreadsheetDrawing is the conventional prefix for SpreadsheetDrawing elements.
	PrefixSpreadsheetDrawing = "xdr"

	// PrefixWordprocessingDrawing is the conventional prefix for WordprocessingDrawing elements.
	PrefixWordprocessingDrawing = "wp"

	// PrefixLockedCanvas is the conventional prefix for LockedCanvas elements.
	PrefixLockedCanvas = "lc"

	// PrefixDrawingML2010 is the prefix for Office 2010 DrawingML extensions.
	PrefixDrawingML2010 = "a14"

	// PrefixDrawingML2012 is the prefix for Office 2012 DrawingML extensions.
	PrefixDrawingML2012 = "a15"

	// PrefixDrawingML2014 is the prefix for Office 2014 DrawingML extensions.
	PrefixDrawingML2014 = "a16"

	// PrefixWordprocessingDrawing2010 is the prefix for Office 2010 WordprocessingDrawing extensions.
	PrefixWordprocessingDrawing2010 = "wp14"

	// PrefixChart2014 is the prefix for Office 2014 Chart extensions.
	PrefixChart2014 = "c15"

	// PrefixChart2016 is the prefix for Office 2016 Chart extensions.
	PrefixChart2016 = "c16"
)

// NamespacePrefixMap maps DrawingML namespace URIs to their conventional prefixes.
var NamespacePrefixMap = map[string]string{
	NamespaceMain:                      PrefixMain,
	NamespacePicture:                   PrefixPicture,
	NamespaceChart:                     PrefixChart,
	NamespaceChartDrawing:              PrefixChartDrawing,
	NamespaceDiagram:                   PrefixDiagram,
	NamespaceSpreadsheetDrawing:        PrefixSpreadsheetDrawing,
	NamespaceWordprocessingDrawing:     PrefixWordprocessingDrawing,
	NamespaceLockedCanvas:              PrefixLockedCanvas,
	NamespaceDrawingML2010Main:         PrefixDrawingML2010,
	NamespaceDrawingML2012Main:         PrefixDrawingML2012,
	NamespaceDrawingML2014Main:         PrefixDrawingML2014,
	NamespaceWordprocessingDrawing2010: PrefixWordprocessingDrawing2010,
	NamespaceChart2014:                 PrefixChart2014,
	NamespaceChart2016:                 PrefixChart2016,
}

// GetPrefixForNamespace returns the conventional prefix for a DrawingML namespace URI.
// Returns an empty string if the namespace is not recognized.
func GetPrefixForNamespace(
	namespaceURI string,
) string {
	if prefix, ok := NamespacePrefixMap[namespaceURI]; ok {
		return prefix
	}

	return ""
}

// GetNamespaceForPrefix returns the namespace URI for a DrawingML prefix.
// Returns an empty string if the prefix is not recognized.
func GetNamespaceForPrefix(prefix string) string {
	for ns, p := range NamespacePrefixMap {
		if p == prefix {
			return ns
		}
	}

	return ""
}
