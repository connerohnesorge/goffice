// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"errors"
	"fmt"
	"strings"

	"github.com/connerohnesorge/goffice/wordprocessing"
)

// MailMerge is the main engine for performing mail merge operations on Word documents.
// It combines a template document with a data source to produce merged output.
type MailMerge struct {
	// template is the Word document template containing merge fields
	template *wordprocessing.Document

	// dataSource provides the data records for merging
	dataSource DataSource

	// options controls merge behavior
	options *MergeOptions

	// metadata contains information about the merge fields found in the template
	// (used in later phases)
	metadata *MailMergeMetadata

	// ooxmlMetadata contains OOXML-specific metadata loaded from document settings
	ooxmlMetadata *OoxmlMailMergeMetadata
}

// MergeOptions configures the behavior of the mail merge operation.
type MergeOptions struct {
	// StrictFields determines whether to error on missing fields in the data source.
	// If true, the merge will fail if a MERGEFIELD references a field not present
	// in the data source. If false, missing fields will be replaced with empty strings.
	StrictFields bool

	// RemoveUnusedFields determines whether to remove merge fields that were not replaced.
	// If true, any MERGEFIELD that couldn't be replaced (e.g., missing from data source
	// when StrictFields is false) will be removed from the output. If false, the field
	// code will remain in the document.
	RemoveUnusedFields bool
}

// MailMergeMetadata contains metadata about merge fields discovered in the template.
// This will be populated during field finding and used during merge execution.
type MailMergeMetadata struct {
	// Fields is the list of all merge fields found in the template
	Fields []*MergeField

	// FieldNames is the set of unique field names referenced in the template
	FieldNames []string
}

// New creates a new MailMerge engine with the specified template document.
// The template should contain MERGEFIELD or GREETINGLINE fields.
//
// Example:
//
//	doc, _ := wordprocessing.Open("template.docx", true)
//	mm := mailmerge.New(doc)
func New(
	template *wordprocessing.Document,
) *MailMerge {
	return &MailMerge{
		template: template,
		options:  DefaultMergeOptions(),
		metadata: &MailMergeMetadata{
			Fields:     make([]*MergeField, 0),
			FieldNames: make([]string, 0),
		},
	}
}

// DataSource sets the data source for the merge operation.
// Returns the MailMerge instance for method chaining.
//
// Example:
//
//	ds := mailmerge.NewCSVDataSource("data.csv")
//	mm := mailmerge.New(doc).DataSource(ds)
func (m *MailMerge) DataSource(
	ds DataSource,
) *MailMerge {
	m.dataSource = ds

	return m
}

// Options sets the merge options for the merge operation.
// Returns the MailMerge instance for method chaining.
//
// Example:
//
//	opts := &mailmerge.MergeOptions{
//	    StrictFields: true,
//	    RemoveUnusedFields: true,
//	}
//	mm := mailmerge.New(doc).Options(opts)
func (m *MailMerge) Options(
	opts *MergeOptions,
) *MailMerge {
	if opts != nil {
		m.options = opts
	}

	return m
}

// DefaultMergeOptions returns the default merge options.
// By default:
//   - StrictFields is false (missing fields won't cause errors)
//   - RemoveUnusedFields is true (unreplaced fields are removed)
func DefaultMergeOptions() *MergeOptions {
	return &MergeOptions{
		StrictFields:       false,
		RemoveUnusedFields: true,
	}
}

// Execute performs the mail merge operation, combining the template with the data source.
// Returns a new Document containing the merged output.
//
// The merge process:
//  1. Validates that a data source is configured
//  2. Opens the data source and finds all merge fields
//  3. Validates fields against data source (if StrictFields is enabled)
//  4. Creates output document structure
//  5. For each record in the data source:
//     - Clones the template body content
//     - Replaces merge fields with record values
//     - Appends to output (with page breaks between records)
//  6. Returns the merged document
//
// Example:
//
//	mm := mailmerge.New(template)
//	mm.DataSource(mailmerge.NewCSVDataSource("data.csv"))
//	merged, err := mm.Execute()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	merged.Save()
//
//nolint:revive // function-length: merge execution requires multi-step workflow
func (m *MailMerge) Execute() (*wordprocessing.Document, error) {
	// Step 1: Load OOXML metadata from document settings
	ooxmlMeta, err := loadMailMergeMetadata(
		m.template,
	)
	if err != nil {
		return nil, err
	}
	m.ooxmlMetadata = ooxmlMeta

	// Step 2: Validate data source is set
	if m.dataSource == nil {
		return nil, ErrNoDataSource
	}

	// Step 3: Open data source
	if err := m.dataSource.Open(); err != nil {
		return nil, err
	}
	defer func() {
		_ = m.dataSource.Close()
	}()

	// Find all merge fields in template
	fields := m.findMergeFields()
	m.metadata.Fields = fields

	// Populate unique field names
	uniqueNames := make(map[string]bool)
	for _, field := range fields {
		uniqueNames[field.FieldName] = true
	}
	m.metadata.FieldNames = make(
		[]string,
		0,
		len(uniqueNames),
	)
	for name := range uniqueNames {
		m.metadata.FieldNames = append(
			m.metadata.FieldNames,
			name,
		)
	}

	// Step 3: Validate fields against data source
	if err := m.validateFields(fields, m.dataSource); err != nil {
		return nil, err
	}

	// Get template body
	mainPart := m.template.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	doc := mainPart.Document()
	if doc == nil {
		return nil, ErrNoDocument
	}

	body := doc.Body()
	if body == nil {
		return nil, ErrNoBody
	}

	// Count records first
	recordCount := 0
	for m.dataSource.Next() {
		recordCount++
	}

	// Step 4: Handle zero records case
	if recordCount == 0 {
		return nil, ErrNoRecords
	}

	// Re-open data source to reset iterator
	_ = m.dataSource.Close()
	if err := m.dataSource.Open(); err != nil {
		return nil, err
	}

	// Step 5: Create output by processing each record
	// For first pass implementation, only process first record
	// TODO: Implement proper body cloning in later phases for multiple records
	if m.dataSource.Next() {
		// Find fields for this record
		currentFields := m.findMergeFields()

		// Replace fields with current record values
		// Apply field mapping from OOXML metadata
		getValue := func(fieldName string) (string, error) {
			// Apply field mapping from OOXML metadata
			mappedFieldName := applyFieldMapping(
				fieldName,
				m.ooxmlMetadata,
			)

			return m.dataSource.Get(
				mappedFieldName,
			)
		}

		if err := m.replaceFields(currentFields, getValue); err != nil {
			return nil, err
		}
	}

	// Return the modified template document
	return m.template, nil
}

// ExecuteToDocuments performs the mail merge operation and returns a separate document for each record.
// This is useful when you need individual documents rather than a single merged document.
//
// The merge process:
//  1. Validates that a data source is configured
//  2. Opens the data source and finds all merge fields
//  3. Validates fields against data source (if StrictFields is enabled)
//  4. For each record in the data source:
//     - Checks SKIPIF conditions (skip if condition is true)
//     - Clones the template document
//     - Replaces merge fields with record values
//     - Adds to output slice
//  5. Returns slice of merged documents
//
// Example:
//
//	mm := mailmerge.New(template)
//	mm.DataSource(mailmerge.NewCSVDataSource("data.csv"))
//	documents, err := mm.ExecuteToDocuments()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for i, doc := range documents {
//	    doc.SaveAs(fmt.Sprintf("output_%d.docx", i))
//	}
//
//nolint:revive // function-length: merge execution requires multi-step workflow
func (m *MailMerge) ExecuteToDocuments() ([]*wordprocessing.Document, error) {
	// Step 1: Load OOXML metadata from document settings
	ooxmlMeta, err := loadMailMergeMetadata(
		m.template,
	)
	if err != nil {
		return nil, err
	}
	m.ooxmlMetadata = ooxmlMeta

	// Step 2: Validate data source is set
	if m.dataSource == nil {
		return nil, ErrNoDataSource
	}

	// Step 3: Open data source
	if err := m.dataSource.Open(); err != nil {
		return nil, err
	}
	defer func() {
		_ = m.dataSource.Close()
	}()

	// Find all merge fields in template
	fields := m.findMergeFields()
	m.metadata.Fields = fields

	// Populate unique field names
	uniqueNames := make(map[string]bool)
	for _, field := range fields {
		uniqueNames[field.FieldName] = true
	}
	m.metadata.FieldNames = make(
		[]string,
		0,
		len(uniqueNames),
	)
	for name := range uniqueNames {
		m.metadata.FieldNames = append(
			m.metadata.FieldNames,
			name,
		)
	}

	// Step 4: Validate fields against data source
	if err := m.validateFields(fields, m.dataSource); err != nil {
		return nil, err
	}

	// Get template structure for cloning
	mainPart := m.template.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	doc := mainPart.Document()
	if doc == nil {
		return nil, ErrNoDocument
	}

	body := doc.Body()
	if body == nil {
		return nil, ErrNoBody
	}

	// Step 5: Process each record
	var documents []*wordprocessing.Document
	recordIndex := 0

	for m.dataSource.Next() {
		recordIndex++

		// Check SKIPIF conditions
		shouldSkip, err := m.shouldSkipRecord(
			fields,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"error evaluating SKIPIF for record %d: %w",
				recordIndex,
				err,
			)
		}
		if shouldSkip {
			continue
		}

		// Clone the template document
		clonedDoc, err := m.cloneTemplateDocument()
		if err != nil {
			return nil, fmt.Errorf(
				"error cloning template for record %d: %w",
				recordIndex,
				err,
			)
		}

		// Find fields in cloned document
		clonedMM := &MailMerge{
			template:      clonedDoc,
			dataSource:    m.dataSource,
			options:       m.options,
			ooxmlMetadata: m.ooxmlMetadata,
		}
		clonedFields := clonedMM.findMergeFields()

		// Replace fields with current record values
		getValue := func(fieldName string) (string, error) {
			// Apply field mapping from OOXML metadata
			mappedFieldName := applyFieldMapping(
				fieldName,
				m.ooxmlMetadata,
			)

			return m.dataSource.Get(
				mappedFieldName,
			)
		}

		if err := clonedMM.replaceFields(clonedFields, getValue); err != nil {
			return nil, fmt.Errorf(
				"error replacing fields for record %d: %w",
				recordIndex,
				err,
			)
		}

		documents = append(documents, clonedDoc)
	}

	// Step 6: Check if we got any documents
	if len(documents) == 0 {
		return nil, ErrNoRecords
	}

	return documents, nil
}

// cloneTemplateDocument creates a deep copy of the template document.
// This is done by saving the template to a bytes buffer and loading it back.
func (m *MailMerge) cloneTemplateDocument() (*wordprocessing.Document, error) {
	// Save template to bytes buffer
	var buf strings.Builder
	if err := m.template.SaveTo(&buf); err != nil {
		return nil, fmt.Errorf(
			"failed to save template: %w",
			err,
		)
	}

	// Create a new document from the bytes
	reader := strings.NewReader(buf.String())
	clonedDoc, err := wordprocessing.OpenReader(
		reader,
		int64(buf.Len()),
		true,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to open cloned document: %w",
			err,
		)
	}

	return clonedDoc, nil
}

// shouldSkipRecord checks if the current record should be skipped based on SKIPIF fields.
// Returns true if any SKIPIF condition evaluates to true.
func (m *MailMerge) shouldSkipRecord(
	fields []*MergeField,
) (bool, error) {
	for _, field := range fields {
		if field.FieldName != FieldNameSkipIf {
			continue
		}

		skip, err := evaluateSkipIf(
			m.dataSource,
			field.FieldCode,
		)
		if err != nil {
			return false, err
		}
		if skip {
			return true, nil
		}
	}

	return false, nil
}

// Error variables for Execute
var (
	ErrNoDataSource = errors.New(
		"no data source configured",
	)
	ErrNoMainPart = errors.New(
		"template has no main document part",
	)
	ErrNoDocument = errors.New(
		"template has no document element",
	)
	ErrNoBody = errors.New(
		"template has no body element",
	)
	ErrNoRecords = errors.New(
		"data source has no records",
	)
)
