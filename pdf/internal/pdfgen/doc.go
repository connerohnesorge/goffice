//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package pdfgen provides a PDF generation abstraction layer over pdfcpu.
//
// This internal package isolates the dependency on pdfcpu and provides a
// simplified API tailored to the needs of the goffice PDF renderers. If a
// different PDF library is needed in the future, only this package needs
// to be modified.
//
// # Responsibilities
//
// This package handles:
//   - PDF document structure creation
//   - Page addition and management
//   - Content stream generation
//   - Font embedding and subsetting
//   - Image embedding and compression
//   - Resource dictionary management
//   - Cross-reference table generation
//   - PDF stream compression
//
// # Content Streams
//
// PDF content streams use a stack-based graphics model with operators for:
//   - Text (Tj, TJ, Tf, Tm, etc.)
//   - Graphics (m, l, c, h, re, S, f, etc.)
//   - Color (rg, RG, k, K, cs, CS, sc, SC)
//   - State (q, Q, cm, w, J, j, M, d)
//
// This package provides a builder pattern for constructing these streams:
//
//	stream := pdfgen.NewContentStream()
//	stream.SaveState()
//	stream.SetFont("F1", 12)
//	stream.MoveTo(72, 720)
//	stream.ShowText("Hello")
//	stream.RestoreState()
//
// # Thread Safety
//
// Document construction is not thread-safe. Each document should be built
// from a single goroutine. However, multiple documents can be built
// concurrently in separate goroutines.
//
// # Dependencies
//
// This package uses pdfcpu (github.com/pdfcpu/pdfcpu) for low-level PDF
// generation. pdfcpu is an MIT-licensed, actively maintained PDF library
// that provides comprehensive PDF manipulation capabilities.
package pdfgen

import (
	// pdfcpu provides the underlying PDF generation capabilities.
	// The pdfcpu/pkg/pdfcpu package contains the core PDF model types.
	_ "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)
