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
	// ContinueOnError indicates whether to continue validation after finding an error.
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
	seenIDs map[string]interface{}
	// Package is the package being validated (if any).
	// Uses interface{} to avoid circular import with openxml package.
	Package interface{}
	// CurrentPart is the current part being validated (if any).
	CurrentPart interface{}
}

// NewValidationContext creates a new validation context with the given settings and version.
func NewValidationContext(
	settings *ValidationSettings,
	version FileFormatVersions,
) *ValidationContext {
	if settings == nil {
		settings = DefaultSettings()
	}
	return &ValidationContext{
		Settings:  settings,
		Version:   version,
		errors:    make(ValidationErrors, 0),
		pathStack: make([]string, 0, 16),
		seenIDs:   make(map[string]interface{}),
	}
}

// AddError adds a validation error to the context.
// Returns true if validation should continue, false if it should stop.
func (ctx *ValidationContext) AddError(
	err *ValidationError,
) bool {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	if ctx.stopped {
		return false
	}

	// In strict mode, treat warnings as errors
	if ctx.Settings.StrictMode &&
		err.Severity == SeverityWarning {
		err.Severity = SeverityError
	}

	ctx.errors = append(ctx.errors, err)

	// Check if we should stop
	if ctx.Settings.MaxErrors > 0 &&
		len(
			ctx.errors,
		) >= ctx.Settings.MaxErrors {
		ctx.stopped = true
		return false
	}

	if !ctx.Settings.ContinueOnError &&
		err.Severity == SeverityError {
		ctx.stopped = true
		return false
	}

	return true
}

// AddErrors adds multiple validation errors to the context.
// Returns true if validation should continue, false if it should stop.
func (ctx *ValidationContext) AddErrors(
	errs ...*ValidationError,
) bool {
	for _, err := range errs {
		if !ctx.AddError(err) {
			return false
		}
	}
	return true
}

// Errors returns all collected validation errors.
func (ctx *ValidationContext) Errors() ValidationErrors {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	result := make(
		ValidationErrors,
		len(ctx.errors),
	)
	copy(result, ctx.errors)
	return result
}

// ShouldStop returns true if validation should stop.
func (ctx *ValidationContext) ShouldStop() bool {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.stopped
}

// Stop explicitly stops validation.
func (ctx *ValidationContext) Stop() {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.stopped = true
}

// PushPath pushes an element onto the path stack.
func (ctx *ValidationContext) PushPath(
	element string,
) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.pathStack = append(ctx.pathStack, element)
}

// PopPath pops the last element from the path stack.
func (ctx *ValidationContext) PopPath() {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if len(ctx.pathStack) > 0 {
		ctx.pathStack = ctx.pathStack[:len(ctx.pathStack)-1]
	}
}

// CurrentPath returns the current XPath-like path.
func (ctx *ValidationContext) CurrentPath() string {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return "/" + strings.Join(ctx.pathStack, "/")
}

// PathWithIndex returns the current path with an index suffix for the given element.
func (ctx *ValidationContext) PathWithIndex(
	element string,
	index int,
) string {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	basePath := "/" + strings.Join(
		ctx.pathStack,
		"/",
	)
	if basePath == "/" {
		return "/" + element + "[" + itoa(
			index,
		) + "]"
	}
	return basePath + "/" + element + "[" + itoa(
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

func uitoa(u uint) string {
	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)
	return string(buf[i:])
}

// TrackID tracks a unique ID for duplicate checking.
// Returns the existing element if a duplicate, or nil if new.
func (ctx *ValidationContext) TrackID(
	id string,
	element interface{},
) interface{} {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	if existing, ok := ctx.seenIDs[id]; ok {
		return existing
	}
	ctx.seenIDs[id] = element
	return nil
}

// HasSeenID returns true if the ID has already been seen.
func (ctx *ValidationContext) HasSeenID(
	id string,
) bool {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	_, ok := ctx.seenIDs[id]
	return ok
}

// ClearIDs clears the tracked IDs (useful between parts).
func (ctx *ValidationContext) ClearIDs() {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.seenIDs = make(map[string]interface{})
}

// Reset resets the context for reuse.
func (ctx *ValidationContext) Reset() {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.errors = ctx.errors[:0]
	ctx.pathStack = ctx.pathStack[:0]
	ctx.stopped = false
	ctx.seenIDs = make(map[string]interface{})
	ctx.Package = nil
	ctx.CurrentPart = nil
}

// ErrorCount returns the number of errors collected.
func (ctx *ValidationContext) ErrorCount() int {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return len(ctx.errors)
}

// IsVersionAvailable checks if a feature is available in the current validation version.
func (ctx *ValidationContext) IsVersionAvailable(
	avail *VersionAvailability,
) bool {
	if avail == nil {
		return true
	}
	return avail.IsAvailableIn(ctx.Version)
}
