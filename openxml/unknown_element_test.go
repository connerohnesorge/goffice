package openxml

import (
	"bytes"
	"testing"
)

func TestUnknownElement_Creation(t *testing.T) {
	rawXML := []byte(
		`<custom:element xmlns:custom="http://example.com/custom">content</custom:element>`,
	)
	ue := NewUnknownElement(
		"http://example.com/custom",
		"element",
		"custom",
		rawXML,
	)

	if ue == nil {
		t.Fatal(
			"NewUnknownElement() returned nil",
		)
	}

	if ue.LocalName() != "element" {
		t.Errorf(
			"LocalName() = %q, want %q",
			ue.LocalName(),
			"element",
		)
	}

	if ue.NamespaceURI() != "http://example.com/custom" {
		t.Errorf(
			"NamespaceURI() = %q, want %q",
			ue.NamespaceURI(),
			"http://example.com/custom",
		)
	}

	if ue.Prefix() != "custom" {
		t.Errorf(
			"Prefix() = %q, want %q",
			ue.Prefix(),
			"custom",
		)
	}
}

func TestUnknownElement_QName(t *testing.T) {
	rawXML := []byte(
		`<test:elem xmlns:test="http://test.com">data</test:elem>`,
	)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	qname := ue.QName()
	if qname.LocalName() != "elem" {
		t.Errorf(
			"QName().LocalName() = %q, want %q",
			qname.LocalName(),
			"elem",
		)
	}
	if qname.NamespaceURI() != "http://test.com" {
		t.Errorf(
			"QName().NamespaceURI() = %q, want %q",
			qname.NamespaceURI(),
			"http://test.com",
		)
	}
}

func TestUnknownElement_RawXML(t *testing.T) {
	rawXML := []byte(
		`<unknown:element xmlns:unknown="http://unknown.com">preserved content</unknown:element>`,
	)
	ue := NewUnknownElement(
		"http://unknown.com",
		"element",
		"unknown",
		rawXML,
	)

	retrievedRaw := ue.RawXML()
	if !bytes.Equal(retrievedRaw, rawXML) {
		t.Errorf(
			"RawXML() = %q, want %q",
			string(retrievedRaw),
			string(rawXML),
		)
	}
}

func TestUnknownElement_WriteXML(t *testing.T) {
	rawXML := []byte(
		`<future:feature xmlns:future="http://future-office.com/2030">
  <future:data>important</future:data>
</future:feature>`,
	)
	ue := NewUnknownElement(
		"http://future-office.com/2030",
		"feature",
		"future",
		rawXML,
	)

	var buf bytes.Buffer
	err := ue.WriteXML(&buf)
	if err != nil {
		t.Fatalf("WriteXML() error = %v", err)
	}

	written := buf.Bytes()
	if !bytes.Equal(written, rawXML) {
		t.Errorf(
			"WriteXML() wrote:\n%s\nwant:\n%s",
			string(written),
			string(rawXML),
		)
	}
}

func TestUnknownElement_OuterXml(t *testing.T) {
	rawXML := []byte(
		`<x:elem xmlns:x="http://x.com">text</x:elem>`,
	)
	ue := NewUnknownElement(
		"http://x.com",
		"elem",
		"x",
		rawXML,
	)

	outerXml := ue.OuterXml()
	if outerXml != string(rawXML) {
		t.Errorf(
			"OuterXml() = %q, want %q",
			outerXml,
			string(rawXML),
		)
	}
}

func TestUnknownElement_InnerXml(t *testing.T) {
	rawXML := []byte(
		`<test:element>inner content</test:element>`,
	)
	ue := NewUnknownElement(
		"http://test.com",
		"element",
		"test",
		rawXML,
	)

	// InnerXml should return empty for unknown elements since we don't parse them
	innerXml := ue.InnerXml()
	if innerXml != "" {
		t.Errorf(
			"InnerXml() = %q, want empty string",
			innerXml,
		)
	}
}

func TestUnknownElement_Clone(t *testing.T) {
	rawXML := []byte(
		`<original:elem xmlns:original="http://original.com">data</original:elem>`,
	)
	ue := NewUnknownElement(
		"http://original.com",
		"elem",
		"original",
		rawXML,
	)

	clone := ue.Clone()
	if clone == nil {
		t.Fatal("Clone() returned nil")
	}

	clonedUE, ok := clone.(*UnknownElement)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, want *UnknownElement",
			clone,
		)
	}

	// Check that the clone has the same properties
	if clonedUE.LocalName() != ue.LocalName() {
		t.Errorf(
			"Clone().LocalName() = %q, want %q",
			clonedUE.LocalName(),
			ue.LocalName(),
		)
	}
	if clonedUE.NamespaceURI() != ue.NamespaceURI() {
		t.Errorf(
			"Clone().NamespaceURI() = %q, want %q",
			clonedUE.NamespaceURI(),
			ue.NamespaceURI(),
		)
	}
	if clonedUE.Prefix() != ue.Prefix() {
		t.Errorf(
			"Clone().Prefix() = %q, want %q",
			clonedUE.Prefix(),
			ue.Prefix(),
		)
	}

	// Check that the raw XML is copied
	if !bytes.Equal(
		clonedUE.RawXML(),
		ue.RawXML(),
	) {
		t.Error(
			"Clone().RawXML() != original.RawXML()",
		)
	}

	// Check that modifying the clone's raw XML doesn't affect the original
	// (they should be separate byte slices)
	clonedRaw := clonedUE.RawXML()
	if len(clonedRaw) == 0 {
		return
	}

	clonedRaw[0] = 'X'
	if bytes.Equal(
		clonedUE.RawXML(),
		ue.RawXML(),
	) {
		t.Error(
			"Modifying clone's raw XML affected the original (not a deep copy)",
		)
	}
}

func TestUnknownElement_CloneNode(t *testing.T) {
	rawXML := []byte(
		`<test:node>content</test:node>`,
	)
	ue := NewUnknownElement(
		"http://test.com",
		"node",
		"test",
		rawXML,
	)

	// For unknown elements, deep parameter doesn't matter
	shallowClone := ue.CloneNode(false)
	deepClone := ue.CloneNode(true)

	if shallowClone == nil || deepClone == nil {
		t.Fatal("CloneNode() returned nil")
	}

	// Both should be equal since unknown elements are opaque
	shallowUE, ok := shallowClone.(*UnknownElement)
	if !ok {
		t.Fatal(
			"shallowClone is not *UnknownElement",
		)
	}
	deepUE, ok := deepClone.(*UnknownElement)
	if !ok {
		t.Fatal(
			"deepClone is not *UnknownElement",
		)
	}

	if !bytes.Equal(
		shallowUE.RawXML(),
		deepUE.RawXML(),
	) {
		t.Error(
			"CloneNode(false) and CloneNode(true) produced different results",
		)
	}
}

func TestUnknownElement_Parent(t *testing.T) {
	rawXML := []byte(
		`<child:elem>data</child:elem>`,
	)
	ue := NewUnknownElement(
		"http://child.com",
		"elem",
		"child",
		rawXML,
	)

	// Initially no parent
	if ue.Parent() != nil {
		t.Error(
			"Parent() should be nil initially",
		)
	}

	// Set a parent
	parent := NewCompositeElement(
		"http://parent.com",
		"parent",
		"p",
	)
	parent.AppendChild(ue)

	if ue.Parent() == nil {
		t.Error(
			"Parent() should not be nil after AppendChild()",
		)
	}
	if ue.Parent() != parent {
		t.Error("Parent() returned wrong parent")
	}
}

func TestUnknownElement_Attributes(t *testing.T) {
	rawXML := []byte(
		`<elem attr="value">content</elem>`,
	)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	// Unknown elements don't parse attributes
	attrs := ue.Attributes()
	if len(attrs) != 0 {
		t.Errorf(
			"Attributes() returned %d attributes, want 0",
			len(attrs),
		)
	}
}

func TestUnknownElement_GetAttribute(
	t *testing.T,
) {
	rawXML := []byte(
		`<elem attr="value">content</elem>`,
	)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	// Unknown elements don't support attribute access
	_, found := ue.GetAttribute("attr", "")
	if found {
		t.Error(
			"GetAttribute() returned true, want false for unknown elements",
		)
	}
}

func TestUnknownElement_SetAttribute(
	t *testing.T,
) {
	rawXML := []byte(`<elem>content</elem>`)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	// SetAttribute should be a no-op
	ue.SetAttribute(
		NewSimpleAttribute("test", "value"),
	)

	// Verify it didn't change anything
	if !bytes.Equal(ue.RawXML(), rawXML) {
		t.Error(
			"SetAttribute() modified the raw XML",
		)
	}
}

func TestUnknownElement_RemoveAttribute(
	t *testing.T,
) {
	rawXML := []byte(
		`<elem attr="value">content</elem>`,
	)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	// RemoveAttribute should return false
	removed := ue.RemoveAttribute("attr", "")
	if removed {
		t.Error(
			"RemoveAttribute() returned true, want false for unknown elements",
		)
	}

	// Verify it didn't change anything
	if !bytes.Equal(ue.RawXML(), rawXML) {
		t.Error(
			"RemoveAttribute() modified the raw XML",
		)
	}
}

func TestUnknownElement_Roundtrip(t *testing.T) {
	// Test that we can read, store, and write back complex unknown XML
	rawXML := []byte(
		`<future:complex xmlns:future="http://future.com/2030" attr1="value1" attr2="value2">
  <future:nested>
    <future:deep attr="nested-value">
      Text content here
    </future:deep>
  </future:nested>
  <future:another>More content</future:another>
</future:complex>`,
	)

	ue := NewUnknownElement(
		"http://future.com/2030",
		"complex",
		"future",
		rawXML,
	)

	// Write it back
	var buf bytes.Buffer
	err := ue.WriteXML(&buf)
	if err != nil {
		t.Fatalf("WriteXML() error = %v", err)
	}

	// Should be identical
	if !bytes.Equal(buf.Bytes(), rawXML) {
		t.Errorf(
			"Roundtrip failed:\nGot:\n%s\n\nWant:\n%s",
			buf.String(),
			string(rawXML),
		)
	}
}

func TestUnknownElement_Features(t *testing.T) {
	rawXML := []byte(`<elem>content</elem>`)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	features := ue.Features()
	if features == nil {
		t.Error("Features() returned nil")
	}
}

func TestUnknownElement_SetPrefix(t *testing.T) {
	rawXML := []byte(`<elem>content</elem>`)
	ue := NewUnknownElement(
		"http://test.com",
		"elem",
		"test",
		rawXML,
	)

	// Test SetPrefix
	ue.SetPrefix("newprefix")
	if ue.Prefix() != "newprefix" {
		t.Errorf(
			"After SetPrefix('newprefix'), Prefix() = %q, want 'newprefix'",
			ue.Prefix(),
		)
	}
}
