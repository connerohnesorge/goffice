package validation

import (
	"reflect"
	"sync"
)

// Validator is the interface for all validators.
type Validator interface {
	// Validate validates the given element within the context.
	// Returns a slice of validation errors (may be empty).
	Validate(
		ctx *ValidationContext,
		element interface{},
	) []*ValidationError
}

// Constraint is the interface for custom validation constraints.
// Constraints are registered per element type and checked during validation.
type Constraint interface {
	// Check checks the constraint on the given element.
	// Returns validation errors if the constraint is violated.
	Check(
		ctx *ValidationContext,
		element interface{},
	) []*ValidationError
}

// SchemaValidator validates element structure against schema particles.
type SchemaValidator struct {
	// Particle is the schema particle defining the allowed structure.
	Particle Particle
	// AllowedAttributes lists the allowed attributes.
	AllowedAttributes []AttributeSchema
	// RequiredAttributes lists attributes that must be present.
	RequiredAttributes []string
}

// AttributeSchema defines the schema for an attribute.
type AttributeSchema struct {
	// LocalName is the attribute's local name.
	LocalName string
	// NamespaceURI is the attribute's namespace URI (empty for no namespace).
	NamespaceURI string
	// Required indicates if this attribute is required.
	Required bool
	// DefaultValue is the default value if not specified.
	DefaultValue string
	// Type is the type name for validation.
	Type string
	// Availability specifies version availability.
	Availability *VersionAvailability
	// Values lists allowed values for enumeration types.
	Values []string
	// MinValue is the minimum value for numeric types.
	MinValue *int64
	// MaxValue is the maximum value for numeric types.
	MaxValue *int64
}

// NewSchemaValidator creates a new schema validator with the given particle.
func NewSchemaValidator(
	particle Particle,
) *SchemaValidator {
	return &SchemaValidator{
		Particle: particle,
		AllowedAttributes: make(
			[]AttributeSchema,
			0,
		),
		RequiredAttributes: make([]string, 0),
	}
}

// WithAttribute adds an allowed attribute to the schema.
func (v *SchemaValidator) WithAttribute(
	attr AttributeSchema,
) *SchemaValidator {
	v.AllowedAttributes = append(
		v.AllowedAttributes,
		attr,
	)
	if attr.Required {
		v.RequiredAttributes = append(
			v.RequiredAttributes,
			attr.LocalName,
		)
	}
	return v
}

// Validate validates an element against the schema.
func (v *SchemaValidator) Validate(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	var errors []*ValidationError

	// Get element info using reflection or interface
	info := GetElementInfoFromInterface(element)
	if info == nil {
		return errors
	}

	// Validate child structure
	if v.Particle != nil {
		children := GetChildElementInfos(element)
		childErrors, _ := v.Particle.Validate(
			ctx,
			children,
			ctx.CurrentPath(),
		)
		errors = append(errors, childErrors...)
	}

	// Validate attributes
	attrErrors := v.validateAttributes(
		ctx,
		element,
	)
	errors = append(errors, attrErrors...)

	return errors
}

// validateAttributes validates the element's attributes.
func (v *SchemaValidator) validateAttributes(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	var errors []*ValidationError

	// Get attributes from element
	attrs := GetElementAttributes(element)
	attrMap := make(map[string]string)
	for name, value := range attrs {
		attrMap[name] = value
	}

	// Check required attributes
	for _, required := range v.RequiredAttributes {
		if _, ok := attrMap[required]; !ok {
			errors = append(
				errors,
				NewValidationError(
					Schema_MissingRequiredAttribute,
					"Required attribute '"+required+"' is missing",
					ctx.CurrentPath(),
					element,
				),
			)
		}
	}

	// Validate each attribute against its schema
	for _, schema := range v.AllowedAttributes {
		value, ok := attrMap[schema.LocalName]
		if !ok {
			continue
		}

		// Check version availability
		if schema.Availability != nil &&
			!ctx.IsVersionAvailable(
				schema.Availability,
			) {
			errors = append(
				errors,
				NewValidationError(
					Schema_AttributeNotAvailable,
					"Attribute '"+schema.LocalName+"' is not available in "+ctx.Version.String(),
					ctx.CurrentPath(),
					element,
				),
			)
			continue
		}

		// Check enumeration values
		if len(schema.Values) > 0 {
			found := false
			for _, allowed := range schema.Values {
				if value == allowed {
					found = true
					break
				}
			}
			if !found {
				errors = append(
					errors,
					NewValidationError(
						Schema_ValueNotInEnumeration,
						"Attribute '"+schema.LocalName+"' value '"+value+"' is not in allowed values",
						ctx.CurrentPath(),
						element,
					),
				)
			}
		}
	}

	return errors
}

// AttributeValidator validates individual attribute values.
type AttributeValidator struct {
	// Schema is the attribute schema.
	Schema AttributeSchema
}

// NewAttributeValidator creates a new attribute validator.
func NewAttributeValidator(
	schema AttributeSchema,
) *AttributeValidator {
	return &AttributeValidator{Schema: schema}
}

// Validate validates an attribute value.
func (v *AttributeValidator) Validate(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	var errors []*ValidationError

	attrs := GetElementAttributes(element)
	value, ok := attrs[v.Schema.LocalName]
	if !ok {
		if v.Schema.Required {
			errors = append(
				errors,
				NewValidationError(
					Schema_MissingRequiredAttribute,
					"Required attribute '"+v.Schema.LocalName+"' is missing",
					ctx.CurrentPath(),
					element,
				),
			)
		}
		return errors
	}

	// Check enumeration
	if len(v.Schema.Values) > 0 {
		found := false
		for _, allowed := range v.Schema.Values {
			if value == allowed {
				found = true
				break
			}
		}
		if !found {
			errors = append(
				errors,
				NewValidationError(
					Schema_ValueNotInEnumeration,
					"Attribute '"+v.Schema.LocalName+"' value '"+value+"' is not in allowed values",
					ctx.CurrentPath(),
					element,
				),
			)
		}
	}

	return errors
}

// SemanticValidator validates business rules and semantic constraints.
type SemanticValidator struct {
	// constraints is the list of constraints to check.
	constraints []Constraint
}

// NewSemanticValidator creates a new semantic validator.
func NewSemanticValidator() *SemanticValidator {
	return &SemanticValidator{
		constraints: make([]Constraint, 0),
	}
}

// AddConstraint adds a constraint to check.
func (v *SemanticValidator) AddConstraint(
	c Constraint,
) *SemanticValidator {
	v.constraints = append(v.constraints, c)
	return v
}

// Validate validates an element against semantic constraints.
func (v *SemanticValidator) Validate(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	var errors []*ValidationError

	for _, c := range v.constraints {
		errs := c.Check(ctx, element)
		errors = append(errors, errs...)
	}

	return errors
}

// Constraint registry for per-element-type constraints
var constraintRegistry = struct {
	mu          sync.RWMutex
	constraints map[reflect.Type][]Constraint
}{
	constraints: make(
		map[reflect.Type][]Constraint,
	),
}

// RegisterConstraint registers a constraint for a specific element type.
func RegisterConstraint(
	elementType reflect.Type,
	constraint Constraint,
) {
	constraintRegistry.mu.Lock()
	defer constraintRegistry.mu.Unlock()

	constraintRegistry.constraints[elementType] = append(
		constraintRegistry.constraints[elementType],
		constraint,
	)
}

// GetConstraints returns all registered constraints for the given element type.
func GetConstraints(
	elementType reflect.Type,
) []Constraint {
	constraintRegistry.mu.RLock()
	defer constraintRegistry.mu.RUnlock()

	return constraintRegistry.constraints[elementType]
}

// ClearConstraints clears all registered constraints (useful for testing).
func ClearConstraints() {
	constraintRegistry.mu.Lock()
	defer constraintRegistry.mu.Unlock()

	constraintRegistry.constraints = make(
		map[reflect.Type][]Constraint,
	)
}

// Built-in constraint implementations

// UniqueIDConstraint ensures unique IDs within a scope.
type UniqueIDConstraint struct {
	// GetID extracts the ID from an element.
	GetID func(element interface{}) string
}

// NewUniqueIDConstraint creates a constraint that checks for unique IDs.
func NewUniqueIDConstraint(
	getID func(element interface{}) string,
) *UniqueIDConstraint {
	return &UniqueIDConstraint{GetID: getID}
}

// Check checks that the element's ID is unique.
func (c *UniqueIDConstraint) Check(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	id := c.GetID(element)
	if id == "" {
		return nil
	}

	if existing := ctx.TrackID(id, element); existing != nil {
		return []*ValidationError{
			NewValidationError(
				Semantic_DuplicateID,
				"Duplicate ID '"+id+"' found",
				ctx.CurrentPath(),
				element,
			).WithRelatedInfo("First occurrence at previously validated element"),
		}
	}

	return nil
}

// RelationshipReferenceConstraint ensures referenced relationships exist.
type RelationshipReferenceConstraint struct {
	// GetRelID extracts the relationship ID from an element.
	GetRelID func(element interface{}) string
	// CheckExists checks if the relationship exists.
	CheckExists func(ctx *ValidationContext, relID string) bool
}

// NewRelationshipReferenceConstraint creates a constraint that checks relationship references.
func NewRelationshipReferenceConstraint(
	getRelID func(element interface{}) string,
	checkExists func(ctx *ValidationContext, relID string) bool,
) *RelationshipReferenceConstraint {
	return &RelationshipReferenceConstraint{
		GetRelID:    getRelID,
		CheckExists: checkExists,
	}
}

// Check checks that the referenced relationship exists.
func (c *RelationshipReferenceConstraint) Check(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	relID := c.GetRelID(element)
	if relID == "" {
		return nil
	}

	if !c.CheckExists(ctx, relID) {
		return []*ValidationError{
			NewValidationError(
				Semantic_RelationshipNotFound,
				"Referenced relationship '"+relID+"' not found",
				ctx.CurrentPath(),
				element,
			),
		}
	}

	return nil
}

// MutuallyExclusiveConstraint ensures certain attributes don't coexist.
type MutuallyExclusiveConstraint struct {
	// Attributes is the list of mutually exclusive attribute names.
	Attributes []string
}

// NewMutuallyExclusiveConstraint creates a constraint for mutually exclusive attributes.
func NewMutuallyExclusiveConstraint(
	attrs ...string,
) *MutuallyExclusiveConstraint {
	return &MutuallyExclusiveConstraint{
		Attributes: attrs,
	}
}

// Check checks that no more than one of the mutually exclusive attributes is present.
func (c *MutuallyExclusiveConstraint) Check(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	attrs := GetElementAttributes(element)
	present := make([]string, 0)

	for _, attrName := range c.Attributes {
		if _, ok := attrs[attrName]; ok {
			present = append(present, attrName)
		}
	}

	if len(present) > 1 {
		return []*ValidationError{
			NewValidationError(
				Semantic_MutuallyExclusiveAttributes,
				"Mutually exclusive attributes are present: "+joinStrings(
					present,
					", ",
				),
				ctx.CurrentPath(),
				element,
			),
		}
	}

	return nil
}

// ParentTypeConstraint ensures an element has a specific parent type.
type ParentTypeConstraint struct {
	// AllowedParents is the list of allowed parent type names.
	AllowedParents []string
}

// NewParentTypeConstraint creates a constraint for allowed parent types.
func NewParentTypeConstraint(
	parents ...string,
) *ParentTypeConstraint {
	return &ParentTypeConstraint{
		AllowedParents: parents,
	}
}

// Check checks that the element's parent is of an allowed type.
func (c *ParentTypeConstraint) Check(
	ctx *ValidationContext,
	element interface{},
) []*ValidationError {
	parentName := GetParentTypeName(element)
	if parentName == "" {
		return nil // No parent or can't determine
	}

	for _, allowed := range c.AllowedParents {
		if parentName == allowed {
			return nil
		}
	}

	return []*ValidationError{NewValidationError(
		Semantic_InvalidParentType,
		"Element cannot have parent of type '"+parentName+"', expected one of: "+joinStrings(
			c.AllowedParents,
			", ",
		),
		ctx.CurrentPath(),
		element,
	)}
}

// joinStrings joins strings with a separator.
func joinStrings(
	strs []string,
	sep string,
) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}
