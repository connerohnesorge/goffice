// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	testValueJohn = "John"
	testValueUSA  = "USA"
)

func TestCSVDataSource_ValidCSV(t *testing.T) {
	// Create a temporary CSV file
	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "test.csv")

	csvContent := `Name,Age,City
John Doe,30,New York
Jane Smith,25,Los Angeles
Bob Johnson,35,Chicago`

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Check fields
	fields := ds.Fields()
	expectedFields := []string{
		"Name",
		"Age",
		"City",
	}
	if len(fields) != len(expectedFields) {
		t.Errorf(
			"Expected %d fields, got %d",
			len(expectedFields),
			len(fields),
		)
	}
	for i, expected := range expectedFields {
		if fields[i] != expected {
			t.Errorf(
				"Field %d: expected %q, got %q",
				i,
				expected,
				fields[i],
			)
		}
	}

	// Read first record
	if !ds.Next() {
		t.Fatal("Expected first record to exist")
	}

	name, err := ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "John Doe" {
		t.Errorf(
			"Expected Name=John Doe, got %q",
			name,
		)
	}

	age, err := ds.Get("Age")
	if err != nil {
		t.Errorf("Failed to get Age: %v", err)
	}
	if age != "30" {
		t.Errorf("Expected Age=30, got %q", age)
	}

	// Read second record
	if !ds.Next() {
		t.Fatal("Expected second record to exist")
	}

	name, err = ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "Jane Smith" {
		t.Errorf(
			"Expected Name=Jane Smith, got %q",
			name,
		)
	}

	// Read third record
	if !ds.Next() {
		t.Fatal("Expected third record to exist")
	}

	name, err = ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "Bob Johnson" {
		t.Errorf(
			"Expected Name=Bob Johnson, got %q",
			name,
		)
	}

	// No more records
	if ds.Next() {
		t.Error("Expected no more records")
	}
}

func TestCSVDataSource_HeadersOnly(t *testing.T) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(
		tmpDir,
		"headers_only.csv",
	)

	csvContent := `Name,Age,City`

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Fields should be present
	fields := ds.Fields()
	if len(fields) != 3 {
		t.Errorf(
			"Expected 3 fields, got %d",
			len(fields),
		)
	}

	// No records should exist
	if ds.Next() {
		t.Error(
			"Expected no records, but Next() returned true",
		)
	}
}

func TestCSVDataSource_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "empty.csv")

	if err := os.WriteFile(csvPath, []byte(""), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	err := ds.Open()
	if err == nil {
		_ = ds.Close()
		t.Fatal(
			"Expected error when opening empty CSV",
		)
	}
	if !strings.Contains(err.Error(), "empty") {
		t.Errorf(
			"Expected 'empty' error, got: %v",
			err,
		)
	}
}

func TestCSVDataSource_CaseInsensitive(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(
		tmpDir,
		"case_test.csv",
	)

	csvContent := "Name,Age,City\n" + testValueJohn + ",30,NYC"

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	// Test various case combinations
	testCases := []string{
		"name",
		"NAME",
		"nAmE",
		"Age",
		"age",
		"AGE",
		"city",
		"CITY",
	}
	for _, fieldName := range testCases {
		_, err := ds.Get(fieldName)
		if err != nil {
			t.Errorf(
				"Failed to get field %q (case-insensitive): %v",
				fieldName,
				err,
			)
		}
	}
}

func TestCSVDataSource_MissingColumns(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(
		tmpDir,
		"missing_cols.csv",
	)

	// Second row has fewer columns than headers
	csvContent := `Name,Age,City
John,30
Jane,25,LA`

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// First record has missing City
	if !ds.Next() {
		t.Fatal("Expected first record to exist")
	}

	city, err := ds.Get("City")
	if err != nil {
		t.Errorf(
			"Expected no error for missing column, got: %v",
			err,
		)
	}
	if city != "" {
		t.Errorf(
			"Expected empty string for missing column, got %q",
			city,
		)
	}

	// Second record has City
	if !ds.Next() {
		t.Fatal("Expected second record to exist")
	}

	city, err = ds.Get("City")
	if err != nil {
		t.Errorf("Failed to get City: %v", err)
	}
	if city != "LA" {
		t.Errorf("Expected City=LA, got %q", city)
	}
}

func TestCSVDataSource_WhitespaceTrimming(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(
		tmpDir,
		"whitespace.csv",
	)

	// CSV with extra whitespace
	csvContent := `  Name  ,  Age  ,  City
  ` + testValueJohn + `  ,  30  ,  NYC  `

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Check fields are trimmed
	fields := ds.Fields()
	if fields[0] != "Name" ||
		fields[1] != "Age" ||
		fields[2] != "City" {
		t.Errorf(
			"Fields not trimmed correctly: %v",
			fields,
		)
	}

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	// Check values are trimmed
	name, _ := ds.Get("Name")
	if name != testValueJohn {
		t.Errorf(
			"Expected trimmed value '%s', got %q",
			testValueJohn,
			name,
		)
	}

	age, _ := ds.Get("Age")
	if age != "30" {
		t.Errorf(
			"Expected trimmed value '30', got %q",
			age,
		)
	}
}

func TestCSVDataSource_NonExistentField(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "test.csv")

	csvContent := `Name,Age
John,30`

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	_, err := ds.Get("NonExistent")
	if err == nil {
		t.Error(
			"Expected error for non-existent field",
		)
	}
	if !strings.Contains(
		err.Error(),
		"not found",
	) {
		t.Errorf(
			"Expected 'not found' error, got: %v",
			err,
		)
	}
}

func TestCSVDataSource_GetBeforeNext(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "test.csv")

	csvContent := `Name,Age
John,30`

	if err := os.WriteFile(csvPath, []byte(csvContent), 0o644); err != nil {
		t.Fatalf(
			"Failed to create test CSV: %v",
			err,
		)
	}

	ds := NewCSVDataSource(csvPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open CSV: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Try to get without calling Next()
	_, err := ds.Get("Name")
	if err == nil {
		t.Error(
			"Expected error when calling Get() before Next()",
		)
	}
	if !strings.Contains(
		err.Error(),
		"no current record",
	) {
		t.Errorf(
			"Expected 'no current record' error, got: %v",
			err,
		)
	}
}

func TestMapDataSource_MultipleRecords(
	t *testing.T,
) {
	records := []map[string]string{
		{
			"Name": "John",
			"Age":  "30",
			"City": "NYC",
		},
		{
			"Name": "Jane",
			"Age":  "25",
			"City": "LA",
		},
		{
			"Name": "Bob",
			"Age":  "35",
			"City": "Chicago",
		},
	}

	ds := NewMapDataSource(records)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open MapDataSource: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Check fields
	fields := ds.Fields()
	if len(fields) != 3 {
		t.Errorf(
			"Expected 3 fields, got %d",
			len(fields),
		)
	}

	// Verify all expected fields are present
	fieldSet := make(map[string]bool)
	for _, f := range fields {
		fieldSet[f] = true
	}
	for _, expectedField := range []string{"Name", "Age", "City"} {
		if !fieldSet[expectedField] {
			t.Errorf(
				"Expected field %q not found",
				expectedField,
			)
		}
	}

	// Read first record
	if !ds.Next() {
		t.Fatal("Expected first record to exist")
	}

	name, err := ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != testValueJohn {
		t.Errorf(
			"Expected Name=%s, got %q",
			testValueJohn,
			name,
		)
	}

	// Read second record
	if !ds.Next() {
		t.Fatal("Expected second record to exist")
	}

	name, err = ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "Jane" {
		t.Errorf(
			"Expected Name=Jane, got %q",
			name,
		)
	}

	// Read third record
	if !ds.Next() {
		t.Fatal("Expected third record to exist")
	}

	name, err = ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "Bob" {
		t.Errorf(
			"Expected Name=Bob, got %q",
			name,
		)
	}

	// No more records
	if ds.Next() {
		t.Error("Expected no more records")
	}
}

func TestMapDataSource_EmptySlice(t *testing.T) {
	var records []map[string]string

	ds := NewMapDataSource(records)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open MapDataSource: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Fields should be empty
	fields := ds.Fields()
	if len(fields) != 0 {
		t.Errorf(
			"Expected 0 fields for empty records, got %d",
			len(fields),
		)
	}

	// No records should exist
	if ds.Next() {
		t.Error(
			"Expected no records, but Next() returned true",
		)
	}
}

func TestMapDataSource_CaseInsensitive(
	t *testing.T,
) {
	records := []map[string]string{
		{"Name": "John", "Age": "30"},
	}

	ds := NewMapDataSource(records)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open MapDataSource: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	// Test various case combinations
	testCases := []string{
		"name",
		"NAME",
		"nAmE",
		"Age",
		"age",
		"AGE",
	}
	for _, fieldName := range testCases {
		_, err := ds.Get(fieldName)
		if err != nil {
			t.Errorf(
				"Failed to get field %q (case-insensitive): %v",
				fieldName,
				err,
			)
		}
	}
}

func TestMapDataSource_HeterogeneousKeys(
	t *testing.T,
) {
	// Records with different sets of keys
	records := []map[string]string{
		{"Name": "John", "Age": "30"},
		{"Name": "Jane", "City": "LA"},
		{
			"Age":     "35",
			"City":    "Chicago",
			"Country": "USA",
		},
	}

	ds := NewMapDataSource(records)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open MapDataSource: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Fields should include all unique keys
	fields := ds.Fields()
	expectedFields := []string{
		"Name",
		"Age",
		"City",
		"Country",
	}
	if len(fields) != len(expectedFields) {
		t.Errorf(
			"Expected %d fields, got %d",
			len(expectedFields),
			len(fields),
		)
	}

	// Verify all fields are present
	fieldSet := make(map[string]bool)
	for _, f := range fields {
		fieldSet[f] = true
	}
	for _, expectedField := range expectedFields {
		if !fieldSet[expectedField] {
			t.Errorf(
				"Expected field %q not found in %v",
				expectedField,
				fields,
			)
		}
	}

	// First record: has Name and Age, missing City and Country
	if !ds.Next() {
		t.Fatal("Expected first record to exist")
	}

	name, err := ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != testValueJohn {
		t.Errorf(
			"Expected Name=%s, got %q",
			testValueJohn,
			name,
		)
	}

	city, err := ds.Get("City")
	if err != nil {
		t.Errorf(
			"Expected no error for missing field, got: %v",
			err,
		)
	}
	if city != "" {
		t.Errorf(
			"Expected empty string for missing field, got %q",
			city,
		)
	}

	// Second record: has Name and City, missing Age and Country
	if !ds.Next() {
		t.Fatal("Expected second record to exist")
	}

	age, err := ds.Get("Age")
	if err != nil {
		t.Errorf(
			"Expected no error for missing field, got: %v",
			err,
		)
	}
	if age != "" {
		t.Errorf(
			"Expected empty string for missing field, got %q",
			age,
		)
	}

	// Third record: missing Name
	if !ds.Next() {
		t.Fatal("Expected third record to exist")
	}

	country, err := ds.Get("Country")
	if err != nil {
		t.Errorf("Failed to get Country: %v", err)
	}
	if country != testValueUSA {
		t.Errorf(
			"Expected Country=%s, got %q",
			testValueUSA,
			country,
		)
	}
}

func TestMapDataSource_GetBeforeNext(
	t *testing.T,
) {
	records := []map[string]string{
		{"Name": "John", "Age": "30"},
	}

	ds := NewMapDataSource(records)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open MapDataSource: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Try to get without calling Next()
	_, err := ds.Get("Name")
	if err == nil {
		t.Error(
			"Expected error when calling Get() before Next()",
		)
	}
	if !strings.Contains(
		err.Error(),
		"no current record",
	) {
		t.Errorf(
			"Expected 'no current record' error, got: %v",
			err,
		)
	}
}

func TestMapDataSource_GetAfterEnd(t *testing.T) {
	records := []map[string]string{
		{"Name": "John"},
	}

	ds := NewMapDataSource(records)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open MapDataSource: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Read the only record
	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	// Move past the end
	if ds.Next() {
		t.Error("Expected no more records")
	}

	// Try to get after end
	_, err := ds.Get("Name")
	if err == nil {
		t.Error(
			"Expected error when calling Get() after end",
		)
	}
	if !strings.Contains(
		err.Error(),
		"no current record",
	) {
		t.Errorf(
			"Expected 'no current record' error, got: %v",
			err,
		)
	}
}

// Helper function to create test JSON files
func createTestJSONFile(
	t *testing.T,
	data string,
) string {
	t.Helper()
	tmpFile, err := os.CreateTemp(
		"",
		"test*.json",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp file: %v",
			err,
		)
	}
	_, err = tmpFile.WriteString(data)
	if err != nil {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
		t.Fatalf(
			"Failed to write to temp file: %v",
			err,
		)
	}
	_ = tmpFile.Close()
	t.Cleanup(
		func() { _ = os.Remove(tmpFile.Name()) },
	)

	return tmpFile.Name()
}

func TestJSONDataSource_ValidJSON(t *testing.T) {
	jsonContent := `[
		{"Name": "John Doe", "Age": "30", "City": "New York"},
		{"Name": "Jane Smith", "Age": "25", "City": "Los Angeles"},
		{"Name": "Bob Johnson", "Age": "35", "City": "Chicago"}
	]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Check fields
	fields := ds.Fields()
	if len(fields) != 3 {
		t.Errorf(
			"Expected 3 fields, got %d",
			len(fields),
		)
	}

	// Verify all expected fields are present (sorted)
	expectedFields := []string{
		"Age",
		"City",
		"Name",
	}
	for i, expected := range expectedFields {
		if fields[i] != expected {
			t.Errorf(
				"Field %d: expected %q, got %q",
				i,
				expected,
				fields[i],
			)
		}
	}

	// Read first record
	if !ds.Next() {
		t.Fatal("Expected first record to exist")
	}

	name, err := ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "John Doe" {
		t.Errorf(
			"Expected Name=John Doe, got %q",
			name,
		)
	}

	age, err := ds.Get("Age")
	if err != nil {
		t.Errorf("Failed to get Age: %v", err)
	}
	if age != "30" {
		t.Errorf("Expected Age=30, got %q", age)
	}

	// Read second record
	if !ds.Next() {
		t.Fatal("Expected second record to exist")
	}

	name, err = ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "Jane Smith" {
		t.Errorf(
			"Expected Name=Jane Smith, got %q",
			name,
		)
	}

	// Read third record
	if !ds.Next() {
		t.Fatal("Expected third record to exist")
	}

	name, err = ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != "Bob Johnson" {
		t.Errorf(
			"Expected Name=Bob Johnson, got %q",
			name,
		)
	}

	// No more records
	if ds.Next() {
		t.Error("Expected no more records")
	}
}

func TestJSONDataSource_MatchingKeys(
	t *testing.T,
) {
	jsonContent := `[
		{"Name": "John", "Age": "30", "City": "NYC"},
		{"Name": "Jane", "Age": "25", "City": "LA"}
	]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// All records have the same keys
	fields := ds.Fields()
	if len(fields) != 3 {
		t.Errorf(
			"Expected 3 fields, got %d",
			len(fields),
		)
	}

	// Verify records can be read
	count := 0
	for ds.Next() {
		count++
		name, err := ds.Get("Name")
		if err != nil {
			t.Errorf(
				"Failed to get Name in record %d: %v",
				count,
				err,
			)
		}
		if name == "" {
			t.Errorf(
				"Expected non-empty Name in record %d",
				count,
			)
		}
	}

	if count != 2 {
		t.Errorf(
			"Expected 2 records, got %d",
			count,
		)
	}
}

func TestJSONDataSource_HeterogeneousKeys(
	t *testing.T,
) {
	jsonContent := `[
		{"Name": "John", "Age": "30"},
		{"Name": "Jane", "City": "LA"},
		{"Age": "35", "City": "Chicago", "Country": "USA"}
	]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Fields should include all unique keys
	fields := ds.Fields()
	expectedFields := []string{
		"Age",
		"City",
		"Country",
		"Name",
	}
	if len(fields) != len(expectedFields) {
		t.Errorf(
			"Expected %d fields, got %d",
			len(expectedFields),
			len(fields),
		)
	}

	// Verify all fields are present (sorted)
	for i, expected := range expectedFields {
		if fields[i] != expected {
			t.Errorf(
				"Field %d: expected %q, got %q",
				i,
				expected,
				fields[i],
			)
		}
	}

	// First record: has Name and Age, missing City and Country
	if !ds.Next() {
		t.Fatal("Expected first record to exist")
	}

	name, err := ds.Get("Name")
	if err != nil {
		t.Errorf("Failed to get Name: %v", err)
	}
	if name != testValueJohn {
		t.Errorf(
			"Expected Name=%s, got %q",
			testValueJohn,
			name,
		)
	}

	// Missing field should return error
	_, err = ds.Get("City")
	if err == nil {
		t.Error(
			"Expected error for missing field City",
		)
	}

	// Second record: has Name and City, missing Age and Country
	if !ds.Next() {
		t.Fatal("Expected second record to exist")
	}

	city, err := ds.Get("City")
	if err != nil {
		t.Errorf("Failed to get City: %v", err)
	}
	if city != "LA" {
		t.Errorf("Expected City=LA, got %q", city)
	}

	// Third record: has Age, City, Country but missing Name
	if !ds.Next() {
		t.Fatal("Expected third record to exist")
	}

	country, err := ds.Get("Country")
	if err != nil {
		t.Errorf("Failed to get Country: %v", err)
	}
	if country != testValueUSA {
		t.Errorf(
			"Expected Country=%s, got %q",
			testValueUSA,
			country,
		)
	}
}

func TestJSONDataSource_EmptyArray(t *testing.T) {
	jsonContent := `[]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf(
			"Failed to open JSON with empty array: %v",
			err,
		)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Fields should be empty
	fields := ds.Fields()
	if len(fields) != 0 {
		t.Errorf(
			"Expected 0 fields for empty array, got %d",
			len(fields),
		)
	}

	// Next should return false immediately
	if ds.Next() {
		t.Error(
			"Expected Next() to return false for empty array",
		)
	}
}

func TestJSONDataSource_NonArrayJSON(
	t *testing.T,
) {
	jsonContent := `{"Name": "John", "Age": "30"}`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	err := ds.Open()
	if err == nil {
		_ = ds.Close()
		t.Fatal(
			"Expected error when opening non-array JSON",
		)
	}
	if !strings.Contains(
		err.Error(),
		"array of objects",
	) {
		t.Errorf(
			"Expected 'array of objects' error, got: %v",
			err,
		)
	}
}

func TestJSONDataSource_CaseInsensitive(
	t *testing.T,
) {
	jsonContent := `[{"Name": "John", "Age": "30"}]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	// Test various case combinations
	testCases := []string{
		"name",
		"NAME",
		"nAmE",
		"Age",
		"age",
		"AGE",
	}
	for _, fieldName := range testCases {
		_, err := ds.Get(fieldName)
		if err != nil {
			t.Errorf(
				"Failed to get field %q (case-insensitive): %v",
				fieldName,
				err,
			)
		}
	}
}

func TestJSONDataSource_NonStringValues(
	t *testing.T,
) {
	jsonContent := `[
		{"Name": "John", "Age": 30, "Active": true, "Score": 95.5, "Notes": null}
	]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	// Number should be converted to string
	age, err := ds.Get("Age")
	if err != nil {
		t.Errorf("Failed to get Age: %v", err)
	}
	if age != "30" {
		t.Errorf(
			"Expected Age=30 (string), got %q",
			age,
		)
	}

	// Boolean should be converted to string
	active, err := ds.Get("Active")
	if err != nil {
		t.Errorf("Failed to get Active: %v", err)
	}
	if active != "true" {
		t.Errorf(
			"Expected Active=true (string), got %q",
			active,
		)
	}

	// Float should be converted to string
	score, err := ds.Get("Score")
	if err != nil {
		t.Errorf("Failed to get Score: %v", err)
	}
	if score != "95.5" {
		t.Errorf(
			"Expected Score=95.5 (string), got %q",
			score,
		)
	}

	// Null should be converted to empty string
	notes, err := ds.Get("Notes")
	if err != nil {
		t.Errorf("Failed to get Notes: %v", err)
	}
	if notes != "" {
		t.Errorf(
			"Expected Notes='' (empty string), got %q",
			notes,
		)
	}
}

func TestJSONDataSource_MissingField(
	t *testing.T,
) {
	jsonContent := `[{"Name": "John", "Age": "30"}]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	if !ds.Next() {
		t.Fatal("Expected record to exist")
	}

	_, err := ds.Get("NonExistent")
	if err == nil {
		t.Error(
			"Expected error for non-existent field",
		)
	}
	if !strings.Contains(
		err.Error(),
		"not found",
	) {
		t.Errorf(
			"Expected 'not found' error, got: %v",
			err,
		)
	}
}

func TestJSONDataSource_GetBeforeNext(
	t *testing.T,
) {
	jsonContent := `[{"Name": "John"}]`

	jsonPath := createTestJSONFile(t, jsonContent)

	ds := NewJSONDataSource(jsonPath)
	if err := ds.Open(); err != nil {
		t.Fatalf("Failed to open JSON: %v", err)
	}
	defer func() {
		_ = ds.Close()
	}()

	// Try to get without calling Next()
	_, err := ds.Get("Name")
	if err == nil {
		t.Error(
			"Expected error when calling Get() before Next()",
		)
	}
	if !strings.Contains(
		err.Error(),
		"no current record",
	) {
		t.Errorf(
			"Expected 'no current record' error, got: %v",
			err,
		)
	}
}

func TestValidateDataSourceFields_AllFieldsPresent(
	t *testing.T,
) {
	records := []map[string]string{
		{
			"Name": "John",
			"Age":  "30",
			"City": "NYC",
		},
	}

	ds := NewMapDataSource(records)
	err := ValidateDataSourceFields(
		ds,
		[]string{"Name", "Age", "City"},
	)
	if err != nil {
		t.Errorf(
			"Expected no error when all fields present, got: %v",
			err,
		)
	}
}

func TestValidateDataSourceFields_MissingField(
	t *testing.T,
) {
	records := []map[string]string{
		{"Name": "John", "Age": "30"},
	}

	ds := NewMapDataSource(records)
	err := ValidateDataSourceFields(
		ds,
		[]string{"Name", "City"},
	)
	if err == nil {
		t.Fatal(
			"Expected error for missing field",
		)
	}
	if !strings.Contains(
		err.Error(),
		"missing fields",
	) {
		t.Errorf(
			"Expected 'missing fields' error, got: %v",
			err,
		)
	}
	if !strings.Contains(err.Error(), "City") {
		t.Errorf(
			"Expected error to mention 'City', got: %v",
			err,
		)
	}
}

func TestValidateDataSourceFields_CaseInsensitive(
	t *testing.T,
) {
	records := []map[string]string{
		{"Name": "John", "Age": "30"},
	}

	ds := NewMapDataSource(records)
	// Request fields with different case
	err := ValidateDataSourceFields(
		ds,
		[]string{"name", "AGE"},
	)
	if err != nil {
		t.Errorf(
			"Expected no error with case variations, got: %v",
			err,
		)
	}
}

func TestValidateDataSourceFields_MultipleMissing(
	t *testing.T,
) {
	records := []map[string]string{
		{"Name": "John"},
	}

	ds := NewMapDataSource(records)
	err := ValidateDataSourceFields(
		ds,
		[]string{
			"Name",
			"Age",
			"City",
			"Country",
		},
	)
	if err == nil {
		t.Fatal(
			"Expected error for multiple missing fields",
		)
	}
	if !strings.Contains(
		err.Error(),
		"missing fields",
	) {
		t.Errorf(
			"Expected 'missing fields' error, got: %v",
			err,
		)
	}
	// Should list all missing fields
	if !strings.Contains(err.Error(), "Age") {
		t.Errorf(
			"Expected error to mention 'Age', got: %v",
			err,
		)
	}
	if !strings.Contains(err.Error(), "City") {
		t.Errorf(
			"Expected error to mention 'City', got: %v",
			err,
		)
	}
	if !strings.Contains(err.Error(), "Country") {
		t.Errorf(
			"Expected error to mention 'Country', got: %v",
			err,
		)
	}
}
