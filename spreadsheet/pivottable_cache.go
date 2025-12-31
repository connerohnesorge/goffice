// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements pivot table cache population logic.
package spreadsheet

import (
	"errors"
	"fmt"
	"strings"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// parseRangeReference parses a range reference string into sheet name and cell range.
// Accepts formats like "Sheet1!A1:D100", "'Sheet Name'!A1:D100", or "A1:D100" (current sheet).
//
//nolint:revive // cognitive-complexity: parsing logic requires complexity
func parseRangeReference(
	ref string,
) (sheetName, cellRange string, err error) {
	if ref == "" {
		return "", "", errors.New(
			"empty range reference",
		)
	}

	// Check for quoted sheet name
	if ref[0] == '\'' {
		// Find the closing quote, accounting for escaped quotes ('')
		pos := 1
		for pos < len(ref) {
			quotePos := strings.Index(
				ref[pos:],
				"'",
			)
			if quotePos == -1 {
				return "", "", fmt.Errorf(
					"unclosed quote in range reference: %s",
					ref,
				)
			}
			pos += quotePos
			// Check if this is an escaped quote ('')
			if pos+1 < len(ref) &&
				ref[pos+1] == '\'' {
				// Escaped quote, continue searching
				pos += 2

				continue
			}
			// Found the closing quote
			break
		}
		if pos >= len(ref) || ref[pos] != '\'' {
			return "", "", fmt.Errorf(
				"unclosed quote in range reference: %s",
				ref,
			)
		}
		if pos+1 >= len(ref) ||
			ref[pos+1] != '!' {
			return "", "", fmt.Errorf(
				"expected '!' after quoted sheet name: %s",
				ref,
			)
		}
		sheetName = ref[1:pos]
		cellRange = ref[pos+2:]

		return sheetName, cellRange, nil
	}

	// Check for unquoted sheet name
	bangIdx := strings.Index(ref, "!")
	if bangIdx != -1 {
		sheetName = ref[:bangIdx]
		cellRange = ref[bangIdx+1:]

		return sheetName, cellRange, nil
	}

	// No sheet name, just cell range
	return "", ref, nil
}

// extractHeaderRow extracts field names from the first row of the given range.
func extractHeaderRow(
	sheet *Sheet,
	rangeRef string,
) ([]string, error) {
	// Parse the range reference
	rng, err := ParseRangeRef(rangeRef)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse range reference: %w",
			err,
		)
	}

	// Extract headers from the first row
	var headers []string
	for col := rng.Start.Col; col <= rng.End.Col; col++ {
		cellRef := CellRef{
			Col: col,
			Row: rng.Start.Row,
		}
		cell := sheet.Cell(cellRef.String())
		if cell == nil {
			headers = append(
				headers,
				fmt.Sprintf("Column%d", col),
			)

			continue
		}

		// Get cell value as string
		value := cell.GetString()
		if value == "" {
			// Use default column name if empty
			headers = append(
				headers,
				fmt.Sprintf("Column%d", col),
			)
		} else {
			headers = append(headers, value)
		}
	}

	return headers, nil
}

// extractDataRows extracts all data rows from the given range, excluding the header row.
//
//nolint:revive // cognitive-complexity: function handles multiple edge cases for data extraction
func extractDataRows(
	sheet *Sheet,
	rangeRef string,
) ([][]any, error) {
	// Parse the range reference
	rng, err := ParseRangeRef(rangeRef)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse range reference: %w",
			err,
		)
	}

	// Extract data rows (skip first row which is the header)
	var rows [][]any
	for row := rng.Start.Row + 1; row <= rng.End.Row; row++ {
		var rowData []any
		for col := rng.Start.Col; col <= rng.End.Col; col++ {
			cellRef := CellRef{
				Col: col,
				Row: row,
			}
			cell := sheet.Cell(cellRef.String())
			if cell == nil {
				rowData = append(rowData, nil)

				continue
			}

			// Get cell value - try to determine type based on cell's data type
			cellType := cell.DataType()

			switch cellType {
			case elements.CellTypeBoolean:
				rowData = append(
					rowData,
					cell.GetBoolean(),
				)
			case elements.CellTypeSharedString,
				elements.CellTypeInlineString,
				elements.CellTypeFormulaString:
				rowData = append(
					rowData,
					cell.GetString(),
				)
			case elements.CellTypeError:
				rowData = append(
					rowData,
					cell.GetString(),
				)
			case elements.CellTypeDate:
				// Date values
				rowData = append(
					rowData,
					cell.GetString(),
				)
			case elements.CellTypeNumber:
				num := cell.GetNumber()
				if num == 0 &&
					cell.GetString() == "" {
					// Empty cell
					rowData = append(rowData, nil)
				} else {
					rowData = append(rowData, num)
				}
			default:
				// Empty or unknown type
				str := cell.GetString()
				if str == "" {
					rowData = append(rowData, nil)
				} else {
					rowData = append(rowData, str)
				}
			}
		}
		rows = append(rows, rowData)
	}

	return rows, nil
}

// populateCacheFields populates the cache field definitions in the cache definition.
func populateCacheFields(
	cacheDef *elements.PivotCacheDefinition,
	headers []string,
) error {
	if cacheDef == nil {
		return errors.New(
			"cache definition is nil",
		)
	}

	// Get or create cache fields collection
	cacheFields := cacheDef.GetOrCreateCacheFields()

	// Remove existing fields
	for field := range cacheFields.CacheFields() {
		cacheFields.RemoveChild(field)
	}

	// Add a cache field for each header
	for _, header := range headers {
		field := cacheFields.AddField(header)
		// Set field as database field (source from data)
		field.SetDatabaseField(true)
	}

	// Update the count attribute
	cacheFields.SetCount(uint32(len(headers)))

	return nil
}

// populateCacheRecords populates the cache records with data rows.
func populateCacheRecords(
	cacheRecords *elements.PivotCacheRecords,
	rows [][]any,
	headers []string,
) error {
	if cacheRecords == nil {
		return errors.New("cache records is nil")
	}

	// Remove existing records
	for record := range cacheRecords.Records() {
		cacheRecords.RemoveChild(record)
	}

	// Add a record for each data row
	for _, rowData := range rows {
		record := cacheRecords.AddRecord()

		// Add field values to the record
		for i, value := range rowData {
			if i >= len(headers) {
				break // Skip extra columns beyond headers
			}

			if value == nil {
				// Add missing value
				record.AddMissing()

				continue
			}

			// Add value based on type
			switch v := value.(type) {
			case bool:
				record.AddBoolean(v)
			case float64:
				record.AddNumber(v)
			case int:
				record.AddNumber(float64(v))
			case string:
				record.AddString(v)
			default:
				// Convert to string as fallback
				record.AddString(fmt.Sprintf("%v", v))
			}
		}
	}

	// Update the count attribute
	cacheRecords.SetCount(uint32(len(rows)))

	return nil
}

// getCacheDefinitionPart returns the pivot cache definition part associated with this pivot table.
func (p *PivotTable) getCacheDefinitionPart() (*parts.PivotTableCacheDefinitionPart, error) {
	if p.pivotPart == nil {
		return nil, errors.New(
			"pivot table part is nil",
		)
	}

	// Check if cache definition part already exists
	cachePart := p.pivotPart.PivotTableCacheDefinitionPart()
	if cachePart != nil {
		return cachePart, nil
	}

	// Create new cache definition part
	cachePart, err := p.pivotPart.AddPivotTableCacheDefinitionPart()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create cache definition part: %w",
			err,
		)
	}

	return cachePart, nil
}

// getCacheRecordsPart returns the pivot cache records part associated with this pivot table.
func (p *PivotTable) getCacheRecordsPart() (*parts.PivotTableCacheRecordsPart, error) {
	// First get the cache definition part
	cacheDef, err := p.getCacheDefinitionPart()
	if err != nil {
		return nil, err
	}

	// Check if cache records part already exists
	recordsPart := cacheDef.PivotTableCacheRecordsPart()
	if recordsPart != nil {
		return recordsPart, nil
	}

	// Create new cache records part
	recordsPart, err = cacheDef.AddPivotTableCacheRecordsPart()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create cache records part: %w",
			err,
		)
	}

	return recordsPart, nil
}

// Refresh re-populates the pivot cache from the source range.
// This should be called after the source data has been modified
// to update the pivot table cache.
//
//nolint:revive // function-length: refresh requires multiple steps
func (p *PivotTable) Refresh() error {
	if p.sourceRange == "" {
		return fmt.Errorf("source range is empty")
	}

	// Parse the source range to extract sheet name and cell range
	sheetName, cellRange, err := parseRangeReference(
		p.sourceRange,
	)
	if err != nil {
		return fmt.Errorf(
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
			return fmt.Errorf("source sheet not found: %s", sheetName)
		}
	}

	// Build full range reference with sheet name for parsing
	fullRangeRef := cellRange
	if sheetName != "" {
		// Quote sheet name if needed
		if strings.Contains(sheetName, " ") ||
			strings.Contains(sheetName, "'") {
			fullRangeRef = "'" + sheetName + "'!" + cellRange
		} else {
			fullRangeRef = sheetName + "!" + cellRange
		}
	}

	// Extract headers and data
	headers, err := extractHeaderRow(
		sourceSheet,
		fullRangeRef,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to extract headers: %w",
			err,
		)
	}

	rows, err := extractDataRows(
		sourceSheet,
		fullRangeRef,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to extract data rows: %w",
			err,
		)
	}

	// Get or create cache definition part
	cacheDefPart, err := p.getCacheDefinitionPart()
	if err != nil {
		return fmt.Errorf(
			"failed to get cache definition part: %w",
			err,
		)
	}

	cacheDef := cacheDefPart.PivotCacheDefinition()
	if cacheDef == nil {
		cacheDef = elements.NewPivotCacheDefinition()
		cacheDefPart.SetRootElement(cacheDef)
	}

	// Populate cache fields
	if err := populateCacheFields(cacheDef, headers); err != nil {
		return fmt.Errorf(
			"failed to populate cache fields: %w",
			err,
		)
	}

	// Update cache source
	cacheSource := cacheDef.GetOrCreateCacheSource()
	cacheSource.SetType("worksheet")
	wsSource := cacheSource.GetOrCreateWorksheetSource()
	wsSource.SetRef(cellRange)
	if sheetName != "" {
		wsSource.SetSheet(sheetName)
	}

	// Update record count
	cacheDef.SetRecordCount(uint32(len(rows)))

	// Get or create cache records part
	recordsPart, err := p.getCacheRecordsPart()
	if err != nil {
		return fmt.Errorf(
			"failed to get cache records part: %w",
			err,
		)
	}

	cacheRecords := recordsPart.PivotCacheRecords()
	if cacheRecords == nil {
		cacheRecords = elements.NewPivotCacheRecords()
		recordsPart.SetRootElement(cacheRecords)
	}

	// Populate cache records
	if err := populateCacheRecords(cacheRecords, rows, headers); err != nil {
		return fmt.Errorf(
			"failed to populate cache records: %w",
			err,
		)
	}

	// Update the relationship from cache definition to cache records
	relID := recordsPart.RelationshipID()
	if relID != "" {
		cacheDef.SetRelationshipId(relID)
	}

	return nil
}
