package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// TestParseRangeReference tests parsing of range references.
func TestParseRangeReference(t *testing.T) {
	tests := []struct {
		name          string
		ref           string
		wantSheet     string
		wantCellRange string
		wantErr       bool
	}{
		{
			name:          "simple range no sheet",
			ref:           "A1:D100",
			wantSheet:     "",
			wantCellRange: "A1:D100",
			wantErr:       false,
		},
		{
			name:          "range with unquoted sheet",
			ref:           "Sheet1!A1:D100",
			wantSheet:     "Sheet1",
			wantCellRange: "A1:D100",
			wantErr:       false,
		},
		{
			name:          "range with quoted sheet",
			ref:           "'Sales Data'!A1:D100",
			wantSheet:     "Sales Data",
			wantCellRange: "A1:D100",
			wantErr:       false,
		},
		{
			name:          "range with single quotes in name",
			ref:           "'Bob''s Sheet'!A1:C10",
			wantSheet:     "Bob''s Sheet",
			wantCellRange: "A1:C10",
			wantErr:       false,
		},
		{
			name:      "empty reference",
			ref:       "",
			wantSheet: "",
			wantErr:   true,
		},
		{
			name:      "unclosed quote",
			ref:       "'Sheet!A1:D100",
			wantSheet: "",
			wantErr:   true,
		},
		{
			name:      "missing bang after quote",
			ref:       "'Sheet'A1:D100",
			wantSheet: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSheet, gotCellRange, err := parseRangeReference(
				tt.ref,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"parseRangeReference() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if err != nil {
				return
			}
			if gotSheet != tt.wantSheet {
				t.Errorf(
					"parseRangeReference() gotSheet = %v, want %v",
					gotSheet,
					tt.wantSheet,
				)
			}
			if gotCellRange != tt.wantCellRange {
				t.Errorf(
					"parseRangeReference() gotCellRange = %v, want %v",
					gotCellRange,
					tt.wantCellRange,
				)
			}
		})
	}
}

// TestExtractHeaderRow tests extracting headers from the first row.
func TestExtractHeaderRow(t *testing.T) {
	// Create a test document
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sheet, err := doc.AddSheet("TestSheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Populate header row
	sheet.Cell("A1").SetString("Product")
	sheet.Cell("B1").SetString("Category")
	sheet.Cell("C1").SetString("Sales")
	sheet.Cell("D1").SetString("Quantity")

	tests := []struct {
		name        string
		rangeRef    string
		wantHeaders []string
		wantErr     bool
	}{
		{
			name:     "extract full header row",
			rangeRef: "A1:D10",
			wantHeaders: []string{
				"Product",
				"Category",
				"Sales",
				"Quantity",
			},
			wantErr: false,
		},
		{
			name:     "extract partial header row",
			rangeRef: "B1:C5",
			wantHeaders: []string{
				"Category",
				"Sales",
			},
			wantErr: false,
		},
		{
			name:        "single column",
			rangeRef:    "A1:A10",
			wantHeaders: []string{"Product"},
			wantErr:     false,
		},
		{
			name:     "invalid range",
			rangeRef: "invalid",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeaders, err := extractHeaderRow(
				sheet,
				tt.rangeRef,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"extractHeaderRow() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if err != nil {
				return
			}
			if len(
				gotHeaders,
			) != len(
				tt.wantHeaders,
			) {
				t.Errorf(
					"extractHeaderRow() got %d headers, want %d",
					len(gotHeaders),
					len(tt.wantHeaders),
				)

				return
			}
			for i, header := range gotHeaders {
				if header != tt.wantHeaders[i] {
					t.Errorf(
						"extractHeaderRow() header[%d] = %v, want %v",
						i,
						header,
						tt.wantHeaders[i],
					)

					break
				}
			}
		})
	}
}

// TestExtractDataRows tests extracting data rows excluding the header.
func TestExtractDataRows(t *testing.T) {
	// Create a test document
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sheet, err := doc.AddSheet("TestSheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Populate header row
	sheet.Cell("A1").SetString("Product")
	sheet.Cell("B1").SetString("Sales")
	sheet.Cell("C1").SetString("InStock")

	// Populate data rows
	sheet.Cell("A2").SetString("Widget")
	sheet.Cell("B2").SetNumber(1250.50)
	sheet.Cell("C2").SetBoolean(true)

	sheet.Cell("A3").SetString("Gadget")
	sheet.Cell("B3").SetNumber(850.25)
	sheet.Cell("C3").SetBoolean(false)

	sheet.Cell("A4").SetString("Tool")
	sheet.Cell("B4").SetNumber(500)
	// C4 is empty

	tests := []struct {
		name     string
		rangeRef string
		wantRows int
		wantErr  bool
	}{
		{
			name:     "extract all data rows",
			rangeRef: "A1:C4",
			wantRows: 3, // Excludes header row
			wantErr:  false,
		},
		{
			name:     "extract partial range",
			rangeRef: "A1:B3",
			wantRows: 2,
			wantErr:  false,
		},
		{
			name:     "single row",
			rangeRef: "A1:C2",
			wantRows: 1,
			wantErr:  false,
		},
		{
			name:     "invalid range",
			rangeRef: "invalid",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRows, err := extractDataRows(
				sheet,
				tt.rangeRef,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"extractDataRows() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if err != nil {
				return
			}
			if len(gotRows) != tt.wantRows {
				t.Errorf(
					"extractDataRows() got %d rows, want %d",
					len(gotRows),
					tt.wantRows,
				)
			}
		})
	}
}

// TestExtractDataRowsTypes tests that data types are correctly extracted.
func TestExtractDataRowsTypes(t *testing.T) {
	// Create a test document
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sheet, err := doc.AddSheet("TestSheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Populate header row
	sheet.Cell("A1").SetString("Text")
	sheet.Cell("B1").SetString("Number")
	sheet.Cell("C1").SetString("Empty")

	// Populate data row with different types
	sheet.Cell("A2").SetString("Hello")
	sheet.Cell("B2").SetNumber(42.5)
	// C2 is empty

	rows, err := extractDataRows(sheet, "A1:C2")
	if err != nil {
		t.Fatalf(
			"extractDataRows() error = %v",
			err,
		)
	}

	if len(rows) != 1 {
		t.Fatalf(
			"extractDataRows() got %d rows, want 1",
			len(rows),
		)
	}

	row := rows[0]
	if len(row) != 3 {
		t.Fatalf(
			"extractDataRows() got %d columns, want 3",
			len(row),
		)
	}

	// Check text value
	if str, ok := row[0].(string); !ok ||
		str != "Hello" {
		t.Errorf(
			"row[0] = %v (type %T), want string 'Hello'",
			row[0],
			row[0],
		)
	}

	// Check numeric value
	if num, ok := row[1].(float64); !ok ||
		num != 42.5 {
		t.Errorf(
			"row[1] = %v (type %T), want float64 42.5",
			row[1],
			row[1],
		)
	}

	// Check nil value for empty cell
	if row[2] != nil {
		t.Errorf("row[2] = %v, want nil", row[2])
	}
}

// TestPopulateCacheFields tests populating cache field definitions.
func TestPopulateCacheFields(t *testing.T) {
	cacheDef := elements.NewPivotCacheDefinition()

	headers := []string{
		"Product",
		"Category",
		"Sales",
		"Quantity",
	}

	err := populateCacheFields(cacheDef, headers)
	if err != nil {
		t.Fatalf(
			"populateCacheFields() error = %v",
			err,
		)
	}

	cacheFields := cacheDef.CacheFields()
	if cacheFields == nil {
		t.Fatal("CacheFields() returned nil")
	}

	// Check count attribute
	if cacheFields.Count() != uint32(
		len(headers),
	) {
		t.Errorf(
			"CacheFields.Count() = %d, want %d",
			cacheFields.Count(),
			len(headers),
		)
	}

	// Check field count
	fieldCount := cacheFields.FieldCount()
	if fieldCount != len(headers) {
		t.Errorf(
			"FieldCount() = %d, want %d",
			fieldCount,
			len(headers),
		)
	}

	// Check each field name
	i := 0
	for field := range cacheFields.CacheFields() {
		if i >= len(headers) {
			t.Errorf(
				"Too many fields: got %d, want %d",
				i+1,
				len(headers),
			)

			break
		}
		if field.Name() != headers[i] {
			t.Errorf(
				"Field[%d].Name() = %v, want %v",
				i,
				field.Name(),
				headers[i],
			)
		}
		i++
	}
}

// TestPopulateCacheRecords tests populating cache records with data.
func TestPopulateCacheRecords(t *testing.T) {
	cacheRecords := elements.NewPivotCacheRecords()

	headers := []string{
		"Product",
		"Sales",
		"InStock",
	}
	rows := [][]any{
		{"Widget", 1250.5, true},
		{"Gadget", 850.25, false},
		{
			"Tool",
			500.0,
			nil,
		}, // nil for empty cell
	}

	err := populateCacheRecords(
		cacheRecords,
		rows,
		headers,
	)
	if err != nil {
		t.Fatalf(
			"populateCacheRecords() error = %v",
			err,
		)
	}

	// Check count attribute
	if cacheRecords.Count() != uint32(len(rows)) {
		t.Errorf(
			"CacheRecords.Count() = %d, want %d",
			cacheRecords.Count(),
			len(rows),
		)
	}

	// Check record count
	recordCount := cacheRecords.RecordCount()
	if recordCount != len(rows) {
		t.Errorf(
			"RecordCount() = %d, want %d",
			recordCount,
			len(rows),
		)
	}
}

// TestPopulateCacheRecordsNilInput tests error handling for nil input.
func TestPopulateCacheRecordsNilInput(
	t *testing.T,
) {
	err := populateCacheRecords(nil, nil, nil)
	if err == nil {
		t.Error(
			"populateCacheRecords(nil) should return error",
		)
	}
}

// TestPopulateCacheFieldsNilInput tests error handling for nil input.
func TestPopulateCacheFieldsNilInput(
	t *testing.T,
) {
	err := populateCacheFields(nil, nil)
	if err == nil {
		t.Error(
			"populateCacheFields(nil) should return error",
		)
	}
}

// TestRefreshRoundtrip tests the full refresh cycle.
func TestRefreshRoundtrip(t *testing.T) {
	// Create a test document with source data
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sheet, err := doc.AddSheet("Sales")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Populate source data
	sheet.Cell("A1").SetString("Product")
	sheet.Cell("B1").SetString("Category")
	sheet.Cell("C1").SetString("Sales")

	sheet.Cell("A2").SetString("Widget")
	sheet.Cell("B2").SetString("Electronics")
	sheet.Cell("C2").SetNumber(1250.50)

	sheet.Cell("A3").SetString("Gadget")
	sheet.Cell("B3").SetString("Electronics")
	sheet.Cell("C3").SetNumber(850.25)

	sheet.Cell("A4").SetString("Tool")
	sheet.Cell("B4").SetString("Hardware")
	sheet.Cell("C4").SetNumber(500)

	// Create a pivot table
	pivotPart, err := sheet.worksheetPart.AddPivotTablePart()
	if err != nil {
		t.Fatalf(
			"AddPivotTablePart() error = %v",
			err,
		)
	}

	destCell := MustParseCellRef("E1")
	pivot, err := newPivotTable(
		sheet,
		pivotPart,
		"A1:C4",
		destCell,
	)
	if err != nil {
		t.Fatalf(
			"newPivotTable() error = %v",
			err,
		)
	}

	// Refresh the pivot cache
	err = pivot.Refresh()
	if err != nil {
		t.Fatalf("Refresh() error = %v", err)
	}

	// Verify cache definition was created
	cacheDef, err := pivot.getCacheDefinitionPart()
	if err != nil {
		t.Fatalf(
			"getCacheDefinitionPart() error = %v",
			err,
		)
	}
	if cacheDef == nil {
		t.Fatal("Cache definition part is nil")
	}

	// Verify cache fields
	cacheDefElem := cacheDef.PivotCacheDefinition()
	if cacheDefElem == nil {
		t.Fatal(
			"PivotCacheDefinition element is nil",
		)
	}

	cacheFields := cacheDefElem.CacheFields()
	if cacheFields == nil {
		t.Fatal("CacheFields is nil")
	}

	if cacheFields.Count() != 3 {
		t.Errorf(
			"CacheFields.Count() = %d, want 3",
			cacheFields.Count(),
		)
	}

	// Verify cache records
	recordsPart, err := pivot.getCacheRecordsPart()
	if err != nil {
		t.Fatalf(
			"getCacheRecordsPart() error = %v",
			err,
		)
	}
	if recordsPart == nil {
		t.Fatal("Cache records part is nil")
	}

	cacheRecords := recordsPart.PivotCacheRecords()
	if cacheRecords == nil {
		t.Fatal(
			"PivotCacheRecords element is nil",
		)
	}

	if cacheRecords.Count() != 3 {
		t.Errorf(
			"CacheRecords.Count() = %d, want 3 (excludes header)",
			cacheRecords.Count(),
		)
	}

	// Verify record count in definition
	if cacheDefElem.RecordCount() != 3 {
		t.Errorf(
			"PivotCacheDefinition.RecordCount() = %d, want 3",
			cacheDefElem.RecordCount(),
		)
	}
}

// TestRefreshWithSheetReference tests refresh with explicit sheet reference.
func TestRefreshWithSheetReference(t *testing.T) {
	// Create a test document with source data on different sheet
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sourceSheet, err := doc.AddSheet("SourceData")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}
	pivotSheet, err := doc.AddSheet("PivotSheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Populate source data
	sourceSheet.Cell("A1").SetString("Product")
	sourceSheet.Cell("B1").SetString("Sales")

	sourceSheet.Cell("A2").SetString("Widget")
	sourceSheet.Cell("B2").SetNumber(1000)

	sourceSheet.Cell("A3").SetString("Gadget")
	sourceSheet.Cell("B3").SetNumber(2000)

	// Create a pivot table on different sheet with cross-sheet reference
	pivotPart, err := pivotSheet.worksheetPart.AddPivotTablePart()
	if err != nil {
		t.Fatalf(
			"AddPivotTablePart() error = %v",
			err,
		)
	}

	destCell := MustParseCellRef("A1")
	pivot, err := newPivotTable(
		pivotSheet,
		pivotPart,
		"SourceData!A1:B3",
		destCell,
	)
	if err != nil {
		t.Fatalf(
			"newPivotTable() error = %v",
			err,
		)
	}

	// Refresh the pivot cache
	err = pivot.Refresh()
	if err != nil {
		t.Fatalf("Refresh() error = %v", err)
	}

	// Verify cache was populated
	cacheDef, err := pivot.getCacheDefinitionPart()
	if err != nil {
		t.Fatalf(
			"getCacheDefinitionPart() error = %v",
			err,
		)
	}

	cacheDefElem := cacheDef.PivotCacheDefinition()
	if cacheDefElem == nil {
		t.Fatal(
			"PivotCacheDefinition element is nil",
		)
	}

	// Verify record count
	if cacheDefElem.RecordCount() != 2 {
		t.Errorf(
			"RecordCount() = %d, want 2",
			cacheDefElem.RecordCount(),
		)
	}

	// Verify cache source has sheet reference
	cacheSource := cacheDefElem.CacheSource()
	if cacheSource == nil {
		t.Fatal("CacheSource is nil")
	}

	wsSource := cacheSource.WorksheetSource()
	if wsSource == nil {
		t.Fatal("WorksheetSource is nil")
	}

	if wsSource.Sheet() != "SourceData" {
		t.Errorf(
			"WorksheetSource.Sheet() = %v, want 'SourceData'",
			wsSource.Sheet(),
		)
	}

	if wsSource.Ref() != "A1:B3" {
		t.Errorf(
			"WorksheetSource.Ref() = %v, want 'A1:B3'",
			wsSource.Ref(),
		)
	}
}

// TestRefreshEmptySourceRange tests error handling for empty source range.
func TestRefreshEmptySourceRange(t *testing.T) {
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sheet, err := doc.AddSheet("TestSheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	pivotPart, err := sheet.worksheetPart.AddPivotTablePart()
	if err != nil {
		t.Fatalf(
			"AddPivotTablePart() error = %v",
			err,
		)
	}

	destCell := MustParseCellRef("A1")
	pivot, err := newPivotTable(
		sheet,
		pivotPart,
		"",
		destCell,
	)
	if err != nil {
		t.Fatalf(
			"newPivotTable() error = %v",
			err,
		)
	}

	err = pivot.Refresh()
	if err == nil {
		t.Error(
			"Refresh() with empty source range should return error",
		)
	}
}

// TestRefreshInvalidSheetReference tests error handling for non-existent sheet.
func TestRefreshInvalidSheetReference(
	t *testing.T,
) {
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close(); _ = os.Remove(tmpFile) }()

	sheet, err := doc.AddSheet("TestSheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	pivotPart, err := sheet.worksheetPart.AddPivotTablePart()
	if err != nil {
		t.Fatalf(
			"AddPivotTablePart() error = %v",
			err,
		)
	}

	destCell := MustParseCellRef("A1")
	pivot, err := newPivotTable(
		sheet,
		pivotPart,
		"NonExistentSheet!A1:B10",
		destCell,
	)
	if err != nil {
		t.Fatalf(
			"newPivotTable() error = %v",
			err,
		)
	}

	err = pivot.Refresh()
	if err == nil {
		t.Error(
			"Refresh() with non-existent sheet should return error",
		)
	}
}
