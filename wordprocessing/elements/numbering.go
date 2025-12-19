package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Numbering represents the root element of the numbering part (w:numbering).
type Numbering struct {
	*openxml.CompositeElementBase
	nextAbstractNumId int
	nextNumId         int
}

// NewNumbering creates a new Numbering element.
func NewNumbering() *Numbering {
	elem := openxml.NewCompositeElement(NamespaceWML, "numbering", PrefixW)
	return &Numbering{
		CompositeElementBase: elem,
		nextAbstractNumId:    0,
		nextNumId:            1,
	}
}

// AbstractNums returns an iterator over all AbstractNum elements.
func (n *Numbering) AbstractNums() iter.Seq[*AbstractNum] {
	return func(yield func(*AbstractNum) bool) {
		for child := range n.Children() {
			if child.LocalName() == "abstractNum" && child.NamespaceURI() == NamespaceWML {
				var an *AbstractNum
				if absNum, ok := child.(*AbstractNum); ok {
					an = absNum
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					an = &AbstractNum{CompositeElementBase: comp}
				}
				if an != nil && !yield(an) {
					return
				}
			}
		}
	}
}

// NumInstances returns an iterator over all NumberingInstance elements.
func (n *Numbering) NumInstances() iter.Seq[*NumberingInstance] {
	return func(yield func(*NumberingInstance) bool) {
		for child := range n.Children() {
			if child.LocalName() == "num" && child.NamespaceURI() == NamespaceWML {
				var ni *NumberingInstance
				if numInst, ok := child.(*NumberingInstance); ok {
					ni = numInst
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					ni = &NumberingInstance{CompositeElementBase: comp}
				}
				if ni != nil && !yield(ni) {
					return
				}
			}
		}
	}
}

// GetAbstractNum returns the AbstractNum with the given ID, or nil if not found.
func (n *Numbering) GetAbstractNum(id int) *AbstractNum {
	for an := range n.AbstractNums() {
		if an.AbstractNumId() == id {
			return an
		}
	}
	return nil
}

// GetNumInstance returns the NumberingInstance with the given ID, or nil if not found.
func (n *Numbering) GetNumInstance(id int) *NumberingInstance {
	for ni := range n.NumInstances() {
		if ni.NumId() == id {
			return ni
		}
	}
	return nil
}

// AddAbstractNum adds an abstract numbering definition and returns the assigned ID.
func (n *Numbering) AddAbstractNum(abstractNum *AbstractNum) int {
	// Find the highest existing ID
	maxId := -1
	for an := range n.AbstractNums() {
		if an.AbstractNumId() > maxId {
			maxId = an.AbstractNumId()
		}
	}
	if maxId >= n.nextAbstractNumId {
		n.nextAbstractNumId = maxId + 1
	}

	// Assign a unique ID
	id := n.nextAbstractNumId
	abstractNum.SetAbstractNumId(id)
	n.nextAbstractNumId++

	// Insert before num elements (abstractNum elements come first)
	var insertBefore openxml.Element
	for child := range n.Children() {
		if child.LocalName() == "num" && child.NamespaceURI() == NamespaceWML {
			insertBefore = child
			break
		}
	}
	if insertBefore != nil {
		n.InsertBefore(abstractNum, insertBefore)
	} else {
		n.AppendChild(abstractNum)
	}

	return id
}

// AddNumInstance adds a numbering instance and returns the assigned ID.
func (n *Numbering) AddNumInstance(instance *NumberingInstance) int {
	// Find the highest existing ID
	maxId := 0
	for ni := range n.NumInstances() {
		if ni.NumId() > maxId {
			maxId = ni.NumId()
		}
	}
	if maxId >= n.nextNumId {
		n.nextNumId = maxId + 1
	}

	// Assign a unique ID
	id := n.nextNumId
	instance.SetNumId(id)
	n.nextNumId++

	n.AppendChild(instance)
	return id
}

// CreateNumberingInstance creates a new numbering instance referencing an abstract num.
func (n *Numbering) CreateNumberingInstance(abstractNumId int) *NumberingInstance {
	ni := NewNumberingInstance(abstractNumId)
	n.AddNumInstance(ni)
	return ni
}

// Clone creates a deep copy of this Numbering element.
func (n *Numbering) Clone() openxml.Element {
	return &Numbering{
		CompositeElementBase: n.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
		nextAbstractNumId:    n.nextAbstractNumId,
		nextNumId:            n.nextNumId,
	}
}

// CloneNode creates a copy of this Numbering element.
func (n *Numbering) CloneNode(deep bool) openxml.Element {
	return &Numbering{
		CompositeElementBase: n.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
		nextAbstractNumId:    n.nextAbstractNumId,
		nextNumId:            n.nextNumId,
	}
}

// NumberingInstance represents a numbering instance (w:num).
type NumberingInstance struct {
	*openxml.CompositeElementBase
}

// NewNumberingInstance creates a new NumberingInstance element.
func NewNumberingInstance(abstractNumId int) *NumberingInstance {
	elem := openxml.NewCompositeElement(NamespaceWML, "num", PrefixW)
	ni := &NumberingInstance{CompositeElementBase: elem}
	ni.SetAbstractNumIdRef(abstractNumId)
	return ni
}

// NumId returns the numbering instance ID.
func (ni *NumberingInstance) NumId() int {
	attr, found := ni.GetAttribute("numId", NamespaceWML)
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}
	return val
}

// SetNumId sets the numbering instance ID.
func (ni *NumberingInstance) SetNumId(id int) {
	ni.SetAttribute(openxml.NewAttribute(NamespaceWML, "numId", PrefixW, strconv.Itoa(id)))
}

// AbstractNumId returns the referenced abstract numbering definition ID.
func (ni *NumberingInstance) AbstractNumId() int {
	elem := ni.GetElement("abstractNumId", NamespaceWML)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute("val", NamespaceWML)
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}
	return val
}

// SetAbstractNumIdRef sets the referenced abstract numbering definition ID.
func (ni *NumberingInstance) SetAbstractNumIdRef(id int) {
	elem := ni.GetElement("abstractNumId", NamespaceWML)
	if elem == nil {
		elem = openxml.NewCompositeElement(NamespaceWML, "abstractNumId", PrefixW)
		ni.AppendChild(elem)
	}
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, strconv.Itoa(id)))
}

// LevelOverrides returns an iterator over all level overrides.
func (ni *NumberingInstance) LevelOverrides() iter.Seq[*LevelOverride] {
	return func(yield func(*LevelOverride) bool) {
		for child := range ni.Children() {
			if child.LocalName() == "lvlOverride" && child.NamespaceURI() == NamespaceWML {
				var lo *LevelOverride
				if lvlOvr, ok := child.(*LevelOverride); ok {
					lo = lvlOvr
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					lo = &LevelOverride{CompositeElementBase: comp}
				}
				if lo != nil && !yield(lo) {
					return
				}
			}
		}
	}
}

// AddLevelOverride adds a level override at the specified level.
func (ni *NumberingInstance) AddLevelOverride(levelIndex int) *LevelOverride {
	lo := NewLevelOverride(levelIndex)
	ni.AppendChild(lo)
	return lo
}

// Clone creates a deep copy of this NumberingInstance element.
func (ni *NumberingInstance) Clone() openxml.Element {
	return &NumberingInstance{
		CompositeElementBase: ni.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// LevelOverride represents a level override within a numbering instance (w:lvlOverride).
type LevelOverride struct {
	*openxml.CompositeElementBase
}

// NewLevelOverride creates a new LevelOverride element.
func NewLevelOverride(levelIndex int) *LevelOverride {
	elem := openxml.NewCompositeElement(NamespaceWML, "lvlOverride", PrefixW)
	lo := &LevelOverride{CompositeElementBase: elem}
	lo.SetLevelIndex(levelIndex)
	return lo
}

// LevelIndex returns the level being overridden (0-8).
func (lo *LevelOverride) LevelIndex() int {
	attr, found := lo.GetAttribute("ilvl", NamespaceWML)
	if !found {
		return 0
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return 0
	}
	return val
}

// SetLevelIndex sets the level being overridden (0-8).
func (lo *LevelOverride) SetLevelIndex(index int) {
	lo.SetAttribute(openxml.NewAttribute(NamespaceWML, "ilvl", PrefixW, strconv.Itoa(index)))
}

// StartOverride returns the overridden start value, or -1 if not set.
func (lo *LevelOverride) StartOverride() int {
	elem := lo.GetElement("startOverride", NamespaceWML)
	if elem == nil {
		return -1
	}
	attr, found := elem.GetAttribute("val", NamespaceWML)
	if !found {
		return -1
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return -1
	}
	return val
}

// SetStartOverride sets the overridden start value.
func (lo *LevelOverride) SetStartOverride(start int) {
	elem := lo.GetElement("startOverride", NamespaceWML)
	if elem == nil {
		elem = openxml.NewCompositeElement(NamespaceWML, "startOverride", PrefixW)
		lo.AppendChild(elem)
	}
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, strconv.Itoa(start)))
}

// Level returns the complete level override definition, or nil if not set.
func (lo *LevelOverride) Level() *Level {
	elem := lo.GetElement("lvl", NamespaceWML)
	if elem == nil {
		return nil
	}
	if lvl, ok := elem.(*Level); ok {
		return lvl
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Level{CompositeElementBase: comp}
	}
	return nil
}

// SetLevel sets the complete level override definition.
func (lo *LevelOverride) SetLevel(level *Level) {
	// Remove existing level
	if existing := lo.GetElement("lvl", NamespaceWML); existing != nil {
		lo.RemoveChild(existing)
	}
	if level != nil {
		lo.AppendChild(level)
	}
}

// Clone creates a deep copy of this LevelOverride element.
func (lo *LevelOverride) Clone() openxml.Element {
	return &LevelOverride{
		CompositeElementBase: lo.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
