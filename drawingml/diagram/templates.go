package diagram

import "strings"

// TemplateType represents the type of diagram template.
type TemplateType string

const (
	// TemplateTypeList represents a basic list layout template.
	TemplateTypeList TemplateType = "list"
	// TemplateTypeHierarchy represents a hierarchical organization chart layout.
	TemplateTypeHierarchy TemplateType = "hierarchy"
)

// Templates contains embedded XML templates for diagram parts.
// These are minimal default templates that allow creating functional SmartArt diagrams.
var Templates = struct {
	// LayoutTemplates contains layout definition templates by type.
	LayoutTemplates map[TemplateType]string
	// StyleTemplates contains style definition templates by type.
	StyleTemplates map[TemplateType]string
	// ColorTemplates contains color definition templates by type.
	ColorTemplates map[TemplateType]string
}{
	LayoutTemplates: map[TemplateType]string{
		TemplateTypeList: `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dgm:layoutDef xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" uniqueId="urn:microsoft.com/office/officeart/2005/8/layout/list1" minVer="12.0">
  <dgm:title val=""/>
  <dgm:desc val=""/>
  <dgm:catLst>
    <dgm:cat type="list" pri="4000"/>
  </dgm:catLst>
  <dgm:sampData>
    <dgm:dataModel>
      <dgm:ptLst>
        <dgm:pt modelId="0" type="doc"/>
        <dgm:pt modelId="1"><dgm:prSet phldr="1"/></dgm:pt>
        <dgm:pt modelId="2"><dgm:prSet phldr="1"/></dgm:pt>
        <dgm:pt modelId="3"><dgm:prSet phldr="1"/></dgm:pt>
      </dgm:ptLst>
      <dgm:cxnLst>
        <dgm:cxn modelId="4" srcId="0" destId="1" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="5" srcId="0" destId="2" srcOrd="1" destOrd="0"/>
        <dgm:cxn modelId="6" srcId="0" destId="3" srcOrd="2" destOrd="0"/>
      </dgm:cxnLst>
      <dgm:bg/>
      <dgm:whole/>
    </dgm:dataModel>
  </dgm:sampData>
  <dgm:styleData>
    <dgm:dataModel>
      <dgm:ptLst>
        <dgm:pt modelId="0" type="doc"/>
        <dgm:pt modelId="1"/>
        <dgm:pt modelId="2"/>
      </dgm:ptLst>
      <dgm:cxnLst>
        <dgm:cxn modelId="4" srcId="0" destId="1" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="5" srcId="0" destId="2" srcOrd="1" destOrd="0"/>
      </dgm:cxnLst>
      <dgm:bg/>
      <dgm:whole/>
    </dgm:dataModel>
  </dgm:styleData>
  <dgm:clrData>
    <dgm:dataModel>
      <dgm:ptLst>
        <dgm:pt modelId="0" type="doc"/>
        <dgm:pt modelId="1"/>
        <dgm:pt modelId="2"/>
        <dgm:pt modelId="3"/>
        <dgm:pt modelId="4"/>
      </dgm:ptLst>
      <dgm:cxnLst>
        <dgm:cxn modelId="5" srcId="0" destId="1" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="6" srcId="0" destId="2" srcOrd="1" destOrd="0"/>
        <dgm:cxn modelId="7" srcId="0" destId="3" srcOrd="2" destOrd="0"/>
        <dgm:cxn modelId="8" srcId="0" destId="4" srcOrd="3" destOrd="0"/>
      </dgm:cxnLst>
      <dgm:bg/>
      <dgm:whole/>
    </dgm:dataModel>
  </dgm:clrData>
  <dgm:layoutNode name="linear">
    <dgm:varLst>
      <dgm:dir/>
      <dgm:animLvl val="lvl"/>
      <dgm:resizeHandles val="exact"/>
    </dgm:varLst>
    <dgm:choose name="Name0">
      <dgm:if name="Name1" func="var" arg="dir" op="equ" val="norm">
        <dgm:alg type="lin">
          <dgm:param type="linDir" val="fromT"/>
          <dgm:param type="vertAlign" val="mid"/>
          <dgm:param type="horzAlign" val="l"/>
          <dgm:param type="nodeHorzAlign" val="l"/>
        </dgm:alg>
      </dgm:if>
      <dgm:else name="Name2">
        <dgm:alg type="lin">
          <dgm:param type="linDir" val="fromT"/>
          <dgm:param type="vertAlign" val="mid"/>
          <dgm:param type="horzAlign" val="r"/>
          <dgm:param type="nodeHorzAlign" val="r"/>
        </dgm:alg>
      </dgm:else>
    </dgm:choose>
    <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
      <dgm:adjLst/>
    </dgm:shape>
    <dgm:presOf/>
    <dgm:constrLst>
      <dgm:constr type="w" for="ch" forName="parentLin" refType="w"/>
      <dgm:constr type="h" for="ch" forName="parentLin" val="INF"/>
      <dgm:constr type="w" for="des" forName="parentLeftMargin" refType="w" fact="0.05"/>
      <dgm:constr type="w" for="des" forName="parentText" refType="w" fact="0.7"/>
      <dgm:constr type="h" for="des" forName="parentText" refType="primFontSz" refFor="des" refForName="parentText" fact="0.82"/>
      <dgm:constr type="h" for="ch" forName="negativeSpace" refType="primFontSz" refFor="des" refForName="parentText" fact="-0.41"/>
      <dgm:constr type="h" for="ch" forName="negativeSpace" refType="h" refFor="des" refForName="parentText" op="lte" fact="-0.82"/>
      <dgm:constr type="h" for="ch" forName="negativeSpace" refType="h" refFor="des" refForName="parentText" op="gte" fact="-0.82"/>
      <dgm:constr type="w" for="ch" forName="childText" refType="w"/>
      <dgm:constr type="h" for="ch" forName="childText" refType="primFontSz" refFor="des" refForName="parentText" fact="0.7"/>
      <dgm:constr type="primFontSz" for="des" forName="parentText" val="65"/>
      <dgm:constr type="primFontSz" for="ch" forName="childText" refType="primFontSz" refFor="des" refForName="parentText"/>
      <dgm:constr type="tMarg" for="ch" forName="childText" refType="primFontSz" refFor="des" refForName="parentText" fact="1.64"/>
      <dgm:constr type="tMarg" for="ch" forName="childText" refType="h" refFor="des" refForName="parentText" op="lte" fact="3.28"/>
      <dgm:constr type="tMarg" for="ch" forName="childText" refType="h" refFor="des" refForName="parentText" op="gte" fact="3.28"/>
      <dgm:constr type="lMarg" for="ch" forName="childText" refType="w" fact="0.22"/>
      <dgm:constr type="rMarg" for="ch" forName="childText" refType="lMarg" refFor="ch" refForName="childText"/>
      <dgm:constr type="lMarg" for="des" forName="parentText" refType="w" fact="0.075"/>
      <dgm:constr type="rMarg" for="des" forName="parentText" refType="lMarg" refFor="des" refForName="parentText"/>
      <dgm:constr type="h" for="ch" forName="spaceBetweenRectangles" refType="primFontSz" refFor="des" refForName="parentText" fact="0.15"/>
    </dgm:constrLst>
    <dgm:ruleLst>
      <dgm:rule type="primFontSz" for="des" forName="parentText" val="5" fact="NaN" max="NaN"/>
    </dgm:ruleLst>
    <dgm:forEach name="Name3" axis="ch" ptType="node">
      <dgm:layoutNode name="parentLin">
        <dgm:choose name="Name4">
          <dgm:if name="Name5" func="var" arg="dir" op="equ" val="norm">
            <dgm:alg type="lin">
              <dgm:param type="linDir" val="fromL"/>
              <dgm:param type="horzAlign" val="l"/>
              <dgm:param type="nodeHorzAlign" val="l"/>
            </dgm:alg>
          </dgm:if>
          <dgm:else name="Name6">
            <dgm:alg type="lin">
              <dgm:param type="linDir" val="fromR"/>
              <dgm:param type="horzAlign" val="r"/>
              <dgm:param type="nodeHorzAlign" val="r"/>
            </dgm:alg>
          </dgm:else>
        </dgm:choose>
        <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
          <dgm:adjLst/>
        </dgm:shape>
        <dgm:presOf/>
        <dgm:constrLst/>
        <dgm:ruleLst/>
        <dgm:layoutNode name="parentLeftMargin">
          <dgm:alg type="sp"/>
          <dgm:shape type="rect" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="" hideGeom="1">
            <dgm:adjLst/>
          </dgm:shape>
          <dgm:presOf axis="self"/>
          <dgm:constrLst>
            <dgm:constr type="h"/>
          </dgm:constrLst>
          <dgm:ruleLst/>
        </dgm:layoutNode>
        <dgm:layoutNode name="parentText" styleLbl="node1">
          <dgm:varLst>
            <dgm:chMax val="0"/>
            <dgm:bulletEnabled val="1"/>
          </dgm:varLst>
          <dgm:choose name="Name7">
            <dgm:if name="Name8" func="var" arg="dir" op="equ" val="norm">
              <dgm:alg type="tx">
                <dgm:param type="parTxLTRAlign" val="l"/>
                <dgm:param type="parTxRTLAlign" val="l"/>
              </dgm:alg>
            </dgm:if>
            <dgm:else name="Name9">
              <dgm:alg type="tx">
                <dgm:param type="parTxLTRAlign" val="r"/>
                <dgm:param type="parTxRTLAlign" val="r"/>
              </dgm:alg>
            </dgm:else>
          </dgm:choose>
          <dgm:shape type="roundRect" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
            <dgm:adjLst/>
          </dgm:shape>
          <dgm:presOf axis="self" ptType="node"/>
          <dgm:constrLst>
            <dgm:constr type="tMarg"/>
            <dgm:constr type="bMarg"/>
          </dgm:constrLst>
          <dgm:ruleLst/>
        </dgm:layoutNode>
      </dgm:layoutNode>
      <dgm:layoutNode name="negativeSpace">
        <dgm:alg type="sp"/>
        <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
          <dgm:adjLst/>
        </dgm:shape>
        <dgm:presOf/>
        <dgm:constrLst/>
        <dgm:ruleLst/>
      </dgm:layoutNode>
      <dgm:layoutNode name="childText" styleLbl="conFgAcc1">
        <dgm:varLst>
          <dgm:bulletEnabled val="1"/>
        </dgm:varLst>
        <dgm:alg type="tx">
          <dgm:param type="stBulletLvl" val="1"/>
        </dgm:alg>
        <dgm:shape type="rect" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="" zOrderOff="-2">
          <dgm:adjLst/>
        </dgm:shape>
        <dgm:presOf axis="des" ptType="node"/>
        <dgm:constrLst>
          <dgm:constr type="secFontSz" refType="primFontSz"/>
        </dgm:constrLst>
        <dgm:ruleLst>
          <dgm:rule type="h" val="INF" fact="NaN" max="NaN"/>
        </dgm:ruleLst>
      </dgm:layoutNode>
      <dgm:forEach name="Name10" axis="followSib" ptType="sibTrans" cnt="1">
        <dgm:layoutNode name="spaceBetweenRectangles">
          <dgm:alg type="sp"/>
          <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
            <dgm:adjLst/>
          </dgm:shape>
          <dgm:presOf/>
          <dgm:constrLst/>
          <dgm:ruleLst/>
        </dgm:layoutNode>
      </dgm:forEach>
    </dgm:forEach>
  </dgm:layoutNode>
</dgm:layoutDef>`,
		TemplateTypeHierarchy: `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dgm:layoutDef xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" uniqueId="urn:microsoft.com/office/officeart/2005/8/layout/orgChart1" minVer="12.0">
  <dgm:title val=""/>
  <dgm:desc val=""/>
  <dgm:catLst>
    <dgm:cat type="hierarchy" pri="1000"/>
    <dgm:cat type="convert" pri="6000"/>
  </dgm:catLst>
  <dgm:sampData>
    <dgm:dataModel>
      <dgm:ptLst>
        <dgm:pt modelId="0" type="doc"/>
        <dgm:pt modelId="1"><dgm:prSet phldr="1"/></dgm:pt>
        <dgm:pt modelId="2" type="asst"><dgm:prSet phldr="1"/></dgm:pt>
        <dgm:pt modelId="3"><dgm:prSet phldr="1"/></dgm:pt>
        <dgm:pt modelId="4"><dgm:prSet phldr="1"/></dgm:pt>
        <dgm:pt modelId="5"><dgm:prSet phldr="1"/></dgm:pt>
      </dgm:ptLst>
      <dgm:cxnLst>
        <dgm:cxn modelId="5" srcId="0" destId="1" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="6" srcId="1" destId="2" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="7" srcId="1" destId="3" srcOrd="1" destOrd="0"/>
        <dgm:cxn modelId="8" srcId="1" destId="4" srcOrd="2" destOrd="0"/>
        <dgm:cxn modelId="9" srcId="1" destId="5" srcOrd="3" destOrd="0"/>
      </dgm:cxnLst>
      <dgm:bg/>
      <dgm:whole/>
    </dgm:dataModel>
  </dgm:sampData>
  <dgm:styleData>
    <dgm:dataModel>
      <dgm:ptLst>
        <dgm:pt modelId="0" type="doc"/>
        <dgm:pt modelId="1"/>
        <dgm:pt modelId="12"/>
        <dgm:pt modelId="13"/>
      </dgm:ptLst>
      <dgm:cxnLst>
        <dgm:cxn modelId="2" srcId="0" destId="1" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="16" srcId="1" destId="12" srcOrd="1" destOrd="0"/>
        <dgm:cxn modelId="17" srcId="1" destId="13" srcOrd="2" destOrd="0"/>
      </dgm:cxnLst>
      <dgm:bg/>
      <dgm:whole/>
    </dgm:dataModel>
  </dgm:styleData>
  <dgm:clrData>
    <dgm:dataModel>
      <dgm:ptLst>
        <dgm:pt modelId="0" type="doc"/>
        <dgm:pt modelId="1"/>
        <dgm:pt modelId="11" type="asst"/>
        <dgm:pt modelId="12"/>
        <dgm:pt modelId="13"/>
        <dgm:pt modelId="14"/>
      </dgm:ptLst>
      <dgm:cxnLst>
        <dgm:cxn modelId="2" srcId="0" destId="1" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="15" srcId="1" destId="11" srcOrd="0" destOrd="0"/>
        <dgm:cxn modelId="16" srcId="1" destId="12" srcOrd="1" destOrd="0"/>
        <dgm:cxn modelId="17" srcId="1" destId="13" srcOrd="2" destOrd="0"/>
        <dgm:cxn modelId="18" srcId="1" destId="14" srcOrd="2" destOrd="0"/>
      </dgm:cxnLst>
      <dgm:bg/>
      <dgm:whole/>
    </dgm:dataModel>
  </dgm:clrData>
  <dgm:layoutNode name="hierChild1">
    <dgm:varLst>
      <dgm:orgChart val="1"/>
      <dgm:chPref val="1"/>
      <dgm:dir/>
      <dgm:animOne val="branch"/>
      <dgm:animLvl val="lvl"/>
      <dgm:resizeHandles/>
    </dgm:varLst>
    <dgm:choose name="Name0">
      <dgm:if name="Name1" func="var" arg="dir" op="equ" val="norm">
        <dgm:alg type="hierChild">
          <dgm:param type="linDir" val="fromL"/>
        </dgm:alg>
      </dgm:if>
      <dgm:else name="Name2">
        <dgm:alg type="hierChild">
          <dgm:param type="linDir" val="fromR"/>
        </dgm:alg>
      </dgm:else>
    </dgm:choose>
    <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
      <dgm:adjLst/>
    </dgm:shape>
    <dgm:presOf/>
    <dgm:constrLst>
      <dgm:constr type="sibSp" refType="w" refFor="des" refForName="rootComposite1" fact="0.21"/>
      <dgm:constr type="secSibSp" refType="w" refFor="des" refForName="rootComposite1" fact="0.21"/>
      <dgm:constr type="primFontSz" for="des" ptType="node" op="equ"/>
    </dgm:constrLst>
    <dgm:ruleLst/>
    <dgm:forEach name="Name3" axis="ch">
      <dgm:forEach name="Name4" axis="self" ptType="node">
        <dgm:layoutNode name="hierRoot1">
          <dgm:varLst>
            <dgm:hierBranch val="init"/>
          </dgm:varLst>
          <dgm:choose name="Name5">
            <dgm:if name="Name6" func="var" arg="hierBranch" op="equ" val="l">
              <dgm:choose name="Name7">
                <dgm:if name="Name8" axis="ch" ptType="asst" func="cnt" op="gte" val="1">
                  <dgm:alg type="hierRoot">
                    <dgm:param type="hierAlign" val="tR"/>
                  </dgm:alg>
                </dgm:if>
                <dgm:else name="Name9">
                  <dgm:alg type="hierRoot">
                    <dgm:param type="hierAlign" val="tR"/>
                  </dgm:alg>
                </dgm:else>
              </dgm:choose>
            </dgm:if>
            <dgm:if name="Name10" func="var" arg="hierBranch" op="equ" val="r">
              <dgm:choose name="Name11">
                <dgm:if name="Name12" axis="ch" ptType="asst" func="cnt" op="gte" val="1">
                  <dgm:alg type="hierRoot">
                    <dgm:param type="hierAlign" val="tL"/>
                  </dgm:alg>
                </dgm:if>
                <dgm:else name="Name13">
                  <dgm:alg type="hierRoot">
                    <dgm:param type="hierAlign" val="tL"/>
                  </dgm:alg>
                </dgm:else>
              </dgm:choose>
            </dgm:if>
            <dgm:else name="Name14">
              <dgm:alg type="hierRoot"/>
            </dgm:else>
          </dgm:choose>
          <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
            <dgm:adjLst/>
          </dgm:shape>
          <dgm:presOf/>
          <dgm:ruleLst/>
          <dgm:layoutNode name="rootComposite1">
            <dgm:alg type="composite"/>
            <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
              <dgm:adjLst/>
            </dgm:shape>
            <dgm:presOf axis="self" ptType="node" cnt="1"/>
            <dgm:ruleLst/>
            <dgm:layoutNode name="rootText1" styleLbl="node0">
              <dgm:varLst>
                <dgm:chPref val="3"/>
              </dgm:varLst>
              <dgm:alg type="tx"/>
              <dgm:shape type="rect" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
                <dgm:adjLst/>
              </dgm:shape>
              <dgm:presOf axis="self" ptType="node" cnt="1"/>
              <dgm:constrLst>
                <dgm:constr type="primFontSz" val="100"/>
              </dgm:constrLst>
              <dgm:ruleLst>
                <dgm:rule type="primFontSz" val="2" fact="NaN" max="NaN"/>
              </dgm:ruleLst>
            </dgm:layoutNode>
            <dgm:layoutNode name="rootConnector1" moveWith="rootText1">
              <dgm:alg type="sp"/>
              <dgm:shape type="rect" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="" hideGeom="1">
                <dgm:adjLst/>
              </dgm:shape>
              <dgm:presOf axis="self" ptType="node" cnt="1"/>
              <dgm:constrLst/>
              <dgm:ruleLst/>
            </dgm:layoutNode>
          </dgm:layoutNode>
        </dgm:layoutNode>
        <dgm:layoutNode name="hierChild2">
          <dgm:alg type="hierChild"/>
          <dgm:shape xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" r:blip="">
            <dgm:adjLst/>
          </dgm:shape>
          <dgm:presOf/>
          <dgm:constrLst/>
          <dgm:ruleLst/>
        </dgm:layoutNode>
      </dgm:forEach>
    </dgm:forEach>
  </dgm:layoutNode>
</dgm:layoutDef>`,
	},
	StyleTemplates: map[TemplateType]string{
		TemplateTypeList: `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dgm:styleDef xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" uniqueId="urn:microsoft.com/office/officeart/2005/8/quickstyle/simple1">
  <dgm:title val=""/>
  <dgm:desc val=""/>
  <dgm:catLst>
    <dgm:cat type="simple" pri="10100"/>
  </dgm:catLst>
  <dgm:scene3d>
    <a:camera prst="orthographicFront"/>
    <a:lightRig rig="threePt" dir="t"/>
  </dgm:scene3d>
  <dgm:styleLbl name="node0">
    <dgm:scene3d>
      <a:camera prst="orthographicFront"/>
      <a:lightRig rig="threePt" dir="t"/>
    </dgm:scene3d>
    <dgm:sp3d/>
    <dgm:txPr/>
    <dgm:style>
      <a:lnRef idx="2">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:lnRef>
      <a:fillRef idx="1">
        <a:schemeClr val="accent1"/>
      </a:fillRef>
      <a:effectRef idx="0">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:effectRef>
      <a:fontRef idx="minor">
        <a:schemeClr val="lt1"/>
      </a:fontRef>
    </dgm:style>
  </dgm:styleLbl>
  <dgm:styleLbl name="node1">
    <dgm:scene3d>
      <a:camera prst="orthographicFront"/>
      <a:lightRig rig="threePt" dir="t"/>
    </dgm:scene3d>
    <dgm:sp3d/>
    <dgm:txPr/>
    <dgm:style>
      <a:lnRef idx="2">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:lnRef>
      <a:fillRef idx="1">
        <a:schemeClr val="accent1"/>
      </a:fillRef>
      <a:effectRef idx="0">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:effectRef>
      <a:fontRef idx="minor">
        <a:schemeClr val="lt1"/>
      </a:fontRef>
    </dgm:style>
  </dgm:styleLbl>
  <dgm:styleLbl name="conFgAcc1">
    <dgm:scene3d>
      <a:camera prst="orthographicFront"/>
      <a:lightRig rig="threePt" dir="t"/>
    </dgm:scene3d>
    <dgm:sp3d/>
    <dgm:txPr/>
    <dgm:style>
      <a:lnRef idx="2">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:lnRef>
      <a:fillRef idx="1">
        <a:schemeClr val="accent1"/>
      </a:fillRef>
      <a:effectRef idx="0">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:effectRef>
      <a:fontRef idx="minor">
        <a:schemeClr val="lt1"/>
      </a:fontRef>
    </dgm:style>
  </dgm:styleLbl>
</dgm:styleDef>`,
		TemplateTypeHierarchy: `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dgm:styleDef xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" uniqueId="urn:microsoft.com/office/officeart/2005/8/quickstyle/simple1">
  <dgm:title val=""/>
  <dgm:desc val=""/>
  <dgm:catLst>
    <dgm:cat type="simple" pri="10100"/>
  </dgm:catLst>
  <dgm:scene3d>
    <a:camera prst="orthographicFront"/>
    <a:lightRig rig="threePt" dir="t"/>
  </dgm:scene3d>
  <dgm:styleLbl name="node0">
    <dgm:scene3d>
      <a:camera prst="orthographicFront"/>
      <a:lightRig rig="threePt" dir="t"/>
    </dgm:scene3d>
    <dgm:sp3d/>
    <dgm:txPr/>
    <dgm:style>
      <a:lnRef idx="2">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:lnRef>
      <a:fillRef idx="1">
        <a:schemeClr val="accent2"/>
      </a:fillRef>
      <a:effectRef idx="0">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:effectRef>
      <a:fontRef idx="minor">
        <a:schemeClr val="lt1"/>
      </a:fontRef>
    </dgm:style>
  </dgm:styleLbl>
  <dgm:styleLbl name="node1">
    <dgm:scene3d>
      <a:camera prst="orthographicFront"/>
      <a:lightRig rig="threePt" dir="t"/>
    </dgm:scene3d>
    <dgm:sp3d/>
    <dgm:txPr/>
    <dgm:style>
      <a:lnRef idx="2">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:lnRef>
      <a:fillRef idx="1">
        <a:schemeClr val="accent2"/>
      </a:fillRef>
      <a:effectRef idx="0">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:effectRef>
      <a:fontRef idx="minor">
        <a:schemeClr val="lt1"/>
      </a:fontRef>
    </dgm:style>
  </dgm:styleLbl>
  <dgm:styleLbl name="asst0">
    <dgm:scene3d>
      <a:camera prst="orthographicFront"/>
      <a:lightRig rig="threePt" dir="t"/>
    </dgm:scene3d>
    <dgm:sp3d/>
    <dgm:txPr/>
    <dgm:style>
      <a:lnRef idx="2">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:lnRef>
      <a:fillRef idx="1">
        <a:schemeClr val="accent3"/>
      </a:fillRef>
      <a:effectRef idx="0">
        <a:scrgbClr r="0" g="0" b="0"/>
      </a:effectRef>
      <a:fontRef idx="minor">
        <a:schemeClr val="lt1"/>
      </a:fontRef>
    </dgm:style>
  </dgm:styleLbl>
</dgm:styleDef>`,
	},
	ColorTemplates: map[TemplateType]string{
		TemplateTypeList: `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dgm:colorsDef xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" uniqueId="urn:microsoft.com/office/officeart/2005/8/colors/accent1_2">
  <dgm:title val=""/>
  <dgm:desc val=""/>
  <dgm:catLst>
    <dgm:cat type="accent1" pri="11200"/>
  </dgm:catLst>
  <dgm:styleLbl name="node0">
    <dgm:fillClrLst meth="repeat">
      <a:schemeClr val="accent1"/>
    </dgm:fillClrLst>
    <dgm:linClrLst meth="repeat">
      <a:schemeClr val="lt1"/>
    </dgm:linClrLst>
    <dgm:effectClrLst/>
    <dgm:txLinClrLst/>
    <dgm:txFillClrLst/>
    <dgm:txEffectClrLst/>
  </dgm:styleLbl>
  <dgm:styleLbl name="node1">
    <dgm:fillClrLst meth="repeat">
      <a:schemeClr val="accent1"/>
    </dgm:fillClrLst>
    <dgm:linClrLst meth="repeat">
      <a:schemeClr val="lt1"/>
    </dgm:linClrLst>
    <dgm:effectClrLst/>
    <dgm:txLinClrLst/>
    <dgm:txFillClrLst/>
    <dgm:txEffectClrLst/>
  </dgm:styleLbl>
  <dgm:styleLbl name="conFgAcc1">
    <dgm:fillClrLst meth="repeat">
      <a:schemeClr val="lt1">
        <a:alpha val="90000"/>
      </a:schemeClr>
    </dgm:fillClrLst>
    <dgm:linClrLst meth="repeat">
      <a:schemeClr val="accent1"/>
    </dgm:linClrLst>
    <dgm:effectClrLst/>
    <dgm:txLinClrLst/>
    <dgm:txFillClrLst meth="repeat">
      <a:schemeClr val="dk1"/>
    </dgm:txFillClrLst>
    <dgm:txEffectClrLst/>
  </dgm:styleLbl>
</dgm:colorsDef>`,
		TemplateTypeHierarchy: `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dgm:colorsDef xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" uniqueId="urn:microsoft.com/office/officeart/2005/8/colors/accent2_3">
  <dgm:title val=""/>
  <dgm:desc val=""/>
  <dgm:catLst>
    <dgm:cat type="accent2" pri="11300"/>
  </dgm:catLst>
  <dgm:styleLbl name="node0">
    <dgm:fillClrLst meth="repeat">
      <a:schemeClr val="accent2"/>
    </dgm:fillClrLst>
    <dgm:linClrLst meth="repeat">
      <a:schemeClr val="lt1"/>
    </dgm:linClrLst>
    <dgm:effectClrLst/>
    <dgm:txLinClrLst/>
    <dgm:txFillClrLst/>
    <dgm:txEffectClrLst/>
  </dgm:styleLbl>
  <dgm:styleLbl name="node1">
    <dgm:fillClrLst meth="repeat">
      <a:schemeClr val="accent2"/>
    </dgm:fillClrLst>
    <dgm:linClrLst meth="repeat">
      <a:schemeClr val="lt1"/>
    </dgm:linClrLst>
    <dgm:effectClrLst/>
    <dgm:txLinClrLst/>
    <dgm:txFillClrLst/>
    <dgm:txEffectClrLst/>
  </dgm:styleLbl>
  <dgm:styleLbl name="asst0">
    <dgm:fillClrLst meth="repeat">
      <a:schemeClr val="accent3"/>
    </dgm:fillClrLst>
    <dgm:linClrLst meth="repeat">
      <a:schemeClr val="lt1"/>
    </dgm:linClrLst>
    <dgm:effectClrLst/>
    <dgm:txLinClrLst/>
    <dgm:txFillClrLst/>
    <dgm:txEffectClrLst/>
  </dgm:styleLbl>
</dgm:colorsDef>`,
	},
}

// GetLayoutTemplate returns the layout template XML for the specified type.
// If the template type is not found, it returns the default list template.
func GetLayoutTemplate(templateType TemplateType) string {
	if template, ok := Templates.LayoutTemplates[templateType]; ok {
		return strings.TrimSpace(template)
	}
	// Return default list template if type not found
	return strings.TrimSpace(Templates.LayoutTemplates[TemplateTypeList])
}

// GetStyleTemplate returns the style template XML for the specified type.
// If the template type is not found, it returns the default list template.
func GetStyleTemplate(templateType TemplateType) string {
	if template, ok := Templates.StyleTemplates[templateType]; ok {
		return strings.TrimSpace(template)
	}
	// Return default list template if type not found
	return strings.TrimSpace(Templates.StyleTemplates[TemplateTypeList])
}

// GetColorTemplate returns the color template XML for the specified type.
// If the template type is not found, it returns the default list template.
func GetColorTemplate(templateType TemplateType) string {
	if template, ok := Templates.ColorTemplates[templateType]; ok {
		return strings.TrimSpace(template)
	}
	// Return default list template if type not found
	return strings.TrimSpace(Templates.ColorTemplates[TemplateTypeList])
}
