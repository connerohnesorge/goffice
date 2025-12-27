package parts

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestFontsPart_GetFont(t *testing.T) {
	// Create a FontsPart
	fp := &FontsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/fontTable.xml",
			ContentTypeFontTable,
			nil,
			nil,
		),
	}

	// Create and set up the Fonts root element
	fonts := elements.NewFonts()
	fonts.AddFontByName("Calibri")
	fonts.AddFontByName("Times New Roman")
	fonts.AddFontByName("Arial")
	fp.SetRootElement(fonts)

	// Test getting an existing font
	t.Run(
		"Get existing font",
		func(t *testing.T) {
			font := fp.GetFont("Calibri")
			if font == nil {
				t.Fatal(
					"GetFont returned nil for existing font",
				)
			}
			if font.Name() != "Calibri" {
				t.Errorf(
					"GetFont returned font with name %q, want %q",
					font.Name(),
					"Calibri",
				)
			}
		},
	)

	// Test getting another existing font
	t.Run(
		"Get another existing font",
		func(t *testing.T) {
			font := fp.GetFont("Times New Roman")
			if font == nil {
				t.Fatal(
					"GetFont returned nil for existing font",
				)
			}
			if font.Name() != "Times New Roman" {
				t.Errorf(
					"GetFont returned font with name %q, want %q",
					font.Name(),
					"Times New Roman",
				)
			}
		},
	)

	// Test getting non-existent font
	t.Run(
		"Get non-existent font",
		func(t *testing.T) {
			font := fp.GetFont("NonExistent")
			if font != nil {
				t.Errorf(
					"GetFont returned non-nil for non-existent font: %v",
					font,
				)
			}
		},
	)
}

func TestStylesPart_GetStyleById(t *testing.T) {
	// Create a StylesPart
	sp := &StylesPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/styles.xml",
			ContentTypeStyles,
			nil,
			nil,
		),
	}

	// Create and set up the Styles root element
	styles := elements.NewStyles()
	normalStyle := elements.NewParagraphStyle(
		"Normal",
		"Normal",
	)
	heading1Style := elements.NewParagraphStyle(
		"Heading1",
		"Heading 1",
	)
	heading2Style := elements.NewParagraphStyle(
		"Heading2",
		"Heading 2",
	)
	styles.AddStyle(normalStyle)
	styles.AddStyle(heading1Style)
	styles.AddStyle(heading2Style)
	sp.SetRootElement(styles)

	// Test getting an existing style by ID
	t.Run(
		"Get existing style by ID",
		func(t *testing.T) {
			style := sp.GetStyleById("Heading1")
			if style == nil {
				t.Fatal(
					"GetStyleById returned nil for existing style",
				)
			}
			if style.StyleId() != "Heading1" {
				t.Errorf(
					"GetStyleById returned style with ID %q, want %q",
					style.StyleId(),
					"Heading1",
				)
			}
		},
	)

	// Test getting non-existent style by ID
	t.Run(
		"Get non-existent style by ID",
		func(t *testing.T) {
			style := sp.GetStyleById(
				"NonExistent",
			)
			if style != nil {
				t.Errorf(
					"GetStyleById returned non-nil for non-existent style: %v",
					style,
				)
			}
		},
	)
}

func TestStylesPart_GetStyleByName(t *testing.T) {
	// Create a StylesPart
	sp := &StylesPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/styles.xml",
			ContentTypeStyles,
			nil,
			nil,
		),
	}

	// Create and set up the Styles root element
	styles := elements.NewStyles()
	normalStyle := elements.NewParagraphStyle(
		"Normal",
		"Normal",
	)
	heading1Style := elements.NewParagraphStyle(
		"Heading1",
		"Heading 1",
	)
	heading2Style := elements.NewParagraphStyle(
		"Heading2",
		"Heading 2",
	)
	styles.AddStyle(normalStyle)
	styles.AddStyle(heading1Style)
	styles.AddStyle(heading2Style)
	sp.SetRootElement(styles)

	// Test getting an existing style by name
	t.Run(
		"Get existing style by name",
		func(t *testing.T) {
			style := sp.GetStyleByName(
				"Heading 1",
			)
			if style == nil {
				t.Fatal(
					"GetStyleByName returned nil for existing style",
				)
			}
			if style.StyleName() != "Heading 1" {
				t.Errorf(
					"GetStyleByName returned style with name %q, want %q",
					style.StyleName(),
					"Heading 1",
				)
			}
			if style.StyleId() != "Heading1" {
				t.Errorf(
					"GetStyleByName returned style with ID %q, want %q",
					style.StyleId(),
					"Heading1",
				)
			}
		},
	)

	// Test getting non-existent style by name
	t.Run(
		"Get non-existent style by name",
		func(t *testing.T) {
			style := sp.GetStyleByName(
				"NonExistent",
			)
			if style != nil {
				t.Errorf(
					"GetStyleByName returned non-nil for non-existent style: %v",
					style,
				)
			}
		},
	)
}

func TestNumberingPart_GetAbstractNum(
	t *testing.T,
) {
	// Create a NumberingPart
	np := &NumberingPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/numbering.xml",
			ContentTypeNumbering,
			nil,
			nil,
		),
	}

	// Create and set up the Numbering root element
	numbering := elements.NewNumbering()

	// Add some abstract numbering definitions
	absNum1 := elements.NewDecimalList()
	absNum1.SetAbstractNumId(0)
	numbering.AddAbstractNum(absNum1)

	absNum2 := elements.NewBulletList(
		"\u2022",
		"Symbol",
	)
	absNum2.SetAbstractNumId(1)
	numbering.AddAbstractNum(absNum2)

	np.SetRootElement(numbering)

	// Test getting an existing abstract num
	t.Run(
		"Get existing abstract num",
		func(t *testing.T) {
			absNum := np.GetAbstractNum(0)
			if absNum == nil {
				t.Fatal(
					"GetAbstractNum returned nil for existing abstract num",
				)
			}
			if absNum.AbstractNumId() != 0 {
				t.Errorf(
					"GetAbstractNum returned abstract num with ID %d, want %d",
					absNum.AbstractNumId(),
					0,
				)
			}
		},
	)

	// Test getting another existing abstract num
	t.Run(
		"Get another existing abstract num",
		func(t *testing.T) {
			absNum := np.GetAbstractNum(1)
			if absNum == nil {
				t.Fatal(
					"GetAbstractNum returned nil for existing abstract num",
				)
			}
			if absNum.AbstractNumId() != 1 {
				t.Errorf(
					"GetAbstractNum returned abstract num with ID %d, want %d",
					absNum.AbstractNumId(),
					1,
				)
			}
		},
	)

	// Test getting non-existent abstract num
	t.Run(
		"Get non-existent abstract num",
		func(t *testing.T) {
			absNum := np.GetAbstractNum(999)
			if absNum != nil {
				t.Errorf(
					"GetAbstractNum returned non-nil for non-existent abstract num: %v",
					absNum,
				)
			}
		},
	)
}

func TestNumberingPart_GetNumInstance(
	t *testing.T,
) {
	// Create a NumberingPart
	np := &NumberingPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/numbering.xml",
			ContentTypeNumbering,
			nil,
			nil,
		),
	}

	// Create and set up the Numbering root element
	numbering := elements.NewNumbering()

	// Add some abstract numbering definitions first
	absNum1 := elements.NewDecimalList()
	absNum1Id := numbering.AddAbstractNum(absNum1)

	// Add numbering instances
	numInst1 := elements.NewNumberingInstance(
		absNum1Id,
	)
	numInst1Id := numbering.AddNumInstance(
		numInst1,
	)

	numInst2 := elements.NewNumberingInstance(
		absNum1Id,
	)
	numInst2Id := numbering.AddNumInstance(
		numInst2,
	)

	np.SetRootElement(numbering)

	// Test getting an existing num instance
	t.Run(
		"Get existing num instance",
		func(t *testing.T) {
			numInst := np.GetNumInstance(
				numInst1Id,
			)
			if numInst == nil {
				t.Fatal(
					"GetNumInstance returned nil for existing num instance",
				)
			}
			if numInst.NumId() != numInst1Id {
				t.Errorf(
					"GetNumInstance returned num instance with ID %d, want %d",
					numInst.NumId(),
					numInst1Id,
				)
			}
		},
	)

	// Test getting another existing num instance
	t.Run(
		"Get another existing num instance",
		func(t *testing.T) {
			numInst := np.GetNumInstance(
				numInst2Id,
			)
			if numInst == nil {
				t.Fatal(
					"GetNumInstance returned nil for existing num instance",
				)
			}
			if numInst.NumId() != numInst2Id {
				t.Errorf(
					"GetNumInstance returned num instance with ID %d, want %d",
					numInst.NumId(),
					numInst2Id,
				)
			}
		},
	)

	// Test getting non-existent num instance
	t.Run(
		"Get non-existent num instance",
		func(t *testing.T) {
			numInst := np.GetNumInstance(999)
			if numInst != nil {
				t.Errorf(
					"GetNumInstance returned non-nil for non-existent num instance: %v",
					numInst,
				)
			}
		},
	)
}

func TestFontsPart_GetFont_NilRoot(t *testing.T) {
	// Create a FontsPart without setting a root element
	fp := &FontsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/fontTable.xml",
			ContentTypeFontTable,
			nil,
			nil,
		),
	}

	// Test getting font when root is nil
	font := fp.GetFont("Calibri")
	if font != nil {
		t.Errorf(
			"GetFont returned non-nil when root is nil: %v",
			font,
		)
	}
}

func TestStylesPart_GetStyleById_NilRoot(
	t *testing.T,
) {
	// Create a StylesPart without setting a root element
	sp := &StylesPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/styles.xml",
			ContentTypeStyles,
			nil,
			nil,
		),
	}

	// Test getting style when root is nil
	style := sp.GetStyleById("Normal")
	if style != nil {
		t.Errorf(
			"GetStyleById returned non-nil when root is nil: %v",
			style,
		)
	}
}

func TestStylesPart_GetStyleByName_NilRoot(
	t *testing.T,
) {
	// Create a StylesPart without setting a root element
	sp := &StylesPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/styles.xml",
			ContentTypeStyles,
			nil,
			nil,
		),
	}

	// Test getting style when root is nil
	style := sp.GetStyleByName("Normal")
	if style != nil {
		t.Errorf(
			"GetStyleByName returned non-nil when root is nil: %v",
			style,
		)
	}
}

func TestNumberingPart_GetAbstractNum_NilRoot(
	t *testing.T,
) {
	// Create a NumberingPart without setting a root element
	np := &NumberingPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/numbering.xml",
			ContentTypeNumbering,
			nil,
			nil,
		),
	}

	// Test getting abstract num when root is nil
	absNum := np.GetAbstractNum(0)
	if absNum != nil {
		t.Errorf(
			"GetAbstractNum returned non-nil when root is nil: %v",
			absNum,
		)
	}
}

func TestNumberingPart_GetNumInstance_NilRoot(
	t *testing.T,
) {
	// Create a NumberingPart without setting a root element
	np := &NumberingPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/numbering.xml",
			ContentTypeNumbering,
			nil,
			nil,
		),
	}

	// Test getting num instance when root is nil
	numInst := np.GetNumInstance(1)
	if numInst != nil {
		t.Errorf(
			"GetNumInstance returned non-nil when root is nil: %v",
			numInst,
		)
	}
}
