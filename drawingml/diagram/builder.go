package diagram

import (
	"fmt"

	"github.com/connerohnesorge/goffice/openxml/types"
)

// DataModelBuilder provides a fluent API for constructing DataModelRoot
// with points (nodes) and connections (relationships).
//
// The builder simplifies the process of creating SmartArt diagram data models
// by handling the boilerplate of creating and linking elements.
//
// Example usage:
//
//	builder := diagram.NewDataModelBuilder()
//	builder.AddPoint("node1", diagram.PointValuesNode).
//	       AddPoint("node2", diagram.PointValuesNode).
//	       AddConnection("node1", "node2", diagram.ConnectionValuesParof)
//	dataModel := builder.Build()
type DataModelBuilder struct {
	dataModel         *DataModelRoot
	pointList         *PointList
	connectionList    *ConnectionList
	pointCounter      int
	connectionCounter int
	lastPoint         *Point
}

// NewDataModelBuilder creates a new DataModelBuilder initialized with
// an empty DataModelRoot.
func NewDataModelBuilder() *DataModelBuilder {
	builder := &DataModelBuilder{
		dataModel:      NewDataModelRoot(),
		pointList:      nil,
		connectionList: nil,
	}

	return builder
}

// AddPoint adds a point (node) to the diagram with a given model ID.
// The point type defaults to PointValuesNode if not specified.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) AddPoint(modelID string, pointType ...PointValues) *DataModelBuilder {
	pt := NewPoint()
	pt.ModelId = types.NewStringValue(modelID)

	// Set point type (default to node if not specified)
	if len(pointType) > 0 && pointType[0] != "" {
		pt.Type = types.NewEnumValue(pointType[0])
	} else {
		pt.Type = types.NewEnumValue(PointValuesNode)
	}

	// Create point list if needed
	if b.pointList == nil {
		b.pointList = NewPointList()
		b.dataModel.PointList = b.pointList
	}

	// Append point to point list
	b.pointList.AppendChild(pt)
	b.pointCounter++
	b.lastPoint = pt

	return b
}

// WithText sets the text body of the last added point.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) WithText(text string) *DataModelBuilder {
	if b.lastPoint == nil {
		return b
	}

	// Create text body if needed
	if b.lastPoint.TextBody == nil {
		b.lastPoint.TextBody = NewTextBody()
	}

	// Add paragraph with text
	b.lastPoint.TextBody.AddParagraph(text)

	return b
}

// WithShapeProperties sets the shape properties of the last added point.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) WithShapeProperties(spPr *ShapeProperties) *DataModelBuilder {
	if b.lastPoint == nil {
		return b
	}

	if spPr != nil {
		b.lastPoint.ShapeProperties = spPr
	}

	return b
}

// AddConnection adds a connection (relationship) between two points.
// The connection type defaults to ConnectionValuesParof (parent-of) if not specified.
//
// Parameters:
//   - sourceID: The model ID of the source point
//   - destID: The model ID of the destination point
//   - connType: Optional connection type (defaults to ConnectionValuesParof)
//
// Returns builder for method chaining.
func (b *DataModelBuilder) AddConnection(sourceID, destID string, connType ...ConnectionValues) *DataModelBuilder {
	conn := NewConnection()
	conn.ModelId = types.NewStringValue(fmt.Sprintf("conn%d", b.connectionCounter+1))
	conn.SourceId = types.NewStringValue(sourceID)
	conn.DestinationId = types.NewStringValue(destID)

	// Set connection type (default to parOf if not specified)
	if len(connType) > 0 && connType[0] != "" {
		conn.Type = types.NewEnumValue(connType[0])
	} else {
		conn.Type = types.NewEnumValue(ConnectionValuesParof)
	}

	// Create connection list if needed
	if b.connectionList == nil {
		b.connectionList = NewConnectionList()
		b.dataModel.ConnectionList = b.connectionList
	}

	// Append connection to connection list
	b.connectionList.AppendChild(conn)
	b.connectionCounter++

	return b
}

// AddAssistantPoint adds an assistant point to the diagram.
// This is a convenience method equivalent to AddPoint with PointValuesAsst.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) AddAssistantPoint(modelID string) *DataModelBuilder {
	return b.AddPoint(modelID, PointValuesAsst)
}

// AddDocumentPoint adds a document point to the diagram.
// This is a convenience method equivalent to AddPoint with PointValuesDoc.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) AddDocumentPoint(modelID string) *DataModelBuilder {
	return b.AddPoint(modelID, PointValuesDoc)
}

// AddParentOfConnection adds a parent-of relationship connection.
// This is a convenience method equivalent to AddConnection with ConnectionValuesParof.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) AddParentOfConnection(sourceID, destID string) *DataModelBuilder {
	return b.AddConnection(sourceID, destID, ConnectionValuesParof)
}

// AddPresentationOfConnection adds a presentation-of relationship connection.
// This is a convenience method equivalent to AddConnection with ConnectionValuesPresof.
//
// Returns builder for method chaining.
func (b *DataModelBuilder) AddPresentationOfConnection(sourceID, destID string) *DataModelBuilder {
	return b.AddConnection(sourceID, destID, ConnectionValuesPresof)
}

// deepCloneDataModel performs a deep clone of a DataModelRoot including
// all child elements (points and connections). This is needed because
// the generated Clone() method only clones explicit fields, not children
// stored in the CompositeElementBase.
//
//nolint:revive // cognitive complexity is necessary for proper deep cloning
func deepCloneDataModel(source *DataModelRoot) *DataModelRoot {
	if source == nil {
		return nil
	}

	clone := NewDataModelRoot()

	if source.PointList != nil {
		clonePointList := NewPointList()
		for child := range source.PointList.Children() {
			if pt, ok := child.(*Point); ok {
				clonePointList.AppendChild(pt.Clone())
			}
		}
		clone.PointList = clonePointList
	}

	if source.ConnectionList != nil {
		cloneConnectionList := NewConnectionList()
		for child := range source.ConnectionList.Children() {
			if conn, ok := child.(*Connection); ok {
				cloneConnectionList.AppendChild(conn.Clone())
			}
		}
		clone.ConnectionList = cloneConnectionList
	}

	if source.Background != nil {
		if bg, ok := source.Background.Clone().(*Background); ok {
			clone.Background = bg
		}
	}

	if source.Whole != nil {
		if whole, ok := source.Whole.Clone().(*Whole); ok {
			clone.Whole = whole
		}
	}

	if source.DataModelExtensionList != nil {
		if extList, ok := source.DataModelExtensionList.Clone().(*DataModelExtensionList); ok {
			clone.DataModelExtensionList = extList
		}
	}

	return clone
}

// Build constructs and returns DataModelRoot with all added points
// and connections.
//
// After calling Build(), the builder can be reused to create another data model.
//
// Returns the constructed DataModelRoot.
func (b *DataModelBuilder) Build() *DataModelRoot {
	result := b.dataModel

	// Reset builder for reuse (optional, but allows chaining multiple builds)
	b.dataModel = NewDataModelRoot()
	b.pointList = nil
	b.connectionList = nil
	b.pointCounter = 0
	b.connectionCounter = 0

	return result
}

// BuildRetainState constructs and returns a deep clone of DataModelRoot
// but does not reset the builder's internal state. This allows continuing to add
// more elements after building.
//
// Use this when you want to build incrementally or export intermediate states.
//
// Returns a deep clone of the constructed DataModelRoot.
func (b *DataModelBuilder) BuildRetainState() *DataModelRoot {
	return deepCloneDataModel(b.dataModel)
}

// PointCount returns the number of points added to the builder.
func (b *DataModelBuilder) PointCount() int {
	return b.pointCounter
}

// ConnectionCount returns the number of connections added to the builder.
func (b *DataModelBuilder) ConnectionCount() int {
	return b.connectionCounter
}

// HasPoints returns true if at least one point has been added.
func (b *DataModelBuilder) HasPoints() bool {
	return b.pointCounter > 0
}

// HasConnections returns true if at least one connection has been added.
func (b *DataModelBuilder) HasConnections() bool {
	return b.connectionCounter > 0
}

// Clear resets the builder, removing all points and connections.
// The builder can then be used to create a new data model.
func (b *DataModelBuilder) Clear() {
	b.dataModel = NewDataModelRoot()
	b.pointList = nil
	b.connectionList = nil
	b.pointCounter = 0
	b.connectionCounter = 0
}
