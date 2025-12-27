// Package main provides an example of using ExtendedProperties in Word
// documents. This demonstrates how to set and read application-specific
// metadata like application name, company, manager, and document statistics.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/wordprocessing"
)

// This example demonstrates how to use ExtendedProperties in Word documents.
// ExtendedProperties provide access to application-specific metadata like
// application name, company, manager, document statistics, and more.
func main() {
	if err := run(); err != nil {
		log.Fatalf("Example failed: %v", err)
	}
}

func run() error {
	// Create and configure document
	doc, err := createDocumentWithExtendedProperties()
	if err != nil {
		return err
	}
	defer func() { _ = doc.Close() }()

	// Display the properties we just set
	displayExtendedProperties(
		doc.ExtendedProperties(),
	)

	// Clean up the example file
	defer func() { _ = os.Remove("extended_props_example.docx") }()

	// Read properties from the saved document
	if err := readExtendedPropertiesExample(); err != nil {
		return err
	}

	fmt.Println(
		"\nExample completed successfully!",
	)

	return nil
}

func createDocumentWithExtendedProperties() (
	*wordprocessing.Document,
	error,
) {
	// Create a new Word document
	doc, err := wordprocessing.New(
		"extended_props_example.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create document: %w",
			err,
		)
	}

	// Configure extended properties
	if err := configureExtendedProperties(doc); err != nil {
		_ = doc.Close()

		return nil, err
	}

	// Save the document
	if err := doc.Save(); err != nil {
		_ = doc.Close()

		return nil, fmt.Errorf(
			"failed to save document: %w",
			err,
		)
	}

	return doc, nil
}

func configureExtendedProperties(
	doc *wordprocessing.Document,
) error {
	// Get the extended properties
	ep := doc.ExtendedProperties()

	// Set application information
	ep.SetApplication(
		"Go Office - Extended Properties Example",
	)
	ep.SetApplicationVersion("1.0.0")
	ep.SetCompany("Example Corporation")
	ep.SetManager("Jane Doe")
	ep.SetTemplate("Normal.dotm")

	// Set document statistics
	//nolint:revive // Example values for demonstration
	ep.SetPages(10)
	//nolint:revive // Example values for demonstration
	ep.SetWords(2500)
	//nolint:revive // Example values for demonstration
	ep.SetCharacters(15000)
	//nolint:revive // Example values for demonstration
	ep.SetCharactersWithSpaces(17500)
	//nolint:revive // Example values for demonstration
	ep.SetLines(300)
	//nolint:revive // Example values for demonstration
	ep.SetParagraphs(150)

	// Set editing time (in minutes)
	//nolint:revive // Example value: 120 minutes = 2 hours
	ep.SetTotalTime(120)

	// Set other properties
	ep.SetHyperlinkBase("https://example.com")
	ep.SetScaleCrop(false)
	ep.SetLinksUpToDate(true)

	// Ensure the extended properties part is created and saved
	pkg := doc.Package().Package()
	if err := pkg.EnsureExtendedPropertiesPart(); err != nil {
		return fmt.Errorf(
			"failed to ensure extended properties part: %w",
			err,
		)
	}

	return nil
}

func displayExtendedProperties(
	ep *packaging.ExtendedProperties,
) {
	fmt.Println(
		"Document created with extended properties!",
	)
	fmt.Println("\nExtended Properties:")
	fmt.Printf(
		"  Application: %s\n",
		ep.Application(),
	)
	fmt.Printf(
		"  Version: %s\n",
		ep.ApplicationVersion(),
	)
	fmt.Printf("  Company: %s\n", ep.Company())
	fmt.Printf("  Manager: %s\n", ep.Manager())
	fmt.Printf("  Template: %s\n", ep.Template())
	fmt.Printf("  Pages: %d\n", ep.Pages())
	fmt.Printf("  Words: %d\n", ep.Words())
	fmt.Printf(
		"  Characters: %d\n",
		ep.Characters(),
	)
	fmt.Printf(
		"  Total Time: %d minutes\n",
		ep.TotalTime(),
	)
}

func readExtendedPropertiesExample() error {
	fmt.Println(
		"\n--- Reading Extended Properties ---",
	)

	doc2, err := wordprocessing.Open(
		"extended_props_example.docx",
		false,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to open document: %w",
			err,
		)
	}
	defer func() { _ = doc2.Close() }()

	ep2 := doc2.ExtendedProperties()
	fmt.Printf(
		"Read Application: %s\n",
		ep2.Application(),
	)
	fmt.Printf(
		"Read Company: %s\n",
		ep2.Company(),
	)
	fmt.Printf("Read Pages: %d\n", ep2.Pages())
	fmt.Printf("Read Words: %d\n", ep2.Words())

	return nil
}
