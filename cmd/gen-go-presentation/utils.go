// Package main provides a code generator for PresentationML elements.
package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// isStructValueType returns true if the type should be handled as a struct.
// It is used for types like Offset or Point2D that are typically values.
func isStructValueType(name string) bool {
	return name == "Offset" ||
		name == "Point2D" ||
		name == "Extent" ||
		name == "EffectExtent"
}

// isPointerValueType returns true if the type should be a pointer.
// It includes specialized DrawingML and PresentationML types.
func isPointerValueType(name string) bool {
	pointerTypes := map[string]bool{
		"RgbColor": true, "HslColor": true, "SystemColor": true,
		"SchemeColor": true, "PresetColor": true, "NoFill": true,
		"SolidFill": true, "GradientFill": true, "BlipFill": true,
		"PatternFill": true, "EffectList": true, "EffectDag": true,
		"ShapeProperties": true, "TextBody": true, "LineProperties": true,
		"Outline": true, "Chart": true,
	}

	return pointerTypes[name]
}

// getXMLName converts a schema type name into an XML name string.
// It includes the full namespace URI for encoding/xml tags.
func getXMLName(name string) string {
	parts := strings.Split(name, "/")
	last := parts[len(parts)-1]

	tagParts := strings.Split(last, colon)
	if len(tagParts) == 2 {
		prefix := tagParts[0]
		local := tagParts[1]
		uri, ok := namespaceMap[prefix]
		if ok {
			return fmt.Sprintf(
				"%s %s",
				uri,
				local,
			)
		}

		return local
	}

	return last
}

// getLocalName extracts the local XML name from a schema type name string.
// For example, "p:sld" becomes "sld".
func getLocalName(name string) string {
	parts := strings.Split(name, "/")
	last := parts[len(parts)-1]
	tagParts := strings.Split(last, colon)
	if len(tagParts) == 2 {
		return tagParts[1]
	}

	return last
}

// getChildXMLName returns the local name for a child element.
func getChildXMLName(name string) string {
	return getLocalName(name)
}

// getAttrXMLName strips the leading colon from qualified attribute names.
// Some schema definitions use ":attrName" which needs to be cleaned up.
func getAttrXMLName(qname string) string {
	if strings.HasPrefix(qname, colon) {
		return qname[1:]
	}

	return qname
}

// getGoPackage returns the Go package name associated with an XML namespace.
// It determines if a type belongs to drawingml or is local to elements.
func getGoPackage(ns string) string {
	switch {
	case ns == nsDrawingML:
		return drawingMLPkg
	case ns == nsPresentationML:
		return ""
	case strings.Contains(ns, drawingMLPkg):
		return drawingMLPkg
	case strings.HasPrefix(ns, nsDrawingMLPrefix):
		return drawingMLPkg
	case ns == "http://schemas.microsoft.com/office/tasks/2019/documenttasks":
		return ""
	case ns == "http://schemas.microsoft.com/office/comments/2020/reactions":
		return ""
	default:
		return ""
	}
}

// getGoNamespace returns the Go variable representing an XML namespace URI.
func getGoNamespace(uri string) string {
	switch uri {
	case nsPresentationML:
		return "openxml.NamespacePresentationML"
	case nsRelationships:
		return "openxml.NamespaceRelationships"
	case nsDrawingML:
		return "openxml.NamespaceDrawingML"
	default:
		return fmt.Sprintf("%q", uri)
	}
}

// mapType maps a schema type to its Go type.
// It handles numeric, boolean, string, and complex types.
func mapType(schemaType string) string {
	if t := mapNumericType(schemaType); t != "" {
		return t
	}
	if t := mapBooleanType(schemaType); t != "" {
		return t
	}
	if t := mapStringLikeType(schemaType); t != "" {
		return t
	}

	return mapComplexType(schemaType)
}

// mapNumericType maps schema numeric types to Go types.
func mapNumericType(schemaType string) string {
	switch schemaType {
	case "Int32Value":
		return "*types.Int32Value"
	case "UInt32Value":
		return "*types.UInt32Value"
	case "DoubleValue":
		return "*types.DoubleValue"
	case "Int64Value":
		return "*types.Int64Value"
	case "UInt64Value":
		return "*types.UInt64Value"
	case "DecimalValue":
		return "*types.DecimalValue"
	case "ByteValue":
		return "*types.ByteValue"
	default:
		return ""
	}
}

// mapBooleanType maps schema boolean-like types to Go types.
func mapBooleanType(schemaType string) string {
	switch schemaType {
	case "BooleanValue":
		return "*types.BooleanValue"
	case "OnOffValue":
		return "*types.OnOffValue"
	case "TrueFalseValue":
		return "*types.TrueFalseValue"
	case "TrueFalseBlankValue":
		return "*types.TrueFalseBlankValue"
	default:
		return ""
	}
}

// mapStringLikeType maps schema string-like types to Go types.
func mapStringLikeType(schemaType string) string {
	switch schemaType {
	case "StringValue",
		"DateTimeValue",
		"SByteValue",
		"SingleValue",
		"HexBinaryValue",
		"Base64BinaryValue":
		return stringValueType
	default:
		return ""
	}
}

// mapComplexType handles EnumValue and ListValue schema types.
func mapComplexType(schemaType string) string {
	if strings.HasPrefix(
		schemaType,
		"EnumValue",
	) {
		start := strings.LastIndex(
			schemaType,
			".",
		)
		if start != -1 &&
			start < len(schemaType)-1 {
			enumName := schemaType[start+1 : len(schemaType)-1]

			return fmt.Sprintf(
				"*types.EnumValue[%s]",
				enumName,
			)
		}

		return stringValueType + " // " + schemaType
	}
	if strings.HasPrefix(
		schemaType,
		"ListValue",
	) {
		return stringValueType
	}

	return stringValueType + " // " + schemaType
}

// toPascalCase converts a string to PascalCase using English rules.
func toPascalCase(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.FieldsFunc(
		s,
		func(r rune) bool {
			return r == '-' || r == '_' ||
				r == ' '
		},
	)

	var result strings.Builder
	caser := cases.Title(language.English)
	for _, p := range parts {
		if len(p) > 0 {
			result.WriteString(caser.String(p))
		}
	}
	if result.Len() == 0 {
		return caser.String(s)
	}

	return result.String()
}

// existsInDrawingML returns true if the type is in DrawingML pkg.
func existsInDrawingML(name string) bool {
	return drawingMLTypes[name]
}
