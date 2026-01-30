//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// Counter for generating unique slide layout filenames.
var slideLayoutUniqueCounter uint64

// SlideLayoutPart represents a slide layout part (ppt/slideLayouts/slideLayout1.xml, etc.).
type SlideLayoutPart struct {
	*openxml.OpenXmlPartData
}

// newSlideLayoutPart creates a new slide layout part.
func newSlideLayoutPart(
	slideMasterPart *SlideMasterPart,
	uri string,
) (*SlideLayoutPart, error) {
	packPart, relID, err := slideMasterPart.addChildPart(
		uri,
		ContentTypeSlideLayout,
		RelationshipTypeSlideLayout,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideLayout,
		packPart,
		slideMasterPart,
	)
	partData.SetRelationshipID(relID)

	slp := &SlideLayoutPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide layout content
	slp.initializeContent()

	// Add to slide master part's child parts
	if err := slideMasterPart.AddPart(slp, relID); err != nil {
		return nil, err
	}

	return slp, nil
}

// NewSlideLayoutPart creates a new slide layout part from a presentation part.
// This is used when creating layouts directly from the presentation.
func NewSlideLayoutPart(
	presPart *PresentationPart,
) (*SlideLayoutPart, error) {
	num := atomic.AddUint64(
		&slideLayoutUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/slideLayouts/slideLayout%d.xml",
		num,
	)

	packPart, relID, err := presPart.addChildPart(
		uri,
		ContentTypeSlideLayout,
		RelationshipTypeSlideLayout,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideLayout,
		packPart,
		presPart,
	)
	partData.SetRelationshipID(relID)

	slp := &SlideLayoutPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide layout content
	slp.initializeContent()

	// Add to presentation part's child parts
	if err := presPart.AddPart(slp, relID); err != nil {
		return nil, err
	}

	return slp, nil
}

// initializeContent sets up minimal slide layout content.
func (slp *SlideLayoutPart) initializeContent() {
	sl := elements.NewSlideLayout()
	slp.SetRootElement(sl)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlideLayoutPart) FixedContentType() string {
	return ContentTypeSlideLayout
}

// SlideLayout returns the root SlideLayout element.
func (slp *SlideLayoutPart) SlideLayout() *elements.SlideLayout {
	root := slp.RootElement()
	if root == nil {
		return nil
	}
	if sl, ok := root.(*elements.SlideLayout); ok {
		return sl
	}

	return nil
}

// EffectiveColorMap returns the effective PresentationColorMap for this layout,
// retrieved from the associated SlideMaster.
func (slp *SlideLayoutPart) EffectiveColorMap() *elements.PresentationColorMap {
	if smp := slp.SlideMasterPart(); smp != nil {
		return smp.SlideMaster().ColorMap()
	}

	return nil
}

// EffectiveTransition returns the effective SlideTransition for this layout,
// resolving from SlideMaster if not present locally.
func (slp *SlideLayoutPart) EffectiveTransition() *elements.SlideTransition {
	if tr := slp.SlideLayout().Transition(); tr != nil {
		return tr
	}
	if smp := slp.SlideMasterPart(); smp != nil {
		return smp.SlideMaster().Transition()
	}

	return nil
}

// EffectiveTiming returns the effective SlideTiming for this layout,
// resolving from SlideMaster if not present locally.
func (slp *SlideLayoutPart) EffectiveTiming() *elements.SlideTiming {
	if t := slp.SlideLayout().Timing(); t != nil {
		return t
	}
	if smp := slp.SlideMasterPart(); smp != nil {
		return smp.SlideMaster().Timing()
	}

	return nil
}

// EffectiveHeaderFooter returns the effective ExtHeaderFooter for this layout,
// resolving from SlideMaster if not present locally.
func (slp *SlideLayoutPart) EffectiveHeaderFooter() *elements.ExtHeaderFooter {
	if hf := slp.SlideLayout().HeaderFooter(); hf != nil {
		return hf
	}
	if smp := slp.SlideMasterPart(); smp != nil {
		return smp.SlideMaster().HeaderFooter()
	}

	return nil
}

// EffectiveTextStyles returns the effective TextStyles for this layout,
// resolving from SlideMaster if not present locally.
func (slp *SlideLayoutPart) EffectiveTextStyles() *elements.TextStyles {
	if ts := slp.SlideLayout().TextStyles(); ts != nil {
		return ts
	}
	if smp := slp.SlideMasterPart(); smp != nil {
		return smp.SlideMaster().TextStyles()
	}

	return nil
}

// EffectiveExtensionList returns the effective ExtensionListModify for this layout,
// resolving from SlideMaster if not present locally.
func (slp *SlideLayoutPart) EffectiveExtensionList() *elements.ExtensionListModify {
	if el := slp.SlideLayout().ExtensionList(); el != nil {
		return el
	}
	if smp := slp.SlideMasterPart(); smp != nil {
		return smp.SlideMaster().ExtensionList()
	}

	return nil
}

// SlideMasterPart returns the slide master part associated with this layout.
func (slp *SlideLayoutPart) SlideMasterPart() *SlideMasterPart {
	// A slide layout can be owned by a slide master or by the presentation directly.
	// We check for SlideMasterPart in parent containers.
	container := slp.Container()
	if smp, ok := container.(*SlideMasterPart); ok {
		return smp
	}

	// Also check related parts
	for part := range slp.Parts() {
		if smp, ok := part.(*SlideMasterPart); ok {
			return smp
		}
	}

	return nil
}

// AddPlaceholder adds a placeholder to the layout.
func (slp *SlideLayoutPart) AddPlaceholder(phType elements.PlaceholderType, idx int) *elements.Shape {
	return slp.SlideLayout().AddPlaceholder(phType, idx)
}

// GetStream returns a reader for the part content.
func (slp *SlideLayoutPart) GetStream() io.Reader {
	return slp.OpenXmlPartData.GetStream()
}

// Ensure SlideLayoutPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlideLayoutPart)(
	nil,
)

// SlideLayoutPartFactory creates a SlideLayoutPart from a URI and container.
func SlideLayoutPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideLayout,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewSlideLayout()
		},
	)

	return &SlideLayoutPart{
		OpenXmlPartData: partData,
	}
}

// Register the SlideLayoutPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSlideLayout,
			RelationshipType:   RelationshipTypeSlideLayout,
			Factory:            SlideLayoutPartFactory,
			DefaultURI:         "/ppt/slideLayouts/slideLayout1.xml",
			IsFixedContentType: true,
		},
	)
}
