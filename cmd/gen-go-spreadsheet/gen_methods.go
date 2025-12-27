// Package main provides a code generator for SpreadsheetML elements.
package main

import (
	"os"
	"strings"
)

// generateConstructor writes the constructor function for the struct.
// It initializes the embedded base element with correct namespace.
func generateConstructor(
	f *os.File,
	t SchemaType,
	partRoot bool,
) {
	// Check if constructor name would conflict with a type name
	constructorName := "New" + t.ClassName
	if generatedTypes[constructorName] ||
		existingTypes[constructorName] {
		// Skip generating constructor to avoid name collision
		return
	}

	safeFprintf(
		f,
		"func New%s() *%s {\n\tret := &%s{}\n",
		t.ClassName,
		t.ClassName,
		t.ClassName,
	)

	uri := t.TargetNamespace
	prefix := ""
	for p, u := range namespaceMap {
		if u == uri {
			prefix = p

			break
		}
	}

	local := getLocalName(t.Name)
	safeFprintf(
		f,
		"\tns := %s\n",
		getGoNamespace(uri),
	)
	switch {
	case t.IsLeafElement:
		safeFprintf(f, "\tret.LeafElementBase = "+
			"openxml.NewLeafElement(ns, \"%s\", \"%s\")\n", local, prefix)
	case partRoot:
		safeFprintf(
			f,
			"\tret.PartRootElementBase = "+
				"openxml.NewPartRootElement(ns, \"%s\", \"%s\")\n",
			local,
			prefix,
		)
	default:
		safeFprintf(
			f,
			"\tret.CompositeElementBase = "+
				"openxml.NewCompositeElement(ns, \"%s\", \"%s\")\n",
			local,
			prefix,
		)
	}
	safeFprintf(f, "\treturn ret\n}\n\n")
}

// generateClone writes the Clone() method for the struct.
// It deep copies all attributes and children of the element.
func generateClone(f *os.File, t SchemaType) {
	safeFprintf(
		f,
		"func (m *%s) Clone() openxml.Element {\n\tret := New%s()\n",
		t.ClassName,
		t.ClassName,
	)
	generateCloneAttributes(f, t)
	generateCloneChildren(f, t)
	safeFprintf(f, "\treturn ret\n}\n\n")
}

// generateCloneAttributes writes the attribute cloning logic.
// It iterates through all schema attributes and copies their values.
func generateCloneAttributes(
	f *os.File,
	t SchemaType,
) {
	for _, attr := range t.Attributes {
		propName := attr.PropertyName
		if propName == "" {
			propName = toPascalCaseFromQName(
				attr.QName,
			)
		}
		safeFprintf(
			f,
			"\tif m.%s != nil {\n\t\tv := *m.%s\n\t\t"+
				"ret.%s = &v\n\t}\n",
			propName,
			propName,
			propName,
		)
	}
}

// toPascalCaseFromQName extracts local name from QName and converts.
func toPascalCaseFromQName(qname string) string {
	parts := strings.Split(qname, colon)
	if len(parts) == 2 {
		return toPascalCase(parts[1])
	}

	return toPascalCase(qname)
}

// generateCloneChildren writes the child element cloning logic.
// It handles both struct-value types and pointer-based elements.
func generateCloneChildren(
	f *os.File,
	t SchemaType,
) {
	seen := make(map[string]bool)
	for _, child := range t.Children {
		info, ok := typeMap[child.Name]
		if !ok {
			continue
		}
		propName := child.PropertyName
		if propName == "" {
			propName = info.ClassName
		}

		// Skip duplicate property names
		if seen[propName] {
			continue
		}
		seen[propName] = true

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

		safeFprintf(
			f,
			"\tif m.%s != nil {\n",
			propName,
		)
		if isStructValueType(info.ClassName) {
			safeFprintf(
				f,
				"\t\tv := m.%s.Clone()\n\t\tret.%s = &v\n",
				propName,
				propName,
			)
		} else {
			safeFprintf(f, "\t\tret.%s = m.%s.Clone().(*%s)\n",
				propName, propName, childType)
		}
		safeFprintf(f, "\t}\n")
	}
}

// generateValidate writes the Validate() method for the struct.
// It recursively validates all child elements.
func generateValidate(f *os.File, t SchemaType) {
	safeFprintf(
		f,
		"func (m *%s) Validate() error {\n",
		t.ClassName,
	)
	seen := make(map[string]bool)
	for _, child := range t.Children {
		info, ok := typeMap[child.Name]
		if !ok {
			continue
		}
		propName := child.PropertyName
		if propName == "" {
			propName = info.ClassName
		}

		// Skip duplicate property names
		if seen[propName] {
			continue
		}
		seen[propName] = true

		pkg := getGoPackage(info.Namespace)
		if shouldSkipValidation(
			pkg,
			info.ClassName,
		) {
			continue
		}

		safeFprintf(
			f,
			"\tif m.%s != nil {\n\t\tif err := m.%s.Validate(); "+
				"err != nil {\n\t\t	return err\n\t\t}\n\t}\n",
			propName,
			propName,
		)
	}
	safeFprintf(f, "\treturn nil\n}\n\n")
}

// shouldSkipValidation returns true if the element is not to be validated.
func shouldSkipValidation(
	pkg, className string,
) bool {
	if pkg == drawingMLPkg &&
		!existsInDrawingML(className) {
		return true
	}
	// Reordered operands for performance and revive compliance.
	return existingTypes[className] ||
		isStructValueType(className) ||
		isPointerValueType(className)
}
