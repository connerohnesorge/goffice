// Package main provides a code generator for WordprocessingML elements.
package main

import (
	"os"
	"unicode"
)

const (
	valueSuffix = "Value"
)

// generateEnum generates a Go string-based enumeration for a schema type.
// It includes XML attribute marshaling and unmarshaling implementations.
func generateEnum(f *os.File, t *SchemaType) {
	if existingTypes[t.Name] ||
		generatedTypes[t.Name] {
		return
	}
	generatedTypes[t.Name] = true

	// Write type definition
	safeFprintf(
		f,
		"// %s\ntype %s string\nconst (\n",
		t.Name,
		t.Name,
	)

	// Write constant values
	for _, facet := range t.Facets {
		name := toPascalCase(facet.Value)
		if name == "" {
			name = "None"
		}
		// Prefix with type name if it starts with a digit
		// to ensure valid identifier
		if unicode.IsDigit(rune(name[0])) {
			name = t.Name + name
		}

		// Check if the constant name would conflict with a type name
		constName := t.Name + name
		_, exists := typeMap["/"+constName]
		if exists {
			// Add "Value" suffix to avoid conflict
			name += valueSuffix
		}
		_, exists = typeMap["/w:"+constName]
		if exists {
			name += valueSuffix
		}
		if generatedTypes[constName] {
			name += valueSuffix
		}

		safeFprintf(f, "\t%s%s %s = \"%s\"\n",
			t.Name, name, t.Name, facet.Value)
	}
	safeFprintf(f, ")\n\n")

	// Write marshaling methods
	generateEnumMarshaling(f, t.Name)
}

// generateEnumMarshaling writes XML marshaling methods.
// It includes MarshalXMLAttr and UnmarshalXMLAttr.
func generateEnumMarshaling(
	f *os.File,
	typeName string,
) {
	safeFprintf(
		f,
		"func (e %s) MarshalXMLAttr(name xml.Name) "+
			"(xml.Attr, error) {\n",
		typeName,
	)
	safeFprintf(
		f,
		"\treturn xml.Attr{Name: name, Value: string(e)}, nil\n}\n\n",
	)

	safeFprintf(
		f,
		"func (e *%s) UnmarshalXMLAttr(attr xml.Attr) error {\n",
		typeName,
	)
	safeFprintf(
		f,
		"\t*e = %s(attr.Value)\n\treturn nil\n}\n\n",
		typeName,
	)
}
