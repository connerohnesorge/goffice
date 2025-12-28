package spreadsheet

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
)

const minimalFixturePath = "../testdata/fixtures/minimal.xlsx"

// BenchmarkWorkbookCreation benchmarks creating a new workbook.
// Target: 10MB workbook in <1s
func BenchmarkWorkbookCreation(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-create-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	b.ResetTimer()
	for i := range b.N {
		testPath := filepath.Join(
			tmpDir,
			"bench_"+strconv.Itoa(i)+".xlsx",
		)
		doc, err := Create(
			testPath,
			DocTypeWorkbook,
		)
		if err != nil {
			b.Fatalf("Create() error = %v", err)
		}

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			b.Fatalf("AddSheet() error = %v", err)
		}

		// Add some content
		for row := 1; row <= 100; row++ {
			for col := 'A'; col <= 'J'; col++ {
				cellRef := string(
					col,
				) + strconv.Itoa(
					row,
				)
				_ = sheet.SetCellValue(
					cellRef,
					row*int(col-'A'+1),
				)
			}
		}

		if err := doc.Save(); err != nil {
			b.Fatalf("Save() error = %v", err)
		}
		_ = doc.Close()
	}
}

// BenchmarkLargeWorksheet benchmarks handling large worksheets.
// Target: 1M cells in <5s
func BenchmarkLargeWorksheet(b *testing.B) {
	if testing.Short() {
		b.Skip(
			"Skipping large worksheet benchmark in short mode",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-large-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	b.ResetTimer()
	for i := range b.N {
		testPath := filepath.Join(
			tmpDir,
			"large_"+strconv.Itoa(i)+".xlsx",
		)
		doc, err := Create(
			testPath,
			DocTypeWorkbook,
		)
		if err != nil {
			b.Fatalf("Create() error = %v", err)
		}

		sheet, err := doc.AddSheet("Large Sheet")
		if err != nil {
			b.Fatalf("AddSheet() error = %v", err)
		}

		// Add 1000 rows x 10 columns = 10,000 cells
		// For full benchmark, increase to 100,000 rows = 1M cells
		const numRows = 1000
		const numCols = 10

		for row := 1; row <= numRows; row++ {
			for col := range numCols {
				colLetter := string(
					rune('A' + col),
				)
				cellRef := colLetter + strconv.Itoa(
					row,
				)
				_ = sheet.SetCellValue(
					cellRef,
					row*col,
				)
			}
		}

		if err := doc.Save(); err != nil {
			b.Fatalf("Save() error = %v", err)
		}
		_ = doc.Close()
	}
	b.ReportMetric(float64(1000*10), "cells/op")
}

// BenchmarkValidation benchmarks document validation.
// Target: <2s for typical docs
func BenchmarkValidation(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-validate-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a document to validate
	testPath := filepath.Join(
		tmpDir,
		"validate.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}

	sheet, _ := doc.AddSheet("Sheet1")
	for row := 1; row <= 100; row++ {
		for col := 'A'; col <= 'E'; col++ {
			cellRef := string(
				col,
			) + strconv.Itoa(
				row,
			)
			_ = sheet.SetCellValue(cellRef, row)
		}
	}

	if err := doc.Save(); err != nil {
		b.Fatalf("Save() error = %v", err)
	}
	_ = doc.Close()

	// Benchmark validation
	b.ResetTimer()
	for range b.N {
		doc, err := Open(testPath, false)
		if err != nil {
			b.Fatalf("Open() error = %v", err)
		}

		_ = doc.Validate(validation.Office2016)
		_ = doc.Close()
	}
}

// BenchmarkSheetAccess benchmarks accessing sheets.
func BenchmarkSheetAccess(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-access-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a document with multiple sheets
	testPath := filepath.Join(
		tmpDir,
		"sheets.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}

	const numSheets = 10
	for i := range numSheets {
		sheet, _ := doc.AddSheet(
			"Sheet" + strconv.Itoa(i),
		)
		_ = sheet.SetCellValue("A1", i)
	}

	if err := doc.Save(); err != nil {
		b.Fatalf("Save() error = %v", err)
	}
	_ = doc.Close()

	// Benchmark sheet access
	doc2, err := Open(testPath, false)
	if err != nil {
		b.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	b.ResetTimer()
	for range b.N {
		for j := range numSheets {
			_, _ = doc2.Sheet(j)
		}
	}
}

// BenchmarkSheetByName benchmarks accessing sheets by name.
func BenchmarkSheetByName(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-byname-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a document with multiple sheets
	testPath := filepath.Join(
		tmpDir,
		"sheets.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}

	sheetNames := []string{
		"Alpha",
		"Beta",
		"Gamma",
		"Delta",
		"Epsilon",
	}
	for _, name := range sheetNames {
		sheet, _ := doc.AddSheet(name)
		_ = sheet.SetCellValue("A1", name)
	}

	if err := doc.Save(); err != nil {
		b.Fatalf("Save() error = %v", err)
	}
	_ = doc.Close()

	// Benchmark sheet access by name
	doc2, err := Open(testPath, false)
	if err != nil {
		b.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	b.ResetTimer()
	for range b.N {
		for _, name := range sheetNames {
			_, _ = doc2.SheetByName(name)
		}
	}
}

// BenchmarkCellValue benchmarks setting cell values.
func BenchmarkCellValue(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-cellvalue-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"cellvalue.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Values")
	if err != nil {
		b.Fatalf("AddSheet() error = %v", err)
	}

	b.ResetTimer()
	for i := range b.N {
		row := (i % 1000) + 1
		col := string(rune('A' + (i/1000)%26))
		cellRef := col + strconv.Itoa(row)
		_ = sheet.SetCellValue(cellRef, i)
	}
}

// BenchmarkCellFormula benchmarks setting cell formulas.
func BenchmarkCellFormula(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-formula-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"formula.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Formulas")
	if err != nil {
		b.Fatalf("AddSheet() error = %v", err)
	}

	b.ResetTimer()
	for i := range b.N {
		row := (i % 1000) + 1
		col := string(rune('A' + (i/1000)%26))
		cellRef := col + strconv.Itoa(row)
		_ = sheet.SetCellFormula(
			cellRef,
			"=A1+B1",
		)
	}
}

// BenchmarkSaveAs benchmarks saving documents.
func BenchmarkSaveAs(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-saveas-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a document
	sourcePath := filepath.Join(
		tmpDir,
		"source.xlsx",
	)
	doc, err := Create(
		sourcePath,
		DocTypeWorkbook,
	)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}

	sheet, _ := doc.AddSheet("Data")
	for row := 1; row <= 100; row++ {
		for col := 'A'; col <= 'J'; col++ {
			cellRef := string(
				col,
			) + strconv.Itoa(
				row,
			)
			_ = sheet.SetCellValue(cellRef, row)
		}
	}

	if err := doc.Save(); err != nil {
		b.Fatalf("Save() error = %v", err)
	}
	_ = doc.Close()

	// Benchmark SaveAs
	b.ResetTimer()
	for i := range b.N {
		doc, err := Open(sourcePath, true)
		if err != nil {
			b.Fatalf("Open() error = %v", err)
		}

		destPath := filepath.Join(
			tmpDir,
			"dest_"+strconv.Itoa(i)+".xlsx",
		)
		if err := doc.SaveAs(destPath); err != nil {
			b.Fatalf("SaveAs() error = %v", err)
		}
		_ = doc.Close()
	}
}

// BenchmarkOpenFile benchmarks opening files.
func BenchmarkOpenFile(b *testing.B) {
	if _, err := os.Stat(minimalFixturePath); os.IsNotExist(
		err,
	) {
		b.Skip(
			"Skipping: minimal.xlsx fixture not found",
		)
	}

	b.ResetTimer()
	for range b.N {
		doc, err := Open(
			minimalFixturePath,
			false,
		)
		if err != nil {
			b.Fatalf("Open() error = %v", err)
		}
		_ = doc.Close()
	}
}

// BenchmarkStreamWrite benchmarks writing to an io.Writer.
func BenchmarkStreamWrite(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		var buf bytes.Buffer
		doc, err := CreateFromStream(
			&buf,
			DocTypeWorkbook,
		)
		if err != nil {
			b.Fatalf(
				"CreateFromStream() error = %v",
				err,
			)
		}

		sheet, _ := doc.AddSheet("Stream")
		for row := 1; row <= 50; row++ {
			for col := 'A'; col <= 'E'; col++ {
				cellRef := string(
					col,
				) + strconv.Itoa(
					row,
				)
				_ = sheet.SetCellValue(
					cellRef,
					row,
				)
			}
		}

		_ = doc.Close()
	}
}

// BenchmarkAddSheet benchmarks adding sheets.
func BenchmarkAddSheet(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-addsheet-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"sheets.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	b.ResetTimer()
	for i := range b.N {
		_, _ = doc.AddSheet(
			"Sheet" + strconv.Itoa(i),
		)
	}
}

// BenchmarkSheetIteration benchmarks iterating over sheets.
func BenchmarkSheetIteration(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-iter-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create document with many sheets
	testPath := filepath.Join(tmpDir, "iter.xlsx")
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}

	const numSheets = 20
	for i := range numSheets {
		_, _ = doc.AddSheet(
			"Sheet" + strconv.Itoa(i),
		)
	}

	if err := doc.Save(); err != nil {
		b.Fatalf("Save() error = %v", err)
	}
	_ = doc.Close()

	doc2, err := Open(testPath, false)
	if err != nil {
		b.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	b.ResetTimer()
	for range b.N {
		count := 0
		for range doc2.Sheets() {
			count++
		}
		if count != numSheets {
			b.Errorf(
				"Expected %d sheets, got %d",
				numSheets,
				count,
			)
		}
	}
}

// BenchmarkMemoryUsage reports memory allocation during workbook creation.
func BenchmarkMemoryUsage(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-memory-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	b.ReportAllocs()
	b.ResetTimer()

	for i := range b.N {
		testPath := filepath.Join(
			tmpDir,
			"memory_"+strconv.Itoa(i)+".xlsx",
		)
		doc, err := Create(
			testPath,
			DocTypeWorkbook,
		)
		if err != nil {
			b.Fatalf("Create() error = %v", err)
		}

		sheet, _ := doc.AddSheet("Memory")
		for row := 1; row <= 100; row++ {
			for col := 'A'; col <= 'J'; col++ {
				cellRef := string(
					col,
				) + strconv.Itoa(
					row,
				)
				_ = sheet.SetCellValue(
					cellRef,
					"test data "+strconv.Itoa(
						row,
					),
				)
			}
		}

		if err := doc.Save(); err != nil {
			b.Fatalf("Save() error = %v", err)
		}
		_ = doc.Close()
	}
}

// BenchmarkDocumentTypes benchmarks creating different document types.
func BenchmarkDocumentTypes(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-types-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	types := []DocType{
		DocTypeWorkbook,
		DocTypeTemplate,
	}

	for _, docType := range types {
		b.Run(
			docType.String(),
			func(b *testing.B) {
				for i := range b.N {
					testPath := filepath.Join(
						tmpDir,
						docType.String()+"_"+strconv.Itoa(
							i,
						)+docType.Extension(),
					)
					doc, err := Create(
						testPath,
						docType,
					)
					if err != nil {
						b.Fatalf(
							"Create() error = %v",
							err,
						)
					}

					_, _ = doc.AddSheet("Test")

					if err := doc.Save(); err != nil {
						b.Fatalf(
							"Save() error = %v",
							err,
						)
					}
					_ = doc.Close()
				}
			},
		)
	}
}

// BenchmarkRoundtrip benchmarks the complete open-save-reopen cycle.
func BenchmarkRoundtrip(b *testing.B) {
	if _, err := os.Stat(minimalFixturePath); os.IsNotExist(
		err,
	) {
		b.Skip(
			"Skipping: minimal.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-roundtrip-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	b.ResetTimer()
	for i := range b.N {
		// Open
		doc, err := Open(minimalFixturePath, true)
		if err != nil {
			b.Fatalf("Open() error = %v", err)
		}

		// Save
		tempPath := filepath.Join(
			tmpDir,
			"roundtrip_"+strconv.Itoa(i)+".xlsx",
		)
		if err := doc.SaveAs(tempPath); err != nil {
			b.Fatalf("SaveAs() error = %v", err)
		}
		_ = doc.Close()

		// Reopen
		doc2, err := Open(tempPath, false)
		if err != nil {
			b.Fatalf("Reopen() error = %v", err)
		}
		_ = doc2.Close()
	}
}

// BenchmarkManySmallWrites benchmarks many small cell writes.
func BenchmarkManySmallWrites(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-smallwrites-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"smallwrites.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("Small Writes")

	b.ResetTimer()
	for i := range b.N {
		row := (i % 1000) + 1
		col := string(rune('A' + (i/1000)%26))
		cellRef := col + strconv.Itoa(row)
		_ = sheet.SetCellValue(cellRef, "small")
	}
}

// BenchmarkLargeStrings benchmarks writing large string values.
func BenchmarkLargeStrings(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-largestr-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"largestrings.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		b.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("Large Strings")

	// Create a large string
	largeString := string(make([]byte, 1000))
	for i := range largeString {
		largeString = largeString[:i] + "X" + largeString[i+1:]
	}

	b.ResetTimer()
	for i := range b.N {
		row := (i % 100) + 1
		cellRef := "A" + strconv.Itoa(row)
		_ = sheet.SetCellValue(
			cellRef,
			largeString,
		)
	}
}

// StreamingRowWriter provides streaming write capabilities for large datasets.
// This is a placeholder for the streaming row writer implementation.
type StreamingRowWriter struct {
	sheet  *Sheet
	rowNum int
}

// NewStreamingRowWriter creates a new streaming row writer.
func NewStreamingRowWriter(
	sheet *Sheet,
) *StreamingRowWriter {
	return &StreamingRowWriter{
		sheet:  sheet,
		rowNum: 1,
	}
}

// WriteRow writes a row of values.
func (s *StreamingRowWriter) WriteRow(
	values ...any,
) error {
	for col, val := range values {
		colLetter := string(rune('A' + col))
		cellRef := colLetter + strconv.Itoa(
			s.rowNum,
		)
		_ = s.sheet.SetCellValue(cellRef, val)
	}
	s.rowNum++

	return nil
}

// Flush flushes any buffered data.
func (*StreamingRowWriter) Flush() error {
	// In a full implementation, this would flush buffered rows
	return nil
}

// Close closes the streaming writer.
func (s *StreamingRowWriter) Close() error {
	return s.Flush()
}

// Ensure StreamingRowWriter implements io.Closer.
var _ io.Closer = (*StreamingRowWriter)(nil)

// BenchmarkStreamingRowWriter benchmarks the streaming row writer.
func BenchmarkStreamingRowWriter(b *testing.B) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-bench-streaming-*",
	)
	if err != nil {
		b.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	b.ResetTimer()
	for i := range b.N {
		testPath := filepath.Join(
			tmpDir,
			"streaming_"+strconv.Itoa(i)+".xlsx",
		)
		doc, err := Create(
			testPath,
			DocTypeWorkbook,
		)
		if err != nil {
			b.Fatalf("Create() error = %v", err)
		}

		sheet, _ := doc.AddSheet("Streaming")
		writer := NewStreamingRowWriter(sheet)

		// Write 1000 rows
		for row := range 1000 {
			_ = writer.WriteRow(
				row,
				row*2,
				row*3,
				row*4,
				row*5,
			)
		}

		_ = writer.Close()
		if err := doc.Save(); err != nil {
			b.Fatalf("Save() error = %v", err)
		}
		_ = doc.Close()
	}
	b.ReportMetric(1000, "rows/op")
}
