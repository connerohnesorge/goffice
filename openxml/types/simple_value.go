package types

// SimpleValue is the base interface for all typed attribute values.
// It provides a consistent way to check if a value is set, get its
// XML string representation, and parse from XML strings.
type SimpleValue interface {
	// HasValue returns true if the value is set, false if nil/unset.
	HasValue() bool

	// InnerText returns the string representation suitable for XML
	// serialization.
	// Returns an empty string if the value is not set.
	InnerText() string

	// SetInnerText parses a value from its XML string representation.
	// An empty string typically unsets the value.
	// Returns an error if the text cannot be parsed for the target type.
	SetInnerText(text string) error
}

// Resettable is an optional interface for types that support explicit
// nil/unset state.
type Resettable interface {
	// SetNil clears the value, making HasValue() return false.
	SetNil()
}
