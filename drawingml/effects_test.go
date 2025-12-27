package drawingml

import (
	"strings"
	"testing"
)

func TestNewEffectList(t *testing.T) {
	effectList := NewEffectList()
	if effectList == nil {
		t.Fatal("NewEffectList returned nil")
	}
	if effectList.LocalName() != "effectLst" {
		t.Errorf(
			"expected local name 'effectLst', got '%s'",
			effectList.LocalName(),
		)
	}
	if effectList.NamespaceURI() != NamespaceMain {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceMain,
			effectList.NamespaceURI(),
		)
	}
}

func TestEffectListBlur(t *testing.T) {
	effectList := NewEffectList()

	// Initially nil
	if effectList.Blur() != nil {
		t.Error(
			"expected Blur to be nil initially",
		)
	}

	// Set blur
	blur := NewBlur(50000)
	effectList.SetBlur(blur)

	// Get blur
	retrieved := effectList.Blur()
	if retrieved == nil {
		t.Fatal("expected Blur to be set")
	}
	if retrieved.Radius() != 50000 {
		t.Errorf(
			"expected radius 50000, got %d",
			retrieved.Radius(),
		)
	}

	// Remove blur
	effectList.SetBlur(nil)
	if effectList.Blur() != nil {
		t.Error(
			"expected Blur to be nil after removal",
		)
	}
}

func TestEffectListGlow(t *testing.T) {
	effectList := NewEffectList()

	// Set glow with RGB color
	glow := NewGlowWithParams(100000, "FF0000")
	effectList.SetGlow(glow)

	// Get glow
	retrieved := effectList.Glow()
	if retrieved == nil {
		t.Fatal("expected Glow to be set")
	}
	if retrieved.Radius() != 100000 {
		t.Errorf(
			"expected radius 100000, got %d",
			retrieved.Radius(),
		)
	}
}

func TestEffectListShadows(t *testing.T) {
	effectList := NewEffectList()

	// Test outer shadow
	outerShadow := NewDropShadow(
		50000,
		100000,
		5400000,
		"808080",
	)
	effectList.SetOuterShadow(outerShadow)
	if effectList.OuterShadow() == nil {
		t.Error("expected OuterShadow to be set")
	}

	// Test inner shadow
	innerShadow := NewInnerShadowWithParams(
		30000,
		50000,
		2700000,
		"404040",
	)
	effectList.SetInnerShadow(innerShadow)
	if effectList.InnerShadow() == nil {
		t.Error("expected InnerShadow to be set")
	}

	// Test preset shadow
	presetShadow := NewPresetShadow(
		PresetShadowTopLeftDropShadow,
	)
	effectList.SetPresetShadow(presetShadow)
	if effectList.PresetShadow() == nil {
		t.Error("expected PresetShadow to be set")
	}
}

func TestOuterShadow(t *testing.T) {
	shadow := NewOuterShadow()

	// Test blur radius
	shadow.SetBlurRadius(75000)
	if shadow.BlurRadius() != 75000 {
		t.Errorf(
			"expected blur radius 75000, got %d",
			shadow.BlurRadius(),
		)
	}

	// Test distance
	shadow.SetDistance(100000)
	if shadow.Distance() != 100000 {
		t.Errorf(
			"expected distance 100000, got %d",
			shadow.Distance(),
		)
	}

	// Test direction
	shadow.SetDirection(5400000)
	if shadow.Direction() != 5400000 {
		t.Errorf(
			"expected direction 5400000, got %d",
			shadow.Direction(),
		)
	}

	// Test horizontal ratio
	shadow.SetHorizontalRatio(120000)
	if shadow.HorizontalRatio() != 120000 {
		t.Errorf(
			"expected horizontal ratio 120000, got %d",
			shadow.HorizontalRatio(),
		)
	}

	// Test vertical ratio
	shadow.SetVerticalRatio(80000)
	if shadow.VerticalRatio() != 80000 {
		t.Errorf(
			"expected vertical ratio 80000, got %d",
			shadow.VerticalRatio(),
		)
	}

	// Test horizontal skew
	shadow.SetHorizontalSkew(300000)
	if shadow.HorizontalSkew() != 300000 {
		t.Errorf(
			"expected horizontal skew 300000, got %d",
			shadow.HorizontalSkew(),
		)
	}

	// Test vertical skew
	shadow.SetVerticalSkew(-300000)
	if shadow.VerticalSkew() != -300000 {
		t.Errorf(
			"expected vertical skew -300000, got %d",
			shadow.VerticalSkew(),
		)
	}

	// Test alignment
	shadow.SetAlignment(RectAlignTopLeft)
	if shadow.Alignment() != RectAlignTopLeft {
		t.Errorf(
			"expected alignment '%s', got '%s'",
			RectAlignTopLeft,
			shadow.Alignment(),
		)
	}

	// Test rotate with shape
	shadow.SetRotateWithShape(false)
	if shadow.RotateWithShape() {
		t.Error(
			"expected RotateWithShape to be false",
		)
	}
	shadow.SetRotateWithShape(true)
	if !shadow.RotateWithShape() {
		t.Error(
			"expected RotateWithShape to be true",
		)
	}

	// Test RGB color
	shadow.SetRgbColor("FF5500")
	rgb := shadow.RgbColor()
	if rgb == nil {
		t.Fatal("expected RGB color to be set")
	}
	if rgb.Value() != "FF5500" {
		t.Errorf(
			"expected color 'FF5500', got '%s'",
			rgb.Value(),
		)
	}

	// Test scheme color
	shadow.SetSchemeColor(SchemeColorAccent1)
	if shadow.RgbColor() != nil {
		t.Error(
			"expected RGB color to be nil after setting scheme color",
		)
	}
	scheme := shadow.SchemeColor()
	if scheme == nil {
		t.Fatal("expected scheme color to be set")
	}
	if scheme.Value() != SchemeColorAccent1 {
		t.Errorf(
			"expected scheme color '%s', got '%s'",
			SchemeColorAccent1,
			scheme.Value(),
		)
	}
}

func TestDropShadowFactory(t *testing.T) {
	shadow := NewDropShadow(
		50000,
		100000,
		5400000,
		"808080",
	)

	if shadow.BlurRadius() != 50000 {
		t.Errorf(
			"expected blur radius 50000, got %d",
			shadow.BlurRadius(),
		)
	}
	if shadow.Distance() != 100000 {
		t.Errorf(
			"expected distance 100000, got %d",
			shadow.Distance(),
		)
	}
	if shadow.Direction() != 5400000 {
		t.Errorf(
			"expected direction 5400000, got %d",
			shadow.Direction(),
		)
	}
	if shadow.RgbColor() == nil ||
		shadow.RgbColor().Value() != "808080" {
		t.Error("expected RGB color '808080'")
	}
}

func TestInnerShadow(t *testing.T) {
	shadow := NewInnerShadow()

	// Test blur radius
	shadow.SetBlurRadius(40000)
	if shadow.BlurRadius() != 40000 {
		t.Errorf(
			"expected blur radius 40000, got %d",
			shadow.BlurRadius(),
		)
	}

	// Test distance
	shadow.SetDistance(60000)
	if shadow.Distance() != 60000 {
		t.Errorf(
			"expected distance 60000, got %d",
			shadow.Distance(),
		)
	}

	// Test direction
	shadow.SetDirection(2700000)
	if shadow.Direction() != 2700000 {
		t.Errorf(
			"expected direction 2700000, got %d",
			shadow.Direction(),
		)
	}

	// Test RGB color
	shadow.SetRgbColor("404040")
	rgb := shadow.RgbColor()
	if rgb == nil {
		t.Fatal("expected RGB color to be set")
	}
	if rgb.Value() != "404040" {
		t.Errorf(
			"expected color '404040', got '%s'",
			rgb.Value(),
		)
	}
}

func TestPresetShadow(t *testing.T) {
	shadow := NewPresetShadow(
		PresetShadowBottomRightDropShadow,
	)

	if shadow.Preset() != PresetShadowBottomRightDropShadow {
		t.Errorf(
			"expected preset '%s', got '%s'",
			PresetShadowBottomRightDropShadow,
			shadow.Preset(),
		)
	}

	// Test distance
	shadow.SetDistance(50000)
	if shadow.Distance() != 50000 {
		t.Errorf(
			"expected distance 50000, got %d",
			shadow.Distance(),
		)
	}

	// Test direction
	shadow.SetDirection(7200000)
	if shadow.Direction() != 7200000 {
		t.Errorf(
			"expected direction 7200000, got %d",
			shadow.Direction(),
		)
	}

	// Test RGB color
	shadow.SetRgbColor("333333")
	rgb := shadow.RgbColor()
	if rgb == nil || rgb.Value() != "333333" {
		t.Error("expected RGB color '333333'")
	}
}

func TestGlow(t *testing.T) {
	glow := NewGlow()

	// Test radius
	glow.SetRadius(150000)
	if glow.Radius() != 150000 {
		t.Errorf(
			"expected radius 150000, got %d",
			glow.Radius(),
		)
	}

	// Test RGB color
	glow.SetRgbColor("00FF00")
	rgb := glow.RgbColor()
	if rgb == nil {
		t.Fatal("expected RGB color to be set")
	}
	if rgb.Value() != "00FF00" {
		t.Errorf(
			"expected color '00FF00', got '%s'",
			rgb.Value(),
		)
	}

	// Test scheme color
	glow.SetSchemeColor(SchemeColorAccent2)
	scheme := glow.SchemeColor()
	if scheme == nil {
		t.Fatal("expected scheme color to be set")
	}
	if scheme.Value() != SchemeColorAccent2 {
		t.Errorf(
			"expected scheme color '%s', got '%s'",
			SchemeColorAccent2,
			scheme.Value(),
		)
	}
}

func TestGlowFactories(t *testing.T) {
	// Test NewGlowWithRadius
	glow1 := NewGlowWithRadius(100000)
	if glow1.Radius() != 100000 {
		t.Errorf(
			"expected radius 100000, got %d",
			glow1.Radius(),
		)
	}

	// Test NewGlowWithParams
	glow2 := NewGlowWithParams(200000, "FF00FF")
	if glow2.Radius() != 200000 {
		t.Errorf(
			"expected radius 200000, got %d",
			glow2.Radius(),
		)
	}
	if glow2.RgbColor() == nil ||
		glow2.RgbColor().Value() != "FF00FF" {
		t.Error("expected RGB color 'FF00FF'")
	}
}

func TestSoftEdge(t *testing.T) {
	softEdge := NewSoftEdge(80000)

	if softEdge.Radius() != 80000 {
		t.Errorf(
			"expected radius 80000, got %d",
			softEdge.Radius(),
		)
	}

	// Test updating radius
	softEdge.SetRadius(120000)
	if softEdge.Radius() != 120000 {
		t.Errorf(
			"expected radius 120000, got %d",
			softEdge.Radius(),
		)
	}
}

func TestReflection(t *testing.T) {
	reflection := NewReflection()

	// Test blur radius
	reflection.SetBlurRadius(30000)
	if reflection.BlurRadius() != 30000 {
		t.Errorf(
			"expected blur radius 30000, got %d",
			reflection.BlurRadius(),
		)
	}

	// Test start opacity
	reflection.SetStartOpacity(50000)
	if reflection.StartOpacity() != 50000 {
		t.Errorf(
			"expected start opacity 50000, got %d",
			reflection.StartOpacity(),
		)
	}

	// Test end alpha
	reflection.SetEndAlpha(10000)
	if reflection.EndAlpha() != 10000 {
		t.Errorf(
			"expected end alpha 10000, got %d",
			reflection.EndAlpha(),
		)
	}

	// Test positions
	reflection.SetStartPosition(10000)
	if reflection.StartPosition() != 10000 {
		t.Errorf(
			"expected start position 10000, got %d",
			reflection.StartPosition(),
		)
	}
	reflection.SetEndPosition(90000)
	if reflection.EndPosition() != 90000 {
		t.Errorf(
			"expected end position 90000, got %d",
			reflection.EndPosition(),
		)
	}

	// Test distance
	reflection.SetDistance(50000)
	if reflection.Distance() != 50000 {
		t.Errorf(
			"expected distance 50000, got %d",
			reflection.Distance(),
		)
	}

	// Test direction
	reflection.SetDirection(5400000)
	if reflection.Direction() != 5400000 {
		t.Errorf(
			"expected direction 5400000, got %d",
			reflection.Direction(),
		)
	}

	// Test fade direction
	reflection.SetFadeDirection(10800000)
	if reflection.FadeDirection() != 10800000 {
		t.Errorf(
			"expected fade direction 10800000, got %d",
			reflection.FadeDirection(),
		)
	}

	// Test scaling
	reflection.SetHorizontalRatio(90000)
	if reflection.HorizontalRatio() != 90000 {
		t.Errorf(
			"expected horizontal ratio 90000, got %d",
			reflection.HorizontalRatio(),
		)
	}
	reflection.SetVerticalRatio(-100000)
	if reflection.VerticalRatio() != -100000 {
		t.Errorf(
			"expected vertical ratio -100000, got %d",
			reflection.VerticalRatio(),
		)
	}

	// Test skew
	reflection.SetHorizontalSkew(100000)
	if reflection.HorizontalSkew() != 100000 {
		t.Errorf(
			"expected horizontal skew 100000, got %d",
			reflection.HorizontalSkew(),
		)
	}
	reflection.SetVerticalSkew(-50000)
	if reflection.VerticalSkew() != -50000 {
		t.Errorf(
			"expected vertical skew -50000, got %d",
			reflection.VerticalSkew(),
		)
	}

	// Test alignment
	reflection.SetAlignment(RectAlignTop)
	if reflection.Alignment() != RectAlignTop {
		t.Errorf(
			"expected alignment '%s', got '%s'",
			RectAlignTop,
			reflection.Alignment(),
		)
	}

	// Test rotate with shape
	reflection.SetRotateWithShape(false)
	if reflection.RotateWithShape() {
		t.Error(
			"expected RotateWithShape to be false",
		)
	}
}

func TestReflectionWithDefaults(t *testing.T) {
	reflection := NewReflectionWithDefaults()

	if reflection.StartOpacity() != 100000 {
		t.Errorf(
			"expected start opacity 100000, got %d",
			reflection.StartOpacity(),
		)
	}
	if reflection.EndAlpha() != 0 {
		t.Errorf(
			"expected end alpha 0, got %d",
			reflection.EndAlpha(),
		)
	}
	if reflection.StartPosition() != 0 {
		t.Errorf(
			"expected start position 0, got %d",
			reflection.StartPosition(),
		)
	}
	if reflection.EndPosition() != 100000 {
		t.Errorf(
			"expected end position 100000, got %d",
			reflection.EndPosition(),
		)
	}
	if reflection.Direction() != 5400000 {
		t.Errorf(
			"expected direction 5400000, got %d",
			reflection.Direction(),
		)
	}
}

func TestBlur(t *testing.T) {
	blur := NewBlur(60000)

	if blur.Radius() != 60000 {
		t.Errorf(
			"expected radius 60000, got %d",
			blur.Radius(),
		)
	}

	// Test grow
	if !blur.Grow() {
		t.Error(
			"expected Grow to be true by default",
		)
	}
	blur.SetGrow(false)
	if blur.Grow() {
		t.Error("expected Grow to be false")
	}
	blur.SetGrow(true)
	if !blur.Grow() {
		t.Error("expected Grow to be true")
	}
}

func TestFillOverlay(t *testing.T) {
	overlay := NewFillOverlay(BlendModeMultiply)

	if overlay.Blend() != BlendModeMultiply {
		t.Errorf(
			"expected blend mode '%s', got '%s'",
			BlendModeMultiply,
			overlay.Blend(),
		)
	}

	// Test solid fill
	solidFill := NewSolidFillWithRgb("00AAFF")
	overlay.SetSolidFill(solidFill)
	if overlay.SolidFill() == nil {
		t.Error("expected SolidFill to be set")
	}

	// Test no fill
	overlay.SetNoFill()
	if overlay.SolidFill() != nil {
		t.Error(
			"expected SolidFill to be nil after SetNoFill",
		)
	}
}

func TestEffectContainer(t *testing.T) {
	container := NewEffectContainer()

	if container.LocalName() != "effectDag" {
		t.Errorf(
			"expected local name 'effectDag', got '%s'",
			container.LocalName(),
		)
	}

	// Test type
	container.SetType("tree")
	if container.Type() != "tree" {
		t.Errorf(
			"expected type 'tree', got '%s'",
			container.Type(),
		)
	}

	// Test name
	container.SetName("myEffect")
	if container.Name() != "myEffect" {
		t.Errorf(
			"expected name 'myEffect', got '%s'",
			container.Name(),
		)
	}

	// Test empty name
	container.SetName("")
	if container.Name() != "" {
		t.Errorf(
			"expected empty name, got '%s'",
			container.Name(),
		)
	}
}

func TestEffectListClone(t *testing.T) {
	effectList := NewEffectList()
	effectList.SetBlur(NewBlur(50000))
	effectList.SetGlow(NewGlowWithRadius(100000))

	cloned := effectList.Clone()
	clone, ok := cloned.(*EffectList)
	if !ok {
		t.Fatal(
			"Clone did not return *EffectList",
		)
	}

	if clone == effectList {
		t.Error(
			"clone should not be the same instance",
		)
	}
	if clone.Blur() == nil {
		t.Error(
			"cloned effect list should have blur",
		)
	}
	if clone.Blur().Radius() != 50000 {
		t.Errorf(
			"expected blur radius 50000, got %d",
			clone.Blur().Radius(),
		)
	}
	if clone.Glow() == nil {
		t.Error(
			"cloned effect list should have glow",
		)
	}
}

func TestOuterShadowClone(t *testing.T) {
	shadow := NewDropShadow(
		50000,
		100000,
		5400000,
		"FF0000",
	)
	shadow.SetAlignment(RectAlignCenter)

	cloned := shadow.Clone()
	clone, ok := cloned.(*OuterShadow)
	if !ok {
		t.Fatal(
			"Clone did not return *OuterShadow",
		)
	}

	if clone == shadow {
		t.Error(
			"clone should not be the same instance",
		)
	}
	if clone.BlurRadius() != 50000 {
		t.Errorf(
			"expected blur radius 50000, got %d",
			clone.BlurRadius(),
		)
	}
	if clone.Distance() != 100000 {
		t.Errorf(
			"expected distance 100000, got %d",
			clone.Distance(),
		)
	}
	if clone.Direction() != 5400000 {
		t.Errorf(
			"expected direction 5400000, got %d",
			clone.Direction(),
		)
	}
	if clone.Alignment() != RectAlignCenter {
		t.Errorf(
			"expected alignment '%s', got '%s'",
			RectAlignCenter,
			clone.Alignment(),
		)
	}
}

func TestPresetShadowValues(t *testing.T) {
	tests := []struct {
		value PresetShadowValue
		xml   string
	}{
		{PresetShadowTopLeftDropShadow, "shdw1"},
		{PresetShadowTopRightDropShadow, "shdw2"},
		{
			PresetShadowBackLeftPerspective,
			"shdw3",
		},
		{
			PresetShadowBackRightPerspective,
			"shdw4",
		},
		{
			PresetShadowBottomLeftDropShadow,
			"shdw5",
		},
		{
			PresetShadowBottomRightDropShadow,
			"shdw6",
		},
		{
			PresetShadowFrontLeftPerspective,
			"shdw7",
		},
		{
			PresetShadowFrontRightPerspective,
			"shdw8",
		},
		{
			PresetShadowTopLeftSmallDropShadow,
			"shdw9",
		},
		{
			PresetShadowTopLeftLargeDropShadow,
			"shdw10",
		},
		{
			PresetShadowBackLeftLongPerspective,
			"shdw11",
		},
		{
			PresetShadowBackRightLongPerspective,
			"shdw12",
		},
		{
			PresetShadowTopLeftDoubleDropShadow,
			"shdw13",
		},
		{
			PresetShadowBottomRightSmallDropShadow,
			"shdw14",
		},
		{
			PresetShadowFrontLeftLongPerspective,
			"shdw15",
		},
		{
			PresetShadowFrontRightLongPerspective,
			"shdw16",
		},
		{PresetShadow3DOuterBox, "shdw17"},
		{PresetShadow3DInnerBox, "shdw18"},
		{
			PresetShadowBackCenterPerspective,
			"shdw19",
		},
		{PresetShadowFrontBottom, "shdw20"},
	}

	for _, tt := range tests {
		if string(tt.value) != tt.xml {
			t.Errorf(
				"expected value '%s', got '%s'",
				tt.xml,
				string(tt.value),
			)
		}
	}
}

func TestBlendModeValues(t *testing.T) {
	tests := []struct {
		value BlendModeValue
		xml   string
	}{
		{BlendModeOverlay, "over"},
		{BlendModeMultiply, "mult"},
		{BlendModeScreen, "screen"},
		{BlendModeDarken, "darken"},
		{BlendModeLighten, "lighten"},
	}

	for _, tt := range tests {
		if string(tt.value) != tt.xml {
			t.Errorf(
				"expected value '%s', got '%s'",
				tt.xml,
				string(tt.value),
			)
		}
	}
}

func TestEffectXmlOutput(t *testing.T) {
	effectList := NewEffectList()

	// Add various effects
	effectList.SetBlur(NewBlur(50000))
	glow := NewGlowWithParams(100000, "FF0000")
	effectList.SetGlow(glow)
	shadow := NewDropShadow(
		30000,
		50000,
		5400000,
		"000000",
	)
	effectList.SetOuterShadow(shadow)

	xml := effectList.OuterXml()

	// Check that the XML contains expected elements
	if !strings.Contains(xml, "effectLst") {
		t.Error(
			"expected XML to contain 'effectLst'",
		)
	}
	if !strings.Contains(xml, "blur") {
		t.Error("expected XML to contain 'blur'")
	}
	if !strings.Contains(xml, "glow") {
		t.Error("expected XML to contain 'glow'")
	}
	if !strings.Contains(xml, "outerShdw") {
		t.Error(
			"expected XML to contain 'outerShdw'",
		)
	}
	if !strings.Contains(xml, "rad=\"100000\"") {
		t.Error(
			"expected XML to contain glow radius",
		)
	}
	if !strings.Contains(xml, "srgbClr") {
		t.Error(
			"expected XML to contain color element",
		)
	}
}

func TestSoftEdgeXmlOutput(t *testing.T) {
	softEdge := NewSoftEdge(120000)
	xml := softEdge.OuterXml()

	if !strings.Contains(xml, "softEdge") {
		t.Error(
			"expected XML to contain 'softEdge'",
		)
	}
	if !strings.Contains(xml, "rad=\"120000\"") {
		t.Error(
			"expected XML to contain radius attribute",
		)
	}
}

func TestReflectionXmlOutput(t *testing.T) {
	reflection := NewReflection()
	reflection.SetBlurRadius(50000)
	reflection.SetStartOpacity(80000)
	reflection.SetEndAlpha(10000)
	reflection.SetDirection(5400000)

	xml := reflection.OuterXml()

	if !strings.Contains(xml, "reflection") {
		t.Error(
			"expected XML to contain 'reflection'",
		)
	}
	if !strings.Contains(
		xml,
		"blurRad=\"50000\"",
	) {
		t.Error(
			"expected XML to contain blur radius",
		)
	}
	if !strings.Contains(xml, "stA=\"80000\"") {
		t.Error(
			"expected XML to contain start opacity",
		)
	}
	if !strings.Contains(xml, "endA=\"10000\"") {
		t.Error(
			"expected XML to contain end alpha",
		)
	}
	if !strings.Contains(xml, "dir=\"5400000\"") {
		t.Error(
			"expected XML to contain direction",
		)
	}
}
