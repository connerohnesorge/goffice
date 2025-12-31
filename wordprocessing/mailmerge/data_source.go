// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package mailmerge provides mail merge functionality for Word documents.
//
// Mail merge allows you to create personalized documents by combining a template
// document with data from various sources (CSV files, JSON files, or in-memory maps).
//
// # Basic Usage
//
//	doc, _ := wordprocessing.Open("template.docx", true)
//	result, err := mailmerge.New(doc).
//		DataSource(mailmerge.NewCSVDataSource("contacts.csv")).
//		Execute()
//	if err != nil {
//		log.Fatal(err)
//	}
//	result.Save()
//
// # Data Sources
//
// The package supports three built-in data sources:
//   - CSVDataSource: reads data from CSV files with headers in the first row
//   - JSONDataSource: reads data from JSON array files (array of objects)
//   - MapDataSource: uses in-memory slice of maps for programmatic data
//
// You can also implement custom data sources by implementing the DataSource interface.
//
// # Field Types
//
// Word documents can contain merge fields in two formats:
//   - SimpleField: single XML element (w:fldSimple) with field code as attribute
//   - ComplexField: span of elements with begin/separate/end markers (w:fldChar)
//
// Both formats are automatically detected and replaced during merge.
//
// # Advanced Fields
//
// The package supports special field types:
//   - GREETINGLINE: generates personalized greetings (e.g., "Dear John Smith,")
//   - SKIPIF: conditionally skips records based on field value comparisons
//
// # Merge Options
//
// Configure merge behavior using MergeOptions:
//   - StrictFields: whether to error on missing fields (default: false)
//   - RemoveUnusedFields: whether to remove unreplaced fields (default: true)
//
// Example:
//
//	opts := &mailmerge.MergeOptions{
//		StrictFields:       true,  // error on missing fields
//		RemoveUnusedFields: true,  // remove unreplaced fields
//	}
//	result, _ := mailmerge.New(doc).Options(opts).DataSource(ds).Execute()
//
// # OOXML Metadata Integration
//
// The package automatically reads mail merge metadata from the document's settings.xml
// (w:mailMerge element) to support field name mappings and data source configurations
// created by Microsoft Word. This ensures compatibility with templates created in Word.
package mailmerge

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// DataSource is the interface for mail merge data sources.
// It provides methods to iterate through records and retrieve field values.
type DataSource interface {
	// Open initializes the data source and prepares it for reading.
	// It must be called before any other operations.
	Open() error

	// Close releases any resources held by the data source.
	// It should be called when done reading the data source.
	Close() error

	// Next advances to the next record in the data source.
	// It returns true if there is a next record, false if no more records exist.
	// Must call Open() before calling Next().
	Next() bool

	// Fields returns the list of field names available in the data source.
	// The order may vary between implementations.
	Fields() []string

	// Get retrieves the value for the specified field name in the current record.
	// Field name lookup is case-insensitive.
	// Returns an error if the field does not exist or if Next() has not been called.
	Get(fieldName string) (string, error)
}

// CSVDataSource is a DataSource implementation that reads from a CSV file.
//
// The first row of the CSV file must contain headers that define the field names.
// Subsequent rows contain the data records. Field name matching is case-insensitive.
//
// Example CSV file:
//
//	FirstName,LastName,Email
//	John,Doe,john@example.com
//	Jane,Smith,jane@example.com
//
// The CSV reader allows variable number of fields per record. If a row has fewer
// columns than the header row, missing values are treated as empty strings.
type CSVDataSource struct {
	filePath   string
	file       *os.File
	reader     *csv.Reader
	headers    []string
	currentRow []string
}

// NewCSVDataSource creates a new CSV data source from the specified file path.
// The file is not opened until Open() is called.
//
// Example:
//
//	ds := mailmerge.NewCSVDataSource("contacts.csv")
//	defer ds.Close()
//	if err := ds.Open(); err != nil {
//		log.Fatal(err)
//	}
func NewCSVDataSource(
	filePath string,
) *CSVDataSource {
	return &CSVDataSource{
		filePath: filePath,
	}
}

// Open opens the CSV file and reads the header row.
func (c *CSVDataSource) Open() error {
	file, err := os.Open(c.filePath)
	if err != nil {
		return fmt.Errorf(
			"failed to open CSV file: %w",
			err,
		)
	}
	c.file = file
	c.reader = csv.NewReader(file)
	// Allow variable number of fields per record
	c.reader.FieldsPerRecord = -1

	// Read the header row
	headers, err := c.reader.Read()
	if err != nil {
		_ = c.file.Close()
		if err == io.EOF {
			return errors.New("CSV file is empty")
		}

		return fmt.Errorf(
			"failed to read CSV headers: %w",
			err,
		)
	}

	// Trim whitespace from headers
	c.headers = make([]string, len(headers))
	for i, h := range headers {
		c.headers[i] = strings.TrimSpace(h)
	}

	return nil
}

// Close closes the underlying file handle.
func (c *CSVDataSource) Close() error {
	if c.file != nil {
		return c.file.Close()
	}

	return nil
}

// Next reads the next row from the CSV file.
func (c *CSVDataSource) Next() bool {
	row, err := c.reader.Read()
	if err != nil {
		c.currentRow = nil

		return false
	}

	// Trim whitespace from values
	c.currentRow = make([]string, len(row))
	for i, v := range row {
		c.currentRow[i] = strings.TrimSpace(v)
	}

	return true
}

// Fields returns the header row as field names.
func (c *CSVDataSource) Fields() []string {
	return c.headers
}

// Get retrieves the value for the specified field name (case-insensitive).
func (c *CSVDataSource) Get(
	fieldName string,
) (string, error) {
	if c.currentRow == nil {
		return "", errors.New(
			"no current record (call Next() first)",
		)
	}

	// Case-insensitive lookup
	fieldNameLower := strings.ToLower(fieldName)
	for i, header := range c.headers {
		if strings.EqualFold(
			header,
			fieldNameLower,
		) {
			// Handle case where row has fewer columns than headers
			if i >= len(c.currentRow) {
				return "", nil
			}

			return c.currentRow[i], nil
		}
	}

	return "", fmt.Errorf(
		"field %q not found",
		fieldName,
	)
}

// MapDataSource is a DataSource implementation that reads from a slice of maps.
//
// Each map represents one record, with keys as field names and values as field values.
// This is useful for programmatic data generation or when you already have data in memory.
//
// Field names are collected from all records (union of all keys). Records may have
// different sets of fields - missing fields in a record are treated as empty strings.
// Field name matching is case-insensitive.
//
// Example:
//
//	records := []map[string]string{
//		{"FirstName": "John", "LastName": "Doe", "Email": "john@example.com"},
//		{"FirstName": "Jane", "LastName": "Smith", "Email": "jane@example.com"},
//	}
//	ds := mailmerge.NewMapDataSource(records)
type MapDataSource struct {
	records []map[string]string
	index   int
	fields  []string
}

// NewMapDataSource creates a new map data source from a slice of maps.
// The records parameter should contain one map per data record.
//
// Example:
//
//	records := []map[string]string{
//		{"Name": "Alice", "Age": "30"},
//		{"Name": "Bob", "Age": "25"},
//	}
//	ds := mailmerge.NewMapDataSource(records)
func NewMapDataSource(
	records []map[string]string,
) *MapDataSource {
	return &MapDataSource{
		records: records,
		index:   -1, // Start before first record
	}
}

// Open computes the unique set of field names across all records.
func (m *MapDataSource) Open() error {
	// Compute unique fields from all records
	fieldSet := make(map[string]bool)
	for _, record := range m.records {
		for key := range record {
			fieldSet[key] = true
		}
	}

	// Convert to slice
	m.fields = make([]string, 0, len(fieldSet))
	for field := range fieldSet {
		m.fields = append(m.fields, field)
	}

	return nil
}

// Close is a no-op for MapDataSource.
func (*MapDataSource) Close() error {
	return nil
}

// Next advances to the next record.
func (m *MapDataSource) Next() bool {
	m.index++

	return m.index < len(m.records)
}

// Fields returns the list of unique field names across all records.
func (m *MapDataSource) Fields() []string {
	return m.fields
}

// Get retrieves the value for the specified field name (case-insensitive).
func (m *MapDataSource) Get(
	fieldName string,
) (string, error) {
	if m.index < 0 || m.index >= len(m.records) {
		return "", errors.New(
			"no current record (call Next() first)",
		)
	}

	currentRecord := m.records[m.index]

	// Case-insensitive lookup
	fieldNameLower := strings.ToLower(fieldName)
	for key, value := range currentRecord {
		if strings.EqualFold(
			key,
			fieldNameLower,
		) {
			return value, nil
		}
	}

	// Field not found in this record (may exist in others)
	return "", nil
}

// JSONDataSource is a DataSource implementation that reads from a JSON file.
//
// The JSON file must contain an array of objects at the top level. Each object represents
// one data record, with object keys as field names and values as field values.
//
// Field names are collected from all records (union of all keys). Non-string values are
// converted to strings using fmt.Sprint. Null values are treated as empty strings.
// Field name matching is case-insensitive.
//
// Example JSON file:
//
//	[
//		{"FirstName": "John", "LastName": "Doe", "Age": 30},
//		{"FirstName": "Jane", "LastName": "Smith", "Age": 25}
//	]
//
// The Age field (numeric) will be automatically converted to strings "30" and "25".
type JSONDataSource struct {
	filePath string
	records  []map[string]any
	fields   []string
	index    int
}

// NewJSONDataSource creates a new JSON data source from the specified file path.
//
// The JSON file must contain an array of objects at the top level. The file is not
// read until Open() is called.
//
// Example:
//
//	ds := mailmerge.NewJSONDataSource("contacts.json")
//	defer ds.Close()
//	if err := ds.Open(); err != nil {
//		log.Fatal(err)
//	}
func NewJSONDataSource(
	filePath string,
) *JSONDataSource {
	return &JSONDataSource{
		filePath: filePath,
		index:    -1, // Start before first record
	}
}

// Open reads and parses the JSON file, extracting field names from all records.
func (j *JSONDataSource) Open() error {
	data, err := os.ReadFile(j.filePath)
	if err != nil {
		return fmt.Errorf(
			"failed to read JSON file: %w",
			err,
		)
	}

	var arr []map[string]any
	if err := json.Unmarshal(data, &arr); err != nil {
		return fmt.Errorf(
			"JSON must be an array of objects: %w",
			err,
		)
	}

	j.records = arr
	j.index = -1

	// Extract unique field names from all records
	fieldSet := make(map[string]bool)
	for _, record := range j.records {
		for key := range record {
			fieldSet[key] = true
		}
	}

	// Convert to sorted slice for consistent ordering
	j.fields = make([]string, 0, len(fieldSet))
	for key := range fieldSet {
		j.fields = append(j.fields, key)
	}
	sort.Strings(j.fields)

	return nil
}

// Close is a no-op for JSONDataSource.
func (*JSONDataSource) Close() error {
	return nil
}

// Next advances to the next record.
func (j *JSONDataSource) Next() bool {
	j.index++

	return j.index < len(j.records)
}

// Fields returns the list of unique field names across all records.
func (j *JSONDataSource) Fields() []string {
	return j.fields
}

// Get retrieves the value for the specified field name (case-insensitive).
// Non-string values are converted to strings using fmt.Sprint.
// Nil values are returned as empty strings.
func (j *JSONDataSource) Get(
	fieldName string,
) (string, error) {
	if j.index < 0 || j.index >= len(j.records) {
		return "", errors.New(
			"no current record (call Next() first)",
		)
	}

	currentRecord := j.records[j.index]

	// Case-insensitive lookup
	for key, value := range currentRecord {
		if strings.EqualFold(key, fieldName) {
			if value == nil {
				return "", nil
			}

			return fmt.Sprint(value), nil
		}
	}

	// Field not found in this record
	return "", fmt.Errorf(
		"field %q not found",
		fieldName,
	)
}

// ValidateDataSourceFields validates that all requested field names exist in the data source.
// It performs case-insensitive comparison of field names.
// Returns an error listing all missing fields if any are not found.
func ValidateDataSourceFields(
	ds DataSource,
	fieldNames []string,
) error {
	if err := ds.Open(); err != nil {
		return fmt.Errorf(
			"failed to open data source: %w",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	availableFields := ds.Fields()
	availableFieldsLower := make(map[string]bool)
	for _, field := range availableFields {
		availableFieldsLower[strings.ToLower(field)] = true
	}

	var missingFields []string
	for _, requestedField := range fieldNames {
		if !availableFieldsLower[strings.ToLower(requestedField)] {
			missingFields = append(
				missingFields,
				requestedField,
			)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf(
			"missing fields: %s",
			strings.Join(missingFields, ", "),
		)
	}

	return nil
}
