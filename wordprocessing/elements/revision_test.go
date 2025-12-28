package elements

import (
	"strings"
	"testing"
	"time"
)

func TestInsertedRun(t *testing.T) {
	id := 1
	author := "Test Author"
	date := time.Date(
		2024,
		1,
		15,
		10,
		30,
		0,
		0,
		time.UTC,
	)

	ins := NewInsertedRun(id, author, date)

	if ins.Id() != id {
		t.Errorf(
			"Expected ID %d, got %d",
			id,
			ins.Id(),
		)
	}

	if ins.Author() != author {
		t.Errorf(
			"Expected author '%s', got '%s'",
			author,
			ins.Author(),
		)
	}

	parsedDate := ins.Date()
	if !parsedDate.Equal(date) {
		t.Errorf(
			"Expected date %v, got %v",
			date,
			parsedDate,
		)
	}
}

func TestInsertedRunWithContent(t *testing.T) {
	ins := NewInsertedRun(
		1,
		testAuthorName,
		time.Now(),
	)

	ins.AppendRun("Inserted text")

	text := ins.InnerText()
	if text != "Inserted text" {
		t.Errorf(
			"Expected 'Inserted text', got '%s'",
			text,
		)
	}

	// Count runs
	count := 0
	for range ins.Runs() {
		count++
	}
	if count != 1 {
		t.Errorf("Expected 1 run, got %d", count)
	}
}

func TestInsertedRunSetters(t *testing.T) {
	ins := NewInsertedRun(
		1,
		testAuthorName,
		time.Now(),
	)

	ins.SetId(5)
	if ins.Id() != 5 {
		t.Errorf(
			"Expected ID 5, got %d",
			ins.Id(),
		)
	}

	ins.SetAuthor("New Author")
	if ins.Author() != "New Author" {
		t.Errorf(
			"Expected 'New Author', got '%s'",
			ins.Author(),
		)
	}

	newDate := time.Date(
		2024,
		6,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	ins.SetDate(newDate)
	if !ins.Date().Equal(newDate) {
		t.Errorf(
			"Expected date %v, got %v",
			newDate,
			ins.Date(),
		)
	}
}

func TestDeletedRun(t *testing.T) {
	id := 2
	author := "Test Author"
	date := time.Date(
		2024,
		1,
		15,
		10,
		30,
		0,
		0,
		time.UTC,
	)

	del := NewDeletedRun(id, author, date)

	if del.Id() != id {
		t.Errorf(
			"Expected ID %d, got %d",
			id,
			del.Id(),
		)
	}

	if del.Author() != author {
		t.Errorf(
			"Expected author '%s', got '%s'",
			author,
			del.Author(),
		)
	}

	parsedDate := del.Date()
	if !parsedDate.Equal(date) {
		t.Errorf(
			"Expected date %v, got %v",
			date,
			parsedDate,
		)
	}
}

func TestDeletedRunWithContent(t *testing.T) {
	del := NewDeletedRun(
		1,
		testAuthorName,
		time.Now(),
	)

	del.AppendDeletedRun("Deleted text")

	text := del.InnerText()
	if text != "Deleted text" {
		t.Errorf(
			"Expected 'Deleted text', got '%s'",
			text,
		)
	}

	// Count deleted texts
	count := 0
	for range del.DeletedTexts() {
		count++
	}
	if count != 1 {
		t.Errorf(
			"Expected 1 deleted text, got %d",
			count,
		)
	}
}

func TestDeletedText(t *testing.T) {
	dt := NewDeletedText("Some deleted text")

	if dt.InnerText() != "Some deleted text" {
		t.Errorf(
			"Expected 'Some deleted text', got '%s'",
			dt.InnerText(),
		)
	}

	dt.SetText("New deleted text")
	if dt.InnerText() != "New deleted text" {
		t.Errorf(
			"Expected 'New deleted text', got '%s'",
			dt.InnerText(),
		)
	}
}

func TestDeletedTextSpacePreserve(t *testing.T) {
	dt := NewDeletedText("  spaces  ")

	if dt.Space() != "preserve" {
		t.Error(
			"Expected space='preserve' for text with leading/trailing spaces",
		)
	}
}

func TestMoveFromRun(t *testing.T) {
	mf := NewMoveFromRun(
		1,
		testAuthorName,
		time.Now(),
	)

	if mf.Id() != 1 {
		t.Errorf("Expected ID 1, got %d", mf.Id())
	}

	if mf.Author() != testAuthorName {
		t.Errorf(
			"Expected author 'Author', got '%s'",
			mf.Author(),
		)
	}
}

func TestMoveToRun(t *testing.T) {
	mt := NewMoveToRun(
		1,
		testAuthorName,
		time.Now(),
	)

	if mt.Id() != 1 {
		t.Errorf("Expected ID 1, got %d", mt.Id())
	}

	if mt.Author() != testAuthorName {
		t.Errorf(
			"Expected author 'Author', got '%s'",
			mt.Author(),
		)
	}
}

func TestRunPropertiesChange(t *testing.T) {
	rpc := NewRunPropertiesChange(
		1,
		testAuthorName,
		time.Now(),
	)

	if rpc.Id() != 1 {
		t.Errorf(
			"Expected ID 1, got %d",
			rpc.Id(),
		)
	}

	if rpc.Author() != testAuthorName {
		t.Errorf(
			"Expected author 'Author', got '%s'",
			rpc.Author(),
		)
	}

	// Test previous properties
	rp := NewRunProperties()
	rp.SetBold(true)
	rpc.SetPreviousRunProperties(rp)

	prevRp := rpc.PreviousRunProperties()
	if prevRp == nil {
		t.Fatal(
			"PreviousRunProperties returned nil",
		)
	}

	xml := prevRp.OuterXml()
	if !strings.Contains(xml, "b") {
		t.Error(
			"Expected bold property in previous run properties",
		)
	}
}

func TestParagraphPropertiesChange(t *testing.T) {
	ppc := NewParagraphPropertiesChange(
		1,
		testAuthorName,
		time.Now(),
	)

	if ppc.Id() != 1 {
		t.Errorf(
			"Expected ID 1, got %d",
			ppc.Id(),
		)
	}

	if ppc.Author() != testAuthorName {
		t.Errorf(
			"Expected author 'Author', got '%s'",
			ppc.Author(),
		)
	}

	// Test previous properties
	pp := NewParagraphProperties()
	pp.SetJustification(JustificationCenter)
	ppc.SetPreviousParagraphProperties(pp)

	prevPp := ppc.PreviousParagraphProperties()
	if prevPp == nil {
		t.Fatal(
			"PreviousParagraphProperties returned nil",
		)
	}

	xml := prevPp.OuterXml()
	if !strings.Contains(xml, "jc") {
		t.Error(
			"Expected justification in previous paragraph properties",
		)
	}
}

func TestInsertedRunClone(t *testing.T) {
	ins := NewInsertedRun(
		1,
		testAuthorName,
		time.Now(),
	)
	ins.AppendRun("Text")

	cloned := ins.Clone()
	clone, ok := cloned.(*InsertedRun)
	if !ok {
		t.Fatal(
			"Clone did not return *InsertedRun",
		)
	}

	if clone.Id() != ins.Id() {
		t.Error("Cloned ID doesn't match")
	}
	if clone.Author() != ins.Author() {
		t.Error("Cloned author doesn't match")
	}
	if clone.InnerText() != ins.InnerText() {
		t.Error("Cloned text doesn't match")
	}
}

func TestDeletedRunClone(t *testing.T) {
	del := NewDeletedRun(
		1,
		testAuthorName,
		time.Now(),
	)
	del.AppendDeletedRun("Text")

	cloned := del.Clone()
	clone, ok := cloned.(*DeletedRun)
	if !ok {
		t.Fatal(
			"Clone did not return *DeletedRun",
		)
	}

	if clone.Id() != del.Id() {
		t.Error("Cloned ID doesn't match")
	}
	if clone.Author() != del.Author() {
		t.Error("Cloned author doesn't match")
	}
}
