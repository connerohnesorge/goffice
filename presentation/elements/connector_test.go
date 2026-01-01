package elements

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/routing"
)

func TestConnectionShape_StartConnection(t *testing.T) {
	t.Run("no connection initially", func(t *testing.T) {
		connector := NewConnectionShape()

		_, ok := connector.StartConnection()
		if ok {
			t.Error("Expected no start connection initially, but got one")
		}
	})

	t.Run("set and get start connection", func(t *testing.T) {
		connector := NewConnectionShape()

		connector.SetStartConnection("shape123", 2)

		conn, ok := connector.StartConnection()
		if !ok {
			t.Fatal("Expected start connection to be set")
		}

		if conn.ShapeID != "shape123" {
			t.Errorf("Expected ShapeID 'shape123', got '%s'", conn.ShapeID)
		}

		if conn.ConnectionSiteIndex != 2 {
			t.Errorf("Expected ConnectionSiteIndex 2, got %d", conn.ConnectionSiteIndex)
		}
	})

	t.Run("clear start connection", func(t *testing.T) {
		connector := NewConnectionShape()

		connector.SetStartConnection("shape123", 2)

		_, ok := connector.StartConnection()
		if !ok {
			t.Fatal("Expected start connection to be set")
		}

		connector.ClearStartConnection()

		_, ok = connector.StartConnection()
		if ok {
			t.Error("Expected start connection to be cleared")
		}
	})
}

func TestConnectionShape_EndConnection(t *testing.T) {
	t.Run("no connection initially", func(t *testing.T) {
		connector := NewConnectionShape()

		_, ok := connector.EndConnection()
		if ok {
			t.Error("Expected no end connection initially, but got one")
		}
	})

	t.Run("set and get end connection", func(t *testing.T) {
		connector := NewConnectionShape()

		connector.SetEndConnection("shape456", 3)

		conn, ok := connector.EndConnection()
		if !ok {
			t.Fatal("Expected end connection to be set")
		}

		if conn.ShapeID != "shape456" {
			t.Errorf("Expected ShapeID 'shape456', got '%s'", conn.ShapeID)
		}

		if conn.ConnectionSiteIndex != 3 {
			t.Errorf("Expected ConnectionSiteIndex 3, got %d", conn.ConnectionSiteIndex)
		}
	})

	t.Run("clear end connection", func(t *testing.T) {
		connector := NewConnectionShape()

		connector.SetEndConnection("shape456", 3)

		_, ok := connector.EndConnection()
		if !ok {
			t.Fatal("Expected end connection to be set")
		}

		connector.ClearEndConnection()

		_, ok = connector.EndConnection()
		if ok {
			t.Error("Expected end connection to be cleared")
		}
	})
}

func TestConnectionShape_BothConnections(t *testing.T) {
	connector := NewConnectionShape()

	// Set both connections
	connector.SetStartConnection("shape1", 0)
	connector.SetEndConnection("shape2", 2)

	// Verify start connection
	startConn, ok := connector.StartConnection()
	if !ok {
		t.Fatal("Expected start connection to be set")
	}
	if startConn.ShapeID != "shape1" || startConn.ConnectionSiteIndex != 0 {
		t.Errorf("Start connection mismatch: got %+v", startConn)
	}

	// Verify end connection
	endConn, ok := connector.EndConnection()
	if !ok {
		t.Fatal("Expected end connection to be set")
	}
	if endConn.ShapeID != "shape2" || endConn.ConnectionSiteIndex != 2 {
		t.Errorf("End connection mismatch: got %+v", endConn)
	}

	// Clear start, verify end remains
	connector.ClearStartConnection()
	_, ok = connector.StartConnection()
	if ok {
		t.Error("Expected start connection to be cleared")
	}
	_, ok = connector.EndConnection()
	if !ok {
		t.Error("Expected end connection to remain after clearing start")
	}
}

func TestGenerateDefaultConnectionSites(t *testing.T) {
	width := drawingml.EMU(100000)  // 100000 EMU
	height := drawingml.EMU(50000)  // 50000 EMU

	sites := routing.GenerateDefaultConnectionSites(width, height)

	if len(sites) != 4 {
		t.Fatalf("Expected 4 default connection sites, got %d", len(sites))
	}

	// Test site 0 (top center)
	if sites[0].Index != 0 {
		t.Errorf("Site 0: expected index 0, got %d", sites[0].Index)
	}
	if sites[0].X != width/2 {
		t.Errorf("Site 0: expected X=%d, got %d", width/2, sites[0].X)
	}
	if sites[0].Y != 0 {
		t.Errorf("Site 0: expected Y=0, got %d", sites[0].Y)
	}
	if sites[0].Angle != 270*60000 {
		t.Errorf("Site 0: expected angle 270*60000 (up), got %d", sites[0].Angle)
	}

	// Test site 1 (right center)
	if sites[1].Index != 1 {
		t.Errorf("Site 1: expected index 1, got %d", sites[1].Index)
	}
	if sites[1].X != width {
		t.Errorf("Site 1: expected X=%d, got %d", width, sites[1].X)
	}
	if sites[1].Y != height/2 {
		t.Errorf("Site 1: expected Y=%d, got %d", height/2, sites[1].Y)
	}
	if sites[1].Angle != 0 {
		t.Errorf("Site 1: expected angle 0 (right), got %d", sites[1].Angle)
	}

	// Test site 2 (bottom center)
	if sites[2].Index != 2 {
		t.Errorf("Site 2: expected index 2, got %d", sites[2].Index)
	}
	if sites[2].X != width/2 {
		t.Errorf("Site 2: expected X=%d, got %d", width/2, sites[2].X)
	}
	if sites[2].Y != height {
		t.Errorf("Site 2: expected Y=%d, got %d", height, sites[2].Y)
	}
	if sites[2].Angle != 90*60000 {
		t.Errorf("Site 2: expected angle 90*60000 (down), got %d", sites[2].Angle)
	}

	// Test site 3 (left center)
	if sites[3].Index != 3 {
		t.Errorf("Site 3: expected index 3, got %d", sites[3].Index)
	}
	if sites[3].X != 0 {
		t.Errorf("Site 3: expected X=0, got %d", sites[3].X)
	}
	if sites[3].Y != height/2 {
		t.Errorf("Site 3: expected Y=%d, got %d", height/2, sites[3].Y)
	}
	if sites[3].Angle != 180*60000 {
		t.Errorf("Site 3: expected angle 180*60000 (left), got %d", sites[3].Angle)
	}
}
