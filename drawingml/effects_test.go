package drawingml

import (
	"fmt"
	"strings"
	"testing"
)

// TestConstructorsBasic verifies that all effect constructors return valid objects
// with expected basic properties. This consolidates simple constructor tests
// that previously only checked for nil returns.
func TestConstructorsBasic(t *testing.T) {
	tests := []struct {
		name      string
		obj       any
		localName string
		namespace string
	}{
		{
			name:      "EffectList",
			obj:       NewEffectList(),
			localName: "effectLst",
			namespace: NamespaceMain,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.obj == nil {
				t.Fatalf("%s constructor returned nil", tt.name)
			}

			// Check if the object has the expected methods
			type localNamer interface {
				LocalName() string
			}
			type namespaceURIer interface {
				NamespaceURI() string
			}

			if ln, ok := tt.obj.(localNamer); ok {
				if got := ln.LocalName(); got != tt.localName {
					t.Errorf("expected local name '%s', got '%s'", tt.localName, got)
				}
			}

			nu, ok := tt.obj.(namespaceURIer)
			if !ok {
				return
			}
			if got := nu.NamespaceURI(); got != tt.namespace {
				t.Errorf("expected namespace '%s', got '%s'", tt.namespace, got)
			}
		})
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

func TestSoftEdgeXmlOutput(t *testing.T) {
	softEdge := NewSoftEdge(120000)
	xml := softEdge.OuterXml()

	// Parse XML to verify structure
	if !strings.HasPrefix(xml, "<a:softEdge") {
		t.Error("expected XML to start with <a:softEdge")
	}
	if !strings.HasSuffix(xml, "/>") {
		t.Error("expected XML to be self-closing")
	}
	// Verify the radius attribute is properly formatted
	if !strings.Contains(xml, `rad="120000"`) {
		t.Error("expected XML to contain radius attribute with value 120000")
	}
}

func TestReflectionXmlOutput(t *testing.T) {
	reflection := NewReflection()
	reflection.SetBlurRadius(50000)
	reflection.SetStartOpacity(80000)
	reflection.SetEndAlpha(10000)
	reflection.SetDirection(5400000)

	xml := reflection.OuterXml()

	// Verify XML structure - should be self-contained reflection element
	if !strings.HasPrefix(xml, "<a:reflection") {
		t.Error("expected XML to start with <a:reflection")
	}
	if !strings.HasSuffix(xml, "/>") {
		t.Error("expected XML to be self-closing")
	}

	// Verify all attributes are present with correct values
	attrs := []struct {
		name  string
		value string
	}{
		{"blurRad", "50000"},
		{"stA", "80000"},
		{"endA", "10000"},
		{"dir", "5400000"},
	}

	for _, attr := range attrs {
		expected := fmt.Sprintf(`%s=%q`, attr.name, attr.value)
		if !strings.Contains(xml, expected) {
			t.Errorf("expected XML to contain %s attribute with value %s", attr.name, attr.value)
		}
	}
}
