package elements

//revive:disable:file-length-limit many cache definition types
//revive:disable:max-public-structs many cache definition types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PivotCacheDefinition represents the pivot cache definition root element
// (x:pivotCacheDefinition).
// This element is the root of a pivot cache definition part and contains
// the cache field definitions and source information.
type PivotCacheDefinition struct {
	*openxml.PartRootElementBase
}

// NewPivotCacheDefinition creates a new PivotCacheDefinition element.
func NewPivotCacheDefinition() *PivotCacheDefinition {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"pivotCacheDefinition",
		PrefixDefault,
	)

	return &PivotCacheDefinition{
		PartRootElementBase: elem,
	}
}

// RelationshipId returns the relationship ID to the cache records.
// Attribute: r:id.
func (p *PivotCacheDefinition) RelationshipId() string {
	attr, found := p.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID to the cache records.
// Attribute: r:id.
func (p *PivotCacheDefinition) SetRelationshipId(
	id string,
) {
	if id == "" {
		p.RemoveAttribute(
			"id",
			NamespaceRelationships,
		)

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Invalid returns whether the cache is invalid. Attribute: invalid.
func (p *PivotCacheDefinition) Invalid() bool {
	attr, found := p.GetAttribute("invalid", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetInvalid sets whether the cache is invalid. Attribute: invalid.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetInvalid(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"invalid",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("invalid", "")
	}
}

// SaveData returns whether to save data. Attribute: saveData.
func (p *PivotCacheDefinition) SaveData() bool {
	attr, found := p.GetAttribute(
		"saveData",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSaveData sets whether to save data. Attribute: saveData.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetSaveData(
	value bool,
) {
	if value {
		p.RemoveAttribute(
			"saveData",
			"",
		) // true is default
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "saveData", "", attrValueFalse),
		)
	}
}

// RefreshOnLoad returns whether to refresh on load. Attribute: refreshOnLoad.
func (p *PivotCacheDefinition) RefreshOnLoad() bool {
	attr, found := p.GetAttribute(
		"refreshOnLoad",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetRefreshOnLoad sets whether to refresh on load. Attribute: refreshOnLoad.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetRefreshOnLoad(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"refreshOnLoad",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("refreshOnLoad", "")
	}
}

// OptimizeMemory returns whether to optimize memory. Attribute: optimizeMemory.
func (p *PivotCacheDefinition) OptimizeMemory() bool {
	attr, found := p.GetAttribute(
		"optimizeMemory",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetOptimizeMemory sets whether to optimize memory. Attribute: optimizeMemory.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetOptimizeMemory(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"optimizeMemory",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("optimizeMemory", "")
	}
}

// EnableRefresh returns whether refresh is enabled. Attribute: enableRefresh.
func (p *PivotCacheDefinition) EnableRefresh() bool {
	attr, found := p.GetAttribute(
		"enableRefresh",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnableRefresh sets whether refresh is enabled. Attribute: enableRefresh.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetEnableRefresh(
	value bool,
) {
	if value {
		p.RemoveAttribute(
			"enableRefresh",
			"",
		) // true is default
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "enableRefresh", "", attrValueFalse),
		)
	}
}

// RefreshedBy returns the name of the user who refreshed the cache.
// Attribute: refreshedBy.
func (p *PivotCacheDefinition) RefreshedBy() string {
	attr, found := p.GetAttribute(
		"refreshedBy",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRefreshedBy sets the name of the user who refreshed the cache.
// Attribute: refreshedBy.
func (p *PivotCacheDefinition) SetRefreshedBy(
	name string,
) {
	if name == "" {
		p.RemoveAttribute("refreshedBy", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"refreshedBy",
			"",
			name,
		),
	)
}

// RefreshedDate returns the date when the cache was refreshed.
// Attribute: refreshedDate.
func (p *PivotCacheDefinition) RefreshedDate() float64 {
	attr, found := p.GetAttribute(
		"refreshedDate",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetRefreshedDate sets the date when the cache was refreshed.
// Attribute: refreshedDate.
func (p *PivotCacheDefinition) SetRefreshedDate(
	date float64,
) {
	if date == 0 {
		p.RemoveAttribute("refreshedDate", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"refreshedDate",
			"",
			strconv.FormatFloat(
				date,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// CreatedVersion returns the created version. Attribute: createdVersion.
func (p *PivotCacheDefinition) CreatedVersion() uint8 {
	attr, found := p.GetAttribute(
		"createdVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetCreatedVersion sets the created version. Attribute: createdVersion.
func (p *PivotCacheDefinition) SetCreatedVersion(
	version uint8,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"createdVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// RefreshedVersion returns the refreshed version. Attribute: refreshedVersion.
func (p *PivotCacheDefinition) RefreshedVersion() uint8 {
	attr, found := p.GetAttribute(
		"refreshedVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetRefreshedVersion sets the refreshed version. Attribute: refreshedVersion.
func (p *PivotCacheDefinition) SetRefreshedVersion(
	version uint8,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"refreshedVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// MinRefreshableVersion returns the minimum refreshable version.
// Attribute: minRefreshableVersion.
func (p *PivotCacheDefinition) MinRefreshableVersion() uint8 {
	attr, found := p.GetAttribute(
		"minRefreshableVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetMinRefreshableVersion sets the minimum refreshable version.
// Attribute: minRefreshableVersion.
func (p *PivotCacheDefinition) SetMinRefreshableVersion(
	version uint8,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"minRefreshableVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// RecordCount returns the record count. Attribute: recordCount.
func (p *PivotCacheDefinition) RecordCount() uint32 {
	attr, found := p.GetAttribute(
		"recordCount",
		"",
	)
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

// SetRecordCount sets the record count. Attribute: recordCount.
func (p *PivotCacheDefinition) SetRecordCount(
	count uint32,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"recordCount",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// UpgradeOnRefresh returns whether to upgrade on refresh.
// Attribute: upgradeOnRefresh.
func (p *PivotCacheDefinition) UpgradeOnRefresh() bool {
	attr, found := p.GetAttribute(
		"upgradeOnRefresh",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetUpgradeOnRefresh sets whether to upgrade on refresh.
// Attribute: upgradeOnRefresh.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetUpgradeOnRefresh(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"upgradeOnRefresh",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("upgradeOnRefresh", "")
	}
}

// TupleCache returns whether this is a tuple cache. Attribute: tupleCache.
func (p *PivotCacheDefinition) TupleCache() bool {
	attr, found := p.GetAttribute(
		"tupleCache",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetTupleCache sets whether this is a tuple cache. Attribute: tupleCache.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetTupleCache(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"tupleCache",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("tupleCache", "")
	}
}

// SupportSubquery returns whether subquery is supported.
// Attribute: supportSubquery.
func (p *PivotCacheDefinition) SupportSubquery() bool {
	attr, found := p.GetAttribute(
		"supportSubquery",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSupportSubquery sets whether subquery is supported.
// Attribute: supportSubquery.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetSupportSubquery(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"supportSubquery",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("supportSubquery", "")
	}
}

// SupportAdvancedDrill returns whether advanced drill is supported.
// Attribute: supportAdvancedDrill.
func (p *PivotCacheDefinition) SupportAdvancedDrill() bool {
	attr, found := p.GetAttribute(
		"supportAdvancedDrill",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSupportAdvancedDrill sets whether advanced drill is supported.
// Attribute: supportAdvancedDrill.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotCacheDefinition) SetSupportAdvancedDrill(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"supportAdvancedDrill",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("supportAdvancedDrill", "")
	}
}

// CacheSource returns the CacheSource child element, or nil if not present.
func (p *PivotCacheDefinition) CacheSource() *CacheSource {
	elem := p.GetElement(
		"cacheSource",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cs, ok := elem.(*CacheSource); ok {
		return cs
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CacheSource{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCacheSource returns the CacheSource child element,
// creating it if needed.
func (p *PivotCacheDefinition) GetOrCreateCacheSource() *CacheSource {
	cs := p.CacheSource()
	if cs != nil {
		return cs
	}
	cs = NewCacheSource()
	// CacheSource should be first
	if first := p.FirstChild(); first != nil {
		p.InsertBefore(cs, first)
	} else {
		p.AppendChild(cs)
	}

	return cs
}

// CacheFields returns the CacheFields child element, or nil if not present.
func (p *PivotCacheDefinition) CacheFields() *CacheFields {
	elem := p.GetElement(
		"cacheFields",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cf, ok := elem.(*CacheFields); ok {
		return cf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CacheFields{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCacheFields returns the CacheFields child element,
// creating it if needed.
func (p *PivotCacheDefinition) GetOrCreateCacheFields() *CacheFields {
	cf := p.CacheFields()
	if cf != nil {
		return cf
	}
	cf = NewCacheFields()
	// CacheFields should come after cacheSource
	cs := p.CacheSource()
	if cs != nil {
		p.InsertAfter(cf, cs)
	} else {
		p.AppendChild(cf)
	}

	return cf
}

// Clone creates a deep copy of this PivotCacheDefinition element.
func (p *PivotCacheDefinition) Clone() openxml.Element {
	cloned := p.PartRootElementBase.Clone()

	return &PivotCacheDefinition{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this PivotCacheDefinition element.
func (p *PivotCacheDefinition) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.PartRootElementBase.CloneNode(
		deep,
	)

	return &PivotCacheDefinition{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CacheSource represents the cache source element (x:cacheSource).
type CacheSource struct {
	*openxml.CompositeElementBase
}

// NewCacheSource creates a new CacheSource element.
func NewCacheSource() *CacheSource {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cacheSource",
		PrefixDefault,
	)

	return &CacheSource{
		CompositeElementBase: elem,
	}
}

// Type returns the source type. Attribute: type.
// Valid values: worksheet, external, consolidation, scenario
func (cs *CacheSource) Type() string {
	attr, found := cs.GetAttribute("type", "")
	if !found {
		return attrValueWorksheet // Default is worksheet
	}

	return attr.Value()
}

// SetType sets the source type. Attribute: type.
func (cs *CacheSource) SetType(
	sourceType string,
) {
	if sourceType == "" ||
		sourceType == attrValueWorksheet {
		cs.RemoveAttribute(
			"type",
			"",
		) // worksheet is default

		return
	}
	cs.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			sourceType,
		),
	)
}

// ConnectionId returns the connection ID. Attribute: connectionId.
func (cs *CacheSource) ConnectionId() (uint32, bool) {
	attr, found := cs.GetAttribute(
		"connectionId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetConnectionId sets the connection ID. Attribute: connectionId.
func (cs *CacheSource) SetConnectionId(
	id uint32,
) {
	cs.SetAttribute(
		openxml.NewAttribute(
			"",
			"connectionId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// ClearConnectionId removes the connection ID attribute.
func (cs *CacheSource) ClearConnectionId() {
	cs.RemoveAttribute("connectionId", "")
}

// WorksheetSource returns the WorksheetSource child element,
// or nil if not present.
func (cs *CacheSource) WorksheetSource() *WorksheetSource {
	elem := cs.GetElement(
		"worksheetSource",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ws, ok := elem.(*WorksheetSource); ok {
		return ws
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &WorksheetSource{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateWorksheetSource returns the WorksheetSource child element,
// creating it if needed.
func (cs *CacheSource) GetOrCreateWorksheetSource() *WorksheetSource {
	ws := cs.WorksheetSource()
	if ws != nil {
		return ws
	}
	ws = NewWorksheetSource()
	cs.AppendChild(ws)

	return ws
}

// Clone creates a deep copy of this CacheSource element.
func (cs *CacheSource) Clone() openxml.Element {
	cloned := cs.CompositeElementBase.Clone()

	return &CacheSource{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CacheSource element.
func (cs *CacheSource) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cs.CompositeElementBase.CloneNode(
		deep,
	)

	return &CacheSource{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// WorksheetSource represents the worksheet source element (x:worksheetSource).
type WorksheetSource struct {
	*openxml.LeafElementBase
}

// NewWorksheetSource creates a new WorksheetSource element.
func NewWorksheetSource() *WorksheetSource {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"worksheetSource",
		PrefixDefault,
	)

	return &WorksheetSource{LeafElementBase: elem}
}

// Ref returns the cell range reference. Attribute: ref.
func (ws *WorksheetSource) Ref() string {
	attr, found := ws.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell range reference. Attribute: ref.
func (ws *WorksheetSource) SetRef(ref string) {
	if ref == "" {
		ws.RemoveAttribute("ref", "")

		return
	}
	ws.SetAttribute(
		openxml.NewAttribute("", "ref", "", ref),
	)
}

// Name returns the defined name reference. Attribute: name.
func (ws *WorksheetSource) Name() string {
	attr, found := ws.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the defined name reference. Attribute: name.
func (ws *WorksheetSource) SetName(name string) {
	if name == "" {
		ws.RemoveAttribute("name", "")

		return
	}
	ws.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Sheet returns the sheet name. Attribute: sheet.
func (ws *WorksheetSource) Sheet() string {
	attr, found := ws.GetAttribute("sheet", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSheet sets the sheet name. Attribute: sheet.
func (ws *WorksheetSource) SetSheet(
	sheet string,
) {
	if sheet == "" {
		ws.RemoveAttribute("sheet", "")

		return
	}
	ws.SetAttribute(
		openxml.NewAttribute(
			"",
			"sheet",
			"",
			sheet,
		),
	)
}

// RelationshipId returns the relationship ID. Attribute: r:id.
//
//nolint:revive // add-constant: attribute name standard pattern
func (ws *WorksheetSource) RelationshipId() string {
	attr, found := ws.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID. Attribute: r:id.
func (ws *WorksheetSource) SetRelationshipId(
	id string,
) {
	if id == "" {
		ws.RemoveAttribute(
			"id",
			NamespaceRelationships,
		)

		return
	}
	ws.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this WorksheetSource element.
func (ws *WorksheetSource) Clone() openxml.Element {
	cloned := ws.LeafElementBase.Clone()

	return &WorksheetSource{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this WorksheetSource element.
func (ws *WorksheetSource) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ws.LeafElementBase.CloneNode(deep)

	return &WorksheetSource{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CacheFields represents the cache fields collection element (x:cacheFields).
type CacheFields struct {
	*openxml.CompositeElementBase
}

// NewCacheFields creates a new CacheFields element.
func NewCacheFields() *CacheFields {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cacheFields",
		PrefixDefault,
	)

	return &CacheFields{
		CompositeElementBase: elem,
	}
}

// Count returns the count of cache fields. Attribute: count.
func (cf *CacheFields) Count() uint32 {
	attr, found := cf.GetAttribute("count", "")
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

// SetCount sets the count of cache fields. Attribute: count.
func (cf *CacheFields) SetCount(count uint32) {
	cf.SetAttribute(
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

// CacheFields returns an iterator over all CacheField elements.
func (cf *CacheFields) CacheFields() iter.Seq[*CacheField] {
	return func(yield func(*CacheField) bool) {
		for child := range cf.Children() {
			if child.LocalName() != "cacheField" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var field *CacheField
			if f, ok := child.(*CacheField); ok {
				field = f
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				field = &CacheField{CompositeElementBase: comp}
			}
			if field != nil && !yield(field) {
				return
			}
		}
	}
}

// FieldCount returns the count of cache field elements.
func (cf *CacheFields) FieldCount() int {
	count := 0
	for range cf.CacheFields() {
		count++
	}

	return count
}

// GetField returns the cache field at the given index, or nil if out of range.
func (cf *CacheFields) GetField(
	index int,
) *CacheField {
	if index < 0 {
		return nil
	}
	i := 0
	for field := range cf.CacheFields() {
		if i == index {
			return field
		}
		i++
	}

	return nil
}

// AddField adds a new CacheField element with the given name.
func (cf *CacheFields) AddField(
	name string,
) *CacheField {
	field := NewCacheField()
	field.SetName(name)
	cf.AppendChild(field)

	return field
}

// Clone creates a deep copy of this CacheFields element.
func (cf *CacheFields) Clone() openxml.Element {
	cloned := cf.CompositeElementBase.Clone()

	return &CacheFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CacheFields element.
func (cf *CacheFields) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.CompositeElementBase.CloneNode(
		deep,
	)

	return &CacheFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CacheField represents a single cache field element (x:cacheField).
type CacheField struct {
	*openxml.CompositeElementBase
}

// NewCacheField creates a new CacheField element.
func NewCacheField() *CacheField {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cacheField",
		PrefixDefault,
	)

	return &CacheField{CompositeElementBase: elem}
}

// Name returns the field name. Attribute: name.
//
//nolint:revive // add-constant: attribute name standard pattern
func (cf *CacheField) Name() string {
	attr, found := cf.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the field name. Attribute: name.
func (cf *CacheField) SetName(name string) {
	if name == "" {
		cf.RemoveAttribute("name", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Caption returns the field caption. Attribute: caption.
func (cf *CacheField) Caption() string {
	attr, found := cf.GetAttribute("caption", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCaption sets the field caption. Attribute: caption.
func (cf *CacheField) SetCaption(caption string) {
	if caption == "" {
		cf.RemoveAttribute("caption", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"caption",
			"",
			caption,
		),
	)
}

// PropertyName returns the property name. Attribute: propertyName.
func (cf *CacheField) PropertyName() string {
	attr, found := cf.GetAttribute(
		"propertyName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPropertyName sets the property name. Attribute: propertyName.
func (cf *CacheField) SetPropertyName(
	name string,
) {
	if name == "" {
		cf.RemoveAttribute("propertyName", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"propertyName",
			"",
			name,
		),
	)
}

// ServerField returns whether this is a server field. Attribute: serverField.
func (cf *CacheField) ServerField() bool {
	attr, found := cf.GetAttribute(
		"serverField",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetServerField sets whether this is a server field. Attribute: serverField.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CacheField) SetServerField(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"serverField",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("serverField", "")
	}
}

// UniqueList returns whether the field has unique list. Attribute: uniqueList.
func (cf *CacheField) UniqueList() bool {
	attr, found := cf.GetAttribute(
		"uniqueList",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetUniqueList sets whether the field has unique list. Attribute: uniqueList.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CacheField) SetUniqueList(value bool) {
	if value {
		cf.RemoveAttribute(
			"uniqueList",
			"",
		) // true is default
	} else {
		cf.SetAttribute(
			openxml.NewAttribute("", "uniqueList", "", attrValueFalse),
		)
	}
}

// NumFmtId returns the number format ID. Attribute: numFmtId.
func (cf *CacheField) NumFmtId() (uint32, bool) {
	attr, found := cf.GetAttribute("numFmtId", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetNumFmtId sets the number format ID. Attribute: numFmtId.
func (cf *CacheField) SetNumFmtId(id uint32) {
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"numFmtId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// ClearNumFmtId removes the number format ID attribute.
func (cf *CacheField) ClearNumFmtId() {
	cf.RemoveAttribute("numFmtId", "")
}

// Formula returns the formula for calculated field. Attribute: formula.
func (cf *CacheField) Formula() string {
	attr, found := cf.GetAttribute("formula", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFormula sets the formula for calculated field. Attribute: formula.
func (cf *CacheField) SetFormula(formula string) {
	if formula == "" {
		cf.RemoveAttribute("formula", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"formula",
			"",
			formula,
		),
	)
}

// SqlType returns the SQL type. Attribute: sqlType.
func (cf *CacheField) SqlType() int32 {
	attr, found := cf.GetAttribute("sqlType", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return int32(val)
}

// SetSqlType sets the SQL type. Attribute: sqlType.
func (cf *CacheField) SetSqlType(sqlType int32) {
	if sqlType == 0 {
		cf.RemoveAttribute("sqlType", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"sqlType",
			"",
			strconv.FormatInt(
				int64(sqlType),
				parseBase10,
			),
		),
	)
}

// Hierarchy returns the hierarchy index. Attribute: hierarchy.
func (cf *CacheField) Hierarchy() int32 {
	attr, found := cf.GetAttribute(
		"hierarchy",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return int32(val)
}

// SetHierarchy sets the hierarchy index. Attribute: hierarchy.
func (cf *CacheField) SetHierarchy(
	hierarchy int32,
) {
	if hierarchy == 0 {
		cf.RemoveAttribute("hierarchy", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"hierarchy",
			"",
			strconv.FormatInt(
				int64(hierarchy),
				parseBase10,
			),
		),
	)
}

// Level returns the level index. Attribute: level.
func (cf *CacheField) Level() uint32 {
	attr, found := cf.GetAttribute("level", "")
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

// SetLevel sets the level index. Attribute: level.
func (cf *CacheField) SetLevel(level uint32) {
	if level == 0 {
		cf.RemoveAttribute("level", "")

		return
	}
	cf.SetAttribute(
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

// DatabaseField returns whether this is a database field.
// Attribute: databaseField.
func (cf *CacheField) DatabaseField() bool {
	attr, found := cf.GetAttribute(
		"databaseField",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDatabaseField sets whether this is a database field.
// Attribute: databaseField.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CacheField) SetDatabaseField(
	value bool,
) {
	if value {
		cf.RemoveAttribute(
			"databaseField",
			"",
		) // true is default
	} else {
		cf.SetAttribute(
			openxml.NewAttribute("", "databaseField", "", attrValueFalse),
		)
	}
}

// MappingCount returns the mapping count. Attribute: mappingCount.
func (cf *CacheField) MappingCount() uint32 {
	attr, found := cf.GetAttribute(
		"mappingCount",
		"",
	)
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

// SetMappingCount sets the mapping count. Attribute: mappingCount.
func (cf *CacheField) SetMappingCount(
	count uint32,
) {
	if count == 0 {
		cf.RemoveAttribute("mappingCount", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"mappingCount",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// MemberPropertyField returns whether this is a member property field.
// Attribute: memberPropertyField.
func (cf *CacheField) MemberPropertyField() bool {
	attr, found := cf.GetAttribute(
		"memberPropertyField",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMemberPropertyField sets whether this is a member property field.
// Attribute: memberPropertyField.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CacheField) SetMemberPropertyField(
	value bool,
) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"memberPropertyField",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("memberPropertyField", "")
	}
}

// SharedItems returns the SharedItems child element, or nil if not present.
func (cf *CacheField) SharedItems() *SharedItems {
	elem := cf.GetElement(
		"sharedItems",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if si, ok := elem.(*SharedItems); ok {
		return si
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SharedItems{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSharedItems returns the SharedItems child element,
// creating it if needed.
func (cf *CacheField) GetOrCreateSharedItems() *SharedItems {
	si := cf.SharedItems()
	if si != nil {
		return si
	}
	si = NewSharedItems()
	cf.AppendChild(si)

	return si
}

// Clone creates a deep copy of this CacheField element.
func (cf *CacheField) Clone() openxml.Element {
	cloned := cf.CompositeElementBase.Clone()

	return &CacheField{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CacheField element.
func (cf *CacheField) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.CompositeElementBase.CloneNode(
		deep,
	)

	return &CacheField{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// SharedItems represents the shared items element (x:sharedItems).
type SharedItems struct {
	*openxml.CompositeElementBase
}

// NewSharedItems creates a new SharedItems element.
func NewSharedItems() *SharedItems {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sharedItems",
		PrefixDefault,
	)

	return &SharedItems{
		CompositeElementBase: elem,
	}
}

// Count returns the count of shared items. Attribute: count.
func (si *SharedItems) Count() uint32 {
	attr, found := si.GetAttribute("count", "")
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

// SetCount sets the count of shared items. Attribute: count.
//
//nolint:revive // add-constant: attribute name standard pattern
func (si *SharedItems) SetCount(count uint32) {
	si.SetAttribute(
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

// ContainsSemiMixedTypes returns whether items contain semi-mixed types.
// Attribute: containsSemiMixedTypes.
func (si *SharedItems) ContainsSemiMixedTypes() bool {
	attr, found := si.GetAttribute(
		"containsSemiMixedTypes",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetContainsSemiMixedTypes sets whether items contain semi-mixed types.
// Attribute: containsSemiMixedTypes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsSemiMixedTypes(
	value bool,
) {
	if value {
		si.RemoveAttribute(
			"containsSemiMixedTypes",
			"",
		) // true is default
	} else {
		si.SetAttribute(
			openxml.NewAttribute("", "containsSemiMixedTypes", "", attrValueFalse),
		)
	}
}

// ContainsNonDate returns whether items contain non-date values.
// Attribute: containsNonDate.
func (si *SharedItems) ContainsNonDate() bool {
	attr, found := si.GetAttribute(
		"containsNonDate",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetContainsNonDate sets whether items contain non-date values.
// Attribute: containsNonDate.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsNonDate(
	value bool,
) {
	if value {
		si.RemoveAttribute(
			"containsNonDate",
			"",
		) // true is default
	} else {
		si.SetAttribute(
			openxml.NewAttribute("", "containsNonDate", "", attrValueFalse),
		)
	}
}

// ContainsDate returns whether items contain date values.
// Attribute: containsDate.
func (si *SharedItems) ContainsDate() bool {
	attr, found := si.GetAttribute(
		"containsDate",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetContainsDate sets whether items contain date values.
// Attribute: containsDate.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsDate(
	value bool,
) {
	if value {
		si.SetAttribute(
			openxml.NewAttribute(
				"",
				"containsDate",
				"",
				attrValueTrue,
			),
		)
	} else {
		si.RemoveAttribute("containsDate", "")
	}
}

// ContainsString returns whether items contain string values.
// Attribute: containsString.
func (si *SharedItems) ContainsString() bool {
	attr, found := si.GetAttribute(
		"containsString",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetContainsString sets whether items contain string values.
// Attribute: containsString.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsString(
	value bool,
) {
	if value {
		si.RemoveAttribute(
			"containsString",
			"",
		) // true is default
	} else {
		si.SetAttribute(
			openxml.NewAttribute("", "containsString", "", attrValueFalse),
		)
	}
}

// ContainsBlank returns whether items contain blank values.
// Attribute: containsBlank.
func (si *SharedItems) ContainsBlank() bool {
	attr, found := si.GetAttribute(
		"containsBlank",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetContainsBlank sets whether items contain blank values.
// Attribute: containsBlank.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsBlank(
	value bool,
) {
	if value {
		si.SetAttribute(
			openxml.NewAttribute(
				"",
				"containsBlank",
				"",
				attrValueTrue,
			),
		)
	} else {
		si.RemoveAttribute("containsBlank", "")
	}
}

// ContainsMixedTypes returns whether items contain mixed types.
// Attribute: containsMixedTypes.
func (si *SharedItems) ContainsMixedTypes() bool {
	attr, found := si.GetAttribute(
		"containsMixedTypes",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetContainsMixedTypes sets whether items contain mixed types.
// Attribute: containsMixedTypes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsMixedTypes(
	value bool,
) {
	if value {
		si.SetAttribute(
			openxml.NewAttribute(
				"",
				"containsMixedTypes",
				"",
				attrValueTrue,
			),
		)
	} else {
		si.RemoveAttribute("containsMixedTypes", "")
	}
}

// ContainsNumber returns whether items contain number values.
// Attribute: containsNumber.
func (si *SharedItems) ContainsNumber() bool {
	attr, found := si.GetAttribute(
		"containsNumber",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetContainsNumber sets whether items contain number values.
// Attribute: containsNumber.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsNumber(
	value bool,
) {
	if value {
		si.SetAttribute(
			openxml.NewAttribute(
				"",
				"containsNumber",
				"",
				attrValueTrue,
			),
		)
	} else {
		si.RemoveAttribute("containsNumber", "")
	}
}

// ContainsInteger returns whether items contain integer values.
// Attribute: containsInteger.
func (si *SharedItems) ContainsInteger() bool {
	attr, found := si.GetAttribute(
		"containsInteger",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetContainsInteger sets whether items contain integer values.
// Attribute: containsInteger.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetContainsInteger(
	value bool,
) {
	if value {
		si.SetAttribute(
			openxml.NewAttribute(
				"",
				"containsInteger",
				"",
				attrValueTrue,
			),
		)
	} else {
		si.RemoveAttribute("containsInteger", "")
	}
}

// MinValue returns the minimum numeric value. Attribute: minValue.
func (si *SharedItems) MinValue() (float64, bool) {
	attr, found := si.GetAttribute("minValue", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val, true
}

// SetMinValue sets the minimum numeric value. Attribute: minValue.
func (si *SharedItems) SetMinValue(
	value float64,
) {
	si.SetAttribute(
		openxml.NewAttribute(
			"",
			"minValue",
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

// ClearMinValue removes the minimum value attribute.
func (si *SharedItems) ClearMinValue() {
	si.RemoveAttribute("minValue", "")
}

// MaxValue returns the maximum numeric value. Attribute: maxValue.
func (si *SharedItems) MaxValue() (float64, bool) {
	attr, found := si.GetAttribute("maxValue", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val, true
}

// SetMaxValue sets the maximum numeric value. Attribute: maxValue.
func (si *SharedItems) SetMaxValue(
	value float64,
) {
	si.SetAttribute(
		openxml.NewAttribute(
			"",
			"maxValue",
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

// ClearMaxValue removes the maximum value attribute.
func (si *SharedItems) ClearMaxValue() {
	si.RemoveAttribute("maxValue", "")
}

// MinDate returns the minimum date value. Attribute: minDate.
func (si *SharedItems) MinDate() string {
	attr, found := si.GetAttribute("minDate", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMinDate sets the minimum date value. Attribute: minDate.
func (si *SharedItems) SetMinDate(date string) {
	if date == "" {
		si.RemoveAttribute("minDate", "")

		return
	}
	si.SetAttribute(
		openxml.NewAttribute(
			"",
			"minDate",
			"",
			date,
		),
	)
}

// MaxDate returns the maximum date value. Attribute: maxDate.
func (si *SharedItems) MaxDate() string {
	attr, found := si.GetAttribute("maxDate", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMaxDate sets the maximum date value. Attribute: maxDate.
func (si *SharedItems) SetMaxDate(date string) {
	if date == "" {
		si.RemoveAttribute("maxDate", "")

		return
	}
	si.SetAttribute(
		openxml.NewAttribute(
			"",
			"maxDate",
			"",
			date,
		),
	)
}

// LongText returns whether items contain long text. Attribute: longText.
func (si *SharedItems) LongText() bool {
	attr, found := si.GetAttribute("longText", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetLongText sets whether items contain long text. Attribute: longText.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (si *SharedItems) SetLongText(value bool) {
	if value {
		si.SetAttribute(
			openxml.NewAttribute(
				"",
				"longText",
				"",
				attrValueTrue,
			),
		)
	} else {
		si.RemoveAttribute("longText", "")
	}
}

// AddString adds a string item to the shared items.
func (si *SharedItems) AddString(
	value string,
) *CacheItemString {
	item := NewCacheItemString()
	item.SetV(value)
	si.AppendChild(item)

	return item
}

// AddNumber adds a number item to the shared items.
func (si *SharedItems) AddNumber(
	value float64,
) *CacheItemNumber {
	item := NewCacheItemNumber()
	item.SetV(value)
	si.AppendChild(item)

	return item
}

// AddBoolean adds a boolean item to the shared items.
func (si *SharedItems) AddBoolean(
	value bool,
) *CacheItemBoolean {
	item := NewCacheItemBoolean()
	item.SetV(value)
	si.AppendChild(item)

	return item
}

// AddMissing adds a missing item to the shared items.
func (si *SharedItems) AddMissing() *CacheItemMissing {
	item := NewCacheItemMissing()
	si.AppendChild(item)

	return item
}

// AddError adds an error item to the shared items.
func (si *SharedItems) AddError(
	value string,
) *CacheItemError {
	item := NewCacheItemError()
	item.SetV(value)
	si.AppendChild(item)

	return item
}

// AddDateTime adds a date/time item to the shared items.
func (si *SharedItems) AddDateTime(
	value string,
) *CacheItemDateTime {
	item := NewCacheItemDateTime()
	item.SetV(value)
	si.AppendChild(item)

	return item
}

// Clone creates a deep copy of this SharedItems element.
func (si *SharedItems) Clone() openxml.Element {
	cloned := si.CompositeElementBase.Clone()

	return &SharedItems{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SharedItems element.
func (si *SharedItems) CloneNode(
	deep bool,
) openxml.Element {
	cloned := si.CompositeElementBase.CloneNode(
		deep,
	)

	return &SharedItems{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CacheItemString represents a string item element (x:s).
type CacheItemString struct {
	*openxml.LeafElementBase
}

// NewCacheItemString creates a new CacheItemString element.
func NewCacheItemString() *CacheItemString {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"s",
		PrefixDefault,
	)

	return &CacheItemString{LeafElementBase: elem}
}

// V returns the string value. Attribute: v.
func (c *CacheItemString) V() string {
	attr, found := c.GetAttribute("v", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetV sets the string value. Attribute: v.
func (c *CacheItemString) SetV(value string) {
	c.SetAttribute(
		openxml.NewAttribute("", "v", "", value),
	)
}

// U returns whether item is unused. Attribute: u.
func (c *CacheItemString) U() bool {
	attr, found := c.GetAttribute("u", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetU sets whether item is unused. Attribute: u.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemString) SetU(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"u",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("u", "")
	}
}

// F returns whether item has calculated value. Attribute: f.
func (c *CacheItemString) F() bool {
	attr, found := c.GetAttribute("f", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetF sets whether item has calculated value. Attribute: f.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemString) SetF(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"f",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("f", "")
	}
}

// C returns the caption. Attribute: c.
func (c *CacheItemString) C() string {
	attr, found := c.GetAttribute("c", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetC sets the caption. Attribute: c.
func (c *CacheItemString) SetC(caption string) {
	if caption == "" {
		c.RemoveAttribute("c", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"c",
			"",
			caption,
		),
	)
}

// Cp returns the member property count. Attribute: cp.
func (c *CacheItemString) Cp() uint32 {
	attr, found := c.GetAttribute("cp", "")
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

// SetCp sets the member property count. Attribute: cp.
func (c *CacheItemString) SetCp(count uint32) {
	if count == 0 {
		c.RemoveAttribute("cp", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"cp",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// In returns the index. Attribute: in.
func (c *CacheItemString) In() (uint32, bool) {
	attr, found := c.GetAttribute("in", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetIn sets the index. Attribute: in.
func (c *CacheItemString) SetIn(index uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"in",
			"",
			strconv.FormatUint(
				uint64(index),
				parseBase10,
			),
		),
	)
}

// ClearIn removes the index attribute.
func (c *CacheItemString) ClearIn() {
	c.RemoveAttribute("in", "")
}

// Bc returns the background color. Attribute: bc.
func (c *CacheItemString) Bc() string {
	attr, found := c.GetAttribute("bc", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetBc sets the background color. Attribute: bc.
func (c *CacheItemString) SetBc(color string) {
	if color == "" {
		c.RemoveAttribute("bc", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute("", "bc", "", color),
	)
}

// Fc returns the foreground color. Attribute: fc.
func (c *CacheItemString) Fc() string {
	attr, found := c.GetAttribute("fc", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFc sets the foreground color. Attribute: fc.
func (c *CacheItemString) SetFc(color string) {
	if color == "" {
		c.RemoveAttribute("fc", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute("", "fc", "", color),
	)
}

// I returns whether item is italic. Attribute: i.
func (c *CacheItemString) I() bool {
	attr, found := c.GetAttribute("i", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetI sets whether item is italic. Attribute: i.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemString) SetI(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"i",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("i", "")
	}
}

// Un returns whether item is underlined. Attribute: un.
func (c *CacheItemString) Un() bool {
	attr, found := c.GetAttribute("un", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetUn sets whether item is underlined. Attribute: un.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemString) SetUn(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"un",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("un", "")
	}
}

// St returns whether item has strikethrough. Attribute: st.
func (c *CacheItemString) St() bool {
	attr, found := c.GetAttribute("st", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSt sets whether item has strikethrough. Attribute: st.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemString) SetSt(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"st",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("st", "")
	}
}

// B returns whether item is bold. Attribute: b.
func (c *CacheItemString) B() bool {
	attr, found := c.GetAttribute("b", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetB sets whether item is bold. Attribute: b.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemString) SetB(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"b",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("b", "")
	}
}

// Clone creates a deep copy of this CacheItemString element.
func (c *CacheItemString) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CacheItemString{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CacheItemString element.
func (c *CacheItemString) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CacheItemString{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CacheItemNumber represents a number item element (x:n).
type CacheItemNumber struct {
	*openxml.LeafElementBase
}

// NewCacheItemNumber creates a new CacheItemNumber element.
func NewCacheItemNumber() *CacheItemNumber {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"n",
		PrefixDefault,
	)

	return &CacheItemNumber{LeafElementBase: elem}
}

// V returns the numeric value. Attribute: v.
func (c *CacheItemNumber) V() float64 {
	attr, found := c.GetAttribute("v", "")
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
//
//nolint:revive // add-constant: attribute name standard pattern
func (c *CacheItemNumber) SetV(value float64) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"v",
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

// U returns whether item is unused. Attribute: u.
func (c *CacheItemNumber) U() bool {
	attr, found := c.GetAttribute(
		"u", //nolint:revive // add-constant
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetU sets whether item is unused. Attribute: u.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemNumber) SetU(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"u",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("u", "")
	}
}

// F returns whether item has calculated value. Attribute: f.
func (c *CacheItemNumber) F() bool {
	attr, found := c.GetAttribute(
		"f", //nolint:revive // add-constant
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetF sets whether item has calculated value. Attribute: f.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CacheItemNumber) SetF(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"f",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("f", "")
	}
}

// C returns the caption. Attribute: c.
func (c *CacheItemNumber) C() string {
	attr, found := c.GetAttribute(
		"c", //nolint:revive // add-constant
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetC sets the caption. Attribute: c.
func (c *CacheItemNumber) SetC(caption string) {
	if caption == "" {
		c.RemoveAttribute("c", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"c",
			"",
			caption,
		),
	)
}

// Clone creates a deep copy of this CacheItemNumber element.
func (c *CacheItemNumber) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CacheItemNumber{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CacheItemNumber element.
func (c *CacheItemNumber) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CacheItemNumber{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CacheItemBoolean represents a boolean item element (x:b).
type CacheItemBoolean struct {
	*openxml.LeafElementBase
}

// NewCacheItemBoolean creates a new CacheItemBoolean element.
func NewCacheItemBoolean() *CacheItemBoolean {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"b", //nolint:revive // add-constant: element name
		PrefixDefault,
	)

	return &CacheItemBoolean{
		LeafElementBase: elem,
	}
}

// V returns the boolean value. Attribute: v.
func (c *CacheItemBoolean) V() bool {
	attr, found := c.GetAttribute("v", "")
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
func (c *CacheItemBoolean) SetV(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"v",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.SetAttribute(
			openxml.NewAttribute("", "v", "", attrValueFalse),
		)
	}
}

// Clone creates a deep copy of this CacheItemBoolean element.
func (c *CacheItemBoolean) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CacheItemBoolean{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CacheItemBoolean element.
func (c *CacheItemBoolean) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CacheItemBoolean{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CacheItemMissing represents a missing item element (x:m).
type CacheItemMissing struct {
	*openxml.LeafElementBase
}

// NewCacheItemMissing creates a new CacheItemMissing element.
func NewCacheItemMissing() *CacheItemMissing {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"m",
		PrefixDefault,
	)

	return &CacheItemMissing{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy of this CacheItemMissing element.
func (c *CacheItemMissing) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CacheItemMissing{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CacheItemMissing element.
func (c *CacheItemMissing) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CacheItemMissing{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CacheItemError represents an error item element (x:e).
type CacheItemError struct {
	*openxml.LeafElementBase
}

// NewCacheItemError creates a new CacheItemError element.
func NewCacheItemError() *CacheItemError {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"e",
		PrefixDefault,
	)

	return &CacheItemError{LeafElementBase: elem}
}

// V returns the error value. Attribute: v.
func (c *CacheItemError) V() string {
	attr, found := c.GetAttribute("v", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetV sets the error value. Attribute: v.
func (c *CacheItemError) SetV(value string) {
	c.SetAttribute(
		openxml.NewAttribute("", "v", "", value),
	)
}

// Clone creates a deep copy of this CacheItemError element.
func (c *CacheItemError) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CacheItemError{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CacheItemError element.
func (c *CacheItemError) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CacheItemError{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CacheItemDateTime represents a date/time item element (x:d).
type CacheItemDateTime struct {
	*openxml.LeafElementBase
}

// NewCacheItemDateTime creates a new CacheItemDateTime element.
func NewCacheItemDateTime() *CacheItemDateTime {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"d",
		PrefixDefault,
	)

	return &CacheItemDateTime{
		LeafElementBase: elem,
	}
}

// V returns the date/time value (ISO 8601 format). Attribute: v.
func (c *CacheItemDateTime) V() string {
	attr, found := c.GetAttribute("v", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetV sets the date/time value (ISO 8601 format). Attribute: v.
func (c *CacheItemDateTime) SetV(value string) {
	c.SetAttribute(
		openxml.NewAttribute("", "v", "", value),
	)
}

// Clone creates a deep copy of this CacheItemDateTime element.
func (c *CacheItemDateTime) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CacheItemDateTime{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CacheItemDateTime element.
func (c *CacheItemDateTime) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CacheItemDateTime{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
