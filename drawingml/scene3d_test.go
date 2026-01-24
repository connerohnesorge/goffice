package drawingml

import (
	"strings"
	"testing"
)

func TestScene3D(t *testing.T) {
	t.Run("NewScene3D", func(t *testing.T) {
		scene := NewScene3D()
		if !strings.Contains(scene.OuterXml(), "scene3d") {
			t.Error("Expected scene3d element")
		}
	})

	t.Run("Camera", func(t *testing.T) {
		scene := NewScene3D()
		cam := NewCamera()
		cam.SetPreset("orthographicFront")
		cam.SetFieldOfView(5400000)
		cam.SetZoom(100000)
		
		rot := NewRotation3D()
		rot.SetLat(0)
		rot.SetLon(0)
		rot.SetRev(0)
		cam.SetRotation(rot)

		scene.SetCamera(cam)

		if scene.Camera() == nil {
			t.Error("Expected camera")
		}
		if scene.Camera().Preset() != "orthographicFront" {
			t.Errorf("Expected orthographicFront, got %s", scene.Camera().Preset())
		}
	})

	t.Run("LightRig", func(t *testing.T) {
		scene := NewScene3D()
		rig := NewLightRig()
		rig.SetRig("threePt")
		rig.SetDirection("t")
		scene.SetLightRig(rig)

		if scene.LightRig() == nil {
			t.Error("Expected light rig")
		}
		if scene.LightRig().Rig() != "threePt" {
			t.Errorf("Expected threePt, got %s", scene.LightRig().Rig())
		}
	})
}

func TestShape3D(t *testing.T) {
	t.Run("NewShape3D", func(t *testing.T) {
		sp3d := NewShape3D()
		if !strings.Contains(sp3d.OuterXml(), "sp3d") {
			t.Error("Expected sp3d element")
		}
	})

	t.Run("Properties", func(t *testing.T) {
		sp3d := NewShape3D()
		sp3d.SetZ(1000)
		sp3d.SetExtrusionHeight(2000)
		sp3d.SetContourWidth(500)
		sp3d.SetPresetMaterial("plastic")

		if sp3d.Z() != 1000 {
			t.Errorf("Expected Z 1000, got %d", sp3d.Z())
		}
		if sp3d.ExtrusionHeight() != 2000 {
			t.Errorf("Expected ExtrusionHeight 2000, got %d", sp3d.ExtrusionHeight())
		}
		if sp3d.PresetMaterial() != "plastic" {
			t.Errorf("Expected plastic, got %s", sp3d.PresetMaterial())
		}
	})

	t.Run("Bevels", func(t *testing.T) {
		sp3d := NewShape3D()
		
		top := NewBevel("bevelT")
		top.SetWidth(76200)
		top.SetHeight(76200)
		top.SetPreset("circle")
		sp3d.SetBevelTop(top)

		bottom := NewBevel("bevelB")
		bottom.SetPreset("slope")
		sp3d.SetBevelBottom(bottom)

		if sp3d.BevelTop() == nil {
			t.Error("Expected top bevel")
		}
		if sp3d.BevelTop().Preset() != "circle" {
			t.Errorf("Expected circle bevel, got %s", sp3d.BevelTop().Preset())
		}

		if sp3d.BevelBottom() == nil {
			t.Error("Expected bottom bevel")
		}
		if sp3d.BevelBottom().Preset() != "slope" {
			t.Errorf("Expected slope bevel, got %s", sp3d.BevelBottom().Preset())
		}
	})

	t.Run("Colors", func(t *testing.T) {
		sp3d := NewShape3D()
		
		// Test RGB extrusion color
		sp3d.SetExtrusionColorRgb("FF0000")
		if sp3d.ExtrusionColor() == nil {
			t.Error("Expected extrusion color")
		}
		
		// Test scheme extrusion color
		sp3d.SetExtrusionColorScheme(SchemeColorAccent1)
		if sp3d.ExtrusionColor() == nil {
			t.Error("Expected extrusion color")
		}
		
		// Test RGB contour color
		sp3d.SetContourColorRgb("0000FF")
		if sp3d.ContourColor() == nil {
			t.Error("Expected contour color")
		}
		
		// Test scheme contour color
		sp3d.SetContourColorScheme(SchemeColorAccent2)
		if sp3d.ContourColor() == nil {
			t.Error("Expected contour color")
		}
	})
}
