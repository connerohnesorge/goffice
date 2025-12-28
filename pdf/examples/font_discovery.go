// This example demonstrates how to discover system fonts using the PDF renderer.
package main

import (
	"fmt"
	"sort"

	pdf "github.com/connerohnesorge/goffice-pdf"
)

func main() {
	// Discover all available system fonts
	fonts := pdf.DiscoverFonts()

	fmt.Printf(
		"Discovered %d fonts on your system\n\n",
		len(fonts),
	)

	// Group fonts by family
	familyMap := make(map[string][]string)
	for _, font := range fonts {
		familyMap[font.Family] = append(
			familyMap[font.Family],
			font.Style,
		)
	}

	// Sort families for consistent output
	families := make([]string, 0, len(familyMap))
	for family := range familyMap {
		families = append(families, family)
	}
	sort.Strings(families)

	// Print first 20 font families
	count := 20
	if len(families) < 20 {
		count = len(families)
	}

	fmt.Println(
		"Sample font families and their variants:",
	)
	fmt.Println(
		"=========================================",
	)
	for i := 0; i < count; i++ {
		family := families[i]
		styles := familyMap[family]

		fmt.Printf("\n%s:\n", family)
		for _, style := range styles {
			fmt.Printf("  - %s\n", style)
		}
	}

	if len(families) > 20 {
		fmt.Printf(
			"\n... and %d more font families\n",
			len(families)-20,
		)
	}
}
