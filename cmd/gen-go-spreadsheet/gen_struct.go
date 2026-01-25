package main

import (
	"os"
	"strings"
)

// generateStruct generates a Go struct for a schema element type.
// It creates the struct definition, constructor, and methods.
func generateStruct(f *os.File, t *SchemaType) {
	if existingTypes[t.ClassName] ||
		generatedTypes[t.ClassName] ||
		skipTypes[t.ClassName] {
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
// It checks against a known list of SpreadsheetML root elements.
func isPartRootType(name string) bool {
	roots := []string{
		// Main SpreadsheetML namespace (x:)
		"/x:workbook",             // WorkbookPart
		"/x:worksheet",            // WorksheetPart
		"/x:chartsheet",           // ChartsheetPart
		"/x:styleSheet",           // WorkbookStylesPart
		"/x:connections",          // ConnectionsPart
		"/x:queryTable",           // QueryTablePart
		"/x:externalLink",         // ExternalWorkbookPart
		"/x:calcChain",            // CalculationChainPart
		"/x:pivotTableDefinition", // PivotTablePart
		"/x:pivotCacheDefinition", // PivotTableCacheDefinitionPart
		"/x:pivotCacheRecords",    // PivotTableCacheRecordsPart
		"/x:comments",             // WorksheetCommentsPart
		"/x:table",                // TableDefinitionPart
		"/x:sst",                  // SharedStringTablePart

		// SpreadsheetML 2009/9 namespace (x14:)
		"/x14:slicers",               // SlicerPart
		"/x14:slicerCacheDefinition", // SlicerCachePart

		// SpreadsheetML 2010/11 namespace (x15:)
		"/x15:timelines",               // TimeLinePart
		"/x15:timelineCacheDefinition", // TimeLineCachePart

		// DrawingML SpreadsheetDrawing namespace (xdr:)
		"/xdr:wsDr", // DrawingsPart
	}
	for _, root := range roots {
		if strings.HasSuffix(name, root) {
			return true
		}
	}

	return false
}
