// Package openxml provides advanced relationship management capabilities.
package openxml

import (
	"fmt"
	"sync"
)

// RelationshipGraph manages part dependencies and relationship chains.
type RelationshipGraph struct {
	mu           sync.RWMutex
	dependencies map[string][]string // part URI -> dependent part URIs
	reverse      map[string][]string // part URI -> parts that depend on it
}

// NewRelationshipGraph creates a new relationship dependency graph.
func NewRelationshipGraph() *RelationshipGraph {
	return &RelationshipGraph{
		dependencies: make(map[string][]string),
		reverse:      make(map[string][]string),
	}
}

// AddDependency records that sourcePart depends on targetPart.
func (g *RelationshipGraph) AddDependency(sourcePart, targetPart string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Add forward dependency
	if !contains(g.dependencies[sourcePart], targetPart) {
		g.dependencies[sourcePart] = append(g.dependencies[sourcePart], targetPart)
	}

	// Add reverse dependency
	if !contains(g.reverse[targetPart], sourcePart) {
		g.reverse[targetPart] = append(g.reverse[targetPart], sourcePart)
	}
}

// RemoveDependency removes a dependency between sourcePart and targetPart.
func (g *RelationshipGraph) RemoveDependency(sourcePart, targetPart string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.dependencies[sourcePart] = remove(g.dependencies[sourcePart], targetPart)
	g.reverse[targetPart] = remove(g.reverse[targetPart], sourcePart)
}

// GetDependencies returns all parts that the given part depends on.
func (g *RelationshipGraph) GetDependencies(partURI string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	deps := g.dependencies[partURI]
	result := make([]string, len(deps))
	copy(result, deps)
	return result
}

// GetDependents returns all parts that depend on the given part.
func (g *RelationshipGraph) GetDependents(partURI string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	deps := g.reverse[partURI]
	result := make([]string, len(deps))
	copy(result, deps)
	return result
}

// DetectCircularDependency checks if adding a dependency would create a cycle.
func (g *RelationshipGraph) DetectCircularDependency(from, to string) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()

	// If "to" is reachable from "from", adding this edge creates a cycle
	visited := make(map[string]bool)
	return g.isReachable(to, from, visited)
}

// isReachable checks if target is reachable from source using DFS.
func (g *RelationshipGraph) isReachable(source, target string, visited map[string]bool) bool {
	if source == target {
		return true
	}

	if visited[source] {
		return false
	}

	visited[source] = true

	for _, dep := range g.dependencies[source] {
		if g.isReachable(dep, target, visited) {
			return true
		}
	}

	return false
}

// TopologicalSort returns parts in topological order (dependencies first).
// Returns error if circular dependencies exist.
func (g *RelationshipGraph) TopologicalSort(parts []string) ([]string, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	inDegree := make(map[string]int)
	for _, part := range parts {
		inDegree[part] = 0
	}

	// Calculate in-degrees
	for _, part := range parts {
		for _, dep := range g.dependencies[part] {
			if _, exists := inDegree[dep]; exists {
				inDegree[dep]++
			}
		}
	}

	// Queue nodes with no incoming edges
	queue := make([]string, 0)
	for part, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, part)
		}
	}

	result := make([]string, 0, len(parts))
	for len(queue) > 0 {
		// Dequeue
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// Reduce in-degree of dependents
		for _, dependent := range g.reverse[current] {
			if _, exists := inDegree[dependent]; exists {
				inDegree[dependent]--
				if inDegree[dependent] == 0 {
					queue = append(queue, dependent)
				}
			}
		}
	}

	// If not all nodes processed, there's a cycle
	if len(result) != len(parts) {
		return nil, fmt.Errorf("circular dependency detected")
	}

	return result, nil
}

// GetOrphanedParts returns parts that are not referenced by any other part.
func (g *RelationshipGraph) GetOrphanedParts(allParts []string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	orphaned := make([]string, 0)
	for _, part := range allParts {
		if len(g.reverse[part]) == 0 {
			orphaned = append(orphaned, part)
		}
	}
	return orphaned
}

// GetTransitiveDependencies returns all parts reachable from the given part.
func (g *RelationshipGraph) GetTransitiveDependencies(partURI string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	visited := make(map[string]bool)
	result := make([]string, 0)
	g.visitDependencies(partURI, visited, &result)
	return result
}

// visitDependencies performs DFS to collect all dependencies.
func (g *RelationshipGraph) visitDependencies(part string, visited map[string]bool, result *[]string) {
	if visited[part] {
		return
	}
	visited[part] = true

	for _, dep := range g.dependencies[part] {
		*result = append(*result, dep)
		g.visitDependencies(dep, visited, result)
	}
}

// RelationshipCloneOptions configures relationship cloning behavior.
type RelationshipCloneOptions struct {
	// PreserveIDs keeps original relationship IDs if true
	PreserveIDs bool
	// IDMapping maps old relationship IDs to new ones
	IDMapping map[string]string
	// UpdateCallback is called for each cloned relationship
	UpdateCallback func(oldRel, newRel OpenXmlRelationship) error
}

// RelationshipCloner handles cloning relationships with proper ID rewriting.
type RelationshipCloner struct {
	options RelationshipCloneOptions
	idGen   *RelationshipIDGenerator
}

// NewRelationshipCloner creates a new relationship cloner.
func NewRelationshipCloner(options RelationshipCloneOptions) *RelationshipCloner {
	return &RelationshipCloner{
		options: options,
		idGen:   NewRelationshipIDGenerator(),
	}
}

// CloneRelationship creates a copy of a relationship with optional ID remapping.
func (c *RelationshipCloner) CloneRelationship(
	rel OpenXmlRelationship,
	newContainer OpenXmlPartContainer,
) (OpenXmlRelationship, error) {
	// Determine new ID
	var newID string
	if c.options.PreserveIDs {
		newID = rel.ID()
	} else if mappedID, exists := c.options.IDMapping[rel.ID()]; exists {
		newID = mappedID
	} else {
		newID = c.idGen.Next()
	}

	// Clone based on relationship type
	var newRel OpenXmlRelationship
	switch r := rel.(type) {
	case *PartRelationship:
		newRel = NewPartRelationship(newID, r.Type(), r.TargetPart(), newContainer)
	case *ExternalRelationship:
		newRel = NewExternalRelationship(newID, r.Type(), r.Target(), newContainer)
	case *HyperlinkRelationship:
		newRel = NewHyperlinkRelationship(newID, r.Target(), r.IsExternal(), newContainer)
	case *DataPartReferenceRelationship:
		newRel = NewDataPartReferenceRelationship(newID, r.Type(), r.Target(), newContainer)
	default:
		return nil, fmt.Errorf("unsupported relationship type: %T", rel)
	}

	// Call update callback if provided
	if c.options.UpdateCallback != nil {
		if err := c.options.UpdateCallback(rel, newRel); err != nil {
			return nil, fmt.Errorf("clone callback failed: %w", err)
		}
	}

	return newRel, nil
}

// RelationshipValidator validates relationship consistency.
type RelationshipValidator struct {
	graph *RelationshipGraph
}

// NewRelationshipValidator creates a new validator.
func NewRelationshipValidator(graph *RelationshipGraph) *RelationshipValidator {
	return &RelationshipValidator{
		graph: graph,
	}
}

// ValidateNoCycles checks for circular dependencies.
func (v *RelationshipValidator) ValidateNoCycles(parts []string) error {
	_, err := v.graph.TopologicalSort(parts)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

// ValidateNoOrphans checks for orphaned parts (excluding root parts).
func (v *RelationshipValidator) ValidateNoOrphans(allParts, rootParts []string) error {
	orphaned := v.graph.GetOrphanedParts(allParts)

	// Filter out root parts (they're allowed to be orphaned)
	actualOrphans := make([]string, 0)
	for _, orphan := range orphaned {
		isRoot := false
		for _, root := range rootParts {
			if orphan == root {
				isRoot = true
				break
			}
		}
		if !isRoot {
			actualOrphans = append(actualOrphans, orphan)
		}
	}

	if len(actualOrphans) > 0 {
		return fmt.Errorf("orphaned parts detected: %v", actualOrphans)
	}

	return nil
}

// RelationshipPath represents a path through relationships.
type RelationshipPath struct {
	Parts []OpenXmlPart
}

// Length returns the number of parts in the path minus one (the number of hops).
func (p *RelationshipPath) Length() int {
	if len(p.Parts) == 0 {
		return 0
	}
	return len(p.Parts) - 1
}

// LastPart returns the final part in the path.
func (p *RelationshipPath) LastPart() OpenXmlPart {
	if len(p.Parts) == 0 {
		return nil
	}
	return p.Parts[len(p.Parts)-1]
}

// RelationshipPathFinder finds paths through part hierarchies.
type RelationshipPathFinder struct {
	maxDepth int
}

// NewRelationshipPathFinder creates a new path finder.
func NewRelationshipPathFinder(maxDepth int) *RelationshipPathFinder {
	if maxDepth <= 0 {
		maxDepth = 10 // Default max depth
	}
	return &RelationshipPathFinder{
		maxDepth: maxDepth,
	}
}

// FindPath finds a path from source to target through the part hierarchy.
func (f *RelationshipPathFinder) FindPath(
	source OpenXmlPartContainer,
	targetURI string,
) (*RelationshipPath, error) {
	visited := make(map[string]bool)
	path := &RelationshipPath{}

	if f.findPathRecursive(source, targetURI, 0, visited, path) {
		return path, nil
	}

	return nil, fmt.Errorf("no path found to %s", targetURI)
}

// findPathRecursive performs DFS to find a path.
func (f *RelationshipPathFinder) findPathRecursive(
	container OpenXmlPartContainer,
	targetURI string,
	depth int,
	visited map[string]bool,
	path *RelationshipPath,
) bool {
	if depth >= f.maxDepth {
		return false
	}

	// Check all child parts
	for part := range container.Parts() {
		partURI := part.URI()
		
		if visited[partURI] {
			continue
		}
		
		if partURI == targetURI {
			// Found target
			path.Parts = append(path.Parts, part)
			return true
		}

		// Mark as visited
		visited[partURI] = true

		// Try going deeper - OpenXmlPart also implements OpenXmlPartContainer
		path.Parts = append(path.Parts, part)

		if partContainer, ok := part.(OpenXmlPartContainer); ok {
			if f.findPathRecursive(partContainer, targetURI, depth+1, visited, path) {
				return true
			}
		}

		// Backtrack
		path.Parts = path.Parts[:len(path.Parts)-1]
		visited[partURI] = false
	}

	return false
}

// Helper functions

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func remove(slice []string, item string) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}
