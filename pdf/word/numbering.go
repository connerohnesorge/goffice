// numbering.go provides numbering and list support for Word document rendering.
// It handles bullet rendering, number formatting, list indentation, and multi-level lists.

package word

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/parts"
)

// NumberingState tracks the current state of numbering across paragraphs.
type NumberingState struct {
	// counters maintains the counter for each numId and level.
	// Key format: "numId:level"
	counters map[string]int

	// numberingPart provides access to numbering definitions.
	numberingPart *parts.NumberingPart
}

// NewNumberingState creates a new numbering state tracker.
func NewNumberingState(
	numberingPart *parts.NumberingPart,
) *NumberingState {
	return &NumberingState{
		counters:      make(map[string]int),
		numberingPart: numberingPart,
	}
}

// GetNumberingText returns the formatted numbering text for a paragraph.
// Returns empty string if the paragraph doesn't have numbering.
func (ns *NumberingState) GetNumberingText(
	numId, level int,
) string {
	if ns.numberingPart == nil {
		return ""
	}

	// Get the numbering instance
	root := ns.numberingPart.RootElement()
	if root == nil {
		return ""
	}

	numbering := ns.getNumbering(root)
	if numbering == nil {
		return ""
	}

	numInstance := numbering.GetNumInstance(numId)
	if numInstance == nil {
		return ""
	}

	// Get the abstract numbering definition
	abstractNumId := numInstance.AbstractNumId()
	abstractNum := numbering.GetAbstractNum(
		abstractNumId,
	)
	if abstractNum == nil {
		return ""
	}

	// Get the level definition
	lvl := abstractNum.GetLevel(level)
	if lvl == nil {
		return ""
	}

	// Check for level override
	for override := range numInstance.LevelOverrides() {
		if override.LevelIndex() == level {
			if overrideLevel := override.Level(); overrideLevel != nil {
				lvl = overrideLevel
			}
			// Check for start override
			if startVal := override.StartOverride(); startVal >= 0 {
				ns.setCounter(
					numId,
					level,
					startVal-1,
				)
			}
			break
		}
	}

	// Increment the counter for this level
	ns.incrementCounter(numId, level)
	currentValue := ns.getCounter(numId, level)

	// Reset child levels when a parent level increments
	for childLevel := level + 1; childLevel < 9; childLevel++ {
		ns.resetCounter(numId, childLevel)
	}

	// Get the number format
	numFmt := lvl.NumberFormat()
	levelText := lvl.LevelText()

	// Format the numbering text
	return ns.formatNumberingText(
		levelText,
		numFmt,
		currentValue,
		numId,
		level,
	)
}

// GetNumberingIndent returns the left and hanging indentation for a numbered paragraph.
func (ns *NumberingState) GetNumberingIndent(
	numId, level int,
) (left, hanging float64) {
	if ns.numberingPart == nil {
		return 0, 0
	}

	root := ns.numberingPart.RootElement()
	if root == nil {
		return 0, 0
	}

	numbering := ns.getNumbering(root)
	if numbering == nil {
		return 0, 0
	}

	numInstance := numbering.GetNumInstance(numId)
	if numInstance == nil {
		return 0, 0
	}

	abstractNumId := numInstance.AbstractNumId()
	abstractNum := numbering.GetAbstractNum(
		abstractNumId,
	)
	if abstractNum == nil {
		return 0, 0
	}

	lvl := abstractNum.GetLevel(level)
	if lvl == nil {
		return 0, 0
	}

	// Get paragraph properties for indentation
	pPr := lvl.ParagraphProperties()
	if pPr == nil {
		return 0, 0
	}

	left = float64(
		pPr.Left(),
	) / 20.0 // Convert twips to points
	hanging = float64(pPr.Hanging()) / 20.0

	return left, hanging
}

// GetNumberingFont returns the font name for the numbering symbol.
func (ns *NumberingState) GetNumberingFont(
	numId, level int,
) string {
	if ns.numberingPart == nil {
		return ""
	}

	root := ns.numberingPart.RootElement()
	if root == nil {
		return ""
	}

	numbering := ns.getNumbering(root)
	if numbering == nil {
		return ""
	}

	numInstance := numbering.GetNumInstance(numId)
	if numInstance == nil {
		return ""
	}

	abstractNumId := numInstance.AbstractNumId()
	abstractNum := numbering.GetAbstractNum(
		abstractNumId,
	)
	if abstractNum == nil {
		return ""
	}

	lvl := abstractNum.GetLevel(level)
	if lvl == nil {
		return ""
	}

	// Get run properties for font
	rPr := lvl.RunProperties()
	if rPr == nil {
		return ""
	}

	// Get font from rFonts element
	rFonts := rPr.GetElement(
		"rFonts",
		elements.NamespaceWML,
	)
	if rFonts == nil {
		return ""
	}

	// Try ascii attribute first
	if attr, found := rFonts.GetAttribute("ascii", elements.NamespaceWML); found {
		return attr.Value()
	}

	// Fall back to hAnsi
	if attr, found := rFonts.GetAttribute("hAnsi", elements.NamespaceWML); found {
		return attr.Value()
	}

	return ""
}

// formatNumberingText formats the level text with the current counter value.
func (ns *NumberingState) formatNumberingText(
	levelText string,
	numFmt elements.NumberFormatValue,
	currentValue int,
	numId, level int,
) string {
	// Handle bullet format specially
	if numFmt == elements.NumberFormatBullet {
		return levelText
	}

	// Replace placeholders like %1, %2, %3 with actual values
	result := levelText

	// Process all levels that might be referenced in the text
	for lvl := 0; lvl <= level; lvl++ {
		placeholder := fmt.Sprintf("%%%d", lvl+1)
		if !strings.Contains(
			result,
			placeholder,
		) {
			continue
		}

		// Get the counter for this level
		var value int
		if lvl == level {
			value = currentValue
		} else {
			value = ns.getCounter(numId, lvl)
		}

		// Get the format for this level
		var fmt elements.NumberFormatValue
		if lvl == level {
			fmt = numFmt
		} else {
			// Need to look up the format for this level
			fmt = ns.getNumberFormat(numId, lvl)
		}

		// Format the value
		formatted := ns.formatNumber(value, fmt)
		result = strings.ReplaceAll(
			result,
			placeholder,
			formatted,
		)
	}

	return result
}

// formatNumber formats a number according to the specified format.
func (ns *NumberingState) formatNumber(
	value int,
	format elements.NumberFormatValue,
) string {
	switch format {
	case elements.NumberFormatDecimal:
		return strconv.Itoa(value)

	case elements.NumberFormatDecimalZero:
		return fmt.Sprintf("%02d", value)

	case elements.NumberFormatUpperRoman:
		return toRomanUpper(value)

	case elements.NumberFormatLowerRoman:
		return toRomanLower(value)

	case elements.NumberFormatUpperLetter:
		return toLetterUpper(value)

	case elements.NumberFormatLowerLetter:
		return toLetterLower(value)

	case elements.NumberFormatOrdinal:
		return toOrdinal(value)

	case elements.NumberFormatBullet:
		return "\u2022" // Default bullet

	default:
		return strconv.Itoa(value)
	}
}

// getNumberFormat retrieves the number format for a specific level.
func (ns *NumberingState) getNumberFormat(
	numId, level int,
) elements.NumberFormatValue {
	if ns.numberingPart == nil {
		return elements.NumberFormatDecimal
	}

	root := ns.numberingPart.RootElement()
	if root == nil {
		return elements.NumberFormatDecimal
	}

	numbering := ns.getNumbering(root)
	if numbering == nil {
		return elements.NumberFormatDecimal
	}

	numInstance := numbering.GetNumInstance(numId)
	if numInstance == nil {
		return elements.NumberFormatDecimal
	}

	abstractNumId := numInstance.AbstractNumId()
	abstractNum := numbering.GetAbstractNum(
		abstractNumId,
	)
	if abstractNum == nil {
		return elements.NumberFormatDecimal
	}

	lvl := abstractNum.GetLevel(level)
	if lvl == nil {
		return elements.NumberFormatDecimal
	}

	return lvl.NumberFormat()
}

// Counter management helpers

func (ns *NumberingState) getCounter(
	numId, level int,
) int {
	key := fmt.Sprintf("%d:%d", numId, level)
	if val, ok := ns.counters[key]; ok {
		return val
	}

	// Get the start value from the level definition
	start := ns.getStartValue(numId, level)
	return start
}

func (ns *NumberingState) setCounter(
	numId, level, value int,
) {
	key := fmt.Sprintf("%d:%d", numId, level)
	ns.counters[key] = value
}

func (ns *NumberingState) incrementCounter(
	numId, level int,
) {
	current := ns.getCounter(numId, level)
	ns.setCounter(numId, level, current+1)
}

func (ns *NumberingState) resetCounter(
	numId, level int,
) {
	start := ns.getStartValue(numId, level)
	ns.setCounter(numId, level, start)
}

func (ns *NumberingState) getStartValue(
	numId, level int,
) int {
	if ns.numberingPart == nil {
		return 1
	}

	root := ns.numberingPart.RootElement()
	if root == nil {
		return 1
	}

	numbering := ns.getNumbering(root)
	if numbering == nil {
		return 1
	}

	numInstance := numbering.GetNumInstance(numId)
	if numInstance == nil {
		return 1
	}

	abstractNumId := numInstance.AbstractNumId()
	abstractNum := numbering.GetAbstractNum(
		abstractNumId,
	)
	if abstractNum == nil {
		return 1
	}

	lvl := abstractNum.GetLevel(level)
	if lvl == nil {
		return 1
	}

	return lvl.Start()
}

// Number formatting helper functions

func toRomanUpper(num int) string {
	if num <= 0 {
		return ""
	}

	values := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	symbols := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}

	var result strings.Builder
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result.WriteString(symbols[i])
			num -= values[i]
		}
	}

	return result.String()
}

func toRomanLower(num int) string {
	return strings.ToLower(toRomanUpper(num))
}

func toLetterUpper(num int) string {
	if num <= 0 {
		return ""
	}

	var result strings.Builder
	num-- // Convert to 0-based

	for {
		result.WriteByte(byte('A' + (num % 26)))
		num = num / 26
		if num == 0 {
			break
		}
		num-- // Adjust for 1-based alphabet
	}

	// Reverse the string
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func toLetterLower(num int) string {
	return strings.ToLower(toLetterUpper(num))
}

func toOrdinal(num int) string {
	suffix := "th"

	switch num % 10 {
	case 1:
		if num%100 != 11 {
			suffix = "st"
		}
	case 2:
		if num%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if num%100 != 13 {
			suffix = "rd"
		}
	}

	return strconv.Itoa(num) + suffix
}

// getNumbering converts a PartRootElement to a Numbering element.
func (ns *NumberingState) getNumbering(
	root interface{},
) *elements.Numbering {
	// Try direct type assertion
	if num, ok := root.(*elements.Numbering); ok {
		return num
	}

	// Try CompositeElementBase wrapper
	if comp, ok := root.(*openxml.CompositeElementBase); ok {
		return &elements.Numbering{
			CompositeElementBase: comp,
		}
	}

	return nil
}
