package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// CfvoType represents the type of conditional format value object.
type CfvoType string

const (
	// CfvoTypeNum indicates a numeric value.
	CfvoTypeNum CfvoType = "num"
	// CfvoTypePercent indicates a percentage value.
	CfvoTypePercent CfvoType = "percent"
	// CfvoTypeMax indicates the maximum value in the range.
	CfvoTypeMax CfvoType = "max"
	// CfvoTypeMin indicates the minimum value in the range.
	CfvoTypeMin CfvoType = "min"
	// CfvoTypeFormula indicates a formula value.
	CfvoTypeFormula CfvoType = "formula"
	// CfvoTypePercentile indicates a percentile value.
	CfvoTypePercentile CfvoType = "percentile"
	// CfvoTypeAutoMin indicates automatic minimum value (Excel 2010+).
	CfvoTypeAutoMin CfvoType = "autoMin"
	// CfvoTypeAutoMax indicates automatic maximum value (Excel 2010+).
	CfvoTypeAutoMax CfvoType = "autoMax"
)

// Cfvo represents a conditional format value object element (x:cfvo).
// This element specifies the value thresholds for color scales,
// data bars, and icon sets.
type Cfvo struct {
	*openxml.LeafElementBase
}

// NewCfvo creates a new Cfvo element.
func NewCfvo() *Cfvo {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"cfvo",
		PrefixDefault,
	)

	return &Cfvo{LeafElementBase: elem}
}

// Type returns the type of the value object. Attribute: type.
func (c *Cfvo) Type() CfvoType {
	attr, found := c.GetAttribute("type", "")
	if !found {
		return ""
	}

	return CfvoType(attr.Value())
}

// SetType sets the type of the value object. Attribute: type.
func (c *Cfvo) SetType(cfvoType CfvoType) {
	if cfvoType == "" {
		c.RemoveAttribute("type", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(cfvoType),
		),
	)
}

// Val returns the value for this threshold. Attribute: val.
func (c *Cfvo) Val() string {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the value for this threshold. Attribute: val.
func (c *Cfvo) SetVal(val string) {
	if val == "" {
		c.RemoveAttribute("val", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			val,
		),
	)
}

// Gte returns whether values >= this threshold are included.
// Attribute: gte. Default is true.
func (c *Cfvo) Gte() bool {
	attr, found := c.GetAttribute("gte", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetGte sets whether values greater than or equal to this threshold are included.
// Attribute: gte.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Cfvo) SetGte(value bool) {
	if value {
		// Default is true, so remove attribute
		c.RemoveAttribute("gte", "")
	} else {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"gte",
				"",
				attrValueFalse,
			),
		)
	}
}

// Clone creates a deep copy of this Cfvo element.
func (c *Cfvo) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &Cfvo{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Cfvo element.
func (c *Cfvo) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &Cfvo{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
