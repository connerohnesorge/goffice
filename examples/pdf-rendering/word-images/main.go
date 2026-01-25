//go:build ignore
// +build ignore

// Package main demonstrates Word document to PDF rendering with DrawingML images.
//
// This example shows how to:
// 1. Create a Word document with inline images
// 2. Render the document to PDF using DrawingML rendering
// 3. Generate a professional-looking PDF with embedded images
//
// The DrawingML rendering system handles:
// - Image placement and scaling
// - Multiple image formats (PNG, JPEG)
// - Mixed text and image content
// - Professional layout and spacing
package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	pdfword "github.com/connerohnesorge/goffice-pdf/word"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/parts"
)

const (
	docxFilename = "word_images_example.docx"
	pdfFilename  = "word_images_example.pdf"
)

func main() {
	fmt.Println(
		"Creating Word document with images...",
	)

	// Create the Word document
	if err := createWordDocument(); err != nil {
		log.Fatalf(
			"Failed to create Word document: %v",
			err,
		)
	}

	fmt.Printf("Created %s\n", docxFilename)
	fmt.Println("Rendering to PDF...")

	// Render to PDF
	if err := renderToPDF(); err != nil {
		log.Fatalf(
			"Failed to render PDF: %v",
			err,
		)
	}

	fmt.Printf(
		"Successfully rendered to %s\n",
		pdfFilename,
	)
	fmt.Println(
		"\nThe PDF demonstrates DrawingML image rendering:",
	)
	fmt.Println(
		"  - Inline images with proper scaling",
	)
	fmt.Println("  - Multiple image sizes")
	fmt.Println(
		"  - Mixed text and image content",
	)
	fmt.Println(
		"  - Professional layout and spacing",
	)
}

// createWordDocument creates a Word document with multiple images.
func createWordDocument() error {
	doc, err := wordprocessing.New(
		docxFilename,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		return fmt.Errorf(
			"creating document: %w",
			err,
		)
	}
	defer doc.Close()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, err = doc.AddMainPart()
		if err != nil {
			return fmt.Errorf(
				"adding main part: %w",
				err,
			)
		}
	}
	mainPart.InitializeContent()
	if err := mainPart.Reload(); err != nil {
		return fmt.Errorf(
			"reloading main part: %w",
			err,
		)
	}

	docElem := mainPart.Document()
	body := docElem.Body()

	// Add title
	titlePara := body.AppendParagraph("")
	titleRun := titlePara.AppendRun(
		"DrawingML Image Rendering Example",
	)
	titleRun.SetBold(true)
	titleRun.SetFontSize(32) // 16pt (half-points)

	// Add introduction
	body.AppendParagraph(
		"This document demonstrates the DrawingML to PDF rendering capabilities of goffice. " +
			"Images are rendered with high fidelity, maintaining their aspect ratios and quality.",
	)

	// Simple 1x1 PNG image (smallest valid PNG)
	pngData := createSimplePNG()

	// Add small image
	body.AppendParagraph(
		"Small Image (1.5 x 1.5 inches):",
	)
	if err := addImage(mainPart, body, pngData, 1.5, 1.5); err != nil {
		return fmt.Errorf(
			"adding small image: %w",
			err,
		)
	}

	// Add medium image
	body.AppendParagraph(
		"Medium Image (3 x 2 inches):",
	)
	if err := addImage(mainPart, body, pngData, 3.0, 2.0); err != nil {
		return fmt.Errorf(
			"adding medium image: %w",
			err,
		)
	}

	// Add large image
	body.AppendParagraph(
		"Large Image (5 x 3.5 inches):",
	)
	if err := addImage(mainPart, body, pngData, 5.0, 3.5); err != nil {
		return fmt.Errorf(
			"adding large image: %w",
			err,
		)
	}

	// Add conclusion
	body.AppendParagraph(
		"All images above are rendered using the DrawingML rendering system, " +
			"which ensures consistent appearance between the Word document and the generated PDF.",
	)

	if err := doc.Save(); err != nil {
		return fmt.Errorf(
			"saving document: %w",
			err,
		)
	}

	return nil
}

// renderToPDF renders the Word document to PDF.
func renderToPDF() error {
	// Open the document
	doc, err := wordprocessing.Open(
		docxFilename,
		false,
	)
	if err != nil {
		return fmt.Errorf(
			"opening document: %w",
			err,
		)
	}
	defer doc.Close()

	// Set up font cache and layout engine
	fontCache := font.NewFontCache(10)
	layoutEngine := layout.NewTextLayoutEngine(
		fontCache,
	)

	// Create renderer
	renderer, err := pdfword.NewWordRenderer(
		doc,
		layoutEngine,
	)
	if err != nil {
		return fmt.Errorf(
			"creating renderer: %w",
			err,
		)
	}

	// Render to PDF
	absPath, err := filepath.Abs(pdfFilename)
	if err != nil {
		return fmt.Errorf(
			"resolving path: %w",
			err,
		)
	}

	if err := renderer.Render(absPath); err != nil {
		return fmt.Errorf(
			"rendering PDF: %w",
			err,
		)
	}

	return nil
}

// addImage adds an inline image to the document body.
func addImage(
	mainPart *parts.MainPart,
	body *elements.Body,
	imageData []byte,
	widthInches float64,
	heightInches float64,
) error {
	p := body.AppendParagraph("")
	r := p.AppendRun("")

	// Add image part
	imagePart, err := mainPart.AddImagePart(
		parts.ImageTypePng,
	)
	if err != nil {
		return fmt.Errorf(
			"adding image part: %w",
			err,
		)
	}

	imagePart.FeedDataBytes(imageData)
	relID := imagePart.RelationshipID()

	// Convert inches to EMU (English Metric Units)
	width := int64(
		widthInches * float64(
			elements.EMUsPerInch,
		),
	)
	height := int64(
		heightInches * float64(
			elements.EMUsPerInch,
		),
	)

	// Create inline drawing
	drawing := elements.NewInlineDrawing(
		width,
		height,
		relID,
	)
	r.AppendChild(drawing)

	// Add spacing
	body.AppendParagraph("")

	return nil
}

// createSimplePNG creates a minimal valid PNG image (1x1 pixel, white).
func createSimplePNG() []byte {
	return []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG signature
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52, // IHDR chunk
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, // 1x1
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41, // IDAT chunk
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xCF, 0xC0, 0x00,
		0x00, 0x03, 0x01, 0x01, 0x00, 0x18, 0xDD, 0x8D,
		0xB4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, // IEND chunk
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}
}
