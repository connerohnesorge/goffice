package elements

//revive:disable:file-length-limit many data bar properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// DataBarDirection represents the direction of a data bar.
type DataBarDirection string

const (
	// DataBarDirectionContext indicates direction based on context (default).
	DataBarDirectionContext DataBarDirection = "context"
	// DataBarDirectionLTR indicates left-to-right direction.
	DataBarDirectionLTR DataBarDirection = "leftToRight"
	// DataBarDirectionRTL indicates right-to-left direction.
	DataBarDirectionRTL DataBarDirection = "rightToLeft"
)

// DataBarAxisPosition represents the axis position for a data bar.
type DataBarAxisPosition string

const (
	// DataBarAxisPositionAutomatic indicates automatic axis position (default).
	DataBarAxisPositionAutomatic DataBarAxisPosition = "automatic"
	// DataBarAxisPositionMiddle indicates axis in the middle of the cell.
	DataBarAxisPositionMiddle DataBarAxisPosition = "middle"
	// DataBarAxisPositionNone indicates no axis is shown.
	DataBarAxisPositionNone DataBarAxisPosition = "none"
)

// DataBar represents a data bar element (x:dataBar) for conditional formatting.
// A data bar displays a colored bar whose length represents the cell value.
type DataBar struct {
	*openxml.CompositeElementBase
}

// NewDataBar creates a new DataBar element.
func NewDataBar() *DataBar {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dataBar",
		PrefixDefault,
	)

	return &DataBar{CompositeElementBase: elem}
}

// MinLength returns the minimum length of the data bar (0-100).
// Attribute: minLength. Default is 10.
func (d *DataBar) MinLength() uint32 {
	attr, found := d.GetAttribute("minLength", "")
	if !found {
		return 10 //nolint:revive // add-constant: default min length
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size 32
	)

	return uint32(val)
}

// SetMinLength sets the minimum length of the data bar. Attribute: minLength.
func (d *DataBar) SetMinLength(length uint32) {
	if length == 10 { //nolint:revive // add-constant: default min length
		d.RemoveAttribute("minLength", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"minLength",
			"",
			strconv.FormatUint(
				uint64(length),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// MaxLength returns the maximum length of the data bar (0-100).
// Attribute: maxLength.
// Default is 90.
func (d *DataBar) MaxLength() uint32 {
	attr, found := d.GetAttribute("maxLength", "")
	if !found {
		return 90 //nolint:revive // add-constant: default max length
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetMaxLength sets the maximum length of the data bar. Attribute: maxLength.
func (d *DataBar) SetMaxLength(length uint32) {
	if length == 90 { //nolint:revive // add-constant: default max length
		d.RemoveAttribute("maxLength", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"maxLength",
			"",
			strconv.FormatUint(
				uint64(length),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ShowValue returns whether the cell value is shown. Attribute: showValue.
// Default is true.
func (d *DataBar) ShowValue() bool {
	attr, found := d.GetAttribute("showValue", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowValue sets whether the cell value is shown. Attribute: showValue.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataBar) SetShowValue(value bool) {
	if value {
		d.RemoveAttribute("showValue", "")
	} else {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"showValue",
				"",
				attrValueFalse,
			),
		)
	}
}

// Gradient returns whether the data bar has a gradient fill.
// Attribute: gradient. Default is true.
func (d *DataBar) Gradient() bool {
	attr, found := d.GetAttribute("gradient", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetGradient sets whether the data bar has a gradient fill. Attribute: gradient.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataBar) SetGradient(value bool) {
	if value {
		d.RemoveAttribute("gradient", "")
	} else {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"gradient",
				"",
				attrValueFalse,
			),
		)
	}
}

// Direction returns the direction of the data bar. Attribute: direction.
func (d *DataBar) Direction() DataBarDirection {
	attr, found := d.GetAttribute("direction", "")
	if !found {
		return DataBarDirectionContext
	}

	return DataBarDirection(attr.Value())
}

// SetDirection sets the direction of the data bar. Attribute: direction.
func (d *DataBar) SetDirection(
	direction DataBarDirection,
) {
	if direction == "" ||
		direction == DataBarDirectionContext {
		d.RemoveAttribute("direction", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"direction",
			"",
			string(direction),
		),
	)
}

// NegativeBarColorSameAsPositive returns whether negative bars use the same
// color as positive bars. Attribute: negativeBarColorSameAsPositive.
// Default is true.
func (d *DataBar) NegativeBarColorSameAsPositive() bool {
	attr, found := d.GetAttribute(
		"negativeBarColorSameAsPositive",
		"",
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetNegativeBarColorSameAsPositive sets whether negative bars use the same color
// as positive bars. Attribute: negativeBarColorSameAsPositive.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataBar) SetNegativeBarColorSameAsPositive(
	value bool,
) {
	if value {
		d.RemoveAttribute(
			"negativeBarColorSameAsPositive",
			"",
		)
	} else {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"negativeBarColorSameAsPositive",
				"",
				attrValueFalse,
			),
		)
	}
}

// NegativeBarBorderColorSameAsPositive returns whether negative bar borders use
// the same color as positive bar borders.
// Attribute: negativeBarBorderColorSameAsPositive.
// Default is true.
func (d *DataBar) NegativeBarBorderColorSameAsPositive() bool {
	attr, found := d.GetAttribute(
		"negativeBarBorderColorSameAsPositive",
		"",
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetNegativeBarBorderColorSameAsPositive sets whether negative bar borders use
// the same color as positive bar borders.
// Attribute: negativeBarBorderColorSameAsPositive.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataBar) SetNegativeBarBorderColorSameAsPositive(
	value bool,
) {
	if value {
		d.RemoveAttribute(
			"negativeBarBorderColorSameAsPositive",
			"",
		)
	} else {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"negativeBarBorderColorSameAsPositive",
				"",
				attrValueFalse,
			),
		)
	}
}

// AxisPosition returns the axis position for the data bar.
// Attribute: axisPosition.
func (d *DataBar) AxisPosition() DataBarAxisPosition {
	attr, found := d.GetAttribute(
		"axisPosition",
		"",
	)
	if !found {
		return DataBarAxisPositionAutomatic
	}

	return DataBarAxisPosition(attr.Value())
}

// SetAxisPosition sets the axis position for the data bar.
// Attribute: axisPosition.
func (d *DataBar) SetAxisPosition(
	position DataBarAxisPosition,
) {
	if position == "" ||
		position == DataBarAxisPositionAutomatic {
		d.RemoveAttribute("axisPosition", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"axisPosition",
			"",
			string(position),
		),
	)
}

// Cfvos returns an iterator over all Cfvo elements in this data bar.
func (d *DataBar) Cfvos() iter.Seq[*Cfvo] {
	return func(yield func(*Cfvo) bool) {
		for child := range d.Children() {
			if child.LocalName() != "cfvo" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var cfvo *Cfvo
			if c, ok := child.(*Cfvo); ok {
				cfvo = c
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				cfvo = &Cfvo{LeafElementBase: leaf}
			}
			if cfvo != nil && !yield(cfvo) {
				return
			}
		}
	}
}

// AddCfvo adds a new Cfvo element to the data bar.
func (d *DataBar) AddCfvo() *Cfvo {
	cfvo := NewCfvo()
	// Insert before colors
	var insertBefore openxml.Element
	for child := range d.Children() {
		if child.LocalName() == "color" &&
			child.NamespaceURI() == NamespaceSML {
			insertBefore = child

			break
		}
	}
	if insertBefore != nil {
		d.InsertBefore(cfvo, insertBefore)
	} else {
		d.AppendChild(cfvo)
	}

	return cfvo
}

// Color returns the bar fill color element, or nil if not present.
func (d *DataBar) Color() *Color {
	elem := d.GetElement("color", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Color); ok {
		return c
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Color{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateColor returns the bar fill color, creating it if needed.
func (d *DataBar) GetOrCreateColor() *Color {
	c := d.Color()
	if c != nil {
		return c
	}
	c = NewColor()
	d.AppendChild(c)

	return c
}

// NegativeColor returns the negative bar color element, or nil if not present.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) NegativeColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "negativeColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}

	return nil
}

// GetOrCreateNegativeColor returns the negative bar color, creating it if needed.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) GetOrCreateNegativeColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "negativeColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}
	// Create as negativeColor element
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"negativeColor",
		PrefixDefault,
	)
	c := &Color{LeafElementBase: elem}
	d.AppendChild(c)

	return c
}

// BorderColor returns the bar border color element, or nil if not present.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) BorderColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "borderColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}

	return nil
}

// GetOrCreateBorderColor returns the bar border color, creating it if needed.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) GetOrCreateBorderColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "borderColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"borderColor",
		PrefixDefault,
	)
	c := &Color{LeafElementBase: elem}
	d.AppendChild(c)

	return c
}

// NegativeBorderColor returns the negative bar border color,
// or nil if not present.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) NegativeBorderColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "negativeBorderColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}

	return nil
}

// GetOrCreateNegativeBorderColor returns the negative bar border color,
// creating it if needed.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) GetOrCreateNegativeBorderColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "negativeBorderColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"negativeBorderColor",
		PrefixDefault,
	)
	c := &Color{LeafElementBase: elem}
	d.AppendChild(c)

	return c
}

// AxisColor returns the axis color element, or nil if not present.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) AxisColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "axisColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}

	return nil
}

// GetOrCreateAxisColor returns the axis color, creating it if needed.
//
//nolint:revive // early-return: nested checks for type assertion clarity
func (d *DataBar) GetOrCreateAxisColor() *Color {
	for child := range d.Children() {
		if child.LocalName() == "axisColor" &&
			child.NamespaceURI() == NamespaceSML {
			if leaf, ok := child.(*openxml.LeafElementBase); ok {
				return &Color{
					LeafElementBase: leaf,
				}
			}
		}
	}
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"axisColor",
		PrefixDefault,
	)
	c := &Color{LeafElementBase: elem}
	d.AppendChild(c)

	return c
}

// SetDataBar sets up a data bar with min and max values and a bar color.
//
//nolint:revive // enforce-repeated-arg-type-style: API design for consistency
func (d *DataBar) SetDataBar(
	minType CfvoType, minVal string,
	maxType CfvoType, maxVal string,
	barColor string,
) {
	// Clear existing
	d.RemoveAllChildren()

	// Add cfvo elements
	minCfvo := d.AddCfvo()
	minCfvo.SetType(minType)
	if minVal != "" {
		minCfvo.SetVal(minVal)
	}

	maxCfvo := d.AddCfvo()
	maxCfvo.SetType(maxType)
	if maxVal != "" {
		maxCfvo.SetVal(maxVal)
	}

	// Add color
	color := d.GetOrCreateColor()
	color.SetRGB(barColor)
}

// Clone creates a deep copy of this DataBar element.
func (d *DataBar) Clone() openxml.Element {
	cloned := d.CompositeElementBase.Clone()

	return &DataBar{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DataBar element.
func (d *DataBar) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.CompositeElementBase.CloneNode(
		deep,
	)

	return &DataBar{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
