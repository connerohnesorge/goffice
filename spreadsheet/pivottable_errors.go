// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements error types for PivotTable validation.
package spreadsheet

import (
	"errors"
	"fmt"
)

var (
	// ErrFieldNotFound indicates that the specified field name does not exist
	// in the source data headers.
	ErrFieldNotFound = errors.New(
		"field not found in source data",
	)

	// ErrDuplicateField indicates that the field is already assigned to an axis
	// (row, column, page, or data area).
	ErrDuplicateField = errors.New(
		"field already exists on an axis",
	)

	// ErrInvalidSourceRange indicates that the source range reference is invalid
	// or cannot be parsed.
	ErrInvalidSourceRange = errors.New(
		"invalid source range",
	)
)

// validateFieldName checks if the field name exists in the pivot cache headers.
// If the cache is not populated yet, it attempts to extract headers from the source range.
func (p *PivotTable) validateFieldName(
	name string,
) error {
	if name == "" {
		return fmt.Errorf(
			"%w: field name is empty",
			ErrFieldNotFound,
		)
	}

	// Try to get headers from cache definition
	headers, err := p.getCachedHeaders()
	if err != nil {
		// Cache not populated yet, try to extract headers from source range
		headers, err = p.extractSourceHeaders()
		if err != nil {
			return fmt.Errorf(
				"%w: failed to extract headers: %v",
				ErrInvalidSourceRange,
				err,
			)
		}
	}

	// Check if the field name exists in headers
	for _, header := range headers {
		if header == name {
			return nil
		}
	}

	return fmt.Errorf(
		"%w: '%s'",
		ErrFieldNotFound,
		name,
	)
}

// getCachedHeaders retrieves field names from the pivot cache definition.
func (p *PivotTable) getCachedHeaders() ([]string, error) {
	cacheDefPart, err := p.getCacheDefinitionPart()
	if err != nil {
		return nil, err
	}

	cacheDef := cacheDefPart.PivotCacheDefinition()
	if cacheDef == nil {
		return nil, errors.New(
			"cache definition is nil",
		)
	}

	cacheFields := cacheDef.CacheFields()
	if cacheFields == nil {
		return nil, errors.New(
			"cache fields is nil",
		)
	}

	var headers []string
	for field := range cacheFields.CacheFields() {
		name := field.Name()
		if name != "" {
			headers = append(headers, name)
		}
	}

	if len(headers) == 0 {
		return nil, errors.New(
			"no fields in cache",
		)
	}

	return headers, nil
}

// extractSourceHeaders extracts headers from the source range without populating the cache.
func (p *PivotTable) extractSourceHeaders() ([]string, error) {
	if p.sourceRange == "" {
		return nil, errors.New(
			"source range is empty",
		)
	}

	// Parse the source range to extract sheet name and cell range
	sheetName, cellRange, err := parseRangeReference(
		p.sourceRange,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse source range: %w",
			err,
		)
	}

	// Get the source sheet
	var sourceSheet *Sheet
	if sheetName == "" {
		// Use the current sheet
		sourceSheet = p.sheet
	} else {
		// Find the sheet by name
		var found bool
		sourceSheet, found = p.sheet.doc.SheetByName(sheetName)
		if !found {
			return nil, fmt.Errorf("source sheet not found: %s", sheetName)
		}
	}

	// Extract headers from the first row
	headers, err := extractHeaderRow(
		sourceSheet,
		cellRange,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to extract headers: %w",
			err,
		)
	}

	return headers, nil
}

// containsField checks if the specified field name is already assigned to any axis
// (row, column, page, or data area).
func (p *PivotTable) containsField(
	name string,
) bool {
	// Check row fields
	for _, field := range p.rowFields {
		if field == name {
			return true
		}
	}

	// Check column fields
	for _, field := range p.colFields {
		if field == name {
			return true
		}
	}

	// Check page fields
	for _, field := range p.pageFields {
		if field == name {
			return true
		}
	}

	// Check data fields
	for _, field := range p.dataFields {
		if field.name == name {
			return true
		}
	}

	return false
}
