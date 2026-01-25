package core

import (
	"errors"
	"fmt"
)

// Error variables for common error conditions.
var (
	// ErrDocumentClosed indicates an operation was attempted on a closed document.
	ErrDocumentClosed = errors.New(
		"core: document is closed",
	)

	// ErrDocumentNotWritable indicates the document cannot be written to.
	ErrDocumentNotWritable = errors.New(
		"core: document is not writable",
	)

	// ErrInvalidPageSize indicates an invalid page size was specified.
	ErrInvalidPageSize = errors.New(
		"core: invalid page size",
	)

	// ErrInvalidVersion indicates an invalid PDF version was specified.
	ErrInvalidVersion = errors.New(
		"core: invalid PDF version",
	)

	// ErrNilWriter indicates a nil writer was passed to a write operation.
	ErrNilWriter = errors.New(
		"core: writer is nil",
	)

	// ErrPageNotFound indicates the requested page does not exist.
	ErrPageNotFound = errors.New(
		"core: page not found",
	)

	// ErrInvalidPageIndex indicates an invalid page index was specified.
	ErrInvalidPageIndex = errors.New(
		"core: invalid page index",
	)

	// ErrNoPages indicates the document has no pages.
	ErrNoPages = errors.New(
		"core: document has no pages",
	)
)

// DocumentError represents an error that occurred during document processing.
type DocumentError struct {
	Op  string // The operation that failed.
	Err error  // The underlying error.
}

// Error implements the error interface.
func (e *DocumentError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(
			"core: %s: %v",
			e.Op,
			e.Err,
		)
	}

	return fmt.Sprintf("core: %s", e.Op)
}

// Unwrap returns the underlying error for errors.Is and errors.As support.
func (e *DocumentError) Unwrap() error {
	return e.Err
}

// NewDocumentError creates a new DocumentError with the given operation and error.
func NewDocumentError(
	op string,
	err error,
) *DocumentError {
	return &DocumentError{
		Op:  op,
		Err: err,
	}
}

// PageError represents an error that occurred during page processing.
type PageError struct {
	PageNum int    // The page number (1-based) where the error occurred.
	Op      string // The operation that failed.
	Err     error  // The underlying error.
}

// Error implements the error interface.
func (e *PageError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(
			"core: page %d: %s: %v",
			e.PageNum,
			e.Op,
			e.Err,
		)
	}

	return fmt.Sprintf(
		"core: page %d: %s",
		e.PageNum,
		e.Op,
	)
}

// Unwrap returns the underlying error for errors.Is and errors.As support.
func (e *PageError) Unwrap() error {
	return e.Err
}

// NewPageError creates a new PageError with the given page number, operation, and error.
func NewPageError(
	pageNum int,
	op string,
	err error,
) *PageError {
	return &PageError{
		PageNum: pageNum,
		Op:      op,
		Err:     err,
	}
}

// WriteError represents an error that occurred during PDF writing.
type WriteError struct {
	Op  string // The write operation that failed.
	Err error  // The underlying error.
}

// Error implements the error interface.
func (e *WriteError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(
			"core: write %s: %v",
			e.Op,
			e.Err,
		)
	}

	return fmt.Sprintf(
		"core: write %s failed",
		e.Op,
	)
}

// Unwrap returns the underlying error for errors.Is and errors.As support.
func (e *WriteError) Unwrap() error {
	return e.Err
}

// NewWriteError creates a new WriteError with the given operation and error.
func NewWriteError(
	op string,
	err error,
) *WriteError {
	return &WriteError{
		Op:  op,
		Err: err,
	}
}
