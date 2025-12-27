package elements

import (
	"testing"
)

func TestNewPivotCacheRecords(t *testing.T) {
	pcr := NewPivotCacheRecords()

	if pcr == nil {
		t.Fatal(
			"NewPivotCacheRecords returned nil",
		)
	}

	if pcr.LocalName() != "pivotCacheRecords" {
		t.Errorf(
			"Expected LocalName 'pivotCacheRecords', got '%s'",
			pcr.LocalName(),
		)
	}

	if pcr.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"Expected NamespaceURI '%s', got '%s'",
			NamespaceSML,
			pcr.NamespaceURI(),
		)
	}
}

func TestPivotCacheRecordsCount(t *testing.T) {
	pcr := NewPivotCacheRecords()

	// Initially 0
	if pcr.Count() != 0 {
		t.Errorf(
			"Expected Count 0, got %d",
			pcr.Count(),
		)
	}

	// Set count
	pcr.SetCount(100)
	if pcr.Count() != 100 {
		t.Errorf(
			"Expected Count 100, got %d",
			pcr.Count(),
		)
	}
}

func TestPivotCacheRecordsAddRecord(
	t *testing.T,
) {
	pcr := NewPivotCacheRecords()

	// Add a record
	r := pcr.AddRecord()
	if r == nil {
		t.Fatal("AddRecord returned nil")
	}

	if r.LocalName() != "r" {
		t.Errorf(
			"Expected LocalName 'r', got '%s'",
			r.LocalName(),
		)
	}

	// Check record count
	if pcr.RecordCount() != 1 {
		t.Errorf(
			"Expected RecordCount 1, got %d",
			pcr.RecordCount(),
		)
	}
}

func TestPivotCacheRecordsIteration(
	t *testing.T,
) {
	pcr := NewPivotCacheRecords()

	// Add records
	pcr.AddRecord()
	pcr.AddRecord()
	pcr.AddRecord()

	// Iterate
	count := 0
	for r := range pcr.Records() {
		count++
		if r == nil {
			t.Error("Iterator yielded nil record")
		}
	}
	if count != 3 {
		t.Errorf(
			"Expected 3 records in iterator, got %d",
			count,
		)
	}
}

func TestPivotCacheRecordsGetRecord(
	t *testing.T,
) {
	pcr := NewPivotCacheRecords()

	// Add records with data
	r0 := pcr.AddRecord()
	r0.AddString("first")
	r1 := pcr.AddRecord()
	r1.AddString("second")
	r2 := pcr.AddRecord()
	r2.AddString("third")

	// Get by index
	rec := pcr.GetRecord(1)
	if rec == nil {
		t.Fatal("GetRecord(1) returned nil")
	}

	// Invalid index
	if pcr.GetRecord(-1) != nil {
		t.Error("Expected nil for negative index")
	}

	if pcr.GetRecord(99) != nil {
		t.Error(
			"Expected nil for out of range index",
		)
	}
}

func TestRecordAddString(t *testing.T) {
	r := NewR()

	s := r.AddString("TestValue")
	if s == nil {
		t.Fatal("AddString returned nil")
	}

	if s.V() != "TestValue" {
		t.Errorf(
			"Expected V 'TestValue', got '%s'",
			s.V(),
		)
	}

	if s.LocalName() != "s" {
		t.Errorf(
			"Expected LocalName 's', got '%s'",
			s.LocalName(),
		)
	}
}

func TestRecordAddNumber(t *testing.T) {
	r := NewR()

	n := r.AddNumber(123.456)
	if n == nil {
		t.Fatal("AddNumber returned nil")
	}

	if n.V() != 123.456 {
		t.Errorf(
			"Expected V 123.456, got %f",
			n.V(),
		)
	}

	if n.LocalName() != "n" {
		t.Errorf(
			"Expected LocalName 'n', got '%s'",
			n.LocalName(),
		)
	}
}

func TestRecordAddBoolean(t *testing.T) {
	r := NewR()

	b := r.AddBoolean(true)
	if b == nil {
		t.Fatal("AddBoolean returned nil")
	}

	if !b.V() {
		t.Error("Expected V true")
	}

	if b.LocalName() != "b" {
		t.Errorf(
			"Expected LocalName 'b', got '%s'",
			b.LocalName(),
		)
	}

	// Test false
	b2 := r.AddBoolean(false)
	if b2.V() {
		t.Error("Expected V false")
	}
}

func TestRecordAddMissing(t *testing.T) {
	r := NewR()

	m := r.AddMissing()
	if m == nil {
		t.Fatal("AddMissing returned nil")
	}

	if m.LocalName() != "m" {
		t.Errorf(
			"Expected LocalName 'm', got '%s'",
			m.LocalName(),
		)
	}
}

func TestRecordAddError(t *testing.T) {
	r := NewR()

	e := r.AddError("#DIV/0!")
	if e == nil {
		t.Fatal("AddError returned nil")
	}

	if e.V() != "#DIV/0!" {
		t.Errorf(
			"Expected V '#DIV/0!', got '%s'",
			e.V(),
		)
	}

	if e.LocalName() != "e" {
		t.Errorf(
			"Expected LocalName 'e', got '%s'",
			e.LocalName(),
		)
	}
}

func TestRecordAddDateTime(t *testing.T) {
	r := NewR()

	d := r.AddDateTime("2023-12-20T15:30:00")
	if d == nil {
		t.Fatal("AddDateTime returned nil")
	}

	if d.V() != "2023-12-20T15:30:00" {
		t.Errorf(
			"Expected V '2023-12-20T15:30:00', got '%s'",
			d.V(),
		)
	}

	if d.LocalName() != "d" {
		t.Errorf(
			"Expected LocalName 'd', got '%s'",
			d.LocalName(),
		)
	}
}

func TestRecordAddIndex(t *testing.T) {
	r := NewR()

	x := r.AddIndex(5)
	if x == nil {
		t.Fatal("AddIndex returned nil")
	}

	if x.V() != 5 {
		t.Errorf("Expected V 5, got %d", x.V())
	}

	if x.LocalName() != "x" {
		t.Errorf(
			"Expected LocalName 'x', got '%s'",
			x.LocalName(),
		)
	}
}

func TestRecordMixedValues(t *testing.T) {
	r := NewR()

	// Add mixed value types
	r.AddString("Hello")
	r.AddNumber(42)
	r.AddBoolean(true)
	r.AddMissing()
	r.AddIndex(0)

	// Verify the record has 5 children
	count := 0
	for range r.Children() {
		count++
	}
	if count != 5 {
		t.Errorf(
			"Expected 5 children, got %d",
			count,
		)
	}
}

func TestRecordString(t *testing.T) {
	rs := NewRecordString()

	if rs == nil {
		t.Fatal("NewRecordString returned nil")
	}

	// Test V
	rs.SetV("test string")
	if rs.V() != "test string" {
		t.Errorf(
			"Expected V 'test string', got '%s'",
			rs.V(),
		)
	}
}

func TestRecordNumber(t *testing.T) {
	rn := NewRecordNumber()

	if rn == nil {
		t.Fatal("NewRecordNumber returned nil")
	}

	// Test V
	rn.SetV(99.99)
	if rn.V() != 99.99 {
		t.Errorf(
			"Expected V 99.99, got %f",
			rn.V(),
		)
	}

	// Test zero
	rn.SetV(0)
	if rn.V() != 0 {
		t.Errorf("Expected V 0, got %f", rn.V())
	}

	// Test negative
	rn.SetV(-123.45)
	if rn.V() != -123.45 {
		t.Errorf(
			"Expected V -123.45, got %f",
			rn.V(),
		)
	}
}

func TestRecordBoolean(t *testing.T) {
	rb := NewRecordBoolean()

	if rb == nil {
		t.Fatal("NewRecordBoolean returned nil")
	}

	// Test true
	rb.SetV(true)
	if !rb.V() {
		t.Error("Expected V true")
	}

	// Test false
	rb.SetV(false)
	if rb.V() {
		t.Error("Expected V false")
	}
}

func TestRecordMissing(t *testing.T) {
	rm := NewRecordMissing()

	if rm == nil {
		t.Fatal("NewRecordMissing returned nil")
	}

	if rm.LocalName() != "m" {
		t.Errorf(
			"Expected LocalName 'm', got '%s'",
			rm.LocalName(),
		)
	}
}

func TestRecordError(t *testing.T) {
	re := NewRecordError()

	if re == nil {
		t.Fatal("NewRecordError returned nil")
	}

	// Test V
	re.SetV("#VALUE!")
	if re.V() != "#VALUE!" {
		t.Errorf(
			"Expected V '#VALUE!', got '%s'",
			re.V(),
		)
	}
}

func TestRecordDateTime(t *testing.T) {
	rd := NewRecordDateTime()

	if rd == nil {
		t.Fatal("NewRecordDateTime returned nil")
	}

	// Test V
	rd.SetV("2023-06-15T12:00:00")
	if rd.V() != "2023-06-15T12:00:00" {
		t.Errorf(
			"Expected V '2023-06-15T12:00:00', got '%s'",
			rd.V(),
		)
	}
}

func TestRecordIndex(t *testing.T) {
	ri := NewRecordIndex()

	if ri == nil {
		t.Fatal("NewRecordIndex returned nil")
	}

	// Test V
	ri.SetV(10)
	if ri.V() != 10 {
		t.Errorf("Expected V 10, got %d", ri.V())
	}

	// Test zero
	ri.SetV(0)
	if ri.V() != 0 {
		t.Errorf("Expected V 0, got %d", ri.V())
	}
}

func TestPivotCacheRecordsClone(t *testing.T) {
	pcr := NewPivotCacheRecords()
	pcr.SetCount(10)
	r := pcr.AddRecord()
	r.AddString("test")

	clonedResult := pcr.Clone()
	cloned, ok := clonedResult.(*PivotCacheRecords)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *PivotCacheRecords",
			clonedResult,
		)
	}

	if cloned == pcr {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Count() != 10 {
		t.Errorf(
			"Cloned Count should be 10, got %d",
			cloned.Count(),
		)
	}

	if cloned.RecordCount() != 1 {
		t.Errorf(
			"Cloned RecordCount should be 1, got %d",
			cloned.RecordCount(),
		)
	}
}

func TestRecordClone(t *testing.T) {
	r := NewR()
	r.AddString("value1")
	r.AddNumber(123)

	clonedResult := r.Clone()
	cloned, ok := clonedResult.(*R)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *R",
			clonedResult,
		)
	}

	if cloned == r {
		t.Error(
			"Clone should return a new instance",
		)
	}

	// Verify children count
	origCount := 0
	for range r.Children() {
		origCount++
	}

	clonedCount := 0
	for range cloned.Children() {
		clonedCount++
	}

	if origCount != clonedCount {
		t.Errorf(
			"Expected %d children in clone, got %d",
			origCount,
			clonedCount,
		)
	}
}

//nolint:revive // unchecked-type-assertion: test code with controlled types
func TestRecordElementClones(t *testing.T) {
	// Test RecordString clone
	rs := NewRecordString()
	rs.SetV("test")
	rsCloned := rs.Clone().(*RecordString)
	if rsCloned.V() != "test" {
		t.Error("RecordString clone failed")
	}

	// Test RecordNumber clone
	rn := NewRecordNumber()
	rn.SetV(42.5)
	rnCloned := rn.Clone().(*RecordNumber)
	if rnCloned.V() != 42.5 {
		t.Error("RecordNumber clone failed")
	}

	// Test RecordBoolean clone
	rb := NewRecordBoolean()
	rb.SetV(true)
	rbCloned := rb.Clone().(*RecordBoolean)
	if !rbCloned.V() {
		t.Error("RecordBoolean clone failed")
	}

	// Test RecordMissing clone
	rm := NewRecordMissing()
	rmCloned := rm.Clone().(*RecordMissing)
	if rmCloned.LocalName() != "m" {
		t.Error("RecordMissing clone failed")
	}

	// Test RecordError clone
	re := NewRecordError()
	re.SetV("#N/A")
	reCloned := re.Clone().(*RecordError)
	if reCloned.V() != "#N/A" {
		t.Error("RecordError clone failed")
	}

	// Test RecordDateTime clone
	rd := NewRecordDateTime()
	rd.SetV("2023-01-01")
	rdCloned := rd.Clone().(*RecordDateTime)
	if rdCloned.V() != "2023-01-01" {
		t.Error("RecordDateTime clone failed")
	}

	// Test RecordIndex clone
	ri := NewRecordIndex()
	ri.SetV(5)
	riCloned := ri.Clone().(*RecordIndex)
	if riCloned.V() != 5 {
		t.Error("RecordIndex clone failed")
	}
}
