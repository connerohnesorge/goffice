package openxml

import (
	"io"
	"iter"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/packaging"
)

func TestRelationshipGraph_AddDependency(t *testing.T) {
	graph := NewRelationshipGraph()

	graph.AddDependency(testDocumentXML, "/word/styles.xml")
	graph.AddDependency(testDocumentXML, "/word/numbering.xml")

	deps := graph.GetDependencies(testDocumentXML)
	if len(deps) != 2 {
		t.Errorf("expected 2 dependencies, got %d", len(deps))
	}

	dependents := graph.GetDependents("/word/styles.xml")
	if len(dependents) != 1 || dependents[0] != testDocumentXML {
		t.Errorf("expected document.xml as dependent of styles.xml")
	}
}

func TestRelationshipGraph_RemoveDependency(t *testing.T) {
	graph := NewRelationshipGraph()

	graph.AddDependency(testDocumentXML, "/word/styles.xml")
	graph.RemoveDependency(testDocumentXML, "/word/styles.xml")

	deps := graph.GetDependencies(testDocumentXML)
	if len(deps) != 0 {
		t.Errorf("expected no dependencies after removal, got %d", len(deps))
	}
}

func TestRelationshipGraph_DetectCircularDependency(t *testing.T) {
	graph := NewRelationshipGraph()

	// Create a chain: A -> B -> C
	graph.AddDependency("A", "B")
	graph.AddDependency("B", "C")

	// Try to add C -> A (would create cycle)
	hasCycle := graph.DetectCircularDependency("C", "A")
	if !hasCycle {
		t.Error("expected circular dependency to be detected")
	}

	// Try to add C -> D (no cycle)
	hasCycle = graph.DetectCircularDependency("C", "D")
	if hasCycle {
		t.Error("expected no circular dependency")
	}
}

func TestRelationshipGraph_TopologicalSort(t *testing.T) {
	graph := NewRelationshipGraph()

	// Create dependencies: document -> styles, numbering
	// numbering -> abstractNum
	graph.AddDependency("/word/document.xml", "/word/styles.xml")
	graph.AddDependency("/word/document.xml", "/word/numbering.xml")
	graph.AddDependency("/word/numbering.xml", "/word/abstractNum.xml")

	parts := []string{
		"/word/document.xml",
		"/word/styles.xml",
		"/word/numbering.xml",
		"/word/abstractNum.xml",
	}

	sorted, err := graph.TopologicalSort(parts)
	if err != nil {
		t.Fatalf("topological sort failed: %v", err)
	}

	// Verify abstractNum comes before numbering
	abstractNumIdx := -1
	numberingIdx := -1
	docIdx := -1

	for i, part := range sorted {
		switch part {
		case "/word/abstractNum.xml":
			abstractNumIdx = i
		case "/word/numbering.xml":
			numberingIdx = i
		case testDocumentXML:
			docIdx = i
		}
	}

	if abstractNumIdx > numberingIdx {
		t.Error("abstractNum should come before numbering in topological order")
	}
	if numberingIdx > docIdx {
		t.Error("numbering should come before document in topological order")
	}
}

func TestRelationshipGraph_TopologicalSortWithCycle(t *testing.T) {
	graph := NewRelationshipGraph()

	// Create a cycle: A -> B -> C -> A
	graph.AddDependency("A", "B")
	graph.AddDependency("B", "C")
	graph.AddDependency("C", "A")

	parts := []string{"A", "B", "C"}

	_, err := graph.TopologicalSort(parts)
	if err == nil {
		t.Error("expected error for circular dependencies")
	}
}

func TestRelationshipGraph_GetOrphanedParts(t *testing.T) {
	graph := NewRelationshipGraph()

	// Create dependencies
	graph.AddDependency("/word/document.xml", "/word/styles.xml")
	graph.AddDependency("/word/document.xml", "/word/numbering.xml")

	allParts := []string{
		"/word/document.xml",
		"/word/styles.xml",
		"/word/numbering.xml",
		"/word/orphaned.xml",
	}

	orphaned := graph.GetOrphanedParts(allParts)

	// document.xml and orphaned.xml should be orphaned (no one depends on them)
	if len(orphaned) != 2 {
		t.Errorf("expected 2 orphaned parts, got %d: %v", len(orphaned), orphaned)
	}

	hasDocument := false
	hasOrphaned := false
	for _, part := range orphaned {
		if part == testDocumentXML {
			hasDocument = true
		}
		if part == "/word/orphaned.xml" {
			hasOrphaned = true
		}
	}

	if !hasDocument || !hasOrphaned {
		t.Errorf("expected document.xml and orphaned.xml to be orphaned, got: %v", orphaned)
	}
}

func TestRelationshipGraph_GetTransitiveDependencies(t *testing.T) {
	graph := NewRelationshipGraph()

	// Create chain: A -> B -> C -> D
	graph.AddDependency("A", "B")
	graph.AddDependency("B", "C")
	graph.AddDependency("C", "D")

	transitive := graph.GetTransitiveDependencies("A")

	// Should include B, C, D
	if len(transitive) != 3 {
		t.Errorf("expected 3 transitive dependencies, got %d: %v", len(transitive), transitive)
	}

	expected := map[string]bool{"B": true, "C": true, "D": true}
	for _, dep := range transitive {
		if !expected[dep] {
			t.Errorf("unexpected dependency: %s", dep)
		}
	}
}

func TestRelationshipCloner_CloneRelationship(t *testing.T) {
	// Create a mock container
	container := &mockPartContainer{}

	// Create original relationship
	mockPart := &mockPart{uri: "/word/styles.xml"}
	origRel := NewPartRelationship("rId5", RelationshipTypeStyles, mockPart, container)

	// Clone with explicit ID mapping
	options := RelationshipCloneOptions{
		PreserveIDs: false,
		IDMapping:   map[string]string{"rId5": "rId10"},
	}
	cloner := NewRelationshipCloner(options)

	newRel, err := cloner.CloneRelationship(origRel, container)
	if err != nil {
		t.Fatalf("cloning failed: %v", err)
	}

	t.Logf("Original ID: %s, New ID: %s", origRel.ID(), newRel.ID())
	if newRel.ID() == origRel.ID() {
		t.Errorf("cloned relationship should have different ID, both are %s", origRel.ID())
	}

	if newRel.ID() != "rId10" {
		t.Errorf("cloned relationship should have mapped ID rId10, got %s", newRel.ID())
	}

	if newRel.Type() != origRel.Type() {
		t.Error("cloned relationship should have same type")
	}
}

func TestRelationshipCloner_CloneWithCallback(t *testing.T) {
	container := &mockPartContainer{}
	mockPart := &mockPart{uri: "/word/styles.xml"}
	origRel := NewPartRelationship("rId1", RelationshipTypeStyles, mockPart, container)

	callbackCalled := false
	options := RelationshipCloneOptions{
		PreserveIDs: false,
		UpdateCallback: func(oldRel, newRel OpenXmlRelationship) error {
			callbackCalled = true
			if oldRel.Type() != newRel.Type() {
				t.Error("callback received different types")
			}

			return nil
		},
	}

	cloner := NewRelationshipCloner(options)
	_, err := cloner.CloneRelationship(origRel, container)
	if err != nil {
		t.Fatalf("cloning failed: %v", err)
	}

	if !callbackCalled {
		t.Error("update callback was not called")
	}
}

func TestRelationshipValidator_ValidateNoCycles(t *testing.T) {
	graph := NewRelationshipGraph()
	validator := NewRelationshipValidator(graph)

	// Create valid dependencies
	graph.AddDependency("A", "B")
	graph.AddDependency("B", "C")

	parts := []string{"A", "B", "C"}
	err := validator.ValidateNoCycles(parts)
	if err != nil {
		t.Errorf("validation should pass for acyclic graph: %v", err)
	}

	// Add cycle
	graph.AddDependency("C", "A")

	err = validator.ValidateNoCycles(parts)
	if err == nil {
		t.Error("validation should fail for cyclic graph")
	}
}

func TestRelationshipValidator_ValidateNoOrphans(t *testing.T) {
	graph := NewRelationshipGraph()
	validator := NewRelationshipValidator(graph)

	// Create dependencies
	graph.AddDependency("/word/document.xml", "/word/styles.xml")

	allParts := []string{
		"/word/document.xml",
		"/word/styles.xml",
		"/word/orphaned.xml",
	}
	rootParts := []string{"/word/document.xml"}

	err := validator.ValidateNoOrphans(allParts, rootParts)
	if err == nil {
		t.Error("validation should fail with orphaned parts")
	}

	// Add dependency to orphaned part
	graph.AddDependency("/word/document.xml", "/word/orphaned.xml")

	err = validator.ValidateNoOrphans(allParts, rootParts)
	if err != nil {
		t.Errorf("validation should pass with no orphans: %v", err)
	}
}

func TestRelationshipPathFinder_FindPath(t *testing.T) {
	// Create mock parts
	styles := &mockPart{uri: "/word/styles.xml"}
	numbering := &mockPart{uri: "/word/numbering.xml"}
	theme := &mockPart{uri: "/word/theme/theme1.xml"}

	// Add theme as a child of styles
	styles.parts = []OpenXmlPart{theme}

	container := &mockPartContainer{
		parts: []OpenXmlPart{styles, numbering},
	}

	finder := NewRelationshipPathFinder(10)
	path, err := finder.FindPath(container, "/word/styles.xml")
	if err != nil {
		t.Fatalf("path finding failed: %v", err)
	}

	if path.Length() != 0 {
		t.Errorf("expected path length 0 (direct child), got %d", path.Length())
	}

	// Find path through chain: container -> styles -> theme
	path, err = finder.FindPath(container, "/word/theme/theme1.xml")
	if err != nil {
		t.Fatalf("path finding through chain failed: %v", err)
	}

	if len(path.Parts) < 2 {
		t.Errorf("expected at least 2 parts in path, got %d", len(path.Parts))
	}
}

func TestRelationshipPathFinder_NoPath(t *testing.T) {
	container := &mockPartContainer{
		relationships: []OpenXmlRelationship{},
	}

	finder := NewRelationshipPathFinder(10)
	_, err := finder.FindPath(container, "/word/nonexistent.xml")
	if err == nil {
		t.Error("expected error when no path exists")
	}
}

func TestRelationshipPath_Methods(t *testing.T) {
	styles := &mockPart{uri: "/word/styles.xml"}
	numbering := &mockPart{uri: "/word/numbering.xml"}

	path := &RelationshipPath{
		Parts: []OpenXmlPart{styles, numbering},
	}

	if path.Length() != 1 {
		t.Errorf("expected length 1 (one hop), got %d", path.Length())
	}

	lastPart := path.LastPart()
	if lastPart.URI() != numbering.URI() {
		t.Error("LastPart() returned wrong part")
	}
}

// Mock types for testing

type mockPartContainer struct {
	relationships []OpenXmlRelationship
	parts         []OpenXmlPart
}

func (m *mockPartContainer) Parts() iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		for _, part := range m.parts {
			if !yield(part) {
				return
			}
		}
	}
}

func (m *mockPartContainer) GetPartById(id string) (OpenXmlPart, error) {
	return nil, nil
}

func (m *mockPartContainer) GetPartsOfType(contentType string) iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {}
}

func (m *mockPartContainer) AddPart(part OpenXmlPart, id string) error {
	m.parts = append(m.parts, part)

	return nil
}

func (m *mockPartContainer) DeletePart(id string) error {
	return nil
}

func (m *mockPartContainer) Features() *features.FeatureCollection {
	return nil
}

func (m *mockPartContainer) GetPackagingPart(uri string) *packaging.Part {
	return nil
}

func (m *mockPartContainer) Package() *packaging.Package {
	return nil
}

func (m *mockPartContainer) URI() string {
	return "/mock/container"
}

type mockPart struct {
	uri   string
	parts []OpenXmlPart
}

func (m *mockPart) URI() string {
	return m.uri
}

func (m *mockPart) Parts() iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		for _, part := range m.parts {
			if !yield(part) {
				return
			}
		}
	}
}

func (m *mockPart) GetPartById(id string) (OpenXmlPart, error) {
	return nil, nil
}

func (m *mockPart) GetPartsOfType(contentType string) iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {}
}

func (m *mockPart) AddPart(part OpenXmlPart, id string) error {
	m.parts = append(m.parts, part)

	return nil
}

func (m *mockPart) DeletePart(id string) error {
	return nil
}

func (m *mockPart) ContentType() string {
	return "application/xml"
}

func (m *mockPart) Root() Element {
	return nil
}

func (m *mockPart) Features() *features.FeatureCollection {
	return nil
}

func (m *mockPart) GetStream() io.Reader {
	return nil
}

func (m *mockPart) RootElement() PartRootElement {
	return nil
}

func (m *mockPart) SetData(data []byte) {
}

func (m *mockPart) Package() *packaging.Package {
	return nil
}

func (m *mockPart) GetPackagingPart(uri string) *packaging.Part {
	return nil
}

// Ensure mockPart implements both interfaces
var (
	_ OpenXmlPart          = (*mockPart)(nil)
	_ OpenXmlPartContainer = (*mockPart)(nil)
)
