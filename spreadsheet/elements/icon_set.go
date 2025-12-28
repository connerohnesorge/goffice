package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// IconSetType represents the type of icon set for conditional formatting.
type IconSetType string

const (
	// IconSet3Arrows is 3 arrows (colored).
	IconSet3Arrows IconSetType = "3Arrows"
	// IconSet3ArrowsGray is 3 arrows (gray).
	IconSet3ArrowsGray IconSetType = "3ArrowsGray"
	// IconSet3Flags is 3 flags.
	IconSet3Flags IconSetType = "3Flags"
	// IconSet3TrafficLights1 is 3 traffic lights (unrimmed).
	IconSet3TrafficLights1 IconSetType = "3TrafficLights1"
	// IconSet3TrafficLights2 is 3 traffic lights (rimmed).
	IconSet3TrafficLights2 IconSetType = "3TrafficLights2"
	// IconSet3Signs is 3 signs.
	IconSet3Signs IconSetType = "3Signs"
	// IconSet3Symbols is 3 symbols (circled).
	IconSet3Symbols IconSetType = "3Symbols"
	// IconSet3Symbols2 is 3 symbols (uncircled).
	IconSet3Symbols2 IconSetType = "3Symbols2"
	// IconSet4Arrows is 4 arrows (colored).
	IconSet4Arrows IconSetType = "4Arrows"
	// IconSet4ArrowsGray is 4 arrows (gray).
	IconSet4ArrowsGray IconSetType = "4ArrowsGray"
	// IconSet4RedToBlack is 4 circles red to black.
	IconSet4RedToBlack IconSetType = "4RedToBlack"
	// IconSet4Rating is 4 ratings.
	IconSet4Rating IconSetType = "4Rating"
	// IconSet4TrafficLights is 4 traffic lights.
	IconSet4TrafficLights IconSetType = "4TrafficLights"
	// IconSet5Arrows is 5 arrows (colored).
	IconSet5Arrows IconSetType = "5Arrows"
	// IconSet5ArrowsGray is 5 arrows (gray).
	IconSet5ArrowsGray IconSetType = "5ArrowsGray"
	// IconSet5Rating is 5 ratings.
	IconSet5Rating IconSetType = "5Rating"
	// IconSet5Quarters is 5 quarters.
	IconSet5Quarters IconSetType = "5Quarters"
)

// IconSet represents an icon set element (x:iconSet) for conditional
// formatting. An icon set displays icons based on cell values.
type IconSet struct {
	*openxml.CompositeElementBase
}

// NewIconSet creates a new IconSet element.
func NewIconSet() *IconSet {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"iconSet",
		PrefixDefault,
	)

	return &IconSet{CompositeElementBase: elem}
}

// IconSetType returns the type of icon set. Attribute: iconSet.
// Default is IconSet3TrafficLights1.
func (is *IconSet) IconSetType() IconSetType {
	attr, found := is.GetAttribute(
		elemNameIconSet,
		"",
	)
	if !found {
		return IconSet3TrafficLights1
	}

	return IconSetType(attr.Value())
}

// SetIconSetType sets the type of icon set. Attribute: iconSet.
func (is *IconSet) SetIconSetType(
	iconSetType IconSetType,
) {
	if iconSetType == "" ||
		iconSetType == IconSet3TrafficLights1 {
		is.RemoveAttribute(elemNameIconSet, "")

		return
	}
	is.SetAttribute(
		openxml.NewAttribute(
			"",
			elemNameIconSet,
			"",
			string(iconSetType),
		),
	)
}

// ShowValue returns whether the cell value is shown. Attribute: showValue.
// Default is true.
func (is *IconSet) ShowValue() bool {
	attr, found := is.GetAttribute(
		"showValue",
		"",
	)
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
func (is *IconSet) SetShowValue(value bool) {
	if value {
		is.RemoveAttribute("showValue", "")
	} else {
		is.SetAttribute(
			openxml.NewAttribute(
				"",
				"showValue",
				"",
				attrValueFalse,
			),
		)
	}
}

// Percent returns whether the thresholds are percentages. Attribute: percent.
// Default is true.
func (is *IconSet) Percent() bool {
	attr, found := is.GetAttribute("percent", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPercent sets whether the thresholds are percentages. Attribute: percent.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (is *IconSet) SetPercent(value bool) {
	if value {
		is.RemoveAttribute("percent", "")
	} else {
		is.SetAttribute(
			openxml.NewAttribute(
				"",
				"percent",
				"",
				attrValueFalse,
			),
		)
	}
}

// Reverse returns whether the icon order is reversed. Attribute: reverse.
// Default is false.
func (is *IconSet) Reverse() bool {
	attr, found := is.GetAttribute("reverse", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetReverse sets whether the icon order is reversed. Attribute: reverse.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (is *IconSet) SetReverse(value bool) {
	if value {
		is.SetAttribute(
			openxml.NewAttribute(
				"",
				"reverse",
				"",
				attrValueTrue,
			),
		)
	} else {
		is.RemoveAttribute("reverse", "")
	}
}

// Cfvos returns an iterator over all Cfvo elements in this icon set.
func (is *IconSet) Cfvos() iter.Seq[*Cfvo] {
	return func(yield func(*Cfvo) bool) {
		for child := range is.Children() {
			if child.LocalName() != elemNameCfvo ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var cfvo *Cfvo
			switch v := child.(type) {
			case *Cfvo:
				cfvo = v
			case *openxml.LeafElementBase:
				cfvo = &Cfvo{LeafElementBase: v}
			}
			if cfvo != nil && !yield(cfvo) {
				return
			}
		}
	}
}

// CfvoCount returns the number of Cfvo elements.
func (is *IconSet) CfvoCount() int {
	count := 0
	for range is.Cfvos() {
		count++
	}

	return count
}

// AddCfvo adds a new Cfvo element to the icon set.
func (is *IconSet) AddCfvo() *Cfvo {
	cfvo := NewCfvo()
	is.AppendChild(cfvo)

	return cfvo
}

// Set3IconSet sets up a 3-icon icon set with the given type and thresholds.
//
//nolint:revive // argument-limit
func (is *IconSet) Set3IconSet(
	iconSetType IconSetType,
	threshold1Type CfvoType, threshold1Val string,
	threshold2Type CfvoType, threshold2Val string,
) {
	// Clear existing
	is.RemoveAllChildren()
	is.SetIconSetType(iconSetType)

	// Add minimum cfvo (usually min or percent 0)
	minCfvo := is.AddCfvo()
	minCfvo.SetType(CfvoTypePercent)
	minCfvo.SetVal("0")

	// Add first threshold
	t1Cfvo := is.AddCfvo()
	t1Cfvo.SetType(threshold1Type)
	if threshold1Val != "" {
		t1Cfvo.SetVal(threshold1Val)
	}

	// Add second threshold
	t2Cfvo := is.AddCfvo()
	t2Cfvo.SetType(threshold2Type)
	if threshold2Val != "" {
		t2Cfvo.SetVal(threshold2Val)
	}
}

// Set4IconSet sets up a 4-icon icon set with the given type and thresholds.
//
//nolint:revive // argument-limit
func (is *IconSet) Set4IconSet(
	iconSetType IconSetType,
	threshold1Type CfvoType, threshold1Val string,
	threshold2Type CfvoType, threshold2Val string,
	threshold3Type CfvoType, threshold3Val string,
) {
	// Clear existing
	is.RemoveAllChildren()
	is.SetIconSetType(iconSetType)

	// Add minimum cfvo
	minCfvo := is.AddCfvo()
	minCfvo.SetType(CfvoTypePercent)
	minCfvo.SetVal("0")

	// Add thresholds
	t1Cfvo := is.AddCfvo()
	t1Cfvo.SetType(threshold1Type)
	if threshold1Val != "" {
		t1Cfvo.SetVal(threshold1Val)
	}

	t2Cfvo := is.AddCfvo()
	t2Cfvo.SetType(threshold2Type)
	if threshold2Val != "" {
		t2Cfvo.SetVal(threshold2Val)
	}

	t3Cfvo := is.AddCfvo()
	t3Cfvo.SetType(threshold3Type)
	if threshold3Val != "" {
		t3Cfvo.SetVal(threshold3Val)
	}
}

// Set5IconSet sets up a 5-icon icon set with the given type and thresholds.
//
//nolint:revive // argument-limit
func (is *IconSet) Set5IconSet(
	iconSetType IconSetType,
	threshold1Type CfvoType, threshold1Val string,
	threshold2Type CfvoType, threshold2Val string,
	threshold3Type CfvoType, threshold3Val string,
	threshold4Type CfvoType, threshold4Val string,
) {
	// Clear existing
	is.RemoveAllChildren()
	is.SetIconSetType(iconSetType)

	// Add minimum cfvo
	minCfvo := is.AddCfvo()
	minCfvo.SetType(CfvoTypePercent)
	minCfvo.SetVal("0")

	// Add thresholds
	t1Cfvo := is.AddCfvo()
	t1Cfvo.SetType(threshold1Type)
	if threshold1Val != "" {
		t1Cfvo.SetVal(threshold1Val)
	}

	t2Cfvo := is.AddCfvo()
	t2Cfvo.SetType(threshold2Type)
	if threshold2Val != "" {
		t2Cfvo.SetVal(threshold2Val)
	}

	t3Cfvo := is.AddCfvo()
	t3Cfvo.SetType(threshold3Type)
	if threshold3Val != "" {
		t3Cfvo.SetVal(threshold3Val)
	}

	t4Cfvo := is.AddCfvo()
	t4Cfvo.SetType(threshold4Type)
	if threshold4Val != "" {
		t4Cfvo.SetVal(threshold4Val)
	}
}

// Clone creates a deep copy of this IconSet element.
func (is *IconSet) Clone() openxml.Element {
	cloned := is.CompositeElementBase.Clone()

	return &IconSet{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this IconSet element.
func (is *IconSet) CloneNode(
	deep bool,
) openxml.Element {
	cloned := is.CompositeElementBase.CloneNode(
		deep,
	)

	return &IconSet{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
