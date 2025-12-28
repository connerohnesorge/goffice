package elements

//revive:disable:file-length-limit many slicer types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Slicer namespace constants.
const (
	// NamespaceSlicerX14 is the Excel 2010 slicer namespace.
	//nolint:revive // line-length-limit: namespace URL
	NamespaceSlicerX14 = "http://schemas.microsoft.com/office/spreadsheetml/2009/9/main"

	// NamespaceSlicerX15 is the Excel 2013+ slicer namespace.
	//nolint:revive // line-length-limit: namespace URL
	NamespaceSlicerX15 = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"
)

// Slicers represents the slicers collection root element (x14:slicers).
// This element is the root of a slicers part.
type Slicers struct {
	*openxml.PartRootElementBase
}

// NewSlicers creates a new Slicers element.
func NewSlicers() *Slicers {
	elem := openxml.NewPartRootElement(
		NamespaceSlicerX14,
		"slicers",
		"x14",
	)

	return &Slicers{PartRootElementBase: elem}
}

// Slicers returns an iterator over all Slicer elements.
func (s *Slicers) Slicers() iter.Seq[*Slicer] {
	return func(yield func(*Slicer) bool) {
		for child := range s.Children() {
			if child.LocalName() != "slicer" {
				continue
			}
			var slicer *Slicer
			switch v := child.(type) {
			case *Slicer:
				slicer = v
			case *openxml.LeafElementBase:
				slicer = &Slicer{LeafElementBase: v}
			}
			if slicer != nil && !yield(slicer) {
				return
			}
		}
	}
}

// SlicerCount returns the count of slicer elements.
func (s *Slicers) SlicerCount() int {
	count := 0
	for range s.Slicers() {
		count++
	}

	return count
}

// GetSlicerByName returns the slicer with the given name, or nil if not found.
func (s *Slicers) GetSlicerByName(
	name string,
) *Slicer {
	for slicer := range s.Slicers() {
		if slicer.Name() == name {
			return slicer
		}
	}

	return nil
}

// AddSlicer adds a new Slicer element.
func (s *Slicers) AddSlicer(name string) *Slicer {
	slicer := NewSlicer()
	slicer.SetName(name)
	s.AppendChild(slicer)

	return slicer
}

// RemoveSlicer removes a slicer from the collection.
func (s *Slicers) RemoveSlicer(
	slicer *Slicer,
) bool {
	return s.RemoveChild(slicer)
}

// Clone creates a deep copy of this Slicers element.
func (s *Slicers) Clone() openxml.Element {
	cloned := s.PartRootElementBase.Clone()

	return &Slicers{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Slicers element.
func (s *Slicers) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.PartRootElementBase.CloneNode(
		deep,
	)

	return &Slicers{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// Slicer represents a single slicer element (x14:slicer).
type Slicer struct {
	*openxml.LeafElementBase
}

// NewSlicer creates a new Slicer element.
func NewSlicer() *Slicer {
	elem := openxml.NewLeafElement(
		NamespaceSlicerX14,
		"slicer",
		"x14",
	)

	return &Slicer{LeafElementBase: elem}
}

// Name returns the slicer name. Attribute: name.
func (s *Slicer) Name() string {
	attr, found := s.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the slicer name. Attribute: name.
func (s *Slicer) SetName(name string) {
	if name == "" {
		s.RemoveAttribute("name", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Cache returns the cache name. Attribute: cache.
func (s *Slicer) Cache() string {
	attr, found := s.GetAttribute("cache", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCache sets the cache name. Attribute: cache.
func (s *Slicer) SetCache(cache string) {
	if cache == "" {
		s.RemoveAttribute("cache", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"cache",
			"",
			cache,
		),
	)
}

// Caption returns the slicer caption. Attribute: caption.
func (s *Slicer) Caption() string {
	attr, found := s.GetAttribute("caption", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCaption sets the slicer caption. Attribute: caption.
func (s *Slicer) SetCaption(caption string) {
	if caption == "" {
		s.RemoveAttribute("caption", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"caption",
			"",
			caption,
		),
	)
}

// StartItem returns the start item index. Attribute: startItem.
func (s *Slicer) StartItem() uint32 {
	attr, found := s.GetAttribute("startItem", "")
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

// SetStartItem sets the start item index. Attribute: startItem.
func (s *Slicer) SetStartItem(index uint32) {
	if index == 0 {
		s.RemoveAttribute("startItem", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"startItem",
			"",
			strconv.FormatUint(
				uint64(index),
				parseBase10,
			),
		),
	)
}

// ColumnCount returns the column count. Attribute: columnCount.
func (s *Slicer) ColumnCount() uint32 {
	attr, found := s.GetAttribute(
		"columnCount",
		"",
	)
	if !found {
		return 1 // Default is 1
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetColumnCount sets the column count. Attribute: columnCount.
func (s *Slicer) SetColumnCount(count uint32) {
	if count == 1 {
		s.RemoveAttribute(
			"columnCount",
			"",
		) // 1 is default

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"columnCount",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// ShowCaption returns whether to show caption. Attribute: showCaption.
func (s *Slicer) ShowCaption() bool {
	attr, found := s.GetAttribute(
		"showCaption",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowCaption sets whether to show caption. Attribute: showCaption.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *Slicer) SetShowCaption(value bool) {
	if value {
		s.RemoveAttribute(
			"showCaption",
			"",
		) // true is default
	} else {
		s.SetAttribute(
			openxml.NewAttribute("", "showCaption", "", attrValueFalse),
		)
	}
}

// Level returns the OLAP level. Attribute: level.
func (s *Slicer) Level() uint32 {
	attr, found := s.GetAttribute("level", "")
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

// SetLevel sets the OLAP level. Attribute: level.
func (s *Slicer) SetLevel(level uint32) {
	if level == 0 {
		s.RemoveAttribute("level", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"level",
			"",
			strconv.FormatUint(
				uint64(level),
				parseBase10,
			),
		),
	)
}

// Style returns the slicer style name. Attribute: style.
func (s *Slicer) Style() string {
	attr, found := s.GetAttribute("style", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetStyle sets the slicer style name. Attribute: style.
func (s *Slicer) SetStyle(style string) {
	if style == "" {
		s.RemoveAttribute("style", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"style",
			"",
			style,
		),
	)
}

// LockedPosition returns whether position is locked. Attribute: lockedPosition.
func (s *Slicer) LockedPosition() bool {
	attr, found := s.GetAttribute(
		"lockedPosition",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetLockedPosition sets whether position is locked. Attribute: lockedPosition.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *Slicer) SetLockedPosition(value bool) {
	if value {
		s.SetAttribute(
			openxml.NewAttribute(
				"",
				"lockedPosition",
				"",
				attrValueTrue,
			),
		)
	} else {
		s.RemoveAttribute("lockedPosition", "")
	}
}

// RowHeight returns the row height in EMUs. Attribute: rowHeight.
func (s *Slicer) RowHeight() uint32 {
	attr, found := s.GetAttribute("rowHeight", "")
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

// SetRowHeight sets the row height in EMUs. Attribute: rowHeight.
func (s *Slicer) SetRowHeight(height uint32) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"rowHeight",
			"",
			strconv.FormatUint(
				uint64(height),
				parseBase10,
			),
		),
	)
}

// Clone creates a deep copy of this Slicer element.
func (s *Slicer) Clone() openxml.Element {
	cloned := s.LeafElementBase.Clone()

	return &Slicer{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Slicer element.
func (s *Slicer) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &Slicer{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// SlicerCacheDefinition represents the slicer cache definition root element
// (x14:slicerCacheDefinition).
type SlicerCacheDefinition struct {
	*openxml.PartRootElementBase
}

// NewSlicerCacheDefinition creates a new SlicerCacheDefinition element.
func NewSlicerCacheDefinition() *SlicerCacheDefinition {
	elem := openxml.NewPartRootElement(
		NamespaceSlicerX14,
		"slicerCacheDefinition",
		"x14",
	)

	return &SlicerCacheDefinition{
		PartRootElementBase: elem,
	}
}

// Name returns the cache name. Attribute: name.
func (s *SlicerCacheDefinition) Name() string {
	attr, found := s.GetAttribute(
		"name", //nolint:revive // add-constant
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the cache name. Attribute: name.
func (s *SlicerCacheDefinition) SetName(
	name string,
) {
	if name == "" {
		s.RemoveAttribute("name", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"name", //nolint:revive // add-constant: attribute name
			"",
			name,
		),
	)
}

// SourceName returns the source name. Attribute: sourceName.
func (s *SlicerCacheDefinition) SourceName() string {
	attr, found := s.GetAttribute(
		"sourceName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSourceName sets the source name. Attribute: sourceName.
func (s *SlicerCacheDefinition) SetSourceName(
	name string,
) {
	if name == "" {
		s.RemoveAttribute("sourceName", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"sourceName",
			"",
			name,
		),
	)
}

// PivotTables returns the pivot tables element, or nil if not present.
func (s *SlicerCacheDefinition) PivotTables() *SlicerCachePivotTables {
	elem := s.GetElement(
		"pivotTables",
		NamespaceSlicerX14,
	)
	if elem == nil {
		return nil
	}
	if pt, ok := elem.(*SlicerCachePivotTables); ok {
		return pt
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SlicerCachePivotTables{
			CompositeElementBase: comp,
		}
	}

	return nil
}

//
//nolint:lll,revive // long function signature
func (s *SlicerCacheDefinition) GetOrCreatePivotTables() *SlicerCachePivotTables {
	pt := s.PivotTables()
	if pt != nil {
		return pt
	}
	pt = NewSlicerCachePivotTables()
	s.AppendChild(pt)

	return pt
}

// Data returns the data element, or nil if not present.
func (s *SlicerCacheDefinition) Data() *SlicerCacheData {
	elem := s.GetElement(
		"data",
		NamespaceSlicerX14,
	)
	if elem == nil {
		return nil
	}
	if d, ok := elem.(*SlicerCacheData); ok {
		return d
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SlicerCacheData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateData returns the data element, creating if needed.
func (s *SlicerCacheDefinition) GetOrCreateData() *SlicerCacheData {
	d := s.Data()
	if d != nil {
		return d
	}
	d = NewSlicerCacheData()
	s.AppendChild(d)

	return d
}

// Clone creates a deep copy of this SlicerCacheDefinition element.
func (s *SlicerCacheDefinition) Clone() openxml.Element {
	cloned := s.PartRootElementBase.Clone()

	return &SlicerCacheDefinition{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this SlicerCacheDefinition element.
func (s *SlicerCacheDefinition) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.PartRootElementBase.CloneNode(
		deep,
	)

	return &SlicerCacheDefinition{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// SlicerCachePivotTables represents the pivot tables element
// (x14:pivotTables).
type SlicerCachePivotTables struct {
	*openxml.CompositeElementBase
}

// NewSlicerCachePivotTables creates a new SlicerCachePivotTables element.
func NewSlicerCachePivotTables() *SlicerCachePivotTables {
	elem := openxml.NewCompositeElement(
		NamespaceSlicerX14,
		"pivotTables",
		"x14", //nolint:revive // add-constant: namespace prefix
	)

	return &SlicerCachePivotTables{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this SlicerCachePivotTables element.
func (p *SlicerCachePivotTables) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &SlicerCachePivotTables{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SlicerCachePivotTables element.
func (p *SlicerCachePivotTables) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &SlicerCachePivotTables{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// SlicerCacheData represents the data element (x14:data).
type SlicerCacheData struct {
	*openxml.CompositeElementBase
}

// NewSlicerCacheData creates a new SlicerCacheData element.
func NewSlicerCacheData() *SlicerCacheData {
	elem := openxml.NewCompositeElement(
		NamespaceSlicerX14,
		"data",
		"x14",
	)

	return &SlicerCacheData{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this SlicerCacheData element.
func (s *SlicerCacheData) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &SlicerCacheData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SlicerCacheData element.
func (s *SlicerCacheData) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.CompositeElementBase.CloneNode(
		deep,
	)

	return &SlicerCacheData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
