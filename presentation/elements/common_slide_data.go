//nolint:revive // This file contains common slide data element implementation.
package elements

import (
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// CommonSlideData represents the common slide data element (p:cSld).
// This contains the visual content of a slide, layout, or master.
type CommonSlideData struct {
	*openxml.CompositeElementBase
}

// NewCommonSlideData creates a new CommonSlideData element.
func NewCommonSlideData() *CommonSlideData {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cSld",
		PrefixP,
	)
	csd := &CommonSlideData{
		CompositeElementBase: elem,
	}

	// Add required shape tree
	csd.AppendChild(NewShapeTree())

	return csd
}

// Name returns the name attribute.
func (csd *CommonSlideData) Name() string {
	attr, found := csd.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the name attribute.
func (csd *CommonSlideData) SetName(name string) {
	if name == "" {
		csd.RemoveAttribute("name", "")

		return
	}
	csd.SetAttribute(openxml.NewAttribute(
		"",
		"name",
		"",
		name,
	))
}

// Background returns the background element.
func (csd *CommonSlideData) Background() *SlideBackground {
	elem := csd.GetElement(
		"bg",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if bg, ok := elem.(*SlideBackground); ok {
		return bg
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideBackground{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetBackground sets the background element.
func (csd *CommonSlideData) SetBackground(
	bg *SlideBackground,
) {
	// Remove existing background
	if existing := csd.Background(); existing != nil {
		csd.RemoveChild(existing)
	}
	if bg != nil {
		// Insert before shape tree
		if st := csd.ShapeTree(); st != nil {
			csd.InsertBefore(bg, st)
		} else {
			csd.PrependChild(bg)
		}
	}
}

// ShapeTree returns the shape tree element.
func (csd *CommonSlideData) ShapeTree() *ShapeTree {
	elem := csd.GetElement(
		"spTree",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if st, ok := elem.(*ShapeTree); ok {
		return st
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ShapeTree{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateShapeTree returns the shape tree, creating if needed.
func (csd *CommonSlideData) GetOrCreateShapeTree() *ShapeTree {
	st := csd.ShapeTree()
	if st != nil {
		return st
	}
	st = NewShapeTree()
	// Insert after background if present, otherwise as first child after existing elements
	if bg := csd.Background(); bg != nil {
		csd.InsertAfter(st, bg)
	} else {
		csd.AppendChild(st)
	}

	return st
}

// Controls returns the controls element.
func (csd *CommonSlideData) Controls() *ControlList {
	elem := csd.GetElement(
		"controls",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if cl, ok := elem.(*ControlList); ok {
		return cl
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ControlList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this CommonSlideData element.
func (csd *CommonSlideData) Clone() openxml.Element {
	cloned := csd.CompositeElementBase.Clone()

	return &CommonSlideData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideBackground (p:bg)
// ===========================================================================

// SlideBackground represents the slide background element (p:bg).
type SlideBackground struct {
	*openxml.CompositeElementBase
}

// NewSlideBackground creates a new SlideBackground element.
func NewSlideBackground() *SlideBackground {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"bg",
		PrefixP,
	)

	return &SlideBackground{
		CompositeElementBase: elem,
	}
}

// BackgroundProperties returns the background properties.
func (bg *SlideBackground) BackgroundProperties() *BackgroundProperties {
	elem := bg.GetElement(
		"bgPr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if bgPr, ok := elem.(*BackgroundProperties); ok {
		return bgPr
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &BackgroundProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateBackgroundProperties returns the background properties, creating if needed.
func (bg *SlideBackground) GetOrCreateBackgroundProperties() *BackgroundProperties {
	bgPr := bg.BackgroundProperties()
	if bgPr != nil {
		return bgPr
	}
	bgPr = NewBackgroundProperties()
	bg.AppendChild(bgPr)

	return bgPr
}

// Clone creates a deep copy of this SlideBackground element.
func (bg *SlideBackground) Clone() openxml.Element {
	cloned := bg.CompositeElementBase.Clone()

	return &SlideBackground{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// SetSolidFill sets a solid color fill for the background.
func (bp *BackgroundProperties) SetSolidFill(hexColor string) {
	bp.removeFill()
	solidFill := drawingml.NewSolidFillWithRgb(hexColor)
	bp.AppendChild(solidFill)
}

// removeFill removes any existing fill elements.
func (bp *BackgroundProperties) removeFill() {
	if nf := bp.GetElement("noFill", NamespaceDrawingML); nf != nil {
		bp.RemoveChild(nf)
	}
	if sf := bp.GetElement("solidFill", NamespaceDrawingML); sf != nil {
		bp.RemoveChild(sf)
	}
	if gf := bp.GetElement("gradFill", NamespaceDrawingML); gf != nil {
		bp.RemoveChild(gf)
	}
	if pf := bp.GetElement("pattFill", NamespaceDrawingML); pf != nil {
		bp.RemoveChild(pf)
	}
	if bf := bp.GetElement("blipFill", NamespaceDrawingML); bf != nil {
		bp.RemoveChild(bf)
	}
}

// ===========================================================================
// ControlList (p:controls)
// ===========================================================================

// ControlList represents the controls element (p:controls).
type ControlList struct {
	*openxml.CompositeElementBase
}

// NewControlList creates a new ControlList element.
func NewControlList() *ControlList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"controls",
		PrefixP,
	)

	return &ControlList{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this ControlList element.
func (cl *ControlList) Clone() openxml.Element {
	cloned := cl.CompositeElementBase.Clone()

	return &ControlList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
