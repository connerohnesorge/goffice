//go:build ignore
// +build ignore

// Package main provides an example of using the comment API to add comments to Word documents.
package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func main() {
	// Create a new document
	doc, err := wordprocessing.New(
		"comment_example.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		log.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}

	// Get the main part and body
	mainPart := doc.MainPart()
	docElem := mainPart.Document()
	body := docElem.GetOrCreateBody()

	// Add a paragraph with multiple runs
	para := elements.NewParagraph()
	para.AppendRun("This is ")
	para.AppendRun("some text ")
	para.AppendRun("with a comment.")
	body.AppendChild(para)

	// Apply a comment to runs 1-2 (covering "some text ")
	comment, err := doc.ApplyComment(
		para,
		1, // Start at "some text "
		2, // End at "with a comment."
		"John Doe",
		"This needs revision!",
	)
	if err != nil {
		_ = doc.Close()
		log.Fatalf(
			"Failed to apply comment: %v",
			err,
		)
	}

	fmt.Printf(
		"Created comment with ID: %d\n",
		comment.Id(),
	)
	fmt.Printf(
		"Comment author: %s\n",
		comment.Author(),
	)
	fmt.Printf(
		"Comment date: %s\n",
		comment.Date(),
	)

	// Save the document
	if err := doc.Save(); err != nil {
		_ = doc.Close()
		log.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	// Close the document
	if err := doc.Close(); err != nil {
		log.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	fmt.Println("Document saved successfully!")
}
