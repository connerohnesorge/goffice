package diagram_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml/types"
)

const (
	testNode1ID = "node1"
	testConn1ID = "conn1"
)

// TestNewDataModelRoot verifies that DataModelRoot can be created.
func TestNewDataModelRoot(t *testing.T) {
	dm := diagram.NewDataModelRoot()
	if dm == nil {
		t.Fatal("NewDataModelRoot returned nil")
	}
}

// TestNewPoint verifies that Point can be created.
func TestNewPoint(t *testing.T) {
	pt := diagram.NewPoint()
	if pt == nil {
		t.Fatal("NewPoint returned nil")
	}
}

// TestNewConnection verifies that Connection can be created.
func TestNewConnection(t *testing.T) {
	conn := diagram.NewConnection()
	if conn == nil {
		t.Fatal("NewConnection returned nil")
	}
}

// TestNewLayoutDefinition verifies that LayoutDefinition can be created.
func TestNewLayoutDefinition(t *testing.T) {
	layout := diagram.NewLayoutDefinition()
	if layout == nil {
		t.Fatal(
			"NewLayoutDefinition returned nil",
		)
	}
}

// TestNewStyleDefinition verifies that StyleDefinition can be created.
func TestNewStyleDefinition(t *testing.T) {
	style := diagram.NewStyleDefinition()
	if style == nil {
		t.Fatal("NewStyleDefinition returned nil")
	}
}

// TestNewColorsDefinition verifies that ColorsDefinition can be created.
func TestNewColorsDefinition(t *testing.T) {
	colors := diagram.NewColorsDefinition()
	if colors == nil {
		t.Fatal(
			"NewColorsDefinition returned nil",
		)
	}
}

// TestDataModelStructure verifies the basic structure of DataModelRoot.
func TestDataModelStructure(t *testing.T) {
	dm := diagram.NewDataModelRoot()

	// Create a point list
	pointList := diagram.NewPointList()
	dm.PointList = pointList

	// Create a connection list
	connList := diagram.NewConnectionList()
	dm.ConnectionList = connList

	// Verify the structure
	if dm.PointList == nil {
		t.Error(
			"PointList should not be nil after assignment",
		)
	}
	if dm.ConnectionList == nil {
		t.Error(
			"ConnectionList should not be nil after assignment",
		)
	}
}

// TestPoint_Properties tests Point property setting and getting.
func TestPoint_Properties(t *testing.T) {
	// Create a point with properties
	pt := diagram.NewPoint()
	pt.ModelId = types.NewStringValue(testNode1ID)
	pt.Type = types.NewEnumValue(
		diagram.PointValuesNode,
	)

	// Verify properties
	if pt.ModelId == nil ||
		pt.ModelId.Value() != testNode1ID {
		t.Errorf(
			"ModelId mismatch: got %v, want %q",
			pt.ModelId,
			testNode1ID,
		)
	}
	if pt.Type == nil ||
		pt.Type.Value() != diagram.PointValuesNode {
		t.Errorf(
			"Type mismatch: got %v, want %v",
			pt.Type,
			diagram.PointValuesNode,
		)
	}
}

// TestConnection_Properties tests Connection property setting and getting.
func TestConnection_Properties(t *testing.T) {
	// Create a connection with properties
	conn := diagram.NewConnection()
	conn.ModelId = types.NewStringValue(
		testConn1ID,
	)
	conn.Type = types.NewEnumValue(
		diagram.ConnectionValuesParof,
	)
	conn.SourceId = types.NewStringValue(
		testNode1ID,
	)
	conn.DestinationId = types.NewStringValue(
		"node2",
	)

	// Verify properties
	if conn.ModelId == nil ||
		conn.ModelId.Value() != testConn1ID {
		t.Errorf(
			"ModelId mismatch: got %v, want %q",
			conn.ModelId,
			testConn1ID,
		)
	}
	if conn.Type == nil ||
		conn.Type.Value() != diagram.ConnectionValuesParof {
		t.Errorf(
			"Type mismatch: got %v, want %v",
			conn.Type,
			diagram.ConnectionValuesParof,
		)
	}
	if conn.SourceId == nil ||
		conn.SourceId.Value() != testNode1ID {
		t.Errorf(
			"SourceId mismatch: got %v, want %q",
			conn.SourceId,
			testNode1ID,
		)
	}
	if conn.DestinationId == nil ||
		conn.DestinationId.Value() != "node2" {
		t.Errorf(
			"DestinationId mismatch: got %v, want %q",
			conn.DestinationId,
			"node2",
		)
	}
}

// TestDataModelRoot_Structure tests DataModelRoot structure.
func TestDataModelRoot_Structure(t *testing.T) {
	// Create a data model with points and connections
	dm := diagram.NewDataModelRoot()

	// Add point list with one point
	pointList := diagram.NewPointList()
	pt := diagram.NewPoint()
	pt.ModelId = types.NewStringValue(testNode1ID)
	pt.Type = types.NewEnumValue(
		diagram.PointValuesNode,
	)
	pointList.Point = pt
	dm.PointList = pointList

	// Add connection list with one connection
	connList := diagram.NewConnectionList()
	conn := diagram.NewConnection()
	conn.ModelId = types.NewStringValue(
		testConn1ID,
	)
	conn.Type = types.NewEnumValue(
		diagram.ConnectionValuesParof,
	)
	conn.SourceId = types.NewStringValue(
		testNode1ID,
	)
	conn.DestinationId = types.NewStringValue(
		"node2",
	)
	connList.Connection = conn
	dm.ConnectionList = connList

	// Verify structure
	if dm.PointList == nil {
		t.Fatal("PointList is nil")
	}
	if dm.PointList.Point == nil {
		t.Fatal("Expected 1 point, got nil")
	}
	if dm.PointList.Point.ModelId == nil ||
		dm.PointList.Point.ModelId.Value() != testNode1ID {
		t.Errorf(
			"Point ModelId mismatch: got %v, want %q",
			dm.PointList.Point.ModelId,
			testNode1ID,
		)
	}

	if dm.ConnectionList == nil {
		t.Fatal("ConnectionList is nil")
	}
	if dm.ConnectionList.Connection == nil {
		t.Fatal("Expected 1 connection, got nil")
	}
	if dm.ConnectionList.Connection.ModelId == nil ||
		dm.ConnectionList.Connection.ModelId.Value() != testConn1ID {
		t.Errorf(
			"Connection ModelId mismatch: got %v, want %q",
			dm.ConnectionList.Connection.ModelId,
			testConn1ID,
		)
	}
}

// TestLayoutDefinition_Creation tests LayoutDefinition creation.
func TestLayoutDefinition_Creation(t *testing.T) {
	layout := diagram.NewLayoutDefinition()
	if layout == nil {
		t.Fatal(
			"NewLayoutDefinition returned nil",
		)
	}

	// Verify validation works
	err := layout.Validate()
	if err != nil {
		t.Errorf(
			"LayoutDefinition.Validate() failed: %v",
			err,
		)
	}
}

// TestStyleDefinition_Creation tests StyleDefinition creation.
func TestStyleDefinition_Creation(t *testing.T) {
	style := diagram.NewStyleDefinition()
	if style == nil {
		t.Fatal("NewStyleDefinition returned nil")
	}

	// Verify validation works
	err := style.Validate()
	if err != nil {
		t.Errorf(
			"StyleDefinition.Validate() failed: %v",
			err,
		)
	}
}

// TestColorsDefinition_Creation tests ColorsDefinition creation.
func TestColorsDefinition_Creation(t *testing.T) {
	colors := diagram.NewColorsDefinition()
	if colors == nil {
		t.Fatal(
			"NewColorsDefinition returned nil",
		)
	}

	// Verify validation works
	err := colors.Validate()
	if err != nil {
		t.Errorf(
			"ColorsDefinition.Validate() failed: %v",
			err,
		)
	}
}

// TestPoint_Validation tests Point validation.
func TestPoint_Validation(t *testing.T) {
	pt := diagram.NewPoint()
	pt.ModelId = types.NewStringValue(testNode1ID)

	err := pt.Validate()
	if err != nil {
		t.Errorf(
			"Point.Validate() failed: %v",
			err,
		)
	}
}

// TestConnection_Validation tests Connection validation.
func TestConnection_Validation(t *testing.T) {
	conn := diagram.NewConnection()
	conn.ModelId = types.NewStringValue(
		testConn1ID,
	)
	conn.SourceId = types.NewStringValue(
		testNode1ID,
	)
	conn.DestinationId = types.NewStringValue(
		"node2",
	)

	err := conn.Validate()
	if err != nil {
		t.Errorf(
			"Connection.Validate() failed: %v",
			err,
		)
	}
}

// TestDataModelRoot_Validation tests DataModelRoot validation.
func TestDataModelRoot_Validation(t *testing.T) {
	dm := diagram.NewDataModelRoot()

	// Add point list
	pointList := diagram.NewPointList()
	pt := diagram.NewPoint()
	pt.ModelId = types.NewStringValue(testNode1ID)
	pointList.Point = pt
	dm.PointList = pointList

	err := dm.Validate()
	if err != nil {
		t.Errorf(
			"DataModelRoot.Validate() failed: %v",
			err,
		)
	}
}

// TestPoint_Clone tests Point cloning.
func TestPoint_Clone(t *testing.T) {
	pt := diagram.NewPoint()
	pt.ModelId = types.NewStringValue(testNode1ID)
	pt.Type = types.NewEnumValue(
		diagram.PointValuesNode,
	)

	cloned := pt.Clone()
	pt2, ok := cloned.(*diagram.Point)
	if !ok {
		t.Fatal(
			"Clone() did not return *diagram.Point",
		)
	}
	if pt2.ModelId == nil ||
		pt2.ModelId.Value() != testNode1ID {
		t.Errorf(
			"Clone ModelId mismatch: got %v, want %q",
			pt2.ModelId,
			testNode1ID,
		)
	}
	if pt2.Type == nil ||
		pt2.Type.Value() != diagram.PointValuesNode {
		t.Errorf(
			"Clone Type mismatch: got %v, want %v",
			pt2.Type,
			diagram.PointValuesNode,
		)
	}

	// Verify it's a deep copy
	pt.ModelId = types.NewStringValue("node2")
	if pt2.ModelId.Value() != testNode1ID {
		t.Error(
			"Clone should be independent of original",
		)
	}
}

// TestConnection_Clone tests Connection cloning.
func TestConnection_Clone(t *testing.T) {
	conn := diagram.NewConnection()
	conn.ModelId = types.NewStringValue(
		testConn1ID,
	)
	conn.SourceId = types.NewStringValue(
		testNode1ID,
	)
	conn.DestinationId = types.NewStringValue(
		"node2",
	)

	cloned := conn.Clone()
	conn2, ok := cloned.(*diagram.Connection)
	if !ok {
		t.Fatal(
			"Clone() did not return *diagram.Connection",
		)
	}
	if conn2.ModelId == nil ||
		conn2.ModelId.Value() != testConn1ID {
		t.Errorf(
			"Clone ModelId mismatch: got %v, want %q",
			conn2.ModelId,
			testConn1ID,
		)
	}

	// Verify it's a deep copy
	conn.ModelId = types.NewStringValue("conn2")
	if conn2.ModelId.Value() != testConn1ID {
		t.Error(
			"Clone should be independent of original",
		)
	}
}
