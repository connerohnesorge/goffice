//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
)

// NotesMasterPart represents a notes master part (ppt/notesMasters/notesMaster1.xml).
type NotesMasterPart struct {
	*openxml.OpenXmlPartData
}

// newNotesMasterPart creates a new notes master part.
func newNotesMasterPart(
	presentationPart *PresentationPart,
	uri string,
) (*NotesMasterPart, error) {
	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeNotesMaster,
		RelationshipTypeNotesMaster,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNotesMaster,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	nmp := &NotesMasterPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal notes master content
	nmp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(nmp, relID); err != nil {
		return nil, err
	}

	return nmp, nil
}

// initializeContent sets up minimal notes master content.
func (nmp *NotesMasterPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:notesMaster xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
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
          <p:cNvPr id="3" name="Slide Image Placeholder 2"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1" noRot="1" noChangeAspect="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="sldImg" idx="2"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="1371600" y="1143000"/>
            <a:ext cx="4114800" cy="3086100"/>
          </a:xfrm>
          <a:prstGeom prst="rect">
            <a:avLst/>
          </a:prstGeom>
          <a:noFill/>
          <a:ln w="12700">
            <a:solidFill>
              <a:prstClr val="black"/>
            </a:solidFill>
          </a:ln>
        </p:spPr>
        <p:txBody>
          <a:bodyPr vert="horz" lIns="91440" tIns="45720" rIns="91440" bIns="45720" rtlCol="0" anchor="ctr"/>
          <a:lstStyle/>
          <a:p>
            <a:endParaRPr lang="en-US"/>
          </a:p>
        </p:txBody>
      </p:sp>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="4" name="Notes Placeholder 3"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="body" sz="quarter" idx="3"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="685800" y="4400550"/>
            <a:ext cx="5486400" cy="3600450"/>
          </a:xfrm>
          <a:prstGeom prst="rect">
            <a:avLst/>
          </a:prstGeom>
        </p:spPr>
        <p:txBody>
          <a:bodyPr vert="horz" lIns="91440" tIns="45720" rIns="91440" bIns="45720" rtlCol="0"/>
          <a:lstStyle/>
          <a:p>
            <a:pPr lvl="0"/>
            <a:r>
              <a:rPr lang="en-US"/>
              <a:t>Click to edit Master text styles</a:t>
            </a:r>
          </a:p>
        </p:txBody>
      </p:sp>
    </p:spTree>
  </p:cSld>
  <p:clrMap bg1="lt1" tx1="dk1" bg2="lt2" tx2="dk2" accent1="accent1" accent2="accent2" accent3="accent3" accent4="accent4" accent5="accent5" accent6="accent6" hlink="hlink" folHlink="folHlink"/>
</p:notesMaster>`
	nmp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*NotesMasterPart) FixedContentType() string {
	return ContentTypeNotesMaster
}

// NotesMaster returns the root NotesMaster element.
// TODO: Return a proper NotesMaster element type when elements are implemented.
func (nmp *NotesMasterPart) NotesMaster() openxml.PartRootElement {
	return nmp.RootElement()
}

// AddThemePart adds a theme part to this notes master.
func (nmp *NotesMasterPart) AddThemePart() (*ThemePart, error) {
	return newThemePartForNotesMaster(nmp)
}

// ThemePart returns the theme part if present.
func (nmp *NotesMasterPart) ThemePart() *ThemePart {
	for part := range nmp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (nmp *NotesMasterPart) GetStream() io.Reader {
	return nmp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (nmp *NotesMasterPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		nmp,
		uri,
		contentType,
		relType,
	)
}

// Ensure NotesMasterPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*NotesMasterPart)(
	nil,
)

// NotesMasterPartFactory creates a NotesMasterPart from a URI and container.
func NotesMasterPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNotesMaster,
		packPart,
		container,
	)

	return &NotesMasterPart{
		OpenXmlPartData: partData,
	}
}

// Register the NotesMasterPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeNotesMaster,
			RelationshipType:   RelationshipTypeNotesMaster,
			Factory:            NotesMasterPartFactory,
			DefaultURI:         "/ppt/notesMasters/notesMaster1.xml",
			IsFixedContentType: true,
		},
	)
}
