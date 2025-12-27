//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
)

// HandoutMasterPart represents a handout master part (ppt/handoutMasters/handoutMaster1.xml).
type HandoutMasterPart struct {
	*openxml.OpenXmlPartData
}

// newHandoutMasterPart creates a new handout master part.
func newHandoutMasterPart(
	presentationPart *PresentationPart,
	uri string,
) (*HandoutMasterPart, error) {
	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeHandoutMaster,
		RelationshipTypeHandoutMaster,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeHandoutMaster,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	hmp := &HandoutMasterPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal handout master content
	hmp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(hmp, relID); err != nil {
		return nil, err
	}

	return hmp, nil
}

// initializeContent sets up minimal handout master content.
func (hmp *HandoutMasterPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:handoutMaster xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:bg>
      <p:bgRef idx="1001">
        <a:schemeClr val="bg1"/>
      </p:bgRef>
    </p:bg>
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr/>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="2" name="Header Placeholder 1"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="hdr" sz="quarter"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="0" y="0"/>
            <a:ext cx="2971800" cy="458788"/>
          </a:xfrm>
          <a:prstGeom prst="rect">
            <a:avLst/>
          </a:prstGeom>
        </p:spPr>
        <p:txBody>
          <a:bodyPr vert="horz" lIns="91440" tIns="45720" rIns="91440" bIns="45720" rtlCol="0"/>
          <a:lstStyle>
            <a:lvl1pPr algn="l">
              <a:defRPr sz="1200"/>
            </a:lvl1pPr>
          </a:lstStyle>
          <a:p>
            <a:endParaRPr lang="en-US"/>
          </a:p>
        </p:txBody>
      </p:sp>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="3" name="Date Placeholder 2"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="dt" sz="quarter" idx="1"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="3884613" y="0"/>
            <a:ext cx="2971800" cy="458788"/>
          </a:xfrm>
          <a:prstGeom prst="rect">
            <a:avLst/>
          </a:prstGeom>
        </p:spPr>
        <p:txBody>
          <a:bodyPr vert="horz" lIns="91440" tIns="45720" rIns="91440" bIns="45720" rtlCol="0"/>
          <a:lstStyle>
            <a:lvl1pPr algn="r">
              <a:defRPr sz="1200"/>
            </a:lvl1pPr>
          </a:lstStyle>
          <a:p>
            <a:fld id="{5282F153-3F37-4BDA-A5BB-2F8B22000000}" type="datetimeFigureOut">
              <a:rPr lang="en-US"/>
              <a:t>1/1/2024</a:t>
            </a:fld>
            <a:endParaRPr lang="en-US"/>
          </a:p>
        </p:txBody>
      </p:sp>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="4" name="Footer Placeholder 3"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="ftr" sz="quarter" idx="2"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="0" y="8685213"/>
            <a:ext cx="2971800" cy="458787"/>
          </a:xfrm>
          <a:prstGeom prst="rect">
            <a:avLst/>
          </a:prstGeom>
        </p:spPr>
        <p:txBody>
          <a:bodyPr vert="horz" lIns="91440" tIns="45720" rIns="91440" bIns="45720" rtlCol="0" anchor="b"/>
          <a:lstStyle>
            <a:lvl1pPr algn="l">
              <a:defRPr sz="1200"/>
            </a:lvl1pPr>
          </a:lstStyle>
          <a:p>
            <a:endParaRPr lang="en-US"/>
          </a:p>
        </p:txBody>
      </p:sp>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="5" name="Slide Number Placeholder 4"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="sldNum" sz="quarter" idx="3"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="3884613" y="8685213"/>
            <a:ext cx="2971800" cy="458787"/>
          </a:xfrm>
          <a:prstGeom prst="rect">
            <a:avLst/>
          </a:prstGeom>
        </p:spPr>
        <p:txBody>
          <a:bodyPr vert="horz" lIns="91440" tIns="45720" rIns="91440" bIns="45720" rtlCol="0" anchor="b"/>
          <a:lstStyle>
            <a:lvl1pPr algn="r">
              <a:defRPr sz="1200"/>
            </a:lvl1pPr>
          </a:lstStyle>
          <a:p>
            <a:fld id="{5282F153-3F37-4BDA-A5BB-2F8B22000001}" type="slidenum">
              <a:rPr lang="en-US"/>
              <a:t>1</a:t>
            </a:fld>
            <a:endParaRPr lang="en-US"/>
          </a:p>
        </p:txBody>
      </p:sp>
    </p:spTree>
  </p:cSld>
  <p:clrMap bg1="lt1" tx1="dk1" bg2="lt2" tx2="dk2" accent1="accent1" accent2="accent2" accent3="accent3" accent4="accent4" accent5="accent5" accent6="accent6" hlink="hlink" folHlink="folHlink"/>
</p:handoutMaster>`
	hmp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*HandoutMasterPart) FixedContentType() string {
	return ContentTypeHandoutMaster
}

// HandoutMaster returns the root HandoutMaster element.
// TODO: Return a proper HandoutMaster element type when elements are implemented.
func (hmp *HandoutMasterPart) HandoutMaster() openxml.PartRootElement {
	return hmp.RootElement()
}

// AddThemePart adds a theme part to this handout master.
func (hmp *HandoutMasterPart) AddThemePart() (*ThemePart, error) {
	return newThemePartForHandoutMaster(hmp)
}

// ThemePart returns the theme part if present.
func (hmp *HandoutMasterPart) ThemePart() *ThemePart {
	for part := range hmp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (hmp *HandoutMasterPart) GetStream() io.Reader {
	return hmp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (hmp *HandoutMasterPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		hmp,
		uri,
		contentType,
		relType,
	)
}

// Ensure HandoutMasterPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*HandoutMasterPart)(
	nil,
)

// HandoutMasterPartFactory creates a HandoutMasterPart from a URI and container.
func HandoutMasterPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeHandoutMaster,
		packPart,
		container,
	)

	return &HandoutMasterPart{
		OpenXmlPartData: partData,
	}
}

// Register the HandoutMasterPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeHandoutMaster,
			RelationshipType:   RelationshipTypeHandoutMaster,
			Factory:            HandoutMasterPartFactory,
			DefaultURI:         "/ppt/handoutMasters/handoutMaster1.xml",
			IsFixedContentType: true,
		},
	)
}
