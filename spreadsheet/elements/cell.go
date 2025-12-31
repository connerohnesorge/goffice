package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// CellType represents the data type of a cell.
type CellType string

const (
	// CellTypeBoolean indicates a boolean cell value.
	CellTypeBoolean CellType = "b"
	// CellTypeDate indicates a date cell value (ISO 8601 format).
	CellTypeDate CellType = "d"
	// CellTypeError indicates an error value.
	CellTypeError CellType = "e"
	// CellTypeInlineString indicates an inline string (not in shared strings).
	CellTypeInlineString CellType = "inlineStr"
	// CellTypeNumber indicates a numeric value (default if omitted).
	CellTypeNumber CellType = "n"
	// CellTypeSharedString indicates a shared string reference.
	CellTypeSharedString CellType = "s"
	// CellTypeFormulaString indicates a formula string result.
	CellTypeFormulaString CellType = "str"
)

// Cell represents a cell element (x:c) in a worksheet row.
type Cell struct {
	*openxml.CompositeElementBase
}

// CreateCell creates a new Cell element.
// Renamed from NewCell to avoid conflict with the NewCell type from the schema.
func CreateCell() *Cell {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"c",
		PrefixDefault,
	)

	return &Cell{CompositeElementBase: elem}
}

// Reference returns the cell reference (e.g., "A1"). Attribute: r.
func (c *Cell) Reference() string {
	attr, found := c.GetAttribute("r", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetReference sets the cell reference (e.g., "A1"). Attribute: r.
func (c *Cell) SetReference(ref string) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"r",
			"",
			ref,
		),
	)
}

// StyleIndex returns the style index for the cell. Attribute: s.
func (c *Cell) StyleIndex() uint32 {
	attr, found := c.GetAttribute("s", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return uint32(val)
}

// SetStyleIndex sets the style index for the cell. Attribute: s.
func (c *Cell) SetStyleIndex(index uint32) {
	if index == 0 {
		c.RemoveAttribute("s", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"s",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// DataType returns the data type of the cell. Attribute: t.
func (c *Cell) DataType() CellType {
	attr, found := c.GetAttribute("t", "")
	if !found {
		return CellTypeNumber // Default is number
	}

	return CellType(attr.Value())
}

// SetDataType sets the data type of the cell. Attribute: t.
func (c *Cell) SetDataType(cellType CellType) {
	if cellType == "" ||
		cellType == CellTypeNumber {
		c.RemoveAttribute("t", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"t",
			"",
			string(cellType),
		),
	)
}

// CellMetaIndex returns the cell metadata index. Attribute: cm.
func (c *Cell) CellMetaIndex() uint32 {
	attr, found := c.GetAttribute("cm", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return uint32(val)
}

// SetCellMetaIndex sets the cell metadata index. Attribute: cm.
func (c *Cell) SetCellMetaIndex(index uint32) {
	if index == 0 {
		c.RemoveAttribute("cm", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"cm",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ValueMetaIndex returns the value metadata index. Attribute: vm.
func (c *Cell) ValueMetaIndex() uint32 {
	attr, found := c.GetAttribute("vm", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return uint32(val)
}

// SetValueMetaIndex sets the value metadata index. Attribute: vm.
func (c *Cell) SetValueMetaIndex(index uint32) {
	if index == 0 {
		c.RemoveAttribute("vm", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"vm",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ShowPhonetic returns whether phonetic information should be shown.
// Attribute: ph.
func (c *Cell) ShowPhonetic() bool {
	attr, found := c.GetAttribute("ph", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowPhonetic sets whether phonetic information should be shown.
// Attribute: ph.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Cell) SetShowPhonetic(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"ph",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("ph", "")
	}
}

// CellValue returns the cell value element (x:v), or nil if not present.
//
//nolint:revive // cognitive-complexity: necessary complexity for XML reload handling
func (c *Cell) CellValue() *CellValue {
	elem := c.GetElement("v", NamespaceSML)
	if elem == nil {
		return nil
	}
	if cv, ok := elem.(*CellValue); ok {
		return cv
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &CellValue{LeafElementBase: leaf}
	}
	// Handle generic elements created during XML reload
	// Convert CompositeElement to CellValue if it matches
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		if comp.LocalName() == "v" &&
			comp.NamespaceURI() == NamespaceSML {
			// Create a proper CellValue element from the CompositeElement
			// Get the inner text content from the special #text child node
			innerText := ""
			if comp.FirstChild() != nil {
				// Check if it's a text node (local name "#text")
				//nolint:revive // max-control-nesting: necessary for XML reload fix
				if comp.FirstChild().
					LocalName() ==
					"#text" {
					if leafChild, ok := comp.FirstChild().(openxml.LeafElement); ok {
						innerText = leafChild.InnerText()
					}
				}
			}
			// Create new CellValue element and replace the old one
			cv := NewCellValueWithText(innerText)
			c.ReplaceChild(cv, comp)

			return cv
		}
	}

	return nil
}

// GetOrCreateCellValue returns the cell value element, creating it if needed.
func (c *Cell) GetOrCreateCellValue() *CellValue {
	cv := c.CellValue()
	if cv != nil {
		return cv
	}
	cv = NewCellValue()
	c.AppendChild(cv)

	return cv
}

// CellFormula returns the cell formula element (x:f), or nil if not present.
func (c *Cell) CellFormula() *CellFormula {
	elem := c.GetElement("f", NamespaceSML)
	if elem == nil {
		return nil
	}
	if cf, ok := elem.(*CellFormula); ok {
		return cf
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &CellFormula{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateCellFormula returns the cell formula element,
// creating it if needed.
func (c *Cell) GetOrCreateCellFormula() *CellFormula {
	cf := c.CellFormula()
	if cf != nil {
		return cf
	}
	cf = NewCellFormula()
	// Formula should come before value
	cv := c.CellValue()
	if cv != nil {
		c.InsertBefore(cf, cv)
	} else {
		c.AppendChild(cf)
	}

	return cf
}

// InlineString returns the inline string element (x:is), or nil if not present.
func (c *Cell) InlineString() *InlineString {
	elem := c.GetElement("is", NamespaceSML)
	if elem == nil {
		return nil
	}
	if is, ok := elem.(*InlineString); ok {
		return is
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &InlineString{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateInlineString returns the inline string element,
// creating it if needed.
func (c *Cell) GetOrCreateInlineString() *InlineString {
	is := c.InlineString()
	if is != nil {
		return is
	}
	is = NewInlineString()
	c.AppendChild(is)

	return is
}

// Value returns the cell value as a string.
// This is a convenience method that returns the text content
// of the value element.
func (c *Cell) Value() string {
	cv := c.CellValue()
	if cv == nil {
		return ""
	}

	return cv.Value()
}

// SetValue sets the cell value as a string.
// This is a convenience method.
func (c *Cell) SetValue(value string) {
	cv := c.GetOrCreateCellValue()
	cv.SetValue(value)
}

// Clone creates a deep copy of this Cell element.
func (c *Cell) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Cell{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Cell element.
func (c *Cell) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &Cell{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
