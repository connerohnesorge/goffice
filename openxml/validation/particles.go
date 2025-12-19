package validation

// Particle represents a schema particle for validating element structure.
// Particles define the allowed structure of child elements within a parent.
type Particle interface {
	// Validate validates the given children against this particle.
	// Returns validation errors and the number of children consumed.
	Validate(
		ctx *ValidationContext,
		children []ElementInfo,
		path string,
	) (errors []*ValidationError, consumed int)

	// MinOccurs returns the minimum number of times this particle must occur.
	MinOccurs() int

	// MaxOccurs returns the maximum number of times this particle can occur.
	// -1 means unbounded.
	MaxOccurs() int

	// Description returns a human-readable description of this particle.
	Description() string
}

// ElementInfo provides information about an element for validation.
// This avoids direct dependency on the openxml.Element interface.
type ElementInfo struct {
	// LocalName is the local name of the element.
	LocalName string
	// NamespaceURI is the namespace URI of the element.
	NamespaceURI string
	// Element is the actual element (as interface{} to avoid circular imports).
	Element interface{}
}

// ElementMatcher matches elements by name and namespace.
type ElementMatcher struct {
	// LocalName is the expected local name.
	LocalName string
	// NamespaceURI is the expected namespace URI.
	NamespaceURI string
	// Availability is the version availability for this element.
	Availability *VersionAvailability
}

// Matches returns true if the given element matches this matcher.
func (m *ElementMatcher) Matches(
	info ElementInfo,
) bool {
	if m.LocalName != info.LocalName {
		return false
	}
	// Empty namespace in matcher means any namespace is OK
	if m.NamespaceURI != "" &&
		m.NamespaceURI != info.NamespaceURI {
		return false
	}

	return true
}

// QualifiedName returns the qualified name for error messages.
func (m *ElementMatcher) QualifiedName() string {
	if m.NamespaceURI == "" {
		return m.LocalName
	}

	return m.NamespaceURI + ":" + m.LocalName
}

// baseParticle provides common particle functionality.
type baseParticle struct {
	minOccurs int
	maxOccurs int
}

func (p *baseParticle) MinOccurs() int { return p.minOccurs }

func (p *baseParticle) MaxOccurs() int { return p.maxOccurs }

// ElementParticle matches a single element type.
type ElementParticle struct {
	baseParticle
	Matcher ElementMatcher
}

// NewElementParticle creates a particle that matches a single element.
func NewElementParticle(
	localName, namespaceURI string,
	minOccurs, maxOccurs int,
) *ElementParticle {
	return &ElementParticle{
		baseParticle: baseParticle{
			minOccurs: minOccurs,
			maxOccurs: maxOccurs,
		},
		Matcher: ElementMatcher{
			LocalName:    localName,
			NamespaceURI: namespaceURI,
		},
	}
}

// WithAvailability sets the version availability for this particle.
func (p *ElementParticle) WithAvailability(
	avail *VersionAvailability,
) *ElementParticle {
	p.Matcher.Availability = avail

	return p
}

// Validate validates children against this element particle.
func (p *ElementParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	var errors []*ValidationError
	consumed := 0

	// Check version availability
	if p.Matcher.Availability != nil &&
		!ctx.IsVersionAvailable(
			p.Matcher.Availability,
		) {
		// Element not available in this version, skip matching
		if p.minOccurs > 0 {
			return errors, 0
		}

		return nil, 0
	}

	// Count matching elements
	for consumed < len(children) && p.Matcher.Matches(children[consumed]) {
		consumed++
		// Check max occurs
		if p.maxOccurs >= 0 &&
			consumed > p.maxOccurs {
			errors = append(
				errors,
				NewValidationError(
					Schema_TooManyElements,
					"Too many occurrences of element "+p.Matcher.QualifiedName(),
					path,
					children[consumed-1].Element,
				),
			)
		}
	}

	// Check min occurs
	if consumed < p.minOccurs {
		errors = append(
			errors,
			NewValidationError(
				Schema_MissingRequiredElement,
				"Required element "+p.Matcher.QualifiedName()+" is missing (expected at least "+itoa(
					p.minOccurs,
				)+")",
				path,
				nil,
			),
		)
	}

	return errors, consumed
}

// Description returns a human-readable description.
func (p *ElementParticle) Description() string {
	return p.Matcher.QualifiedName()
}

// SequenceParticle requires children to appear in a specific order.
type SequenceParticle struct {
	baseParticle
	Particles []Particle
}

// NewSequenceParticle creates a sequence particle with the given child particles.
func NewSequenceParticle(
	minOccurs, maxOccurs int,
	particles ...Particle,
) *SequenceParticle {
	return &SequenceParticle{
		baseParticle: baseParticle{
			minOccurs: minOccurs,
			maxOccurs: maxOccurs,
		},
		Particles: particles,
	}
}

// Validate validates children against this sequence particle.
func (p *SequenceParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	var allErrors []*ValidationError
	totalConsumed := 0
	occurrences := 0

	// Try to match the sequence multiple times up to maxOccurs
	for !ctx.ShouldStop() {

		// Check if we've reached max occurrences
		if p.maxOccurs >= 0 &&
			occurrences >= p.maxOccurs {
			break
		}

		// Try to match all particles in sequence
		sequenceConsumed := 0
		sequenceErrors := make(
			[]*ValidationError,
			0,
		)
		remaining := children[totalConsumed:]

		matched := true
		for _, particle := range p.Particles {
			if len(remaining) == 0 &&
				particle.MinOccurs() > 0 {
				matched = false

				break
			}
			errs, consumed := particle.Validate(
				ctx,
				remaining,
				path,
			)
			sequenceErrors = append(
				sequenceErrors,
				errs...)
			remaining = remaining[consumed:]
			sequenceConsumed += consumed
		}

		// If nothing was consumed and we need more occurrences, break
		if sequenceConsumed == 0 && !matched {
			if occurrences < p.minOccurs {
				allErrors = append(
					allErrors,
					sequenceErrors...)
			}

			break
		}

		totalConsumed += sequenceConsumed
		occurrences++
		allErrors = append(
			allErrors,
			sequenceErrors...)

		// If nothing was consumed, we can't make progress
		if sequenceConsumed == 0 {
			break
		}
	}

	// Check min occurrences
	if occurrences < p.minOccurs {
		allErrors = append(
			allErrors,
			NewValidationError(
				Schema_MissingRequiredElement,
				"Sequence is required but not found (expected at least "+itoa(
					p.minOccurs,
				)+" occurrences)",
				path,
				nil,
			),
		)
	}

	return allErrors, totalConsumed
}

// Description returns a human-readable description.
func (p *SequenceParticle) Description() string {
	return "sequence"
}

// ChoiceParticle requires exactly one of the child particles to match.
type ChoiceParticle struct {
	baseParticle
	Particles []Particle
}

// NewChoiceParticle creates a choice particle with the given alternatives.
func NewChoiceParticle(
	minOccurs, maxOccurs int,
	particles ...Particle,
) *ChoiceParticle {
	return &ChoiceParticle{
		baseParticle: baseParticle{
			minOccurs: minOccurs,
			maxOccurs: maxOccurs,
		},
		Particles: particles,
	}
}

// Validate validates children against this choice particle.
func (p *ChoiceParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	var errors []*ValidationError
	totalConsumed := 0
	occurrences := 0

	// Try to match one of the choices multiple times up to maxOccurs
	for totalConsumed < len(children) {
		if ctx.ShouldStop() {
			break
		}

		// Check if we've reached max occurrences
		if p.maxOccurs >= 0 &&
			occurrences >= p.maxOccurs {
			break
		}

		// Try each alternative
		bestConsumed := 0
		var bestErrors []*ValidationError
		remaining := children[totalConsumed:]

		for _, particle := range p.Particles {
			errs, consumed := particle.Validate(
				ctx,
				remaining,
				path,
			)
			if consumed > bestConsumed ||
				(consumed == bestConsumed && len(errs) < len(bestErrors)) {
				bestConsumed = consumed
				bestErrors = errs
			}
		}

		// If nothing matched, stop
		if bestConsumed == 0 {
			break
		}

		totalConsumed += bestConsumed
		occurrences++
		errors = append(errors, bestErrors...)
	}

	// Check min occurrences
	if occurrences < p.minOccurs {
		errors = append(
			errors,
			NewValidationError(
				Schema_MissingRequiredElement,
				"Choice is required but no alternative matched (expected at least "+itoa(
					p.minOccurs,
				)+" occurrences)",
				path,
				nil,
			),
		)
	}

	return errors, totalConsumed
}

// Description returns a human-readable description.
func (p *ChoiceParticle) Description() string {
	return "choice"
}

// AllParticle requires all child particles to appear but in any order.
type AllParticle struct {
	baseParticle
	Particles []Particle
}

// NewAllParticle creates an all particle with the given child particles.
func NewAllParticle(
	minOccurs, maxOccurs int,
	particles ...Particle,
) *AllParticle {
	return &AllParticle{
		baseParticle: baseParticle{
			minOccurs: minOccurs,
			maxOccurs: maxOccurs,
		},
		Particles: particles,
	}
}

// Validate validates children against this all particle.
func (p *AllParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	var errors []*ValidationError
	totalConsumed := 0

	// Track which particles have been matched
	matched := make([]bool, len(p.Particles))
	remaining := make(
		[]ElementInfo,
		len(children),
	)
	copy(remaining, children)

	// Try to match each child to a particle
	for len(remaining) > 0 {
		if ctx.ShouldStop() {
			break
		}

		child := remaining[0]
		foundMatch := false

		for i, particle := range p.Particles {
			if matched[i] {
				continue
			}

			errs, consumed := particle.Validate(
				ctx,
				[]ElementInfo{child},
				path,
			)
			if consumed > 0 {
				matched[i] = true
				foundMatch = true
				totalConsumed++
				remaining = remaining[1:]
				errors = append(errors, errs...)

				break
			}
		}

		if !foundMatch {
			// Unexpected element
			break
		}
	}

	// Check that required particles were matched
	for i, particle := range p.Particles {
		if !matched[i] &&
			particle.MinOccurs() > 0 {
			errors = append(
				errors,
				NewValidationError(
					Schema_MissingRequiredElement,
					"Required element is missing (in 'all' group): "+particle.Description(),
					path,
					nil,
				),
			)
		}
	}

	return errors, totalConsumed
}

// Description returns a human-readable description.
func (p *AllParticle) Description() string {
	return "all"
}

// AnyParticle matches any element (wildcard).
type AnyParticle struct {
	baseParticle
	// Namespace specifies the allowed namespace (empty means any).
	Namespace string
	// ProcessContents indicates how to process the matched content.
	ProcessContents ProcessContentsMode
}

// ProcessContentsMode specifies how to process wildcard content.
type ProcessContentsMode int

const (
	// ProcessStrict means the element must be valid according to schema.
	ProcessStrict ProcessContentsMode = iota
	// ProcessLax means validate if schema is available, otherwise allow.
	ProcessLax
	// ProcessSkip means do not validate the content.
	ProcessSkip
)

// NewAnyParticle creates a wildcard particle.
func NewAnyParticle(
	minOccurs, maxOccurs int,
	namespace string,
	mode ProcessContentsMode,
) *AnyParticle {
	return &AnyParticle{
		baseParticle: baseParticle{
			minOccurs: minOccurs,
			maxOccurs: maxOccurs,
		},
		Namespace:       namespace,
		ProcessContents: mode,
	}
}

// Validate validates children against this any particle.
func (p *AnyParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	var errors []*ValidationError
	consumed := 0

	for consumed < len(children) {
		// Check max occurs
		if p.maxOccurs >= 0 &&
			consumed >= p.maxOccurs {
			break
		}

		// Check namespace constraint
		if p.Namespace != "" &&
			p.Namespace != "##any" {
			if p.Namespace == "##other" {
				// Must be from a different namespace (not implemented fully)
			} else if p.Namespace != children[consumed].NamespaceURI {
				break
			}
		}

		consumed++
	}

	// Check min occurs
	if consumed < p.minOccurs {
		errors = append(
			errors,
			NewValidationError(
				Schema_MissingRequiredElement,
				"Expected at least "+itoa(
					p.minOccurs,
				)+" element(s) matching wildcard",
				path,
				nil,
			),
		)
	}

	return errors, consumed
}

// Description returns a human-readable description.
func (p *AnyParticle) Description() string {
	if p.Namespace != "" {
		return "any (" + p.Namespace + ")"
	}

	return "any"
}

// EmptyParticle matches no children (element must have no children).
type EmptyParticle struct{}

// NewEmptyParticle creates an empty particle.
func NewEmptyParticle() *EmptyParticle {
	return &EmptyParticle{}
}

// Validate validates that there are no children.
func (p *EmptyParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	if len(children) > 0 {
		return []*ValidationError{
			NewValidationError(
				Schema_UnexpectedElement,
				"Element should have no children but found "+children[0].LocalName,
				path,
				children[0].Element,
			),
		}, 0
	}

	return nil, 0
}

// MinOccurs returns 1 for empty (the element itself must appear).
func (p *EmptyParticle) MinOccurs() int { return 1 }

// MaxOccurs returns 1 for empty.
func (p *EmptyParticle) MaxOccurs() int { return 1 }

// Description returns a human-readable description.
func (p *EmptyParticle) Description() string { return "empty" }

// TextOnlyParticle indicates the element can only contain text.
type TextOnlyParticle struct{}

// NewTextOnlyParticle creates a text-only particle.
func NewTextOnlyParticle() *TextOnlyParticle {
	return &TextOnlyParticle{}
}

// Validate validates that there are no element children (text is allowed).
func (p *TextOnlyParticle) Validate(
	ctx *ValidationContext,
	children []ElementInfo,
	path string,
) ([]*ValidationError, int) {
	// For text-only, we shouldn't have element children
	if len(children) > 0 {
		return []*ValidationError{
			NewValidationError(
				Schema_UnexpectedElement,
				"Element should contain only text but found child element "+children[0].LocalName,
				path,
				children[0].Element,
			),
		}, 0
	}

	return nil, 0
}

// MinOccurs returns 0 for text-only.
func (p *TextOnlyParticle) MinOccurs() int { return 0 }

// MaxOccurs returns 0 for text-only (no element children allowed).
func (p *TextOnlyParticle) MaxOccurs() int { return 0 }

// Description returns a human-readable description.
func (p *TextOnlyParticle) Description() string { return "text-only" }
