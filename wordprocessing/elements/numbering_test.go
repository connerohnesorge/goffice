package elements

import (
	"testing"
)

func TestLevelIndentation(t *testing.T) {
	lvl := NewLevel(0)
	lvl.SetIndentation(720, 360)

	pp := lvl.ParagraphProperties()
	if pp == nil {
		t.Fatal(
			"ParagraphProperties() should not be nil after SetIndentation",
		)
	}

	if pp.Left() != 720 {
		t.Errorf(
			"Expected Left 720, got %d",
			pp.Left(),
		)
	}
	if pp.Hanging() != 360 {
		t.Errorf(
			"Expected Hanging 360, got %d",
			pp.Hanging(),
		)
	}
}

func TestNumberingAddAbstractNum(t *testing.T) {
	numbering := NewNumbering()

	an1 := NewAbstractNum()
	id1 := numbering.AddAbstractNum(an1)
	if id1 != 0 {
		t.Errorf(
			"Expected first AbstractNum ID 0, got %d",
			id1,
		)
	}

	an2 := NewAbstractNum()
	id2 := numbering.AddAbstractNum(an2)
	if id2 != 1 {
		t.Errorf(
			"Expected second AbstractNum ID 1, got %d",
			id2,
		)
	}

	// Verify we can find them
	found := numbering.GetAbstractNum(0)
	if found == nil {
		t.Error(
			"GetAbstractNum(0) should not return nil",
		)
	}

	found = numbering.GetAbstractNum(1)
	if found == nil {
		t.Error(
			"GetAbstractNum(1) should not return nil",
		)
	}
}

func TestNumberingAddNumInstance(t *testing.T) {
	numbering := NewNumbering()

	// Add abstract num first
	an := NewAbstractNum()
	anId := numbering.AddAbstractNum(an)

	// Add numbering instances
	ni1 := NewNumberingInstance(anId)
	id1 := numbering.AddNumInstance(ni1)
	if id1 != 1 {
		t.Errorf(
			"Expected first NumInstance ID 1, got %d",
			id1,
		)
	}

	ni2 := NewNumberingInstance(anId)
	id2 := numbering.AddNumInstance(ni2)
	if id2 != 2 {
		t.Errorf(
			"Expected second NumInstance ID 2, got %d",
			id2,
		)
	}

	// Verify we can find them
	found := numbering.GetNumInstance(1)
	if found == nil {
		t.Error(
			"GetNumInstance(1) should not return nil",
		)
	}
}

func TestCreateNumberingInstance(t *testing.T) {
	numbering := NewNumbering()

	an := NewAbstractNum()
	anId := numbering.AddAbstractNum(an)

	ni := numbering.CreateNumberingInstance(anId)
	if ni == nil {
		t.Fatal(
			"CreateNumberingInstance() returned nil",
		)
	}

	if ni.AbstractNumId() != anId {
		t.Errorf(
			"Expected AbstractNumId %d, got %d",
			anId,
			ni.AbstractNumId(),
		)
	}
}

func TestParagraphNumberingHelpers(t *testing.T) {
	p := NewParagraph("Test paragraph")

	// Initially not numbered
	if p.IsNumbered() {
		t.Error(
			"New paragraph should not be numbered",
		)
	}

	// Apply numbering
	p.ApplyNumbering(1, 0)
	if !p.IsNumbered() {
		t.Error(
			"Paragraph should be numbered after ApplyNumbering",
		)
	}
	if p.NumberingId() != 1 {
		t.Errorf(
			"Expected NumberingId 1, got %d",
			p.NumberingId(),
		)
	}
	if p.NumberingLevel() != 0 {
		t.Errorf(
			"Expected NumberingLevel 0, got %d",
			p.NumberingLevel(),
		)
	}

	// Change level
	p.SetNumberingLevel(2)
	if p.NumberingLevel() != 2 {
		t.Errorf(
			"Expected NumberingLevel 2, got %d",
			p.NumberingLevel(),
		)
	}

	// Remove numbering
	p.RemoveNumbering()
	if p.IsNumbered() {
		t.Error(
			"Paragraph should not be numbered after RemoveNumbering",
		)
	}
}

func TestAbstractNumIterators(t *testing.T) {
	numbering := NewNumbering()

	an1 := NewAbstractNum()
	an2 := NewAbstractNum()
	numbering.AddAbstractNum(an1)
	numbering.AddAbstractNum(an2)

	count := 0
	for range numbering.AbstractNums() {
		count++
	}
	if count != 2 {
		t.Errorf(
			"Expected 2 abstract nums, got %d",
			count,
		)
	}

	ni1 := NewNumberingInstance(0)
	numbering.AddNumInstance(ni1)

	count = 0
	for range numbering.NumInstances() {
		count++
	}
	if count != 1 {
		t.Errorf(
			"Expected 1 num instance, got %d",
			count,
		)
	}
}
