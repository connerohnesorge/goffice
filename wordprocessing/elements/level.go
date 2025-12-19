package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// NumberFormatValue represents number format values.
type NumberFormatValue string

const (
	// NumberFormatDecimal is decimal numbering (1, 2, 3...).
	NumberFormatDecimal NumberFormatValue = "decimal"
	// NumberFormatLowerLetter is lowercase letter numbering (a, b, c...).
	NumberFormatLowerLetter NumberFormatValue = "lowerLetter"
	// NumberFormatUpperLetter is uppercase letter numbering (A, B, C...).
	NumberFormatUpperLetter NumberFormatValue = "upperLetter"
	// NumberFormatLowerRoman is lowercase Roman numeral numbering (i, ii, iii...).
	NumberFormatLowerRoman NumberFormatValue = "lowerRoman"
	// NumberFormatUpperRoman is uppercase Roman numeral numbering (I, II, III...).
	NumberFormatUpperRoman NumberFormatValue = "upperRoman"
	// NumberFormatBullet indicates bullet formatting.
	NumberFormatBullet NumberFormatValue = "bullet"
	// NumberFormatNone indicates no numbering.
	NumberFormatNone NumberFormatValue = "none"
	// NumberFormatCardinalText is cardinal text (one, two, three...).
	NumberFormatCardinalText NumberFormatValue = "cardinalText"
	// NumberFormatOrdinalText is ordinal text (first, second, third...).
	NumberFormatOrdinalText NumberFormatValue = "ordinalText"
	// NumberFormatOrdinal is ordinal numbering (1st, 2nd, 3rd...).
	NumberFormatOrdinal NumberFormatValue = "ordinal"
	// NumberFormatDecimalZero is decimal with leading zeros (01, 02, 03...).
	NumberFormatDecimalZero NumberFormatValue = "decimalZero"
	// NumberFormatChicago is Chicago Manual of Style footnote numbering.
	NumberFormatChicago NumberFormatValue = "chicago"
)

// LevelSuffixValue represents level suffix values.
type LevelSuffixValue string

const (
	// LevelSuffixTab adds a tab after the number.
	LevelSuffixTab LevelSuffixValue = "tab"
	// LevelSuffixSpace adds a space after the number.
	LevelSuffixSpace LevelSuffixValue = "space"
	// LevelSuffixNothing adds nothing after the number.
	LevelSuffixNothing LevelSuffixValue = "nothing"
)

// Level represents a level definition within an abstract numbering (w:lvl).
type Level struct {
	*openxml.CompositeElementBase
}

// NewLevel creates a new Level element at the specified index (0-8).
func NewLevel(index int) *Level {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"lvl",
		PrefixW,
	)
	lvl := &Level{CompositeElementBase: elem}
	lvl.SetLevelIndex(index)

	return lvl
}

// LevelIndex returns the level index (0-8).
func (l *Level) LevelIndex() int {
	attr, found := l.GetAttribute(
		"ilvl",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}

	return val
}

// SetLevelIndex sets the level index (0-8).
func (l *Level) SetLevelIndex(index int) {
	l.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"ilvl",
			PrefixW,
			strconv.Itoa(index),
		),
	)
}

// Start returns the starting number for this level.
func (l *Level) Start() int {
	elem := l.GetElement("start", NamespaceWML)
	if elem == nil {
		return 1
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 1
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 1
	}

	return val
}

// SetStart sets the starting number for this level.
func (l *Level) SetStart(start int) {
	elem := l.getOrCreateElement("start")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(start),
		),
	)
}

// NumberFormat returns the number format for this level.
func (l *Level) NumberFormat() NumberFormatValue {
	elem := l.GetElement("numFmt", NamespaceWML)
	if elem == nil {
		return NumberFormatDecimal
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return NumberFormatDecimal
	}

	return NumberFormatValue(attr.Value())
}

// SetNumberFormat sets the number format for this level.
func (l *Level) SetNumberFormat(
	f NumberFormatValue,
) {
	elem := l.getOrCreateElement("numFmt")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(f),
		),
	)
}

// LevelText returns the level text pattern (e.g., "%1.", "%1.%2").
func (l *Level) LevelText() string {
	elem := l.GetElement("lvlText", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLevelText sets the level text pattern.
func (l *Level) SetLevelText(text string) {
	elem := l.getOrCreateElement("lvlText")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			text,
		),
	)
}

// LevelJustification returns the justification for this level.
func (l *Level) LevelJustification() JustificationValue {
	elem := l.GetElement("lvlJc", NamespaceWML)
	if elem == nil {
		return JustificationLeft
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return JustificationLeft
	}

	return JustificationValue(attr.Value())
}

// SetLevelJustification sets the justification for this level.
func (l *Level) SetLevelJustification(
	j JustificationValue,
) {
	elem := l.getOrCreateElement("lvlJc")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(j),
		),
	)
}

// LevelRestart returns the level at which this level restarts, or -1 if not set.
func (l *Level) LevelRestart() int {
	elem := l.GetElement(
		"lvlRestart",
		NamespaceWML,
	)
	if elem == nil {
		return -1
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return -1
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return -1
	}

	return val
}

// SetLevelRestart sets the level at which this level restarts.
func (l *Level) SetLevelRestart(level int) {
	if level < 0 {
		l.removeElement("lvlRestart")

		return
	}
	elem := l.getOrCreateElement("lvlRestart")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(level),
		),
	)
}

// LevelSuffix returns the suffix after the number.
func (l *Level) LevelSuffix() LevelSuffixValue {
	elem := l.GetElement("suff", NamespaceWML)
	if elem == nil {
		return LevelSuffixTab
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return LevelSuffixTab
	}

	return LevelSuffixValue(attr.Value())
}

// SetLevelSuffix sets the suffix after the number.
func (l *Level) SetLevelSuffix(
	suffix LevelSuffixValue,
) {
	if suffix == LevelSuffixTab {
		l.removeElement("suff")

		return
	}
	elem := l.getOrCreateElement("suff")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(suffix),
		),
	)
}

// IsLegalNumbering returns whether legal numbering format is used.
func (l *Level) IsLegalNumbering() bool {
	elem := l.GetElement("isLgl", NamespaceWML)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if found {
		val := attr.Value()

		return val != "false" && val != "0" &&
			val != "off"
	}

	return true
}

// SetLegalNumbering sets whether legal numbering format is used.
func (l *Level) SetLegalNumbering(b bool) {
	if b {
		l.getOrCreateElement("isLgl")
	} else {
		l.removeElement("isLgl")
	}
}

// ParagraphProperties returns the paragraph properties for this level.
func (l *Level) ParagraphProperties() *LevelParagraphProperties {
	elem := l.GetElement("pPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pp, ok := elem.(*LevelParagraphProperties); ok {
		return pp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &LevelParagraphProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateParagraphProperties returns paragraph properties, creating if needed.
func (l *Level) GetOrCreateParagraphProperties() *LevelParagraphProperties {
	pp := l.ParagraphProperties()
	if pp != nil {
		return pp
	}
	pp = NewLevelParagraphProperties()
	l.AppendChild(pp)

	return pp
}

// RunProperties returns the run properties for the numbering symbol.
func (l *Level) RunProperties() *NumberingRunProperties {
	elem := l.GetElement("rPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*NumberingRunProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NumberingRunProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNumberingRunProperties returns run properties, creating if needed.
func (l *Level) GetOrCreateNumberingRunProperties() *NumberingRunProperties {
	rp := l.RunProperties()
	if rp != nil {
		return rp
	}
	rp = NewNumberingRunProperties()
	l.AppendChild(rp)

	return rp
}

// SetIndentation is a convenience method to set the indentation for this level.
func (l *Level) SetIndentation(
	left, hanging int,
) {
	pp := l.GetOrCreateParagraphProperties()
	pp.SetIndentation(left, hanging)
}

func (l *Level) getOrCreateElement(
	name string,
) openxml.Element {
	elem := l.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	l.AppendChild(newElem)

	return newElem
}

func (l *Level) removeElement(name string) {
	elem := l.GetElement(name, NamespaceWML)
	if elem != nil {
		l.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Level element.
func (l *Level) Clone() openxml.Element {
	return &Level{
		CompositeElementBase: l.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Level element.
func (l *Level) CloneNode(
	deep bool,
) openxml.Element {
	return &Level{
		CompositeElementBase: l.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// LevelParagraphProperties represents paragraph properties within a level (w:pPr).
type LevelParagraphProperties struct {
	*openxml.CompositeElementBase
}

// NewLevelParagraphProperties creates a new LevelParagraphProperties element.
func NewLevelParagraphProperties() *LevelParagraphProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pPr",
		PrefixW,
	)

	return &LevelParagraphProperties{
		CompositeElementBase: elem,
	}
}

// SetIndentation sets the indentation for the level.
func (pp *LevelParagraphProperties) SetIndentation(
	left, hanging int,
) {
	ind := pp.GetElement("ind", NamespaceWML)
	if ind == nil {
		ind = openxml.NewCompositeElement(
			NamespaceWML,
			"ind",
			PrefixW,
		)
		pp.AppendChild(ind)
	}
	ind.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"left",
			PrefixW,
			strconv.Itoa(left),
		),
	)
	ind.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hanging",
			PrefixW,
			strconv.Itoa(hanging),
		),
	)
}

// Left returns the left indentation in twips.
func (pp *LevelParagraphProperties) Left() int {
	ind := pp.GetElement("ind", NamespaceWML)
	if ind == nil {
		return 0
	}
	attr, found := ind.GetAttribute(
		"left",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Hanging returns the hanging indentation in twips.
func (pp *LevelParagraphProperties) Hanging() int {
	ind := pp.GetElement("ind", NamespaceWML)
	if ind == nil {
		return 0
	}
	attr, found := ind.GetAttribute(
		"hanging",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Clone creates a deep copy of this LevelParagraphProperties element.
func (pp *LevelParagraphProperties) Clone() openxml.Element {
	return &LevelParagraphProperties{
		CompositeElementBase: pp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// NumberingRunProperties represents run properties for the numbering symbol (w:rPr).
type NumberingRunProperties struct {
	*openxml.CompositeElementBase
}

// NewNumberingRunProperties creates a new NumberingRunProperties element.
func NewNumberingRunProperties() *NumberingRunProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"rPr",
		PrefixW,
	)

	return &NumberingRunProperties{
		CompositeElementBase: elem,
	}
}

// SetFont sets the font name for the numbering symbol.
func (rp *NumberingRunProperties) SetFont(
	fontName string,
) {
	rf := rp.getOrCreateElement("rFonts")
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"ascii",
			PrefixW,
			fontName,
		),
	)
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hAnsi",
			PrefixW,
			fontName,
		),
	)
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hint",
			PrefixW,
			"default",
		),
	)
}

// SetFontSize sets the font size in half-points.
func (rp *NumberingRunProperties) SetFontSize(
	halfPoints int,
) {
	elem := rp.getOrCreateElement("sz")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(halfPoints),
		),
	)
}

// SetBold sets bold formatting.
func (rp *NumberingRunProperties) SetBold(
	b bool,
) {
	if b {
		rp.getOrCreateElement("b")
	} else {
		rp.removeElement("b")
	}
}

// SetItalic sets italic formatting.
func (rp *NumberingRunProperties) SetItalic(
	b bool,
) {
	if b {
		rp.getOrCreateElement("i")
	} else {
		rp.removeElement("i")
	}
}

// SetColor sets the text color.
func (rp *NumberingRunProperties) SetColor(
	hex string,
) {
	if hex == "" {
		rp.removeElement("color")

		return
	}
	elem := rp.getOrCreateElement("color")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			hex,
		),
	)
}

func (rp *NumberingRunProperties) getOrCreateElement(
	name string,
) openxml.Element {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	rp.AppendChild(newElem)

	return newElem
}

func (rp *NumberingRunProperties) removeElement(
	name string,
) {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		rp.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this NumberingRunProperties element.
func (rp *NumberingRunProperties) Clone() openxml.Element {
	return &NumberingRunProperties{
		CompositeElementBase: rp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
