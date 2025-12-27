package elements

//revive:disable:file-length-limit many cache record types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PivotCacheRecords represents the pivot cache records root element
// (x:pivotCacheRecords).
// This element is the root of a pivot cache records part and contains
// the actual data records for a pivot cache.
type PivotCacheRecords struct {
	*openxml.PartRootElementBase
}

// NewPivotCacheRecords creates a new PivotCacheRecords element.
func NewPivotCacheRecords() *PivotCacheRecords {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"pivotCacheRecords",
		PrefixDefault,
	)

	return &PivotCacheRecords{
		PartRootElementBase: elem,
	}
}

// Count returns the count of records. Attribute: count.
func (p *PivotCacheRecords) Count() uint32 {
	attr, found := p.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetCount sets the count of records. Attribute: count.
func (p *PivotCacheRecords) SetCount(
	count uint32,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Records returns an iterator over all R (record) elements.
func (p *PivotCacheRecords) Records() iter.Seq[*R] {
	return func(yield func(*R) bool) {
		for child := range p.Children() {
			if child.LocalName() != "r" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var record *R
			if r, ok := child.(*R); ok {
				record = r
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				record = &R{CompositeElementBase: comp}
			}
			if record != nil && !yield(record) {
				return
			}
		}
	}
}

// RecordCount returns the count of record elements.
func (p *PivotCacheRecords) RecordCount() int {
	count := 0
	for range p.Records() {
		count++
	}

	return count
}

// GetRecord returns the record at the given index, or nil if out of range.
func (p *PivotCacheRecords) GetRecord(
	index int,
) *R {
	if index < 0 {
		return nil
	}
	i := 0
	for record := range p.Records() {
		if i == index {
			return record
		}
		i++
	}

	return nil
}

// AddRecord adds a new R (record) element.
func (p *PivotCacheRecords) AddRecord() *R {
	record := NewR()
	p.AppendChild(record)

	return record
}

// Clone creates a deep copy of this PivotCacheRecords element.
func (p *PivotCacheRecords) Clone() openxml.Element {
	cloned := p.PartRootElementBase.Clone()

	return &PivotCacheRecords{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this PivotCacheRecords element.
func (p *PivotCacheRecords) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.PartRootElementBase.CloneNode(
		deep,
	)

	return &PivotCacheRecords{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// R represents a single record element (x:r).
type R struct {
	*openxml.CompositeElementBase
}

// NewR creates a new R (record) element.
func NewR() *R {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"r",
		PrefixDefault,
	)

	return &R{CompositeElementBase: elem}
}

// AddString adds a string value to the record.
func (r *R) AddString(
	value string,
) *RecordString {
	item := NewRecordString()
	item.SetV(value)
	r.AppendChild(item)

	return item
}

// AddNumber adds a number value to the record.
func (r *R) AddNumber(
	value float64,
) *RecordNumber {
	item := NewRecordNumber()
	item.SetV(value)
	r.AppendChild(item)

	return item
}

// AddBoolean adds a boolean value to the record.
func (r *R) AddBoolean(
	value bool,
) *RecordBoolean {
	item := NewRecordBoolean()
	item.SetV(value)
	r.AppendChild(item)

	return item
}

// AddMissing adds a missing value to the record.
func (r *R) AddMissing() *RecordMissing {
	item := NewRecordMissing()
	r.AppendChild(item)

	return item
}

// AddError adds an error value to the record.
func (r *R) AddError(value string) *RecordError {
	item := NewRecordError()
	item.SetV(value)
	r.AppendChild(item)

	return item
}

// AddDateTime adds a date/time value to the record.
func (r *R) AddDateTime(
	value string,
) *RecordDateTime {
	item := NewRecordDateTime()
	item.SetV(value)
	r.AppendChild(item)

	return item
}

// AddIndex adds a shared item index to the record.
func (r *R) AddIndex(value uint32) *RecordIndex {
	item := NewRecordIndex()
	item.SetV(value)
	r.AppendChild(item)

	return item
}

// Clone creates a deep copy of this R element.
func (r *R) Clone() openxml.Element {
	cloned := r.CompositeElementBase.Clone()

	return &R{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this R element.
func (r *R) CloneNode(deep bool) openxml.Element {
	cloned := r.CompositeElementBase.CloneNode(
		deep,
	)

	return &R{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// RecordString represents a string value in a record (x:s).
type RecordString struct {
	*openxml.LeafElementBase
}

// NewRecordString creates a new RecordString element.
func NewRecordString() *RecordString {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"s",
		PrefixDefault,
	)

	return &RecordString{LeafElementBase: elem}
}

// V returns the string value. Attribute: v.
func (rs *RecordString) V() string {
	attr, found := rs.GetAttribute("v", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetV sets the string value. Attribute: v.
func (rs *RecordString) SetV(value string) {
	rs.SetAttribute(
		openxml.NewAttribute("", "v", "", value),
	)
}

// Clone creates a deep copy of this RecordString element.
func (rs *RecordString) Clone() openxml.Element {
	cloned := rs.LeafElementBase.Clone()

	return &RecordString{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordString element.
func (rs *RecordString) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rs.LeafElementBase.CloneNode(deep)

	return &RecordString{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RecordNumber represents a number value in a record (x:n).
type RecordNumber struct {
	*openxml.LeafElementBase
}

// NewRecordNumber creates a new RecordNumber element.
func NewRecordNumber() *RecordNumber {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"n",
		PrefixDefault,
	)

	return &RecordNumber{LeafElementBase: elem}
}

// V returns the numeric value. Attribute: v.
func (rn *RecordNumber) V() float64 {
	attr, found := rn.GetAttribute("v", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetV sets the numeric value. Attribute: v.
func (rn *RecordNumber) SetV(value float64) {
	rn.SetAttribute(
		openxml.NewAttribute(
			"",
			"v", //nolint:revive // add-constant: attribute name
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// Clone creates a deep copy of this RecordNumber element.
func (rn *RecordNumber) Clone() openxml.Element {
	cloned := rn.LeafElementBase.Clone()

	return &RecordNumber{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordNumber element.
func (rn *RecordNumber) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rn.LeafElementBase.CloneNode(deep)

	return &RecordNumber{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RecordBoolean represents a boolean value in a record (x:b).
type RecordBoolean struct {
	*openxml.LeafElementBase
}

// NewRecordBoolean creates a new RecordBoolean element.
func NewRecordBoolean() *RecordBoolean {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"b",
		PrefixDefault,
	)

	return &RecordBoolean{LeafElementBase: elem}
}

// V returns the boolean value. Attribute: v.
func (rb *RecordBoolean) V() bool {
	attr, found := rb.GetAttribute("v", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetV sets the boolean value. Attribute: v.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (rb *RecordBoolean) SetV(value bool) {
	if value {
		rb.SetAttribute(
			openxml.NewAttribute(
				"",
				"v",
				"",
				attrValueTrue,
			),
		)
	} else {
		rb.SetAttribute(
			openxml.NewAttribute("", "v", "", attrValueFalse),
		)
	}
}

// Clone creates a deep copy of this RecordBoolean element.
func (rb *RecordBoolean) Clone() openxml.Element {
	cloned := rb.LeafElementBase.Clone()

	return &RecordBoolean{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordBoolean element.
func (rb *RecordBoolean) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rb.LeafElementBase.CloneNode(deep)

	return &RecordBoolean{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RecordMissing represents a missing value in a record (x:m).
type RecordMissing struct {
	*openxml.LeafElementBase
}

// NewRecordMissing creates a new RecordMissing element.
func NewRecordMissing() *RecordMissing {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"m",
		PrefixDefault,
	)

	return &RecordMissing{LeafElementBase: elem}
}

// Clone creates a deep copy of this RecordMissing element.
func (rm *RecordMissing) Clone() openxml.Element {
	cloned := rm.LeafElementBase.Clone()

	return &RecordMissing{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordMissing element.
func (rm *RecordMissing) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rm.LeafElementBase.CloneNode(deep)

	return &RecordMissing{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RecordError represents an error value in a record (x:e).
type RecordError struct {
	*openxml.LeafElementBase
}

// NewRecordError creates a new RecordError element.
func NewRecordError() *RecordError {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"e",
		PrefixDefault,
	)

	return &RecordError{LeafElementBase: elem}
}

// V returns the error value. Attribute: v.
func (re *RecordError) V() string {
	attr, found := re.GetAttribute("v", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetV sets the error value. Attribute: v.
func (re *RecordError) SetV(value string) {
	re.SetAttribute(
		openxml.NewAttribute("", "v", "", value),
	)
}

// Clone creates a deep copy of this RecordError element.
func (re *RecordError) Clone() openxml.Element {
	cloned := re.LeafElementBase.Clone()

	return &RecordError{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordError element.
func (re *RecordError) CloneNode(
	deep bool,
) openxml.Element {
	cloned := re.LeafElementBase.CloneNode(deep)

	return &RecordError{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RecordDateTime represents a date/time value in a record (x:d).
type RecordDateTime struct {
	*openxml.LeafElementBase
}

// NewRecordDateTime creates a new RecordDateTime element.
func NewRecordDateTime() *RecordDateTime {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"d",
		PrefixDefault,
	)

	return &RecordDateTime{LeafElementBase: elem}
}

// V returns the date/time value (ISO 8601 format). Attribute: v.
func (rd *RecordDateTime) V() string {
	attr, found := rd.GetAttribute("v", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetV sets the date/time value (ISO 8601 format). Attribute: v.
func (rd *RecordDateTime) SetV(value string) {
	rd.SetAttribute(
		openxml.NewAttribute("", "v", "", value),
	)
}

// Clone creates a deep copy of this RecordDateTime element.
func (rd *RecordDateTime) Clone() openxml.Element {
	cloned := rd.LeafElementBase.Clone()

	return &RecordDateTime{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordDateTime element.
func (rd *RecordDateTime) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rd.LeafElementBase.CloneNode(deep)

	return &RecordDateTime{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RecordIndex represents a shared item index in a record (x:x).
type RecordIndex struct {
	*openxml.LeafElementBase
}

// NewRecordIndex creates a new RecordIndex element.
func NewRecordIndex() *RecordIndex {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"x",
		PrefixDefault,
	)

	return &RecordIndex{LeafElementBase: elem}
}

// V returns the index value. Attribute: v.
func (ri *RecordIndex) V() uint32 {
	attr, found := ri.GetAttribute("v", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetV sets the index value. Attribute: v.
func (ri *RecordIndex) SetV(value uint32) {
	ri.SetAttribute(
		openxml.NewAttribute(
			"",
			"v",
			"",
			strconv.FormatUint(
				uint64(value),
				parseBase10,
			),
		),
	)
}

// Clone creates a deep copy of this RecordIndex element.
func (ri *RecordIndex) Clone() openxml.Element {
	cloned := ri.LeafElementBase.Clone()

	return &RecordIndex{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RecordIndex element.
func (ri *RecordIndex) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ri.LeafElementBase.CloneNode(deep)

	return &RecordIndex{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
