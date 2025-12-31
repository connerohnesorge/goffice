package validation

import (
	"strings"
	"sync"
)

// ValidationSettings configures the validation behavior.
type ValidationSettings struct {
	// MaxErrors is the maximum number of errors to collect before stopping.
	// 0 means unlimited.
	MaxErrors int
	// ContinueOnError indicates whether to continue validation after
	// finding an error.
	ContinueOnError bool
	// SemanticValidation enables semantic validation (business rules).
	// Default is true.
	SemanticValidation bool
	// SchemaValidation enables schema structure validation.
	// Default is true.
	SchemaValidation bool
	// ValidateReferences enables validation of relationship references.
	// Default is true.
	ValidateReferences bool
	// StrictMode enables strict validation that treats warnings as errors.
	StrictMode bool
}

// DefaultSettings returns the default validation settings.
func DefaultSettings() *ValidationSettings {
	return &ValidationSettings{
		MaxErrors:          0, // unlimited
		ContinueOnError:    true,
		SemanticValidation: true,
		SchemaValidation:   true,
		ValidateReferences: true,
		StrictMode:         false,
	}
}

// FastSettings returns settings optimized for quick validation.
// It stops at the first error and skips semantic validation.
func FastSettings() *ValidationSettings {
	return &ValidationSettings{
		MaxErrors:          1,
		ContinueOnError:    false,
		SemanticValidation: false,
		SchemaValidation:   true,
		ValidateReferences: false,
		StrictMode:         false,
	}
}

// StrictSettings returns settings for strict validation.
// Treats warnings as errors and enables all validation.
func StrictSettings() *ValidationSettings {
	return &ValidationSettings{
		MaxErrors:          0,
		ContinueOnError:    true,
		SemanticValidation: true,
		SchemaValidation:   true,
		ValidateReferences: true,
		StrictMode:         true,
	}
}

// ValidationContext holds the state during validation.
type ValidationContext struct {
	mu sync.Mutex

	// Settings for this validation run.
	Settings *ValidationSettings
	// Version is the target Office version for validation.
	Version FileFormatVersions
	// errors is the collection of validation errors found.
	errors ValidationErrors
	// pathStack is the current element path stack.
	pathStack []string
	// stopped indicates if validation should stop.
	stopped bool
	// seenIDs tracks unique IDs for duplicate checking.
	seenIDs map[string]any
	// Package is the package being validated (if any).
	// Uses any to avoid circular import with openxml package.
	Package any
	// CurrentPart is the current part being validated (if any).
	CurrentPart any
}

// NewValidationContext creates a new validation context with the given
// settings and version.
// nolint:revive // modifies-parameter: intentional nil check with default
func NewValidationContext(
	settings *ValidationSettings,
	version FileFormatVersions,
) *ValidationContext {
	if settings == nil {
		settings = DefaultSettings()
	}

	const defaultPathStackCap = 16

	return &ValidationContext{
		Settings: settings,
		Version:  version,
		errors:   make(ValidationErrors, 0),
		pathStack: make(
			[]string,
			0,
			defaultPathStackCap,
		),
		seenIDs: make(map[string]any),
	}
}

// AddError adds a validation error to the context.
// Returns true if validation should continue, false if it should stop.
func (c *ValidationContext) AddError(
	err *ValidationError,
) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.stopped {
		return false
	}

	// In strict mode, treat warnings as errors
	if c.Settings.StrictMode &&
		err.Severity == SeverityWarning {
		err.Severity = SeverityError
	}

	c.errors = append(c.errors, err)

	// Check if we should stop
	if c.Settings.MaxErrors > 0 &&
		len(
			c.errors,
		) >= c.Settings.MaxErrors {
		c.stopped = true

		return false
	}

	if !c.Settings.ContinueOnError &&
		err.Severity == SeverityError {
		c.stopped = true

		return false
	}

	return true
}

// AddErrors adds multiple validation errors to the context.
// Returns true if validation should continue, false if it should stop.
func (c *ValidationContext) AddErrors(
	errs ...*ValidationError,
) bool {
	for _, err := range errs {
		if !c.AddError(err) {
			return false
		}
	}

	return true
}

// Errors returns all collected validation errors.
func (c *ValidationContext) Errors() ValidationErrors {
	c.mu.Lock()
	defer c.mu.Unlock()
	result := make(
		ValidationErrors,
		len(c.errors),
	)
	copy(result, c.errors)

	return result
}

// ShouldStop returns true if validation should stop.
func (c *ValidationContext) ShouldStop() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.stopped
}

// Stop explicitly stops validation.
func (c *ValidationContext) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stopped = true
}

// PushPath pushes an element onto the path stack.
func (c *ValidationContext) PushPath(
	element string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pathStack = append(c.pathStack, element)
}

// PopPath pops the last element from the path stack.
func (c *ValidationContext) PopPath() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.pathStack) > 0 {
		c.pathStack = c.pathStack[:len(c.pathStack)-1]
	}
}

// CurrentPath returns the current XPath-like path.
func (c *ValidationContext) CurrentPath() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	const pathSep = "/"

	return pathSep + strings.Join(
		c.pathStack,
		pathSep,
	)
}

// PathWithIndex returns the current path with an index suffix
// for the given element.
func (c *ValidationContext) PathWithIndex(
	element string,
	index int,
) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	const pathSep = "/"

	basePath := pathSep + strings.Join(
		c.pathStack,
		pathSep,
	)
	if basePath == pathSep {
		return pathSep + element + "[" + itoa(
			index,
		) + "]"
	}

	return basePath + pathSep + element + "[" + itoa(
		index,
	) + "]"
}

// itoa converts an integer to a string without importing strconv.
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	if i < 0 {
		return "-" + uitoa(uint(-i))
	}

	return uitoa(uint(i))
}

// nolint:revive // modifies-parameter: loop variable modification is intentional
func uitoa(u uint) string {
	const (
		bufSize   = 20
		base      = 10
		zeroDigit = '0'
	)
	var buf [bufSize]byte
	i := len(buf)
	for u >= base {
		i--
		q := u / base
		buf[i] = byte(zeroDigit + u - q*base)
		u = q
	}
	i--
	buf[i] = byte(zeroDigit + u)

	return string(buf[i:])
}

// TrackID tracks a unique ID for duplicate checking.
// Returns the existing element if a duplicate, or nil if new.
func (c *ValidationContext) TrackID(
	id string,
	element any,
) any {
	c.mu.Lock()
	defer c.mu.Unlock()

	if existing, ok := c.seenIDs[id]; ok {
		return existing
	}
	c.seenIDs[id] = element

	return nil
}

// HasSeenID returns true if the ID has already been seen.
func (c *ValidationContext) HasSeenID(
	id string,
) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.seenIDs[id]

	return ok
}

// ClearIDs clears the tracked IDs (useful between parts).
func (c *ValidationContext) ClearIDs() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.seenIDs = make(map[string]any)
}

// Reset resets the context for reuse.
func (c *ValidationContext) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.errors = c.errors[:0]
	c.pathStack = c.pathStack[:0]
	c.stopped = false
	c.seenIDs = make(map[string]any)
	c.Package = nil
	c.CurrentPart = nil
}

// ErrorCount returns the number of errors collected.
func (c *ValidationContext) ErrorCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.errors)
}

// IsVersionAvailable checks if a feature is available in the current
// validation version.
func (c *ValidationContext) IsVersionAvailable(
	avail *VersionAvailability,
) bool {
	if avail == nil {
		return true
	}

	return avail.IsAvailableIn(c.Version)
}

// WithTargetVersion returns a new ValidationContext with the specified
// target version. This is useful for validating documents against specific
// Office versions.
func (c *ValidationContext) WithTargetVersion(
	version FileFormatVersions,
) *ValidationContext {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Version = version

	return c
}

// TargetVersion returns the target Office version for validation.
// This is an alias for the Version field to match the naming convention
// in the specification.
func (c *ValidationContext) TargetVersion() FileFormatVersions {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Version
}
