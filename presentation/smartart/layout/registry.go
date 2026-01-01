package layout

// Registry stores layout engine implementations
type Registry struct {
	engines map[LayoutType]LayoutEngine
}

var defaultRegistry = &Registry{
	engines: make(map[LayoutType]LayoutEngine),
}

// Register adds a layout engine to the registry
func (r *Registry) Register(engine LayoutEngine) {
	r.engines[engine.Type()] = engine
}

// GetEngine returns engine for the given type
func (r *Registry) GetEngine(t LayoutType) (LayoutEngine, bool) {
	e, ok := r.engines[t]

	return e, ok
}

// DefaultRegistry returns the global registry
func DefaultRegistry() *Registry {
	return defaultRegistry
}
