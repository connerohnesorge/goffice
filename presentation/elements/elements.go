// nolint
//
//revive:disable:file-length-limit,function-length,unchecked-type-assertion,line-length-limit,max-public-structs,cognitive-complexity,cyclomatic,comments-density // Generated code
package elements

import (
	"encoding/xml"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/types"
)

// Defines the NonVisualContentPartProperties Class.
type NonVisualContentPartProperties struct {
	*openxml.CompositeElementBase
	XMLName                               xml.Name                               `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main nvContentPartPr"`
	NonVisualDrawingProperties            *NonVisualDrawingProperties            `xml:"cNvPr,omitempty"`
	NonVisualInkContentPartProperties     *NonVisualInkContentPartProperties     `xml:"cNvContentPartPr,omitempty"`
	ApplicationNonVisualDrawingProperties *ApplicationNonVisualDrawingProperties `xml:"nvPr,omitempty"`
}

func NewNonVisualContentPartProperties() *NonVisualContentPartProperties {
	ret := &NonVisualContentPartProperties{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"nvContentPartPr",
		"p14",
	)
	return ret
}

func (m *NonVisualContentPartProperties) Clone() openxml.Element {
	ret := NewNonVisualContentPartProperties()
	if m.NonVisualDrawingProperties != nil {
		ret.NonVisualDrawingProperties = m.NonVisualDrawingProperties.Clone().(*NonVisualDrawingProperties)
	}
	if m.NonVisualInkContentPartProperties != nil {
		ret.NonVisualInkContentPartProperties = m.NonVisualInkContentPartProperties.Clone().(*NonVisualInkContentPartProperties)
	}
	if m.ApplicationNonVisualDrawingProperties != nil {
		ret.ApplicationNonVisualDrawingProperties = m.ApplicationNonVisualDrawingProperties.Clone().(*ApplicationNonVisualDrawingProperties)
	}
	return ret
}

func (m *NonVisualContentPartProperties) Validate() error {
	if m.NonVisualDrawingProperties != nil {
		if err := m.NonVisualDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	if m.NonVisualInkContentPartProperties != nil {
		if err := m.NonVisualInkContentPartProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ApplicationNonVisualDrawingProperties != nil {
		if err := m.ApplicationNonVisualDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Transform2D Class.
type Transform2D struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main xfrm"`
	Rotation       *types.Int32Value   `xml:"rot,attr,omitempty"`
	HorizontalFlip *types.BooleanValue `xml:"flipH,attr,omitempty"`
	VerticalFlip   *types.BooleanValue `xml:"flipV,attr,omitempty"`
	Offset         *drawingml.Offset   `xml:"off,omitempty"`
	// Skipped DrawingML type not yet implemented: Extents
}

func NewTransform2D() *Transform2D {
	ret := &Transform2D{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"xfrm",
		"p14",
	)
	return ret
}

func (m *Transform2D) Clone() openxml.Element {
	ret := NewTransform2D()
	if m.Rotation != nil {
		v := *m.Rotation
		ret.Rotation = &v
	}
	if m.HorizontalFlip != nil {
		v := *m.HorizontalFlip
		ret.HorizontalFlip = &v
	}
	if m.VerticalFlip != nil {
		v := *m.VerticalFlip
		ret.VerticalFlip = &v
	}
	if m.Offset != nil {
		v := m.Offset.Clone()
		ret.Offset = &v
	}
	return ret
}

func (m *Transform2D) Validate() error {
	return nil
}

// Defines the ExtensionListModify Class.
type ExtensionListModify struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main extLst"`
	Modify    *types.BooleanValue `xml:"mod,attr,omitempty"`
	Extension *Extension          `xml:"ext,omitempty"`
}

func NewExtensionListModify() *ExtensionListModify {
	ret := &ExtensionListModify{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p14",
	)
	return ret
}

func (m *ExtensionListModify) Clone() openxml.Element {
	ret := NewExtensionListModify()
	if m.Modify != nil {
		v := *m.Modify
		ret.Modify = &v
	}
	if m.Extension != nil {
		ret.Extension = m.Extension.Clone().(*Extension)
	}
	return ret
}

func (m *ExtensionListModify) Validate() error {
	if m.Extension != nil {
		if err := m.Extension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Media Class.
type Media struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main media"`
	Embed             *types.StringValue `xml:"r:embed,attr,omitempty"`
	Link              *types.StringValue `xml:"r:link,attr,omitempty"`
	MediaTrim         *MediaTrim         `xml:"trim,omitempty"`
	MediaFade         *MediaFade         `xml:"fade,omitempty"`
	MediaBookmarkList *MediaBookmarkList `xml:"bmkLst,omitempty"`
	ExtensionList     *ExtensionList     `xml:"extLst,omitempty"`
}

func NewMedia() *Media {
	ret := &Media{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"media",
		"p14",
	)
	return ret
}

func (m *Media) Clone() openxml.Element {
	ret := NewMedia()
	if m.Embed != nil {
		v := *m.Embed
		ret.Embed = &v
	}
	if m.Link != nil {
		v := *m.Link
		ret.Link = &v
	}
	if m.MediaTrim != nil {
		ret.MediaTrim = m.MediaTrim.Clone().(*MediaTrim)
	}
	if m.MediaFade != nil {
		ret.MediaFade = m.MediaFade.Clone().(*MediaFade)
	}
	if m.MediaBookmarkList != nil {
		ret.MediaBookmarkList = m.MediaBookmarkList.Clone().(*MediaBookmarkList)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *Media) Validate() error {
	if m.MediaTrim != nil {
		if err := m.MediaTrim.Validate(); err != nil {
			return err
		}
	}
	if m.MediaFade != nil {
		if err := m.MediaFade.Validate(); err != nil {
			return err
		}
	}
	if m.MediaBookmarkList != nil {
		if err := m.MediaBookmarkList.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the VortexTransition Class.
type VortexTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main vortex"`
}

func NewVortexTransition() *VortexTransition {
	ret := &VortexTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"vortex",
		"p14",
	)
	return ret
}

func (m *VortexTransition) Clone() openxml.Element {
	ret := NewVortexTransition()
	return ret
}

func (m *VortexTransition) Validate() error {
	return nil
}

// Defines the PanTransition Class.
type PanTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main pan"`
}

func NewPanTransition() *PanTransition {
	ret := &PanTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pan",
		"p14",
	)
	return ret
}

func (m *PanTransition) Clone() openxml.Element {
	ret := NewPanTransition()
	return ret
}

func (m *PanTransition) Validate() error {
	return nil
}

// Defines the SideDirectionTransitionType Class.
type SideDirectionTransitionType struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                         `xml:""`
	Direction *types.EnumValue[TransitionSlideDirectionValues] `xml:"dir,attr,omitempty"`
}

func NewSideDirectionTransitionType() *SideDirectionTransitionType {
	ret := &SideDirectionTransitionType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p14",
	)
	return ret
}

func (m *SideDirectionTransitionType) Clone() openxml.Element {
	ret := NewSideDirectionTransitionType()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *SideDirectionTransitionType) Validate() error {
	return nil
}

// Defines the SwitchTransition Class.
type SwitchTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main switch"`
}

func NewSwitchTransition() *SwitchTransition {
	ret := &SwitchTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"switch",
		"p14",
	)
	return ret
}

func (m *SwitchTransition) Clone() openxml.Element {
	ret := NewSwitchTransition()
	return ret
}

func (m *SwitchTransition) Validate() error {
	return nil
}

// Defines the FlipTransition Class.
type FlipTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main flip"`
}

func NewFlipTransition() *FlipTransition {
	ret := &FlipTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"flip",
		"p14",
	)
	return ret
}

func (m *FlipTransition) Clone() openxml.Element {
	ret := NewFlipTransition()
	return ret
}

func (m *FlipTransition) Validate() error {
	return nil
}

// Defines the FerrisTransition Class.
type FerrisTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main ferris"`
}

func NewFerrisTransition() *FerrisTransition {
	ret := &FerrisTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"ferris",
		"p14",
	)
	return ret
}

func (m *FerrisTransition) Clone() openxml.Element {
	ret := NewFerrisTransition()
	return ret
}

func (m *FerrisTransition) Validate() error {
	return nil
}

// Defines the GalleryTransition Class.
type GalleryTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main gallery"`
}

func NewGalleryTransition() *GalleryTransition {
	ret := &GalleryTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"gallery",
		"p14",
	)
	return ret
}

func (m *GalleryTransition) Clone() openxml.Element {
	ret := NewGalleryTransition()
	return ret
}

func (m *GalleryTransition) Validate() error {
	return nil
}

// Defines the ConveyorTransition Class.
type ConveyorTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main conveyor"`
}

func NewConveyorTransition() *ConveyorTransition {
	ret := &ConveyorTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"conveyor",
		"p14",
	)
	return ret
}

func (m *ConveyorTransition) Clone() openxml.Element {
	ret := NewConveyorTransition()
	return ret
}

func (m *ConveyorTransition) Validate() error {
	return nil
}

// Defines the LeftRightDirectionTransitionType Class.
type LeftRightDirectionTransitionType struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                                 `xml:""`
	Direction *types.EnumValue[TransitionLeftRightDirectionTypeValues] `xml:"dir,attr,omitempty"`
}

func NewLeftRightDirectionTransitionType() *LeftRightDirectionTransitionType {
	ret := &LeftRightDirectionTransitionType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p14",
	)
	return ret
}

func (m *LeftRightDirectionTransitionType) Clone() openxml.Element {
	ret := NewLeftRightDirectionTransitionType()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *LeftRightDirectionTransitionType) Validate() error {
	return nil
}

// Defines the RippleTransition Class.
type RippleTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main ripple"`
	Direction *types.StringValue `xml:"dir,attr,omitempty"`
}

func NewRippleTransition() *RippleTransition {
	ret := &RippleTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"ripple",
		"p14",
	)
	return ret
}

func (m *RippleTransition) Clone() openxml.Element {
	ret := NewRippleTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *RippleTransition) Validate() error {
	return nil
}

// Defines the HoneycombTransition Class.
type HoneycombTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main honeycomb"`
}

func NewHoneycombTransition() *HoneycombTransition {
	ret := &HoneycombTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"honeycomb",
		"p14",
	)
	return ret
}

func (m *HoneycombTransition) Clone() openxml.Element {
	ret := NewHoneycombTransition()
	return ret
}

func (m *HoneycombTransition) Validate() error {
	return nil
}

// Defines the FlashTransition Class.
type FlashTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main flash"`
}

func NewFlashTransition() *FlashTransition {
	ret := &FlashTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"flash",
		"p14",
	)
	return ret
}

func (m *FlashTransition) Clone() openxml.Element {
	ret := NewFlashTransition()
	return ret
}

func (m *FlashTransition) Validate() error {
	return nil
}

// Defines the EmptyType Class.
type EmptyType struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:""`
}

func NewEmptyType() *EmptyType {
	ret := &EmptyType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p14",
	)
	return ret
}

func (m *EmptyType) Clone() openxml.Element {
	ret := NewEmptyType()
	return ret
}

func (m *EmptyType) Validate() error {
	return nil
}

// Defines the PrismTransition Class.
type PrismTransition struct {
	*openxml.LeafElementBase
	XMLName    xml.Name                                         `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main prism"`
	Direction  *types.EnumValue[TransitionSlideDirectionValues] `xml:"dir,attr,omitempty"`
	IsContent  *types.BooleanValue                              `xml:"isContent,attr,omitempty"`
	IsInverted *types.BooleanValue                              `xml:"isInverted,attr,omitempty"`
}

func NewPrismTransition() *PrismTransition {
	ret := &PrismTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"prism",
		"p14",
	)
	return ret
}

func (m *PrismTransition) Clone() openxml.Element {
	ret := NewPrismTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	if m.IsContent != nil {
		v := *m.IsContent
		ret.IsContent = &v
	}
	if m.IsInverted != nil {
		v := *m.IsInverted
		ret.IsInverted = &v
	}
	return ret
}

func (m *PrismTransition) Validate() error {
	return nil
}

// Defines the DoorsTransition Class.
type DoorsTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main doors"`
}

func NewDoorsTransition() *DoorsTransition {
	ret := &DoorsTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"doors",
		"p14",
	)
	return ret
}

func (m *DoorsTransition) Clone() openxml.Element {
	ret := NewDoorsTransition()
	return ret
}

func (m *DoorsTransition) Validate() error {
	return nil
}

// Defines the WindowTransition Class.
type WindowTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main window"`
}

func NewWindowTransition() *WindowTransition {
	ret := &WindowTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"window",
		"p14",
	)
	return ret
}

func (m *WindowTransition) Clone() openxml.Element {
	ret := NewWindowTransition()
	return ret
}

func (m *WindowTransition) Validate() error {
	return nil
}

// Defines the OrientationTransitionType Class.
type OrientationTransitionType struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                          `xml:""`
	Direction *types.EnumValue[DirectionValues] `xml:"dir,attr,omitempty"`
}

func NewOrientationTransitionType() *OrientationTransitionType {
	ret := &OrientationTransitionType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p14",
	)
	return ret
}

func (m *OrientationTransitionType) Clone() openxml.Element {
	ret := NewOrientationTransitionType()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *OrientationTransitionType) Validate() error {
	return nil
}

// Defines the GlitterTransition Class.
type GlitterTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                         `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main glitter"`
	Direction *types.EnumValue[TransitionSlideDirectionValues] `xml:"dir,attr,omitempty"`
	Pattern   *types.EnumValue[TransitionPatternValues]        `xml:"pattern,attr,omitempty"`
}

func NewGlitterTransition() *GlitterTransition {
	ret := &GlitterTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"glitter",
		"p14",
	)
	return ret
}

func (m *GlitterTransition) Clone() openxml.Element {
	ret := NewGlitterTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	if m.Pattern != nil {
		v := *m.Pattern
		ret.Pattern = &v
	}
	return ret
}

func (m *GlitterTransition) Validate() error {
	return nil
}

// Defines the WarpTransition Class.
type WarpTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                         `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main warp"`
	Direction *types.EnumValue[TransitionInOutDirectionValues] `xml:"dir,attr,omitempty"`
}

func NewWarpTransition() *WarpTransition {
	ret := &WarpTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"warp",
		"p14",
	)
	return ret
}

func (m *WarpTransition) Clone() openxml.Element {
	ret := NewWarpTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *WarpTransition) Validate() error {
	return nil
}

// Defines the FlythroughTransition Class.
type FlythroughTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                         `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main flythrough"`
	Direction *types.EnumValue[TransitionInOutDirectionValues] `xml:"dir,attr,omitempty"`
	HasBounce *types.BooleanValue                              `xml:"hasBounce,attr,omitempty"`
}

func NewFlythroughTransition() *FlythroughTransition {
	ret := &FlythroughTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"flythrough",
		"p14",
	)
	return ret
}

func (m *FlythroughTransition) Clone() openxml.Element {
	ret := NewFlythroughTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	if m.HasBounce != nil {
		v := *m.HasBounce
		ret.HasBounce = &v
	}
	return ret
}

func (m *FlythroughTransition) Validate() error {
	return nil
}

// Defines the ShredTransition Class.
type ShredTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                         `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main shred"`
	Pattern   *types.EnumValue[TransitionShredPatternValues]   `xml:"pattern,attr,omitempty"`
	Direction *types.EnumValue[TransitionInOutDirectionValues] `xml:"dir,attr,omitempty"`
}

func NewShredTransition() *ShredTransition {
	ret := &ShredTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"shred",
		"p14",
	)
	return ret
}

func (m *ShredTransition) Clone() openxml.Element {
	ret := NewShredTransition()
	if m.Pattern != nil {
		v := *m.Pattern
		ret.Pattern = &v
	}
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *ShredTransition) Validate() error {
	return nil
}

// Defines the RevealTransition Class.
type RevealTransition struct {
	*openxml.LeafElementBase
	XMLName      xml.Name                                                 `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main reveal"`
	ThroughBlack *types.BooleanValue                                      `xml:"thruBlk,attr,omitempty"`
	Direction    *types.EnumValue[TransitionLeftRightDirectionTypeValues] `xml:"dir,attr,omitempty"`
}

func NewRevealTransition() *RevealTransition {
	ret := &RevealTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"reveal",
		"p14",
	)
	return ret
}

func (m *RevealTransition) Clone() openxml.Element {
	ret := NewRevealTransition()
	if m.ThroughBlack != nil {
		v := *m.ThroughBlack
		ret.ThroughBlack = &v
	}
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *RevealTransition) Validate() error {
	return nil
}

// Defines the WheelReverseTransition Class.
type WheelReverseTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main wheelReverse"`
	Spokes  *types.UInt32Value `xml:"spokes,attr,omitempty"`
}

func NewWheelReverseTransition() *WheelReverseTransition {
	ret := &WheelReverseTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"wheelReverse",
		"p14",
	)
	return ret
}

func (m *WheelReverseTransition) Clone() openxml.Element {
	ret := NewWheelReverseTransition()
	if m.Spokes != nil {
		v := *m.Spokes
		ret.Spokes = &v
	}
	return ret
}

func (m *WheelReverseTransition) Validate() error {
	return nil
}

// Defines the BookmarkTarget Class.
type BookmarkTarget struct {
	*openxml.LeafElementBase
	XMLName      xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main bmkTgt"`
	ShapeId      *types.UInt32Value `xml:"spid,attr,omitempty"`
	BookmarkName *types.StringValue `xml:"bmkName,attr,omitempty"`
}

func NewBookmarkTarget() *BookmarkTarget {
	ret := &BookmarkTarget{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bmkTgt",
		"p14",
	)
	return ret
}

func (m *BookmarkTarget) Clone() openxml.Element {
	ret := NewBookmarkTarget()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	if m.BookmarkName != nil {
		v := *m.BookmarkName
		ret.BookmarkName = &v
	}
	return ret
}

func (m *BookmarkTarget) Validate() error {
	return nil
}

// Defines the SectionProperties Class.
type SectionProperties struct {
	*openxml.CompositeElementBase
	XMLName    xml.Name    `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main sectionPr"`
	SectionOld *SectionOld `xml:"section,omitempty"`
}

func NewSectionProperties() *SectionProperties {
	ret := &SectionProperties{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sectionPr",
		"p14",
	)
	return ret
}

func (m *SectionProperties) Clone() openxml.Element {
	ret := NewSectionProperties()
	if m.SectionOld != nil {
		ret.SectionOld = m.SectionOld.Clone().(*SectionOld)
	}
	return ret
}

func (m *SectionProperties) Validate() error {
	if m.SectionOld != nil {
		if err := m.SectionOld.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SectionList Class.
type SectionList struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main sectionLst"`
	Section *Section `xml:"section,omitempty"`
}

func NewSectionList() *SectionList {
	ret := &SectionList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sectionLst",
		"p14",
	)
	return ret
}

func (m *SectionList) Clone() openxml.Element {
	ret := NewSectionList()
	if m.Section != nil {
		ret.Section = m.Section.Clone().(*Section)
	}
	return ret
}

func (m *SectionList) Validate() error {
	if m.Section != nil {
		if err := m.Section.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the BrowseMode Class.
type BrowseMode struct {
	*openxml.LeafElementBase
	XMLName    xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main browseMode"`
	ShowStatus *types.BooleanValue `xml:"showStatus,attr,omitempty"`
}

func NewBrowseMode() *BrowseMode {
	ret := &BrowseMode{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"browseMode",
		"p14",
	)
	return ret
}

func (m *BrowseMode) Clone() openxml.Element {
	ret := NewBrowseMode()
	if m.ShowStatus != nil {
		v := *m.ShowStatus
		ret.ShowStatus = &v
	}
	return ret
}

func (m *BrowseMode) Validate() error {
	return nil
}

// Defines the LaserColor Class.
type LaserColor struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main laserClr"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewLaserColor() *LaserColor {
	ret := &LaserColor{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"laserClr",
		"p14",
	)
	return ret
}

func (m *LaserColor) Clone() openxml.Element {
	ret := NewLaserColor()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *LaserColor) Validate() error {
	return nil
}

// Defines the DefaultImageDpi Class.
type DefaultImageDpi struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main defaultImageDpi"`
	Val     *types.UInt32Value `xml:"val,attr,omitempty"`
}

func NewDefaultImageDpi() *DefaultImageDpi {
	ret := &DefaultImageDpi{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"defaultImageDpi",
		"p14",
	)
	return ret
}

func (m *DefaultImageDpi) Clone() openxml.Element {
	ret := NewDefaultImageDpi()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *DefaultImageDpi) Validate() error {
	return nil
}

// Defines the DiscardImageEditData Class.
type DiscardImageEditData struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main discardImageEditData"`
	Val     *types.BooleanValue `xml:"val,attr,omitempty"`
}

func NewDiscardImageEditData() *DiscardImageEditData {
	ret := &DiscardImageEditData{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"discardImageEditData",
		"p14",
	)
	return ret
}

func (m *DiscardImageEditData) Clone() openxml.Element {
	ret := NewDiscardImageEditData()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *DiscardImageEditData) Validate() error {
	return nil
}

// Defines the ShowMediaControls Class.
type ShowMediaControls struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main showMediaCtrls"`
	Val     *types.BooleanValue `xml:"val,attr,omitempty"`
}

func NewShowMediaControls() *ShowMediaControls {
	ret := &ShowMediaControls{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"showMediaCtrls",
		"p14",
	)
	return ret
}

func (m *ShowMediaControls) Clone() openxml.Element {
	ret := NewShowMediaControls()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *ShowMediaControls) Validate() error {
	return nil
}

// Defines the LaserTraceList Class.
type LaserTraceList struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name        `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main laserTraceLst"`
	TracePointList *TracePointList `xml:"tracePtLst,omitempty"`
}

func NewLaserTraceList() *LaserTraceList {
	ret := &LaserTraceList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"laserTraceLst",
		"p14",
	)
	return ret
}

func (m *LaserTraceList) Clone() openxml.Element {
	ret := NewLaserTraceList()
	if m.TracePointList != nil {
		ret.TracePointList = m.TracePointList.Clone().(*TracePointList)
	}
	return ret
}

func (m *LaserTraceList) Validate() error {
	if m.TracePointList != nil {
		if err := m.TracePointList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CreationId Class.
type CreationId struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main creationId"`
}

func NewCreationId() *CreationId {
	ret := &CreationId{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"creationId",
		"p14",
	)
	return ret
}

func (m *CreationId) Clone() openxml.Element {
	ret := NewCreationId()
	return ret
}

func (m *CreationId) Validate() error {
	return nil
}

// Defines the ModificationId Class.
type ModificationId struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main modId"`
}

func NewModificationId() *ModificationId {
	ret := &ModificationId{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"modId",
		"p14",
	)
	return ret
}

func (m *ModificationId) Clone() openxml.Element {
	ret := NewModificationId()
	return ret
}

func (m *ModificationId) Validate() error {
	return nil
}

// Defines the RandomIdType Class.
type RandomIdType struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:""`
	Val     *types.UInt32Value `xml:"val,attr,omitempty"`
}

func NewRandomIdType() *RandomIdType {
	ret := &RandomIdType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p14",
	)
	return ret
}

func (m *RandomIdType) Clone() openxml.Element {
	ret := NewRandomIdType()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *RandomIdType) Validate() error {
	return nil
}

// Defines the ShowEventRecordList Class.
type ShowEventRecordList struct {
	*openxml.CompositeElementBase
	XMLName            xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main showEvtLst"`
	TriggerEventRecord *TriggerEventRecord `xml:"triggerEvt,omitempty"`
	PlayEventRecord    *PlayEventRecord    `xml:"playEvt,omitempty"`
	StopEventRecord    *StopEventRecord    `xml:"stopEvt,omitempty"`
	PauseEventRecord   *PauseEventRecord   `xml:"pauseEvt,omitempty"`
	ResumeEventRecord  *ResumeEventRecord  `xml:"resumeEvt,omitempty"`
	SeekEventRecord    *SeekEventRecord    `xml:"seekEvt,omitempty"`
	NullEventRecord    *NullEventRecord    `xml:"nullEvt,omitempty"`
}

func NewShowEventRecordList() *ShowEventRecordList {
	ret := &ShowEventRecordList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"showEvtLst",
		"p14",
	)
	return ret
}

func (m *ShowEventRecordList) Clone() openxml.Element {
	ret := NewShowEventRecordList()
	if m.TriggerEventRecord != nil {
		ret.TriggerEventRecord = m.TriggerEventRecord.Clone().(*TriggerEventRecord)
	}
	if m.PlayEventRecord != nil {
		ret.PlayEventRecord = m.PlayEventRecord.Clone().(*PlayEventRecord)
	}
	if m.StopEventRecord != nil {
		ret.StopEventRecord = m.StopEventRecord.Clone().(*StopEventRecord)
	}
	if m.PauseEventRecord != nil {
		ret.PauseEventRecord = m.PauseEventRecord.Clone().(*PauseEventRecord)
	}
	if m.ResumeEventRecord != nil {
		ret.ResumeEventRecord = m.ResumeEventRecord.Clone().(*ResumeEventRecord)
	}
	if m.SeekEventRecord != nil {
		ret.SeekEventRecord = m.SeekEventRecord.Clone().(*SeekEventRecord)
	}
	if m.NullEventRecord != nil {
		ret.NullEventRecord = m.NullEventRecord.Clone().(*NullEventRecord)
	}
	return ret
}

func (m *ShowEventRecordList) Validate() error {
	if m.TriggerEventRecord != nil {
		if err := m.TriggerEventRecord.Validate(); err != nil {
			return err
		}
	}
	if m.PlayEventRecord != nil {
		if err := m.PlayEventRecord.Validate(); err != nil {
			return err
		}
	}
	if m.StopEventRecord != nil {
		if err := m.StopEventRecord.Validate(); err != nil {
			return err
		}
	}
	if m.PauseEventRecord != nil {
		if err := m.PauseEventRecord.Validate(); err != nil {
			return err
		}
	}
	if m.ResumeEventRecord != nil {
		if err := m.ResumeEventRecord.Validate(); err != nil {
			return err
		}
	}
	if m.SeekEventRecord != nil {
		if err := m.SeekEventRecord.Validate(); err != nil {
			return err
		}
	}
	if m.NullEventRecord != nil {
		if err := m.NullEventRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the NonVisualDrawingProperties Class.
type NonVisualDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName     xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main cNvPr"`
	Id          *types.UInt32Value  `xml:"id,attr,omitempty"`
	Name        *types.StringValue  `xml:"name,attr,omitempty"`
	Description *types.StringValue  `xml:"descr,attr,omitempty"`
	Hidden      *types.BooleanValue `xml:"hidden,attr,omitempty"`
	Title       *types.StringValue  `xml:"title,attr,omitempty"`
	// Skipped DrawingML type not yet implemented: HyperlinkOnClick
	// Skipped DrawingML type not yet implemented: HyperlinkOnHover
	// Skipped DrawingML type not yet implemented: NonVisualDrawingPropertiesExtensionList
}

func NewNonVisualDrawingProperties() *NonVisualDrawingProperties {
	ret := &NonVisualDrawingProperties{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvPr",
		"p14",
	)
	return ret
}

func (m *NonVisualDrawingProperties) Clone() openxml.Element {
	ret := NewNonVisualDrawingProperties()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Description != nil {
		v := *m.Description
		ret.Description = &v
	}
	if m.Hidden != nil {
		v := *m.Hidden
		ret.Hidden = &v
	}
	if m.Title != nil {
		v := *m.Title
		ret.Title = &v
	}
	return ret
}

func (m *NonVisualDrawingProperties) Validate() error {
	return nil
}

// Defines the NonVisualInkContentPartProperties Class.
type NonVisualInkContentPartProperties struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main cNvContentPartPr"`
	IsComment *types.BooleanValue `xml:"isComment,attr,omitempty"`
	// Skipped DrawingML type not yet implemented: ContentPartLocks
	// Skipped DrawingML type not yet implemented: OfficeArtExtensionList
}

func NewNonVisualInkContentPartProperties() *NonVisualInkContentPartProperties {
	ret := &NonVisualInkContentPartProperties{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvContentPartPr",
		"p14",
	)
	return ret
}

func (m *NonVisualInkContentPartProperties) Clone() openxml.Element {
	ret := NewNonVisualInkContentPartProperties()
	if m.IsComment != nil {
		v := *m.IsComment
		ret.IsComment = &v
	}
	return ret
}

func (m *NonVisualInkContentPartProperties) Validate() error {
	return nil
}

// Defines the ApplicationNonVisualDrawingProperties Class.
type ApplicationNonVisualDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main nvPr"`
	IsPhoto          *types.BooleanValue `xml:"isPhoto,attr,omitempty"`
	UserDrawn        *types.BooleanValue `xml:"userDrawn,attr,omitempty"`
	PlaceholderShape *PlaceholderShape   `xml:"ph,omitempty"`
	// Skipped DrawingML type not yet implemented: AudioFromCD
	// Skipped DrawingML type not yet implemented: WaveAudioFile
	// Skipped DrawingML type not yet implemented: AudioFromFile
	// Skipped DrawingML type not yet implemented: VideoFromFile
	// Skipped DrawingML type not yet implemented: QuickTimeFromFile
	CustomerDataList                                   *CustomerDataList                                   `xml:"custDataLst,omitempty"`
	ApplicationNonVisualDrawingPropertiesExtensionList *ApplicationNonVisualDrawingPropertiesExtensionList `xml:"extLst,omitempty"`
}

func NewApplicationNonVisualDrawingProperties() *ApplicationNonVisualDrawingProperties {
	ret := &ApplicationNonVisualDrawingProperties{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"nvPr",
		"p14",
	)
	return ret
}

func (m *ApplicationNonVisualDrawingProperties) Clone() openxml.Element {
	ret := NewApplicationNonVisualDrawingProperties()
	if m.IsPhoto != nil {
		v := *m.IsPhoto
		ret.IsPhoto = &v
	}
	if m.UserDrawn != nil {
		v := *m.UserDrawn
		ret.UserDrawn = &v
	}
	if m.PlaceholderShape != nil {
		ret.PlaceholderShape = m.PlaceholderShape.Clone().(*PlaceholderShape)
	}
	if m.CustomerDataList != nil {
		ret.CustomerDataList = m.CustomerDataList.Clone().(*CustomerDataList)
	}
	if m.ApplicationNonVisualDrawingPropertiesExtensionList != nil {
		ret.ApplicationNonVisualDrawingPropertiesExtensionList = m.ApplicationNonVisualDrawingPropertiesExtensionList.Clone().(*ApplicationNonVisualDrawingPropertiesExtensionList)
	}
	return ret
}

func (m *ApplicationNonVisualDrawingProperties) Validate() error {
	if m.PlaceholderShape != nil {
		if err := m.PlaceholderShape.Validate(); err != nil {
			return err
		}
	}
	if m.CustomerDataList != nil {
		if err := m.CustomerDataList.Validate(); err != nil {
			return err
		}
	}
	if m.ApplicationNonVisualDrawingPropertiesExtensionList != nil {
		if err := m.ApplicationNonVisualDrawingPropertiesExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the MediaBookmark Class.
type MediaBookmark struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main bmk"`
	Name    *types.StringValue `xml:"name,attr,omitempty"`
	Time    *types.StringValue `xml:"time,attr,omitempty"`
}

func NewMediaBookmark() *MediaBookmark {
	ret := &MediaBookmark{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bmk",
		"p14",
	)
	return ret
}

func (m *MediaBookmark) Clone() openxml.Element {
	ret := NewMediaBookmark()
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	return ret
}

func (m *MediaBookmark) Validate() error {
	return nil
}

// Defines the MediaTrim Class.
type MediaTrim struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main trim"`
	Start   *types.StringValue `xml:"st,attr,omitempty"`
	End     *types.StringValue `xml:"end,attr,omitempty"`
}

func NewMediaTrim() *MediaTrim {
	ret := &MediaTrim{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"trim",
		"p14",
	)
	return ret
}

func (m *MediaTrim) Clone() openxml.Element {
	ret := NewMediaTrim()
	if m.Start != nil {
		v := *m.Start
		ret.Start = &v
	}
	if m.End != nil {
		v := *m.End
		ret.End = &v
	}
	return ret
}

func (m *MediaTrim) Validate() error {
	return nil
}

// Defines the MediaFade Class.
type MediaFade struct {
	*openxml.LeafElementBase
	XMLName     xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main fade"`
	InDuration  *types.StringValue `xml:"in,attr,omitempty"`
	OutDuration *types.StringValue `xml:"out,attr,omitempty"`
}

func NewMediaFade() *MediaFade {
	ret := &MediaFade{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"fade",
		"p14",
	)
	return ret
}

func (m *MediaFade) Clone() openxml.Element {
	ret := NewMediaFade()
	if m.InDuration != nil {
		v := *m.InDuration
		ret.InDuration = &v
	}
	if m.OutDuration != nil {
		v := *m.OutDuration
		ret.OutDuration = &v
	}
	return ret
}

func (m *MediaFade) Validate() error {
	return nil
}

// Defines the MediaBookmarkList Class.
type MediaBookmarkList struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main bmkLst"`
	MediaBookmark *MediaBookmark `xml:"bmk,omitempty"`
}

func NewMediaBookmarkList() *MediaBookmarkList {
	ret := &MediaBookmarkList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bmkLst",
		"p14",
	)
	return ret
}

func (m *MediaBookmarkList) Clone() openxml.Element {
	ret := NewMediaBookmarkList()
	if m.MediaBookmark != nil {
		ret.MediaBookmark = m.MediaBookmark.Clone().(*MediaBookmark)
	}
	return ret
}

func (m *MediaBookmarkList) Validate() error {
	if m.MediaBookmark != nil {
		if err := m.MediaBookmark.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ExtensionList Class.
type ExtensionList struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name   `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main extLst"`
	Extension *Extension `xml:"ext,omitempty"`
}

func NewExtensionList() *ExtensionList {
	ret := &ExtensionList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p14",
	)
	return ret
}

func (m *ExtensionList) Clone() openxml.Element {
	ret := NewExtensionList()
	if m.Extension != nil {
		ret.Extension = m.Extension.Clone().(*Extension)
	}
	return ret
}

func (m *ExtensionList) Validate() error {
	if m.Extension != nil {
		if err := m.Extension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SectionOld Class.
type SectionOld struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main section"`
	Name          *types.StringValue `xml:"name,attr,omitempty"`
	SlideIdList   *types.StringValue `xml:"slideIdLst,attr,omitempty"`
	Id            *types.StringValue `xml:"id,attr,omitempty"`
	ExtensionList *ExtensionList     `xml:"extLst,omitempty"`
}

func NewSectionOld() *SectionOld {
	ret := &SectionOld{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"section",
		"p14",
	)
	return ret
}

func (m *SectionOld) Clone() openxml.Element {
	ret := NewSectionOld()
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.SlideIdList != nil {
		v := *m.SlideIdList
		ret.SlideIdList = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *SectionOld) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SectionSlideIdListEntry Class.
type SectionSlideIdListEntry struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main sldId"`
	Id      *types.UInt32Value `xml:"id,attr,omitempty"`
}

func NewSectionSlideIdListEntry() *SectionSlideIdListEntry {
	ret := &SectionSlideIdListEntry{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sldId",
		"p14",
	)
	return ret
}

func (m *SectionSlideIdListEntry) Clone() openxml.Element {
	ret := NewSectionSlideIdListEntry()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *SectionSlideIdListEntry) Validate() error {
	return nil
}

// Defines the SectionSlideIdList Class.
type SectionSlideIdList struct {
	*openxml.CompositeElementBase
	XMLName                 xml.Name                 `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main sldIdLst"`
	SectionSlideIdListEntry *SectionSlideIdListEntry `xml:"sldId,omitempty"`
}

func NewSectionSlideIdList() *SectionSlideIdList {
	ret := &SectionSlideIdList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldIdLst",
		"p14",
	)
	return ret
}

func (m *SectionSlideIdList) Clone() openxml.Element {
	ret := NewSectionSlideIdList()
	if m.SectionSlideIdListEntry != nil {
		ret.SectionSlideIdListEntry = m.SectionSlideIdListEntry.Clone().(*SectionSlideIdListEntry)
	}
	return ret
}

func (m *SectionSlideIdList) Validate() error {
	if m.SectionSlideIdListEntry != nil {
		if err := m.SectionSlideIdListEntry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Section Class.
type Section struct {
	*openxml.CompositeElementBase
	XMLName            xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main section"`
	Name               *types.StringValue  `xml:"name,attr,omitempty"`
	Id                 *types.StringValue  `xml:"id,attr,omitempty"`
	SectionSlideIdList *SectionSlideIdList `xml:"sldIdLst,omitempty"`
	ExtensionList      *ExtensionList      `xml:"extLst,omitempty"`
}

func NewSection() *Section {
	ret := &Section{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"section",
		"p14",
	)
	return ret
}

func (m *Section) Clone() openxml.Element {
	ret := NewSection()
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.SectionSlideIdList != nil {
		ret.SectionSlideIdList = m.SectionSlideIdList.Clone().(*SectionSlideIdList)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *Section) Validate() error {
	if m.SectionSlideIdList != nil {
		if err := m.SectionSlideIdList.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TracePoint Class.
type TracePoint struct {
	*openxml.LeafElementBase
	XMLName     xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main tracePt"`
	Time        *types.StringValue `xml:"t,attr,omitempty"`
	XCoordinate *types.Int64Value  `xml:"x,attr,omitempty"`
	YCoordinate *types.Int64Value  `xml:"y,attr,omitempty"`
}

func NewTracePoint() *TracePoint {
	ret := &TracePoint{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"tracePt",
		"p14",
	)
	return ret
}

func (m *TracePoint) Clone() openxml.Element {
	ret := NewTracePoint()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.XCoordinate != nil {
		v := *m.XCoordinate
		ret.XCoordinate = &v
	}
	if m.YCoordinate != nil {
		v := *m.YCoordinate
		ret.YCoordinate = &v
	}
	return ret
}

func (m *TracePoint) Validate() error {
	return nil
}

// Defines the TracePointList Class.
type TracePointList struct {
	*openxml.CompositeElementBase
	XMLName    xml.Name    `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main tracePtLst"`
	TracePoint *TracePoint `xml:"tracePt,omitempty"`
}

func NewTracePointList() *TracePointList {
	ret := &TracePointList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tracePtLst",
		"p14",
	)
	return ret
}

func (m *TracePointList) Clone() openxml.Element {
	ret := NewTracePointList()
	if m.TracePoint != nil {
		ret.TracePoint = m.TracePoint.Clone().(*TracePoint)
	}
	return ret
}

func (m *TracePointList) Validate() error {
	if m.TracePoint != nil {
		if err := m.TracePoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TriggerEventRecord Class.
type TriggerEventRecord struct {
	*openxml.LeafElementBase
	XMLName  xml.Name                             `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main triggerEvt"`
	Type     *types.EnumValue[TriggerEventValues] `xml:"type,attr,omitempty"`
	Time     *types.StringValue                   `xml:"time,attr,omitempty"`
	ObjectId *types.UInt32Value                   `xml:"objId,attr,omitempty"`
}

func NewTriggerEventRecord() *TriggerEventRecord {
	ret := &TriggerEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"triggerEvt",
		"p14",
	)
	return ret
}

func (m *TriggerEventRecord) Clone() openxml.Element {
	ret := NewTriggerEventRecord()
	if m.Type != nil {
		v := *m.Type
		ret.Type = &v
	}
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.ObjectId != nil {
		v := *m.ObjectId
		ret.ObjectId = &v
	}
	return ret
}

func (m *TriggerEventRecord) Validate() error {
	return nil
}

// Defines the PlayEventRecord Class.
type PlayEventRecord struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main playEvt"`
}

func NewPlayEventRecord() *PlayEventRecord {
	ret := &PlayEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"playEvt",
		"p14",
	)
	return ret
}

func (m *PlayEventRecord) Clone() openxml.Element {
	ret := NewPlayEventRecord()
	return ret
}

func (m *PlayEventRecord) Validate() error {
	return nil
}

// Defines the StopEventRecord Class.
type StopEventRecord struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main stopEvt"`
}

func NewStopEventRecord() *StopEventRecord {
	ret := &StopEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"stopEvt",
		"p14",
	)
	return ret
}

func (m *StopEventRecord) Clone() openxml.Element {
	ret := NewStopEventRecord()
	return ret
}

func (m *StopEventRecord) Validate() error {
	return nil
}

// Defines the PauseEventRecord Class.
type PauseEventRecord struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main pauseEvt"`
}

func NewPauseEventRecord() *PauseEventRecord {
	ret := &PauseEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pauseEvt",
		"p14",
	)
	return ret
}

func (m *PauseEventRecord) Clone() openxml.Element {
	ret := NewPauseEventRecord()
	return ret
}

func (m *PauseEventRecord) Validate() error {
	return nil
}

// Defines the ResumeEventRecord Class.
type ResumeEventRecord struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main resumeEvt"`
}

func NewResumeEventRecord() *ResumeEventRecord {
	ret := &ResumeEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"resumeEvt",
		"p14",
	)
	return ret
}

func (m *ResumeEventRecord) Clone() openxml.Element {
	ret := NewResumeEventRecord()
	return ret
}

func (m *ResumeEventRecord) Validate() error {
	return nil
}

// Defines the MediaPlaybackEventRecordType Class.
type MediaPlaybackEventRecordType struct {
	*openxml.LeafElementBase
	XMLName  xml.Name           `xml:""`
	Time     *types.StringValue `xml:"time,attr,omitempty"`
	ObjectId *types.UInt32Value `xml:"objId,attr,omitempty"`
}

func NewMediaPlaybackEventRecordType() *MediaPlaybackEventRecordType {
	ret := &MediaPlaybackEventRecordType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p14",
	)
	return ret
}

func (m *MediaPlaybackEventRecordType) Clone() openxml.Element {
	ret := NewMediaPlaybackEventRecordType()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.ObjectId != nil {
		v := *m.ObjectId
		ret.ObjectId = &v
	}
	return ret
}

func (m *MediaPlaybackEventRecordType) Validate() error {
	return nil
}

// Defines the SeekEventRecord Class.
type SeekEventRecord struct {
	*openxml.LeafElementBase
	XMLName  xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main seekEvt"`
	Time     *types.StringValue `xml:"time,attr,omitempty"`
	ObjectId *types.UInt32Value `xml:"objId,attr,omitempty"`
	Seek     *types.StringValue `xml:"seek,attr,omitempty"`
}

func NewSeekEventRecord() *SeekEventRecord {
	ret := &SeekEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"seekEvt",
		"p14",
	)
	return ret
}

func (m *SeekEventRecord) Clone() openxml.Element {
	ret := NewSeekEventRecord()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.ObjectId != nil {
		v := *m.ObjectId
		ret.ObjectId = &v
	}
	if m.Seek != nil {
		v := *m.Seek
		ret.Seek = &v
	}
	return ret
}

func (m *SeekEventRecord) Validate() error {
	return nil
}

// Defines the NullEventRecord Class.
type NullEventRecord struct {
	*openxml.LeafElementBase
	XMLName  xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2010/main nullEvt"`
	Time     *types.StringValue `xml:"time,attr,omitempty"`
	ObjectId *types.UInt32Value `xml:"objId,attr,omitempty"`
}

func NewNullEventRecord() *NullEventRecord {
	ret := &NullEventRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"nullEvt",
		"p14",
	)
	return ret
}

func (m *NullEventRecord) Clone() openxml.Element {
	ret := NewNullEventRecord()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.ObjectId != nil {
		v := *m.ObjectId
		ret.ObjectId = &v
	}
	return ret
}

func (m *NullEventRecord) Validate() error {
	return nil
}

// Defines the PresetTransition Class.
type PresetTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main prstTrans"`
	Preset  *types.StringValue  `xml:"prst,attr,omitempty"`
	InvX    *types.BooleanValue `xml:"invX,attr,omitempty"`
	InvY    *types.BooleanValue `xml:"invY,attr,omitempty"`
}

func NewPresetTransition() *PresetTransition {
	ret := &PresetTransition{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"prstTrans",
		"p15",
	)
	return ret
}

func (m *PresetTransition) Clone() openxml.Element {
	ret := NewPresetTransition()
	if m.Preset != nil {
		v := *m.Preset
		ret.Preset = &v
	}
	if m.InvX != nil {
		v := *m.InvX
		ret.InvX = &v
	}
	if m.InvY != nil {
		v := *m.InvY
		ret.InvY = &v
	}
	return ret
}

func (m *PresetTransition) Validate() error {
	return nil
}

// Defines the PresenceInfo Class.
type PresenceInfo struct {
	*openxml.LeafElementBase
	XMLName    xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main presenceInfo"`
	UserId     *types.StringValue `xml:"userId,attr,omitempty"`
	ProviderId *types.StringValue `xml:"providerId,attr,omitempty"`
}

func NewPresenceInfo() *PresenceInfo {
	ret := &PresenceInfo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"presenceInfo",
		"p15",
	)
	return ret
}

func (m *PresenceInfo) Clone() openxml.Element {
	ret := NewPresenceInfo()
	if m.UserId != nil {
		v := *m.UserId
		ret.UserId = &v
	}
	if m.ProviderId != nil {
		v := *m.ProviderId
		ret.ProviderId = &v
	}
	return ret
}

func (m *PresenceInfo) Validate() error {
	return nil
}

// Defines the ThreadingInfo Class.
type ThreadingInfo struct {
	*openxml.CompositeElementBase
	XMLName                 xml.Name                 `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main threadingInfo"`
	TimeZoneBias            *types.Int32Value        `xml:"timeZoneBias,attr,omitempty"`
	ParentCommentIdentifier *ParentCommentIdentifier `xml:"parentCm,omitempty"`
}

func NewThreadingInfo() *ThreadingInfo {
	ret := &ThreadingInfo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"threadingInfo",
		"p15",
	)
	return ret
}

func (m *ThreadingInfo) Clone() openxml.Element {
	ret := NewThreadingInfo()
	if m.TimeZoneBias != nil {
		v := *m.TimeZoneBias
		ret.TimeZoneBias = &v
	}
	if m.ParentCommentIdentifier != nil {
		ret.ParentCommentIdentifier = m.ParentCommentIdentifier.Clone().(*ParentCommentIdentifier)
	}
	return ret
}

func (m *ThreadingInfo) Validate() error {
	if m.ParentCommentIdentifier != nil {
		if err := m.ParentCommentIdentifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideGuideList Class.
type SlideGuideList struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main sldGuideLst"`
	ExtendedGuide *ExtendedGuide `xml:"guide,omitempty"`
	ExtensionList *ExtensionList `xml:"extLst,omitempty"`
}

func NewSlideGuideList() *SlideGuideList {
	ret := &SlideGuideList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldGuideLst",
		"p15",
	)
	return ret
}

func (m *SlideGuideList) Clone() openxml.Element {
	ret := NewSlideGuideList()
	if m.ExtendedGuide != nil {
		ret.ExtendedGuide = m.ExtendedGuide.Clone().(*ExtendedGuide)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *SlideGuideList) Validate() error {
	if m.ExtendedGuide != nil {
		if err := m.ExtendedGuide.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the NotesGuideList Class.
type NotesGuideList struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main notesGuideLst"`
	ExtendedGuide *ExtendedGuide `xml:"guide,omitempty"`
	ExtensionList *ExtensionList `xml:"extLst,omitempty"`
}

func NewNotesGuideList() *NotesGuideList {
	ret := &NotesGuideList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"notesGuideLst",
		"p15",
	)
	return ret
}

func (m *NotesGuideList) Clone() openxml.Element {
	ret := NewNotesGuideList()
	if m.ExtendedGuide != nil {
		ret.ExtendedGuide = m.ExtendedGuide.Clone().(*ExtendedGuide)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *NotesGuideList) Validate() error {
	if m.ExtendedGuide != nil {
		if err := m.ExtendedGuide.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ExtendedGuideList Class.
type ExtendedGuideList struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:""`
	ExtendedGuide *ExtendedGuide `xml:"guide,omitempty"`
	ExtensionList *ExtensionList `xml:"extLst,omitempty"`
}

func NewExtendedGuideList() *ExtendedGuideList {
	ret := &ExtendedGuideList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p15",
	)
	return ret
}

func (m *ExtendedGuideList) Clone() openxml.Element {
	ret := NewExtendedGuideList()
	if m.ExtendedGuide != nil {
		ret.ExtendedGuide = m.ExtendedGuide.Clone().(*ExtendedGuide)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *ExtendedGuideList) Validate() error {
	if m.ExtendedGuide != nil {
		if err := m.ExtendedGuide.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ChartTrackingReferenceBased Class.
type ChartTrackingReferenceBased struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main chartTrackingRefBased"`
	Val     *types.BooleanValue `xml:"val,attr,omitempty"`
}

func NewChartTrackingReferenceBased() *ChartTrackingReferenceBased {
	ret := &ChartTrackingReferenceBased{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"chartTrackingRefBased",
		"p15",
	)
	return ret
}

func (m *ChartTrackingReferenceBased) Clone() openxml.Element {
	ret := NewChartTrackingReferenceBased()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *ChartTrackingReferenceBased) Validate() error {
	return nil
}

// Defines the ParentCommentIdentifier Class.
type ParentCommentIdentifier struct {
	*openxml.LeafElementBase
	XMLName  xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main parentCm"`
	AuthorId *types.UInt32Value `xml:"authorId,attr,omitempty"`
	Index    *types.UInt32Value `xml:"idx,attr,omitempty"`
}

func NewParentCommentIdentifier() *ParentCommentIdentifier {
	ret := &ParentCommentIdentifier{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"parentCm",
		"p15",
	)
	return ret
}

func (m *ParentCommentIdentifier) Clone() openxml.Element {
	ret := NewParentCommentIdentifier()
	if m.AuthorId != nil {
		v := *m.AuthorId
		ret.AuthorId = &v
	}
	if m.Index != nil {
		v := *m.Index
		ret.Index = &v
	}
	return ret
}

func (m *ParentCommentIdentifier) Validate() error {
	return nil
}

// Defines the ColorType Class.
type ColorType struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main clr"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewColorType() *ColorType {
	ret := &ColorType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"clr",
		"p15",
	)
	return ret
}

func (m *ColorType) Clone() openxml.Element {
	ret := NewColorType()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *ColorType) Validate() error {
	return nil
}

// Defines the ExtendedGuide Class.
type ExtendedGuide struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name                          `xml:"http://schemas.microsoft.com/office/powerpoint/2012/main guide"`
	Id            *types.UInt32Value                `xml:"id,attr,omitempty"`
	Name          *types.StringValue                `xml:"name,attr,omitempty"`
	Orientation   *types.EnumValue[DirectionValues] `xml:"orient,attr,omitempty"`
	Position      *types.Int32Value                 `xml:"pos,attr,omitempty"`
	IsUserDrawn   *types.BooleanValue               `xml:"userDrawn,attr,omitempty"`
	ColorType     *ColorType                        `xml:"clr,omitempty"`
	ExtensionList *ExtensionList                    `xml:"extLst,omitempty"`
}

func NewExtendedGuide() *ExtendedGuide {
	ret := &ExtendedGuide{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2012/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"guide",
		"p15",
	)
	return ret
}

func (m *ExtendedGuide) Clone() openxml.Element {
	ret := NewExtendedGuide()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Orientation != nil {
		v := *m.Orientation
		ret.Orientation = &v
	}
	if m.Position != nil {
		v := *m.Position
		ret.Position = &v
	}
	if m.IsUserDrawn != nil {
		v := *m.IsUserDrawn
		ret.IsUserDrawn = &v
	}
	if m.ColorType != nil {
		ret.ColorType = m.ColorType.Clone().(*ColorType)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *ExtendedGuide) Validate() error {
	if m.ColorType != nil {
		if err := m.ColorType.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the DesignElement Class.
type DesignElement struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2015/main designElem"`
	Val     *types.BooleanValue `xml:"val,attr,omitempty"`
}

func NewDesignElement() *DesignElement {
	ret := &DesignElement{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2015/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"designElem",
		"p16",
	)
	return ret
}

func (m *DesignElement) Clone() openxml.Element {
	ret := NewDesignElement()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *DesignElement) Validate() error {
	return nil
}

// Defines the ReadonlyRecommended Class.
type ReadonlyRecommended struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2017/10/main readonlyRecommended"`
	Val     *types.BooleanValue `xml:"val,attr,omitempty"`
}

func NewReadonlyRecommended() *ReadonlyRecommended {
	ret := &ReadonlyRecommended{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2017/10/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"readonlyRecommended",
		"p1710",
	)
	return ret
}

func (m *ReadonlyRecommended) Clone() openxml.Element {
	ret := NewReadonlyRecommended()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *ReadonlyRecommended) Validate() error {
	return nil
}

// Defines the TracksInfo Class.
type TracksInfo struct {
	*openxml.CompositeElementBase
	XMLName    xml.Name                          `xml:"http://schemas.microsoft.com/office/powerpoint/2017/3/main tracksInfo"`
	DisplayLoc *types.EnumValue[DisplayLocation] `xml:"displayLoc,attr,omitempty"`
	TrackList  *TrackList                        `xml:"trackLst,omitempty"`
}

func NewTracksInfo() *TracksInfo {
	ret := &TracksInfo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2017/3/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tracksInfo",
		"p173",
	)
	return ret
}

func (m *TracksInfo) Clone() openxml.Element {
	ret := NewTracksInfo()
	if m.DisplayLoc != nil {
		v := *m.DisplayLoc
		ret.DisplayLoc = &v
	}
	if m.TrackList != nil {
		ret.TrackList = m.TrackList.Clone().(*TrackList)
	}
	return ret
}

func (m *TracksInfo) Validate() error {
	if m.TrackList != nil {
		if err := m.TrackList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Track Class.
type Track struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2017/3/main track"`
	Id      *types.StringValue `xml:"id,attr,omitempty"`
	Label   *types.StringValue `xml:"label,attr,omitempty"`
	Lang    *types.StringValue `xml:"lang,attr,omitempty"`
	Embed   *types.StringValue `xml:"r:embed,attr,omitempty"`
	Link    *types.StringValue `xml:"r:link,attr,omitempty"`
}

func NewTrack() *Track {
	ret := &Track{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2017/3/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"track",
		"p173",
	)
	return ret
}

func (m *Track) Clone() openxml.Element {
	ret := NewTrack()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.Label != nil {
		v := *m.Label
		ret.Label = &v
	}
	if m.Lang != nil {
		v := *m.Lang
		ret.Lang = &v
	}
	if m.Embed != nil {
		v := *m.Embed
		ret.Embed = &v
	}
	if m.Link != nil {
		v := *m.Link
		ret.Link = &v
	}
	return ret
}

func (m *Track) Validate() error {
	return nil
}

// Defines the TrackList Class.
type TrackList struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2017/3/main trackLst"`
	Track   *Track   `xml:"track,omitempty"`
}

func NewTrackList() *TrackList {
	ret := &TrackList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2017/3/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"trackLst",
		"p173",
	)
	return ret
}

func (m *TrackList) Clone() openxml.Element {
	ret := NewTrackList()
	if m.Track != nil {
		ret.Track = m.Track.Clone().(*Track)
	}
	return ret
}

func (m *TrackList) Validate() error {
	if m.Track != nil {
		if err := m.Track.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ClassificationOutcome Class.
type ClassificationOutcome struct {
	*openxml.LeafElementBase
	XMLName xml.Name                                    `xml:"http://schemas.microsoft.com/office/powerpoint/2018/4/main classification"`
	Val     *types.EnumValue[ClassificationOutcomeType] `xml:"val,attr,omitempty"`
}

func NewClassificationOutcome() *ClassificationOutcome {
	ret := &ClassificationOutcome{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/4/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"classification",
		"p184",
	)
	return ret
}

func (m *ClassificationOutcome) Clone() openxml.Element {
	ret := NewClassificationOutcome()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *ClassificationOutcome) Validate() error {
	return nil
}

// Defines the CommentUnknownAnchor Class.
type CommentUnknownAnchor struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main unknownAnchor"`
}

func NewCommentUnknownAnchor() *CommentUnknownAnchor {
	ret := &CommentUnknownAnchor{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"unknownAnchor",
		"p188",
	)
	return ret
}

func (m *CommentUnknownAnchor) Clone() openxml.Element {
	ret := NewCommentUnknownAnchor()
	return ret
}

func (m *CommentUnknownAnchor) Validate() error {
	return nil
}

// Defines the TextBodyType Class.
type TextBodyType struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main txBody"`
	// Skipped DrawingML type not yet implemented: BodyProperties
	// Skipped DrawingML type not yet implemented: ListStyle
	// Skipped DrawingML type not yet implemented: Paragraph
}

func NewTextBodyType() *TextBodyType {
	ret := &TextBodyType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"txBody",
		"p188",
	)
	return ret
}

func (m *TextBodyType) Clone() openxml.Element {
	ret := NewTextBodyType()
	return ret
}

func (m *TextBodyType) Validate() error {
	return nil
}

// Defines the CommentPropertiesExtensionList Class.
type CommentPropertiesExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                    xml.Name                    `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main extLst"`
	CommentPropertiesExtension *CommentPropertiesExtension `xml:"ext,omitempty"`
}

func NewCommentPropertiesExtensionList() *CommentPropertiesExtensionList {
	ret := &CommentPropertiesExtensionList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p188",
	)
	return ret
}

func (m *CommentPropertiesExtensionList) Clone() openxml.Element {
	ret := NewCommentPropertiesExtensionList()
	if m.CommentPropertiesExtension != nil {
		ret.CommentPropertiesExtension = m.CommentPropertiesExtension.Clone().(*CommentPropertiesExtension)
	}
	return ret
}

func (m *CommentPropertiesExtensionList) Validate() error {
	if m.CommentPropertiesExtension != nil {
		if err := m.CommentPropertiesExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the AuthorList Class.
type AuthorList struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main authorLst"`
	Author  *Author  `xml:"author,omitempty"`
}

func NewAuthorList() *AuthorList {
	ret := &AuthorList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"authorLst",
		"p188",
	)
	return ret
}

func (m *AuthorList) Clone() openxml.Element {
	ret := NewAuthorList()
	if m.Author != nil {
		ret.Author = m.Author.Clone().(*Author)
	}
	return ret
}

func (m *AuthorList) Validate() error {
	if m.Author != nil {
		if err := m.Author.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommentList Class.
type CommentList struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main cmLst"`
	Comment *Comment `xml:"cm,omitempty"`
}

func NewCommentList() *CommentList {
	ret := &CommentList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cmLst",
		"p188",
	)
	return ret
}

func (m *CommentList) Clone() openxml.Element {
	ret := NewCommentList()
	if m.Comment != nil {
		ret.Comment = m.Comment.Clone().(*Comment)
	}
	return ret
}

func (m *CommentList) Validate() error {
	if m.Comment != nil {
		if err := m.Comment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommentRelationship Class.
type CommentRelationship struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main commentRel"`
	Id      *types.StringValue `xml:"r:id,attr,omitempty"`
}

func NewCommentRelationship() *CommentRelationship {
	ret := &CommentRelationship{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"commentRel",
		"p188",
	)
	return ret
}

func (m *CommentRelationship) Clone() openxml.Element {
	ret := NewCommentRelationship()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *CommentRelationship) Validate() error {
	return nil
}

// Defines the Author Class.
type Author struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main author"`
	Id            *types.StringValue `xml:"id,attr,omitempty"`
	Name          *types.StringValue `xml:"name,attr,omitempty"`
	Initials      *types.StringValue `xml:"initials,attr,omitempty"`
	UserId        *types.StringValue `xml:"userId,attr,omitempty"`
	ProviderId    *types.StringValue `xml:"providerId,attr,omitempty"`
	ExtensionList *ExtensionList     `xml:"extLst,omitempty"`
}

func NewAuthor() *Author {
	ret := &Author{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"author",
		"p188",
	)
	return ret
}

func (m *Author) Clone() openxml.Element {
	ret := NewAuthor()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Initials != nil {
		v := *m.Initials
		ret.Initials = &v
	}
	if m.UserId != nil {
		v := *m.UserId
		ret.UserId = &v
	}
	if m.ProviderId != nil {
		v := *m.ProviderId
		ret.ProviderId = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *Author) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommentReply Class.
type CommentReply struct {
	*openxml.CompositeElementBase
	XMLName                        xml.Name                        `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main reply"`
	Id                             *types.StringValue              `xml:"id,attr,omitempty"`
	AuthorId                       *types.StringValue              `xml:"authorId,attr,omitempty"`
	Status                         *types.EnumValue[CommentStatus] `xml:"status,attr,omitempty"`
	Created                        *types.DateTimeValue            `xml:"created,attr,omitempty"`
	Tags                           *types.StringValue              `xml:"tags,attr,omitempty"`
	Likes                          *types.StringValue              `xml:"likes,attr,omitempty"`
	TextBodyType                   *TextBodyType                   `xml:"txBody,omitempty"`
	CommentPropertiesExtensionList *CommentPropertiesExtensionList `xml:"extLst,omitempty"`
}

func NewCommentReply() *CommentReply {
	ret := &CommentReply{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"reply",
		"p188",
	)
	return ret
}

func (m *CommentReply) Clone() openxml.Element {
	ret := NewCommentReply()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.AuthorId != nil {
		v := *m.AuthorId
		ret.AuthorId = &v
	}
	if m.Status != nil {
		v := *m.Status
		ret.Status = &v
	}
	if m.Created != nil {
		v := *m.Created
		ret.Created = &v
	}
	if m.Tags != nil {
		v := *m.Tags
		ret.Tags = &v
	}
	if m.Likes != nil {
		v := *m.Likes
		ret.Likes = &v
	}
	if m.TextBodyType != nil {
		ret.TextBodyType = m.TextBodyType.Clone().(*TextBodyType)
	}
	if m.CommentPropertiesExtensionList != nil {
		ret.CommentPropertiesExtensionList = m.CommentPropertiesExtensionList.Clone().(*CommentPropertiesExtensionList)
	}
	return ret
}

func (m *CommentReply) Validate() error {
	if m.TextBodyType != nil {
		if err := m.TextBodyType.Validate(); err != nil {
			return err
		}
	}
	if m.CommentPropertiesExtensionList != nil {
		if err := m.CommentPropertiesExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Point2DType Class.
type Point2DType struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main pos"`
	X       *types.Int64Value `xml:"x,attr,omitempty"`
	Y       *types.Int64Value `xml:"y,attr,omitempty"`
}

func NewPoint2DType() *Point2DType {
	ret := &Point2DType{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pos",
		"p188",
	)
	return ret
}

func (m *Point2DType) Clone() openxml.Element {
	ret := NewPoint2DType()
	if m.X != nil {
		v := *m.X
		ret.X = &v
	}
	if m.Y != nil {
		v := *m.Y
		ret.Y = &v
	}
	return ret
}

func (m *Point2DType) Validate() error {
	return nil
}

// Defines the CommentReplyList Class.
type CommentReplyList struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name      `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main replyLst"`
	CommentReply *CommentReply `xml:"reply,omitempty"`
}

func NewCommentReplyList() *CommentReplyList {
	ret := &CommentReplyList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"replyLst",
		"p188",
	)
	return ret
}

func (m *CommentReplyList) Clone() openxml.Element {
	ret := NewCommentReplyList()
	if m.CommentReply != nil {
		ret.CommentReply = m.CommentReply.Clone().(*CommentReply)
	}
	return ret
}

func (m *CommentReplyList) Validate() error {
	if m.CommentReply != nil {
		if err := m.CommentReply.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Comment Class.
type Comment struct {
	*openxml.CompositeElementBase
	XMLName    xml.Name                        `xml:"http://schemas.microsoft.com/office/powerpoint/2018/8/main cm"`
	Id         *types.StringValue              `xml:"id,attr,omitempty"`
	AuthorId   *types.StringValue              `xml:"authorId,attr,omitempty"`
	Status     *types.EnumValue[CommentStatus] `xml:"status,attr,omitempty"`
	Created    *types.DateTimeValue            `xml:"created,attr,omitempty"`
	Tags       *types.StringValue              `xml:"tags,attr,omitempty"`
	Likes      *types.StringValue              `xml:"likes,attr,omitempty"`
	StartDate  *types.DateTimeValue            `xml:"startDate,attr,omitempty"`
	DueDate    *types.DateTimeValue            `xml:"dueDate,attr,omitempty"`
	AssignedTo *types.StringValue              `xml:"assignedTo,attr,omitempty"`
	Complete   *types.Int32Value               `xml:"complete,attr,omitempty"`
	Priority   *types.UInt32Value              `xml:"priority,attr,omitempty"`
	Title      *types.StringValue              `xml:"title,attr,omitempty"`
	// Unknown Child: Unknown (pc:CT_SlideMonikerList/pc:sldMkLst)
	// Unknown Child: Unknown (pc:CT_SlideLayoutMonikerList/pc:sldLayoutMkLst)
	// Unknown Child: Unknown (pc:CT_MainMasterMonikerList/pc:sldMasterMkLst)
	// Unknown Child: Unknown (oac:CT_DrawingElementMonikerList/oac:deMkLst)
	// Unknown Child: Unknown (oac:CT_TextBodyMonikerList/oac:txBodyMkLst)
	// Unknown Child: Unknown (oac:CT_TextCharRangeMonikerList/oac:txMkLst)
	// Unknown Child: Unknown (oac:CT_TableCellMonikerList/oac:tcMkLst)
	// Unknown Child: Unknown (oac:CT_TableRowMonikerList/oac:trMkLst)
	// Unknown Child: Unknown (oac:CT_TableColumnMonikerList/oac:gridColMkLst)
	CommentUnknownAnchor           *CommentUnknownAnchor           `xml:"unknownAnchor,omitempty"`
	Point2DType                    *Point2DType                    `xml:"pos,omitempty"`
	CommentReplyList               *CommentReplyList               `xml:"replyLst,omitempty"`
	TextBodyType                   *TextBodyType                   `xml:"txBody,omitempty"`
	CommentPropertiesExtensionList *CommentPropertiesExtensionList `xml:"extLst,omitempty"`
}

func NewComment() *Comment {
	ret := &Comment{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2018/8/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cm",
		"p188",
	)
	return ret
}

func (m *Comment) Clone() openxml.Element {
	ret := NewComment()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.AuthorId != nil {
		v := *m.AuthorId
		ret.AuthorId = &v
	}
	if m.Status != nil {
		v := *m.Status
		ret.Status = &v
	}
	if m.Created != nil {
		v := *m.Created
		ret.Created = &v
	}
	if m.Tags != nil {
		v := *m.Tags
		ret.Tags = &v
	}
	if m.Likes != nil {
		v := *m.Likes
		ret.Likes = &v
	}
	if m.StartDate != nil {
		v := *m.StartDate
		ret.StartDate = &v
	}
	if m.DueDate != nil {
		v := *m.DueDate
		ret.DueDate = &v
	}
	if m.AssignedTo != nil {
		v := *m.AssignedTo
		ret.AssignedTo = &v
	}
	if m.Complete != nil {
		v := *m.Complete
		ret.Complete = &v
	}
	if m.Priority != nil {
		v := *m.Priority
		ret.Priority = &v
	}
	if m.Title != nil {
		v := *m.Title
		ret.Title = &v
	}
	if m.CommentUnknownAnchor != nil {
		ret.CommentUnknownAnchor = m.CommentUnknownAnchor.Clone().(*CommentUnknownAnchor)
	}
	if m.Point2DType != nil {
		ret.Point2DType = m.Point2DType.Clone().(*Point2DType)
	}
	if m.CommentReplyList != nil {
		ret.CommentReplyList = m.CommentReplyList.Clone().(*CommentReplyList)
	}
	if m.TextBodyType != nil {
		ret.TextBodyType = m.TextBodyType.Clone().(*TextBodyType)
	}
	if m.CommentPropertiesExtensionList != nil {
		ret.CommentPropertiesExtensionList = m.CommentPropertiesExtensionList.Clone().(*CommentPropertiesExtensionList)
	}
	return ret
}

func (m *Comment) Validate() error {
	if m.CommentUnknownAnchor != nil {
		if err := m.CommentUnknownAnchor.Validate(); err != nil {
			return err
		}
	}
	if m.Point2DType != nil {
		if err := m.Point2DType.Validate(); err != nil {
			return err
		}
	}
	if m.CommentReplyList != nil {
		if err := m.CommentReplyList.Validate(); err != nil {
			return err
		}
	}
	if m.TextBodyType != nil {
		if err := m.TextBodyType.Validate(); err != nil {
			return err
		}
	}
	if m.CommentPropertiesExtensionList != nil {
		if err := m.CommentPropertiesExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TaskHistoryDetails Class.
type TaskHistoryDetails struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main taskHistoryDetails"`
	Id            *types.StringValue `xml:"id,attr,omitempty"`
	TaskHistory   *TaskHistory       `xml:"history,omitempty"`
	ExtensionList *ExtensionList     `xml:"extLst,omitempty"`
}

func NewTaskHistoryDetails() *TaskHistoryDetails {
	ret := &TaskHistoryDetails{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"taskHistoryDetails",
		"p1912",
	)
	return ret
}

func (m *TaskHistoryDetails) Clone() openxml.Element {
	ret := NewTaskHistoryDetails()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.TaskHistory != nil {
		ret.TaskHistory = m.TaskHistory.Clone().(*TaskHistory)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *TaskHistoryDetails) Validate() error {
	if m.TaskHistory != nil {
		if err := m.TaskHistory.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommentAnchor Class.
type CommentAnchor struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main comment"`
	Id      *types.StringValue `xml:"id,attr,omitempty"`
}

func NewCommentAnchor() *CommentAnchor {
	ret := &CommentAnchor{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"comment",
		"p1912",
	)
	return ret
}

func (m *CommentAnchor) Clone() openxml.Element {
	ret := NewCommentAnchor()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *CommentAnchor) Validate() error {
	return nil
}

// Defines the AtrbtnTaskAssignUnassignUser Class.
type AtrbtnTaskAssignUnassignUser struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main atrbtn"`
}

func NewAtrbtnTaskAssignUnassignUser() *AtrbtnTaskAssignUnassignUser {
	ret := &AtrbtnTaskAssignUnassignUser{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"atrbtn",
		"p1912",
	)
	return ret
}

func (m *AtrbtnTaskAssignUnassignUser) Clone() openxml.Element {
	ret := NewAtrbtnTaskAssignUnassignUser()
	return ret
}

func (m *AtrbtnTaskAssignUnassignUser) Validate() error {
	return nil
}

// Defines the AsgnTaskAssignUnassignUser Class.
type AsgnTaskAssignUnassignUser struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main asgn"`
}

func NewAsgnTaskAssignUnassignUser() *AsgnTaskAssignUnassignUser {
	ret := &AsgnTaskAssignUnassignUser{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"asgn",
		"p1912",
	)
	return ret
}

func (m *AsgnTaskAssignUnassignUser) Clone() openxml.Element {
	ret := NewAsgnTaskAssignUnassignUser()
	return ret
}

func (m *AsgnTaskAssignUnassignUser) Validate() error {
	return nil
}

// Defines the UnAsgnTaskAssignUnassignUser Class.
type UnAsgnTaskAssignUnassignUser struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main unAsgn"`
}

func NewUnAsgnTaskAssignUnassignUser() *UnAsgnTaskAssignUnassignUser {
	ret := &UnAsgnTaskAssignUnassignUser{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"unAsgn",
		"p1912",
	)
	return ret
}

func (m *UnAsgnTaskAssignUnassignUser) Clone() openxml.Element {
	ret := NewUnAsgnTaskAssignUnassignUser()
	return ret
}

func (m *UnAsgnTaskAssignUnassignUser) Validate() error {
	return nil
}

// Defines the OpenXmlTaskAssignUnassignUserElement Class.
type OpenXmlTaskAssignUnassignUserElement struct {
	*openxml.LeafElementBase
	XMLName  xml.Name           `xml:""`
	AuthorId *types.StringValue `xml:"authorId,attr,omitempty"`
}

func NewOpenXmlTaskAssignUnassignUserElement() *OpenXmlTaskAssignUnassignUserElement {
	ret := &OpenXmlTaskAssignUnassignUserElement{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p1912",
	)
	return ret
}

func (m *OpenXmlTaskAssignUnassignUserElement) Clone() openxml.Element {
	ret := NewOpenXmlTaskAssignUnassignUserElement()
	if m.AuthorId != nil {
		v := *m.AuthorId
		ret.AuthorId = &v
	}
	return ret
}

func (m *OpenXmlTaskAssignUnassignUserElement) Validate() error {
	return nil
}

// Defines the TaskAnchor Class.
type TaskAnchor struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main anchr"`
	CommentAnchor *CommentAnchor `xml:"comment,omitempty"`
	ExtensionList *ExtensionList `xml:"extLst,omitempty"`
}

func NewTaskAnchor() *TaskAnchor {
	ret := &TaskAnchor{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"anchr",
		"p1912",
	)
	return ret
}

func (m *TaskAnchor) Clone() openxml.Element {
	ret := NewTaskAnchor()
	if m.CommentAnchor != nil {
		ret.CommentAnchor = m.CommentAnchor.Clone().(*CommentAnchor)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *TaskAnchor) Validate() error {
	if m.CommentAnchor != nil {
		if err := m.CommentAnchor.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the AddEmpty Class.
type AddEmpty struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main add"`
}

func NewAddEmpty() *AddEmpty {
	ret := &AddEmpty{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"add",
		"p1912",
	)
	return ret
}

func (m *AddEmpty) Clone() openxml.Element {
	ret := NewAddEmpty()
	return ret
}

func (m *AddEmpty) Validate() error {
	return nil
}

// Defines the UnasgnAllEmpty Class.
type UnasgnAllEmpty struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main unasgnAll"`
}

func NewUnasgnAllEmpty() *UnasgnAllEmpty {
	ret := &UnasgnAllEmpty{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"unasgnAll",
		"p1912",
	)
	return ret
}

func (m *UnasgnAllEmpty) Clone() openxml.Element {
	ret := NewUnasgnAllEmpty()
	return ret
}

func (m *UnasgnAllEmpty) Validate() error {
	return nil
}

// Defines the TaskTitleEventInfo Class.
type TaskTitleEventInfo struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main title"`
	Val     *types.StringValue `xml:"val,attr,omitempty"`
}

func NewTaskTitleEventInfo() *TaskTitleEventInfo {
	ret := &TaskTitleEventInfo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"title",
		"p1912",
	)
	return ret
}

func (m *TaskTitleEventInfo) Clone() openxml.Element {
	ret := NewTaskTitleEventInfo()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *TaskTitleEventInfo) Validate() error {
	return nil
}

// Defines the TaskScheduleEventInfo Class.
type TaskScheduleEventInfo struct {
	*openxml.LeafElementBase
	XMLName xml.Name             `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main date"`
	StDt    *types.DateTimeValue `xml:"stDt,attr,omitempty"`
	EndDt   *types.DateTimeValue `xml:"endDt,attr,omitempty"`
}

func NewTaskScheduleEventInfo() *TaskScheduleEventInfo {
	ret := &TaskScheduleEventInfo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"date",
		"p1912",
	)
	return ret
}

func (m *TaskScheduleEventInfo) Clone() openxml.Element {
	ret := NewTaskScheduleEventInfo()
	if m.StDt != nil {
		v := *m.StDt
		ret.StDt = &v
	}
	if m.EndDt != nil {
		v := *m.EndDt
		ret.EndDt = &v
	}
	return ret
}

func (m *TaskScheduleEventInfo) Validate() error {
	return nil
}

// Defines the TaskProgressEventInfo Class.
type TaskProgressEventInfo struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main pcntCmplt"`
	Val     *types.Int32Value `xml:"val,attr,omitempty"`
}

func NewTaskProgressEventInfo() *TaskProgressEventInfo {
	ret := &TaskProgressEventInfo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pcntCmplt",
		"p1912",
	)
	return ret
}

func (m *TaskProgressEventInfo) Clone() openxml.Element {
	ret := NewTaskProgressEventInfo()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *TaskProgressEventInfo) Validate() error {
	return nil
}

// Defines the TaskPriorityRecord Class.
type TaskPriorityRecord struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main pri"`
	Val     *types.Int32Value `xml:"val,attr,omitempty"`
}

func NewTaskPriorityRecord() *TaskPriorityRecord {
	ret := &TaskPriorityRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pri",
		"p1912",
	)
	return ret
}

func (m *TaskPriorityRecord) Clone() openxml.Element {
	ret := NewTaskPriorityRecord()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *TaskPriorityRecord) Validate() error {
	return nil
}

// Defines the TaskUndo Class.
type TaskUndo struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main undo"`
	Id      *types.StringValue `xml:"id,attr,omitempty"`
}

func NewTaskUndo() *TaskUndo {
	ret := &TaskUndo{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"undo",
		"p1912",
	)
	return ret
}

func (m *TaskUndo) Clone() openxml.Element {
	ret := NewTaskUndo()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *TaskUndo) Validate() error {
	return nil
}

// Defines the TaskUnknownRecord Class.
type TaskUnknownRecord struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main unknown"`
}

func NewTaskUnknownRecord() *TaskUnknownRecord {
	ret := &TaskUnknownRecord{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"unknown",
		"p1912",
	)
	return ret
}

func (m *TaskUnknownRecord) Clone() openxml.Element {
	ret := NewTaskUnknownRecord()
	return ret
}

func (m *TaskUnknownRecord) Validate() error {
	return nil
}

// Defines the TaskHistoryEvent Class.
type TaskHistoryEvent struct {
	*openxml.CompositeElementBase
	XMLName                      xml.Name                      `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main event"`
	Time                         *types.DateTimeValue          `xml:"time,attr,omitempty"`
	Id                           *types.StringValue            `xml:"id,attr,omitempty"`
	AtrbtnTaskAssignUnassignUser *AtrbtnTaskAssignUnassignUser `xml:"atrbtn,omitempty"`
	TaskAnchor                   *TaskAnchor                   `xml:"anchr,omitempty"`
	AsgnTaskAssignUnassignUser   *AsgnTaskAssignUnassignUser   `xml:"asgn,omitempty"`
	UnAsgnTaskAssignUnassignUser *UnAsgnTaskAssignUnassignUser `xml:"unAsgn,omitempty"`
	AddEmpty                     *AddEmpty                     `xml:"add,omitempty"`
	TaskTitleEventInfo           *TaskTitleEventInfo           `xml:"title,omitempty"`
	TaskScheduleEventInfo        *TaskScheduleEventInfo        `xml:"date,omitempty"`
	TaskProgressEventInfo        *TaskProgressEventInfo        `xml:"pcntCmplt,omitempty"`
	TaskPriorityRecord           *TaskPriorityRecord           `xml:"pri,omitempty"`
	UnasgnAllEmpty               *UnasgnAllEmpty               `xml:"unasgnAll,omitempty"`
	TaskUndo                     *TaskUndo                     `xml:"undo,omitempty"`
	TaskUnknownRecord            *TaskUnknownRecord            `xml:"unknown,omitempty"`
	ExtensionList                *ExtensionList                `xml:"extLst,omitempty"`
}

func NewTaskHistoryEvent() *TaskHistoryEvent {
	ret := &TaskHistoryEvent{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"event",
		"p1912",
	)
	return ret
}

func (m *TaskHistoryEvent) Clone() openxml.Element {
	ret := NewTaskHistoryEvent()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.AtrbtnTaskAssignUnassignUser != nil {
		ret.AtrbtnTaskAssignUnassignUser = m.AtrbtnTaskAssignUnassignUser.Clone().(*AtrbtnTaskAssignUnassignUser)
	}
	if m.TaskAnchor != nil {
		ret.TaskAnchor = m.TaskAnchor.Clone().(*TaskAnchor)
	}
	if m.AsgnTaskAssignUnassignUser != nil {
		ret.AsgnTaskAssignUnassignUser = m.AsgnTaskAssignUnassignUser.Clone().(*AsgnTaskAssignUnassignUser)
	}
	if m.UnAsgnTaskAssignUnassignUser != nil {
		ret.UnAsgnTaskAssignUnassignUser = m.UnAsgnTaskAssignUnassignUser.Clone().(*UnAsgnTaskAssignUnassignUser)
	}
	if m.AddEmpty != nil {
		ret.AddEmpty = m.AddEmpty.Clone().(*AddEmpty)
	}
	if m.TaskTitleEventInfo != nil {
		ret.TaskTitleEventInfo = m.TaskTitleEventInfo.Clone().(*TaskTitleEventInfo)
	}
	if m.TaskScheduleEventInfo != nil {
		ret.TaskScheduleEventInfo = m.TaskScheduleEventInfo.Clone().(*TaskScheduleEventInfo)
	}
	if m.TaskProgressEventInfo != nil {
		ret.TaskProgressEventInfo = m.TaskProgressEventInfo.Clone().(*TaskProgressEventInfo)
	}
	if m.TaskPriorityRecord != nil {
		ret.TaskPriorityRecord = m.TaskPriorityRecord.Clone().(*TaskPriorityRecord)
	}
	if m.UnasgnAllEmpty != nil {
		ret.UnasgnAllEmpty = m.UnasgnAllEmpty.Clone().(*UnasgnAllEmpty)
	}
	if m.TaskUndo != nil {
		ret.TaskUndo = m.TaskUndo.Clone().(*TaskUndo)
	}
	if m.TaskUnknownRecord != nil {
		ret.TaskUnknownRecord = m.TaskUnknownRecord.Clone().(*TaskUnknownRecord)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *TaskHistoryEvent) Validate() error {
	if m.AtrbtnTaskAssignUnassignUser != nil {
		if err := m.AtrbtnTaskAssignUnassignUser.Validate(); err != nil {
			return err
		}
	}
	if m.TaskAnchor != nil {
		if err := m.TaskAnchor.Validate(); err != nil {
			return err
		}
	}
	if m.AsgnTaskAssignUnassignUser != nil {
		if err := m.AsgnTaskAssignUnassignUser.Validate(); err != nil {
			return err
		}
	}
	if m.UnAsgnTaskAssignUnassignUser != nil {
		if err := m.UnAsgnTaskAssignUnassignUser.Validate(); err != nil {
			return err
		}
	}
	if m.AddEmpty != nil {
		if err := m.AddEmpty.Validate(); err != nil {
			return err
		}
	}
	if m.TaskTitleEventInfo != nil {
		if err := m.TaskTitleEventInfo.Validate(); err != nil {
			return err
		}
	}
	if m.TaskScheduleEventInfo != nil {
		if err := m.TaskScheduleEventInfo.Validate(); err != nil {
			return err
		}
	}
	if m.TaskProgressEventInfo != nil {
		if err := m.TaskProgressEventInfo.Validate(); err != nil {
			return err
		}
	}
	if m.TaskPriorityRecord != nil {
		if err := m.TaskPriorityRecord.Validate(); err != nil {
			return err
		}
	}
	if m.UnasgnAllEmpty != nil {
		if err := m.UnasgnAllEmpty.Validate(); err != nil {
			return err
		}
	}
	if m.TaskUndo != nil {
		if err := m.TaskUndo.Validate(); err != nil {
			return err
		}
	}
	if m.TaskUnknownRecord != nil {
		if err := m.TaskUnknownRecord.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TaskHistory Class.
type TaskHistory struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name          `xml:"http://schemas.microsoft.com/office/powerpoint/2019/12/main history"`
	TaskHistoryEvent *TaskHistoryEvent `xml:"event,omitempty"`
}

func NewTaskHistory() *TaskHistory {
	ret := &TaskHistory{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2019/12/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"history",
		"p1912",
	)
	return ret
}

func (m *TaskHistory) Clone() openxml.Element {
	ret := NewTaskHistory()
	if m.TaskHistoryEvent != nil {
		ret.TaskHistoryEvent = m.TaskHistoryEvent.Clone().(*TaskHistoryEvent)
	}
	return ret
}

func (m *TaskHistory) Validate() error {
	if m.TaskHistoryEvent != nil {
		if err := m.TaskHistoryEvent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the DesignerTagList Class.
type DesignerTagList struct {
	*openxml.CompositeElementBase
	XMLName     xml.Name     `xml:"http://schemas.microsoft.com/office/powerpoint/2020/02/main designTagLst"`
	DesignerTag *DesignerTag `xml:"designTag,omitempty"`
}

func NewDesignerTagList() *DesignerTagList {
	ret := &DesignerTagList{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2020/02/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"designTagLst",
		"p202",
	)
	return ret
}

func (m *DesignerTagList) Clone() openxml.Element {
	ret := NewDesignerTagList()
	if m.DesignerTag != nil {
		ret.DesignerTag = m.DesignerTag.Clone().(*DesignerTag)
	}
	return ret
}

func (m *DesignerTagList) Validate() error {
	if m.DesignerTag != nil {
		if err := m.DesignerTag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the DesignerDrawingProps Class.
type DesignerDrawingProps struct {
	*openxml.CompositeElementBase
	XMLName         xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2020/02/main designPr"`
	EdtDesignElem   *types.BooleanValue `xml:"edtDesignElem,attr,omitempty"`
	DesignerTagList *DesignerTagList    `xml:"designTagLst,omitempty"`
	ExtensionList   *ExtensionList      `xml:"extLst,omitempty"`
}

func NewDesignerDrawingProps() *DesignerDrawingProps {
	ret := &DesignerDrawingProps{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2020/02/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"designPr",
		"p202",
	)
	return ret
}

func (m *DesignerDrawingProps) Clone() openxml.Element {
	ret := NewDesignerDrawingProps()
	if m.EdtDesignElem != nil {
		v := *m.EdtDesignElem
		ret.EdtDesignElem = &v
	}
	if m.DesignerTagList != nil {
		ret.DesignerTagList = m.DesignerTagList.Clone().(*DesignerTagList)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *DesignerDrawingProps) Validate() error {
	if m.DesignerTagList != nil {
		if err := m.DesignerTagList.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the DesignerTag Class.
type DesignerTag struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2020/02/main designTag"`
	Name    *types.StringValue `xml:"name,attr,omitempty"`
	Val     *types.StringValue `xml:"val,attr,omitempty"`
}

func NewDesignerTag() *DesignerTag {
	ret := &DesignerTag{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2020/02/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"designTag",
		"p202",
	)
	return ret
}

func (m *DesignerTag) Clone() openxml.Element {
	ret := NewDesignerTag()
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *DesignerTag) Validate() error {
	return nil
}

// Defines the Reactions Class.
type Reactions struct {
	*openxml.CompositeElementBase
	XMLName  xml.Name  `xml:"http://schemas.microsoft.com/office/powerpoint/2022/03/main reactions"`
	Reaction *Reaction `xml:"rxn,omitempty"`
}

func NewReactions() *Reactions {
	ret := &Reactions{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2022/03/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"reactions",
		"p223",
	)
	return ret
}

func (m *Reactions) Clone() openxml.Element {
	ret := NewReactions()
	if m.Reaction != nil {
		ret.Reaction = m.Reaction.Clone().(*Reaction)
	}
	return ret
}

func (m *Reactions) Validate() error {
	if m.Reaction != nil {
		if err := m.Reaction.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ReactionInstance Class.
type ReactionInstance struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name             `xml:"http://schemas.microsoft.com/office/powerpoint/2022/03/main instance"`
	Time          *types.DateTimeValue `xml:"time,attr,omitempty"`
	AuthorId      *types.StringValue   `xml:"authorId,attr,omitempty"`
	ExtensionList *ExtensionList       `xml:"extLst,omitempty"`
}

func NewReactionInstance() *ReactionInstance {
	ret := &ReactionInstance{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2022/03/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"instance",
		"p223",
	)
	return ret
}

func (m *ReactionInstance) Clone() openxml.Element {
	ret := NewReactionInstance()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.AuthorId != nil {
		v := *m.AuthorId
		ret.AuthorId = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *ReactionInstance) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Reaction Class.
type Reaction struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name           `xml:"http://schemas.microsoft.com/office/powerpoint/2022/03/main rxn"`
	Type             *types.StringValue `xml:"type,attr,omitempty"`
	ReactionInstance *ReactionInstance  `xml:"instance,omitempty"`
}

func NewReaction() *Reaction {
	ret := &Reaction{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2022/03/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"rxn",
		"p223",
	)
	return ret
}

func (m *Reaction) Clone() openxml.Element {
	ret := NewReaction()
	if m.Type != nil {
		v := *m.Type
		ret.Type = &v
	}
	if m.ReactionInstance != nil {
		ret.ReactionInstance = m.ReactionInstance.Clone().(*ReactionInstance)
	}
	return ret
}

func (m *Reaction) Validate() error {
	if m.ReactionInstance != nil {
		if err := m.ReactionInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TaskDetails Class.
type TaskDetails struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2022/08/main taskDetails"`
	Deleted       *types.BooleanValue `xml:"deleted,attr,omitempty"`
	Inactive      *types.BooleanValue `xml:"inactive,attr,omitempty"`
	TaskHistory   *TaskHistory        `xml:"history,omitempty"`
	ExtensionList *ExtensionList      `xml:"extLst,omitempty"`
}

func NewTaskDetails() *TaskDetails {
	ret := &TaskDetails{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2022/08/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"taskDetails",
		"p228",
	)
	return ret
}

func (m *TaskDetails) Clone() openxml.Element {
	ret := NewTaskDetails()
	if m.Deleted != nil {
		v := *m.Deleted
		ret.Deleted = &v
	}
	if m.Inactive != nil {
		v := *m.Inactive
		ret.Inactive = &v
	}
	if m.TaskHistory != nil {
		ret.TaskHistory = m.TaskHistory.Clone().(*TaskHistory)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *TaskDetails) Validate() error {
	if m.TaskHistory != nil {
		if err := m.TaskHistory.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the PlaceholderTypeExtension Class.
type PlaceholderTypeExtension struct {
	*openxml.CompositeElementBase
	XMLName            xml.Name            `xml:"http://schemas.microsoft.com/office/powerpoint/2023/02/main phTypeExt"`
	PlaceholderTypeACB *PlaceholderTypeACB `xml:"type,omitempty"`
}

func NewPlaceholderTypeExtension() *PlaceholderTypeExtension {
	ret := &PlaceholderTypeExtension{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2023/02/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"phTypeExt",
		"p232",
	)
	return ret
}

func (m *PlaceholderTypeExtension) Clone() openxml.Element {
	ret := NewPlaceholderTypeExtension()
	if m.PlaceholderTypeACB != nil {
		ret.PlaceholderTypeACB = m.PlaceholderTypeACB.Clone().(*PlaceholderTypeACB)
	}
	return ret
}

func (m *PlaceholderTypeExtension) Validate() error {
	if m.PlaceholderTypeACB != nil {
		if err := m.PlaceholderTypeACB.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CameoEmpty Class.
type CameoEmpty struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2023/02/main cameo"`
}

func NewCameoEmpty() *CameoEmpty {
	ret := &CameoEmpty{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2023/02/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"cameo",
		"p232",
	)
	return ret
}

func (m *CameoEmpty) Clone() openxml.Element {
	ret := NewCameoEmpty()
	return ret
}

func (m *CameoEmpty) Validate() error {
	return nil
}

// Defines the UnknownEmpty Class.
type UnknownEmpty struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/powerpoint/2023/02/main unknown"`
}

func NewUnknownEmpty() *UnknownEmpty {
	ret := &UnknownEmpty{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2023/02/main"
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"unknown",
		"p232",
	)
	return ret
}

func (m *UnknownEmpty) Clone() openxml.Element {
	ret := NewUnknownEmpty()
	return ret
}

func (m *UnknownEmpty) Validate() error {
	return nil
}

// Defines the PlaceholderTypeACB Class.
type PlaceholderTypeACB struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name      `xml:"http://schemas.microsoft.com/office/powerpoint/2023/02/main type"`
	CameoEmpty   *CameoEmpty   `xml:"cameo,omitempty"`
	UnknownEmpty *UnknownEmpty `xml:"unknown,omitempty"`
}

func NewPlaceholderTypeACB() *PlaceholderTypeACB {
	ret := &PlaceholderTypeACB{}
	ns := "http://schemas.microsoft.com/office/powerpoint/2023/02/main"
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"type",
		"p232",
	)
	return ret
}

func (m *PlaceholderTypeACB) Clone() openxml.Element {
	ret := NewPlaceholderTypeACB()
	if m.CameoEmpty != nil {
		ret.CameoEmpty = m.CameoEmpty.Clone().(*CameoEmpty)
	}
	if m.UnknownEmpty != nil {
		ret.UnknownEmpty = m.UnknownEmpty.Clone().(*UnknownEmpty)
	}
	return ret
}

func (m *PlaceholderTypeACB) Validate() error {
	if m.CameoEmpty != nil {
		if err := m.CameoEmpty.Validate(); err != nil {
			return err
		}
	}
	if m.UnknownEmpty != nil {
		if err := m.UnknownEmpty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// All Slides.
type SlideAll struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldAll"`
}

func NewSlideAll() *SlideAll {
	ret := &SlideAll{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sldAll",
		"p",
	)
	return ret
}

func (m *SlideAll) Clone() openxml.Element {
	ret := NewSlideAll()
	return ret
}

func (m *SlideAll) Validate() error {
	return nil
}

// Presenter Slide Show Mode.
type PresenterSlideMode struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main present"`
}

func NewPresenterSlideMode() *PresenterSlideMode {
	ret := &PresenterSlideMode{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"present",
		"p",
	)
	return ret
}

func (m *PresenterSlideMode) Clone() openxml.Element {
	ret := NewPresenterSlideMode()
	return ret
}

func (m *PresenterSlideMode) Validate() error {
	return nil
}

// Stop Sound Action.
type EndSoundAction struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main endSnd"`
}

func NewEndSoundAction() *EndSoundAction {
	ret := &EndSoundAction{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"endSnd",
		"p",
	)
	return ret
}

func (m *EndSoundAction) Clone() openxml.Element {
	ret := NewEndSoundAction()
	return ret
}

func (m *EndSoundAction) Validate() error {
	return nil
}

// Build As One.
type BuildAsOne struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bldAsOne"`
}

func NewBuildAsOne() *BuildAsOne {
	ret := &BuildAsOne{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bldAsOne",
		"p",
	)
	return ret
}

func (m *BuildAsOne) Clone() openxml.Element {
	ret := NewBuildAsOne()
	return ret
}

func (m *BuildAsOne) Validate() error {
	return nil
}

// Slide Target.
type SlideTarget struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldTgt"`
}

func NewSlideTarget() *SlideTarget {
	ret := &SlideTarget{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sldTgt",
		"p",
	)
	return ret
}

func (m *SlideTarget) Clone() openxml.Element {
	ret := NewSlideTarget()
	return ret
}

func (m *SlideTarget) Validate() error {
	return nil
}

// Background.
type BackgroundAnimation struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bg"`
}

func NewBackgroundAnimation() *BackgroundAnimation {
	ret := &BackgroundAnimation{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bg",
		"p",
	)
	return ret
}

func (m *BackgroundAnimation) Clone() openxml.Element {
	ret := NewBackgroundAnimation()
	return ret
}

func (m *BackgroundAnimation) Validate() error {
	return nil
}

// Defines the CircleTransition Class.
type CircleTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main circle"`
}

func NewCircleTransition() *CircleTransition {
	ret := &CircleTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"circle",
		"p",
	)
	return ret
}

func (m *CircleTransition) Clone() openxml.Element {
	ret := NewCircleTransition()
	return ret
}

func (m *CircleTransition) Validate() error {
	return nil
}

// Defines the DissolveTransition Class.
type DissolveTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main dissolve"`
}

func NewDissolveTransition() *DissolveTransition {
	ret := &DissolveTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"dissolve",
		"p",
	)
	return ret
}

func (m *DissolveTransition) Clone() openxml.Element {
	ret := NewDissolveTransition()
	return ret
}

func (m *DissolveTransition) Validate() error {
	return nil
}

// Defines the DiamondTransition Class.
type DiamondTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main diamond"`
}

func NewDiamondTransition() *DiamondTransition {
	ret := &DiamondTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"diamond",
		"p",
	)
	return ret
}

func (m *DiamondTransition) Clone() openxml.Element {
	ret := NewDiamondTransition()
	return ret
}

func (m *DiamondTransition) Validate() error {
	return nil
}

// Defines the NewsflashTransition Class.
type NewsflashTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main newsflash"`
}

func NewNewsflashTransition() *NewsflashTransition {
	ret := &NewsflashTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"newsflash",
		"p",
	)
	return ret
}

func (m *NewsflashTransition) Clone() openxml.Element {
	ret := NewNewsflashTransition()
	return ret
}

func (m *NewsflashTransition) Validate() error {
	return nil
}

// Defines the PlusTransition Class.
type PlusTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main plus"`
}

func NewPlusTransition() *PlusTransition {
	ret := &PlusTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"plus",
		"p",
	)
	return ret
}

func (m *PlusTransition) Clone() openxml.Element {
	ret := NewPlusTransition()
	return ret
}

func (m *PlusTransition) Validate() error {
	return nil
}

// Defines the RandomTransition Class.
type RandomTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main random"`
}

func NewRandomTransition() *RandomTransition {
	ret := &RandomTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"random",
		"p",
	)
	return ret
}

func (m *RandomTransition) Clone() openxml.Element {
	ret := NewRandomTransition()
	return ret
}

func (m *RandomTransition) Validate() error {
	return nil
}

// Defines the WedgeTransition Class.
type WedgeTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main wedge"`
}

func NewWedgeTransition() *WedgeTransition {
	ret := &WedgeTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"wedge",
		"p",
	)
	return ret
}

func (m *WedgeTransition) Clone() openxml.Element {
	ret := NewWedgeTransition()
	return ret
}

func (m *WedgeTransition) Validate() error {
	return nil
}

// Slide Range.
type SlideRange struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldRg"`
}

func NewSlideRange() *SlideRange {
	ret := &SlideRange{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sldRg",
		"p",
	)
	return ret
}

func (m *SlideRange) Clone() openxml.Element {
	ret := NewSlideRange()
	return ret
}

func (m *SlideRange) Validate() error {
	return nil
}

// Character Range.
type CharRange struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main charRg"`
}

func NewCharRange() *CharRange {
	ret := &CharRange{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"charRg",
		"p",
	)
	return ret
}

func (m *CharRange) Clone() openxml.Element {
	ret := NewCharRange()
	return ret
}

func (m *CharRange) Validate() error {
	return nil
}

// Paragraph Text Range.
type ParagraphIndexRange struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main pRg"`
}

func NewParagraphIndexRange() *ParagraphIndexRange {
	ret := &ParagraphIndexRange{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pRg",
		"p",
	)
	return ret
}

func (m *ParagraphIndexRange) Clone() openxml.Element {
	ret := NewParagraphIndexRange()
	return ret
}

func (m *ParagraphIndexRange) Validate() error {
	return nil
}

// Defines the IndexRangeType Class.
type IndexRangeType struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:""`
	Start   *types.UInt32Value `xml:"st,attr,omitempty"`
	End     *types.UInt32Value `xml:"end,attr,omitempty"`
}

func NewIndexRangeType() *IndexRangeType {
	ret := &IndexRangeType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *IndexRangeType) Clone() openxml.Element {
	ret := NewIndexRangeType()
	if m.Start != nil {
		v := *m.Start
		ret.Start = &v
	}
	if m.End != nil {
		v := *m.End
		ret.End = &v
	}
	return ret
}

func (m *IndexRangeType) Validate() error {
	return nil
}

// Custom Show.
type CustomShowReference struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main custShow"`
	Id      *types.UInt32Value `xml:"id,attr,omitempty"`
}

func NewCustomShowReference() *CustomShowReference {
	ret := &CustomShowReference{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"custShow",
		"p",
	)
	return ret
}

func (m *CustomShowReference) Clone() openxml.Element {
	ret := NewCustomShowReference()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *CustomShowReference) Validate() error {
	return nil
}

// Extension.
type Extension struct {
	*openxml.CompositeElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri     *types.StringValue `xml:"uri,attr,omitempty"`
}

func NewExtension() *Extension {
	ret := &Extension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *Extension) Clone() openxml.Element {
	ret := NewExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	return ret
}

func (m *Extension) Validate() error {
	return nil
}

// Browse Slide Show Mode.
type BrowseSlideMode struct {
	*openxml.LeafElementBase
	XMLName       xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main browse"`
	ShowScrollbar *types.BooleanValue `xml:"showScrollbar,attr,omitempty"`
}

func NewBrowseSlideMode() *BrowseSlideMode {
	ret := &BrowseSlideMode{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"browse",
		"p",
	)
	return ret
}

func (m *BrowseSlideMode) Clone() openxml.Element {
	ret := NewBrowseSlideMode()
	if m.ShowScrollbar != nil {
		v := *m.ShowScrollbar
		ret.ShowScrollbar = &v
	}
	return ret
}

func (m *BrowseSlideMode) Validate() error {
	return nil
}

// Kiosk Slide Show Mode.
type KioskSlideMode struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main kiosk"`
	Restart *types.UInt32Value `xml:"restart,attr,omitempty"`
}

func NewKioskSlideMode() *KioskSlideMode {
	ret := &KioskSlideMode{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"kiosk",
		"p",
	)
	return ret
}

func (m *KioskSlideMode) Clone() openxml.Element {
	ret := NewKioskSlideMode()
	if m.Restart != nil {
		v := *m.Restart
		ret.Restart = &v
	}
	return ret
}

func (m *KioskSlideMode) Validate() error {
	return nil
}

// Color Scheme Map.
type ColorMap struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name                                 `xml:"http://schemas.openxmlformats.org/presentationml/2006/main clrMap"`
	Background1       *types.EnumValue[ColorSchemeIndexValues] `xml:"bg1,attr,omitempty"`
	Text1             *types.EnumValue[ColorSchemeIndexValues] `xml:"tx1,attr,omitempty"`
	Background2       *types.EnumValue[ColorSchemeIndexValues] `xml:"bg2,attr,omitempty"`
	Text2             *types.EnumValue[ColorSchemeIndexValues] `xml:"tx2,attr,omitempty"`
	Accent1           *types.EnumValue[ColorSchemeIndexValues] `xml:"accent1,attr,omitempty"`
	Accent2           *types.EnumValue[ColorSchemeIndexValues] `xml:"accent2,attr,omitempty"`
	Accent3           *types.EnumValue[ColorSchemeIndexValues] `xml:"accent3,attr,omitempty"`
	Accent4           *types.EnumValue[ColorSchemeIndexValues] `xml:"accent4,attr,omitempty"`
	Accent5           *types.EnumValue[ColorSchemeIndexValues] `xml:"accent5,attr,omitempty"`
	Accent6           *types.EnumValue[ColorSchemeIndexValues] `xml:"accent6,attr,omitempty"`
	Hyperlink         *types.EnumValue[ColorSchemeIndexValues] `xml:"hlink,attr,omitempty"`
	FollowedHyperlink *types.EnumValue[ColorSchemeIndexValues] `xml:"folHlink,attr,omitempty"`
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewColorMap() *ColorMap {
	ret := &ColorMap{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"clrMap",
		"p",
	)
	return ret
}

func (m *ColorMap) Clone() openxml.Element {
	ret := NewColorMap()
	if m.Background1 != nil {
		v := *m.Background1
		ret.Background1 = &v
	}
	if m.Text1 != nil {
		v := *m.Text1
		ret.Text1 = &v
	}
	if m.Background2 != nil {
		v := *m.Background2
		ret.Background2 = &v
	}
	if m.Text2 != nil {
		v := *m.Text2
		ret.Text2 = &v
	}
	if m.Accent1 != nil {
		v := *m.Accent1
		ret.Accent1 = &v
	}
	if m.Accent2 != nil {
		v := *m.Accent2
		ret.Accent2 = &v
	}
	if m.Accent3 != nil {
		v := *m.Accent3
		ret.Accent3 = &v
	}
	if m.Accent4 != nil {
		v := *m.Accent4
		ret.Accent4 = &v
	}
	if m.Accent5 != nil {
		v := *m.Accent5
		ret.Accent5 = &v
	}
	if m.Accent6 != nil {
		v := *m.Accent6
		ret.Accent6 = &v
	}
	if m.Hyperlink != nil {
		v := *m.Hyperlink
		ret.Hyperlink = &v
	}
	if m.FollowedHyperlink != nil {
		v := *m.FollowedHyperlink
		ret.FollowedHyperlink = &v
	}
	return ret
}

func (m *ColorMap) Validate() error {
	return nil
}

// Background Properties.
type BackgroundProperties struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bgPr"`
	ShadeToTitle *types.BooleanValue     `xml:"shadeToTitle,attr,omitempty"`
	NoFill       *drawingml.NoFill       `xml:"noFill,omitempty"`
	SolidFill    *drawingml.SolidFill    `xml:"solidFill,omitempty"`
	GradientFill *drawingml.GradientFill `xml:"gradFill,omitempty"`
	BlipFill     *drawingml.BlipFill     `xml:"blipFill,omitempty"`
	PatternFill  *drawingml.PatternFill  `xml:"pattFill,omitempty"`
	EffectList   *drawingml.EffectList   `xml:"effectLst,omitempty"`
	// Skipped DrawingML type not yet implemented: EffectDag
	ExtensionList *ExtensionList `xml:"extLst,omitempty"`
}

func NewBackgroundProperties() *BackgroundProperties {
	ret := &BackgroundProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bgPr",
		"p",
	)
	return ret
}

func (m *BackgroundProperties) Clone() openxml.Element {
	ret := NewBackgroundProperties()
	if m.ShadeToTitle != nil {
		v := *m.ShadeToTitle
		ret.ShadeToTitle = &v
	}
	if m.NoFill != nil {
		ret.NoFill = m.NoFill.Clone().(*drawingml.NoFill)
	}
	if m.SolidFill != nil {
		ret.SolidFill = m.SolidFill.Clone().(*drawingml.SolidFill)
	}
	if m.GradientFill != nil {
		ret.GradientFill = m.GradientFill.Clone().(*drawingml.GradientFill)
	}
	if m.BlipFill != nil {
		ret.BlipFill = m.BlipFill.Clone().(*drawingml.BlipFill)
	}
	if m.PatternFill != nil {
		ret.PatternFill = m.PatternFill.Clone().(*drawingml.PatternFill)
	}
	if m.EffectList != nil {
		ret.EffectList = m.EffectList.Clone().(*drawingml.EffectList)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *BackgroundProperties) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Background Style Reference.
type BackgroundStyleReference struct {
	*openxml.CompositeElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bgRef"`
	Index   *types.UInt32Value `xml:"idx,attr,omitempty"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewBackgroundStyleReference() *BackgroundStyleReference {
	ret := &BackgroundStyleReference{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bgRef",
		"p",
	)
	return ret
}

func (m *BackgroundStyleReference) Clone() openxml.Element {
	ret := NewBackgroundStyleReference()
	if m.Index != nil {
		v := *m.Index
		ret.Index = &v
	}
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *BackgroundStyleReference) Validate() error {
	return nil
}

// Data for the Windows platform..
type CommentPropertiesExtension struct {
	*openxml.CompositeElementBase
	XMLName     xml.Name     `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	TaskDetails *TaskDetails `xml:"taskDetails,omitempty"`
	Reactions   *Reactions   `xml:"reactions,omitempty"`
}

func NewCommentPropertiesExtension() *CommentPropertiesExtension {
	ret := &CommentPropertiesExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *CommentPropertiesExtension) Clone() openxml.Element {
	ret := NewCommentPropertiesExtension()
	if m.TaskDetails != nil {
		ret.TaskDetails = m.TaskDetails.Clone().(*TaskDetails)
	}
	if m.Reactions != nil {
		ret.Reactions = m.Reactions.Clone().(*Reactions)
	}
	return ret
}

func (m *CommentPropertiesExtension) Validate() error {
	if m.TaskDetails != nil {
		if err := m.TaskDetails.Validate(); err != nil {
			return err
		}
	}
	if m.Reactions != nil {
		if err := m.Reactions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// List of Comment Authors.
type CommentAuthorList struct {
	*openxml.PartRootElementBase
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cmAuthorLst"`
	CommentAuthor *CommentAuthor `xml:"cmAuthor,omitempty"`
}

func NewCommentAuthorList() *CommentAuthorList {
	ret := &CommentAuthorList{}
	ns := openxml.NamespacePresentationML
	ret.PartRootElementBase = openxml.NewPartRootElement(
		ns,
		"cmAuthorLst",
		"p",
	)
	return ret
}

func (m *CommentAuthorList) Clone() openxml.Element {
	ret := NewCommentAuthorList()
	if m.CommentAuthor != nil {
		ret.CommentAuthor = m.CommentAuthor.Clone().(*CommentAuthor)
	}
	return ret
}

func (m *CommentAuthorList) Validate() error {
	if m.CommentAuthor != nil {
		if err := m.CommentAuthor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Comment List (PresentationML standard namespace).
type PresentationCommentList struct {
	*openxml.PartRootElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cmLst"`
}

func NewPresentationCommentList() *PresentationCommentList {
	ret := &PresentationCommentList{}
	ns := openxml.NamespacePresentationML
	ret.PartRootElementBase = openxml.NewPartRootElement(
		ns,
		"cmLst",
		"p",
	)
	return ret
}

func (m *PresentationCommentList) Clone() openxml.Element {
	ret := NewPresentationCommentList()
	return ret
}

func (m *PresentationCommentList) Validate() error {
	return nil
}

// Global Element for OLE Objects and Controls.
type OleObject struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main oleObj"`
	ShapeId        *types.StringValue  `xml:"spid,attr,omitempty"`
	Name           *types.StringValue  `xml:"name,attr,omitempty"`
	ShowAsIcon     *types.BooleanValue `xml:"showAsIcon,attr,omitempty"`
	Id             *types.StringValue  `xml:"r:id,attr,omitempty"`
	ImageWidth     *types.Int32Value   `xml:"imgW,attr,omitempty"`
	ImageHeight    *types.Int32Value   `xml:"imgH,attr,omitempty"`
	ProgId         *types.StringValue  `xml:"progId,attr,omitempty"`
	OleObjectEmbed *OleObjectEmbed     `xml:"embed,omitempty"`
	OleObjectLink  *OleObjectLink      `xml:"link,omitempty"`
	Picture        *Picture            `xml:"pic,omitempty"`
}

func NewOleObject() *OleObject {
	ret := &OleObject{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"oleObj",
		"p",
	)
	return ret
}

func (m *OleObject) Clone() openxml.Element {
	ret := NewOleObject()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.ShowAsIcon != nil {
		v := *m.ShowAsIcon
		ret.ShowAsIcon = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.ImageWidth != nil {
		v := *m.ImageWidth
		ret.ImageWidth = &v
	}
	if m.ImageHeight != nil {
		v := *m.ImageHeight
		ret.ImageHeight = &v
	}
	if m.ProgId != nil {
		v := *m.ProgId
		ret.ProgId = &v
	}
	if m.OleObjectEmbed != nil {
		ret.OleObjectEmbed = m.OleObjectEmbed.Clone().(*OleObjectEmbed)
	}
	if m.OleObjectLink != nil {
		ret.OleObjectLink = m.OleObjectLink.Clone().(*OleObjectLink)
	}
	if m.Picture != nil {
		ret.Picture = m.Picture.Clone().(*Picture)
	}
	return ret
}

func (m *OleObject) Validate() error {
	if m.OleObjectEmbed != nil {
		if err := m.OleObjectEmbed.Validate(); err != nil {
			return err
		}
	}
	if m.OleObjectLink != nil {
		if err := m.OleObjectLink.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Presentation-wide Properties.
type PresentationProperties struct {
	*openxml.CompositeElementBase
	XMLName                             xml.Name                             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main presentationPr"`
	HtmlPublishProperties               *HtmlPublishProperties               `xml:"htmlPubPr,omitempty"`
	WebProperties                       *WebProperties                       `xml:"webPr,omitempty"`
	PrintingProperties                  *PrintingProperties                  `xml:"prnPr,omitempty"`
	ShowProperties                      *ShowProperties                      `xml:"showPr,omitempty"`
	ColorMostRecentlyUsed               *ColorMostRecentlyUsed               `xml:"clrMru,omitempty"`
	PresentationPropertiesExtensionList *PresentationPropertiesExtensionList `xml:"extLst,omitempty"`
}

func NewPresentationProperties() *PresentationProperties {
	ret := &PresentationProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"presentationPr",
		"p",
	)
	return ret
}

func (m *PresentationProperties) Clone() openxml.Element {
	ret := NewPresentationProperties()
	if m.HtmlPublishProperties != nil {
		ret.HtmlPublishProperties = m.HtmlPublishProperties.Clone().(*HtmlPublishProperties)
	}
	if m.WebProperties != nil {
		ret.WebProperties = m.WebProperties.Clone().(*WebProperties)
	}
	if m.PrintingProperties != nil {
		ret.PrintingProperties = m.PrintingProperties.Clone().(*PrintingProperties)
	}
	if m.ShowProperties != nil {
		ret.ShowProperties = m.ShowProperties.Clone().(*ShowProperties)
	}
	if m.ColorMostRecentlyUsed != nil {
		ret.ColorMostRecentlyUsed = m.ColorMostRecentlyUsed.Clone().(*ColorMostRecentlyUsed)
	}
	if m.PresentationPropertiesExtensionList != nil {
		ret.PresentationPropertiesExtensionList = m.PresentationPropertiesExtensionList.Clone().(*PresentationPropertiesExtensionList)
	}
	return ret
}

func (m *PresentationProperties) Validate() error {
	if m.HtmlPublishProperties != nil {
		if err := m.HtmlPublishProperties.Validate(); err != nil {
			return err
		}
	}
	if m.WebProperties != nil {
		if err := m.WebProperties.Validate(); err != nil {
			return err
		}
	}
	if m.PrintingProperties != nil {
		if err := m.PrintingProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ShowProperties != nil {
		if err := m.ShowProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ColorMostRecentlyUsed != nil {
		if err := m.ColorMostRecentlyUsed.Validate(); err != nil {
			return err
		}
	}
	if m.PresentationPropertiesExtensionList != nil {
		if err := m.PresentationPropertiesExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Handout Master.
type HandoutMaster struct {
	*openxml.PartRootElementBase
	XMLName                    xml.Name                    `xml:"http://schemas.openxmlformats.org/presentationml/2006/main handoutMaster"`
	CommonSlideData            *CommonSlideData            `xml:"cSld,omitempty"`
	ColorMap                   *ColorMap                   `xml:"clrMap,omitempty"`
	HeaderFooter               *HeaderFooter               `xml:"hf,omitempty"`
	HandoutMasterExtensionList *HandoutMasterExtensionList `xml:"extLst,omitempty"`
}

func NewHandoutMaster() *HandoutMaster {
	ret := &HandoutMaster{}
	ns := openxml.NamespacePresentationML
	ret.PartRootElementBase = openxml.NewPartRootElement(
		ns,
		"handoutMaster",
		"p",
	)
	return ret
}

func (m *HandoutMaster) Clone() openxml.Element {
	ret := NewHandoutMaster()
	if m.CommonSlideData != nil {
		ret.CommonSlideData = m.CommonSlideData.Clone().(*CommonSlideData)
	}
	if m.ColorMap != nil {
		ret.ColorMap = m.ColorMap.Clone().(*ColorMap)
	}
	if m.HeaderFooter != nil {
		ret.HeaderFooter = m.HeaderFooter.Clone().(*HeaderFooter)
	}
	if m.HandoutMasterExtensionList != nil {
		ret.HandoutMasterExtensionList = m.HandoutMasterExtensionList.Clone().(*HandoutMasterExtensionList)
	}
	return ret
}

func (m *HandoutMaster) Validate() error {
	if m.ColorMap != nil {
		if err := m.ColorMap.Validate(); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.Validate(); err != nil {
			return err
		}
	}
	if m.HandoutMasterExtensionList != nil {
		if err := m.HandoutMasterExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Notes Master.
type NotesMaster struct {
	*openxml.PartRootElementBase
	XMLName                  xml.Name                  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main notesMaster"`
	CommonSlideData          *CommonSlideData          `xml:"cSld,omitempty"`
	ColorMap                 *ColorMap                 `xml:"clrMap,omitempty"`
	HeaderFooter             *HeaderFooter             `xml:"hf,omitempty"`
	NotesStyle               *NotesStyle               `xml:"notesStyle,omitempty"`
	NotesMasterExtensionList *NotesMasterExtensionList `xml:"extLst,omitempty"`
}

func NewNotesMaster() *NotesMaster {
	ret := &NotesMaster{}
	ns := openxml.NamespacePresentationML
	ret.PartRootElementBase = openxml.NewPartRootElement(
		ns,
		"notesMaster",
		"p",
	)
	return ret
}

func (m *NotesMaster) Clone() openxml.Element {
	ret := NewNotesMaster()
	if m.CommonSlideData != nil {
		ret.CommonSlideData = m.CommonSlideData.Clone().(*CommonSlideData)
	}
	if m.ColorMap != nil {
		ret.ColorMap = m.ColorMap.Clone().(*ColorMap)
	}
	if m.HeaderFooter != nil {
		ret.HeaderFooter = m.HeaderFooter.Clone().(*HeaderFooter)
	}
	if m.NotesStyle != nil {
		ret.NotesStyle = m.NotesStyle.Clone().(*NotesStyle)
	}
	if m.NotesMasterExtensionList != nil {
		ret.NotesMasterExtensionList = m.NotesMasterExtensionList.Clone().(*NotesMasterExtensionList)
	}
	return ret
}

func (m *NotesMaster) Validate() error {
	if m.ColorMap != nil {
		if err := m.ColorMap.Validate(); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.Validate(); err != nil {
			return err
		}
	}
	if m.NotesStyle != nil {
		if err := m.NotesStyle.Validate(); err != nil {
			return err
		}
	}
	if m.NotesMasterExtensionList != nil {
		if err := m.NotesMasterExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Notes Slide.
type NotesSlide struct {
	*openxml.PartRootElementBase
	XMLName                         xml.Name                       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main notes"`
	ShowMasterShapes                *types.BooleanValue            `xml:"showMasterSp,attr,omitempty"`
	ShowMasterPlaceholderAnimations *types.BooleanValue            `xml:"showMasterPhAnim,attr,omitempty"`
	CommonSlideData                 *CommonSlideData               `xml:"cSld,omitempty"`
	ColorMapOverride                *ColorMapOverride              `xml:"clrMapOvr,omitempty"`
	ExtensionListWithModification   *ExtensionListWithModification `xml:"extLst,omitempty"`
}

func NewNotesSlide() *NotesSlide {
	ret := &NotesSlide{}
	ns := openxml.NamespacePresentationML
	ret.PartRootElementBase = openxml.NewPartRootElement(
		ns,
		"notes",
		"p",
	)
	return ret
}

func (m *NotesSlide) Clone() openxml.Element {
	ret := NewNotesSlide()
	if m.ShowMasterShapes != nil {
		v := *m.ShowMasterShapes
		ret.ShowMasterShapes = &v
	}
	if m.ShowMasterPlaceholderAnimations != nil {
		v := *m.ShowMasterPlaceholderAnimations
		ret.ShowMasterPlaceholderAnimations = &v
	}
	if m.CommonSlideData != nil {
		ret.CommonSlideData = m.CommonSlideData.Clone().(*CommonSlideData)
	}
	if m.ColorMapOverride != nil {
		ret.ColorMapOverride = m.ColorMapOverride.Clone().(*ColorMapOverride)
	}
	if m.ExtensionListWithModification != nil {
		ret.ExtensionListWithModification = m.ExtensionListWithModification.Clone().(*ExtensionListWithModification)
	}
	return ret
}

func (m *NotesSlide) Validate() error {
	if m.ExtensionListWithModification != nil {
		if err := m.ExtensionListWithModification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Slide Synchronization Properties.
type SlideSyncProperties struct {
	*openxml.CompositeElementBase
	XMLName                 xml.Name             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldSyncPr"`
	ServerSlideId           *types.StringValue   `xml:"serverSldId,attr,omitempty"`
	ServerSlideModifiedTime *types.DateTimeValue `xml:"serverSldModifiedTime,attr,omitempty"`
	ClientInsertedTime      *types.DateTimeValue `xml:"clientInsertedTime,attr,omitempty"`
	ExtensionList           *ExtensionList       `xml:"extLst,omitempty"`
}

func NewSlideSyncProperties() *SlideSyncProperties {
	ret := &SlideSyncProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldSyncPr",
		"p",
	)
	return ret
}

func (m *SlideSyncProperties) Clone() openxml.Element {
	ret := NewSlideSyncProperties()
	if m.ServerSlideId != nil {
		v := *m.ServerSlideId
		ret.ServerSlideId = &v
	}
	if m.ServerSlideModifiedTime != nil {
		v := *m.ServerSlideModifiedTime
		ret.ServerSlideModifiedTime = &v
	}
	if m.ClientInsertedTime != nil {
		v := *m.ClientInsertedTime
		ret.ClientInsertedTime = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *SlideSyncProperties) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Programmable Tab List.
type TagList struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tagLst"`
	Tag     *Tag     `xml:"tag,omitempty"`
}

func NewTagList() *TagList {
	ret := &TagList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tagLst",
		"p",
	)
	return ret
}

func (m *TagList) Clone() openxml.Element {
	ret := NewTagList()
	if m.Tag != nil {
		ret.Tag = m.Tag.Clone().(*Tag)
	}
	return ret
}

func (m *TagList) Validate() error {
	if m.Tag != nil {
		if err := m.Tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Presentation-wide View Properties.
type ViewProperties struct {
	*openxml.CompositeElementBase
	XMLName                 xml.Name                     `xml:"http://schemas.openxmlformats.org/presentationml/2006/main viewPr"`
	LastView                *types.EnumValue[ViewValues] `xml:"lastView,attr,omitempty"`
	ShowComments            *types.BooleanValue          `xml:"showComments,attr,omitempty"`
	NormalViewProperties    *NormalViewProperties        `xml:"normalViewPr,omitempty"`
	SlideViewProperties     *SlideViewProperties         `xml:"slideViewPr,omitempty"`
	OutlineViewProperties   *OutlineViewProperties       `xml:"outlineViewPr,omitempty"`
	NotesTextViewProperties *NotesTextViewProperties     `xml:"notesTextViewPr,omitempty"`
	SorterViewProperties    *SorterViewProperties        `xml:"sorterViewPr,omitempty"`
	NotesViewProperties     *NotesViewProperties         `xml:"notesViewPr,omitempty"`
	GridSpacing             *GridSpacing                 `xml:"gridSpacing,omitempty"`
	ExtensionList           *ExtensionList               `xml:"extLst,omitempty"`
}

func NewViewProperties() *ViewProperties {
	ret := &ViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"viewPr",
		"p",
	)
	return ret
}

func (m *ViewProperties) Clone() openxml.Element {
	ret := NewViewProperties()
	if m.LastView != nil {
		v := *m.LastView
		ret.LastView = &v
	}
	if m.ShowComments != nil {
		v := *m.ShowComments
		ret.ShowComments = &v
	}
	if m.NormalViewProperties != nil {
		ret.NormalViewProperties = m.NormalViewProperties.Clone().(*NormalViewProperties)
	}
	if m.SlideViewProperties != nil {
		ret.SlideViewProperties = m.SlideViewProperties.Clone().(*SlideViewProperties)
	}
	if m.OutlineViewProperties != nil {
		ret.OutlineViewProperties = m.OutlineViewProperties.Clone().(*OutlineViewProperties)
	}
	if m.NotesTextViewProperties != nil {
		ret.NotesTextViewProperties = m.NotesTextViewProperties.Clone().(*NotesTextViewProperties)
	}
	if m.SorterViewProperties != nil {
		ret.SorterViewProperties = m.SorterViewProperties.Clone().(*SorterViewProperties)
	}
	if m.NotesViewProperties != nil {
		ret.NotesViewProperties = m.NotesViewProperties.Clone().(*NotesViewProperties)
	}
	if m.GridSpacing != nil {
		ret.GridSpacing = m.GridSpacing.Clone().(*GridSpacing)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *ViewProperties) Validate() error {
	if m.NormalViewProperties != nil {
		if err := m.NormalViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.SlideViewProperties != nil {
		if err := m.SlideViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.OutlineViewProperties != nil {
		if err := m.OutlineViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.NotesTextViewProperties != nil {
		if err := m.NotesTextViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.SorterViewProperties != nil {
		if err := m.SorterViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.NotesViewProperties != nil {
		if err := m.NotesViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.GridSpacing != nil {
		if err := m.GridSpacing.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Sound.
type Sound struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main snd"`
}

func NewSound() *Sound {
	ret := &Sound{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"snd",
		"p",
	)
	return ret
}

func (m *Sound) Clone() openxml.Element {
	ret := NewSound()
	return ret
}

func (m *Sound) Validate() error {
	return nil
}

// Sound Target.
type SoundTarget struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sndTgt"`
}

func NewSoundTarget() *SoundTarget {
	ret := &SoundTarget{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sndTgt",
		"p",
	)
	return ret
}

func (m *SoundTarget) Clone() openxml.Element {
	ret := NewSoundTarget()
	return ret
}

func (m *SoundTarget) Validate() error {
	return nil
}

// Defines the EmbeddedWavAudioFileType Class.
type EmbeddedWavAudioFileType struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:""`
	Embed   *types.StringValue  `xml:"r:embed,attr,omitempty"`
	Name    *types.StringValue  `xml:"name,attr,omitempty"`
	BuiltIn *types.BooleanValue `xml:"builtIn,attr,omitempty"`
}

func NewEmbeddedWavAudioFileType() *EmbeddedWavAudioFileType {
	ret := &EmbeddedWavAudioFileType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *EmbeddedWavAudioFileType) Clone() openxml.Element {
	ret := NewEmbeddedWavAudioFileType()
	if m.Embed != nil {
		v := *m.Embed
		ret.Embed = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.BuiltIn != nil {
		v := *m.BuiltIn
		ret.BuiltIn = &v
	}
	return ret
}

func (m *EmbeddedWavAudioFileType) Validate() error {
	return nil
}

// Start Sound Action.
type StartSoundAction struct {
	*openxml.CompositeElementBase
	XMLName xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main stSnd"`
	Loop    *types.BooleanValue `xml:"loop,attr,omitempty"`
	Sound   *Sound              `xml:"snd,omitempty"`
}

func NewStartSoundAction() *StartSoundAction {
	ret := &StartSoundAction{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"stSnd",
		"p",
	)
	return ret
}

func (m *StartSoundAction) Clone() openxml.Element {
	ret := NewStartSoundAction()
	if m.Loop != nil {
		v := *m.Loop
		ret.Loop = &v
	}
	if m.Sound != nil {
		ret.Sound = m.Sound.Clone().(*Sound)
	}
	return ret
}

func (m *StartSoundAction) Validate() error {
	if m.Sound != nil {
		if err := m.Sound.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Time Absolute.
type TimeAbsolute struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tmAbs"`
	Val     *types.StringValue `xml:"val,attr,omitempty"`
}

func NewTimeAbsolute() *TimeAbsolute {
	ret := &TimeAbsolute{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"tmAbs",
		"p",
	)
	return ret
}

func (m *TimeAbsolute) Clone() openxml.Element {
	ret := NewTimeAbsolute()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *TimeAbsolute) Validate() error {
	return nil
}

// Time Percentage.
type TimePercentage struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tmPct"`
	Val     *types.Int32Value `xml:"val,attr,omitempty"`
}

func NewTimePercentage() *TimePercentage {
	ret := &TimePercentage{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"tmPct",
		"p",
	)
	return ret
}

func (m *TimePercentage) Clone() openxml.Element {
	ret := NewTimePercentage()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *TimePercentage) Validate() error {
	return nil
}

// Time Node.
type TimeNode struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tn"`
	Val     *types.UInt32Value `xml:"val,attr,omitempty"`
}

func NewTimeNode() *TimeNode {
	ret := &TimeNode{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"tn",
		"p",
	)
	return ret
}

func (m *TimeNode) Clone() openxml.Element {
	ret := NewTimeNode()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *TimeNode) Validate() error {
	return nil
}

// Runtime Node Trigger Choice.
type RuntimeNodeTrigger struct {
	*openxml.LeafElementBase
	XMLName xml.Name                                   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main rtn"`
	Val     *types.EnumValue[TriggerRuntimeNodeValues] `xml:"val,attr,omitempty"`
}

func NewRuntimeNodeTrigger() *RuntimeNodeTrigger {
	ret := &RuntimeNodeTrigger{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"rtn",
		"p",
	)
	return ret
}

func (m *RuntimeNodeTrigger) Clone() openxml.Element {
	ret := NewRuntimeNodeTrigger()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *RuntimeNodeTrigger) Validate() error {
	return nil
}

// Defines the EndSync Class.
type EndSync struct {
	*openxml.CompositeElementBase
	XMLName            xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main endSync"`
	TargetElement      *TargetElement      `xml:"tgtEl,omitempty"`
	TimeNode           *TimeNode           `xml:"tn,omitempty"`
	RuntimeNodeTrigger *RuntimeNodeTrigger `xml:"rtn,omitempty"`
}

func NewEndSync() *EndSync {
	ret := &EndSync{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"endSync",
		"p",
	)
	return ret
}

func (m *EndSync) Clone() openxml.Element {
	ret := NewEndSync()
	if m.TargetElement != nil {
		ret.TargetElement = m.TargetElement.Clone().(*TargetElement)
	}
	if m.TimeNode != nil {
		ret.TimeNode = m.TimeNode.Clone().(*TimeNode)
	}
	if m.RuntimeNodeTrigger != nil {
		ret.RuntimeNodeTrigger = m.RuntimeNodeTrigger.Clone().(*RuntimeNodeTrigger)
	}
	return ret
}

func (m *EndSync) Validate() error {
	if m.TimeNode != nil {
		if err := m.TimeNode.Validate(); err != nil {
			return err
		}
	}
	if m.RuntimeNodeTrigger != nil {
		if err := m.RuntimeNodeTrigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TimeListConditionalType Class.
type TimeListConditionalType struct {
	*openxml.CompositeElementBase
	XMLName            xml.Name                             `xml:""`
	Event              *types.EnumValue[TriggerEventValues] `xml:"evt,attr,omitempty"`
	Delay              *types.StringValue                   `xml:"delay,attr,omitempty"`
	TargetElement      *TargetElement                       `xml:"tgtEl,omitempty"`
	TimeNode           *TimeNode                            `xml:"tn,omitempty"`
	RuntimeNodeTrigger *RuntimeNodeTrigger                  `xml:"rtn,omitempty"`
}

func NewTimeListConditionalType() *TimeListConditionalType {
	ret := &TimeListConditionalType{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TimeListConditionalType) Clone() openxml.Element {
	ret := NewTimeListConditionalType()
	if m.Event != nil {
		v := *m.Event
		ret.Event = &v
	}
	if m.Delay != nil {
		v := *m.Delay
		ret.Delay = &v
	}
	if m.TargetElement != nil {
		ret.TargetElement = m.TargetElement.Clone().(*TargetElement)
	}
	if m.TimeNode != nil {
		ret.TimeNode = m.TimeNode.Clone().(*TimeNode)
	}
	if m.RuntimeNodeTrigger != nil {
		ret.RuntimeNodeTrigger = m.RuntimeNodeTrigger.Clone().(*RuntimeNodeTrigger)
	}
	return ret
}

func (m *TimeListConditionalType) Validate() error {
	if m.TimeNode != nil {
		if err := m.TimeNode.Validate(); err != nil {
			return err
		}
	}
	if m.RuntimeNodeTrigger != nil {
		if err := m.RuntimeNodeTrigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Exclusive.
type ExclusiveTimeNode struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main excl"`
	CommonTimeNode *CommonTimeNode `xml:"cTn,omitempty"`
}

func NewExclusiveTimeNode() *ExclusiveTimeNode {
	ret := &ExclusiveTimeNode{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"excl",
		"p",
	)
	return ret
}

func (m *ExclusiveTimeNode) Clone() openxml.Element {
	ret := NewExclusiveTimeNode()
	if m.CommonTimeNode != nil {
		ret.CommonTimeNode = m.CommonTimeNode.Clone().(*CommonTimeNode)
	}
	return ret
}

func (m *ExclusiveTimeNode) Validate() error {
	return nil
}

// Command.
type Command struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name                        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cmd"`
	Type           *types.EnumValue[CommandValues] `xml:"type,attr,omitempty"`
	CommandName    *types.StringValue              `xml:"cmd,attr,omitempty"`
	CommonBehavior *CommonBehavior                 `xml:"cBhvr,omitempty"`
}

func NewCommand() *Command {
	ret := &Command{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cmd",
		"p",
	)
	return ret
}

func (m *Command) Clone() openxml.Element {
	ret := NewCommand()
	if m.Type != nil {
		v := *m.Type
		ret.Type = &v
	}
	if m.CommandName != nil {
		v := *m.CommandName
		ret.CommandName = &v
	}
	if m.CommonBehavior != nil {
		ret.CommonBehavior = m.CommonBehavior.Clone().(*CommonBehavior)
	}
	return ret
}

func (m *Command) Validate() error {
	if m.CommonBehavior != nil {
		if err := m.CommonBehavior.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set Time Node Behavior.
type SetBehavior struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main set"`
	CommonBehavior *CommonBehavior `xml:"cBhvr,omitempty"`
	ToVariantValue *ToVariantValue `xml:"to,omitempty"`
}

func NewSetBehavior() *SetBehavior {
	ret := &SetBehavior{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"set",
		"p",
	)
	return ret
}

func (m *SetBehavior) Clone() openxml.Element {
	ret := NewSetBehavior()
	if m.CommonBehavior != nil {
		ret.CommonBehavior = m.CommonBehavior.Clone().(*CommonBehavior)
	}
	if m.ToVariantValue != nil {
		ret.ToVariantValue = m.ToVariantValue.Clone().(*ToVariantValue)
	}
	return ret
}

func (m *SetBehavior) Validate() error {
	if m.CommonBehavior != nil {
		if err := m.CommonBehavior.Validate(); err != nil {
			return err
		}
	}
	if m.ToVariantValue != nil {
		if err := m.ToVariantValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Audio.
type Audio struct {
	*openxml.CompositeElementBase
	XMLName         xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main audio"`
	IsNarration     *types.BooleanValue `xml:"isNarration,attr,omitempty"`
	CommonMediaNode *CommonMediaNode    `xml:"cMediaNode,omitempty"`
}

func NewAudio() *Audio {
	ret := &Audio{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"audio",
		"p",
	)
	return ret
}

func (m *Audio) Clone() openxml.Element {
	ret := NewAudio()
	if m.IsNarration != nil {
		v := *m.IsNarration
		ret.IsNarration = &v
	}
	if m.CommonMediaNode != nil {
		ret.CommonMediaNode = m.CommonMediaNode.Clone().(*CommonMediaNode)
	}
	return ret
}

func (m *Audio) Validate() error {
	if m.CommonMediaNode != nil {
		if err := m.CommonMediaNode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Video.
type Video struct {
	*openxml.CompositeElementBase
	XMLName         xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main video"`
	FullScreen      *types.BooleanValue `xml:"fullScrn,attr,omitempty"`
	CommonMediaNode *CommonMediaNode    `xml:"cMediaNode,omitempty"`
}

func NewVideo() *Video {
	ret := &Video{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"video",
		"p",
	)
	return ret
}

func (m *Video) Clone() openxml.Element {
	ret := NewVideo()
	if m.FullScreen != nil {
		v := *m.FullScreen
		ret.FullScreen = &v
	}
	if m.CommonMediaNode != nil {
		ret.CommonMediaNode = m.CommonMediaNode.Clone().(*CommonMediaNode)
	}
	return ret
}

func (m *Video) Validate() error {
	if m.CommonMediaNode != nil {
		if err := m.CommonMediaNode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Previous Conditions List.
type PreviousConditionList struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main prevCondLst"`
	Condition *Condition `xml:"cond,omitempty"`
}

func NewPreviousConditionList() *PreviousConditionList {
	ret := &PreviousConditionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"prevCondLst",
		"p",
	)
	return ret
}

func (m *PreviousConditionList) Clone() openxml.Element {
	ret := NewPreviousConditionList()
	if m.Condition != nil {
		ret.Condition = m.Condition.Clone().(*Condition)
	}
	return ret
}

func (m *PreviousConditionList) Validate() error {
	return nil
}

// Next Conditions List.
type NextConditionList struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main nextCondLst"`
	Condition *Condition `xml:"cond,omitempty"`
}

func NewNextConditionList() *NextConditionList {
	ret := &NextConditionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"nextCondLst",
		"p",
	)
	return ret
}

func (m *NextConditionList) Clone() openxml.Element {
	ret := NewNextConditionList()
	if m.Condition != nil {
		ret.Condition = m.Condition.Clone().(*Condition)
	}
	return ret
}

func (m *NextConditionList) Validate() error {
	return nil
}

// Defines the TimeListTimeConditionalListType Class.
type TimeListTimeConditionalListType struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name   `xml:""`
	Condition *Condition `xml:"cond,omitempty"`
}

func NewTimeListTimeConditionalListType() *TimeListTimeConditionalListType {
	ret := &TimeListTimeConditionalListType{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TimeListTimeConditionalListType) Clone() openxml.Element {
	ret := NewTimeListTimeConditionalListType()
	if m.Condition != nil {
		ret.Condition = m.Condition.Clone().(*Condition)
	}
	return ret
}

func (m *TimeListTimeConditionalListType) Validate() error {
	return nil
}

// Attribute Name.
type AttributeName struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main attrName"`
}

func NewAttributeName() *AttributeName {
	ret := &AttributeName{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"attrName",
		"p",
	)
	return ret
}

func (m *AttributeName) Clone() openxml.Element {
	ret := NewAttributeName()
	return ret
}

func (m *AttributeName) Validate() error {
	return nil
}

// Defines the Text Class.
type Text struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main text"`
}

func NewText() *Text {
	ret := &Text{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"text",
		"p",
	)
	return ret
}

func (m *Text) Clone() openxml.Element {
	ret := NewText()
	return ret
}

func (m *Text) Validate() error {
	return nil
}

// Attribute Name List.
type AttributeNameList struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main attrNameLst"`
	AttributeName *AttributeName `xml:"attrName,omitempty"`
}

func NewAttributeNameList() *AttributeNameList {
	ret := &AttributeNameList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"attrNameLst",
		"p",
	)
	return ret
}

func (m *AttributeNameList) Clone() openxml.Element {
	ret := NewAttributeNameList()
	if m.AttributeName != nil {
		ret.AttributeName = m.AttributeName.Clone().(*AttributeName)
	}
	return ret
}

func (m *AttributeNameList) Validate() error {
	if m.AttributeName != nil {
		if err := m.AttributeName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Boolean Variant.
type BooleanVariantValue struct {
	*openxml.LeafElementBase
	XMLName xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main boolVal"`
	Val     *types.BooleanValue `xml:"val,attr,omitempty"`
}

func NewBooleanVariantValue() *BooleanVariantValue {
	ret := &BooleanVariantValue{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"boolVal",
		"p",
	)
	return ret
}

func (m *BooleanVariantValue) Clone() openxml.Element {
	ret := NewBooleanVariantValue()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *BooleanVariantValue) Validate() error {
	return nil
}

// Integer.
type IntegerVariantValue struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main intVal"`
	Val     *types.Int32Value `xml:"val,attr,omitempty"`
}

func NewIntegerVariantValue() *IntegerVariantValue {
	ret := &IntegerVariantValue{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"intVal",
		"p",
	)
	return ret
}

func (m *IntegerVariantValue) Clone() openxml.Element {
	ret := NewIntegerVariantValue()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *IntegerVariantValue) Validate() error {
	return nil
}

// Float Value.
type FloatVariantValue struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main fltVal"`
	Val     *types.StringValue `xml:"val,attr,omitempty"`
}

func NewFloatVariantValue() *FloatVariantValue {
	ret := &FloatVariantValue{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"fltVal",
		"p",
	)
	return ret
}

func (m *FloatVariantValue) Clone() openxml.Element {
	ret := NewFloatVariantValue()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *FloatVariantValue) Validate() error {
	return nil
}

// String Value.
type StringVariantValue struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main strVal"`
	Val     *types.StringValue `xml:"val,attr,omitempty"`
}

func NewStringVariantValue() *StringVariantValue {
	ret := &StringVariantValue{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"strVal",
		"p",
	)
	return ret
}

func (m *StringVariantValue) Clone() openxml.Element {
	ret := NewStringVariantValue()
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *StringVariantValue) Validate() error {
	return nil
}

// Color Value.
type ColorValue struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main clrVal"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewColorValue() *ColorValue {
	ret := &ColorValue{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"clrVal",
		"p",
	)
	return ret
}

func (m *ColorValue) Clone() openxml.Element {
	ret := NewColorValue()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *ColorValue) Validate() error {
	return nil
}

// Pen Color for Slide Show.
type PenColor struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main penClr"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewPenColor() *PenColor {
	ret := &PenColor{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"penClr",
		"p",
	)
	return ret
}

func (m *PenColor) Clone() openxml.Element {
	ret := NewPenColor()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *PenColor) Validate() error {
	return nil
}

// Time Animate Value.
type TimeAnimateValue struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tav"`
	Time         *types.StringValue `xml:"tm,attr,omitempty"`
	Fomula       *types.StringValue `xml:"fmla,attr,omitempty"`
	VariantValue *VariantValue      `xml:"val,omitempty"`
}

func NewTimeAnimateValue() *TimeAnimateValue {
	ret := &TimeAnimateValue{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tav",
		"p",
	)
	return ret
}

func (m *TimeAnimateValue) Clone() openxml.Element {
	ret := NewTimeAnimateValue()
	if m.Time != nil {
		v := *m.Time
		ret.Time = &v
	}
	if m.Fomula != nil {
		v := *m.Fomula
		ret.Fomula = &v
	}
	if m.VariantValue != nil {
		ret.VariantValue = m.VariantValue.Clone().(*VariantValue)
	}
	return ret
}

func (m *TimeAnimateValue) Validate() error {
	if m.VariantValue != nil {
		if err := m.VariantValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// RGB.
type RgbColor struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main rgb"`
	Red     *types.Int32Value `xml:"r,attr,omitempty"`
	Green   *types.Int32Value `xml:"g,attr,omitempty"`
	Blue    *types.Int32Value `xml:"b,attr,omitempty"`
}

func NewRgbColor() *RgbColor {
	ret := &RgbColor{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"rgb",
		"p",
	)
	return ret
}

func (m *RgbColor) Clone() openxml.Element {
	ret := NewRgbColor()
	if m.Red != nil {
		v := *m.Red
		ret.Red = &v
	}
	if m.Green != nil {
		v := *m.Green
		ret.Green = &v
	}
	if m.Blue != nil {
		v := *m.Blue
		ret.Blue = &v
	}
	return ret
}

func (m *RgbColor) Validate() error {
	return nil
}

// HSL.
type HslColor struct {
	*openxml.LeafElementBase
	XMLName    xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main hsl"`
	Hue        *types.Int32Value `xml:"h,attr,omitempty"`
	Saturation *types.Int32Value `xml:"s,attr,omitempty"`
	Lightness  *types.Int32Value `xml:"l,attr,omitempty"`
}

func NewHslColor() *HslColor {
	ret := &HslColor{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"hsl",
		"p",
	)
	return ret
}

func (m *HslColor) Clone() openxml.Element {
	ret := NewHslColor()
	if m.Hue != nil {
		v := *m.Hue
		ret.Hue = &v
	}
	if m.Saturation != nil {
		v := *m.Saturation
		ret.Saturation = &v
	}
	if m.Lightness != nil {
		v := *m.Lightness
		ret.Lightness = &v
	}
	return ret
}

func (m *HslColor) Validate() error {
	return nil
}

// Defines the CommonBehavior Class.
type CommonBehavior struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name                                   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cBhvr"`
	Additive          *types.EnumValue[BehaviorAdditiveValues]   `xml:"additive,attr,omitempty"`
	Accumulate        *types.EnumValue[BehaviorAccumulateValues] `xml:"accumulate,attr,omitempty"`
	TransformType     *types.EnumValue[BehaviorTransformValues]  `xml:"xfrmType,attr,omitempty"`
	From              *types.StringValue                         `xml:"from,attr,omitempty"`
	To                *types.StringValue                         `xml:"to,attr,omitempty"`
	By                *types.StringValue                         `xml:"by,attr,omitempty"`
	RuntimeContext    *types.StringValue                         `xml:"rctx,attr,omitempty"`
	Override          *types.EnumValue[BehaviorOverrideValues]   `xml:"override,attr,omitempty"`
	CommonTimeNode    *CommonTimeNode                            `xml:"cTn,omitempty"`
	TargetElement     *TargetElement                             `xml:"tgtEl,omitempty"`
	AttributeNameList *AttributeNameList                         `xml:"attrNameLst,omitempty"`
}

func NewCommonBehavior() *CommonBehavior {
	ret := &CommonBehavior{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cBhvr",
		"p",
	)
	return ret
}

func (m *CommonBehavior) Clone() openxml.Element {
	ret := NewCommonBehavior()
	if m.Additive != nil {
		v := *m.Additive
		ret.Additive = &v
	}
	if m.Accumulate != nil {
		v := *m.Accumulate
		ret.Accumulate = &v
	}
	if m.TransformType != nil {
		v := *m.TransformType
		ret.TransformType = &v
	}
	if m.From != nil {
		v := *m.From
		ret.From = &v
	}
	if m.To != nil {
		v := *m.To
		ret.To = &v
	}
	if m.By != nil {
		v := *m.By
		ret.By = &v
	}
	if m.RuntimeContext != nil {
		v := *m.RuntimeContext
		ret.RuntimeContext = &v
	}
	if m.Override != nil {
		v := *m.Override
		ret.Override = &v
	}
	if m.CommonTimeNode != nil {
		ret.CommonTimeNode = m.CommonTimeNode.Clone().(*CommonTimeNode)
	}
	if m.TargetElement != nil {
		ret.TargetElement = m.TargetElement.Clone().(*TargetElement)
	}
	if m.AttributeNameList != nil {
		ret.AttributeNameList = m.AttributeNameList.Clone().(*AttributeNameList)
	}
	return ret
}

func (m *CommonBehavior) Validate() error {
	if m.AttributeNameList != nil {
		if err := m.AttributeNameList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Progress.
type Progress struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main progress"`
	FloatVariantValue *FloatVariantValue `xml:"fltVal,omitempty"`
}

func NewProgress() *Progress {
	ret := &Progress{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"progress",
		"p",
	)
	return ret
}

func (m *Progress) Clone() openxml.Element {
	ret := NewProgress()
	if m.FloatVariantValue != nil {
		ret.FloatVariantValue = m.FloatVariantValue.Clone().(*FloatVariantValue)
	}
	return ret
}

func (m *Progress) Validate() error {
	if m.FloatVariantValue != nil {
		if err := m.FloatVariantValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// To.
type ToVariantValue struct {
	*openxml.CompositeElementBase
	XMLName             xml.Name             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main to"`
	BooleanVariantValue *BooleanVariantValue `xml:"boolVal,omitempty"`
	IntegerVariantValue *IntegerVariantValue `xml:"intVal,omitempty"`
	FloatVariantValue   *FloatVariantValue   `xml:"fltVal,omitempty"`
	StringVariantValue  *StringVariantValue  `xml:"strVal,omitempty"`
	ColorValue          *ColorValue          `xml:"clrVal,omitempty"`
}

func NewToVariantValue() *ToVariantValue {
	ret := &ToVariantValue{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"to",
		"p",
	)
	return ret
}

func (m *ToVariantValue) Clone() openxml.Element {
	ret := NewToVariantValue()
	if m.BooleanVariantValue != nil {
		ret.BooleanVariantValue = m.BooleanVariantValue.Clone().(*BooleanVariantValue)
	}
	if m.IntegerVariantValue != nil {
		ret.IntegerVariantValue = m.IntegerVariantValue.Clone().(*IntegerVariantValue)
	}
	if m.FloatVariantValue != nil {
		ret.FloatVariantValue = m.FloatVariantValue.Clone().(*FloatVariantValue)
	}
	if m.StringVariantValue != nil {
		ret.StringVariantValue = m.StringVariantValue.Clone().(*StringVariantValue)
	}
	if m.ColorValue != nil {
		ret.ColorValue = m.ColorValue.Clone().(*ColorValue)
	}
	return ret
}

func (m *ToVariantValue) Validate() error {
	if m.BooleanVariantValue != nil {
		if err := m.BooleanVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.IntegerVariantValue != nil {
		if err := m.IntegerVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.FloatVariantValue != nil {
		if err := m.FloatVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.StringVariantValue != nil {
		if err := m.StringVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.ColorValue != nil {
		if err := m.ColorValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Value.
type VariantValue struct {
	*openxml.CompositeElementBase
	XMLName             xml.Name             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main val"`
	BooleanVariantValue *BooleanVariantValue `xml:"boolVal,omitempty"`
	IntegerVariantValue *IntegerVariantValue `xml:"intVal,omitempty"`
	FloatVariantValue   *FloatVariantValue   `xml:"fltVal,omitempty"`
	StringVariantValue  *StringVariantValue  `xml:"strVal,omitempty"`
	ColorValue          *ColorValue          `xml:"clrVal,omitempty"`
}

func NewVariantValue() *VariantValue {
	ret := &VariantValue{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"val",
		"p",
	)
	return ret
}

func (m *VariantValue) Clone() openxml.Element {
	ret := NewVariantValue()
	if m.BooleanVariantValue != nil {
		ret.BooleanVariantValue = m.BooleanVariantValue.Clone().(*BooleanVariantValue)
	}
	if m.IntegerVariantValue != nil {
		ret.IntegerVariantValue = m.IntegerVariantValue.Clone().(*IntegerVariantValue)
	}
	if m.FloatVariantValue != nil {
		ret.FloatVariantValue = m.FloatVariantValue.Clone().(*FloatVariantValue)
	}
	if m.StringVariantValue != nil {
		ret.StringVariantValue = m.StringVariantValue.Clone().(*StringVariantValue)
	}
	if m.ColorValue != nil {
		ret.ColorValue = m.ColorValue.Clone().(*ColorValue)
	}
	return ret
}

func (m *VariantValue) Validate() error {
	if m.BooleanVariantValue != nil {
		if err := m.BooleanVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.IntegerVariantValue != nil {
		if err := m.IntegerVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.FloatVariantValue != nil {
		if err := m.FloatVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.StringVariantValue != nil {
		if err := m.StringVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.ColorValue != nil {
		if err := m.ColorValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TimeListAnimationVariantType Class.
type TimeListAnimationVariantType struct {
	*openxml.CompositeElementBase
	XMLName             xml.Name             `xml:""`
	BooleanVariantValue *BooleanVariantValue `xml:"boolVal,omitempty"`
	IntegerVariantValue *IntegerVariantValue `xml:"intVal,omitempty"`
	FloatVariantValue   *FloatVariantValue   `xml:"fltVal,omitempty"`
	StringVariantValue  *StringVariantValue  `xml:"strVal,omitempty"`
	ColorValue          *ColorValue          `xml:"clrVal,omitempty"`
}

func NewTimeListAnimationVariantType() *TimeListAnimationVariantType {
	ret := &TimeListAnimationVariantType{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TimeListAnimationVariantType) Clone() openxml.Element {
	ret := NewTimeListAnimationVariantType()
	if m.BooleanVariantValue != nil {
		ret.BooleanVariantValue = m.BooleanVariantValue.Clone().(*BooleanVariantValue)
	}
	if m.IntegerVariantValue != nil {
		ret.IntegerVariantValue = m.IntegerVariantValue.Clone().(*IntegerVariantValue)
	}
	if m.FloatVariantValue != nil {
		ret.FloatVariantValue = m.FloatVariantValue.Clone().(*FloatVariantValue)
	}
	if m.StringVariantValue != nil {
		ret.StringVariantValue = m.StringVariantValue.Clone().(*StringVariantValue)
	}
	if m.ColorValue != nil {
		ret.ColorValue = m.ColorValue.Clone().(*ColorValue)
	}
	return ret
}

func (m *TimeListAnimationVariantType) Validate() error {
	if m.BooleanVariantValue != nil {
		if err := m.BooleanVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.IntegerVariantValue != nil {
		if err := m.IntegerVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.FloatVariantValue != nil {
		if err := m.FloatVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.StringVariantValue != nil {
		if err := m.StringVariantValue.Validate(); err != nil {
			return err
		}
	}
	if m.ColorValue != nil {
		if err := m.ColorValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Common Media Node Properties.
type CommonMediaNode struct {
	*openxml.CompositeElementBase
	XMLName         xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cMediaNode"`
	Volume          *types.Int32Value   `xml:"vol,attr,omitempty"`
	Mute            *types.BooleanValue `xml:"mute,attr,omitempty"`
	SlideCount      *types.UInt32Value  `xml:"numSld,attr,omitempty"`
	ShowWhenStopped *types.BooleanValue `xml:"showWhenStopped,attr,omitempty"`
	CommonTimeNode  *CommonTimeNode     `xml:"cTn,omitempty"`
	TargetElement   *TargetElement      `xml:"tgtEl,omitempty"`
}

func NewCommonMediaNode() *CommonMediaNode {
	ret := &CommonMediaNode{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cMediaNode",
		"p",
	)
	return ret
}

func (m *CommonMediaNode) Clone() openxml.Element {
	ret := NewCommonMediaNode()
	if m.Volume != nil {
		v := *m.Volume
		ret.Volume = &v
	}
	if m.Mute != nil {
		v := *m.Mute
		ret.Mute = &v
	}
	if m.SlideCount != nil {
		v := *m.SlideCount
		ret.SlideCount = &v
	}
	if m.ShowWhenStopped != nil {
		v := *m.ShowWhenStopped
		ret.ShowWhenStopped = &v
	}
	if m.CommonTimeNode != nil {
		ret.CommonTimeNode = m.CommonTimeNode.Clone().(*CommonTimeNode)
	}
	if m.TargetElement != nil {
		ret.TargetElement = m.TargetElement.Clone().(*TargetElement)
	}
	return ret
}

func (m *CommonMediaNode) Validate() error {
	return nil
}

// Template Effects.
type Template struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tmpl"`
	Level        *types.UInt32Value `xml:"lvl,attr,omitempty"`
	TimeNodeList *TimeNodeList      `xml:"tnLst,omitempty"`
}

func NewTemplate() *Template {
	ret := &Template{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tmpl",
		"p",
	)
	return ret
}

func (m *Template) Clone() openxml.Element {
	ret := NewTemplate()
	if m.Level != nil {
		v := *m.Level
		ret.Level = &v
	}
	if m.TimeNodeList != nil {
		ret.TimeNodeList = m.TimeNodeList.Clone().(*TimeNodeList)
	}
	return ret
}

func (m *Template) Validate() error {
	return nil
}

// Template effects.
type TemplateList struct {
	*openxml.CompositeElementBase
	XMLName  xml.Name  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tmplLst"`
	Template *Template `xml:"tmpl,omitempty"`
}

func NewTemplateList() *TemplateList {
	ret := &TemplateList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tmplLst",
		"p",
	)
	return ret
}

func (m *TemplateList) Clone() openxml.Element {
	ret := NewTemplateList()
	if m.Template != nil {
		ret.Template = m.Template.Clone().(*Template)
	}
	return ret
}

func (m *TemplateList) Validate() error {
	if m.Template != nil {
		if err := m.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Build Sub Elements.
type BuildSubElement struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bldSub"`
	// Skipped DrawingML type not yet implemented: BuildDiagram
	// Skipped DrawingML type not yet implemented: BuildChart
}

func NewBuildSubElement() *BuildSubElement {
	ret := &BuildSubElement{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bldSub",
		"p",
	)
	return ret
}

func (m *BuildSubElement) Clone() openxml.Element {
	ret := NewBuildSubElement()
	return ret
}

func (m *BuildSubElement) Validate() error {
	return nil
}

// Build Diagram.
type BuildDiagram struct {
	*openxml.LeafElementBase
	XMLName  xml.Name                             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bldDgm"`
	ShapeId  *types.StringValue                   `xml:"spid,attr,omitempty"`
	GroupId  *types.UInt32Value                   `xml:"grpId,attr,omitempty"`
	UiExpand *types.BooleanValue                  `xml:"uiExpand,attr,omitempty"`
	Build    *types.EnumValue[DiagramBuildValues] `xml:"bld,attr,omitempty"`
}

func NewBuildDiagram() *BuildDiagram {
	ret := &BuildDiagram{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bldDgm",
		"p",
	)
	return ret
}

func (m *BuildDiagram) Clone() openxml.Element {
	ret := NewBuildDiagram()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	if m.GroupId != nil {
		v := *m.GroupId
		ret.GroupId = &v
	}
	if m.UiExpand != nil {
		v := *m.UiExpand
		ret.UiExpand = &v
	}
	if m.Build != nil {
		v := *m.Build
		ret.Build = &v
	}
	return ret
}

func (m *BuildDiagram) Validate() error {
	return nil
}

// Build OLE Chart.
type BuildOleChart struct {
	*openxml.LeafElementBase
	XMLName           xml.Name                              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bldOleChart"`
	ShapeId           *types.StringValue                    `xml:"spid,attr,omitempty"`
	GroupId           *types.UInt32Value                    `xml:"grpId,attr,omitempty"`
	UiExpand          *types.BooleanValue                   `xml:"uiExpand,attr,omitempty"`
	Build             *types.EnumValue[OleChartBuildValues] `xml:"bld,attr,omitempty"`
	AnimateBackground *types.BooleanValue                   `xml:"animBg,attr,omitempty"`
}

func NewBuildOleChart() *BuildOleChart {
	ret := &BuildOleChart{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bldOleChart",
		"p",
	)
	return ret
}

func (m *BuildOleChart) Clone() openxml.Element {
	ret := NewBuildOleChart()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	if m.GroupId != nil {
		v := *m.GroupId
		ret.GroupId = &v
	}
	if m.UiExpand != nil {
		v := *m.UiExpand
		ret.UiExpand = &v
	}
	if m.Build != nil {
		v := *m.Build
		ret.Build = &v
	}
	if m.AnimateBackground != nil {
		v := *m.AnimateBackground
		ret.AnimateBackground = &v
	}
	return ret
}

func (m *BuildOleChart) Validate() error {
	return nil
}

// Build Graphics.
type BuildGraphics struct {
	*openxml.CompositeElementBase
	XMLName         xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bldGraphic"`
	ShapeId         *types.StringValue  `xml:"spid,attr,omitempty"`
	GroupId         *types.UInt32Value  `xml:"grpId,attr,omitempty"`
	UiExpand        *types.BooleanValue `xml:"uiExpand,attr,omitempty"`
	BuildAsOne      *BuildAsOne         `xml:"bldAsOne,omitempty"`
	BuildSubElement *BuildSubElement    `xml:"bldSub,omitempty"`
}

func NewBuildGraphics() *BuildGraphics {
	ret := &BuildGraphics{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bldGraphic",
		"p",
	)
	return ret
}

func (m *BuildGraphics) Clone() openxml.Element {
	ret := NewBuildGraphics()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	if m.GroupId != nil {
		v := *m.GroupId
		ret.GroupId = &v
	}
	if m.UiExpand != nil {
		v := *m.UiExpand
		ret.UiExpand = &v
	}
	if m.BuildAsOne != nil {
		ret.BuildAsOne = m.BuildAsOne.Clone().(*BuildAsOne)
	}
	if m.BuildSubElement != nil {
		ret.BuildSubElement = m.BuildSubElement.Clone().(*BuildSubElement)
	}
	return ret
}

func (m *BuildGraphics) Validate() error {
	if m.BuildAsOne != nil {
		if err := m.BuildAsOne.Validate(); err != nil {
			return err
		}
	}
	if m.BuildSubElement != nil {
		if err := m.BuildSubElement.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ExtensionListWithModification Class.
type ExtensionListWithModification struct {
	*openxml.CompositeElementBase
	XMLName   xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	Modify    *types.BooleanValue `xml:"mod,attr,omitempty"`
	Extension *Extension          `xml:"ext,omitempty"`
}

func NewExtensionListWithModification() *ExtensionListWithModification {
	ret := &ExtensionListWithModification{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *ExtensionListWithModification) Clone() openxml.Element {
	ret := NewExtensionListWithModification()
	if m.Modify != nil {
		v := *m.Modify
		ret.Modify = &v
	}
	if m.Extension != nil {
		ret.Extension = m.Extension.Clone().(*Extension)
	}
	return ret
}

func (m *ExtensionListWithModification) Validate() error {
	if m.Extension != nil {
		if err := m.Extension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// By.
type ByColor struct {
	*openxml.CompositeElementBase
	XMLName  xml.Name  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main by"`
	RgbColor *RgbColor `xml:"rgb,omitempty"`
	HslColor *HslColor `xml:"hsl,omitempty"`
}

func NewByColor() *ByColor {
	ret := &ByColor{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"by",
		"p",
	)
	return ret
}

func (m *ByColor) Clone() openxml.Element {
	ret := NewByColor()
	if m.RgbColor != nil {
		ret.RgbColor = m.RgbColor.Clone().(*RgbColor)
	}
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*HslColor)
	}
	return ret
}

func (m *ByColor) Validate() error {
	return nil
}

// From.
type FromColor struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main from"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewFromColor() *FromColor {
	ret := &FromColor{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"from",
		"p",
	)
	return ret
}

func (m *FromColor) Clone() openxml.Element {
	ret := NewFromColor()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *FromColor) Validate() error {
	return nil
}

// To.
type ToColor struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main to"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewToColor() *ToColor {
	ret := &ToColor{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"to",
		"p",
	)
	return ret
}

func (m *ToColor) Clone() openxml.Element {
	ret := NewToColor()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *ToColor) Validate() error {
	return nil
}

// Defines the Color3Type Class.
type Color3Type struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:""`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewColor3Type() *Color3Type {
	ret := &Color3Type{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *Color3Type) Clone() openxml.Element {
	ret := NewColor3Type()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *Color3Type) Validate() error {
	return nil
}

// Presentation Slide.
type SlideListEntry struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sld"`
	Id      *types.StringValue `xml:"r:id,attr,omitempty"`
}

func NewSlideListEntry() *SlideListEntry {
	ret := &SlideListEntry{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sld",
		"p",
	)
	return ret
}

func (m *SlideListEntry) Clone() openxml.Element {
	ret := NewSlideListEntry()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *SlideListEntry) Validate() error {
	return nil
}

// Customer Data.
type CustomerData struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main custData"`
	Id      *types.StringValue `xml:"r:id,attr,omitempty"`
}

func NewCustomerData() *CustomerData {
	ret := &CustomerData{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"custData",
		"p",
	)
	return ret
}

func (m *CustomerData) Clone() openxml.Element {
	ret := NewCustomerData()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *CustomerData) Validate() error {
	return nil
}

// Customer Data Tags.
type CustomerDataTags struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tags"`
	Id      *types.StringValue `xml:"r:id,attr,omitempty"`
}

func NewCustomerDataTags() *CustomerDataTags {
	ret := &CustomerDataTags{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"tags",
		"p",
	)
	return ret
}

func (m *CustomerDataTags) Clone() openxml.Element {
	ret := NewCustomerDataTags()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *CustomerDataTags) Validate() error {
	return nil
}

// Comment Author.
type CommentAuthor struct {
	*openxml.CompositeElementBase
	XMLName                    xml.Name                    `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cmAuthor"`
	Id                         *types.UInt32Value          `xml:"id,attr,omitempty"`
	Name                       *types.StringValue          `xml:"name,attr,omitempty"`
	Initials                   *types.StringValue          `xml:"initials,attr,omitempty"`
	LastIndex                  *types.UInt32Value          `xml:"lastIdx,attr,omitempty"`
	ColorIndex                 *types.UInt32Value          `xml:"clrIdx,attr,omitempty"`
	CommentAuthorExtensionList *CommentAuthorExtensionList `xml:"extLst,omitempty"`
}

func NewCommentAuthor() *CommentAuthor {
	ret := &CommentAuthor{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cmAuthor",
		"p",
	)
	return ret
}

func (m *CommentAuthor) Clone() openxml.Element {
	ret := NewCommentAuthor()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Initials != nil {
		v := *m.Initials
		ret.Initials = &v
	}
	if m.LastIndex != nil {
		v := *m.LastIndex
		ret.LastIndex = &v
	}
	if m.ColorIndex != nil {
		v := *m.ColorIndex
		ret.ColorIndex = &v
	}
	if m.CommentAuthorExtensionList != nil {
		ret.CommentAuthorExtensionList = m.CommentAuthorExtensionList.Clone().(*CommentAuthorExtensionList)
	}
	return ret
}

func (m *CommentAuthor) Validate() error {
	if m.CommentAuthorExtensionList != nil {
		if err := m.CommentAuthorExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Embedded Control.
type Control struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main control"`
	ShapeId       *types.StringValue  `xml:"spid,attr,omitempty"`
	Name          *types.StringValue  `xml:"name,attr,omitempty"`
	ShowAsIcon    *types.BooleanValue `xml:"showAsIcon,attr,omitempty"`
	Id            *types.StringValue  `xml:"r:id,attr,omitempty"`
	ImageWidth    *types.Int32Value   `xml:"imgW,attr,omitempty"`
	ImageHeight   *types.Int32Value   `xml:"imgH,attr,omitempty"`
	ExtensionList *ExtensionList      `xml:"extLst,omitempty"`
	Picture       *Picture            `xml:"pic,omitempty"`
}

func NewControl() *Control {
	ret := &Control{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"control",
		"p",
	)
	return ret
}

func (m *Control) Clone() openxml.Element {
	ret := NewControl()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.ShowAsIcon != nil {
		v := *m.ShowAsIcon
		ret.ShowAsIcon = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.ImageWidth != nil {
		v := *m.ImageWidth
		ret.ImageWidth = &v
	}
	if m.ImageHeight != nil {
		v := *m.ImageHeight
		ret.ImageHeight = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	if m.Picture != nil {
		ret.Picture = m.Picture.Clone().(*Picture)
	}
	return ret
}

func (m *Control) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Notes Master ID.
type NotesMasterId struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main notesMasterId"`
	Id            *types.StringValue `xml:"r:id,attr,omitempty"`
	ExtensionList *ExtensionList     `xml:"extLst,omitempty"`
}

func NewNotesMasterId() *NotesMasterId {
	ret := &NotesMasterId{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"notesMasterId",
		"p",
	)
	return ret
}

func (m *NotesMasterId) Clone() openxml.Element {
	ret := NewNotesMasterId()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *NotesMasterId) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Handout Master ID.
type HandoutMasterId struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main handoutMasterId"`
	Id            *types.StringValue `xml:"r:id,attr,omitempty"`
	ExtensionList *ExtensionList     `xml:"extLst,omitempty"`
}

func NewHandoutMasterId() *HandoutMasterId {
	ret := &HandoutMasterId{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"handoutMasterId",
		"p",
	)
	return ret
}

func (m *HandoutMasterId) Clone() openxml.Element {
	ret := NewHandoutMasterId()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *HandoutMasterId) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Embedded Font Name.
type Font struct {
	*openxml.LeafElementBase
	XMLName      xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main font"`
	Typeface     *types.StringValue    `xml:"typeface,attr,omitempty"`
	Panose       *types.HexBinaryValue `xml:"panose,attr,omitempty"`
	PitchFamily  *types.StringValue    `xml:"pitchFamily,attr,omitempty"`
	CharacterSet *types.StringValue    `xml:"charset,attr,omitempty"`
}

func NewFont() *Font {
	ret := &Font{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"font",
		"p",
	)
	return ret
}

func (m *Font) Clone() openxml.Element {
	ret := NewFont()
	if m.Typeface != nil {
		v := *m.Typeface
		ret.Typeface = &v
	}
	if m.Panose != nil {
		v := *m.Panose
		ret.Panose = &v
	}
	if m.PitchFamily != nil {
		v := *m.PitchFamily
		ret.PitchFamily = &v
	}
	if m.CharacterSet != nil {
		v := *m.CharacterSet
		ret.CharacterSet = &v
	}
	return ret
}

func (m *Font) Validate() error {
	return nil
}

// Regular Embedded Font.
type RegularFont struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main regular"`
}

func NewRegularFont() *RegularFont {
	ret := &RegularFont{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"regular",
		"p",
	)
	return ret
}

func (m *RegularFont) Clone() openxml.Element {
	ret := NewRegularFont()
	return ret
}

func (m *RegularFont) Validate() error {
	return nil
}

// Bold Embedded Font.
type BoldFont struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bold"`
}

func NewBoldFont() *BoldFont {
	ret := &BoldFont{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"bold",
		"p",
	)
	return ret
}

func (m *BoldFont) Clone() openxml.Element {
	ret := NewBoldFont()
	return ret
}

func (m *BoldFont) Validate() error {
	return nil
}

// Italic Embedded Font.
type ItalicFont struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main italic"`
}

func NewItalicFont() *ItalicFont {
	ret := &ItalicFont{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"italic",
		"p",
	)
	return ret
}

func (m *ItalicFont) Clone() openxml.Element {
	ret := NewItalicFont()
	return ret
}

func (m *ItalicFont) Validate() error {
	return nil
}

// Bold Italic Embedded Font.
type BoldItalicFont struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main boldItalic"`
}

func NewBoldItalicFont() *BoldItalicFont {
	ret := &BoldItalicFont{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"boldItalic",
		"p",
	)
	return ret
}

func (m *BoldItalicFont) Clone() openxml.Element {
	ret := NewBoldItalicFont()
	return ret
}

func (m *BoldItalicFont) Validate() error {
	return nil
}

// Defines the EmbeddedFontDataIdType Class.
type EmbeddedFontDataIdType struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:""`
	Id      *types.StringValue `xml:"r:id,attr,omitempty"`
}

func NewEmbeddedFontDataIdType() *EmbeddedFontDataIdType {
	ret := &EmbeddedFontDataIdType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *EmbeddedFontDataIdType) Clone() openxml.Element {
	ret := NewEmbeddedFontDataIdType()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	return ret
}

func (m *EmbeddedFontDataIdType) Validate() error {
	return nil
}

// Embedded Font.
type EmbeddedFont struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main embeddedFont"`
	Font           *Font           `xml:"font,omitempty"`
	RegularFont    *RegularFont    `xml:"regular,omitempty"`
	BoldFont       *BoldFont       `xml:"bold,omitempty"`
	ItalicFont     *ItalicFont     `xml:"italic,omitempty"`
	BoldItalicFont *BoldItalicFont `xml:"boldItalic,omitempty"`
}

func NewEmbeddedFont() *EmbeddedFont {
	ret := &EmbeddedFont{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"embeddedFont",
		"p",
	)
	return ret
}

func (m *EmbeddedFont) Clone() openxml.Element {
	ret := NewEmbeddedFont()
	if m.Font != nil {
		ret.Font = m.Font.Clone().(*Font)
	}
	if m.RegularFont != nil {
		ret.RegularFont = m.RegularFont.Clone().(*RegularFont)
	}
	if m.BoldFont != nil {
		ret.BoldFont = m.BoldFont.Clone().(*BoldFont)
	}
	if m.ItalicFont != nil {
		ret.ItalicFont = m.ItalicFont.Clone().(*ItalicFont)
	}
	if m.BoldItalicFont != nil {
		ret.BoldItalicFont = m.BoldItalicFont.Clone().(*BoldItalicFont)
	}
	return ret
}

func (m *EmbeddedFont) Validate() error {
	if m.Font != nil {
		if err := m.Font.Validate(); err != nil {
			return err
		}
	}
	if m.RegularFont != nil {
		if err := m.RegularFont.Validate(); err != nil {
			return err
		}
	}
	if m.BoldFont != nil {
		if err := m.BoldFont.Validate(); err != nil {
			return err
		}
	}
	if m.ItalicFont != nil {
		if err := m.ItalicFont.Validate(); err != nil {
			return err
		}
	}
	if m.BoldItalicFont != nil {
		if err := m.BoldItalicFont.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// List of Presentation Slides.
type SlideList struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldLst"`
	SlideListEntry *SlideListEntry `xml:"sld,omitempty"`
}

func NewSlideList() *SlideList {
	ret := &SlideList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldLst",
		"p",
	)
	return ret
}

func (m *SlideList) Clone() openxml.Element {
	ret := NewSlideList()
	if m.SlideListEntry != nil {
		ret.SlideListEntry = m.SlideListEntry.Clone().(*SlideListEntry)
	}
	return ret
}

func (m *SlideList) Validate() error {
	if m.SlideListEntry != nil {
		if err := m.SlideListEntry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Custom Show.
type CustomShow struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main custShow"`
	Name          *types.StringValue `xml:"name,attr,omitempty"`
	Id            *types.UInt32Value `xml:"id,attr,omitempty"`
	SlideList     *SlideList         `xml:"sldLst,omitempty"`
	ExtensionList *ExtensionList     `xml:"extLst,omitempty"`
}

func NewCustomShow() *CustomShow {
	ret := &CustomShow{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"custShow",
		"p",
	)
	return ret
}

func (m *CustomShow) Clone() openxml.Element {
	ret := NewCustomShow()
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.SlideList != nil {
		ret.SlideList = m.SlideList.Clone().(*SlideList)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *CustomShow) Validate() error {
	if m.SlideList != nil {
		if err := m.SlideList.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Non-Visual Drawing Properties for a Shape.
type NonVisualShapeDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cNvSpPr"`
	TextBox *types.BooleanValue `xml:"txBox,attr,omitempty"`
	// Skipped DrawingML type not yet implemented: ShapeLocks
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewNonVisualShapeDrawingProperties() *NonVisualShapeDrawingProperties {
	ret := &NonVisualShapeDrawingProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvSpPr",
		"p",
	)
	return ret
}

func (m *NonVisualShapeDrawingProperties) Clone() openxml.Element {
	ret := NewNonVisualShapeDrawingProperties()
	if m.TextBox != nil {
		v := *m.TextBox
		ret.TextBox = &v
	}
	return ret
}

func (m *NonVisualShapeDrawingProperties) Validate() error {
	return nil
}

// Non-Visual Connector Shape Drawing Properties.
type NonVisualConnectorShapeDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cNvCxnSpPr"`
	// Skipped DrawingML type not yet implemented: ConnectionShapeLocks
	// Skipped DrawingML type not yet implemented: StartConnection
	// Skipped DrawingML type not yet implemented: EndConnection
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewNonVisualConnectorShapeDrawingProperties() *NonVisualConnectorShapeDrawingProperties {
	ret := &NonVisualConnectorShapeDrawingProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvCxnSpPr",
		"p",
	)
	return ret
}

func (m *NonVisualConnectorShapeDrawingProperties) Clone() openxml.Element {
	ret := NewNonVisualConnectorShapeDrawingProperties()
	return ret
}

func (m *NonVisualConnectorShapeDrawingProperties) Validate() error {
	return nil
}

// Non-Visual Properties for a Connection Shape.
type NonVisualConnectionShapeProperties struct {
	*openxml.CompositeElementBase
	XMLName                                  xml.Name                                  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main nvCxnSpPr"`
	NonVisualDrawingProperties               *NonVisualDrawingProperties               `xml:"cNvPr,omitempty"`
	NonVisualConnectorShapeDrawingProperties *NonVisualConnectorShapeDrawingProperties `xml:"cNvCxnSpPr,omitempty"`
	ApplicationNonVisualDrawingProperties    *ApplicationNonVisualDrawingProperties    `xml:"nvPr,omitempty"`
}

func NewNonVisualConnectionShapeProperties() *NonVisualConnectionShapeProperties {
	ret := &NonVisualConnectionShapeProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"nvCxnSpPr",
		"p",
	)
	return ret
}

func (m *NonVisualConnectionShapeProperties) Clone() openxml.Element {
	ret := NewNonVisualConnectionShapeProperties()
	if m.NonVisualDrawingProperties != nil {
		ret.NonVisualDrawingProperties = m.NonVisualDrawingProperties.Clone().(*NonVisualDrawingProperties)
	}
	if m.NonVisualConnectorShapeDrawingProperties != nil {
		ret.NonVisualConnectorShapeDrawingProperties = m.NonVisualConnectorShapeDrawingProperties.Clone().(*NonVisualConnectorShapeDrawingProperties)
	}
	if m.ApplicationNonVisualDrawingProperties != nil {
		ret.ApplicationNonVisualDrawingProperties = m.ApplicationNonVisualDrawingProperties.Clone().(*ApplicationNonVisualDrawingProperties)
	}
	return ret
}

func (m *NonVisualConnectionShapeProperties) Validate() error {
	if m.NonVisualDrawingProperties != nil {
		if err := m.NonVisualDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	if m.NonVisualConnectorShapeDrawingProperties != nil {
		if err := m.NonVisualConnectorShapeDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ApplicationNonVisualDrawingProperties != nil {
		if err := m.ApplicationNonVisualDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Non-Visual Picture Drawing Properties.
type NonVisualPictureDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cNvPicPr"`
	PreferRelativeResize *types.BooleanValue `xml:"preferRelativeResize,attr,omitempty"`
	// Skipped DrawingML type not yet implemented: PictureLocks
	// Skipped DrawingML type not yet implemented: NonVisualPicturePropertiesExtensionList
}

func NewNonVisualPictureDrawingProperties() *NonVisualPictureDrawingProperties {
	ret := &NonVisualPictureDrawingProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvPicPr",
		"p",
	)
	return ret
}

func (m *NonVisualPictureDrawingProperties) Clone() openxml.Element {
	ret := NewNonVisualPictureDrawingProperties()
	if m.PreferRelativeResize != nil {
		v := *m.PreferRelativeResize
		ret.PreferRelativeResize = &v
	}
	return ret
}

func (m *NonVisualPictureDrawingProperties) Validate() error {
	return nil
}

// Non-Visual Graphic Frame Drawing Properties.
type NonVisualGraphicFrameDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cNvGraphicFramePr"`
	// Skipped DrawingML type not yet implemented: GraphicFrameLocks
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewNonVisualGraphicFrameDrawingProperties() *NonVisualGraphicFrameDrawingProperties {
	ret := &NonVisualGraphicFrameDrawingProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvGraphicFramePr",
		"p",
	)
	return ret
}

func (m *NonVisualGraphicFrameDrawingProperties) Clone() openxml.Element {
	ret := NewNonVisualGraphicFrameDrawingProperties()
	return ret
}

func (m *NonVisualGraphicFrameDrawingProperties) Validate() error {
	return nil
}

// Non-Visual Properties for a Graphic Frame.
type NonVisualGraphicFrameProperties struct {
	*openxml.CompositeElementBase
	XMLName                                xml.Name                                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main nvGraphicFramePr"`
	NonVisualDrawingProperties             *NonVisualDrawingProperties             `xml:"cNvPr,omitempty"`
	NonVisualGraphicFrameDrawingProperties *NonVisualGraphicFrameDrawingProperties `xml:"cNvGraphicFramePr,omitempty"`
	ApplicationNonVisualDrawingProperties  *ApplicationNonVisualDrawingProperties  `xml:"nvPr,omitempty"`
}

func NewNonVisualGraphicFrameProperties() *NonVisualGraphicFrameProperties {
	ret := &NonVisualGraphicFrameProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"nvGraphicFramePr",
		"p",
	)
	return ret
}

func (m *NonVisualGraphicFrameProperties) Clone() openxml.Element {
	ret := NewNonVisualGraphicFrameProperties()
	if m.NonVisualDrawingProperties != nil {
		ret.NonVisualDrawingProperties = m.NonVisualDrawingProperties.Clone().(*NonVisualDrawingProperties)
	}
	if m.NonVisualGraphicFrameDrawingProperties != nil {
		ret.NonVisualGraphicFrameDrawingProperties = m.NonVisualGraphicFrameDrawingProperties.Clone().(*NonVisualGraphicFrameDrawingProperties)
	}
	if m.ApplicationNonVisualDrawingProperties != nil {
		ret.ApplicationNonVisualDrawingProperties = m.ApplicationNonVisualDrawingProperties.Clone().(*ApplicationNonVisualDrawingProperties)
	}
	return ret
}

func (m *NonVisualGraphicFrameProperties) Validate() error {
	if m.NonVisualDrawingProperties != nil {
		if err := m.NonVisualDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	if m.NonVisualGraphicFrameDrawingProperties != nil {
		if err := m.NonVisualGraphicFrameDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ApplicationNonVisualDrawingProperties != nil {
		if err := m.ApplicationNonVisualDrawingProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// 2D Transform for Graphic Frame.
type Transform struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main xfrm"`
	Rotation       *types.Int32Value   `xml:"rot,attr,omitempty"`
	HorizontalFlip *types.BooleanValue `xml:"flipH,attr,omitempty"`
	VerticalFlip   *types.BooleanValue `xml:"flipV,attr,omitempty"`
	Offset         *drawingml.Offset   `xml:"off,omitempty"`
	// Skipped DrawingML type not yet implemented: Extents
}

func NewTransform() *Transform {
	ret := &Transform{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"xfrm",
		"p",
	)
	return ret
}

func (m *Transform) Clone() openxml.Element {
	ret := NewTransform()
	if m.Rotation != nil {
		v := *m.Rotation
		ret.Rotation = &v
	}
	if m.HorizontalFlip != nil {
		v := *m.HorizontalFlip
		ret.HorizontalFlip = &v
	}
	if m.VerticalFlip != nil {
		v := *m.VerticalFlip
		ret.VerticalFlip = &v
	}
	if m.Offset != nil {
		v := m.Offset.Clone()
		ret.Offset = &v
	}
	return ret
}

func (m *Transform) Validate() error {
	return nil
}

// Non-Visual Group Shape Drawing Properties.
type NonVisualGroupShapeDrawingProperties struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cNvGrpSpPr"`
	// Skipped DrawingML type not yet implemented: GroupShapeLocks
	// Skipped DrawingML type not yet implemented: NonVisualGroupDrawingShapePropsExtensionList
}

func NewNonVisualGroupShapeDrawingProperties() *NonVisualGroupShapeDrawingProperties {
	ret := &NonVisualGroupShapeDrawingProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cNvGrpSpPr",
		"p",
	)
	return ret
}

func (m *NonVisualGroupShapeDrawingProperties) Clone() openxml.Element {
	ret := NewNonVisualGroupShapeDrawingProperties()
	return ret
}

func (m *NonVisualGroupShapeDrawingProperties) Validate() error {
	return nil
}

// Slide Master Title Text Style.
type TitleStyle struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main titleStyle"`
	// Skipped DrawingML type not yet implemented: DefaultParagraphProperties
	// Skipped DrawingML type not yet implemented: Level1ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level2ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level3ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level4ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level5ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level6ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level7ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level8ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level9ParagraphProperties
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewTitleStyle() *TitleStyle {
	ret := &TitleStyle{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"titleStyle",
		"p",
	)
	return ret
}

func (m *TitleStyle) Clone() openxml.Element {
	ret := NewTitleStyle()
	return ret
}

func (m *TitleStyle) Validate() error {
	return nil
}

// Slide Master Body Text Style.
type BodyStyle struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bodyStyle"`
	// Skipped DrawingML type not yet implemented: DefaultParagraphProperties
	// Skipped DrawingML type not yet implemented: Level1ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level2ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level3ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level4ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level5ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level6ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level7ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level8ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level9ParagraphProperties
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewBodyStyle() *BodyStyle {
	ret := &BodyStyle{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bodyStyle",
		"p",
	)
	return ret
}

func (m *BodyStyle) Clone() openxml.Element {
	ret := NewBodyStyle()
	return ret
}

func (m *BodyStyle) Validate() error {
	return nil
}

// Slide Master Other Text Style.
type OtherStyle struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main otherStyle"`
	// Skipped DrawingML type not yet implemented: DefaultParagraphProperties
	// Skipped DrawingML type not yet implemented: Level1ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level2ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level3ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level4ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level5ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level6ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level7ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level8ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level9ParagraphProperties
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewOtherStyle() *OtherStyle {
	ret := &OtherStyle{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"otherStyle",
		"p",
	)
	return ret
}

func (m *OtherStyle) Clone() openxml.Element {
	ret := NewOtherStyle()
	return ret
}

func (m *OtherStyle) Validate() error {
	return nil
}

// Defines the NotesStyle Class.
type NotesStyle struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main notesStyle"`
	// Skipped DrawingML type not yet implemented: DefaultParagraphProperties
	// Skipped DrawingML type not yet implemented: Level1ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level2ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level3ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level4ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level5ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level6ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level7ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level8ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level9ParagraphProperties
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewNotesStyle() *NotesStyle {
	ret := &NotesStyle{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"notesStyle",
		"p",
	)
	return ret
}

func (m *NotesStyle) Clone() openxml.Element {
	ret := NewNotesStyle()
	return ret
}

func (m *NotesStyle) Validate() error {
	return nil
}

// Defines the TextListStyleType Class.
type TextListStyleType struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:""`
	// Skipped DrawingML type not yet implemented: DefaultParagraphProperties
	// Skipped DrawingML type not yet implemented: Level1ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level2ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level3ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level4ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level5ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level6ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level7ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level8ParagraphProperties
	// Skipped DrawingML type not yet implemented: Level9ParagraphProperties
	// Skipped DrawingML type not yet implemented: ExtensionList
}

func NewTextListStyleType() *TextListStyleType {
	ret := &TextListStyleType{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TextListStyleType) Clone() openxml.Element {
	ret := NewTextListStyleType()
	return ret
}

func (m *TextListStyleType) Validate() error {
	return nil
}

// Slide Layout Id.
type SlideLayoutId struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldLayoutId"`
	Id             *types.UInt32Value `xml:"id,attr,omitempty"`
	RelationshipId *types.StringValue `xml:"r:id,attr,omitempty"`
	ExtensionList  *ExtensionList     `xml:"extLst,omitempty"`
}

func NewSlideLayoutId() *SlideLayoutId {
	ret := &SlideLayoutId{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldLayoutId",
		"p",
	)
	return ret
}

func (m *SlideLayoutId) Clone() openxml.Element {
	ret := NewSlideLayoutId()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.RelationshipId != nil {
		v := *m.RelationshipId
		ret.RelationshipId = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *SlideLayoutId) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Programmable Extensibility Tag.
type Tag struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tag"`
	Name    *types.StringValue `xml:"name,attr,omitempty"`
	Val     *types.StringValue `xml:"val,attr,omitempty"`
}

func NewTag() *Tag {
	ret := &Tag{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"tag",
		"p",
	)
	return ret
}

func (m *Tag) Clone() openxml.Element {
	ret := NewTag()
	if m.Name != nil {
		v := *m.Name
		ret.Name = &v
	}
	if m.Val != nil {
		v := *m.Val
		ret.Val = &v
	}
	return ret
}

func (m *Tag) Validate() error {
	return nil
}

// Normal View Restored Left Properties.
type RestoredLeft struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main restoredLeft"`
}

func NewRestoredLeft() *RestoredLeft {
	ret := &RestoredLeft{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"restoredLeft",
		"p",
	)
	return ret
}

func (m *RestoredLeft) Clone() openxml.Element {
	ret := NewRestoredLeft()
	return ret
}

func (m *RestoredLeft) Validate() error {
	return nil
}

// Normal View Restored Top Properties.
type RestoredTop struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main restoredTop"`
}

func NewRestoredTop() *RestoredTop {
	ret := &RestoredTop{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"restoredTop",
		"p",
	)
	return ret
}

func (m *RestoredTop) Clone() openxml.Element {
	ret := NewRestoredTop()
	return ret
}

func (m *RestoredTop) Validate() error {
	return nil
}

// Defines the NormalViewPortionType Class.
type NormalViewPortionType struct {
	*openxml.LeafElementBase
	XMLName    xml.Name            `xml:""`
	Size       *types.Int32Value   `xml:"sz,attr,omitempty"`
	AutoAdjust *types.BooleanValue `xml:"autoAdjust,attr,omitempty"`
}

func NewNormalViewPortionType() *NormalViewPortionType {
	ret := &NormalViewPortionType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *NormalViewPortionType) Clone() openxml.Element {
	ret := NewNormalViewPortionType()
	if m.Size != nil {
		v := *m.Size
		ret.Size = &v
	}
	if m.AutoAdjust != nil {
		v := *m.AutoAdjust
		ret.AutoAdjust = &v
	}
	return ret
}

func (m *NormalViewPortionType) Validate() error {
	return nil
}

// View Scale.
type ScaleFactor struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main scale"`
	// Skipped DrawingML type not yet implemented: ScaleX
	// Skipped DrawingML type not yet implemented: ScaleY
}

func NewScaleFactor() *ScaleFactor {
	ret := &ScaleFactor{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"scale",
		"p",
	)
	return ret
}

func (m *ScaleFactor) Clone() openxml.Element {
	ret := NewScaleFactor()
	return ret
}

func (m *ScaleFactor) Validate() error {
	return nil
}

// View Origin.
type Origin struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main origin"`
}

func NewOrigin() *Origin {
	ret := &Origin{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"origin",
		"p",
	)
	return ret
}

func (m *Origin) Clone() openxml.Element {
	ret := NewOrigin()
	return ret
}

func (m *Origin) Validate() error {
	return nil
}

// Defines the Position Class.
type Position struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main pos"`
}

func NewPosition() *Position {
	ret := &Position{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pos",
		"p",
	)
	return ret
}

func (m *Position) Clone() openxml.Element {
	ret := NewPosition()
	return ret
}

func (m *Position) Validate() error {
	return nil
}

// Base properties for Notes View.
type CommonViewProperties struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cViewPr"`
	VariableScale *types.BooleanValue `xml:"varScale,attr,omitempty"`
	ScaleFactor   *ScaleFactor        `xml:"scale,omitempty"`
	Origin        *Origin             `xml:"origin,omitempty"`
}

func NewCommonViewProperties() *CommonViewProperties {
	ret := &CommonViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cViewPr",
		"p",
	)
	return ret
}

func (m *CommonViewProperties) Clone() openxml.Element {
	ret := NewCommonViewProperties()
	if m.VariableScale != nil {
		v := *m.VariableScale
		ret.VariableScale = &v
	}
	if m.ScaleFactor != nil {
		ret.ScaleFactor = m.ScaleFactor.Clone().(*ScaleFactor)
	}
	if m.Origin != nil {
		ret.Origin = m.Origin.Clone().(*Origin)
	}
	return ret
}

func (m *CommonViewProperties) Validate() error {
	if m.ScaleFactor != nil {
		if err := m.ScaleFactor.Validate(); err != nil {
			return err
		}
	}
	if m.Origin != nil {
		if err := m.Origin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Presentation Slide.
type OutlineViewSlideListEntry struct {
	*openxml.LeafElementBase
	XMLName  xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sld"`
	Id       *types.StringValue  `xml:"r:id,attr,omitempty"`
	Collapse *types.BooleanValue `xml:"collapse,attr,omitempty"`
}

func NewOutlineViewSlideListEntry() *OutlineViewSlideListEntry {
	ret := &OutlineViewSlideListEntry{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"sld",
		"p",
	)
	return ret
}

func (m *OutlineViewSlideListEntry) Clone() openxml.Element {
	ret := NewOutlineViewSlideListEntry()
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.Collapse != nil {
		v := *m.Collapse
		ret.Collapse = &v
	}
	return ret
}

func (m *OutlineViewSlideListEntry) Validate() error {
	return nil
}

// List of Presentation Slides.
type OutlineViewSlideList struct {
	*openxml.CompositeElementBase
	XMLName                   xml.Name                   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldLst"`
	OutlineViewSlideListEntry *OutlineViewSlideListEntry `xml:"sld,omitempty"`
}

func NewOutlineViewSlideList() *OutlineViewSlideList {
	ret := &OutlineViewSlideList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldLst",
		"p",
	)
	return ret
}

func (m *OutlineViewSlideList) Clone() openxml.Element {
	ret := NewOutlineViewSlideList()
	if m.OutlineViewSlideListEntry != nil {
		ret.OutlineViewSlideListEntry = m.OutlineViewSlideListEntry.Clone().(*OutlineViewSlideListEntry)
	}
	return ret
}

func (m *OutlineViewSlideList) Validate() error {
	if m.OutlineViewSlideListEntry != nil {
		if err := m.OutlineViewSlideListEntry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// A Guide.
type Guide struct {
	*openxml.LeafElementBase
	XMLName     xml.Name                          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main guide"`
	Orientation *types.EnumValue[DirectionValues] `xml:"orient,attr,omitempty"`
	Position    *types.Int32Value                 `xml:"pos,attr,omitempty"`
}

func NewGuide() *Guide {
	ret := &Guide{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"guide",
		"p",
	)
	return ret
}

func (m *Guide) Clone() openxml.Element {
	ret := NewGuide()
	if m.Orientation != nil {
		v := *m.Orientation
		ret.Orientation = &v
	}
	if m.Position != nil {
		v := *m.Position
		ret.Position = &v
	}
	return ret
}

func (m *Guide) Validate() error {
	return nil
}

// List of Guides.
type GuideList struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main guideLst"`
	Guide   *Guide   `xml:"guide,omitempty"`
}

func NewGuideList() *GuideList {
	ret := &GuideList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"guideLst",
		"p",
	)
	return ret
}

func (m *GuideList) Clone() openxml.Element {
	ret := NewGuideList()
	if m.Guide != nil {
		ret.Guide = m.Guide.Clone().(*Guide)
	}
	return ret
}

func (m *GuideList) Validate() error {
	if m.Guide != nil {
		if err := m.Guide.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommonSlideViewProperties Class.
type CommonSlideViewProperties struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cSldViewPr"`
	SnapToGrid           *types.BooleanValue   `xml:"snapToGrid,attr,omitempty"`
	SnapToObjects        *types.BooleanValue   `xml:"snapToObjects,attr,omitempty"`
	ShowGuides           *types.BooleanValue   `xml:"showGuides,attr,omitempty"`
	CommonViewProperties *CommonViewProperties `xml:"cViewPr,omitempty"`
	GuideList            *GuideList            `xml:"guideLst,omitempty"`
}

func NewCommonSlideViewProperties() *CommonSlideViewProperties {
	ret := &CommonSlideViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"cSldViewPr",
		"p",
	)
	return ret
}

func (m *CommonSlideViewProperties) Clone() openxml.Element {
	ret := NewCommonSlideViewProperties()
	if m.SnapToGrid != nil {
		v := *m.SnapToGrid
		ret.SnapToGrid = &v
	}
	if m.SnapToObjects != nil {
		v := *m.SnapToObjects
		ret.SnapToObjects = &v
	}
	if m.ShowGuides != nil {
		v := *m.ShowGuides
		ret.ShowGuides = &v
	}
	if m.CommonViewProperties != nil {
		ret.CommonViewProperties = m.CommonViewProperties.Clone().(*CommonViewProperties)
	}
	if m.GuideList != nil {
		ret.GuideList = m.GuideList.Clone().(*GuideList)
	}
	return ret
}

func (m *CommonSlideViewProperties) Validate() error {
	if m.CommonViewProperties != nil {
		if err := m.CommonViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.GuideList != nil {
		if err := m.GuideList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Normal View Properties.
type NormalViewProperties struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name                                 `xml:"http://schemas.openxmlformats.org/presentationml/2006/main normalViewPr"`
	ShowOutlineIcons     *types.BooleanValue                      `xml:"showOutlineIcons,attr,omitempty"`
	SnapVerticalSplitter *types.BooleanValue                      `xml:"snapVertSplitter,attr,omitempty"`
	VerticalBarState     *types.EnumValue[SplitterBarStateValues] `xml:"vertBarState,attr,omitempty"`
	HorizontalBarState   *types.EnumValue[SplitterBarStateValues] `xml:"horzBarState,attr,omitempty"`
	PreferSingleView     *types.BooleanValue                      `xml:"preferSingleView,attr,omitempty"`
	RestoredLeft         *RestoredLeft                            `xml:"restoredLeft,omitempty"`
	RestoredTop          *RestoredTop                             `xml:"restoredTop,omitempty"`
	ExtensionList        *ExtensionList                           `xml:"extLst,omitempty"`
}

func NewNormalViewProperties() *NormalViewProperties {
	ret := &NormalViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"normalViewPr",
		"p",
	)
	return ret
}

func (m *NormalViewProperties) Clone() openxml.Element {
	ret := NewNormalViewProperties()
	if m.ShowOutlineIcons != nil {
		v := *m.ShowOutlineIcons
		ret.ShowOutlineIcons = &v
	}
	if m.SnapVerticalSplitter != nil {
		v := *m.SnapVerticalSplitter
		ret.SnapVerticalSplitter = &v
	}
	if m.VerticalBarState != nil {
		v := *m.VerticalBarState
		ret.VerticalBarState = &v
	}
	if m.HorizontalBarState != nil {
		v := *m.HorizontalBarState
		ret.HorizontalBarState = &v
	}
	if m.PreferSingleView != nil {
		v := *m.PreferSingleView
		ret.PreferSingleView = &v
	}
	if m.RestoredLeft != nil {
		ret.RestoredLeft = m.RestoredLeft.Clone().(*RestoredLeft)
	}
	if m.RestoredTop != nil {
		ret.RestoredTop = m.RestoredTop.Clone().(*RestoredTop)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *NormalViewProperties) Validate() error {
	if m.RestoredLeft != nil {
		if err := m.RestoredLeft.Validate(); err != nil {
			return err
		}
	}
	if m.RestoredTop != nil {
		if err := m.RestoredTop.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Slide View Properties.
type SlideViewProperties struct {
	*openxml.CompositeElementBase
	XMLName                   xml.Name                   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main slideViewPr"`
	CommonSlideViewProperties *CommonSlideViewProperties `xml:"cSldViewPr,omitempty"`
	ExtensionList             *ExtensionList             `xml:"extLst,omitempty"`
}

func NewSlideViewProperties() *SlideViewProperties {
	ret := &SlideViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"slideViewPr",
		"p",
	)
	return ret
}

func (m *SlideViewProperties) Clone() openxml.Element {
	ret := NewSlideViewProperties()
	if m.CommonSlideViewProperties != nil {
		ret.CommonSlideViewProperties = m.CommonSlideViewProperties.Clone().(*CommonSlideViewProperties)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *SlideViewProperties) Validate() error {
	if m.CommonSlideViewProperties != nil {
		if err := m.CommonSlideViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Outline View Properties.
type OutlineViewProperties struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main outlineViewPr"`
	CommonViewProperties *CommonViewProperties `xml:"cViewPr,omitempty"`
	OutlineViewSlideList *OutlineViewSlideList `xml:"sldLst,omitempty"`
	ExtensionList        *ExtensionList        `xml:"extLst,omitempty"`
}

func NewOutlineViewProperties() *OutlineViewProperties {
	ret := &OutlineViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"outlineViewPr",
		"p",
	)
	return ret
}

func (m *OutlineViewProperties) Clone() openxml.Element {
	ret := NewOutlineViewProperties()
	if m.CommonViewProperties != nil {
		ret.CommonViewProperties = m.CommonViewProperties.Clone().(*CommonViewProperties)
	}
	if m.OutlineViewSlideList != nil {
		ret.OutlineViewSlideList = m.OutlineViewSlideList.Clone().(*OutlineViewSlideList)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *OutlineViewProperties) Validate() error {
	if m.CommonViewProperties != nil {
		if err := m.CommonViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.OutlineViewSlideList != nil {
		if err := m.OutlineViewSlideList.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Notes Text View Properties.
type NotesTextViewProperties struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main notesTextViewPr"`
	CommonViewProperties *CommonViewProperties `xml:"cViewPr,omitempty"`
	ExtensionList        *ExtensionList        `xml:"extLst,omitempty"`
}

func NewNotesTextViewProperties() *NotesTextViewProperties {
	ret := &NotesTextViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"notesTextViewPr",
		"p",
	)
	return ret
}

func (m *NotesTextViewProperties) Clone() openxml.Element {
	ret := NewNotesTextViewProperties()
	if m.CommonViewProperties != nil {
		ret.CommonViewProperties = m.CommonViewProperties.Clone().(*CommonViewProperties)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *NotesTextViewProperties) Validate() error {
	if m.CommonViewProperties != nil {
		if err := m.CommonViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Slide Sorter View Properties.
type SorterViewProperties struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sorterViewPr"`
	ShowFormatting       *types.BooleanValue   `xml:"showFormatting,attr,omitempty"`
	CommonViewProperties *CommonViewProperties `xml:"cViewPr,omitempty"`
	ExtensionList        *ExtensionList        `xml:"extLst,omitempty"`
}

func NewSorterViewProperties() *SorterViewProperties {
	ret := &SorterViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sorterViewPr",
		"p",
	)
	return ret
}

func (m *SorterViewProperties) Clone() openxml.Element {
	ret := NewSorterViewProperties()
	if m.ShowFormatting != nil {
		v := *m.ShowFormatting
		ret.ShowFormatting = &v
	}
	if m.CommonViewProperties != nil {
		ret.CommonViewProperties = m.CommonViewProperties.Clone().(*CommonViewProperties)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *SorterViewProperties) Validate() error {
	if m.CommonViewProperties != nil {
		if err := m.CommonViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Notes View Properties.
type NotesViewProperties struct {
	*openxml.CompositeElementBase
	XMLName                   xml.Name                   `xml:"http://schemas.openxmlformats.org/presentationml/2006/main notesViewPr"`
	CommonSlideViewProperties *CommonSlideViewProperties `xml:"cSldViewPr,omitempty"`
	ExtensionList             *ExtensionList             `xml:"extLst,omitempty"`
}

func NewNotesViewProperties() *NotesViewProperties {
	ret := &NotesViewProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"notesViewPr",
		"p",
	)
	return ret
}

func (m *NotesViewProperties) Clone() openxml.Element {
	ret := NewNotesViewProperties()
	if m.CommonSlideViewProperties != nil {
		ret.CommonSlideViewProperties = m.CommonSlideViewProperties.Clone().(*CommonSlideViewProperties)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *NotesViewProperties) Validate() error {
	if m.CommonSlideViewProperties != nil {
		if err := m.CommonSlideViewProperties.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Grid Spacing.
type GridSpacing struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main gridSpacing"`
}

func NewGridSpacing() *GridSpacing {
	ret := &GridSpacing{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"gridSpacing",
		"p",
	)
	return ret
}

func (m *GridSpacing) Clone() openxml.Element {
	ret := NewGridSpacing()
	return ret
}

func (m *GridSpacing) Validate() error {
	return nil
}

// Defines the PositiveSize2DType Class.
type PositiveSize2DType struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:""`
	Cx      *types.Int64Value `xml:"cx,attr,omitempty"`
	Cy      *types.Int64Value `xml:"cy,attr,omitempty"`
}

func NewPositiveSize2DType() *PositiveSize2DType {
	ret := &PositiveSize2DType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *PositiveSize2DType) Clone() openxml.Element {
	ret := NewPositiveSize2DType()
	if m.Cx != nil {
		v := *m.Cx
		ret.Cx = &v
	}
	if m.Cy != nil {
		v := *m.Cy
		ret.Cy = &v
	}
	return ret
}

func (m *PositiveSize2DType) Validate() error {
	return nil
}

// Defines the SlideExtension Class.
type SlideExtension struct {
	*openxml.CompositeElementBase
	XMLName             xml.Name             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri                 *types.StringValue   `xml:"uri,attr,omitempty"`
	LaserTraceList      *LaserTraceList      `xml:"laserTraceLst,omitempty"`
	ShowEventRecordList *ShowEventRecordList `xml:"showEvtLst,omitempty"`
	CommentRelationship *CommentRelationship `xml:"commentRel,omitempty"`
}

func NewSlideExtension() *SlideExtension {
	ret := &SlideExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *SlideExtension) Clone() openxml.Element {
	ret := NewSlideExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.LaserTraceList != nil {
		ret.LaserTraceList = m.LaserTraceList.Clone().(*LaserTraceList)
	}
	if m.ShowEventRecordList != nil {
		ret.ShowEventRecordList = m.ShowEventRecordList.Clone().(*ShowEventRecordList)
	}
	if m.CommentRelationship != nil {
		ret.CommentRelationship = m.CommentRelationship.Clone().(*CommentRelationship)
	}
	return ret
}

func (m *SlideExtension) Validate() error {
	if m.LaserTraceList != nil {
		if err := m.LaserTraceList.Validate(); err != nil {
			return err
		}
	}
	if m.ShowEventRecordList != nil {
		if err := m.ShowEventRecordList.Validate(); err != nil {
			return err
		}
	}
	if m.CommentRelationship != nil {
		if err := m.CommentRelationship.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommonSlideDataExtension Class.
type CommonSlideDataExtension struct {
	*openxml.CompositeElementBase
	XMLName    xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri        *types.StringValue `xml:"uri,attr,omitempty"`
	CreationId *CreationId        `xml:"creationId,omitempty"`
}

func NewCommonSlideDataExtension() *CommonSlideDataExtension {
	ret := &CommonSlideDataExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *CommonSlideDataExtension) Clone() openxml.Element {
	ret := NewCommonSlideDataExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.CreationId != nil {
		ret.CreationId = m.CreationId.Clone().(*CreationId)
	}
	return ret
}

func (m *CommonSlideDataExtension) Validate() error {
	if m.CreationId != nil {
		if err := m.CreationId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ShowPropertiesExtension Class.
type ShowPropertiesExtension struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri               *types.StringValue `xml:"uri,attr,omitempty"`
	BrowseMode        *BrowseMode        `xml:"browseMode,omitempty"`
	LaserColor        *LaserColor        `xml:"laserClr,omitempty"`
	ShowMediaControls *ShowMediaControls `xml:"showMediaCtrls,omitempty"`
}

func NewShowPropertiesExtension() *ShowPropertiesExtension {
	ret := &ShowPropertiesExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *ShowPropertiesExtension) Clone() openxml.Element {
	ret := NewShowPropertiesExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.BrowseMode != nil {
		ret.BrowseMode = m.BrowseMode.Clone().(*BrowseMode)
	}
	if m.LaserColor != nil {
		ret.LaserColor = m.LaserColor.Clone().(*LaserColor)
	}
	if m.ShowMediaControls != nil {
		ret.ShowMediaControls = m.ShowMediaControls.Clone().(*ShowMediaControls)
	}
	return ret
}

func (m *ShowPropertiesExtension) Validate() error {
	if m.BrowseMode != nil {
		if err := m.BrowseMode.Validate(); err != nil {
			return err
		}
	}
	if m.LaserColor != nil {
		if err := m.LaserColor.Validate(); err != nil {
			return err
		}
	}
	if m.ShowMediaControls != nil {
		if err := m.ShowMediaControls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the OleObjectEmbed Class.
type OleObjectEmbed struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name                                           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main embed"`
	FollowColorScheme *types.EnumValue[OleObjectFollowColorSchemeValues] `xml:"followColorScheme,attr,omitempty"`
	ExtensionList     *ExtensionList                                     `xml:"extLst,omitempty"`
}

func NewOleObjectEmbed() *OleObjectEmbed {
	ret := &OleObjectEmbed{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"embed",
		"p",
	)
	return ret
}

func (m *OleObjectEmbed) Clone() openxml.Element {
	ret := NewOleObjectEmbed()
	if m.FollowColorScheme != nil {
		v := *m.FollowColorScheme
		ret.FollowColorScheme = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *OleObjectEmbed) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the OleObjectLink Class.
type OleObjectLink struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name            `xml:"http://schemas.openxmlformats.org/presentationml/2006/main link"`
	AutoUpdate    *types.BooleanValue `xml:"updateAutomatic,attr,omitempty"`
	ExtensionList *ExtensionList      `xml:"extLst,omitempty"`
}

func NewOleObjectLink() *OleObjectLink {
	ret := &OleObjectLink{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"link",
		"p",
	)
	return ret
}

func (m *OleObjectLink) Clone() openxml.Element {
	ret := NewOleObjectLink()
	if m.AutoUpdate != nil {
		v := *m.AutoUpdate
		ret.AutoUpdate = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *OleObjectLink) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Slide Transition.
type Transition struct {
	*openxml.CompositeElementBase
	XMLName                       xml.Name                                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main transition"`
	Speed                         *types.EnumValue[TransitionSpeedValues] `xml:"spd,attr,omitempty"`
	Duration                      *types.StringValue                      `xml:"p14:dur,attr,omitempty"`
	AdvanceOnClick                *types.BooleanValue                     `xml:"advClick,attr,omitempty"`
	AdvanceAfterTime              *types.StringValue                      `xml:"advTm,attr,omitempty"`
	BlindsTransition              *BlindsTransition                       `xml:"blinds,omitempty"`
	CheckerTransition             *CheckerTransition                      `xml:"checker,omitempty"`
	CircleTransition              *CircleTransition                       `xml:"circle,omitempty"`
	DissolveTransition            *DissolveTransition                     `xml:"dissolve,omitempty"`
	CombTransition                *CombTransition                         `xml:"comb,omitempty"`
	CoverTransition               *CoverTransition                        `xml:"cover,omitempty"`
	CutTransition                 *CutTransition                          `xml:"cut,omitempty"`
	DiamondTransition             *DiamondTransition                      `xml:"diamond,omitempty"`
	FadeTransition                *FadeTransition                         `xml:"fade,omitempty"`
	NewsflashTransition           *NewsflashTransition                    `xml:"newsflash,omitempty"`
	PlusTransition                *PlusTransition                         `xml:"plus,omitempty"`
	PullTransition                *PullTransition                         `xml:"pull,omitempty"`
	PushTransition                *PushTransition                         `xml:"push,omitempty"`
	RandomTransition              *RandomTransition                       `xml:"random,omitempty"`
	RandomBarTransition           *RandomBarTransition                    `xml:"randomBar,omitempty"`
	SplitTransition               *SplitTransition                        `xml:"split,omitempty"`
	StripsTransition              *StripsTransition                       `xml:"strips,omitempty"`
	WedgeTransition               *WedgeTransition                        `xml:"wedge,omitempty"`
	WheelTransition               *WheelTransition                        `xml:"wheel,omitempty"`
	WipeTransition                *WipeTransition                         `xml:"wipe,omitempty"`
	ZoomTransition                *ZoomTransition                         `xml:"zoom,omitempty"`
	FlashTransition               *FlashTransition                        `xml:"flash,omitempty"`
	VortexTransition              *VortexTransition                       `xml:"vortex,omitempty"`
	SwitchTransition              *SwitchTransition                       `xml:"switch,omitempty"`
	FlipTransition                *FlipTransition                         `xml:"flip,omitempty"`
	RippleTransition              *RippleTransition                       `xml:"ripple,omitempty"`
	GlitterTransition             *GlitterTransition                      `xml:"glitter,omitempty"`
	HoneycombTransition           *HoneycombTransition                    `xml:"honeycomb,omitempty"`
	PrismTransition               *PrismTransition                        `xml:"prism,omitempty"`
	DoorsTransition               *DoorsTransition                        `xml:"doors,omitempty"`
	WindowTransition              *WindowTransition                       `xml:"window,omitempty"`
	ShredTransition               *ShredTransition                        `xml:"shred,omitempty"`
	FerrisTransition              *FerrisTransition                       `xml:"ferris,omitempty"`
	FlythroughTransition          *FlythroughTransition                   `xml:"flythrough,omitempty"`
	WarpTransition                *WarpTransition                         `xml:"warp,omitempty"`
	GalleryTransition             *GalleryTransition                      `xml:"gallery,omitempty"`
	ConveyorTransition            *ConveyorTransition                     `xml:"conveyor,omitempty"`
	PanTransition                 *PanTransition                          `xml:"pan,omitempty"`
	RevealTransition              *RevealTransition                       `xml:"reveal,omitempty"`
	WheelReverseTransition        *WheelReverseTransition                 `xml:"wheelReverse,omitempty"`
	PresetTransition              *PresetTransition                       `xml:"prstTrans,omitempty"`
	SoundAction                   *SoundAction                            `xml:"sndAc,omitempty"`
	ExtensionListWithModification *ExtensionListWithModification          `xml:"extLst,omitempty"`
}

func NewTransition() *Transition {
	ret := &Transition{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"transition",
		"p",
	)
	return ret
}

func (m *Transition) Clone() openxml.Element {
	ret := NewTransition()
	if m.Speed != nil {
		v := *m.Speed
		ret.Speed = &v
	}
	if m.Duration != nil {
		v := *m.Duration
		ret.Duration = &v
	}
	if m.AdvanceOnClick != nil {
		v := *m.AdvanceOnClick
		ret.AdvanceOnClick = &v
	}
	if m.AdvanceAfterTime != nil {
		v := *m.AdvanceAfterTime
		ret.AdvanceAfterTime = &v
	}
	if m.BlindsTransition != nil {
		ret.BlindsTransition = m.BlindsTransition.Clone().(*BlindsTransition)
	}
	if m.CheckerTransition != nil {
		ret.CheckerTransition = m.CheckerTransition.Clone().(*CheckerTransition)
	}
	if m.CircleTransition != nil {
		ret.CircleTransition = m.CircleTransition.Clone().(*CircleTransition)
	}
	if m.DissolveTransition != nil {
		ret.DissolveTransition = m.DissolveTransition.Clone().(*DissolveTransition)
	}
	if m.CombTransition != nil {
		ret.CombTransition = m.CombTransition.Clone().(*CombTransition)
	}
	if m.CoverTransition != nil {
		ret.CoverTransition = m.CoverTransition.Clone().(*CoverTransition)
	}
	if m.CutTransition != nil {
		ret.CutTransition = m.CutTransition.Clone().(*CutTransition)
	}
	if m.DiamondTransition != nil {
		ret.DiamondTransition = m.DiamondTransition.Clone().(*DiamondTransition)
	}
	if m.FadeTransition != nil {
		ret.FadeTransition = m.FadeTransition.Clone().(*FadeTransition)
	}
	if m.NewsflashTransition != nil {
		ret.NewsflashTransition = m.NewsflashTransition.Clone().(*NewsflashTransition)
	}
	if m.PlusTransition != nil {
		ret.PlusTransition = m.PlusTransition.Clone().(*PlusTransition)
	}
	if m.PullTransition != nil {
		ret.PullTransition = m.PullTransition.Clone().(*PullTransition)
	}
	if m.PushTransition != nil {
		ret.PushTransition = m.PushTransition.Clone().(*PushTransition)
	}
	if m.RandomTransition != nil {
		ret.RandomTransition = m.RandomTransition.Clone().(*RandomTransition)
	}
	if m.RandomBarTransition != nil {
		ret.RandomBarTransition = m.RandomBarTransition.Clone().(*RandomBarTransition)
	}
	if m.SplitTransition != nil {
		ret.SplitTransition = m.SplitTransition.Clone().(*SplitTransition)
	}
	if m.StripsTransition != nil {
		ret.StripsTransition = m.StripsTransition.Clone().(*StripsTransition)
	}
	if m.WedgeTransition != nil {
		ret.WedgeTransition = m.WedgeTransition.Clone().(*WedgeTransition)
	}
	if m.WheelTransition != nil {
		ret.WheelTransition = m.WheelTransition.Clone().(*WheelTransition)
	}
	if m.WipeTransition != nil {
		ret.WipeTransition = m.WipeTransition.Clone().(*WipeTransition)
	}
	if m.ZoomTransition != nil {
		ret.ZoomTransition = m.ZoomTransition.Clone().(*ZoomTransition)
	}
	if m.FlashTransition != nil {
		ret.FlashTransition = m.FlashTransition.Clone().(*FlashTransition)
	}
	if m.VortexTransition != nil {
		ret.VortexTransition = m.VortexTransition.Clone().(*VortexTransition)
	}
	if m.SwitchTransition != nil {
		ret.SwitchTransition = m.SwitchTransition.Clone().(*SwitchTransition)
	}
	if m.FlipTransition != nil {
		ret.FlipTransition = m.FlipTransition.Clone().(*FlipTransition)
	}
	if m.RippleTransition != nil {
		ret.RippleTransition = m.RippleTransition.Clone().(*RippleTransition)
	}
	if m.GlitterTransition != nil {
		ret.GlitterTransition = m.GlitterTransition.Clone().(*GlitterTransition)
	}
	if m.HoneycombTransition != nil {
		ret.HoneycombTransition = m.HoneycombTransition.Clone().(*HoneycombTransition)
	}
	if m.PrismTransition != nil {
		ret.PrismTransition = m.PrismTransition.Clone().(*PrismTransition)
	}
	if m.DoorsTransition != nil {
		ret.DoorsTransition = m.DoorsTransition.Clone().(*DoorsTransition)
	}
	if m.WindowTransition != nil {
		ret.WindowTransition = m.WindowTransition.Clone().(*WindowTransition)
	}
	if m.ShredTransition != nil {
		ret.ShredTransition = m.ShredTransition.Clone().(*ShredTransition)
	}
	if m.FerrisTransition != nil {
		ret.FerrisTransition = m.FerrisTransition.Clone().(*FerrisTransition)
	}
	if m.FlythroughTransition != nil {
		ret.FlythroughTransition = m.FlythroughTransition.Clone().(*FlythroughTransition)
	}
	if m.WarpTransition != nil {
		ret.WarpTransition = m.WarpTransition.Clone().(*WarpTransition)
	}
	if m.GalleryTransition != nil {
		ret.GalleryTransition = m.GalleryTransition.Clone().(*GalleryTransition)
	}
	if m.ConveyorTransition != nil {
		ret.ConveyorTransition = m.ConveyorTransition.Clone().(*ConveyorTransition)
	}
	if m.PanTransition != nil {
		ret.PanTransition = m.PanTransition.Clone().(*PanTransition)
	}
	if m.RevealTransition != nil {
		ret.RevealTransition = m.RevealTransition.Clone().(*RevealTransition)
	}
	if m.WheelReverseTransition != nil {
		ret.WheelReverseTransition = m.WheelReverseTransition.Clone().(*WheelReverseTransition)
	}
	if m.PresetTransition != nil {
		ret.PresetTransition = m.PresetTransition.Clone().(*PresetTransition)
	}
	if m.SoundAction != nil {
		ret.SoundAction = m.SoundAction.Clone().(*SoundAction)
	}
	if m.ExtensionListWithModification != nil {
		ret.ExtensionListWithModification = m.ExtensionListWithModification.Clone().(*ExtensionListWithModification)
	}
	return ret
}

func (m *Transition) Validate() error {
	if m.BlindsTransition != nil {
		if err := m.BlindsTransition.Validate(); err != nil {
			return err
		}
	}
	if m.CheckerTransition != nil {
		if err := m.CheckerTransition.Validate(); err != nil {
			return err
		}
	}
	if m.CircleTransition != nil {
		if err := m.CircleTransition.Validate(); err != nil {
			return err
		}
	}
	if m.DissolveTransition != nil {
		if err := m.DissolveTransition.Validate(); err != nil {
			return err
		}
	}
	if m.CombTransition != nil {
		if err := m.CombTransition.Validate(); err != nil {
			return err
		}
	}
	if m.CoverTransition != nil {
		if err := m.CoverTransition.Validate(); err != nil {
			return err
		}
	}
	if m.CutTransition != nil {
		if err := m.CutTransition.Validate(); err != nil {
			return err
		}
	}
	if m.DiamondTransition != nil {
		if err := m.DiamondTransition.Validate(); err != nil {
			return err
		}
	}
	if m.FadeTransition != nil {
		if err := m.FadeTransition.Validate(); err != nil {
			return err
		}
	}
	if m.NewsflashTransition != nil {
		if err := m.NewsflashTransition.Validate(); err != nil {
			return err
		}
	}
	if m.PlusTransition != nil {
		if err := m.PlusTransition.Validate(); err != nil {
			return err
		}
	}
	if m.PullTransition != nil {
		if err := m.PullTransition.Validate(); err != nil {
			return err
		}
	}
	if m.PushTransition != nil {
		if err := m.PushTransition.Validate(); err != nil {
			return err
		}
	}
	if m.RandomTransition != nil {
		if err := m.RandomTransition.Validate(); err != nil {
			return err
		}
	}
	if m.RandomBarTransition != nil {
		if err := m.RandomBarTransition.Validate(); err != nil {
			return err
		}
	}
	if m.SplitTransition != nil {
		if err := m.SplitTransition.Validate(); err != nil {
			return err
		}
	}
	if m.StripsTransition != nil {
		if err := m.StripsTransition.Validate(); err != nil {
			return err
		}
	}
	if m.WedgeTransition != nil {
		if err := m.WedgeTransition.Validate(); err != nil {
			return err
		}
	}
	if m.WheelTransition != nil {
		if err := m.WheelTransition.Validate(); err != nil {
			return err
		}
	}
	if m.WipeTransition != nil {
		if err := m.WipeTransition.Validate(); err != nil {
			return err
		}
	}
	if m.ZoomTransition != nil {
		if err := m.ZoomTransition.Validate(); err != nil {
			return err
		}
	}
	if m.FlashTransition != nil {
		if err := m.FlashTransition.Validate(); err != nil {
			return err
		}
	}
	if m.VortexTransition != nil {
		if err := m.VortexTransition.Validate(); err != nil {
			return err
		}
	}
	if m.SwitchTransition != nil {
		if err := m.SwitchTransition.Validate(); err != nil {
			return err
		}
	}
	if m.FlipTransition != nil {
		if err := m.FlipTransition.Validate(); err != nil {
			return err
		}
	}
	if m.RippleTransition != nil {
		if err := m.RippleTransition.Validate(); err != nil {
			return err
		}
	}
	if m.GlitterTransition != nil {
		if err := m.GlitterTransition.Validate(); err != nil {
			return err
		}
	}
	if m.HoneycombTransition != nil {
		if err := m.HoneycombTransition.Validate(); err != nil {
			return err
		}
	}
	if m.PrismTransition != nil {
		if err := m.PrismTransition.Validate(); err != nil {
			return err
		}
	}
	if m.DoorsTransition != nil {
		if err := m.DoorsTransition.Validate(); err != nil {
			return err
		}
	}
	if m.WindowTransition != nil {
		if err := m.WindowTransition.Validate(); err != nil {
			return err
		}
	}
	if m.ShredTransition != nil {
		if err := m.ShredTransition.Validate(); err != nil {
			return err
		}
	}
	if m.FerrisTransition != nil {
		if err := m.FerrisTransition.Validate(); err != nil {
			return err
		}
	}
	if m.FlythroughTransition != nil {
		if err := m.FlythroughTransition.Validate(); err != nil {
			return err
		}
	}
	if m.WarpTransition != nil {
		if err := m.WarpTransition.Validate(); err != nil {
			return err
		}
	}
	if m.GalleryTransition != nil {
		if err := m.GalleryTransition.Validate(); err != nil {
			return err
		}
	}
	if m.ConveyorTransition != nil {
		if err := m.ConveyorTransition.Validate(); err != nil {
			return err
		}
	}
	if m.PanTransition != nil {
		if err := m.PanTransition.Validate(); err != nil {
			return err
		}
	}
	if m.RevealTransition != nil {
		if err := m.RevealTransition.Validate(); err != nil {
			return err
		}
	}
	if m.WheelReverseTransition != nil {
		if err := m.WheelReverseTransition.Validate(); err != nil {
			return err
		}
	}
	if m.PresetTransition != nil {
		if err := m.PresetTransition.Validate(); err != nil {
			return err
		}
	}
	if m.SoundAction != nil {
		if err := m.SoundAction.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionListWithModification != nil {
		if err := m.ExtensionListWithModification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Slide Timing Information for a Slide.
type Timing struct {
	*openxml.CompositeElementBase
	XMLName                       xml.Name                       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main timing"`
	TimeNodeList                  *TimeNodeList                  `xml:"tnLst,omitempty"`
	BuildList                     *BuildList                     `xml:"bldLst,omitempty"`
	ExtensionListWithModification *ExtensionListWithModification `xml:"extLst,omitempty"`
}

func NewTiming() *Timing {
	ret := &Timing{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"timing",
		"p",
	)
	return ret
}

func (m *Timing) Clone() openxml.Element {
	ret := NewTiming()
	if m.TimeNodeList != nil {
		ret.TimeNodeList = m.TimeNodeList.Clone().(*TimeNodeList)
	}
	if m.BuildList != nil {
		ret.BuildList = m.BuildList.Clone().(*BuildList)
	}
	if m.ExtensionListWithModification != nil {
		ret.ExtensionListWithModification = m.ExtensionListWithModification.Clone().(*ExtensionListWithModification)
	}
	return ret
}

func (m *Timing) Validate() error {
	if m.ExtensionListWithModification != nil {
		if err := m.ExtensionListWithModification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideExtensionList Class.
type SlideExtensionList struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	SlideExtension *SlideExtension `xml:"ext,omitempty"`
}

func NewSlideExtensionList() *SlideExtensionList {
	ret := &SlideExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *SlideExtensionList) Clone() openxml.Element {
	ret := NewSlideExtensionList()
	if m.SlideExtension != nil {
		ret.SlideExtension = m.SlideExtension.Clone().(*SlideExtension)
	}
	return ret
}

func (m *SlideExtensionList) Validate() error {
	if m.SlideExtension != nil {
		if err := m.SlideExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Slide Background.
type Background struct {
	*openxml.CompositeElementBase
	XMLName                  xml.Name                               `xml:"http://schemas.openxmlformats.org/presentationml/2006/main bg"`
	BlackWhiteMode           *types.EnumValue[BlackWhiteModeValues] `xml:"bwMode,attr,omitempty"`
	BackgroundProperties     *BackgroundProperties                  `xml:"bgPr,omitempty"`
	BackgroundStyleReference *BackgroundStyleReference              `xml:"bgRef,omitempty"`
}

func NewBackground() *Background {
	ret := &Background{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"bg",
		"p",
	)
	return ret
}

func (m *Background) Clone() openxml.Element {
	ret := NewBackground()
	if m.BlackWhiteMode != nil {
		v := *m.BlackWhiteMode
		ret.BlackWhiteMode = &v
	}
	if m.BackgroundProperties != nil {
		ret.BackgroundProperties = m.BackgroundProperties.Clone().(*BackgroundProperties)
	}
	if m.BackgroundStyleReference != nil {
		ret.BackgroundStyleReference = m.BackgroundStyleReference.Clone().(*BackgroundStyleReference)
	}
	return ret
}

func (m *Background) Validate() error {
	if m.BackgroundProperties != nil {
		if err := m.BackgroundProperties.Validate(); err != nil {
			return err
		}
	}
	if m.BackgroundStyleReference != nil {
		if err := m.BackgroundStyleReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the GroupShapeType Class.
type GroupShapeType struct {
	*openxml.CompositeElementBase
	XMLName                       xml.Name                       `xml:""`
	NonVisualGroupShapeProperties *NonVisualGroupShapeProperties `xml:"nvGrpSpPr,omitempty"`
	GroupShapeProperties          *GroupShapeProperties          `xml:"grpSpPr,omitempty"`
	Shape                         *Shape                         `xml:"sp,omitempty"`
	GroupShape                    *GroupShape                    `xml:"grpSp,omitempty"`
	GraphicFrame                  *GraphicFrame                  `xml:"graphicFrame,omitempty"`
	ConnectionShape               *ConnectionShape               `xml:"cxnSp,omitempty"`
	Picture                       *Picture                       `xml:"pic,omitempty"`
	ContentPart                   *ContentPart                   `xml:"contentPart,omitempty"`
	ExtensionListWithModification *ExtensionListWithModification `xml:"extLst,omitempty"`
}

func NewGroupShapeType() *GroupShapeType {
	ret := &GroupShapeType{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *GroupShapeType) Clone() openxml.Element {
	ret := NewGroupShapeType()
	if m.NonVisualGroupShapeProperties != nil {
		ret.NonVisualGroupShapeProperties = m.NonVisualGroupShapeProperties.Clone().(*NonVisualGroupShapeProperties)
	}
	if m.GroupShapeProperties != nil {
		ret.GroupShapeProperties = m.GroupShapeProperties.Clone().(*GroupShapeProperties)
	}
	if m.Shape != nil {
		ret.Shape = m.Shape.Clone().(*Shape)
	}
	if m.GroupShape != nil {
		ret.GroupShape = m.GroupShape.Clone().(*GroupShape)
	}
	if m.GraphicFrame != nil {
		ret.GraphicFrame = m.GraphicFrame.Clone().(*GraphicFrame)
	}
	if m.ConnectionShape != nil {
		ret.ConnectionShape = m.ConnectionShape.Clone().(*ConnectionShape)
	}
	if m.Picture != nil {
		ret.Picture = m.Picture.Clone().(*Picture)
	}
	if m.ContentPart != nil {
		ret.ContentPart = m.ContentPart.Clone().(*ContentPart)
	}
	if m.ExtensionListWithModification != nil {
		ret.ExtensionListWithModification = m.ExtensionListWithModification.Clone().(*ExtensionListWithModification)
	}
	return ret
}

func (m *GroupShapeType) Validate() error {
	if m.ExtensionListWithModification != nil {
		if err := m.ExtensionListWithModification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Customer Data List.
type CustomerDataList struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main custDataLst"`
	CustomerData     *CustomerData     `xml:"custData,omitempty"`
	CustomerDataTags *CustomerDataTags `xml:"tags,omitempty"`
}

func NewCustomerDataList() *CustomerDataList {
	ret := &CustomerDataList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"custDataLst",
		"p",
	)
	return ret
}

func (m *CustomerDataList) Clone() openxml.Element {
	ret := NewCustomerDataList()
	if m.CustomerData != nil {
		ret.CustomerData = m.CustomerData.Clone().(*CustomerData)
	}
	if m.CustomerDataTags != nil {
		ret.CustomerDataTags = m.CustomerDataTags.Clone().(*CustomerDataTags)
	}
	return ret
}

func (m *CustomerDataList) Validate() error {
	if m.CustomerData != nil {
		if err := m.CustomerData.Validate(); err != nil {
			return err
		}
	}
	if m.CustomerDataTags != nil {
		if err := m.CustomerDataTags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommonSlideDataExtensionList Class.
type CommonSlideDataExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                  xml.Name                  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	CommonSlideDataExtension *CommonSlideDataExtension `xml:"ext,omitempty"`
}

func NewCommonSlideDataExtensionList() *CommonSlideDataExtensionList {
	ret := &CommonSlideDataExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *CommonSlideDataExtensionList) Clone() openxml.Element {
	ret := NewCommonSlideDataExtensionList()
	if m.CommonSlideDataExtension != nil {
		ret.CommonSlideDataExtension = m.CommonSlideDataExtension.Clone().(*CommonSlideDataExtension)
	}
	return ret
}

func (m *CommonSlideDataExtensionList) Validate() error {
	if m.CommonSlideDataExtension != nil {
		if err := m.CommonSlideDataExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ShowPropertiesExtensionList Class.
type ShowPropertiesExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                 xml.Name                 `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	ShowPropertiesExtension *ShowPropertiesExtension `xml:"ext,omitempty"`
}

func NewShowPropertiesExtensionList() *ShowPropertiesExtensionList {
	ret := &ShowPropertiesExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *ShowPropertiesExtensionList) Clone() openxml.Element {
	ret := NewShowPropertiesExtensionList()
	if m.ShowPropertiesExtension != nil {
		ret.ShowPropertiesExtension = m.ShowPropertiesExtension.Clone().(*ShowPropertiesExtension)
	}
	return ret
}

func (m *ShowPropertiesExtensionList) Validate() error {
	if m.ShowPropertiesExtension != nil {
		if err := m.ShowPropertiesExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Ink Target.
type InkTarget struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main inkTgt"`
}

func NewInkTarget() *InkTarget {
	ret := &InkTarget{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"inkTgt",
		"p",
	)
	return ret
}

func (m *InkTarget) Clone() openxml.Element {
	ret := NewInkTarget()
	return ret
}

func (m *InkTarget) Validate() error {
	return nil
}

// Subshape.
type SubShape struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main subSp"`
}

func NewSubShape() *SubShape {
	ret := &SubShape{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"subSp",
		"p",
	)
	return ret
}

func (m *SubShape) Clone() openxml.Element {
	ret := NewSubShape()
	return ret
}

func (m *SubShape) Validate() error {
	return nil
}

// Defines the TimeListSubShapeIdType Class.
type TimeListSubShapeIdType struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:""`
	ShapeId *types.StringValue `xml:"spid,attr,omitempty"`
}

func NewTimeListSubShapeIdType() *TimeListSubShapeIdType {
	ret := &TimeListSubShapeIdType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TimeListSubShapeIdType) Clone() openxml.Element {
	ret := NewTimeListSubShapeIdType()
	if m.ShapeId != nil {
		v := *m.ShapeId
		ret.ShapeId = &v
	}
	return ret
}

func (m *TimeListSubShapeIdType) Validate() error {
	return nil
}

// Defines the CommentAuthorExtension Class.
type CommentAuthorExtension struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri          *types.StringValue `xml:"uri,attr,omitempty"`
	PresenceInfo *PresenceInfo      `xml:"presenceInfo,omitempty"`
}

func NewCommentAuthorExtension() *CommentAuthorExtension {
	ret := &CommentAuthorExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *CommentAuthorExtension) Clone() openxml.Element {
	ret := NewCommentAuthorExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.PresenceInfo != nil {
		ret.PresenceInfo = m.PresenceInfo.Clone().(*PresenceInfo)
	}
	return ret
}

func (m *CommentAuthorExtension) Validate() error {
	if m.PresenceInfo != nil {
		if err := m.PresenceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommentExtension Class.
type CommentExtension struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri           *types.StringValue `xml:"uri,attr,omitempty"`
	ThreadingInfo *ThreadingInfo     `xml:"threadingInfo,omitempty"`
}

func NewCommentExtension() *CommentExtension {
	ret := &CommentExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *CommentExtension) Clone() openxml.Element {
	ret := NewCommentExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.ThreadingInfo != nil {
		ret.ThreadingInfo = m.ThreadingInfo.Clone().(*ThreadingInfo)
	}
	return ret
}

func (m *CommentExtension) Validate() error {
	if m.ThreadingInfo != nil {
		if err := m.ThreadingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideLayoutExtension Class.
type SlideLayoutExtension struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri            *types.StringValue `xml:"uri,attr,omitempty"`
	SlideGuideList *SlideGuideList    `xml:"sldGuideLst,omitempty"`
}

func NewSlideLayoutExtension() *SlideLayoutExtension {
	ret := &SlideLayoutExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *SlideLayoutExtension) Clone() openxml.Element {
	ret := NewSlideLayoutExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.SlideGuideList != nil {
		ret.SlideGuideList = m.SlideGuideList.Clone().(*SlideGuideList)
	}
	return ret
}

func (m *SlideLayoutExtension) Validate() error {
	if m.SlideGuideList != nil {
		if err := m.SlideGuideList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideMasterExtension Class.
type SlideMasterExtension struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri            *types.StringValue `xml:"uri,attr,omitempty"`
	SlideGuideList *SlideGuideList    `xml:"sldGuideLst,omitempty"`
}

func NewSlideMasterExtension() *SlideMasterExtension {
	ret := &SlideMasterExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *SlideMasterExtension) Clone() openxml.Element {
	ret := NewSlideMasterExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.SlideGuideList != nil {
		ret.SlideGuideList = m.SlideGuideList.Clone().(*SlideGuideList)
	}
	return ret
}

func (m *SlideMasterExtension) Validate() error {
	if m.SlideGuideList != nil {
		if err := m.SlideGuideList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the HandoutMasterExtension Class.
type HandoutMasterExtension struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri            *types.StringValue `xml:"uri,attr,omitempty"`
	SlideGuideList *SlideGuideList    `xml:"sldGuideLst,omitempty"`
}

func NewHandoutMasterExtension() *HandoutMasterExtension {
	ret := &HandoutMasterExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *HandoutMasterExtension) Clone() openxml.Element {
	ret := NewHandoutMasterExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.SlideGuideList != nil {
		ret.SlideGuideList = m.SlideGuideList.Clone().(*SlideGuideList)
	}
	return ret
}

func (m *HandoutMasterExtension) Validate() error {
	if m.SlideGuideList != nil {
		if err := m.SlideGuideList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the NotesMasterExtension Class.
type NotesMasterExtension struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri            *types.StringValue `xml:"uri,attr,omitempty"`
	SlideGuideList *SlideGuideList    `xml:"sldGuideLst,omitempty"`
}

func NewNotesMasterExtension() *NotesMasterExtension {
	ret := &NotesMasterExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *NotesMasterExtension) Clone() openxml.Element {
	ret := NewNotesMasterExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.SlideGuideList != nil {
		ret.SlideGuideList = m.SlideGuideList.Clone().(*SlideGuideList)
	}
	return ret
}

func (m *NotesMasterExtension) Validate() error {
	if m.SlideGuideList != nil {
		if err := m.SlideGuideList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Placeholder Shape.
type PlaceholderShape struct {
	*openxml.CompositeElementBase
	XMLName                       xml.Name                                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ph"`
	Type                          *types.EnumValue[PlaceholderValues]     `xml:"type,attr,omitempty"`
	Orientation                   *types.EnumValue[DirectionValues]       `xml:"orient,attr,omitempty"`
	Size                          *types.EnumValue[PlaceholderSizeValues] `xml:"sz,attr,omitempty"`
	Index                         *types.UInt32Value                      `xml:"idx,attr,omitempty"`
	HasCustomPrompt               *types.BooleanValue                     `xml:"hasCustomPrompt,attr,omitempty"`
	ExtensionListWithModification *ExtensionListWithModification          `xml:"extLst,omitempty"`
}

func NewPlaceholderShape() *PlaceholderShape {
	ret := &PlaceholderShape{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ph",
		"p",
	)
	return ret
}

func (m *PlaceholderShape) Clone() openxml.Element {
	ret := NewPlaceholderShape()
	if m.Type != nil {
		v := *m.Type
		ret.Type = &v
	}
	if m.Orientation != nil {
		v := *m.Orientation
		ret.Orientation = &v
	}
	if m.Size != nil {
		v := *m.Size
		ret.Size = &v
	}
	if m.Index != nil {
		v := *m.Index
		ret.Index = &v
	}
	if m.HasCustomPrompt != nil {
		v := *m.HasCustomPrompt
		ret.HasCustomPrompt = &v
	}
	if m.ExtensionListWithModification != nil {
		ret.ExtensionListWithModification = m.ExtensionListWithModification.Clone().(*ExtensionListWithModification)
	}
	return ret
}

func (m *PlaceholderShape) Validate() error {
	if m.ExtensionListWithModification != nil {
		if err := m.ExtensionListWithModification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ApplicationNonVisualDrawingPropertiesExtensionList Class.
type ApplicationNonVisualDrawingPropertiesExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                                        xml.Name                                        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	ApplicationNonVisualDrawingPropertiesExtension *ApplicationNonVisualDrawingPropertiesExtension `xml:"ext,omitempty"`
}

func NewApplicationNonVisualDrawingPropertiesExtensionList() *ApplicationNonVisualDrawingPropertiesExtensionList {
	ret := &ApplicationNonVisualDrawingPropertiesExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *ApplicationNonVisualDrawingPropertiesExtensionList) Clone() openxml.Element {
	ret := NewApplicationNonVisualDrawingPropertiesExtensionList()
	if m.ApplicationNonVisualDrawingPropertiesExtension != nil {
		ret.ApplicationNonVisualDrawingPropertiesExtension = m.ApplicationNonVisualDrawingPropertiesExtension.Clone().(*ApplicationNonVisualDrawingPropertiesExtension)
	}
	return ret
}

func (m *ApplicationNonVisualDrawingPropertiesExtensionList) Validate() error {
	if m.ApplicationNonVisualDrawingPropertiesExtension != nil {
		if err := m.ApplicationNonVisualDrawingPropertiesExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ApplicationNonVisualDrawingPropertiesExtension Class.
type ApplicationNonVisualDrawingPropertiesExtension struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri            *types.StringValue `xml:"uri,attr,omitempty"`
	Media          *Media             `xml:"media,omitempty"`
	ModificationId *ModificationId    `xml:"modId,omitempty"`
}

func NewApplicationNonVisualDrawingPropertiesExtension() *ApplicationNonVisualDrawingPropertiesExtension {
	ret := &ApplicationNonVisualDrawingPropertiesExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *ApplicationNonVisualDrawingPropertiesExtension) Clone() openxml.Element {
	ret := NewApplicationNonVisualDrawingPropertiesExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.Media != nil {
		ret.Media = m.Media.Clone().(*Media)
	}
	if m.ModificationId != nil {
		ret.ModificationId = m.ModificationId.Clone().(*ModificationId)
	}
	return ret
}

func (m *ApplicationNonVisualDrawingPropertiesExtension) Validate() error {
	if m.Media != nil {
		if err := m.Media.Validate(); err != nil {
			return err
		}
	}
	if m.ModificationId != nil {
		if err := m.ModificationId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Iterate Class.
type Iterate struct {
	*openxml.CompositeElementBase
	XMLName        xml.Name                        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main iterate"`
	Type           *types.EnumValue[IterateValues] `xml:"type,attr,omitempty"`
	Backwards      *types.BooleanValue             `xml:"backwards,attr,omitempty"`
	TimeAbsolute   *TimeAbsolute                   `xml:"tmAbs,omitempty"`
	TimePercentage *TimePercentage                 `xml:"tmPct,omitempty"`
}

func NewIterate() *Iterate {
	ret := &Iterate{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"iterate",
		"p",
	)
	return ret
}

func (m *Iterate) Clone() openxml.Element {
	ret := NewIterate()
	if m.Type != nil {
		v := *m.Type
		ret.Type = &v
	}
	if m.Backwards != nil {
		v := *m.Backwards
		ret.Backwards = &v
	}
	if m.TimeAbsolute != nil {
		ret.TimeAbsolute = m.TimeAbsolute.Clone().(*TimeAbsolute)
	}
	if m.TimePercentage != nil {
		ret.TimePercentage = m.TimePercentage.Clone().(*TimePercentage)
	}
	return ret
}

func (m *Iterate) Validate() error {
	if m.TimeAbsolute != nil {
		if err := m.TimeAbsolute.Validate(); err != nil {
			return err
		}
	}
	if m.TimePercentage != nil {
		if err := m.TimePercentage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SubTimeNodeList Class.
type SubTimeNodeList struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main subTnLst"`
	ParallelTimeNode  *ParallelTimeNode  `xml:"par,omitempty"`
	SequenceTimeNode  *SequenceTimeNode  `xml:"seq,omitempty"`
	ExclusiveTimeNode *ExclusiveTimeNode `xml:"excl,omitempty"`
	Animate           *Animate           `xml:"anim,omitempty"`
	AnimateColor      *AnimateColor      `xml:"animClr,omitempty"`
	AnimateEffect     *AnimateEffect     `xml:"animEffect,omitempty"`
	AnimateMotion     *AnimateMotion     `xml:"animMotion,omitempty"`
	AnimateRotation   *AnimateRotation   `xml:"animRot,omitempty"`
	AnimateScale      *AnimateScale      `xml:"animScale,omitempty"`
	Command           *Command           `xml:"cmd,omitempty"`
	SetBehavior       *SetBehavior       `xml:"set,omitempty"`
	Audio             *Audio             `xml:"audio,omitempty"`
	Video             *Video             `xml:"video,omitempty"`
}

func NewSubTimeNodeList() *SubTimeNodeList {
	ret := &SubTimeNodeList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"subTnLst",
		"p",
	)
	return ret
}

func (m *SubTimeNodeList) Clone() openxml.Element {
	ret := NewSubTimeNodeList()
	if m.ParallelTimeNode != nil {
		ret.ParallelTimeNode = m.ParallelTimeNode.Clone().(*ParallelTimeNode)
	}
	if m.SequenceTimeNode != nil {
		ret.SequenceTimeNode = m.SequenceTimeNode.Clone().(*SequenceTimeNode)
	}
	if m.ExclusiveTimeNode != nil {
		ret.ExclusiveTimeNode = m.ExclusiveTimeNode.Clone().(*ExclusiveTimeNode)
	}
	if m.Animate != nil {
		ret.Animate = m.Animate.Clone().(*Animate)
	}
	if m.AnimateColor != nil {
		ret.AnimateColor = m.AnimateColor.Clone().(*AnimateColor)
	}
	if m.AnimateEffect != nil {
		ret.AnimateEffect = m.AnimateEffect.Clone().(*AnimateEffect)
	}
	if m.AnimateMotion != nil {
		ret.AnimateMotion = m.AnimateMotion.Clone().(*AnimateMotion)
	}
	if m.AnimateRotation != nil {
		ret.AnimateRotation = m.AnimateRotation.Clone().(*AnimateRotation)
	}
	if m.AnimateScale != nil {
		ret.AnimateScale = m.AnimateScale.Clone().(*AnimateScale)
	}
	if m.Command != nil {
		ret.Command = m.Command.Clone().(*Command)
	}
	if m.SetBehavior != nil {
		ret.SetBehavior = m.SetBehavior.Clone().(*SetBehavior)
	}
	if m.Audio != nil {
		ret.Audio = m.Audio.Clone().(*Audio)
	}
	if m.Video != nil {
		ret.Video = m.Video.Clone().(*Video)
	}
	return ret
}

func (m *SubTimeNodeList) Validate() error {
	if m.ExclusiveTimeNode != nil {
		if err := m.ExclusiveTimeNode.Validate(); err != nil {
			return err
		}
	}
	if m.Command != nil {
		if err := m.Command.Validate(); err != nil {
			return err
		}
	}
	if m.SetBehavior != nil {
		if err := m.SetBehavior.Validate(); err != nil {
			return err
		}
	}
	if m.Audio != nil {
		if err := m.Audio.Validate(); err != nil {
			return err
		}
	}
	if m.Video != nil {
		if err := m.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TimeTypeListType Class.
type TimeTypeListType struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name           `xml:""`
	ParallelTimeNode  *ParallelTimeNode  `xml:"par,omitempty"`
	SequenceTimeNode  *SequenceTimeNode  `xml:"seq,omitempty"`
	ExclusiveTimeNode *ExclusiveTimeNode `xml:"excl,omitempty"`
	Animate           *Animate           `xml:"anim,omitempty"`
	AnimateColor      *AnimateColor      `xml:"animClr,omitempty"`
	AnimateEffect     *AnimateEffect     `xml:"animEffect,omitempty"`
	AnimateMotion     *AnimateMotion     `xml:"animMotion,omitempty"`
	AnimateRotation   *AnimateRotation   `xml:"animRot,omitempty"`
	AnimateScale      *AnimateScale      `xml:"animScale,omitempty"`
	Command           *Command           `xml:"cmd,omitempty"`
	SetBehavior       *SetBehavior       `xml:"set,omitempty"`
	Audio             *Audio             `xml:"audio,omitempty"`
	Video             *Video             `xml:"video,omitempty"`
}

func NewTimeTypeListType() *TimeTypeListType {
	ret := &TimeTypeListType{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TimeTypeListType) Clone() openxml.Element {
	ret := NewTimeTypeListType()
	if m.ParallelTimeNode != nil {
		ret.ParallelTimeNode = m.ParallelTimeNode.Clone().(*ParallelTimeNode)
	}
	if m.SequenceTimeNode != nil {
		ret.SequenceTimeNode = m.SequenceTimeNode.Clone().(*SequenceTimeNode)
	}
	if m.ExclusiveTimeNode != nil {
		ret.ExclusiveTimeNode = m.ExclusiveTimeNode.Clone().(*ExclusiveTimeNode)
	}
	if m.Animate != nil {
		ret.Animate = m.Animate.Clone().(*Animate)
	}
	if m.AnimateColor != nil {
		ret.AnimateColor = m.AnimateColor.Clone().(*AnimateColor)
	}
	if m.AnimateEffect != nil {
		ret.AnimateEffect = m.AnimateEffect.Clone().(*AnimateEffect)
	}
	if m.AnimateMotion != nil {
		ret.AnimateMotion = m.AnimateMotion.Clone().(*AnimateMotion)
	}
	if m.AnimateRotation != nil {
		ret.AnimateRotation = m.AnimateRotation.Clone().(*AnimateRotation)
	}
	if m.AnimateScale != nil {
		ret.AnimateScale = m.AnimateScale.Clone().(*AnimateScale)
	}
	if m.Command != nil {
		ret.Command = m.Command.Clone().(*Command)
	}
	if m.SetBehavior != nil {
		ret.SetBehavior = m.SetBehavior.Clone().(*SetBehavior)
	}
	if m.Audio != nil {
		ret.Audio = m.Audio.Clone().(*Audio)
	}
	if m.Video != nil {
		ret.Video = m.Video.Clone().(*Video)
	}
	return ret
}

func (m *TimeTypeListType) Validate() error {
	if m.ExclusiveTimeNode != nil {
		if err := m.ExclusiveTimeNode.Validate(); err != nil {
			return err
		}
	}
	if m.Command != nil {
		if err := m.Command.Validate(); err != nil {
			return err
		}
	}
	if m.SetBehavior != nil {
		if err := m.SetBehavior.Validate(); err != nil {
			return err
		}
	}
	if m.Audio != nil {
		if err := m.Audio.Validate(); err != nil {
			return err
		}
	}
	if m.Video != nil {
		if err := m.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TimeAnimateValueList Class.
type TimeAnimateValueList struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main tavLst"`
	TimeAnimateValue *TimeAnimateValue `xml:"tav,omitempty"`
}

func NewTimeAnimateValueList() *TimeAnimateValueList {
	ret := &TimeAnimateValueList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"tavLst",
		"p",
	)
	return ret
}

func (m *TimeAnimateValueList) Clone() openxml.Element {
	ret := NewTimeAnimateValueList()
	if m.TimeAnimateValue != nil {
		ret.TimeAnimateValue = m.TimeAnimateValue.Clone().(*TimeAnimateValue)
	}
	return ret
}

func (m *TimeAnimateValueList) Validate() error {
	if m.TimeAnimateValue != nil {
		if err := m.TimeAnimateValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ByPosition Class.
type ByPosition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main by"`
}

func NewByPosition() *ByPosition {
	ret := &ByPosition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"by",
		"p",
	)
	return ret
}

func (m *ByPosition) Clone() openxml.Element {
	ret := NewByPosition()
	return ret
}

func (m *ByPosition) Validate() error {
	return nil
}

// Defines the FromPosition Class.
type FromPosition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main from"`
}

func NewFromPosition() *FromPosition {
	ret := &FromPosition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"from",
		"p",
	)
	return ret
}

func (m *FromPosition) Clone() openxml.Element {
	ret := NewFromPosition()
	return ret
}

func (m *FromPosition) Validate() error {
	return nil
}

// Defines the ToPosition Class.
type ToPosition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main to"`
}

func NewToPosition() *ToPosition {
	ret := &ToPosition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"to",
		"p",
	)
	return ret
}

func (m *ToPosition) Clone() openxml.Element {
	ret := NewToPosition()
	return ret
}

func (m *ToPosition) Validate() error {
	return nil
}

// Defines the RotationCenter Class.
type RotationCenter struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main rCtr"`
}

func NewRotationCenter() *RotationCenter {
	ret := &RotationCenter{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"rCtr",
		"p",
	)
	return ret
}

func (m *RotationCenter) Clone() openxml.Element {
	ret := NewRotationCenter()
	return ret
}

func (m *RotationCenter) Validate() error {
	return nil
}

// Defines the TimeListType Class.
type TimeListType struct {
	*openxml.LeafElementBase
	XMLName xml.Name          `xml:""`
	X       *types.Int32Value `xml:"x,attr,omitempty"`
	Y       *types.Int32Value `xml:"y,attr,omitempty"`
}

func NewTimeListType() *TimeListType {
	ret := &TimeListType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *TimeListType) Clone() openxml.Element {
	ret := NewTimeListType()
	if m.X != nil {
		v := *m.X
		ret.X = &v
	}
	if m.Y != nil {
		v := *m.Y
		ret.Y = &v
	}
	return ret
}

func (m *TimeListType) Validate() error {
	return nil
}

// Defines the CommentAuthorExtensionList Class.
type CommentAuthorExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                xml.Name                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	CommentAuthorExtension *CommentAuthorExtension `xml:"ext,omitempty"`
}

func NewCommentAuthorExtensionList() *CommentAuthorExtensionList {
	ret := &CommentAuthorExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *CommentAuthorExtensionList) Clone() openxml.Element {
	ret := NewCommentAuthorExtensionList()
	if m.CommentAuthorExtension != nil {
		ret.CommentAuthorExtension = m.CommentAuthorExtension.Clone().(*CommentAuthorExtension)
	}
	return ret
}

func (m *CommentAuthorExtensionList) Validate() error {
	if m.CommentAuthorExtension != nil {
		if err := m.CommentAuthorExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CommentExtensionList Class.
type CommentExtensionList struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	CommentExtension *CommentExtension `xml:"ext,omitempty"`
}

func NewCommentExtensionList() *CommentExtensionList {
	ret := &CommentExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *CommentExtensionList) Clone() openxml.Element {
	ret := NewCommentExtensionList()
	if m.CommentExtension != nil {
		ret.CommentExtension = m.CommentExtension.Clone().(*CommentExtension)
	}
	return ret
}

func (m *CommentExtensionList) Validate() error {
	if m.CommentExtension != nil {
		if err := m.CommentExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the EmbeddedFontList Class.
type EmbeddedFontList struct {
	*openxml.CompositeElementBase
	XMLName      xml.Name      `xml:"http://schemas.openxmlformats.org/presentationml/2006/main embeddedFontLst"`
	EmbeddedFont *EmbeddedFont `xml:"embeddedFont,omitempty"`
}

func NewEmbeddedFontList() *EmbeddedFontList {
	ret := &EmbeddedFontList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"embeddedFontLst",
		"p",
	)
	return ret
}

func (m *EmbeddedFontList) Clone() openxml.Element {
	ret := NewEmbeddedFontList()
	if m.EmbeddedFont != nil {
		ret.EmbeddedFont = m.EmbeddedFont.Clone().(*EmbeddedFont)
	}
	return ret
}

func (m *EmbeddedFontList) Validate() error {
	if m.EmbeddedFont != nil {
		if err := m.EmbeddedFont.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the CustomShowList Class.
type CustomShowList struct {
	*openxml.CompositeElementBase
	XMLName    xml.Name    `xml:"http://schemas.openxmlformats.org/presentationml/2006/main custShowLst"`
	CustomShow *CustomShow `xml:"custShow,omitempty"`
}

func NewCustomShowList() *CustomShowList {
	ret := &CustomShowList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"custShowLst",
		"p",
	)
	return ret
}

func (m *CustomShowList) Clone() openxml.Element {
	ret := NewCustomShowList()
	if m.CustomShow != nil {
		ret.CustomShow = m.CustomShow.Clone().(*CustomShow)
	}
	return ret
}

func (m *CustomShowList) Validate() error {
	if m.CustomShow != nil {
		if err := m.CustomShow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the PhotoAlbum Class.
type PhotoAlbum struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name                                     `xml:"http://schemas.openxmlformats.org/presentationml/2006/main photoAlbum"`
	BlackWhite    *types.BooleanValue                          `xml:"bw,attr,omitempty"`
	ShowCaptions  *types.BooleanValue                          `xml:"showCaptions,attr,omitempty"`
	Layout        *types.EnumValue[PhotoAlbumLayoutValues]     `xml:"layout,attr,omitempty"`
	Frame         *types.EnumValue[PhotoAlbumFrameShapeValues] `xml:"frame,attr,omitempty"`
	ExtensionList *ExtensionList                               `xml:"extLst,omitempty"`
}

func NewPhotoAlbum() *PhotoAlbum {
	ret := &PhotoAlbum{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"photoAlbum",
		"p",
	)
	return ret
}

func (m *PhotoAlbum) Clone() openxml.Element {
	ret := NewPhotoAlbum()
	if m.BlackWhite != nil {
		v := *m.BlackWhite
		ret.BlackWhite = &v
	}
	if m.ShowCaptions != nil {
		v := *m.ShowCaptions
		ret.ShowCaptions = &v
	}
	if m.Layout != nil {
		v := *m.Layout
		ret.Layout = &v
	}
	if m.Frame != nil {
		v := *m.Frame
		ret.Frame = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *PhotoAlbum) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the Kinsoku Class.
type Kinsoku struct {
	*openxml.LeafElementBase
	XMLName           xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main kinsoku"`
	Language          *types.StringValue `xml:"lang,attr,omitempty"`
	InvalidStartChars *types.StringValue `xml:"invalStChars,attr,omitempty"`
	InvalidEndChars   *types.StringValue `xml:"invalEndChars,attr,omitempty"`
}

func NewKinsoku() *Kinsoku {
	ret := &Kinsoku{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"kinsoku",
		"p",
	)
	return ret
}

func (m *Kinsoku) Clone() openxml.Element {
	ret := NewKinsoku()
	if m.Language != nil {
		v := *m.Language
		ret.Language = &v
	}
	if m.InvalidStartChars != nil {
		v := *m.InvalidStartChars
		ret.InvalidStartChars = &v
	}
	if m.InvalidEndChars != nil {
		v := *m.InvalidEndChars
		ret.InvalidEndChars = &v
	}
	return ret
}

func (m *Kinsoku) Validate() error {
	return nil
}

// Defines the ModificationVerifier Class.
type ModificationVerifier struct {
	*openxml.LeafElementBase
	XMLName                                      xml.Name                                    `xml:"http://schemas.openxmlformats.org/presentationml/2006/main modifyVerifier"`
	CryptographicProviderType                    *types.EnumValue[CryptProviderValues]       `xml:"cryptProviderType,attr,omitempty"`
	CryptographicAlgorithmClass                  *types.EnumValue[CryptAlgorithmClassValues] `xml:"cryptAlgorithmClass,attr,omitempty"`
	CryptographicAlgorithmType                   *types.EnumValue[CryptAlgorithmValues]      `xml:"cryptAlgorithmType,attr,omitempty"`
	CryptographicAlgorithmSid                    *types.UInt32Value                          `xml:"cryptAlgorithmSid,attr,omitempty"`
	SpinCount                                    *types.UInt32Value                          `xml:"spinCount,attr,omitempty"`
	SaltData                                     *types.Base64BinaryValue                    `xml:"saltData,attr,omitempty"`
	HashData                                     *types.StringValue                          `xml:"hashData,attr,omitempty"`
	CryptographicProvider                        *types.StringValue                          `xml:"cryptProvider,attr,omitempty"`
	ExtendedCryptographicAlgorithm               *types.UInt32Value                          `xml:"algIdExt,attr,omitempty"`
	ExtendedCryptographicAlgorithmSource         *types.StringValue                          `xml:"algIdExtSource,attr,omitempty"`
	CryptographicProviderTypeExtensibility       *types.UInt32Value                          `xml:"cryptProviderTypeExt,attr,omitempty"`
	CryptographicProviderTypeExtensibilitySource *types.StringValue                          `xml:"cryptProviderTypeExtSource,attr,omitempty"`
	AlgorithmName                                *types.StringValue                          `xml:"algorithmName,attr,omitempty"`
	HashValue                                    *types.Base64BinaryValue                    `xml:"hashValue,attr,omitempty"`
	SaltValue                                    *types.Base64BinaryValue                    `xml:"saltValue,attr,omitempty"`
	SpinValue                                    *types.UInt32Value                          `xml:"spinValue,attr,omitempty"`
}

func NewModificationVerifier() *ModificationVerifier {
	ret := &ModificationVerifier{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"modifyVerifier",
		"p",
	)
	return ret
}

func (m *ModificationVerifier) Clone() openxml.Element {
	ret := NewModificationVerifier()
	if m.CryptographicProviderType != nil {
		v := *m.CryptographicProviderType
		ret.CryptographicProviderType = &v
	}
	if m.CryptographicAlgorithmClass != nil {
		v := *m.CryptographicAlgorithmClass
		ret.CryptographicAlgorithmClass = &v
	}
	if m.CryptographicAlgorithmType != nil {
		v := *m.CryptographicAlgorithmType
		ret.CryptographicAlgorithmType = &v
	}
	if m.CryptographicAlgorithmSid != nil {
		v := *m.CryptographicAlgorithmSid
		ret.CryptographicAlgorithmSid = &v
	}
	if m.SpinCount != nil {
		v := *m.SpinCount
		ret.SpinCount = &v
	}
	if m.SaltData != nil {
		v := *m.SaltData
		ret.SaltData = &v
	}
	if m.HashData != nil {
		v := *m.HashData
		ret.HashData = &v
	}
	if m.CryptographicProvider != nil {
		v := *m.CryptographicProvider
		ret.CryptographicProvider = &v
	}
	if m.ExtendedCryptographicAlgorithm != nil {
		v := *m.ExtendedCryptographicAlgorithm
		ret.ExtendedCryptographicAlgorithm = &v
	}
	if m.ExtendedCryptographicAlgorithmSource != nil {
		v := *m.ExtendedCryptographicAlgorithmSource
		ret.ExtendedCryptographicAlgorithmSource = &v
	}
	if m.CryptographicProviderTypeExtensibility != nil {
		v := *m.CryptographicProviderTypeExtensibility
		ret.CryptographicProviderTypeExtensibility = &v
	}
	if m.CryptographicProviderTypeExtensibilitySource != nil {
		v := *m.CryptographicProviderTypeExtensibilitySource
		ret.CryptographicProviderTypeExtensibilitySource = &v
	}
	if m.AlgorithmName != nil {
		v := *m.AlgorithmName
		ret.AlgorithmName = &v
	}
	if m.HashValue != nil {
		v := *m.HashValue
		ret.HashValue = &v
	}
	if m.SaltValue != nil {
		v := *m.SaltValue
		ret.SaltValue = &v
	}
	if m.SpinValue != nil {
		v := *m.SpinValue
		ret.SpinValue = &v
	}
	return ret
}

func (m *ModificationVerifier) Validate() error {
	return nil
}

// Defines the PresentationExtensionList Class.
type PresentationExtensionList struct {
	*openxml.CompositeElementBase
	XMLName               xml.Name               `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	PresentationExtension *PresentationExtension `xml:"ext,omitempty"`
}

func NewPresentationExtensionList() *PresentationExtensionList {
	ret := &PresentationExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *PresentationExtensionList) Clone() openxml.Element {
	ret := NewPresentationExtensionList()
	if m.PresentationExtension != nil {
		ret.PresentationExtension = m.PresentationExtension.Clone().(*PresentationExtension)
	}
	return ret
}

func (m *PresentationExtensionList) Validate() error {
	if m.PresentationExtension != nil {
		if err := m.PresentationExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the PresentationExtension Class.
type PresentationExtension struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri               *types.StringValue `xml:"uri,attr,omitempty"`
	SectionProperties *SectionProperties `xml:"sectionPr,omitempty"`
	SectionList       *SectionList       `xml:"sectionLst,omitempty"`
	SlideGuideList    *SlideGuideList    `xml:"sldGuideLst,omitempty"`
	NotesGuideList    *NotesGuideList    `xml:"notesGuideLst,omitempty"`
}

func NewPresentationExtension() *PresentationExtension {
	ret := &PresentationExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *PresentationExtension) Clone() openxml.Element {
	ret := NewPresentationExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.SectionProperties != nil {
		ret.SectionProperties = m.SectionProperties.Clone().(*SectionProperties)
	}
	if m.SectionList != nil {
		ret.SectionList = m.SectionList.Clone().(*SectionList)
	}
	if m.SlideGuideList != nil {
		ret.SlideGuideList = m.SlideGuideList.Clone().(*SlideGuideList)
	}
	if m.NotesGuideList != nil {
		ret.NotesGuideList = m.NotesGuideList.Clone().(*NotesGuideList)
	}
	return ret
}

func (m *PresentationExtension) Validate() error {
	if m.SectionProperties != nil {
		if err := m.SectionProperties.Validate(); err != nil {
			return err
		}
	}
	if m.SectionList != nil {
		if err := m.SectionList.Validate(); err != nil {
			return err
		}
	}
	if m.SlideGuideList != nil {
		if err := m.SlideGuideList.Validate(); err != nil {
			return err
		}
	}
	if m.NotesGuideList != nil {
		if err := m.NotesGuideList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// HTML Publishing Properties.
type HtmlPublishProperties struct {
	*openxml.CompositeElementBase
	XMLName             xml.Name                                             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main htmlPubPr"`
	ShowSpeakerNotes    *types.BooleanValue                                  `xml:"showSpeakerNotes,attr,omitempty"`
	TargetBrowser       *types.EnumValue[HtmlPublishWebBrowserSupportValues] `xml:"pubBrowser,attr,omitempty"`
	Id                  *types.StringValue                                   `xml:"r:id,attr,omitempty"`
	SlideAll            *SlideAll                                            `xml:"sldAll,omitempty"`
	SlideRange          *SlideRange                                          `xml:"sldRg,omitempty"`
	CustomShowReference *CustomShowReference                                 `xml:"custShow,omitempty"`
	ExtensionList       *ExtensionList                                       `xml:"extLst,omitempty"`
}

func NewHtmlPublishProperties() *HtmlPublishProperties {
	ret := &HtmlPublishProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"htmlPubPr",
		"p",
	)
	return ret
}

func (m *HtmlPublishProperties) Clone() openxml.Element {
	ret := NewHtmlPublishProperties()
	if m.ShowSpeakerNotes != nil {
		v := *m.ShowSpeakerNotes
		ret.ShowSpeakerNotes = &v
	}
	if m.TargetBrowser != nil {
		v := *m.TargetBrowser
		ret.TargetBrowser = &v
	}
	if m.Id != nil {
		v := *m.Id
		ret.Id = &v
	}
	if m.SlideAll != nil {
		ret.SlideAll = m.SlideAll.Clone().(*SlideAll)
	}
	if m.SlideRange != nil {
		ret.SlideRange = m.SlideRange.Clone().(*SlideRange)
	}
	if m.CustomShowReference != nil {
		ret.CustomShowReference = m.CustomShowReference.Clone().(*CustomShowReference)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *HtmlPublishProperties) Validate() error {
	if m.SlideAll != nil {
		if err := m.SlideAll.Validate(); err != nil {
			return err
		}
	}
	if m.SlideRange != nil {
		if err := m.SlideRange.Validate(); err != nil {
			return err
		}
	}
	if m.CustomShowReference != nil {
		if err := m.CustomShowReference.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Web Properties.
type WebProperties struct {
	*openxml.CompositeElementBase
	XMLName           xml.Name                              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main webPr"`
	ShowAnimation     *types.BooleanValue                   `xml:"showAnimation,attr,omitempty"`
	ResizeGraphics    *types.BooleanValue                   `xml:"resizeGraphics,attr,omitempty"`
	AllowPng          *types.BooleanValue                   `xml:"allowPng,attr,omitempty"`
	RelyOnVml         *types.BooleanValue                   `xml:"relyOnVml,attr,omitempty"`
	OrganizeInFolders *types.BooleanValue                   `xml:"organizeInFolders,attr,omitempty"`
	UseLongFilenames  *types.BooleanValue                   `xml:"useLongFilenames,attr,omitempty"`
	ImageSize         *types.EnumValue[WebScreenSizeValues] `xml:"imgSz,attr,omitempty"`
	Encoding          *types.StringValue                    `xml:"encoding,attr,omitempty"`
	Color             *types.EnumValue[WebColorValues]      `xml:"clr,attr,omitempty"`
	ExtensionList     *ExtensionList                        `xml:"extLst,omitempty"`
}

func NewWebProperties() *WebProperties {
	ret := &WebProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"webPr",
		"p",
	)
	return ret
}

func (m *WebProperties) Clone() openxml.Element {
	ret := NewWebProperties()
	if m.ShowAnimation != nil {
		v := *m.ShowAnimation
		ret.ShowAnimation = &v
	}
	if m.ResizeGraphics != nil {
		v := *m.ResizeGraphics
		ret.ResizeGraphics = &v
	}
	if m.AllowPng != nil {
		v := *m.AllowPng
		ret.AllowPng = &v
	}
	if m.RelyOnVml != nil {
		v := *m.RelyOnVml
		ret.RelyOnVml = &v
	}
	if m.OrganizeInFolders != nil {
		v := *m.OrganizeInFolders
		ret.OrganizeInFolders = &v
	}
	if m.UseLongFilenames != nil {
		v := *m.UseLongFilenames
		ret.UseLongFilenames = &v
	}
	if m.ImageSize != nil {
		v := *m.ImageSize
		ret.ImageSize = &v
	}
	if m.Encoding != nil {
		v := *m.Encoding
		ret.Encoding = &v
	}
	if m.Color != nil {
		v := *m.Color
		ret.Color = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *WebProperties) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the PrintingProperties Class.
type PrintingProperties struct {
	*openxml.CompositeElementBase
	XMLName         xml.Name                               `xml:"http://schemas.openxmlformats.org/presentationml/2006/main prnPr"`
	PrintWhat       *types.EnumValue[PrintOutputValues]    `xml:"prnWhat,attr,omitempty"`
	ColorMode       *types.EnumValue[PrintColorModeValues] `xml:"clrMode,attr,omitempty"`
	HiddenSlides    *types.BooleanValue                    `xml:"hiddenSlides,attr,omitempty"`
	ScaleToFitPaper *types.BooleanValue                    `xml:"scaleToFitPaper,attr,omitempty"`
	FrameSlides     *types.BooleanValue                    `xml:"frameSlides,attr,omitempty"`
	ExtensionList   *ExtensionList                         `xml:"extLst,omitempty"`
}

func NewPrintingProperties() *PrintingProperties {
	ret := &PrintingProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"prnPr",
		"p",
	)
	return ret
}

func (m *PrintingProperties) Clone() openxml.Element {
	ret := NewPrintingProperties()
	if m.PrintWhat != nil {
		v := *m.PrintWhat
		ret.PrintWhat = &v
	}
	if m.ColorMode != nil {
		v := *m.ColorMode
		ret.ColorMode = &v
	}
	if m.HiddenSlides != nil {
		v := *m.HiddenSlides
		ret.HiddenSlides = &v
	}
	if m.ScaleToFitPaper != nil {
		v := *m.ScaleToFitPaper
		ret.ScaleToFitPaper = &v
	}
	if m.FrameSlides != nil {
		v := *m.FrameSlides
		ret.FrameSlides = &v
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *PrintingProperties) Validate() error {
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ShowProperties Class.
type ShowProperties struct {
	*openxml.CompositeElementBase
	XMLName                     xml.Name                     `xml:"http://schemas.openxmlformats.org/presentationml/2006/main showPr"`
	Loop                        *types.BooleanValue          `xml:"loop,attr,omitempty"`
	ShowNarration               *types.BooleanValue          `xml:"showNarration,attr,omitempty"`
	ShowAnimation               *types.BooleanValue          `xml:"showAnimation,attr,omitempty"`
	UseTimings                  *types.BooleanValue          `xml:"useTimings,attr,omitempty"`
	PresenterSlideMode          *PresenterSlideMode          `xml:"present,omitempty"`
	BrowseSlideMode             *BrowseSlideMode             `xml:"browse,omitempty"`
	KioskSlideMode              *KioskSlideMode              `xml:"kiosk,omitempty"`
	SlideAll                    *SlideAll                    `xml:"sldAll,omitempty"`
	SlideRange                  *SlideRange                  `xml:"sldRg,omitempty"`
	CustomShowReference         *CustomShowReference         `xml:"custShow,omitempty"`
	PenColor                    *PenColor                    `xml:"penClr,omitempty"`
	ShowPropertiesExtensionList *ShowPropertiesExtensionList `xml:"extLst,omitempty"`
}

func NewShowProperties() *ShowProperties {
	ret := &ShowProperties{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"showPr",
		"p",
	)
	return ret
}

func (m *ShowProperties) Clone() openxml.Element {
	ret := NewShowProperties()
	if m.Loop != nil {
		v := *m.Loop
		ret.Loop = &v
	}
	if m.ShowNarration != nil {
		v := *m.ShowNarration
		ret.ShowNarration = &v
	}
	if m.ShowAnimation != nil {
		v := *m.ShowAnimation
		ret.ShowAnimation = &v
	}
	if m.UseTimings != nil {
		v := *m.UseTimings
		ret.UseTimings = &v
	}
	if m.PresenterSlideMode != nil {
		ret.PresenterSlideMode = m.PresenterSlideMode.Clone().(*PresenterSlideMode)
	}
	if m.BrowseSlideMode != nil {
		ret.BrowseSlideMode = m.BrowseSlideMode.Clone().(*BrowseSlideMode)
	}
	if m.KioskSlideMode != nil {
		ret.KioskSlideMode = m.KioskSlideMode.Clone().(*KioskSlideMode)
	}
	if m.SlideAll != nil {
		ret.SlideAll = m.SlideAll.Clone().(*SlideAll)
	}
	if m.SlideRange != nil {
		ret.SlideRange = m.SlideRange.Clone().(*SlideRange)
	}
	if m.CustomShowReference != nil {
		ret.CustomShowReference = m.CustomShowReference.Clone().(*CustomShowReference)
	}
	if m.PenColor != nil {
		ret.PenColor = m.PenColor.Clone().(*PenColor)
	}
	if m.ShowPropertiesExtensionList != nil {
		ret.ShowPropertiesExtensionList = m.ShowPropertiesExtensionList.Clone().(*ShowPropertiesExtensionList)
	}
	return ret
}

func (m *ShowProperties) Validate() error {
	if m.PresenterSlideMode != nil {
		if err := m.PresenterSlideMode.Validate(); err != nil {
			return err
		}
	}
	if m.BrowseSlideMode != nil {
		if err := m.BrowseSlideMode.Validate(); err != nil {
			return err
		}
	}
	if m.KioskSlideMode != nil {
		if err := m.KioskSlideMode.Validate(); err != nil {
			return err
		}
	}
	if m.SlideAll != nil {
		if err := m.SlideAll.Validate(); err != nil {
			return err
		}
	}
	if m.SlideRange != nil {
		if err := m.SlideRange.Validate(); err != nil {
			return err
		}
	}
	if m.CustomShowReference != nil {
		if err := m.CustomShowReference.Validate(); err != nil {
			return err
		}
	}
	if m.PenColor != nil {
		if err := m.PenColor.Validate(); err != nil {
			return err
		}
	}
	if m.ShowPropertiesExtensionList != nil {
		if err := m.ShowPropertiesExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the ColorMostRecentlyUsed Class.
type ColorMostRecentlyUsed struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main clrMru"`
	// Skipped DrawingML type not yet implemented: RgbColorModelPercentage
	// Skipped DrawingML type not yet implemented: RgbColorModelHex
	HslColor    *drawingml.HslColor    `xml:"hslClr,omitempty"`
	SystemColor *drawingml.SystemColor `xml:"sysClr,omitempty"`
	SchemeColor *drawingml.SchemeColor `xml:"schemeClr,omitempty"`
	PresetColor *drawingml.PresetColor `xml:"prstClr,omitempty"`
}

func NewColorMostRecentlyUsed() *ColorMostRecentlyUsed {
	ret := &ColorMostRecentlyUsed{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"clrMru",
		"p",
	)
	return ret
}

func (m *ColorMostRecentlyUsed) Clone() openxml.Element {
	ret := NewColorMostRecentlyUsed()
	if m.HslColor != nil {
		ret.HslColor = m.HslColor.Clone().(*drawingml.HslColor)
	}
	if m.SystemColor != nil {
		ret.SystemColor = m.SystemColor.Clone().(*drawingml.SystemColor)
	}
	if m.SchemeColor != nil {
		ret.SchemeColor = m.SchemeColor.Clone().(*drawingml.SchemeColor)
	}
	if m.PresetColor != nil {
		ret.PresetColor = m.PresetColor.Clone().(*drawingml.PresetColor)
	}
	return ret
}

func (m *ColorMostRecentlyUsed) Validate() error {
	return nil
}

// Defines the PresentationPropertiesExtensionList Class.
type PresentationPropertiesExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                         xml.Name                         `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	PresentationPropertiesExtension *PresentationPropertiesExtension `xml:"ext,omitempty"`
}

func NewPresentationPropertiesExtensionList() *PresentationPropertiesExtensionList {
	ret := &PresentationPropertiesExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *PresentationPropertiesExtensionList) Clone() openxml.Element {
	ret := NewPresentationPropertiesExtensionList()
	if m.PresentationPropertiesExtension != nil {
		ret.PresentationPropertiesExtension = m.PresentationPropertiesExtension.Clone().(*PresentationPropertiesExtension)
	}
	return ret
}

func (m *PresentationPropertiesExtensionList) Validate() error {
	if m.PresentationPropertiesExtension != nil {
		if err := m.PresentationPropertiesExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the PresentationPropertiesExtension Class.
type PresentationPropertiesExtension struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	Uri                  *types.StringValue    `xml:"uri,attr,omitempty"`
	DiscardImageEditData *DiscardImageEditData `xml:"discardImageEditData,omitempty"`
	DefaultImageDpi      *DefaultImageDpi      `xml:"defaultImageDpi,omitempty"`
	// Skipped DrawingML type not yet implemented: TextMath
	ChartTrackingReferenceBased *ChartTrackingReferenceBased `xml:"chartTrackingRefBased,omitempty"`
}

func NewPresentationPropertiesExtension() *PresentationPropertiesExtension {
	ret := &PresentationPropertiesExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *PresentationPropertiesExtension) Clone() openxml.Element {
	ret := NewPresentationPropertiesExtension()
	if m.Uri != nil {
		v := *m.Uri
		ret.Uri = &v
	}
	if m.DiscardImageEditData != nil {
		ret.DiscardImageEditData = m.DiscardImageEditData.Clone().(*DiscardImageEditData)
	}
	if m.DefaultImageDpi != nil {
		ret.DefaultImageDpi = m.DefaultImageDpi.Clone().(*DefaultImageDpi)
	}
	if m.ChartTrackingReferenceBased != nil {
		ret.ChartTrackingReferenceBased = m.ChartTrackingReferenceBased.Clone().(*ChartTrackingReferenceBased)
	}
	return ret
}

func (m *PresentationPropertiesExtension) Validate() error {
	if m.DiscardImageEditData != nil {
		if err := m.DiscardImageEditData.Validate(); err != nil {
			return err
		}
	}
	if m.DefaultImageDpi != nil {
		if err := m.DefaultImageDpi.Validate(); err != nil {
			return err
		}
	}
	if m.ChartTrackingReferenceBased != nil {
		if err := m.ChartTrackingReferenceBased.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the HeaderFooter Class.
type HeaderFooter struct {
	*openxml.CompositeElementBase
	XMLName                       xml.Name                       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main hf"`
	SlideNumber                   *types.BooleanValue            `xml:"sldNum,attr,omitempty"`
	Header                        *types.BooleanValue            `xml:"hdr,attr,omitempty"`
	Footer                        *types.BooleanValue            `xml:"ftr,attr,omitempty"`
	DateTime                      *types.BooleanValue            `xml:"dt,attr,omitempty"`
	ExtensionListWithModification *ExtensionListWithModification `xml:"extLst,omitempty"`
}

func NewHeaderFooter() *HeaderFooter {
	ret := &HeaderFooter{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"hf",
		"p",
	)
	return ret
}

func (m *HeaderFooter) Clone() openxml.Element {
	ret := NewHeaderFooter()
	if m.SlideNumber != nil {
		v := *m.SlideNumber
		ret.SlideNumber = &v
	}
	if m.Header != nil {
		v := *m.Header
		ret.Header = &v
	}
	if m.Footer != nil {
		v := *m.Footer
		ret.Footer = &v
	}
	if m.DateTime != nil {
		v := *m.DateTime
		ret.DateTime = &v
	}
	if m.ExtensionListWithModification != nil {
		ret.ExtensionListWithModification = m.ExtensionListWithModification.Clone().(*ExtensionListWithModification)
	}
	return ret
}

func (m *HeaderFooter) Validate() error {
	if m.ExtensionListWithModification != nil {
		if err := m.ExtensionListWithModification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideLayoutExtensionList Class.
type SlideLayoutExtensionList struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	SlideLayoutExtension *SlideLayoutExtension `xml:"ext,omitempty"`
}

func NewSlideLayoutExtensionList() *SlideLayoutExtensionList {
	ret := &SlideLayoutExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *SlideLayoutExtensionList) Clone() openxml.Element {
	ret := NewSlideLayoutExtensionList()
	if m.SlideLayoutExtension != nil {
		ret.SlideLayoutExtension = m.SlideLayoutExtension.Clone().(*SlideLayoutExtension)
	}
	return ret
}

func (m *SlideLayoutExtensionList) Validate() error {
	if m.SlideLayoutExtension != nil {
		if err := m.SlideLayoutExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideLayoutIdList Class.
type SlideLayoutIdList struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldLayoutIdLst"`
	SlideLayoutId *SlideLayoutId `xml:"sldLayoutId,omitempty"`
}

func NewSlideLayoutIdList() *SlideLayoutIdList {
	ret := &SlideLayoutIdList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sldLayoutIdLst",
		"p",
	)
	return ret
}

func (m *SlideLayoutIdList) Clone() openxml.Element {
	ret := NewSlideLayoutIdList()
	if m.SlideLayoutId != nil {
		ret.SlideLayoutId = m.SlideLayoutId.Clone().(*SlideLayoutId)
	}
	return ret
}

func (m *SlideLayoutIdList) Validate() error {
	if m.SlideLayoutId != nil {
		if err := m.SlideLayoutId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the TextStyles Class.
type TextStyles struct {
	*openxml.CompositeElementBase
	XMLName       xml.Name       `xml:"http://schemas.openxmlformats.org/presentationml/2006/main txStyles"`
	TitleStyle    *TitleStyle    `xml:"titleStyle,omitempty"`
	BodyStyle     *BodyStyle     `xml:"bodyStyle,omitempty"`
	OtherStyle    *OtherStyle    `xml:"otherStyle,omitempty"`
	ExtensionList *ExtensionList `xml:"extLst,omitempty"`
}

func NewTextStyles() *TextStyles {
	ret := &TextStyles{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"txStyles",
		"p",
	)
	return ret
}

func (m *TextStyles) Clone() openxml.Element {
	ret := NewTextStyles()
	if m.TitleStyle != nil {
		ret.TitleStyle = m.TitleStyle.Clone().(*TitleStyle)
	}
	if m.BodyStyle != nil {
		ret.BodyStyle = m.BodyStyle.Clone().(*BodyStyle)
	}
	if m.OtherStyle != nil {
		ret.OtherStyle = m.OtherStyle.Clone().(*OtherStyle)
	}
	if m.ExtensionList != nil {
		ret.ExtensionList = m.ExtensionList.Clone().(*ExtensionList)
	}
	return ret
}

func (m *TextStyles) Validate() error {
	if m.TitleStyle != nil {
		if err := m.TitleStyle.Validate(); err != nil {
			return err
		}
	}
	if m.BodyStyle != nil {
		if err := m.BodyStyle.Validate(); err != nil {
			return err
		}
	}
	if m.OtherStyle != nil {
		if err := m.OtherStyle.Validate(); err != nil {
			return err
		}
	}
	if m.ExtensionList != nil {
		if err := m.ExtensionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the SlideMasterExtensionList Class.
type SlideMasterExtensionList struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	SlideMasterExtension *SlideMasterExtension `xml:"ext,omitempty"`
}

func NewSlideMasterExtensionList() *SlideMasterExtensionList {
	ret := &SlideMasterExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *SlideMasterExtensionList) Clone() openxml.Element {
	ret := NewSlideMasterExtensionList()
	if m.SlideMasterExtension != nil {
		ret.SlideMasterExtension = m.SlideMasterExtension.Clone().(*SlideMasterExtension)
	}
	return ret
}

func (m *SlideMasterExtensionList) Validate() error {
	if m.SlideMasterExtension != nil {
		if err := m.SlideMasterExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the HandoutMasterExtensionList Class.
type HandoutMasterExtensionList struct {
	*openxml.CompositeElementBase
	XMLName                xml.Name                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	HandoutMasterExtension *HandoutMasterExtension `xml:"ext,omitempty"`
}

func NewHandoutMasterExtensionList() *HandoutMasterExtensionList {
	ret := &HandoutMasterExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *HandoutMasterExtensionList) Clone() openxml.Element {
	ret := NewHandoutMasterExtensionList()
	if m.HandoutMasterExtension != nil {
		ret.HandoutMasterExtension = m.HandoutMasterExtension.Clone().(*HandoutMasterExtension)
	}
	return ret
}

func (m *HandoutMasterExtensionList) Validate() error {
	if m.HandoutMasterExtension != nil {
		if err := m.HandoutMasterExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the NotesMasterExtensionList Class.
type NotesMasterExtensionList struct {
	*openxml.CompositeElementBase
	XMLName              xml.Name              `xml:"http://schemas.openxmlformats.org/presentationml/2006/main extLst"`
	NotesMasterExtension *NotesMasterExtension `xml:"ext,omitempty"`
}

func NewNotesMasterExtensionList() *NotesMasterExtensionList {
	ret := &NotesMasterExtensionList{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"extLst",
		"p",
	)
	return ret
}

func (m *NotesMasterExtensionList) Clone() openxml.Element {
	ret := NewNotesMasterExtensionList()
	if m.NotesMasterExtension != nil {
		ret.NotesMasterExtension = m.NotesMasterExtension.Clone().(*NotesMasterExtension)
	}
	return ret
}

func (m *NotesMasterExtensionList) Validate() error {
	if m.NotesMasterExtension != nil {
		if err := m.NotesMasterExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// OLE Chart Element.
type OleChartElement struct {
	*openxml.LeafElementBase
	XMLName xml.Name                                `xml:"http://schemas.openxmlformats.org/presentationml/2006/main oleChartEl"`
	Type    *types.EnumValue[ChartSubElementValues] `xml:"type,attr,omitempty"`
	Level   *types.UInt32Value                      `xml:"lvl,attr,omitempty"`
}

func NewOleChartElement() *OleChartElement {
	ret := &OleChartElement{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"oleChartEl",
		"p",
	)
	return ret
}

func (m *OleChartElement) Clone() openxml.Element {
	ret := NewOleChartElement()
	if m.Type != nil {
		v := *m.Type
		ret.Type = &v
	}
	if m.Level != nil {
		v := *m.Level
		ret.Level = &v
	}
	return ret
}

func (m *OleChartElement) Validate() error {
	return nil
}

// Text Element.
type TextElement struct {
	*openxml.CompositeElementBase
	XMLName             xml.Name             `xml:"http://schemas.openxmlformats.org/presentationml/2006/main txEl"`
	CharRange           *CharRange           `xml:"charRg,omitempty"`
	ParagraphIndexRange *ParagraphIndexRange `xml:"pRg,omitempty"`
}

func NewTextElement() *TextElement {
	ret := &TextElement{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"txEl",
		"p",
	)
	return ret
}

func (m *TextElement) Clone() openxml.Element {
	ret := NewTextElement()
	if m.CharRange != nil {
		ret.CharRange = m.CharRange.Clone().(*CharRange)
	}
	if m.ParagraphIndexRange != nil {
		ret.ParagraphIndexRange = m.ParagraphIndexRange.Clone().(*ParagraphIndexRange)
	}
	return ret
}

func (m *TextElement) Validate() error {
	if m.CharRange != nil {
		if err := m.CharRange.Validate(); err != nil {
			return err
		}
	}
	if m.ParagraphIndexRange != nil {
		if err := m.ParagraphIndexRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Graphic Element.
type GraphicElement struct {
	*openxml.CompositeElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main graphicEl"`
	// Skipped DrawingML type not yet implemented: Diagram
	Chart *drawingml.Chart `xml:"chart,omitempty"`
}

func NewGraphicElement() *GraphicElement {
	ret := &GraphicElement{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"graphicEl",
		"p",
	)
	return ret
}

func (m *GraphicElement) Clone() openxml.Element {
	ret := NewGraphicElement()
	if m.Chart != nil {
		ret.Chart = m.Chart.Clone().(*drawingml.Chart)
	}
	return ret
}

func (m *GraphicElement) Validate() error {
	return nil
}

// Defines the BlindsTransition Class.
type BlindsTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main blinds"`
}

func NewBlindsTransition() *BlindsTransition {
	ret := &BlindsTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"blinds",
		"p",
	)
	return ret
}

func (m *BlindsTransition) Clone() openxml.Element {
	ret := NewBlindsTransition()
	return ret
}

func (m *BlindsTransition) Validate() error {
	return nil
}

// Defines the CheckerTransition Class.
type CheckerTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main checker"`
}

func NewCheckerTransition() *CheckerTransition {
	ret := &CheckerTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"checker",
		"p",
	)
	return ret
}

func (m *CheckerTransition) Clone() openxml.Element {
	ret := NewCheckerTransition()
	return ret
}

func (m *CheckerTransition) Validate() error {
	return nil
}

// Defines the CombTransition Class.
type CombTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main comb"`
}

func NewCombTransition() *CombTransition {
	ret := &CombTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"comb",
		"p",
	)
	return ret
}

func (m *CombTransition) Clone() openxml.Element {
	ret := NewCombTransition()
	return ret
}

func (m *CombTransition) Validate() error {
	return nil
}

// Defines the RandomBarTransition Class.
type RandomBarTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main randomBar"`
}

func NewRandomBarTransition() *RandomBarTransition {
	ret := &RandomBarTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"randomBar",
		"p",
	)
	return ret
}

func (m *RandomBarTransition) Clone() openxml.Element {
	ret := NewRandomBarTransition()
	return ret
}

func (m *RandomBarTransition) Validate() error {
	return nil
}

// Defines the CoverTransition Class.
type CoverTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cover"`
}

func NewCoverTransition() *CoverTransition {
	ret := &CoverTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"cover",
		"p",
	)
	return ret
}

func (m *CoverTransition) Clone() openxml.Element {
	ret := NewCoverTransition()
	return ret
}

func (m *CoverTransition) Validate() error {
	return nil
}

// Defines the PullTransition Class.
type PullTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main pull"`
}

func NewPullTransition() *PullTransition {
	ret := &PullTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"pull",
		"p",
	)
	return ret
}

func (m *PullTransition) Clone() openxml.Element {
	ret := NewPullTransition()
	return ret
}

func (m *PullTransition) Validate() error {
	return nil
}

// Defines the EightDirectionTransitionType Class.
type EightDirectionTransitionType struct {
	*openxml.LeafElementBase
	XMLName   xml.Name           `xml:""`
	Direction *types.StringValue `xml:"dir,attr,omitempty"`
}

func NewEightDirectionTransitionType() *EightDirectionTransitionType {
	ret := &EightDirectionTransitionType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *EightDirectionTransitionType) Clone() openxml.Element {
	ret := NewEightDirectionTransitionType()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *EightDirectionTransitionType) Validate() error {
	return nil
}

// Defines the CutTransition Class.
type CutTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cut"`
}

func NewCutTransition() *CutTransition {
	ret := &CutTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"cut",
		"p",
	)
	return ret
}

func (m *CutTransition) Clone() openxml.Element {
	ret := NewCutTransition()
	return ret
}

func (m *CutTransition) Validate() error {
	return nil
}

// Defines the FadeTransition Class.
type FadeTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main fade"`
}

func NewFadeTransition() *FadeTransition {
	ret := &FadeTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"fade",
		"p",
	)
	return ret
}

func (m *FadeTransition) Clone() openxml.Element {
	ret := NewFadeTransition()
	return ret
}

func (m *FadeTransition) Validate() error {
	return nil
}

// Defines the OptionalBlackTransitionType Class.
type OptionalBlackTransitionType struct {
	*openxml.LeafElementBase
	XMLName      xml.Name            `xml:""`
	ThroughBlack *types.BooleanValue `xml:"thruBlk,attr,omitempty"`
}

func NewOptionalBlackTransitionType() *OptionalBlackTransitionType {
	ret := &OptionalBlackTransitionType{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"",
		"p",
	)
	return ret
}

func (m *OptionalBlackTransitionType) Clone() openxml.Element {
	ret := NewOptionalBlackTransitionType()
	if m.ThroughBlack != nil {
		v := *m.ThroughBlack
		ret.ThroughBlack = &v
	}
	return ret
}

func (m *OptionalBlackTransitionType) Validate() error {
	return nil
}

// Defines the PushTransition Class.
type PushTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main push"`
}

func NewPushTransition() *PushTransition {
	ret := &PushTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"push",
		"p",
	)
	return ret
}

func (m *PushTransition) Clone() openxml.Element {
	ret := NewPushTransition()
	return ret
}

func (m *PushTransition) Validate() error {
	return nil
}

// Defines the WipeTransition Class.
type WipeTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main wipe"`
}

func NewWipeTransition() *WipeTransition {
	ret := &WipeTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"wipe",
		"p",
	)
	return ret
}

func (m *WipeTransition) Clone() openxml.Element {
	ret := NewWipeTransition()
	return ret
}

func (m *WipeTransition) Validate() error {
	return nil
}

// Defines the SplitTransition Class.
type SplitTransition struct {
	*openxml.LeafElementBase
	XMLName     xml.Name                                         `xml:"http://schemas.openxmlformats.org/presentationml/2006/main split"`
	Orientation *types.EnumValue[DirectionValues]                `xml:"orient,attr,omitempty"`
	Direction   *types.EnumValue[TransitionInOutDirectionValues] `xml:"dir,attr,omitempty"`
}

func NewSplitTransition() *SplitTransition {
	ret := &SplitTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"split",
		"p",
	)
	return ret
}

func (m *SplitTransition) Clone() openxml.Element {
	ret := NewSplitTransition()
	if m.Orientation != nil {
		v := *m.Orientation
		ret.Orientation = &v
	}
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *SplitTransition) Validate() error {
	return nil
}

// Defines the StripsTransition Class.
type StripsTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main strips"`
	Direction *types.EnumValue[TransitionCornerDirectionValues] `xml:"dir,attr,omitempty"`
}

func NewStripsTransition() *StripsTransition {
	ret := &StripsTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"strips",
		"p",
	)
	return ret
}

func (m *StripsTransition) Clone() openxml.Element {
	ret := NewStripsTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *StripsTransition) Validate() error {
	return nil
}

// Defines the WheelTransition Class.
type WheelTransition struct {
	*openxml.LeafElementBase
	XMLName xml.Name           `xml:"http://schemas.openxmlformats.org/presentationml/2006/main wheel"`
	Spokes  *types.UInt32Value `xml:"spokes,attr,omitempty"`
}

func NewWheelTransition() *WheelTransition {
	ret := &WheelTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"wheel",
		"p",
	)
	return ret
}

func (m *WheelTransition) Clone() openxml.Element {
	ret := NewWheelTransition()
	if m.Spokes != nil {
		v := *m.Spokes
		ret.Spokes = &v
	}
	return ret
}

func (m *WheelTransition) Validate() error {
	return nil
}

// Defines the ZoomTransition Class.
type ZoomTransition struct {
	*openxml.LeafElementBase
	XMLName   xml.Name                                         `xml:"http://schemas.openxmlformats.org/presentationml/2006/main zoom"`
	Direction *types.EnumValue[TransitionInOutDirectionValues] `xml:"dir,attr,omitempty"`
}

func NewZoomTransition() *ZoomTransition {
	ret := &ZoomTransition{}
	ns := openxml.NamespacePresentationML
	ret.LeafElementBase = openxml.NewLeafElement(
		ns,
		"zoom",
		"p",
	)
	return ret
}

func (m *ZoomTransition) Clone() openxml.Element {
	ret := NewZoomTransition()
	if m.Direction != nil {
		v := *m.Direction
		ret.Direction = &v
	}
	return ret
}

func (m *ZoomTransition) Validate() error {
	return nil
}

// Defines the SoundAction Class.
type SoundAction struct {
	*openxml.CompositeElementBase
	XMLName          xml.Name          `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sndAc"`
	StartSoundAction *StartSoundAction `xml:"stSnd,omitempty"`
	EndSoundAction   *EndSoundAction   `xml:"endSnd,omitempty"`
}

func NewSoundAction() *SoundAction {
	ret := &SoundAction{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"sndAc",
		"p",
	)
	return ret
}

func (m *SoundAction) Clone() openxml.Element {
	ret := NewSoundAction()
	if m.StartSoundAction != nil {
		ret.StartSoundAction = m.StartSoundAction.Clone().(*StartSoundAction)
	}
	if m.EndSoundAction != nil {
		ret.EndSoundAction = m.EndSoundAction.Clone().(*EndSoundAction)
	}
	return ret
}

func (m *SoundAction) Validate() error {
	if m.StartSoundAction != nil {
		if err := m.StartSoundAction.Validate(); err != nil {
			return err
		}
	}
	if m.EndSoundAction != nil {
		if err := m.EndSoundAction.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Defines the PlaceholderExtension Class.
type PlaceholderExtension struct {
	*openxml.CompositeElementBase
	XMLName                  xml.Name                  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ext"`
	PlaceholderTypeExtension *PlaceholderTypeExtension `xml:"phTypeExt,omitempty"`
}

func NewPlaceholderExtension() *PlaceholderExtension {
	ret := &PlaceholderExtension{}
	ns := openxml.NamespacePresentationML
	ret.CompositeElementBase = openxml.NewCompositeElement(
		ns,
		"ext",
		"p",
	)
	return ret
}

func (m *PlaceholderExtension) Clone() openxml.Element {
	ret := NewPlaceholderExtension()
	if m.PlaceholderTypeExtension != nil {
		ret.PlaceholderTypeExtension = m.PlaceholderTypeExtension.Clone().(*PlaceholderTypeExtension)
	}
	return ret
}

func (m *PlaceholderExtension) Validate() error {
	if m.PlaceholderTypeExtension != nil {
		if err := m.PlaceholderTypeExtension.Validate(); err != nil {
			return err
		}
	}
	return nil
}
