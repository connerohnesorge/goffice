//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings, also theme XML is verbose
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// Counter for generating unique theme filenames.
var themeUniqueCounter uint64

// ThemePart represents a theme part (ppt/theme/theme1.xml).
type ThemePart struct {
	*openxml.OpenXmlPartData
}

// newThemePart creates a new theme part for a presentation.
func newThemePart(
	presentationPart *PresentationPart,
) (*ThemePart, error) {
	num := atomic.AddUint64(
		&themeUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/theme/theme%d.xml",
		num,
	)

	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeTheme,
		RelationshipTypeTheme,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTheme,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	tp := &ThemePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal theme content
	tp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(tp, relID); err != nil {
		return nil, err
	}

	return tp, nil
}

// newThemePartForMaster creates a new theme part for a slide master.
func newThemePartForMaster(
	slideMasterPart *SlideMasterPart,
) (*ThemePart, error) {
	num := atomic.AddUint64(
		&themeUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/theme/theme%d.xml",
		num,
	)

	packPart, relID, err := slideMasterPart.addChildPart(
		uri,
		ContentTypeTheme,
		RelationshipTypeTheme,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTheme,
		packPart,
		slideMasterPart,
	)
	partData.SetRelationshipID(relID)

	tp := &ThemePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal theme content
	tp.initializeContent()

	// Add to slide master part's child parts
	if err := slideMasterPart.AddPart(tp, relID); err != nil {
		return nil, err
	}

	return tp, nil
}

// newThemePartForNotesMaster creates a new theme part for a notes master.
func newThemePartForNotesMaster(
	notesMasterPart *NotesMasterPart,
) (*ThemePart, error) {
	num := atomic.AddUint64(
		&themeUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/theme/theme%d.xml",
		num,
	)

	packPart, relID, err := notesMasterPart.addChildPart(
		uri,
		ContentTypeTheme,
		RelationshipTypeTheme,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTheme,
		packPart,
		notesMasterPart,
	)
	partData.SetRelationshipID(relID)

	tp := &ThemePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal theme content
	tp.initializeContent()

	// Add to notes master part's child parts
	if err := notesMasterPart.AddPart(tp, relID); err != nil {
		return nil, err
	}

	return tp, nil
}

// newThemePartForHandoutMaster creates a new theme part for a handout master.
func newThemePartForHandoutMaster(
	handoutMasterPart *HandoutMasterPart,
) (*ThemePart, error) {
	num := atomic.AddUint64(
		&themeUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/theme/theme%d.xml",
		num,
	)

	packPart, relID, err := handoutMasterPart.addChildPart(
		uri,
		ContentTypeTheme,
		RelationshipTypeTheme,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTheme,
		packPart,
		handoutMasterPart,
	)
	partData.SetRelationshipID(relID)

	tp := &ThemePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal theme content
	tp.initializeContent()

	// Add to handout master part's child parts
	if err := handoutMasterPart.AddPart(tp, relID); err != nil {
		return nil, err
	}

	return tp, nil
}

// initializeContent sets up minimal theme content (Office theme).
func (tp *ThemePart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<a:theme xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" name="Office Theme">
  <a:themeElements>
    <a:clrScheme name="Office">
      <a:dk1><a:sysClr val="windowText" lastClr="000000"/></a:dk1>
      <a:lt1><a:sysClr val="window" lastClr="FFFFFF"/></a:lt1>
      <a:dk2><a:srgbClr val="44546A"/></a:dk2>
      <a:lt2><a:srgbClr val="E7E6E6"/></a:lt2>
      <a:accent1><a:srgbClr val="4472C4"/></a:accent1>
      <a:accent2><a:srgbClr val="ED7D31"/></a:accent2>
      <a:accent3><a:srgbClr val="A5A5A5"/></a:accent3>
      <a:accent4><a:srgbClr val="FFC000"/></a:accent4>
      <a:accent5><a:srgbClr val="5B9BD5"/></a:accent5>
      <a:accent6><a:srgbClr val="70AD47"/></a:accent6>
      <a:hlink><a:srgbClr val="0563C1"/></a:hlink>
      <a:folHlink><a:srgbClr val="954F72"/></a:folHlink>
    </a:clrScheme>
    <a:fontScheme name="Office">
      <a:majorFont>
        <a:latin typeface="Calibri Light" panose="020F0302020204030204"/>
        <a:ea typeface=""/>
        <a:cs typeface=""/>
      </a:majorFont>
      <a:minorFont>
        <a:latin typeface="Calibri" panose="020F0502020204030204"/>
        <a:ea typeface=""/>
        <a:cs typeface=""/>
      </a:minorFont>
    </a:fontScheme>
    <a:fmtScheme name="Office">
      <a:fillStyleLst>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
        <a:gradFill rotWithShape="1">
          <a:gsLst>
            <a:gs pos="0"><a:schemeClr val="phClr"><a:lumMod val="110000"/><a:satMod val="105000"/><a:tint val="67000"/></a:schemeClr></a:gs>
            <a:gs pos="50000"><a:schemeClr val="phClr"><a:lumMod val="105000"/><a:satMod val="103000"/><a:tint val="73000"/></a:schemeClr></a:gs>
            <a:gs pos="100000"><a:schemeClr val="phClr"><a:lumMod val="105000"/><a:satMod val="109000"/><a:tint val="81000"/></a:schemeClr></a:gs>
          </a:gsLst>
          <a:lin ang="5400000" scaled="0"/>
        </a:gradFill>
        <a:gradFill rotWithShape="1">
          <a:gsLst>
            <a:gs pos="0"><a:schemeClr val="phClr"><a:satMod val="103000"/><a:lumMod val="102000"/><a:tint val="94000"/></a:schemeClr></a:gs>
            <a:gs pos="50000"><a:schemeClr val="phClr"><a:satMod val="110000"/><a:lumMod val="100000"/><a:shade val="100000"/></a:schemeClr></a:gs>
            <a:gs pos="100000"><a:schemeClr val="phClr"><a:lumMod val="99000"/><a:satMod val="120000"/><a:shade val="78000"/></a:schemeClr></a:gs>
          </a:gsLst>
          <a:lin ang="5400000" scaled="0"/>
        </a:gradFill>
      </a:fillStyleLst>
      <a:lnStyleLst>
        <a:ln w="6350" cap="flat" cmpd="sng" algn="ctr"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:prstDash val="solid"/><a:miter lim="800000"/></a:ln>
        <a:ln w="12700" cap="flat" cmpd="sng" algn="ctr"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:prstDash val="solid"/><a:miter lim="800000"/></a:ln>
        <a:ln w="19050" cap="flat" cmpd="sng" algn="ctr"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:prstDash val="solid"/><a:miter lim="800000"/></a:ln>
      </a:lnStyleLst>
      <a:effectStyleLst>
        <a:effectStyle><a:effectLst/></a:effectStyle>
        <a:effectStyle><a:effectLst/></a:effectStyle>
        <a:effectStyle><a:effectLst><a:outerShdw blurRad="57150" dist="19050" dir="5400000" algn="ctr" rotWithShape="0"><a:srgbClr val="000000"><a:alpha val="63000"/></a:srgbClr></a:outerShdw></a:effectLst></a:effectStyle>
      </a:effectStyleLst>
      <a:bgFillStyleLst>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
        <a:solidFill><a:schemeClr val="phClr"><a:tint val="95000"/><a:satMod val="170000"/></a:schemeClr></a:solidFill>
        <a:gradFill rotWithShape="1">
          <a:gsLst>
            <a:gs pos="0"><a:schemeClr val="phClr"><a:tint val="93000"/><a:satMod val="150000"/><a:shade val="98000"/><a:lumMod val="102000"/></a:schemeClr></a:gs>
            <a:gs pos="50000"><a:schemeClr val="phClr"><a:tint val="98000"/><a:satMod val="130000"/><a:shade val="90000"/><a:lumMod val="103000"/></a:schemeClr></a:gs>
            <a:gs pos="100000"><a:schemeClr val="phClr"><a:shade val="63000"/><a:satMod val="120000"/></a:schemeClr></a:gs>
          </a:gsLst>
          <a:lin ang="5400000" scaled="0"/>
        </a:gradFill>
      </a:bgFillStyleLst>
    </a:fmtScheme>
  </a:themeElements>
  <a:objectDefaults/>
  <a:extraClrSchemeLst/>
</a:theme>`
	tp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*ThemePart) FixedContentType() string {
	return ContentTypeTheme
}

// Theme returns the root Theme element.
func (tp *ThemePart) Theme() *drawingml.Theme {
	root := tp.RootElement()
	if root == nil {
		return nil
	}
	if theme, ok := root.(*drawingml.Theme); ok {
		return theme
	}

	return nil
}

// GetStream returns a reader for the part content.
func (tp *ThemePart) GetStream() io.Reader {
	return tp.OpenXmlPartData.GetStream()
}

// Ensure ThemePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ThemePart)(nil)

// ThemePartFactory creates a ThemePart from a URI and container.
func ThemePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTheme,
		packPart,
		container,
	)

	return &ThemePart{
		OpenXmlPartData: partData,
	}
}

// Register the ThemePart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeTheme,
			RelationshipType:   RelationshipTypeTheme,
			Factory:            ThemePartFactory,
			DefaultURI:         "/ppt/theme/theme1.xml",
			IsFixedContentType: true,
		},
	)
}
