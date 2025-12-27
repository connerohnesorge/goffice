package elements

//revive:disable:file-length-limit many shape properties

import (
	"fmt"

	"github.com/connerohnesorge/goffice/openxml"
)

// Shape represents the shape element (xdr:sp).
// This is used for drawing shapes like rectangles, circles, etc.
type Shape struct {
	*openxml.CompositeElementBase
}

// NewShape creates a new Shape element.
func NewShape() *Shape {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"sp",
		PrefixXDR,
	)

	return &Shape{CompositeElementBase: elem}
}

// Macro returns the macro attribute value.
func (s *Shape) Macro() string {
	attr, found := s.GetAttribute("macro", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMacro sets the macro attribute.
func (s *Shape) SetMacro(macro string) {
	if macro == "" {
		s.RemoveAttribute("macro", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"macro",
			"",
			macro,
		),
	)
}

// TextLink returns the text link attribute value.
func (s *Shape) TextLink() string {
	attr, found := s.GetAttribute("textlink", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTextLink sets the text link attribute.
func (s *Shape) SetTextLink(textLink string) {
	if textLink == "" {
		s.RemoveAttribute("textlink", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"textlink",
			"",
			textLink,
		),
	)
}

// FLocksText returns whether text is locked.
func (s *Shape) FLocksText() bool {
	attr, found := s.GetAttribute(
		"fLocksText",
		"",
	)
	if !found {
		return true // Default
	}

	return attr.Value() != attrValueFalse &&
		attr.Value() != attrValueZero
}

// SetFLocksText sets whether text is locked.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *Shape) SetFLocksText(value bool) {
	if value {
		s.RemoveAttribute(
			"fLocksText",
			"",
		) // true is default

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"fLocksText",
			"",
			attrValueFalse,
		),
	)
}

// FPublished returns whether the shape is published.
func (s *Shape) FPublished() bool {
	attr, found := s.GetAttribute(
		"fPublished",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == attrValueOne
}

// SetFPublished sets whether the shape is published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *Shape) SetFPublished(value bool) {
	if value {
		s.SetAttribute(
			openxml.NewAttribute(
				"",
				"fPublished",
				"",
				attrValueTrue,
			),
		)
	} else {
		s.RemoveAttribute("fPublished", "")
	}
}

// NvSpPr returns the non-visual shape properties element, or nil if not
// present.
func (s *Shape) NvSpPr() *NvSpPr {
	elem := s.GetElement(
		"nvSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if nvSpPr, ok := elem.(*NvSpPr); ok {
		return nvSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NvSpPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateNvSpPr returns the non-visual shape properties, creating if
// needed.
func (s *Shape) GetOrCreateNvSpPr() *NvSpPr {
	nvSpPr := s.NvSpPr()
	if nvSpPr != nil {
		return nvSpPr
	}
	nvSpPr = NewNvSpPr()
	if first := s.FirstChild(); first != nil {
		s.InsertBefore(nvSpPr, first)
	} else {
		s.AppendChild(nvSpPr)
	}

	return nvSpPr
}

// SpPr returns the shape properties element, or nil if not present.
func (s *Shape) SpPr() *ShapeProperties {
	elem := s.GetElement(
		"spPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if spPr, ok := elem.(*ShapeProperties); ok {
		return spPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSpPr returns the shape properties, creating if needed.
func (s *Shape) GetOrCreateSpPr() *ShapeProperties {
	spPr := s.SpPr()
	if spPr != nil {
		return spPr
	}
	spPr = NewShapeProperties()
	nvSpPr := s.NvSpPr()
	if nvSpPr != nil {
		s.InsertAfter(spPr, nvSpPr)
	} else {
		s.AppendChild(spPr)
	}

	return spPr
}

// Style returns the style element, or nil if not present.
func (s *Shape) Style() *ShapeStyle {
	elem := s.GetElement(
		"style",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if style, ok := elem.(*ShapeStyle); ok {
		return style
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ShapeStyle{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateStyle returns the style element, creating if needed.
func (s *Shape) GetOrCreateStyle() *ShapeStyle {
	style := s.Style()
	if style != nil {
		return style
	}
	style = NewShapeStyle()
	spPr := s.SpPr()
	if spPr != nil {
		s.InsertAfter(style, spPr)
	} else {
		s.AppendChild(style)
	}

	return style
}

// TxBody returns the text body element, or nil if not present.
func (s *Shape) TxBody() *TextBody {
	elem := s.GetElement(
		"txBody",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if txBody, ok := elem.(*TextBody); ok {
		return txBody
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextBody{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTxBody returns the text body element, creating if needed.
func (s *Shape) GetOrCreateTxBody() *TextBody {
	txBody := s.TxBody()
	if txBody != nil {
		return txBody
	}
	txBody = NewTextBody()
	s.AppendChild(txBody)

	return txBody
}

// Clone creates a deep copy of this Shape element.
func (s *Shape) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &Shape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Shape element.
func (s *Shape) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.CompositeElementBase.CloneNode(
		deep,
	)

	return &Shape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NvSpPr represents non-visual shape properties (xdr:nvSpPr).
type NvSpPr struct {
	*openxml.CompositeElementBase
}

// NewNvSpPr creates a new NvSpPr element.
func NewNvSpPr() *NvSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"nvSpPr",
		PrefixXDR,
	)

	return &NvSpPr{CompositeElementBase: elem}
}

// CNvPr returns the common non-visual properties element,
// or nil if not present.
func (n *NvSpPr) CNvPr() *CNvPr {
	elem := n.GetElement(
		"cNvPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvPr, ok := elem.(*CNvPr); ok {
		return cNvPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCNvPr returns the common non-visual properties,
// creating if needed.
func (n *NvSpPr) GetOrCreateCNvPr() *CNvPr {
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		return cNvPr
	}
	cNvPr = NewCNvPr()
	if first := n.FirstChild(); first != nil {
		n.InsertBefore(cNvPr, first)
	} else {
		n.AppendChild(cNvPr)
	}

	return cNvPr
}

// CNvSpPr returns the non-visual shape drawing properties element.
func (n *NvSpPr) CNvSpPr() *CNvSpPr {
	elem := n.GetElement(
		"cNvSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvSpPr, ok := elem.(*CNvSpPr); ok {
		return cNvSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvSpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCNvSpPr returns the non-visual shape drawing properties, creating
// if needed.
func (n *NvSpPr) GetOrCreateCNvSpPr() *CNvSpPr {
	cNvSpPr := n.CNvSpPr()
	if cNvSpPr != nil {
		return cNvSpPr
	}
	cNvSpPr = NewCNvSpPr()
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		n.InsertAfter(cNvSpPr, cNvPr)
	} else {
		n.AppendChild(cNvSpPr)
	}

	return cNvSpPr
}

// Clone creates a deep copy of this NvSpPr element.
func (n *NvSpPr) Clone() openxml.Element {
	cloned := n.CompositeElementBase.Clone()

	return &NvSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this NvSpPr element.
func (n *NvSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.CompositeElementBase.CloneNode(
		deep,
	)

	return &NvSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CNvPr represents common non-visual properties (xdr:cNvPr).
type CNvPr struct {
	*openxml.CompositeElementBase
}

// NewCNvPr creates a new CNvPr element.
func NewCNvPr() *CNvPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cNvPr",
		PrefixXDR,
	)

	return &CNvPr{CompositeElementBase: elem}
}

// Id returns the shape ID.
func (c *CNvPr) Id() uint32 {
	attr, found := c.GetAttribute("id", "")
	if !found {
		return 0
	}
	var id uint32
	_, _ = fmt.Sscanf(attr.Value(), "%d", &id)

	return id
}

// SetId sets the shape ID.
func (c *CNvPr) SetId(id uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"id",
			"",
			fmt.Sprintf("%d", id),
		),
	)
}

// Name returns the shape name.
func (c *CNvPr) Name() string {
	attr, found := c.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the shape name.
func (c *CNvPr) SetName(name string) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Descr returns the description/alt text.
func (c *CNvPr) Descr() string {
	attr, found := c.GetAttribute("descr", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDescr sets the description/alt text.
func (c *CNvPr) SetDescr(descr string) {
	if descr == "" {
		c.RemoveAttribute("descr", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"descr",
			"",
			descr,
		),
	)
}

// Hidden returns whether the shape is hidden.
func (c *CNvPr) Hidden() bool {
	attr, found := c.GetAttribute("hidden", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == attrValueOne
}

// SetHidden sets whether the shape is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CNvPr) SetHidden(hidden bool) {
	if hidden {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"hidden",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("hidden", "")
	}
}

// Clone creates a deep copy of this CNvPr element.
func (c *CNvPr) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CNvPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CNvPr element.
func (c *CNvPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CNvPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CNvSpPr represents connection non-visual shape properties (xdr:cNvSpPr).
type CNvSpPr struct {
	*openxml.CompositeElementBase
}

// NewCNvSpPr creates a new CNvSpPr element.
func NewCNvSpPr() *CNvSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cNvSpPr",
		PrefixXDR,
	)

	return &CNvSpPr{CompositeElementBase: elem}
}

// TxBox returns whether this is a text box.
func (c *CNvSpPr) TxBox() bool {
	attr, found := c.GetAttribute("txBox", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == attrValueOne
}

// SetTxBox sets whether this is a text box.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CNvSpPr) SetTxBox(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"txBox",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("txBox", "")
	}
}

// Clone creates a deep copy of this CNvSpPr element.
func (c *CNvSpPr) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CNvSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CNvSpPr element.
func (c *CNvSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CNvSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ShapeProperties represents shape properties (xdr:spPr).
type ShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewShapeProperties creates a new ShapeProperties element.
func NewShapeProperties() *ShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"spPr",
		PrefixXDR,
	)

	return &ShapeProperties{
		CompositeElementBase: elem,
	}
}

// BWMode returns the black and white mode.
func (sp *ShapeProperties) BWMode() string {
	attr, found := sp.GetAttribute("bwMode", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetBWMode sets the black and white mode.
func (sp *ShapeProperties) SetBWMode(
	mode string,
) {
	if mode == "" {
		sp.RemoveAttribute("bwMode", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"bwMode",
			"",
			mode,
		),
	)
}

// Clone creates a deep copy of this ShapeProperties element.
func (sp *ShapeProperties) Clone() openxml.Element {
	cloned := sp.CompositeElementBase.Clone()

	return &ShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ShapeProperties element.
func (sp *ShapeProperties) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sp.CompositeElementBase.CloneNode(
		deep,
	)

	return &ShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ShapeStyle represents shape style (xdr:style).
type ShapeStyle struct {
	*openxml.CompositeElementBase
}

// NewShapeStyle creates a new ShapeStyle element.
func NewShapeStyle() *ShapeStyle {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"style",
		PrefixXDR,
	)

	return &ShapeStyle{CompositeElementBase: elem}
}

// Clone creates a deep copy of this ShapeStyle element.
func (ss *ShapeStyle) Clone() openxml.Element {
	cloned := ss.CompositeElementBase.Clone()

	return &ShapeStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ShapeStyle element.
func (ss *ShapeStyle) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ss.CompositeElementBase.CloneNode(
		deep,
	)

	return &ShapeStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TextBody represents the text body element (xdr:txBody).
type TextBody struct {
	*openxml.CompositeElementBase
}

// NewTextBody creates a new TextBody element.
func NewTextBody() *TextBody {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"txBody",
		PrefixXDR,
	)

	return &TextBody{CompositeElementBase: elem}
}

// Clone creates a deep copy of this TextBody element.
func (tb *TextBody) Clone() openxml.Element {
	cloned := tb.CompositeElementBase.Clone()

	return &TextBody{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TextBody element.
func (tb *TextBody) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tb.CompositeElementBase.CloneNode(
		deep,
	)

	return &TextBody{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
