package diagram_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/types"
)

const (
	testNode1ID = "node1"
	testConn1ID = "conn1"
)

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

// TestDataModelBuilder_NewDataModelBuilder tests builder creation.
func TestDataModelBuilder_NewDataModelBuilder(t *testing.T) {
	builder := diagram.NewDataModelBuilder()
	if builder == nil {
		t.Fatal("NewDataModelBuilder returned nil")
	}
	if builder.PointCount() != 0 {
		t.Errorf(
			"Expected 0 points, got %d",
			builder.PointCount(),
		)
	}
	if builder.ConnectionCount() != 0 {
		t.Errorf(
			"Expected 0 connections, got %d",
			builder.ConnectionCount(),
		)
	}
}

// TestDataModelBuilder_AddPoint tests adding points.
func TestDataModelBuilder_AddPoint(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	// Add single point
	builder.AddPoint("node1")
	if builder.PointCount() != 1 {
		t.Errorf(
			"Expected 1 point, got %d",
			builder.PointCount(),
		)
	}
	if !builder.HasPoints() {
		t.Error("HasPoints should return true")
	}

	// Add second point
	builder.AddPoint("node2")
	if builder.PointCount() != 2 {
		t.Errorf(
			"Expected 2 points, got %d",
			builder.PointCount(),
		)
	}

	// Build and verify
	dataModel := builder.Build()
	if dataModel == nil {
		t.Fatal("Build returned nil")
	}
	if dataModel.PointList == nil {
		t.Fatal("PointList should not be nil")
	}
}

// TestDataModelBuilder_AddPointWithType tests adding points with types.
func TestDataModelBuilder_AddPointWithType(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	// Add node
	builder.AddPoint(
		"node1",
		diagram.PointValuesNode,
	)

	// Add assistant
	builder.AddPoint(
		"asst1",
		diagram.PointValuesAsst,
	)

	dataModel := builder.Build()
	if dataModel.PointList == nil {
		t.Fatal("PointList should not be nil")
	}

	// Verify points have correct types
	pointCount := 0
	for pt := range openxml.Elements[*diagram.Point](
		dataModel.PointList,
	) {
		pointCount++
		switch pointCount {
		case 1:
			if pt.Type == nil ||
				pt.Type.Value() != diagram.PointValuesNode {
				t.Errorf(
					"First point should be node, got %v",
					pt.Type,
				)
			}
		case 2:
			if pt.Type == nil ||
				pt.Type.Value() != diagram.PointValuesAsst {
				t.Errorf(
					"Second point should be assistant, got %v",
					pt.Type,
				)
			}
		}
	}

	if pointCount != 2 {
		t.Errorf(
			"Expected 2 points, got %d",
			pointCount,
		)
	}
}

// TestDataModelBuilder_AddConnection tests adding connections.
func TestDataModelBuilder_AddConnection(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	// Add points first
	builder.AddPoint("node1")
	builder.AddPoint("node2")

	// Add connection
	builder.AddConnection("node1", "node2")

	if builder.ConnectionCount() != 1 {
		t.Errorf(
			"Expected 1 connection, got %d",
			builder.ConnectionCount(),
		)
	}
	if !builder.HasConnections() {
		t.Error("HasConnections should return true")
	}

	dataModel := builder.Build()
	if dataModel == nil {
		t.Fatal("Build returned nil")
	}
	if dataModel.ConnectionList == nil {
		t.Fatal("ConnectionList should not be nil")
	}
}

// TestDataModelBuilder_AddConnectionWithType tests adding connections with types.
func TestDataModelBuilder_AddConnectionWithType(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	builder.AddPoint("parent")
	builder.AddPoint("child")

	// Add parent-of connection
	builder.AddConnection(
		"parent",
		"child",
		diagram.ConnectionValuesParof,
	)

	dataModel := builder.Build()

	connCount := 0
	for conn := range openxml.Elements[*diagram.Connection](
		dataModel.ConnectionList,
	) {
		connCount++
		if conn.Type == nil ||
			conn.Type.Value() != diagram.ConnectionValuesParof {
			t.Errorf(
				"Connection should be parOf, got %v",
				conn.Type,
			)
		}
		if conn.SourceId == nil ||
			conn.SourceId.Value() != "parent" {
			t.Errorf(
				"Source should be parent, got %v",
				conn.SourceId,
			)
		}
		if conn.DestinationId == nil ||
			conn.DestinationId.Value() != "child" {
			t.Errorf(
				"Destination should be child, got %v",
				conn.DestinationId,
			)
		}
	}

	if connCount != 1 {
		t.Errorf(
			"Expected 1 connection, got %d",
			connCount,
		)
	}
}

// TestDataModelBuilder_ConvenienceMethods tests convenience methods.
func TestDataModelBuilder_ConvenienceMethods(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	// Test AddAssistantPoint
	builder.AddAssistantPoint("asst1")
	dataModel := builder.Build()

	ptCount := 0
	for pt := range openxml.Elements[*diagram.Point](
		dataModel.PointList,
	) {
		ptCount++
		if pt.Type == nil ||
			pt.Type.Value() != diagram.PointValuesAsst {
			t.Error("AddAssistantPoint should create assistant type")
		}
	}
	if ptCount != 1 {
		t.Errorf(
			"Expected 1 point from AddAssistantPoint, got %d",
			ptCount,
		)
	}

	// Test AddDocumentPoint
	builder.AddDocumentPoint("doc1")
	dataModel = builder.Build()

	ptCount = 0
	for pt := range openxml.Elements[*diagram.Point](
		dataModel.PointList,
	) {
		ptCount++
		if pt.Type == nil ||
			pt.Type.Value() != diagram.PointValuesDoc {
			t.Error("AddDocumentPoint should create doc type")
		}
	}
	if ptCount != 1 {
		t.Errorf(
			"Expected 1 point from AddDocumentPoint, got %d",
			ptCount,
		)
	}

	// Test AddParentOfConnection
	builder.AddPoint("p1")
	builder.AddPoint("c1")
	builder.AddParentOfConnection("p1", "c1")
	dataModel = builder.Build()

	connCount := 0
	for conn := range openxml.Elements[*diagram.Connection](
		dataModel.ConnectionList,
	) {
		connCount++
		if conn.Type == nil ||
			conn.Type.Value() != diagram.ConnectionValuesParof {
			t.Error("AddParentOfConnection should create parOf type")
		}
	}
	if connCount != 1 {
		t.Errorf(
			"Expected 1 connection from AddParentOfConnection, got %d",
			connCount,
		)
	}

	// Test AddPresentationOfConnection
	builder.AddPoint("pres1")
	builder.AddPoint("node3")
	builder.AddPresentationOfConnection("pres1", "node3")
	dataModel = builder.Build()

	connCount = 0
	for conn := range openxml.Elements[*diagram.Connection](
		dataModel.ConnectionList,
	) {
		connCount++
		if conn.Type == nil ||
			conn.Type.Value() != diagram.ConnectionValuesPresof {
			t.Error("AddPresentationOfConnection should create presOf type")
		}
	}
	if connCount != 1 {
		t.Errorf(
			"Expected 1 connection from AddPresentationOfConnection, got %d",
			connCount,
		)
	}
}

// TestDataModelBuilder_BuildRetainState tests BuildRetainState.
func TestDataModelBuilder_BuildRetainState(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	builder.AddPoint("node1")
	model1 := builder.BuildRetainState()

	// Builder should retain state
	if builder.PointCount() != 1 {
		t.Errorf(
			"PointCount should still be 1 after BuildRetainState, got %d",
			builder.PointCount(),
		)
	}

	// Add another point
	builder.AddPoint("node2")
	model2 := builder.BuildRetainState()

	// model2 should have 2 points
	ptCount1 := 0
	for range openxml.Elements[*diagram.Point](model1.PointList) {
		ptCount1++
	}
	ptCount2 := 0
	for range openxml.Elements[*diagram.Point](model2.PointList) {
		ptCount2++
	}

	if ptCount1 != 1 {
		t.Errorf("model1 should have 1 point, got %d", ptCount1)
	}
	if ptCount2 != 2 {
		t.Errorf("model2 should have 2 points, got %d", ptCount2)
	}
}

// TestDataModelBuilder_Clear tests Clear method.
func TestDataModelBuilder_Clear(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	builder.AddPoint("node1")
	builder.AddConnection("node1", "node2")
	builder.Clear()

	if builder.PointCount() != 0 {
		t.Errorf(
			"PointCount should be 0 after Clear, got %d",
			builder.PointCount(),
		)
	}
	if builder.ConnectionCount() != 0 {
		t.Errorf(
			"ConnectionCount should be 0 after Clear, got %d",
			builder.ConnectionCount(),
		)
	}
	if builder.HasPoints() {
		t.Error("HasPoints should return false after Clear")
	}
	if builder.HasConnections() {
		t.Error("HasConnections should return false after Clear")
	}
}

// TestDataModelBuilder_FullDiagram tests building a complete diagram.
func TestDataModelBuilder_FullDiagram(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	// Create a simple hierarchy
	builder.AddPoint("root")
	builder.AddPoint("child1")
	builder.AddPoint("child2")
	builder.AddAssistantPoint("assistant1")

	// Add connections
	builder.AddParentOfConnection("root", "child1")
	builder.AddParentOfConnection("root", "child2")
	builder.AddParentOfConnection("root", "assistant1")

	// Build
	dataModel := builder.Build()

	if builder.PointCount() != 0 {
		t.Errorf(
			"Builder should be reset after Build, got %d points",
			builder.PointCount(),
		)
	}

	// Verify data model
	if dataModel == nil {
		t.Fatal("Build returned nil")
	}

	ptCount := 0
	connCount := 0

	for range openxml.Elements[*diagram.Point](dataModel.PointList) {
		ptCount++
	}
	for range openxml.Elements[*diagram.Connection](
		dataModel.ConnectionList,
	) {
		connCount++
	}

	if ptCount != 4 {
		t.Errorf("Expected 4 points, got %d", ptCount)
	}
	if connCount != 3 {
		t.Errorf("Expected 3 connections, got %d", connCount)
	}

	// Verify validation works
	err := dataModel.Validate()
	if err != nil {
		t.Errorf("DataModel validation failed: %v", err)
	}
}

// TestDataModelBuilder_Chaining tests method chaining.
func TestDataModelBuilder_Chaining(t *testing.T) {
	dataModel := diagram.NewDataModelBuilder().
		AddPoint("n1").
		AddPoint("n2").
		AddConnection("n1", "n2").
		Build()

	if dataModel == nil {
		t.Fatal("Chained builder returned nil")
	}
	if dataModel.PointList == nil {
		t.Error("PointList should not be nil")
	}
	if dataModel.ConnectionList == nil {
		t.Error("ConnectionList should not be nil")
	}
}
