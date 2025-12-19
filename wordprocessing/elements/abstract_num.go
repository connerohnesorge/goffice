package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// MultiLevelTypeValue represents multi-level numbering type values.
type MultiLevelTypeValue string

const (
	// MultiLevelSingleLevel indicates single-level numbering.
	MultiLevelSingleLevel MultiLevelTypeValue = "singleLevel"
	// MultiLevelMultilevel indicates multi-level numbering.
	MultiLevelMultilevel MultiLevelTypeValue = "multilevel"
	// MultiLevelHybridMultilevel indicates hybrid multi-level numbering.
	MultiLevelHybridMultilevel MultiLevelTypeValue = "hybridMultilevel"
)

// AbstractNum represents an abstract numbering definition (w:abstractNum).
type AbstractNum struct {
	*openxml.CompositeElementBase
}

// NewAbstractNum creates a new AbstractNum element.
func NewAbstractNum() *AbstractNum {
	elem := openxml.NewCompositeElement(NamespaceWML, "abstractNum", PrefixW)
	return &AbstractNum{CompositeElementBase: elem}
}

// AbstractNumId returns the abstract numbering definition ID.
func (an *AbstractNum) AbstractNumId() int {
	attr, found := an.GetAttribute("abstractNumId", NamespaceWML)
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}
	return val
}

// SetAbstractNumId sets the abstract numbering definition ID.
func (an *AbstractNum) SetAbstractNumId(id int) {
	an.SetAttribute(openxml.NewAttribute(NamespaceWML, "abstractNumId", PrefixW, strconv.Itoa(id)))
}

// MultiLevelType returns the multi-level numbering type.
func (an *AbstractNum) MultiLevelType() MultiLevelTypeValue {
	elem := an.GetElement("multiLevelType", NamespaceWML)
	if elem == nil {
		return MultiLevelSingleLevel
	}
	attr, found := elem.GetAttribute("val", NamespaceWML)
	if !found {
		return MultiLevelSingleLevel
	}
	return MultiLevelTypeValue(attr.Value())
}

// SetMultiLevelType sets the multi-level numbering type.
func (an *AbstractNum) SetMultiLevelType(t MultiLevelTypeValue) {
	elem := an.getOrCreateElement("multiLevelType")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, string(t)))
}

// NumberingStyleLink returns the linked numbering style ID.
func (an *AbstractNum) NumberingStyleLink() string {
	elem := an.GetElement("numStyleLink", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute("val", NamespaceWML)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetNumberingStyleLink sets the linked numbering style ID.
func (an *AbstractNum) SetNumberingStyleLink(styleId string) {
	if styleId == "" {
		an.removeElement("numStyleLink")
		return
	}
	elem := an.getOrCreateElement("numStyleLink")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, styleId))
}

// StyleLink returns the linked style ID.
func (an *AbstractNum) StyleLink() string {
	elem := an.GetElement("styleLink", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute("val", NamespaceWML)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetStyleLink sets the linked style ID.
func (an *AbstractNum) SetStyleLink(styleId string) {
	if styleId == "" {
		an.removeElement("styleLink")
		return
	}
	elem := an.getOrCreateElement("styleLink")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, styleId))
}

// Levels returns an iterator over all Level elements.
func (an *AbstractNum) Levels() iter.Seq[*Level] {
	return func(yield func(*Level) bool) {
		for child := range an.Children() {
			if child.LocalName() == "lvl" && child.NamespaceURI() == NamespaceWML {
				var lvl *Level
				if level, ok := child.(*Level); ok {
					lvl = level
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					lvl = &Level{CompositeElementBase: comp}
				}
				if lvl != nil && !yield(lvl) {
					return
				}
			}
		}
	}
}

// GetLevel returns the level at the specified index (0-8), or nil if not found.
func (an *AbstractNum) GetLevel(index int) *Level {
	for lvl := range an.Levels() {
		if lvl.LevelIndex() == index {
			return lvl
		}
	}
	return nil
}

// AddLevel adds a new level at the specified index.
func (an *AbstractNum) AddLevel(index int) *Level {
	lvl := NewLevel(index)
	an.AppendChild(lvl)
	return lvl
}

// GetOrCreateLevel returns the level at the specified index, creating if needed.
func (an *AbstractNum) GetOrCreateLevel(index int) *Level {
	lvl := an.GetLevel(index)
	if lvl != nil {
		return lvl
	}
	return an.AddLevel(index)
}

func (an *AbstractNum) getOrCreateElement(name string) openxml.Element {
	elem := an.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(NamespaceWML, name, PrefixW)
	an.AppendChild(newElem)
	return newElem
}

func (an *AbstractNum) removeElement(name string) {
	elem := an.GetElement(name, NamespaceWML)
	if elem != nil {
		an.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this AbstractNum element.
func (an *AbstractNum) Clone() openxml.Element {
	return &AbstractNum{
		CompositeElementBase: an.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this AbstractNum element.
func (an *AbstractNum) CloneNode(deep bool) openxml.Element {
	return &AbstractNum{
		CompositeElementBase: an.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// Factory functions for common list types

// NewBulletList creates a bullet list abstract numbering definition.
func NewBulletList(bulletChar, fontName string) *AbstractNum {
	an := NewAbstractNum()
	an.SetMultiLevelType(MultiLevelHybridMultilevel)

	// Create 9 levels with the same bullet
	for i := 0; i < 9; i++ {
		lvl := an.AddLevel(i)
		lvl.SetStart(1)
		lvl.SetNumberFormat(NumberFormatBullet)
		lvl.SetLevelText(bulletChar)
		lvl.SetLevelJustification(JustificationLeft)

		// Set indentation: each level indented 720 twips (0.5 inch) more
		indent := (i + 1) * 720
		hanging := 360
		lvl.SetIndentation(indent, hanging)

		// Set the bullet font
		if fontName != "" {
			lvl.GetOrCreateNumberingRunProperties().SetFont(fontName)
		}
	}

	return an
}

// NewStandardBulletList creates a standard bullet list with common bullet characters.
func NewStandardBulletList() *AbstractNum {
	bullets := []struct {
		char string
		font string
	}{
		{"\uF0B7", "Symbol"},       // Level 0: solid bullet
		{"o", "Courier New"},       // Level 1: open bullet
		{"\uF0A7", "Wingdings"},    // Level 2: square bullet
		{"\uF0B7", "Symbol"},       // Level 3: solid bullet
		{"o", "Courier New"},       // Level 4: open bullet
		{"\uF0A7", "Wingdings"},    // Level 5: square bullet
		{"\uF0B7", "Symbol"},       // Level 6: solid bullet
		{"o", "Courier New"},       // Level 7: open bullet
		{"\uF0A7", "Wingdings"},    // Level 8: square bullet
	}

	an := NewAbstractNum()
	an.SetMultiLevelType(MultiLevelHybridMultilevel)

	for i := 0; i < 9; i++ {
		lvl := an.AddLevel(i)
		lvl.SetStart(1)
		lvl.SetNumberFormat(NumberFormatBullet)
		lvl.SetLevelText(bullets[i].char)
		lvl.SetLevelJustification(JustificationLeft)

		// Set indentation
		indent := (i + 1) * 720
		hanging := 360
		lvl.SetIndentation(indent, hanging)

		// Set the bullet font
		lvl.GetOrCreateNumberingRunProperties().SetFont(bullets[i].font)
	}

	return an
}

// NewDecimalList creates a decimal numbered list (1, 2, 3...).
func NewDecimalList() *AbstractNum {
	an := NewAbstractNum()
	an.SetMultiLevelType(MultiLevelHybridMultilevel)

	for i := 0; i < 9; i++ {
		lvl := an.AddLevel(i)
		lvl.SetStart(1)
		lvl.SetNumberFormat(NumberFormatDecimal)
		lvl.SetLevelText("%" + strconv.Itoa(i+1) + ".")
		lvl.SetLevelJustification(JustificationLeft)

		// Set indentation
		indent := (i + 1) * 720
		hanging := 360
		lvl.SetIndentation(indent, hanging)
	}

	return an
}

// NewAlphabeticList creates an alphabetic list (a, b, c... or A, B, C...).
func NewAlphabeticList(lowercase bool) *AbstractNum {
	an := NewAbstractNum()
	an.SetMultiLevelType(MultiLevelHybridMultilevel)

	format := NumberFormatUpperLetter
	if lowercase {
		format = NumberFormatLowerLetter
	}

	for i := 0; i < 9; i++ {
		lvl := an.AddLevel(i)
		lvl.SetStart(1)
		lvl.SetNumberFormat(format)
		lvl.SetLevelText("%" + strconv.Itoa(i+1) + ".")
		lvl.SetLevelJustification(JustificationLeft)

		// Set indentation
		indent := (i + 1) * 720
		hanging := 360
		lvl.SetIndentation(indent, hanging)
	}

	return an
}

// NewRomanNumeralList creates a Roman numeral list (i, ii, iii... or I, II, III...).
func NewRomanNumeralList(lowercase bool) *AbstractNum {
	an := NewAbstractNum()
	an.SetMultiLevelType(MultiLevelHybridMultilevel)

	format := NumberFormatUpperRoman
	if lowercase {
		format = NumberFormatLowerRoman
	}

	for i := 0; i < 9; i++ {
		lvl := an.AddLevel(i)
		lvl.SetStart(1)
		lvl.SetNumberFormat(format)
		lvl.SetLevelText("%" + strconv.Itoa(i+1) + ".")
		lvl.SetLevelJustification(JustificationLeft)

		// Set indentation
		indent := (i + 1) * 720
		hanging := 360
		lvl.SetIndentation(indent, hanging)
	}

	return an
}

// NewOutlineList creates an outline numbering list (1, 1.1, 1.1.1...).
func NewOutlineList() *AbstractNum {
	an := NewAbstractNum()
	an.SetMultiLevelType(MultiLevelMultilevel)

	for i := 0; i < 9; i++ {
		lvl := an.AddLevel(i)
		lvl.SetStart(1)
		lvl.SetNumberFormat(NumberFormatDecimal)

		// Build level text: %1.%2.%3...
		text := ""
		for j := 0; j <= i; j++ {
			if j > 0 {
				text += "."
			}
			text += "%" + strconv.Itoa(j+1)
		}
		lvl.SetLevelText(text)
		lvl.SetLevelJustification(JustificationLeft)

		// Set indentation
		indent := (i + 1) * 720
		hanging := 360 + (i * 180) // Increase hanging for longer numbers
		lvl.SetIndentation(indent, hanging)
	}

	return an
}
