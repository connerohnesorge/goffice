package elements

import (
	"strings"
	"testing"
)

func TestStylesheetCreation(t *testing.T) {
	ss := NewStylesheet()

	if ss == nil {
		t.Fatal("NewStylesheet returned nil")
	}

	if ss.LocalName() != "styleSheet" {
		t.Errorf(
			"expected local name 'styleSheet', got '%s'",
			ss.LocalName(),
		)
	}

	if ss.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			ss.NamespaceURI(),
		)
	}

	// All collections should be nil initially
	if ss.NumFmts() != nil {
		t.Error(
			"expected NumFmts to be nil initially",
		)
	}
	if ss.Fonts() != nil {
		t.Error(
			"expected Fonts to be nil initially",
		)
	}
	if ss.Fills() != nil {
		t.Error(
			"expected Fills to be nil initially",
		)
	}
	if ss.Borders() != nil {
		t.Error(
			"expected Borders to be nil initially",
		)
	}
}

func TestStylesheetInitializeDefaults(
	t *testing.T,
) {
	ss := NewStylesheet()
	ss.InitializeDefaults()

	// Check fonts
	fonts := ss.Fonts()
	if fonts == nil {
		t.Fatal(
			"expected Fonts to exist after InitializeDefaults",
		)
	}
	if fonts.ItemCount() != 1 {
		t.Errorf(
			"expected 1 font, got %d",
			fonts.ItemCount(),
		)
	}
	font := fonts.GetFont(0)
	if font.FontNameValue() != "Calibri" {
		t.Errorf(
			"expected font name 'Calibri', got '%s'",
			font.FontNameValue(),
		)
	}

	// Check fills
	fills := ss.Fills()
	if fills == nil {
		t.Fatal(
			"expected Fills to exist after InitializeDefaults",
		)
	}
	if fills.ItemCount() != 2 {
		t.Errorf(
			"expected 2 fills, got %d",
			fills.ItemCount(),
		)
	}

	// Check borders
	borders := ss.Borders()
	if borders == nil {
		t.Fatal(
			"expected Borders to exist after InitializeDefaults",
		)
	}
	if borders.ItemCount() != 1 {
		t.Errorf(
			"expected 1 border, got %d",
			borders.ItemCount(),
		)
	}

	// Check cellStyleXfs
	cellStyleXfs := ss.CellStyleXfs()
	if cellStyleXfs == nil {
		t.Fatal(
			"expected CellStyleXfs to exist after InitializeDefaults",
		)
	}
	if cellStyleXfs.ItemCount() != 1 {
		t.Errorf(
			"expected 1 cellStyleXf, got %d",
			cellStyleXfs.ItemCount(),
		)
	}

	// Check cellXfs
	cellXfs := ss.CellXfs()
	if cellXfs == nil {
		t.Fatal(
			"expected CellXfs to exist after InitializeDefaults",
		)
	}
	if cellXfs.ItemCount() != 1 {
		t.Errorf(
			"expected 1 cellXf, got %d",
			cellXfs.ItemCount(),
		)
	}

	// Check cellStyles
	cellStyles := ss.CellStyles()
	if cellStyles == nil {
		t.Fatal(
			"expected CellStyles to exist after InitializeDefaults",
		)
	}
	if cellStyles.ItemCount() != 1 {
		t.Errorf(
			"expected 1 cellStyle, got %d",
			cellStyles.ItemCount(),
		)
	}
}

func TestNumFmts(t *testing.T) {
	numFmts := NewNumFmts()

	// Add a custom number format
	nf := numFmts.AddNumFmt(164, "#,##0.00\"$\"")

	if nf.NumFmtId() != 164 {
		t.Errorf(
			"expected numFmtId 164, got %d",
			nf.NumFmtId(),
		)
	}

	if nf.FormatCode() != "#,##0.00\"$\"" {
		t.Errorf(
			"expected formatCode '#,##0.00\"$\"', got '%s'",
			nf.FormatCode(),
		)
	}

	// Check count
	if numFmts.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			numFmts.Count(),
		)
	}

	// Get by ID
	found := numFmts.GetNumFmt(164)
	if found == nil {
		t.Fatal(
			"expected to find numFmt with id 164",
		)
	}

	// Not found
	notFound := numFmts.GetNumFmt(999)
	if notFound != nil {
		t.Error(
			"expected nil for non-existent numFmt",
		)
	}
}

func TestBuiltInNumFmts(t *testing.T) {
	testCases := []struct {
		id        uint32
		expected  string
		isBuiltIn bool
	}{
		{0, "General", true},
		{1, "0", true},
		{2, "0.00", true},
		{14, "m/d/yy", true},
		{49, "@", true},
		{
			164,
			"",
			false,
		}, // Custom formats start at 164
	}

	for _, tc := range testCases {
		code, ok := BuiltInNumFmtCodes[tc.id]
		if tc.isBuiltIn {
			if !ok {
				t.Errorf(
					"expected id %d to be built-in",
					tc.id,
				)
			}
			if code != tc.expected {
				t.Errorf(
					"expected code '%s' for id %d, got '%s'",
					tc.expected,
					tc.id,
					code,
				)
			}
		} else if ok {
			t.Errorf("expected id %d to not be built-in", tc.id)
		}

		if IsBuiltInNumFmt(
			tc.id,
		) != tc.isBuiltIn {
			t.Errorf(
				"IsBuiltInNumFmt(%d) = %v, expected %v",
				tc.id,
				IsBuiltInNumFmt(tc.id),
				tc.isBuiltIn,
			)
		}
	}
}

func TestFonts(t *testing.T) {
	fonts := NewFonts()

	// Add a font
	font := fonts.AddFont()
	font.SetBold(true)
	font.SetItalic(true)
	font.SetFontSize(14)
	font.SetFontName("Arial")
	font.SetFontFamily(2)
	font.SetUnderline(UnderlineStyleDouble)
	font.SetStrikethrough(true)

	// Verify properties
	if !font.IsBold() {
		t.Error("expected font to be bold")
	}
	if !font.IsItalic() {
		t.Error("expected font to be italic")
	}
	if font.FontSizeValue() != 14 {
		t.Errorf(
			"expected font size 14, got %f",
			font.FontSizeValue(),
		)
	}
	if font.FontNameValue() != "Arial" {
		t.Errorf(
			"expected font name 'Arial', got '%s'",
			font.FontNameValue(),
		)
	}
	if font.FontFamilyValue() != 2 {
		t.Errorf(
			"expected font family 2, got %d",
			font.FontFamilyValue(),
		)
	}
	if font.UnderlineStyle() != UnderlineStyleDouble {
		t.Errorf(
			"expected underline style 'double', got '%s'",
			font.UnderlineStyle(),
		)
	}
	if !font.IsStrikethrough() {
		t.Error(
			"expected font to have strikethrough",
		)
	}

	// Test color
	color := font.GetOrCreateColor()
	color.SetRGB("FFFF0000")
	if font.Color().RGB() != "FFFF0000" {
		t.Errorf(
			"expected RGB 'FFFF0000', got '%s'",
			font.Color().RGB(),
		)
	}

	// Test count
	if fonts.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			fonts.Count(),
		)
	}

	// Get font by index
	if fonts.GetFont(0) == nil {
		t.Error("expected to get font at index 0")
	}
	if fonts.GetFont(99) != nil {
		t.Error(
			"expected nil for out of range index",
		)
	}
}

func TestFontScheme(t *testing.T) {
	font := NewFont()

	// Test default
	if font.FontSchemeValue() != FontSchemeNone {
		t.Errorf(
			"expected default scheme 'none', got '%s'",
			font.FontSchemeValue(),
		)
	}

	// Set major
	font.SetFontScheme(FontSchemeMajor)
	if font.FontSchemeValue() != FontSchemeMajor {
		t.Errorf(
			"expected scheme 'major', got '%s'",
			font.FontSchemeValue(),
		)
	}

	// Set minor
	font.SetFontScheme(FontSchemeMinor)
	if font.FontSchemeValue() != FontSchemeMinor {
		t.Errorf(
			"expected scheme 'minor', got '%s'",
			font.FontSchemeValue(),
		)
	}
}

func TestFills(t *testing.T) {
	fills := NewFills()

	// Add pattern fill
	fill := fills.AddFill()
	pf := fill.GetOrCreatePatternFill()
	pf.SetPatternType(PatternTypeSolid)

	fgColor := pf.GetOrCreateFgColor()
	fgColor.SetRGB("FF0000FF")

	// Verify
	if fill.PatternFill() == nil {
		t.Fatal("expected PatternFill to exist")
	}
	if pf.PatternType() != PatternTypeSolid {
		t.Errorf(
			"expected pattern type 'solid', got '%s'",
			pf.PatternType(),
		)
	}
	if pf.FgColor().RGB() != "FF0000FF" {
		t.Errorf(
			"expected fgColor 'FF0000FF', got '%s'",
			pf.FgColor().RGB(),
		)
	}

	// Test gradient fill
	fill2 := fills.AddFill()
	gf := fill2.GetOrCreateGradientFill()
	gf.SetDegree(90)

	stop1 := gf.AddStop(0)
	stop1.GetOrCreateColor().SetRGB("FFFFFFFF")

	stop2 := gf.AddStop(1)
	stop2.GetOrCreateColor().SetRGB("FF000000")

	if gf.Degree() != 90 {
		t.Errorf(
			"expected degree 90, got %f",
			gf.Degree(),
		)
	}

	stopCount := 0
	for range gf.Stops() {
		stopCount++
	}
	if stopCount != 2 {
		t.Errorf(
			"expected 2 stops, got %d",
			stopCount,
		)
	}
}

func TestPatternTypes(t *testing.T) {
	testCases := []PatternType{
		PatternTypeNone,
		PatternTypeSolid,
		PatternTypeMediumGray,
		PatternTypeDarkGray,
		PatternTypeLightGray,
		PatternTypeGray125,
		PatternTypeGray0625,
	}

	for _, pt := range testCases {
		pf := NewPatternFill()
		pf.SetPatternType(pt)
		if pf.PatternType() != pt {
			t.Errorf(
				"expected pattern type '%s', got '%s'",
				pt,
				pf.PatternType(),
			)
		}
	}
}

func TestBorders(t *testing.T) {
	borders := NewBorders()

	border := borders.AddBorder()

	// Set left border
	left := border.GetOrCreateLeft()
	left.SetStyle(BorderStyleThin)
	left.GetOrCreateColor().SetRGB("FF000000")

	// Set right border
	right := border.GetOrCreateRight()
	right.SetStyle(BorderStyleMedium)

	// Set top border
	top := border.GetOrCreateTop()
	top.SetStyle(BorderStyleDashed)

	// Set bottom border
	bottom := border.GetOrCreateBottom()
	bottom.SetStyle(BorderStyleDouble)

	// Set diagonal
	border.SetDiagonalUp(true)
	diagonal := border.GetOrCreateDiagonal()
	diagonal.SetStyle(BorderStyleDotted)

	// Verify
	if border.Left().Style() != BorderStyleThin {
		t.Errorf(
			"expected left style 'thin', got '%s'",
			border.Left().Style(),
		)
	}
	if border.Right().
		Style() !=
		BorderStyleMedium {
		t.Errorf(
			"expected right style 'medium', got '%s'",
			border.Right().Style(),
		)
	}
	if border.Top().Style() != BorderStyleDashed {
		t.Errorf(
			"expected top style 'dashed', got '%s'",
			border.Top().Style(),
		)
	}
	if border.Bottom().
		Style() !=
		BorderStyleDouble {
		t.Errorf(
			"expected bottom style 'double', got '%s'",
			border.Bottom().Style(),
		)
	}
	if !border.DiagonalUp() {
		t.Error("expected diagonalUp to be true")
	}
	if border.Diagonal().
		Style() !=
		BorderStyleDotted {
		t.Errorf(
			"expected diagonal style 'dotted', got '%s'",
			border.Diagonal().Style(),
		)
	}
}

func TestBorderStyles(t *testing.T) {
	testCases := []BorderStyle{
		BorderStyleNone,
		BorderStyleThin,
		BorderStyleMedium,
		BorderStyleDashed,
		BorderStyleDotted,
		BorderStyleThick,
		BorderStyleDouble,
		BorderStyleHair,
		BorderStyleMediumDashed,
		BorderStyleDashDot,
		BorderStyleMediumDashDot,
		BorderStyleDashDotDot,
		BorderStyleMediumDashDotDot,
		BorderStyleSlantDashDot,
	}

	for _, bs := range testCases {
		bp := NewBorderPrLeft()
		bp.SetStyle(bs)
		if bp.Style() != bs {
			t.Errorf(
				"expected border style '%s', got '%s'",
				bs,
				bp.Style(),
			)
		}
	}
}

func TestAlignment(t *testing.T) {
	align := NewAlignment()

	// Test horizontal alignment
	align.SetHorizontal(HorizontalAlignmentCenter)
	if align.Horizontal() != HorizontalAlignmentCenter {
		t.Errorf(
			"expected horizontal 'center', got '%s'",
			align.Horizontal(),
		)
	}

	// Test vertical alignment
	align.SetVertical(VerticalAlignmentTop)
	if align.Vertical() != VerticalAlignmentTop {
		t.Errorf(
			"expected vertical 'top', got '%s'",
			align.Vertical(),
		)
	}

	// Test text rotation
	align.SetTextRotation(45)
	if align.TextRotation() != 45 {
		t.Errorf(
			"expected textRotation 45, got %d",
			align.TextRotation(),
		)
	}

	// Test wrap text
	align.SetWrapText(true)
	if !align.WrapText() {
		t.Error("expected wrapText to be true")
	}

	// Test indent
	align.SetIndent(2)
	if align.Indent() != 2 {
		t.Errorf(
			"expected indent 2, got %d",
			align.Indent(),
		)
	}

	// Test shrink to fit
	align.SetShrinkToFit(true)
	if !align.ShrinkToFit() {
		t.Error("expected shrinkToFit to be true")
	}

	// Test reading order
	align.SetReadingOrder(ReadingOrderRightToLeft)
	if align.ReadingOrder() != ReadingOrderRightToLeft {
		t.Errorf(
			"expected readingOrder 2, got %d",
			align.ReadingOrder(),
		)
	}
}

func TestProtection(t *testing.T) {
	prot := NewProtection()

	// Test default locked (true)
	if !prot.Locked() {
		t.Error(
			"expected default locked to be true",
		)
	}

	// Test set unlocked
	prot.SetLocked(false)
	if prot.Locked() {
		t.Error(
			"expected locked to be false after SetLocked(false)",
		)
	}

	// Test default hidden (false)
	if prot.Hidden() {
		t.Error(
			"expected default hidden to be false",
		)
	}

	// Test set hidden
	prot.SetHidden(true)
	if !prot.Hidden() {
		t.Error(
			"expected hidden to be true after SetHidden(true)",
		)
	}
}

func TestCellXfs(t *testing.T) {
	cellXfs := NewCellXfs()

	xf := cellXfs.AddXf()
	xf.SetNumFmtId(1)
	xf.SetFontId(0)
	xf.SetFillId(0)
	xf.SetBorderId(0)
	xf.SetApplyNumberFormat(true)
	xf.SetApplyFont(true)

	// Verify
	if xf.NumFmtId() != 1 {
		t.Errorf(
			"expected numFmtId 1, got %d",
			xf.NumFmtId(),
		)
	}
	if !xf.ApplyNumberFormat() {
		t.Error(
			"expected applyNumberFormat to be true",
		)
	}
	if !xf.ApplyFont() {
		t.Error("expected applyFont to be true")
	}

	// Test alignment
	align := xf.GetOrCreateAlignment()
	align.SetHorizontal(HorizontalAlignmentRight)
	xf.SetApplyAlignment(true)

	if !xf.ApplyAlignment() {
		t.Error(
			"expected applyAlignment to be true",
		)
	}
	if xf.Alignment().
		Horizontal() !=
		HorizontalAlignmentRight {
		t.Errorf(
			"expected horizontal 'right', got '%s'",
			xf.Alignment().Horizontal(),
		)
	}

	// Test protection
	prot := xf.GetOrCreateProtection()
	prot.SetLocked(false)
	xf.SetApplyProtection(true)

	if !xf.ApplyProtection() {
		t.Error(
			"expected applyProtection to be true",
		)
	}
	if xf.Protection().Locked() {
		t.Error("expected locked to be false")
	}
}

func TestCellStyles(t *testing.T) {
	cellStyles := NewCellStyles()

	style := cellStyles.AddCellStyle()
	style.SetName("Custom Style")
	style.SetXfId(1)
	style.SetBuiltinId(BuiltinStyleNormal)

	// Verify
	if style.Name() != "Custom Style" {
		t.Errorf(
			"expected name 'Custom Style', got '%s'",
			style.Name(),
		)
	}
	if style.XfId() != 1 {
		t.Errorf(
			"expected xfId 1, got %d",
			style.XfId(),
		)
	}
	if !style.HasBuiltinId() {
		t.Error(
			"expected style to have builtinId",
		)
	}
	if style.BuiltinId() != BuiltinStyleNormal {
		t.Errorf(
			"expected builtinId 0, got %d",
			style.BuiltinId(),
		)
	}

	// Get by name
	found := cellStyles.GetCellStyle(
		"Custom Style",
	)
	if found == nil {
		t.Fatal("expected to find style by name")
	}

	// Not found
	notFound := cellStyles.GetCellStyle(
		"Nonexistent",
	)
	if notFound != nil {
		t.Error(
			"expected nil for non-existent style",
		)
	}
}

func TestDxfs(t *testing.T) {
	dxfs := NewDxfs()

	dxf := dxfs.AddDxf()

	// Add font formatting
	font := dxf.GetOrCreateFont()
	font.SetBold(true)
	font.GetOrCreateColor().SetRGB("FFFF0000")

	// Add fill formatting
	fill := dxf.GetOrCreateFill()
	pf := fill.GetOrCreatePatternFill()
	pf.SetPatternType(PatternTypeSolid)
	pf.GetOrCreateFgColor().SetRGB("FFFFFF00")

	// Verify
	if dxfs.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			dxfs.Count(),
		)
	}

	if !dxf.Font().IsBold() {
		t.Error("expected font to be bold")
	}
	if dxf.Fill().
		PatternFill().
		PatternType() !=
		PatternTypeSolid {
		t.Errorf(
			"expected pattern type 'solid', got '%s'",
			dxf.Fill().
				PatternFill().
				PatternType(),
		)
	}
}

func TestTableStyles(t *testing.T) {
	tableStyles := NewTableStyles()
	tableStyles.SetDefaultTableStyle(
		DefaultTableStyle,
	)
	tableStyles.SetDefaultPivotStyle(
		DefaultPivotStyle,
	)

	// Verify defaults
	if tableStyles.DefaultTableStyle() != DefaultTableStyle {
		t.Errorf(
			"expected default table style '%s', got '%s'",
			DefaultTableStyle,
			tableStyles.DefaultTableStyle(),
		)
	}
	if tableStyles.DefaultPivotStyle() != DefaultPivotStyle {
		t.Errorf(
			"expected default pivot style '%s', got '%s'",
			DefaultPivotStyle,
			tableStyles.DefaultPivotStyle(),
		)
	}

	// Add custom table style
	style := tableStyles.AddTableStyle("MyStyle")
	style.AddTableStyleElement(
		TableStyleElementHeaderRow,
		0,
	)
	style.AddTableStyleElement(
		TableStyleElementFirstRowStripe,
		1,
	)

	// Verify
	if style.Name() != "MyStyle" {
		t.Errorf(
			"expected name 'MyStyle', got '%s'",
			style.Name(),
		)
	}
	if style.Count() != 2 {
		t.Errorf(
			"expected 2 elements, got %d",
			style.Count(),
		)
	}

	// Get by name
	found := tableStyles.GetTableStyle("MyStyle")
	if found == nil {
		t.Fatal("expected to find style by name")
	}
}

func TestColors(t *testing.T) {
	colors := NewColors()

	// Add indexed colors
	indexed := colors.GetOrCreateIndexedColors()
	indexed.AddRgbColor("FFFF0000")
	indexed.AddRgbColor("FF00FF00")
	indexed.AddRgbColor("FF0000FF")

	// Verify
	if indexed.ColorCount() != 3 {
		t.Errorf(
			"expected 3 colors, got %d",
			indexed.ColorCount(),
		)
	}

	color := indexed.GetRgbColor(1)
	if color == nil {
		t.Fatal(
			"expected to get color at index 1",
		)
	}
	if color.Rgb() != "FF00FF00" {
		t.Errorf(
			"expected RGB 'FF00FF00', got '%s'",
			color.Rgb(),
		)
	}

	// Add MRU colors
	mru := colors.GetOrCreateMruColors()
	c := mru.AddColor()
	c.SetRGB("FFAABBCC")

	if mru.ColorCount() != 1 {
		t.Errorf(
			"expected 1 MRU color, got %d",
			mru.ColorCount(),
		)
	}
}

func TestThemeColors(t *testing.T) {
	// Verify theme color constants
	if ThemeColorLight1 != 0 {
		t.Errorf(
			"expected ThemeColorLight1 = 0, got %d",
			ThemeColorLight1,
		)
	}
	if ThemeColorDark1 != 1 {
		t.Errorf(
			"expected ThemeColorDark1 = 1, got %d",
			ThemeColorDark1,
		)
	}
	if ThemeColorAccent1 != 4 {
		t.Errorf(
			"expected ThemeColorAccent1 = 4, got %d",
			ThemeColorAccent1,
		)
	}
	if ThemeColorAccent6 != 9 {
		t.Errorf(
			"expected ThemeColorAccent6 = 9, got %d",
			ThemeColorAccent6,
		)
	}
}

func TestColorElement(t *testing.T) {
	color := NewColor()

	// Test auto
	color.SetAuto(true)
	if !color.Auto() {
		t.Error("expected auto to be true")
	}
	color.SetAuto(false)
	if color.Auto() {
		t.Error(
			"expected auto to be false after SetAuto(false)",
		)
	}

	// Test RGB
	color.SetRGB("FF123456")
	if color.RGB() != "FF123456" {
		t.Errorf(
			"expected RGB 'FF123456', got '%s'",
			color.RGB(),
		)
	}

	// Test theme
	color.SetTheme(ThemeColorAccent1)
	if color.Theme() != ThemeColorAccent1 {
		t.Errorf(
			"expected theme %d, got %d",
			ThemeColorAccent1,
			color.Theme(),
		)
	}

	// Test indexed
	color.SetIndexed(10)
	if color.Indexed() != 10 {
		t.Errorf(
			"expected indexed 10, got %d",
			color.Indexed(),
		)
	}

	// Test tint
	color.SetTint(0.5)
	if color.Tint() != 0.5 {
		t.Errorf(
			"expected tint 0.5, got %f",
			color.Tint(),
		)
	}
}

func TestStylesheetXML(t *testing.T) {
	ss := NewStylesheet()
	ss.InitializeDefaults()

	xml := ss.OuterXml()

	// Check XML contains expected elements
	if !strings.Contains(xml, "styleSheet") {
		t.Error(
			"XML should contain styleSheet element",
		)
	}
	if !strings.Contains(xml, "fonts") {
		t.Error(
			"XML should contain fonts element",
		)
	}
	if !strings.Contains(xml, "fills") {
		t.Error(
			"XML should contain fills element",
		)
	}
	if !strings.Contains(xml, "borders") {
		t.Error(
			"XML should contain borders element",
		)
	}
	if !strings.Contains(xml, "cellStyleXfs") {
		t.Error(
			"XML should contain cellStyleXfs element",
		)
	}
	if !strings.Contains(xml, "cellXfs") {
		t.Error(
			"XML should contain cellXfs element",
		)
	}
	if !strings.Contains(xml, "cellStyles") {
		t.Error(
			"XML should contain cellStyles element",
		)
	}
}

func TestStylesheetClone(t *testing.T) {
	ss := NewStylesheet()
	ss.InitializeDefaults()

	// Add custom format
	numFmts := ss.GetOrCreateNumFmts()
	numFmts.AddNumFmt(164, "#,##0.00")

	// Clone
	clonedResult := ss.Clone()
	cloned, ok := clonedResult.(*Stylesheet)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Stylesheet",
			clonedResult,
		)
	}

	// Verify clone has same content
	if cloned.NumFmts().ItemCount() != 1 {
		t.Errorf(
			"expected 1 numFmt in clone, got %d",
			cloned.NumFmts().ItemCount(),
		)
	}

	// Modify clone
	cloned.NumFmts().AddNumFmt(165, "0.00%")

	// Original should be unchanged
	if ss.NumFmts().ItemCount() != 1 {
		t.Errorf(
			"expected 1 numFmt in original after clone modification, got %d",
			ss.NumFmts().ItemCount(),
		)
	}
}

func TestElementOrdering(t *testing.T) {
	ss := NewStylesheet()

	// Add elements in reverse order
	ss.GetOrCreateColors()
	ss.GetOrCreateTableStyles()
	ss.GetOrCreateDxfs()
	ss.GetOrCreateCellStyles()
	ss.GetOrCreateCellXfs()
	ss.GetOrCreateCellStyleXfs()
	ss.GetOrCreateBorders()
	ss.GetOrCreateFills()
	ss.GetOrCreateFonts()
	ss.GetOrCreateNumFmts()

	xml := ss.OuterXml()

	// Check order
	numFmtsPos := strings.Index(xml, "<numFmts")
	fontsPos := strings.Index(xml, "<fonts")
	fillsPos := strings.Index(xml, "<fills")
	bordersPos := strings.Index(xml, "<borders")
	cellStyleXfsPos := strings.Index(
		xml,
		"<cellStyleXfs",
	)
	cellXfsPos := strings.Index(xml, "<cellXfs")
	cellStylesPos := strings.Index(
		xml,
		"<cellStyles",
	)
	dxfsPos := strings.Index(xml, "<dxfs")
	tableStylesPos := strings.Index(
		xml,
		"<tableStyles",
	)
	colorsPos := strings.Index(xml, "<colors")

	if numFmtsPos > fontsPos {
		t.Error(
			"numFmts should come before fonts",
		)
	}
	if fontsPos > fillsPos {
		t.Error("fonts should come before fills")
	}
	if fillsPos > bordersPos {
		t.Error(
			"fills should come before borders",
		)
	}
	if bordersPos > cellStyleXfsPos {
		t.Error(
			"borders should come before cellStyleXfs",
		)
	}
	if cellStyleXfsPos > cellXfsPos {
		t.Error(
			"cellStyleXfs should come before cellXfs",
		)
	}
	if cellXfsPos > cellStylesPos {
		t.Error(
			"cellXfs should come before cellStyles",
		)
	}
	if cellStylesPos > dxfsPos {
		t.Error(
			"cellStyles should come before dxfs",
		)
	}
	if dxfsPos > tableStylesPos {
		t.Error(
			"dxfs should come before tableStyles",
		)
	}
	if tableStylesPos > colorsPos {
		t.Error(
			"tableStyles should come before colors",
		)
	}
}

func TestBuiltinCellStyles(t *testing.T) {
	testCases := []struct {
		id   uint32
		name string
	}{
		{BuiltinStyleNormal, "Normal"},
		{BuiltinStyleComma, "Comma"},
		{BuiltinStyleCurrency, "Currency"},
		{BuiltinStylePercent, "Percent"},
		{BuiltinStyleHyperlink, "Hyperlink"},
		{BuiltinStyleTitle, "Title"},
		{BuiltinStyleHeading1, "Heading 1"},
		{BuiltinStyleGood, "Good"},
		{BuiltinStyleBad, "Bad"},
		{BuiltinStyleNeutral, "Neutral"},
	}

	for _, tc := range testCases {
		name, ok := BuiltinStyleNames[tc.id]
		if !ok {
			t.Errorf(
				"expected builtin style id %d to exist",
				tc.id,
			)

			continue
		}
		if name != tc.name {
			t.Errorf(
				"expected name '%s' for id %d, got '%s'",
				tc.name,
				tc.id,
				name,
			)
		}
	}
}
