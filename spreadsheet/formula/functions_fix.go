package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the file
	content, err := os.ReadFile("/home/connerohnesorge/Documents/001Repos/goffice/spreadsheet/formula/functions.go")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Split into lines
	lines := strings.Split(string(content), "\n")

	// Process each line to add missing opening braces
	var result strings.Builder
	for i, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check if line is a function struct definition that needs opening braces
		if strings.Contains(line, "type") && strings.Contains(line, "struct") && strings.Contains(line, "Function") {
			// Find the struct name
			parts := strings.Fields(line)
			if len(parts) >= 3 && parts[0] == "type" && parts[1] == "Function" {
				funcName := parts[1]
				
				// Check if struct already has opening braces
				if strings.Contains(line, "{") {
					// Already has opening brace, continue
					continue
				}
				
				// Add opening brace after function name
				fixedLine := strings.Replace(line, "type "+funcName+" struct", "type "+funcName+" struct {")
				result.WriteString(fixedLine)
				fmt.Printf("Fixed: %s\n", fixedLine)
			} else {
				result.WriteString(line)
				fmt.Printf("No change needed: %s\n", line)
			}
		} else {
			result.WriteString(line)
			fmt.Printf("No change needed: %s\n", line)
			}
		}
	}

	// Write result
	err := os.WriteFile("/home/connerohnesorge/Documents/001Repos/goffice/spreadsheet/formula/functions.go", result.String())
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
	}
}