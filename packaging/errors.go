// Package packaging provides the OPC (Open Packaging Conventions) layer.
package packaging

import "errors"

// Package-level errors for OPC operations.
var (
	// ErrPartNotFound is returned when a requested part does not exist
	// in the package.
	ErrPartNotFound = errors.New(
		"packaging: part not found",
	)

	// ErrPartExists is returned when attempting to create a part that
	// already exists.
	ErrPartExists = errors.New(
		"packaging: part already exists",
	)

	// ErrInvalidURI is returned when a part URI is malformed or invalid.
	ErrInvalidURI = errors.New(
		"packaging: invalid part URI",
	)

	// ErrRelationshipNotFound is returned when a requested relationship
	// does not exist.
	ErrRelationshipNotFound = errors.New(
		"packaging: relationship not found",
	)

	// ErrRelationshipExists is returned when attempting to create a
	// relationship with a duplicate ID.
	ErrRelationshipExists = errors.New(
		"packaging: relationship already exists",
	)

	// ErrPackageClosed is returned when operations are attempted on a
	// closed package.
	ErrPackageClosed = errors.New(
		"packaging: package is closed",
	)

	// ErrReadOnly is returned when write operations are attempted on a
	// read-only package.
	ErrReadOnly = errors.New(
		"packaging: package is read-only",
	)

	// ErrInvalidPackage is returned when the package structure is invalid.
	ErrInvalidPackage = errors.New(
		"packaging: invalid package structure",
	)

	// ErrContentTypeNotFound is returned when no content type is registered
	// for a part.
	ErrContentTypeNotFound = errors.New(
		"packaging: content type not found",
	)
)
