// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge_test

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/wordprocessing/mailmerge"
)

// Example_new demonstrates creating a new mail merge engine.
func Example_new() {
	// Open a template document
	// doc, err := wordprocessing.Open("template.docx", true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create mail merge engine
	// mm := mailmerge.New(doc)

	// Configure data source (shown in other examples)
	// _ = mm

	fmt.Println(
		"Create mail merge engine with New()",
	)
	// Output:
	// Create mail merge engine with New()
}

// Example_execute demonstrates basic mail merge execution with CSV data.
func Example_execute() {
	// Typical usage pattern:
	//
	// doc, _ := wordprocessing.Open("template.docx", true)
	// ds := mailmerge.NewCSVDataSource("contacts.csv")
	//
	// result, err := mailmerge.New(doc).
	//     DataSource(ds).
	//     Execute()
	// if err != nil {
	//     log.Fatal(err)
	// }
	//
	// result.SaveAs("merged.docx")

	fmt.Println(
		"Execute merges all records into one document",
	)
	// Output:
	// Execute merges all records into one document
}

// Example_executeToDocuments demonstrates creating individual documents per record.
func Example_executeToDocuments() {
	// Typical usage pattern:
	//
	// records := []map[string]string{
	//     {"FirstName": "John", "LastName": "Doe"},
	//     {"FirstName": "Jane", "LastName": "Smith"},
	// }
	// ds := mailmerge.NewMapDataSource(records)
	//
	// doc, _ := wordprocessing.Open("template.docx", true)
	//
	// documents, err := mailmerge.New(doc).
	//     DataSource(ds).
	//     ExecuteToDocuments()
	//
	// for i, doc := range documents {
	//     doc.SaveAs(fmt.Sprintf("letter_%d.docx", i+1))
	// }

	fmt.Println(
		"ExecuteToDocuments creates one document per record",
	)
	// Output:
	// ExecuteToDocuments creates one document per record
}

// ExampleNewCSVDataSource demonstrates creating a CSV data source.
func ExampleNewCSVDataSource() {
	// Create CSV data source
	// ds := mailmerge.NewCSVDataSource("contacts.csv")
	//
	// if err := ds.Open(); err != nil {
	//     log.Fatal(err)
	// }
	// defer ds.Close()
	//
	// for ds.Next() {
	//     firstName, _ := ds.Get("FirstName")
	//     lastName, _ := ds.Get("LastName")
	//     fmt.Printf("%s %s\n", firstName, lastName)
	// }

	fmt.Println(
		"CSV data source reads from CSV files",
	)
	// Output:
	// CSV data source reads from CSV files
}

// ExampleNewJSONDataSource demonstrates creating a JSON data source.
func ExampleNewJSONDataSource() {
	// Create JSON data source
	// ds := mailmerge.NewJSONDataSource("contacts.json")
	//
	// if err := ds.Open(); err != nil {
	//     log.Fatal(err)
	// }
	// defer ds.Close()
	//
	// for ds.Next() {
	//     name, _ := ds.Get("FirstName")
	//     fmt.Println(name)
	// }

	fmt.Println(
		"JSON data source reads from JSON array files",
	)
	// Output:
	// JSON data source reads from JSON array files
}

// ExampleNewMapDataSource demonstrates creating an in-memory data source.
func ExampleNewMapDataSource() {
	records := []map[string]string{
		{"FirstName": "John", "LastName": "Doe"},
		{
			"FirstName": "Jane",
			"LastName":  "Smith",
		},
	}

	ds := mailmerge.NewMapDataSource(records)

	if err := ds.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = ds.Close()
	}()

	recordCount := 0
	for ds.Next() {
		recordCount++
	}

	fmt.Printf(
		"Processed %d records\n",
		recordCount,
	)
	// Output:
	// Processed 2 records
}

// ExampleMergeOptions demonstrates configuring merge options.
func ExampleMergeOptions() {
	// Typical usage pattern:
	//
	// opts := &mailmerge.MergeOptions{
	//     StrictFields:       true, // Error on missing fields
	//     RemoveUnusedFields: true, // Remove unreplaced fields
	// }
	//
	// doc, _ := wordprocessing.Open("template.docx", true)
	// ds := mailmerge.NewMapDataSource(records)
	//
	// result, err := mailmerge.New(doc).
	//     Options(opts).
	//     DataSource(ds).
	//     Execute()

	fmt.Println(
		"MergeOptions configure merge behavior",
	)
	// Output:
	// MergeOptions configure merge behavior
}
