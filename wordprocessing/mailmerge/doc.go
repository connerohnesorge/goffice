//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package mailmerge provides mail merge functionality for Word documents.
//
// Mail merge allows you to create personalized documents by combining a template
// document with data from various sources (CSV files, JSON files, or in-memory maps).
//
// # Basic Usage
//
//      doc, _ := wordprocessing.Open("template.docx", true)
//      result, err := mailmerge.New(doc).
//              DataSource(mailmerge.NewCSVDataSource("contacts.csv")).
//              Execute()
//      if err != nil {
//              log.Fatal(err)
//      }
//      result.Save()
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
//      opts := &mailmerge.MergeOptions{
//              StrictFields:       true,  // error on missing fields
//              RemoveUnusedFields: true,  // remove unreplaced fields
//      }
//      result, _ := mailmerge.New(doc).Options(opts).DataSource(ds).Execute()
//
// # OOXML Metadata Integration
//
// The package automatically reads mail merge metadata from the document's settings.xml
// (w:mailMerge element) to support field name mappings and data source configurations
// created by Microsoft Word. This ensures compatibility with templates created in Word.
package mailmerge
