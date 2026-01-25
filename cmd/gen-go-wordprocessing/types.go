package main

// SchemaType represents a type definition in the schema JSON.
// It contains metadata about the XML element or attribute.
type SchemaType struct {
	// Name is the schema name of the type.
	Name string `json:"Name"`
	// ClassName is the Go class name to be generated.
	ClassName string `json:"ClassName"`
	// Summary is a brief description of the type.
	Summary string `json:"Summary"`
	// BaseClass is the parent class in the schema hierarchy.
	BaseClass string `json:"BaseClass"`
	// IsDerived indicates if the type is derived from another.
	IsDerived bool `json:"IsDerived"`
	// IsLeafElement indicates if the element has no child elements.
	IsLeafElement bool `json:"IsLeafElement"`
	// CompositeType defines the kind of composite element.
	CompositeType string `json:"CompositeType"`
	// Attributes is the list of XML attributes for this type.
	Attributes []SchemaAttribute `json:"Attributes"`
	// Children is the list of child XML elements.
	Children []SchemaChild `json:"Children"`
	// Particle defines the child element ordering and cardinality.
	Particle *SchemaParticle `json:"Particle"`
	// Facets contains constraints for simple types (enums).
	Facets []SchemaFacet `json:"Facets"`
	// TargetNamespace is the XML namespace URI for this type.
	TargetNamespace string
}

// SchemaFacet represents a constraint facet in the schema JSON,
// such as an enumeration value.
type SchemaFacet struct {
	// Value is the literal value of the facet.
	Value string `json:"Value"`
}

// SchemaAttribute represents an attribute definition in the schema JSON.
type SchemaAttribute struct {
	// QName is the qualified name of the attribute.
	QName string `json:"QName"`
	// PropertyName is the Go property name for the attribute.
	PropertyName string `json:"PropertyName"`
	// Type is the schema type of the attribute value.
	Type string `json:"Type"`
	// PropertyComments is the documentation for the property.
	PropertyComments string `json:"PropertyComments"`
}

// SchemaChild represents a child element definition in the schema JSON.
type SchemaChild struct {
	// Name is the schema name of the child element.
	Name string `json:"Name"`
	// PropertyName is the Go property name for the child element.
	PropertyName string `json:"PropertyName"`
}

// SchemaParticle represents a particle definition in the schema JSON,
// defining how child elements are structured (sequence, choice, etc.).
type SchemaParticle struct {
	// Kind is the type of particle (e.g., Sequence, Choice).
	Kind string `json:"Kind"`
	// Items is the list of child particles or elements.
	Items []any `json:"Items"`
}

// SchemaFile represents the structure of a schema JSON file.
type SchemaFile struct {
	// TargetNamespace is the default namespace for types in this file.
	TargetNamespace string `json:"TargetNamespace"`
	// Types is the list of complex type definitions.
	Types []SchemaType `json:"Types"`
	// Enums is the list of simple type enumeration definitions.
	Enums []SchemaType `json:"Enums"`
}

// NamespaceEntry represents a namespace prefix-to-URI mapping.
type NamespaceEntry struct {
	// Prefix is the canonical XML prefix (e.g., "w", "a").
	Prefix string `json:"Prefix"`
	// Uri is the full XML namespace URI.
	Uri string `json:"Uri"`
}

// TypeInfo holds basic metadata about a schema type used for cross-referencing.
type TypeInfo struct {
	// ClassName is the generated Go struct name.
	ClassName string
	// Namespace is the XML namespace URI.
	Namespace string
	// Version is the Office version this type was introduced in.
	Version FileFormatVersion
}

var (
	// namespaceMap maps XML prefixes to their namespace URIs.
	namespaceMap = make(map[string]string)
	// typeMap maps schema type names to their TypeInfo metadata.
	typeMap = make(map[string]TypeInfo)
	// enumMap maps enum names to their SchemaType definitions.
	enumMap = make(map[string]SchemaType)
	// enumNameMap tracks enum names to detect struct naming conflicts.
	enumNameMap = make(map[string]bool)
	// existingTypes tracks types already defined in manually written files.
	existingTypes = make(map[string]bool)
	// generatedTypes prevents duplicate generation of the same type.
	generatedTypes = make(map[string]bool)
	// drawingMLTypes tracks types available in the drawingml package.
	drawingMLTypes = make(map[string]bool)
)

const (
	// drawingMLPkg is the name of the shared DrawingML package.
	drawingMLPkg = "drawingml"
	// stringValueType is the Go type for OpenXML string values.
	stringValueType = "*types.StringValue"
	// nsWordprocessingML is the WordprocessingML main namespace URI.
	nsWordprocessingML = "http://schemas.openxmlformats.org/" +
		"wordprocessingml/2006/main"
	// nsRelationships is the Office document relationships namespace URI.
	nsRelationships = "http://schemas.openxmlformats.org/" +
		"officeDocument/2006/relationships"
	// nsDrawingML is the DrawingML main namespace URI.
	nsDrawingML = "http://schemas.openxmlformats.org/" +
		"drawingml/2006/main"
	// nsDrawingMLPrefix is the common prefix for Microsoft Drawing namespaces.
	nsDrawingMLPrefix = "http://schemas.microsoft.com/office/drawing/"
	// colon is the XML namespace separator.
	colon = ":"
)
