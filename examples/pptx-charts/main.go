//go:build ignore
// +build ignore

// Package main demonstrates creating PowerPoint presentations with various chart types.
//
// This example shows how to create six different chart types using the goffice library:
// 1. Bar Chart (Clustered) - Quarterly sales comparison
// 2. Line Chart - Temperature trends over months
// 3. Pie Chart - Market share distribution
// 4. Area Chart - Cumulative revenue growth
// 5. Scatter Chart - Correlation analysis
// 6. Doughnut Chart - Budget allocation
//
// Each chart is placed on a separate slide with realistic fake data.
package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/presentation"
)

const outputFilename = "charts-example.pptx"

func main() {
	doc, err := presentation.New(
		outputFilename,
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}

	if err := createAllChartSlides(doc); err != nil {
		_ = doc.Close()
		log.Fatalf(
			"Failed to create charts: %v",
			err,
		)
	}

	if err := doc.Save(); err != nil {
		_ = doc.Close()
		log.Fatalf(
			"Failed to save presentation: %v",
			err,
		)
	}

	if err := doc.Close(); err != nil {
		log.Fatalf(
			"Failed to close presentation: %v",
			err,
		)
	}

	printSuccessMessage()
}

// createAllChartSlides creates all six chart slides in the presentation.
func createAllChartSlides(
	doc *presentation.Document,
) error {
	chartCreators := []func(*presentation.Document) error{
		createBarChartSlide,
		createLineChartSlide,
		createPieChartSlide,
		createAreaChartSlide,
		createScatterChartSlide,
		createDoughnutChartSlide,
	}

	for _, creator := range chartCreators {
		if err := creator(doc); err != nil {
			return err
		}
	}

	return nil
}

// printSuccessMessage prints a success message with all chart types created.
func printSuccessMessage() {
	fmt.Printf(
		"Successfully created %s with 6 chart types\n",
		outputFilename,
	)
	fmt.Println("Charts included:")
	fmt.Println(
		"  1. Bar Chart - Quarterly Sales Comparison",
	)
	fmt.Println(
		"  2. Line Chart - Monthly Temperature Trends",
	)
	fmt.Println(
		"  3. Pie Chart - Market Share Distribution",
	)
	fmt.Println(
		"  4. Area Chart - Cumulative Revenue Growth",
	)
	fmt.Println(
		"  5. Scatter Chart - Price vs Demand Analysis",
	)
	fmt.Println(
		"  6. Doughnut Chart - Annual Budget Allocation",
	)
}
