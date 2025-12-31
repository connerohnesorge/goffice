package openxml

import (
	"bytes"
	"io"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// AlternateContent represents the mc:AlternateContent element for version compatibility.
// It contains Choice elements for newer Office versions and a Fallback for older versions.
// This element is part of the Markup Compatibility namespace and allows documents to
// provide both modern and legacy content for different Office versions.
type AlternateContent struct {
	CompositeElementBase
}

// NewAlternateContent creates a new AlternateContent element.
func NewAlternateContent() *AlternateContent {
	ac := &AlternateContent{}
	InitBaseElement(
		&ac.BaseElement,
		NamespaceMarkupCompatibility,
		"AlternateContent",
		GetPrefixForNamespace(
			NamespaceMarkupCompatibility,
		),
		nil,
	)

	return ac
}

// NewAlternateContentWithFeatures creates a new AlternateContent element with parent features.
func NewAlternateContentWithFeatures(
	parentFeatures *features.FeatureCollection,
) *AlternateContent {
	ac := &AlternateContent{}
	InitBaseElement(
		&ac.BaseElement,
		NamespaceMarkupCompatibility,
		"AlternateContent",
		GetPrefixForNamespace(
			NamespaceMarkupCompatibility,
		),
		parentFeatures,
	)

	return ac
}

// Choices returns all Choice child elements.
func (ac *AlternateContent) Choices() []*Choice {
	var choices []*Choice
	for child := range ac.Children() {
		if c, ok := child.(*Choice); ok {
			choices = append(choices, c)
		}
	}

	return choices
}

// Fallback returns the Fallback child element, or nil if not present.
func (ac *AlternateContent) Fallback() *Fallback {
	for child := range ac.Children() {
		if fb, ok := child.(*Fallback); ok {
			return fb
		}
	}

	return nil
}

// SelectContent returns the appropriate content based on target version.
// It evaluates the Choice elements in order and returns the first one whose
// Requires attribute is satisfied by the target version. If no Choice is
// satisfied, it returns the Fallback content.
func (ac *AlternateContent) SelectContent(
	targetVersion FileFormatVersion,
) Element {
	// Check each Choice element
	for _, choice := range ac.Choices() {
		if choice.IsSatisfiedBy(targetVersion) {
			return choice
		}
	}

	// If no Choice is satisfied, return the Fallback
	return ac.Fallback()
}

// Choice represents mc:Choice - content for specific Office version.
// The Requires attribute specifies which namespace prefix is required,
// which corresponds to a specific Office version.
type Choice struct {
	CompositeElementBase
}

// NewChoice creates a new Choice element.
func NewChoice() *Choice {
	c := &Choice{}
	InitBaseElement(
		&c.BaseElement,
		NamespaceMarkupCompatibility,
		"Choice",
		GetPrefixForNamespace(
			NamespaceMarkupCompatibility,
		),
		nil,
	)

	return c
}

// NewChoiceWithFeatures creates a new Choice element with parent features.
func NewChoiceWithFeatures(
	parentFeatures *features.FeatureCollection,
) *Choice {
	c := &Choice{}
	InitBaseElement(
		&c.BaseElement,
		NamespaceMarkupCompatibility,
		"Choice",
		GetPrefixForNamespace(
			NamespaceMarkupCompatibility,
		),
		parentFeatures,
	)

	return c
}

// Requires returns the namespace prefix required for this choice.
// This is stored in the "Requires" attribute.
func (c *Choice) Requires() string {
	if attr, ok := c.GetAttribute("Requires", ""); ok {
		return attr.Value()
	}

	return ""
}

// SetRequires sets the namespace prefix required for this choice.
func (c *Choice) SetRequires(prefix string) {
	c.SetAttribute(
		NewSimpleAttribute("Requires", prefix),
	)
}

// IsSatisfiedBy checks if the given Office version satisfies this Choice's requirements.
// It checks if the Requires prefix corresponds to a version that is supported by
// the target version.
func (c *Choice) IsSatisfiedBy(
	targetVersion FileFormatVersion,
) bool {
	requires := c.Requires()
	if requires == "" {
		return false
	}

	// Get the version required by this Choice
	requiredVersion := prefixToVersion(requires)
	if requiredVersion == FileFormatVersionOffice2007 {
		// Unknown prefix, cannot satisfy
		return false
	}

	// Check if target version is >= required version
	return targetVersion >= requiredVersion
}

// Fallback represents mc:Fallback - content for older Office versions.
// This element contains content that will be used when the Choice requirements
// are not satisfied.
type Fallback struct {
	CompositeElementBase
}

// NewFallback creates a new Fallback element.
func NewFallback() *Fallback {
	fb := &Fallback{}
	InitBaseElement(
		&fb.BaseElement,
		NamespaceMarkupCompatibility,
		"Fallback",
		GetPrefixForNamespace(
			NamespaceMarkupCompatibility,
		),
		nil,
	)

	return fb
}

// NewFallbackWithFeatures creates a new Fallback element with parent features.
func NewFallbackWithFeatures(
	parentFeatures *features.FeatureCollection,
) *Fallback {
	fb := &Fallback{}
	InitBaseElement(
		&fb.BaseElement,
		NamespaceMarkupCompatibility,
		"Fallback",
		GetPrefixForNamespace(
			NamespaceMarkupCompatibility,
		),
		parentFeatures,
	)

	return fb
}

// prefixToVersion converts a namespace prefix to its corresponding FileFormatVersion.
// Common prefixes:
//   - w14 -> Office2010 (Word 2010)
//   - w15 -> Office2013 (Word 2013)
//   - w16 -> Office2016 (Word 2016)
//   - x14 -> Office2010 (Excel 2010)
//   - x15 -> Office2013 (Excel 2013)
//   - a14 -> Office2010 (Drawing 2010)
//
// Returns Office2007 if the prefix is unknown.
func prefixToVersion(
	prefix string,
) FileFormatVersion {
	// Map common version suffixes to Office versions
	// The number after the letter(s) corresponds to the Office version
	if len(prefix) < 2 {
		return FileFormatVersionOffice2007
	}

	// Extract the version suffix (e.g., "14" from "w14", "15" from "x15")
	var suffix string
	for i := len(prefix) - 1; i >= 0; i-- {
		if prefix[i] < '0' || prefix[i] > '9' {
			break
		}
		suffix = prefix[i:len(prefix)-len(suffix)] + suffix
	}

	switch suffix {
	case "14":
		return FileFormatVersionOffice2010
	case "15":
		return FileFormatVersionOffice2013
	case "16":
		return FileFormatVersionOffice2016
	case "19":
		return FileFormatVersionOffice2019
	case "21":
		return FileFormatVersionOffice2021
	case "22":
		return FileFormatVersionOffice2022
	case "23":
		return FileFormatVersionOffice2023
	case "24":
		return FileFormatVersionOffice2024
	case "25":
		return FileFormatVersionOffice2025
	default:
		return FileFormatVersionOffice2007
	}
}

// WriteXML writes the XML representation to the given writer.
func (ac *AlternateContent) WriteXML(
	w io.Writer,
) error {
	return ac.CompositeElementBase.WriteXML(w)
}

// WriteXML writes the XML representation to the given writer.
func (c *Choice) WriteXML(w io.Writer) error {
	return c.CompositeElementBase.WriteXML(w)
}

// WriteXML writes the XML representation to the given writer.
func (fb *Fallback) WriteXML(w io.Writer) error {
	return fb.CompositeElementBase.WriteXML(w)
}

// OuterXml returns the complete XML representation of this element.
func (ac *AlternateContent) OuterXml() string {
	var buf bytes.Buffer
	_ = ac.WriteXML(&buf)

	return buf.String()
}

// OuterXml returns the complete XML representation of this element.
func (c *Choice) OuterXml() string {
	var buf bytes.Buffer
	_ = c.WriteXML(&buf)

	return buf.String()
}

// OuterXml returns the complete XML representation of this element.
func (fb *Fallback) OuterXml() string {
	var buf bytes.Buffer
	_ = fb.WriteXML(&buf)

	return buf.String()
}
