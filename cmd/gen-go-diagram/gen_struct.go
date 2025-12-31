// Package main provides a code generator for Diagram elements.
package main

import (
	"os"
	"strings"
)

// generateStruct generates a Go struct for a schema element type.
// It creates the struct definition, constructor, and methods.
func generateStruct(f *os.File, t *SchemaType) {
	if existingTypes[t.ClassName] ||
		generatedTypes[t.ClassName] {
		return
	}
	generatedTypes[t.ClassName] = true

	// Documentation comment
	safeFprintf(
		f,
		"// %s\ntype %s struct {\n",
		t.Summary,
		t.ClassName,
	)

	isPartRoot := isPartRootType(t.Name)
	generateEmbeddedBase(f, t, isPartRoot)

	// XML Name field for encoding/xml
	safeFprintf(
		f,
		"\tXMLName xml.Name `xml:\"%s\"`\n",
		getXMLName(t.Name),
	)
	generateAttributes(f, t)
	generateChildren(f, t)
	safeFprintf(f, "}\n\n")

	generateConstructor(f, t, isPartRoot)
	generateClone(f, t)
	generateValidate(f, t)
}

// generateEmbeddedBase writes the embedded base type for the struct.
// It chooses between Leaf, PartRoot, or Composite base elements.
func generateEmbeddedBase(
	f *os.File,
	t *SchemaType,
	isPartRoot bool,
) {
	switch {
	case t.IsLeafElement:
		safeFprintf(
			f,
			"\t*openxml.LeafElementBase\n",
		)
	case isPartRoot:
		safeFprintf(
			f,
			"\t*openxml.PartRootElementBase\n",
		)
	default:
		safeFprintf(
			f,
			"\t*openxml.CompositeElementBase\n",
		)
	}
}

// generateAttributes writes the attribute fields for the struct.
// It maps schema types to Go types and generates appropriate XML tags.
func generateAttributes(
	f *os.File,
	t *SchemaType,
) {
	for _, attr := range t.Attributes {
		goType := mapType(attr.Type)
		propName := attr.PropertyName
		if propName == "" {
			propName = toPascalCaseFromQName(
				attr.QName,
			)
		}
		safeFprintf(
			f,
			"\t%s %s `xml:\"%s,attr,omitempty\"`\n",
			propName,
			goType,
			getAttrXMLName(attr.QName),
		)
	}
}

// generateChildren writes the child element fields for the struct.
// It handles optional children and maps them to appropriate Go types.
func generateChildren(f *os.File, t *SchemaType) {
	for _, child := range t.Children {
		info, ok := typeMap[child.Name]
		if !ok {
			continue
		}
		propName := child.PropertyName
		if propName == "" {
			propName = info.ClassName
		}

		childType := info.ClassName
		if pkg := getGoPackage(info.Namespace); pkg != "" {
			childType = pkg + "." + childType
		}

		if getGoPackage(
			info.Namespace,
		) == drawingMLPkg &&
			!existsInDrawingML(info.ClassName) {
			continue
		}

		typeName := childType
		if isStructValueType(info.ClassName) {
			typeName = strings.TrimPrefix(
				typeName,
				"*",
			)
		}
		safeFprintf(
			f,
			"\t%s *%s `xml:\"%s,omitempty\"`\n",
			propName,
			typeName,
			getChildXMLName(child.Name),
		)
	}
}

// isPartRootType returns true if the type is a root element of an OOXML part.
// It checks against a known list of Diagram root elements.
func isPartRootType(name string) bool {
	roots := []string{
		"/dgm:dataModel", "/dgm:layoutDef", "/dgm:colorsDef", "/dgm:styleDef",
	}
	for _, root := range roots {
		if strings.HasSuffix(name, root) {
			return true
		}
	}

	return false
}
