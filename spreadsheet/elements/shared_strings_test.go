package elements

import (
	"strings"
	"testing"
)

func TestSharedStringTableCreation(t *testing.T) {
	sst := NewSharedStringTable()

	if sst == nil {
		t.Fatal(
			"NewSharedStringTable returned nil",
		)
	}

	if sst.LocalName() != "sst" {
		t.Errorf(
			"expected local name 'sst', got '%s'",
			sst.LocalName(),
		)
	}

	if sst.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			sst.NamespaceURI(),
		)
	}

	// Initially empty
	if sst.ItemCount() != 0 {
		t.Errorf(
			"expected 0 items, got %d",
			sst.ItemCount(),
		)
	}

	if sst.Count() != 0 {
		t.Errorf(
			"expected count 0, got %d",
			sst.Count(),
		)
	}

	if sst.UniqueCount() != 0 {
		t.Errorf(
			"expected uniqueCount 0, got %d",
			sst.UniqueCount(),
		)
	}
}

func TestSharedStringTableAddString(
	t *testing.T,
) {
	sst := NewSharedStringTable()

	// Add first string
	idx0 := sst.AddString("Hello")
	if idx0 != 0 {
		t.Errorf("expected index 0, got %d", idx0)
	}

	// Add second string
	idx1 := sst.AddString("World")
	if idx1 != 1 {
		t.Errorf("expected index 1, got %d", idx1)
	}

	// Add duplicate - should return existing index
	idx2 := sst.AddString("Hello")
	if idx2 != 0 {
		t.Errorf(
			"expected index 0 for duplicate, got %d",
			idx2,
		)
	}

	// Check count
	if sst.ItemCount() != 2 {
		t.Errorf(
			"expected 2 unique items, got %d",
			sst.ItemCount(),
		)
	}

	// Check uniqueCount attribute
	if sst.UniqueCount() != 2 {
		t.Errorf(
			"expected uniqueCount 2, got %d",
			sst.UniqueCount(),
		)
	}
}

func TestSharedStringTableGetString(
	t *testing.T,
) {
	sst := NewSharedStringTable()

	sst.AddString("First")
	sst.AddString("Second")
	sst.AddString("Third")

	testCases := []struct {
		index    int
		expected string
	}{
		{0, "First"},
		{1, "Second"},
		{2, "Third"},
		{-1, ""}, // Invalid index
		{10, ""}, // Out of range
	}

	for _, tc := range testCases {
		result := sst.GetString(tc.index)
		if result != tc.expected {
			t.Errorf(
				"GetString(%d): expected '%s', got '%s'",
				tc.index,
				tc.expected,
				result,
			)
		}
	}
}

func TestSharedStringTableIndexOf(t *testing.T) {
	sst := NewSharedStringTable()

	sst.AddString("Apple")
	sst.AddString("Banana")
	sst.AddString("Cherry")

	testCases := []struct {
		text     string
		expected int
	}{
		{"Apple", 0},
		{"Banana", 1},
		{"Cherry", 2},
		{"Date", -1}, // Not found
		{"", -1},     // Empty string not added
	}

	for _, tc := range testCases {
		result := sst.IndexOf(tc.text)
		if result != tc.expected {
			t.Errorf(
				"IndexOf('%s'): expected %d, got %d",
				tc.text,
				tc.expected,
				result,
			)
		}
	}
}

func TestSharedStringTableDeduplication(
	t *testing.T,
) {
	sst := NewSharedStringTable()

	// Add same string multiple times
	testStrings := []string{
		"Same",
		"Same",
		"Different",
		"Same",
		"Different",
	}
	expectedIndices := []int{0, 0, 1, 0, 1}

	for i, s := range testStrings {
		idx := sst.AddString(s)
		if idx != expectedIndices[i] {
			t.Errorf(
				"AddString('%s') at iteration %d: expected %d, got %d",
				s,
				i,
				expectedIndices[i],
				idx,
			)
		}
	}

	// Should only have 2 unique items
	if sst.ItemCount() != 2 {
		t.Errorf(
			"expected 2 unique items after deduplication, got %d",
			sst.ItemCount(),
		)
	}
}

func TestSharedStringTableGetItem(t *testing.T) {
	sst := NewSharedStringTable()

	sst.AddString("Item0")
	sst.AddString("Item1")

	// Valid indices
	item0 := sst.GetItem(0)
	if item0 == nil {
		t.Fatal("GetItem(0) returned nil")
	}
	if item0.PlainText() != "Item0" {
		t.Errorf(
			"expected 'Item0', got '%s'",
			item0.PlainText(),
		)
	}

	item1 := sst.GetItem(1)
	if item1 == nil {
		t.Fatal("GetItem(1) returned nil")
	}
	if item1.PlainText() != "Item1" {
		t.Errorf(
			"expected 'Item1', got '%s'",
			item1.PlainText(),
		)
	}

	// Invalid indices
	if sst.GetItem(-1) != nil {
		t.Error("expected nil for negative index")
	}

	if sst.GetItem(10) != nil {
		t.Error(
			"expected nil for out of range index",
		)
	}
}

func TestSharedStringTableClear(t *testing.T) {
	sst := NewSharedStringTable()

	sst.AddString("A")
	sst.AddString("B")
	sst.SetCount(
		5,
	) // Simulating 5 total references

	sst.Clear()

	if sst.ItemCount() != 0 {
		t.Errorf(
			"expected 0 items after clear, got %d",
			sst.ItemCount(),
		)
	}

	if sst.Count() != 0 {
		t.Errorf(
			"expected count 0 after clear, got %d",
			sst.Count(),
		)
	}

	if sst.UniqueCount() != 0 {
		t.Errorf(
			"expected uniqueCount 0 after clear, got %d",
			sst.UniqueCount(),
		)
	}

	// Should be able to add new strings after clear
	idx := sst.AddString("New")
	if idx != 0 {
		t.Errorf(
			"expected index 0 for first string after clear, got %d",
			idx,
		)
	}
}

func TestSharedStringTableIteration(
	t *testing.T,
) {
	sst := NewSharedStringTable()

	expected := []string{"One", "Two", "Three"}
	for _, s := range expected {
		sst.AddString(s)
	}

	i := 0
	for item := range sst.Items() {
		if item.PlainText() != expected[i] {
			t.Errorf(
				"iteration %d: expected '%s', got '%s'",
				i,
				expected[i],
				item.PlainText(),
			)
		}
		i++
	}

	if i != len(expected) {
		t.Errorf(
			"expected %d iterations, got %d",
			len(expected),
			i,
		)
	}
}

func TestSharedStringItemCreation(t *testing.T) {
	si := NewSharedStringItem()

	if si == nil {
		t.Fatal(
			"NewSharedStringItem returned nil",
		)
	}

	if si.LocalName() != "si" {
		t.Errorf(
			"expected local name 'si', got '%s'",
			si.LocalName(),
		)
	}

	// Initially no text
	if si.PlainText() != "" {
		t.Errorf(
			"expected empty text, got '%s'",
			si.PlainText(),
		)
	}

	if si.IsRichText() {
		t.Error(
			"expected IsRichText() to be false initially",
		)
	}
}

func TestSharedStringItemSimpleText(
	t *testing.T,
) {
	si := NewSharedStringItem()

	si.SetPlainText("Hello World")

	if si.PlainText() != "Hello World" {
		t.Errorf(
			"expected 'Hello World', got '%s'",
			si.PlainText(),
		)
	}

	// Text element should exist
	text := si.Text()
	if text == nil {
		t.Fatal(
			"Text() returned nil after SetPlainText",
		)
	}

	if text.Text() != "Hello World" {
		t.Errorf(
			"expected 'Hello World' from Text(), got '%s'",
			text.Text(),
		)
	}

	// Should not be rich text
	if si.IsRichText() {
		t.Error(
			"expected IsRichText() to be false for simple text",
		)
	}
}

func TestSharedStringItemRichText(t *testing.T) {
	si := NewSharedStringItem()

	// Add rich text runs
	run1 := si.AddRichTextRun()
	run1.SetTextContent("Bold ")

	run2 := si.AddRichTextRun()
	run2.SetTextContent("Italic")

	// Should be rich text
	if !si.IsRichText() {
		t.Error(
			"expected IsRichText() to be true",
		)
	}

	// Plain text should concatenate
	if si.PlainText() != "Bold Italic" {
		t.Errorf(
			"expected 'Bold Italic', got '%s'",
			si.PlainText(),
		)
	}

	// Count runs
	count := 0
	for range si.RichTextRuns() {
		count++
	}
	if count != 2 {
		t.Errorf(
			"expected 2 rich text runs, got %d",
			count,
		)
	}
}

func TestSharedStringItemRichTextWithProperties(
	t *testing.T,
) {
	si := NewSharedStringItem()

	run := si.AddRichTextRunWithText(
		"Styled Text",
	)
	props := run.GetOrCreateRunProperties()
	props.SetBold(true)
	props.SetItalic(true)

	// Verify run properties
	if !props.IsBold() {
		t.Error("expected bold to be true")
	}

	if !props.IsItalic() {
		t.Error("expected italic to be true")
	}

	// Verify text
	if run.TextContent() != "Styled Text" {
		t.Errorf(
			"expected 'Styled Text', got '%s'",
			run.TextContent(),
		)
	}
}

func TestSharedStringItemSwitchToRichText(
	t *testing.T,
) {
	si := NewSharedStringItem()

	// Start with simple text
	si.SetPlainText("Simple")

	if si.Text() == nil {
		t.Fatal("expected text element")
	}

	// Switch to rich text
	run := si.AddRichTextRun()
	run.SetTextContent("Rich")

	// Simple text should be removed
	if si.Text() != nil {
		t.Error(
			"expected text element to be removed when adding rich text",
		)
	}

	// Should now be rich text
	if !si.IsRichText() {
		t.Error(
			"expected IsRichText() to be true after adding run",
		)
	}
}

func TestSharedStringItemSwitchToSimpleText(
	t *testing.T,
) {
	si := NewSharedStringItem()

	// Start with rich text
	run1 := si.AddRichTextRun()
	run1.SetTextContent("Part1")
	run2 := si.AddRichTextRun()
	run2.SetTextContent("Part2")

	if !si.IsRichText() {
		t.Fatal("expected rich text")
	}

	// Switch to simple text
	si.SetPlainText("Simple Again")

	// Rich text runs should be removed
	if si.IsRichText() {
		t.Error(
			"expected IsRichText() to be false after SetPlainText",
		)
	}

	if si.PlainText() != "Simple Again" {
		t.Errorf(
			"expected 'Simple Again', got '%s'",
			si.PlainText(),
		)
	}
}

func TestPhoneticRun(t *testing.T) {
	rph := NewPhoneticRun()

	if rph == nil {
		t.Fatal("NewPhoneticRun returned nil")
	}

	if rph.LocalName() != "rPh" {
		t.Errorf(
			"expected local name 'rPh', got '%s'",
			rph.LocalName(),
		)
	}

	// Set indices
	rph.SetSb(0)
	rph.SetEb(3)

	if rph.Sb() != 0 {
		t.Errorf(
			"expected Sb 0, got %d",
			rph.Sb(),
		)
	}

	if rph.Eb() != 3 {
		t.Errorf(
			"expected Eb 3, got %d",
			rph.Eb(),
		)
	}

	// Set text
	rph.SetTextContent("phonetic")

	if rph.TextContent() != "phonetic" {
		t.Errorf(
			"expected 'phonetic', got '%s'",
			rph.TextContent(),
		)
	}
}

func TestPhoneticPr(t *testing.T) {
	pr := NewPhoneticPr()

	if pr == nil {
		t.Fatal("NewPhoneticPr returned nil")
	}

	if pr.LocalName() != "phoneticPr" {
		t.Errorf(
			"expected local name 'phoneticPr', got '%s'",
			pr.LocalName(),
		)
	}

	// Default values
	if pr.Type() != PhoneticTypeFullwidthKatakana {
		t.Errorf(
			"expected default type '%s', got '%s'",
			PhoneticTypeFullwidthKatakana,
			pr.Type(),
		)
	}

	if pr.Alignment() != PhoneticAlignmentLeft {
		t.Errorf(
			"expected default alignment '%s', got '%s'",
			PhoneticAlignmentLeft,
			pr.Alignment(),
		)
	}

	// Set values
	pr.SetFontId(5)
	pr.SetType(PhoneticTypeHiragana)
	pr.SetAlignment(PhoneticAlignmentCenter)

	if pr.FontId() != 5 {
		t.Errorf(
			"expected fontId 5, got %d",
			pr.FontId(),
		)
	}

	if pr.Type() != PhoneticTypeHiragana {
		t.Errorf("expected type '%s', got '%s'",
			PhoneticTypeHiragana, pr.Type())
	}

	if pr.Alignment() != PhoneticAlignmentCenter {
		t.Errorf(
			"expected alignment '%s', got '%s'",
			PhoneticAlignmentCenter,
			pr.Alignment(),
		)
	}
}

func TestSharedStringItemWithPhonetic(
	t *testing.T,
) {
	si := NewSharedStringItem()

	si.SetPlainText("Japanese Text")

	// Add phonetic run
	rph := si.AddPhoneticRun(0, 8)
	rph.SetTextContent("furigana")

	// Add phonetic properties
	pr := si.GetOrCreatePhoneticPr()
	pr.SetType(PhoneticTypeHiragana)
	pr.SetFontId(1)

	// Check phonetic runs exist
	count := 0
	for range si.PhoneticRuns() {
		count++
	}
	if count != 1 {
		t.Errorf(
			"expected 1 phonetic run, got %d",
			count,
		)
	}

	// Check phonetic properties
	if si.PhoneticPr() == nil {
		t.Fatal("expected phonetic properties")
	}

	if si.PhoneticPr().
		Type() !=
		PhoneticTypeHiragana {
		t.Errorf(
			"expected type '%s', got '%s'",
			PhoneticTypeHiragana,
			si.PhoneticPr().Type(),
		)
	}
}

func TestSharedStringTableRebuildIndex(
	t *testing.T,
) {
	sst := NewSharedStringTable()

	// Manually add items without using AddString
	si1 := NewSharedStringItem()
	si1.SetPlainText("Manual1")
	sst.AppendChild(si1)

	si2 := NewSharedStringItem()
	si2.SetPlainText("Manual2")
	sst.AppendChild(si2)

	// IndexOf should work after rebuilding
	sst.RebuildIndex()

	if sst.IndexOf("Manual1") != 0 {
		t.Errorf(
			"expected index 0 for 'Manual1', got %d",
			sst.IndexOf("Manual1"),
		)
	}

	if sst.IndexOf("Manual2") != 1 {
		t.Errorf(
			"expected index 1 for 'Manual2', got %d",
			sst.IndexOf("Manual2"),
		)
	}
}

func TestSharedStringTableClone(t *testing.T) {
	sst := NewSharedStringTable()
	sst.AddString("Clone1")
	sst.AddString("Clone2")
	sst.SetCount(10)

	clonedResult := sst.Clone()
	cloned, ok := clonedResult.(*SharedStringTable)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *SharedStringTable",
			clonedResult,
		)
	}

	// Verify cloned data
	if cloned.ItemCount() != 2 {
		t.Errorf(
			"expected 2 items in clone, got %d",
			cloned.ItemCount(),
		)
	}

	if cloned.GetString(0) != "Clone1" {
		t.Errorf(
			"expected 'Clone1', got '%s'",
			cloned.GetString(0),
		)
	}

	if cloned.GetString(1) != "Clone2" {
		t.Errorf(
			"expected 'Clone2', got '%s'",
			cloned.GetString(1),
		)
	}

	// Index should be preserved
	if cloned.IndexOf("Clone1") != 0 {
		t.Errorf(
			"expected index 0, got %d",
			cloned.IndexOf("Clone1"),
		)
	}

	// Modifying clone should not affect original
	cloned.AddString("Clone3")
	if sst.IndexOf("Clone3") != -1 {
		t.Error("original should not have Clone3")
	}
}

func TestSharedStringItemClone(t *testing.T) {
	si := NewSharedStringItem()
	run := si.AddRichTextRunWithText("Original")
	run.GetOrCreateRunProperties().SetBold(true)

	clonedResult := si.Clone()
	cloned, ok := clonedResult.(*SharedStringItem)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *SharedStringItem",
			clonedResult,
		)
	}

	// Verify cloned data
	if cloned.PlainText() != "Original" {
		t.Errorf(
			"expected 'Original', got '%s'",
			cloned.PlainText(),
		)
	}

	if !cloned.IsRichText() {
		t.Error(
			"expected cloned item to be rich text",
		)
	}

	// Modifying clone should not affect original
	for run := range cloned.RichTextRuns() {
		run.SetTextContent("Modified")
	}

	if si.PlainText() == "Modified" {
		t.Error("original should not be modified")
	}
}

func TestSharedStringTableXML(t *testing.T) {
	sst := NewSharedStringTable()
	sst.AddString("Test String")
	sst.SetCount(1)

	xml := sst.OuterXml()

	// Check XML contains expected elements
	if !strings.Contains(xml, "<sst") {
		t.Error("XML should contain <sst element")
	}

	if !strings.Contains(
		xml,
		"uniqueCount=\"1\"",
	) {
		t.Error(
			"XML should contain uniqueCount attribute",
		)
	}

	if !strings.Contains(xml, "<si>") ||
		!strings.Contains(xml, "</si>") {
		t.Error("XML should contain <si> element")
	}

	if !strings.Contains(
		xml,
		"<t>Test String</t>",
	) {
		t.Error("XML should contain text content")
	}
}

func TestEmptyString(t *testing.T) {
	sst := NewSharedStringTable()

	// Adding empty string should work
	idx := sst.AddString("")
	if idx != 0 {
		t.Errorf(
			"expected index 0 for empty string, got %d",
			idx,
		)
	}

	// Getting empty string should return empty
	if sst.GetString(0) != "" {
		t.Errorf(
			"expected empty string, got '%s'",
			sst.GetString(0),
		)
	}

	// Adding empty string again should return same index
	idx2 := sst.AddString("")
	if idx2 != 0 {
		t.Errorf(
			"expected index 0 for duplicate empty string, got %d",
			idx2,
		)
	}
}

func TestLargeSharedStringTable(t *testing.T) {
	sst := NewSharedStringTable()

	// Add many strings
	const numStrings = 1000
	for i := range numStrings {
		idx := sst.AddString(
			strings.Repeat("x", i+1),
		)
		if idx != i {
			t.Errorf(
				"expected index %d, got %d",
				i,
				idx,
			)
		}
	}

	// Verify count
	if sst.ItemCount() != numStrings {
		t.Errorf(
			"expected %d items, got %d",
			numStrings,
			sst.ItemCount(),
		)
	}

	// Verify lookup still works
	for i := range numStrings {
		expectedStr := strings.Repeat("x", i+1)
		idx := sst.IndexOf(expectedStr)
		if idx != i {
			t.Errorf(
				"IndexOf for string %d: expected %d, got %d",
				i,
				i,
				idx,
			)
		}
	}
}
