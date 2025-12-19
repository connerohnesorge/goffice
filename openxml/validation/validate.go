package validation

import (
	"reflect"
	"sync"
)

// ValidatableElement is an interface for elements that can be validated.
type ValidatableElement interface {
	// LocalName returns the element's local name.
	LocalName() string
	// NamespaceURI returns the element's namespace URI.
	NamespaceURI() string
	// Parent returns the parent element.
	Parent() interface{}
	// Attributes returns the element's attributes as a map.
	Attributes() []interface{}
}

// ValidatableCompositeElement is an interface for composite elements.
type ValidatableCompositeElement interface {
	ValidatableElement
	// Children returns the child elements.
	Children() []interface{}
}

// ValidatablePackage is an interface for packages that can be validated.
type ValidatablePackage interface {
	// Parts returns an iterator over all parts.
	Parts() []interface{}
	// MainPart returns the main document part.
	MainPart() interface{}
}

// ValidatablePart is an interface for parts that can be validated.
type ValidatablePart interface {
	// URI returns the part URI.
	URI() string
	// ContentType returns the part content type.
	ContentType() string
	// RootElement returns the root element of the part.
	RootElement() interface{}
}

// Validate validates an OpenXML package.
func Validate(
	pkg interface{},
	version FileFormatVersions,
	settings *ValidationSettings,
) ValidationErrors {
	if settings == nil {
		settings = DefaultSettings()
	}

	ctx := NewValidationContext(settings, version)
	ctx.Package = pkg

	// Try to get parts from the package
	parts := getPackageParts(pkg)
	for _, part := range parts {
		if ctx.ShouldStop() {
			break
		}

		ctx.CurrentPart = part
		validatePart(ctx, part)
	}

	return ctx.Errors()
}

// ValidateElement validates a single element and its descendants.
func ValidateElement(
	element interface{},
	ctx *ValidationContext,
) []*ValidationError {
	if element == nil {
		return nil
	}

	var errors []*ValidationError

	// Get element info
	info := GetElementInfoFromInterface(element)
	if info != nil {
		ctx.PushPath(
			buildElementPathSegment(info),
		)
		defer ctx.PopPath()
	}

	// Check version availability
	if verr := CheckElementVersion(element, ctx.Version, ctx.CurrentPath()); verr != nil {
		errors = append(errors, verr)
		if !ctx.AddError(verr) {
			return errors
		}
	}

	// Schema validation
	if ctx.Settings.SchemaValidation {
		schemaErrs := validateElementSchema(
			ctx,
			element,
		)
		for _, err := range schemaErrs {
			errors = append(errors, err)
			if !ctx.AddError(err) {
				return errors
			}
		}
	}

	// Semantic validation
	if ctx.Settings.SemanticValidation {
		semanticErrs := validateElementSemantics(
			ctx,
			element,
		)
		for _, err := range semanticErrs {
			errors = append(errors, err)
			if !ctx.AddError(err) {
				return errors
			}
		}
	}

	// Validate children recursively
	children := getChildElements(element)
	for _, child := range children {
		if ctx.ShouldStop() {
			break
		}
		childErrs := ValidateElement(child, ctx)
		errors = append(errors, childErrs...)
	}

	return errors
}

// validatePart validates a document part.
func validatePart(
	ctx *ValidationContext,
	part interface{},
) {
	// Get root element from part
	root := getPartRootElement(part)
	if root == nil {
		return
	}

	// Clear IDs for this part (IDs are unique within parts)
	ctx.ClearIDs()

	// Validate the root element
	ValidateElement(root, ctx)
}

// validateElementSchema performs schema validation on an element.
func validateElementSchema(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	var errors []*ValidationError

	// Get registered schema validator for this element type
	validator := getSchemaValidator(element)
	if validator != nil {
		errors = validator.Validate(ctx, element)
	}

	return errors
}

// validateElementSemantics performs semantic validation on an element.
func validateElementSemantics(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	var errors []*ValidationError

	// Get registered constraints for this element type
	elementType := reflect.TypeOf(element)
	constraints := GetConstraints(elementType)
	for _, c := range constraints {
		errs := c.Check(ctx, element)
		errors = append(errors, errs...)
	}

	return errors
}

// Schema validator registry
var schemaValidatorRegistry = struct {
	mu         sync.RWMutex
	validators map[reflect.Type]*SchemaValidator
}{
	validators: make(
		map[reflect.Type]*SchemaValidator,
	),
}

// RegisterSchemaValidator registers a schema validator for an element type.
func RegisterSchemaValidator(
	elementType reflect.Type,
	validator *SchemaValidator,
) {
	schemaValidatorRegistry.mu.Lock()
	defer schemaValidatorRegistry.mu.Unlock()

	schemaValidatorRegistry.validators[elementType] = validator
}

// getSchemaValidator returns the registered schema validator for an element.
func getSchemaValidator(
	element interface{},
) *SchemaValidator {
	schemaValidatorRegistry.mu.RLock()
	defer schemaValidatorRegistry.mu.RUnlock()

	elementType := reflect.TypeOf(element)

	return schemaValidatorRegistry.validators[elementType]
}

// Helper functions for extracting information from elements

// GetElementInfoFromInterface extracts ElementInfo from an element.
func GetElementInfoFromInterface(
	element interface{},
) *ElementInfo {
	if element == nil {
		return nil
	}

	// Check for nil interface value
	v := reflect.ValueOf(element)
	if !v.IsValid() ||
		(v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil
	}

	// Try ValidatableElement interface
	if ve, ok := element.(ValidatableElement); ok {
		return &ElementInfo{
			LocalName:    ve.LocalName(),
			NamespaceURI: ve.NamespaceURI(),
			Element:      element,
		}
	}

	// Try using reflection to find LocalName and NamespaceURI methods
	// Keep the original value (don't dereference pointer) to access methods
	localNameMethod := v.MethodByName("LocalName")
	nsURIMethod := v.MethodByName("NamespaceURI")

	if localNameMethod.IsValid() &&
		nsURIMethod.IsValid() {
		localNameResult := localNameMethod.Call(
			nil,
		)
		nsURIResult := nsURIMethod.Call(nil)

		if len(localNameResult) > 0 &&
			len(nsURIResult) > 0 {
			localName, _ := localNameResult[0].Interface().(string)
			nsURI, _ := nsURIResult[0].Interface().(string)

			return &ElementInfo{
				LocalName:    localName,
				NamespaceURI: nsURI,
				Element:      element,
			}
		}
	}

	return nil
}

// GetChildElementInfos extracts ElementInfo for all children of an element.
func GetChildElementInfos(
	element interface{},
) []ElementInfo {
	children := getChildElements(element)
	infos := make([]ElementInfo, 0, len(children))

	for _, child := range children {
		if info := GetElementInfoFromInterface(child); info != nil {
			infos = append(infos, *info)
		}
	}

	return infos
}

// getChildElements extracts child elements from an element.
func getChildElements(
	element interface{},
) []interface{} {
	if element == nil {
		return nil
	}

	// Try ValidatableCompositeElement interface
	if vce, ok := element.(ValidatableCompositeElement); ok {
		return vce.Children()
	}

	// Try using reflection to find Children method
	v := reflect.ValueOf(element)
	childrenMethod := v.MethodByName("Children")
	if childrenMethod.IsValid() {
		results := childrenMethod.Call(nil)
		if len(results) > 0 {
			// Handle different return types
			result := results[0]
			switch result.Kind() {
			case reflect.Slice:
				children := make(
					[]interface{},
					result.Len(),
				)
				for i := 0; i < result.Len(); i++ {
					children[i] = result.Index(i).
						Interface()
				}

				return children
			case reflect.Func:
				// Handle iterator pattern (iter.Seq[Element])
				// This requires calling the iterator function
				return extractFromIterator(result)
			}
		}
	}

	return nil
}

// extractFromIterator extracts elements from a Go iterator function.
func extractFromIterator(
	iterFunc reflect.Value,
) []interface{} {
	if !iterFunc.IsValid() ||
		iterFunc.Kind() != reflect.Func {
		return nil
	}

	var results []interface{}

	// Create a yield function
	yieldType := iterFunc.Type().In(0)
	if yieldType.Kind() != reflect.Func {
		return nil
	}

	// Create the yield callback
	yieldFunc := reflect.MakeFunc(
		yieldType,
		func(args []reflect.Value) []reflect.Value {
			if len(args) > 0 {
				results = append(
					results,
					args[0].Interface(),
				)
			}

			return []reflect.Value{
				reflect.ValueOf(true),
			}
		},
	)

	// Call the iterator with our yield function
	iterFunc.Call([]reflect.Value{yieldFunc})

	return results
}

// GetElementAttributes extracts attributes from an element as a map.
func GetElementAttributes(
	element interface{},
) map[string]string {
	if element == nil {
		return nil
	}

	attrs := make(map[string]string)

	// Try using reflection to find Attributes method
	v := reflect.ValueOf(element)
	if !v.IsValid() ||
		(v.Kind() == reflect.Ptr && v.IsNil()) {
		return attrs
	}

	attrsMethod := v.MethodByName("Attributes")
	if attrsMethod.IsValid() {
		results := attrsMethod.Call(nil)
		if len(results) > 0 {
			result := results[0]
			if result.Kind() == reflect.Slice {
				for i := 0; i < result.Len(); i++ {
					attr := result.Index(i)

					// Handle both struct and pointer types
					if attr.Kind() == reflect.Interface {
						attr = attr.Elem()
					}

					// Try to get LocalName and Value methods
					var localNameMethod, valueMethod reflect.Value

					// First try on the value directly (for struct types)
					if attr.Kind() == reflect.Struct {
						// Need to get addressable value for methods with pointer receiver
						localNameMethod = attr.MethodByName(
							"LocalName",
						)
						valueMethod = attr.MethodByName(
							"Value",
						)

						// If methods not found, try getting address if possible
						if !localNameMethod.IsValid() &&
							attr.CanAddr() {
							localNameMethod = attr.Addr().
								MethodByName("LocalName")
							valueMethod = attr.Addr().
								MethodByName("Value")
						}
					} else {
						localNameMethod = attr.MethodByName("LocalName")
						valueMethod = attr.MethodByName("Value")
					}

					if localNameMethod.IsValid() &&
						valueMethod.IsValid() {
						localNameResult := localNameMethod.Call(
							nil,
						)
						valueResult := valueMethod.Call(
							nil,
						)

						if len(
							localNameResult,
						) > 0 &&
							len(valueResult) > 0 {
							name, _ := localNameResult[0].Interface().(string)
							value, _ := valueResult[0].Interface().(string)
							attrs[name] = value
						}
					}
				}
			}
		}
	}

	return attrs
}

// GetParentTypeName gets the type name of an element's parent.
func GetParentTypeName(
	element interface{},
) string {
	if element == nil {
		return ""
	}

	// Try ValidatableElement interface
	if ve, ok := element.(ValidatableElement); ok {
		parent := ve.Parent()
		if parent == nil {
			return ""
		}

		return reflect.TypeOf(parent).String()
	}

	// Try using reflection
	v := reflect.ValueOf(element)
	parentMethod := v.MethodByName("Parent")
	if parentMethod.IsValid() {
		results := parentMethod.Call(nil)
		if len(results) > 0 &&
			!results[0].IsNil() {
			return results[0].Type().String()
		}
	}

	return ""
}

// getPackageParts extracts parts from a package.
func getPackageParts(
	pkg interface{},
) []interface{} {
	if pkg == nil {
		return nil
	}

	// Try ValidatablePackage interface
	if vp, ok := pkg.(ValidatablePackage); ok {
		return vp.Parts()
	}

	// Try using reflection to find Parts method
	v := reflect.ValueOf(pkg)
	partsMethod := v.MethodByName("Parts")
	if partsMethod.IsValid() {
		results := partsMethod.Call(nil)
		if len(results) > 0 {
			result := results[0]
			switch result.Kind() {
			case reflect.Slice:
				parts := make(
					[]interface{},
					result.Len(),
				)
				for i := 0; i < result.Len(); i++ {
					parts[i] = result.Index(i).
						Interface()
				}

				return parts
			case reflect.Func:
				return extractFromIterator(result)
			}
		}
	}

	return nil
}

// getPartRootElement gets the root element from a part.
func getPartRootElement(
	part interface{},
) interface{} {
	if part == nil {
		return nil
	}

	// Try ValidatablePart interface
	if vp, ok := part.(ValidatablePart); ok {
		return vp.RootElement()
	}

	// Try using reflection
	v := reflect.ValueOf(part)
	rootMethod := v.MethodByName("RootElement")
	if rootMethod.IsValid() {
		results := rootMethod.Call(nil)
		if len(results) > 0 &&
			!results[0].IsNil() {
			return results[0].Interface()
		}
	}

	return nil
}

// buildElementPathSegment builds an XPath-like path segment for an element.
func buildElementPathSegment(
	info *ElementInfo,
) string {
	if info == nil {
		return ""
	}

	// Use namespace prefix if available, otherwise just local name
	if info.NamespaceURI != "" {
		prefix := getNamespacePrefix(
			info.NamespaceURI,
		)
		if prefix != "" {
			return prefix + ":" + info.LocalName
		}
	}

	return info.LocalName
}

// Common namespace prefixes for OOXML
var namespacePrefixes = map[string]string{
	"http://schemas.openxmlformats.org/wordprocessingml/2006/main":              "w",
	"http://schemas.openxmlformats.org/spreadsheetml/2006/main":                 "x",
	"http://schemas.openxmlformats.org/presentationml/2006/main":                "p",
	"http://schemas.openxmlformats.org/drawingml/2006/main":                     "a",
	"http://schemas.openxmlformats.org/drawingml/2006/picture":                  "pic",
	"http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing":    "wp",
	"http://schemas.openxmlformats.org/officeDocument/2006/relationships":       "r",
	"http://schemas.openxmlformats.org/package/2006/relationships":              "rel",
	"http://schemas.openxmlformats.org/officeDocument/2006/math":                "m",
	"http://schemas.openxmlformats.org/markup-compatibility/2006":               "mc",
	"http://schemas.microsoft.com/office/word/2010/wordml":                      "w14",
	"http://schemas.microsoft.com/office/word/2012/wordml":                      "w15",
	"http://schemas.microsoft.com/office/word/2015/wordml/symex":                "w16se",
	"http://schemas.microsoft.com/office/word/2018/wordml":                      "w16",
	"http://purl.org/dc/elements/1.1/":                                          "dc",
	"http://purl.org/dc/terms/":                                                 "dcterms",
	"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes":      "vt",
	"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties": "ep",
	"http://schemas.openxmlformats.org/package/2006/content-types":              "ct",
}

// getNamespacePrefix returns a known prefix for a namespace URI.
func getNamespacePrefix(
	namespaceURI string,
) string {
	if prefix, ok := namespacePrefixes[namespaceURI]; ok {
		return prefix
	}

	return ""
}

// BuildXPath builds an XPath-like string from an element to the root.
func BuildXPath(element interface{}) string {
	if element == nil {
		return "/"
	}

	var segments []string

	// Walk up the tree
	current := element
	for current != nil {
		info := GetElementInfoFromInterface(
			current,
		)
		if info != nil {
			segment := buildElementPathSegment(
				info,
			)
			// Prepend to build path from root to element
			segments = append(
				[]string{segment},
				segments...)
		}

		// Get parent
		parent := getParent(current)
		if parent == nil {
			break
		}
		current = parent
	}

	if len(segments) == 0 {
		return "/"
	}

	return "/" + joinStrings(segments, "/")
}

// getParent gets the parent of an element.
func getParent(element interface{}) interface{} {
	if element == nil {
		return nil
	}

	v := reflect.ValueOf(element)
	if !v.IsValid() ||
		(v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil
	}

	parentMethod := v.MethodByName("Parent")
	if parentMethod.IsValid() {
		results := parentMethod.Call(nil)
		if len(results) > 0 {
			result := results[0]
			// Check for nil in various forms
			if !result.IsValid() {
				return nil
			}
			// Check if the value is a nil pointer or nil interface
			switch result.Kind() {
			case reflect.Ptr, reflect.Interface:
				if result.IsNil() {
					return nil
				}
			}

			return result.Interface()
		}
	}

	return nil
}
