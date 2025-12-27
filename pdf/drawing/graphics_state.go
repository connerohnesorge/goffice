// Package drawing provides graphics state management for PDF content streams.

package drawing

import (
	"fmt"
	"strings"
)

// BlendMode represents the PDF blend mode for transparency compositing.
// These correspond to the /BM entry in extended graphics state dictionaries.
type BlendMode string

const (
	// BlendModeNormal is the default blend mode.
	BlendModeNormal BlendMode = "Normal"
	// BlendModeMultiply multiplies the backdrop and source colors.
	BlendModeMultiply BlendMode = "Multiply"
	// BlendModeScreen screens the colors.
	BlendModeScreen BlendMode = "Screen"
	// BlendModeOverlay combines multiply and screen.
	BlendModeOverlay BlendMode = "Overlay"
	// BlendModeDarken selects the darker of the colors.
	BlendModeDarken BlendMode = "Darken"
	// BlendModeLighten selects the lighter of the colors.
	BlendModeLighten BlendMode = "Lighten"
	// BlendModeColorDodge brightens the backdrop.
	BlendModeColorDodge BlendMode = "ColorDodge"
	// BlendModeColorBurn darkens the backdrop.
	BlendModeColorBurn BlendMode = "ColorBurn"
	// BlendModeHardLight combines multiply and screen like overlay.
	BlendModeHardLight BlendMode = "HardLight"
	// BlendModeSoftLight darkens or lightens depending on source.
	BlendModeSoftLight BlendMode = "SoftLight"
	// BlendModeDifference subtracts the darker from the lighter.
	BlendModeDifference BlendMode = "Difference"
	// BlendModeExclusion produces lower contrast than difference.
	BlendModeExclusion BlendMode = "Exclusion"
	// BlendModeHue uses hue of source with saturation/luminosity of backdrop.
	BlendModeHue BlendMode = "Hue"
	// BlendModeSaturation uses saturation of source.
	BlendModeSaturation BlendMode = "Saturation"
	// BlendModeColor uses hue/saturation of source with luminosity of backdrop.
	BlendModeColor BlendMode = "Color"
	// BlendModeLuminosity uses luminosity of source.
	BlendModeLuminosity BlendMode = "Luminosity"
)

// TextState holds the current text rendering state.
// This corresponds to PDF text state parameters set by Tf, Tc, Tw, etc.
type TextState struct {
	// FontName is the resource name of the current font (e.g., "F1").
	FontName string
	// FontSize is the current font size in points.
	FontSize float64
	// CharacterSpacing is additional spacing between characters (Tc operator).
	CharacterSpacing float64
	// WordSpacing is additional spacing for space characters (Tw operator).
	WordSpacing float64
	// HorizontalScaling is the horizontal text scaling percentage (Tz operator).
	// 100 is normal, 50 is half-width, 200 is double-width.
	HorizontalScaling float64
	// Leading is the distance between baselines (TL operator).
	Leading float64
	// TextRise is vertical offset from baseline (Ts operator).
	TextRise float64
	// RenderingMode specifies how text is painted (Tr operator).
	RenderingMode TextRenderingMode
}

// NewTextState creates a new TextState with default values.
func NewTextState() TextState {
	return TextState{
		FontSize:          12.0,
		HorizontalScaling: 100.0,
		Leading:           14.4, // 1.2 * 12pt
		RenderingMode:     TextRenderFill,
	}
}

// Clone returns a copy of the text state.
func (ts TextState) Clone() TextState {
	return ts
}

// StrokeState holds the current stroke/line rendering state.
type StrokeState struct {
	// Color is the current stroke color.
	Color Color
	// Width is the line width in points (w operator).
	Width float64
	// Cap is the line cap style (J operator).
	Cap LineCap
	// Join is the line join style (j operator).
	Join LineJoin
	// MiterLimit is the miter limit for miter joins (M operator).
	MiterLimit float64
	// Dash is the dash pattern (d operator).
	Dash DashPattern
}

// NewStrokeState creates a new StrokeState with default values.
func NewStrokeState() StrokeState {
	return StrokeState{
		Color:      Black,
		Width:      DefaultLineWidth,
		Cap:        LineCapButt,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashSolid,
	}
}

// Clone returns a copy of the stroke state.
func (ss StrokeState) Clone() StrokeState {
	return StrokeState{
		Color:      ss.Color,
		Width:      ss.Width,
		Cap:        ss.Cap,
		Join:       ss.Join,
		MiterLimit: ss.MiterLimit,
		Dash: DashPattern{
			Array: append(
				[]float64(nil),
				ss.Dash.Array...),
			Phase: ss.Dash.Phase,
		},
	}
}

// FillState holds the current fill state.
type FillState struct {
	// Color is the current fill color.
	Color Color
}

// NewFillState creates a new FillState with default values.
func NewFillState() FillState {
	return FillState{
		Color: Black,
	}
}

// Clone returns a copy of the fill state.
func (fs FillState) Clone() FillState {
	return fs
}

// TransparencyState holds transparency-related state.
type TransparencyState struct {
	// StrokeAlpha is the alpha value for stroking operations (CA in ExtGState).
	// Range is 0 (transparent) to 1 (opaque).
	StrokeAlpha float64
	// FillAlpha is the alpha value for fill operations (ca in ExtGState).
	// Range is 0 (transparent) to 1 (opaque).
	FillAlpha float64
	// BlendMode is the blend mode for transparency compositing.
	BlendMode BlendMode
	// AlphaIsShape indicates if alpha should be interpreted as shape (AIS in ExtGState).
	AlphaIsShape bool
}

// NewTransparencyState creates a new TransparencyState with default values.
func NewTransparencyState() TransparencyState {
	return TransparencyState{
		StrokeAlpha: 1.0,
		FillAlpha:   1.0,
		BlendMode:   BlendModeNormal,
	}
}

// Clone returns a copy of the transparency state.
func (ts TransparencyState) Clone() TransparencyState {
	return ts
}

// IsOpaque returns true if both stroke and fill are fully opaque.
func (ts TransparencyState) IsOpaque() bool {
	return ts.StrokeAlpha >= 1.0 &&
		ts.FillAlpha >= 1.0
}

// ClippingState holds the current clipping path state.
// Note: The actual clipping path cannot be queried in PDF; this tracks metadata.
type ClippingState struct {
	// Depth tracks the nesting depth of clipping regions.
	Depth int
	// Bounds approximates the current clipping bounds (may be nil).
	Bounds *Rect
}

// NewClippingState creates a new ClippingState.
func NewClippingState() ClippingState {
	return ClippingState{
		Depth: 0,
	}
}

// Clone returns a copy of the clipping state.
func (cs ClippingState) Clone() ClippingState {
	clone := ClippingState{
		Depth: cs.Depth,
	}
	if cs.Bounds != nil {
		bounds := *cs.Bounds
		clone.Bounds = &bounds
	}

	return clone
}

// FullGraphicsState represents the complete PDF graphics state.
// This tracks all parameters that can be saved and restored with q/Q operators.
type FullGraphicsState struct {
	// Transform is the current transformation matrix (CTM).
	Transform Transform
	// Stroke holds stroke-related state.
	Stroke StrokeState
	// Fill holds fill-related state.
	Fill FillState
	// Text holds text-related state.
	Text TextState
	// Transparency holds alpha and blending state.
	Transparency TransparencyState
	// Clipping holds clipping path state.
	Clipping ClippingState
	// Flatness is the flatness tolerance (i operator).
	Flatness float64
	// SmoothnessTolerance controls color gradient smoothness.
	SmoothnessTolerance float64
}

// NewFullGraphicsState creates a new FullGraphicsState with default values.
func NewFullGraphicsState() FullGraphicsState {
	return FullGraphicsState{
		Transform:    Identity(),
		Stroke:       NewStrokeState(),
		Fill:         NewFillState(),
		Text:         NewTextState(),
		Transparency: NewTransparencyState(),
		Clipping:     NewClippingState(),
		Flatness:     0, // 0 means device default
	}
}

// Clone returns a deep copy of the graphics state.
func (gs FullGraphicsState) Clone() FullGraphicsState {
	return FullGraphicsState{
		Transform:           gs.Transform,
		Stroke:              gs.Stroke.Clone(),
		Fill:                gs.Fill.Clone(),
		Text:                gs.Text.Clone(),
		Transparency:        gs.Transparency.Clone(),
		Clipping:            gs.Clipping.Clone(),
		Flatness:            gs.Flatness,
		SmoothnessTolerance: gs.SmoothnessTolerance,
	}
}

// GraphicsStateStack manages a stack of complete graphics states.
// It generates PDF operators for state changes and tracks the current state.
type GraphicsStateStack struct {
	stack         []FullGraphicsState
	current       FullGraphicsState
	operators     strings.Builder
	extGStateRefs map[string]bool // Tracks which ExtGState resources have been used
	//nolint:unused // Reserved for future ExtGState name generation
	extGStateNum int // Counter for generating ExtGState names
}

// NewGraphicsStateStack creates a new graphics state stack manager.
func NewGraphicsStateStack() *GraphicsStateStack {
	return &GraphicsStateStack{
		current:       NewFullGraphicsState(),
		extGStateRefs: make(map[string]bool),
	}
}

// Save saves the current graphics state onto the stack (q operator).
func (gss *GraphicsStateStack) Save() *GraphicsStateStack {
	gss.stack = append(
		gss.stack,
		gss.current.Clone(),
	)
	gss.operators.WriteString("q\n")

	return gss
}

// Restore restores the graphics state from the stack (Q operator).
// If the stack is empty, this still outputs Q but doesn't modify internal state.
func (gss *GraphicsStateStack) Restore() *GraphicsStateStack {
	if len(gss.stack) > 0 {
		gss.current = gss.stack[len(gss.stack)-1]
		gss.stack = gss.stack[:len(gss.stack)-1]
	}
	gss.operators.WriteString("Q\n")

	return gss
}

// Depth returns the current stack depth.
func (gss *GraphicsStateStack) Depth() int {
	return len(gss.stack)
}

// CurrentState returns a copy of the current graphics state.
func (gss *GraphicsStateStack) CurrentState() FullGraphicsState {
	return gss.current.Clone()
}

// CurrentTransform returns the current transformation matrix.
func (gss *GraphicsStateStack) CurrentTransform() Transform {
	return gss.current.Transform
}

// SetTransform sets the transformation matrix (cm operator).
// This calculates and applies the delta from current to new transform.
func (gss *GraphicsStateStack) SetTransform(
	t Transform,
) *GraphicsStateStack {
	delta := gss.current.Transform.Inverse().
		Concat(t)
	if !delta.IsIdentity() {
		gss.operators.WriteString(
			delta.ToCMOperator(),
		)
		gss.operators.WriteByte('\n')
		gss.current.Transform = t
	}

	return gss
}

// ApplyTransform concatenates a transform with the current CTM (cm operator).
func (gss *GraphicsStateStack) ApplyTransform(
	t Transform,
) *GraphicsStateStack {
	if !t.IsIdentity() {
		gss.operators.WriteString(
			t.ToCMOperator(),
		)
		gss.operators.WriteByte('\n')
		gss.current.Transform = gss.current.Transform.Concat(
			t,
		)
	}

	return gss
}

// Translate applies a translation transform.
func (gss *GraphicsStateStack) Translate(
	tx, ty float64,
) *GraphicsStateStack {
	return gss.ApplyTransform(
		Identity().Translate(tx, ty),
	)
}

// Scale applies a scaling transform.
func (gss *GraphicsStateStack) Scale(
	sx, sy float64,
) *GraphicsStateStack {
	return gss.ApplyTransform(
		Identity().Scale(sx, sy),
	)
}

// Rotate applies a rotation transform (degrees, counter-clockwise).
func (gss *GraphicsStateStack) Rotate(
	degrees float64,
) *GraphicsStateStack {
	return gss.ApplyTransform(
		Identity().Rotate(degrees),
	)
}

// RotateAround applies rotation around a specific point.
func (gss *GraphicsStateStack) RotateAround(
	degrees, cx, cy float64,
) *GraphicsStateStack {
	return gss.ApplyTransform(
		Identity().RotateAround(degrees, cx, cy),
	)
}

// SetLineWidth sets the line width (w operator).
func (gss *GraphicsStateStack) SetLineWidth(
	width float64,
) *GraphicsStateStack {
	if gss.current.Stroke.Width != width {
		gss.operators.WriteString(
			fmt.Sprintf(
				"%s w\n",
				formatFloat(width),
			),
		)
		gss.current.Stroke.Width = width
	}

	return gss
}

// SetLineCap sets the line cap style (J operator).
func (gss *GraphicsStateStack) SetLineCap(
	cap LineCap,
) *GraphicsStateStack {
	if gss.current.Stroke.Cap != cap {
		gss.operators.WriteString(
			fmt.Sprintf("%d J\n", cap),
		)
		gss.current.Stroke.Cap = cap
	}

	return gss
}

// SetLineJoin sets the line join style (j operator).
func (gss *GraphicsStateStack) SetLineJoin(
	join LineJoin,
) *GraphicsStateStack {
	if gss.current.Stroke.Join != join {
		gss.operators.WriteString(
			fmt.Sprintf("%d j\n", join),
		)
		gss.current.Stroke.Join = join
	}

	return gss
}

// SetMiterLimit sets the miter limit (M operator).
func (gss *GraphicsStateStack) SetMiterLimit(
	limit float64,
) *GraphicsStateStack {
	if gss.current.Stroke.MiterLimit != limit {
		gss.operators.WriteString(
			fmt.Sprintf(
				"%s M\n",
				formatFloat(limit),
			),
		)
		gss.current.Stroke.MiterLimit = limit
	}

	return gss
}

// SetDashPattern sets the dash pattern (d operator).
func (gss *GraphicsStateStack) SetDashPattern(
	dash DashPattern,
) *GraphicsStateStack {
	// Compare dash patterns
	needsUpdate := len(
		gss.current.Stroke.Dash.Array,
	) != len(
		dash.Array,
	) ||
		gss.current.Stroke.Dash.Phase != dash.Phase
	if !needsUpdate {
		for i := range dash.Array {
			if gss.current.Stroke.Dash.Array[i] != dash.Array[i] {
				needsUpdate = true

				break
			}
		}
	}
	if needsUpdate {
		gss.operators.WriteString(
			dash.ContentStream(),
		)
		gss.operators.WriteByte('\n')
		gss.current.Stroke.Dash = DashPattern{
			Array: append(
				[]float64(nil),
				dash.Array...),
			Phase: dash.Phase,
		}
	}

	return gss
}

// SetStrokeColor sets the stroke color (RG operator for RGB).
func (gss *GraphicsStateStack) SetStrokeColor(
	color Color,
) *GraphicsStateStack {
	if gss.current.Stroke.Color != color {
		gss.operators.WriteString(
			color.SetStrokeRGB(),
		)
		gss.operators.WriteByte('\n')
		gss.current.Stroke.Color = color
	}

	return gss
}

// SetFillColor sets the fill color (rg operator for RGB).
func (gss *GraphicsStateStack) SetFillColor(
	color Color,
) *GraphicsStateStack {
	if gss.current.Fill.Color != color {
		gss.operators.WriteString(
			color.SetFillRGB(),
		)
		gss.operators.WriteByte('\n')
		gss.current.Fill.Color = color
	}

	return gss
}

// SetStrokeStyle applies a complete stroke style.
func (gss *GraphicsStateStack) SetStrokeStyle(
	style *StrokeStyle,
) *GraphicsStateStack {
	if style == nil {
		return gss
	}
	gss.SetLineWidth(style.Width)
	gss.SetLineCap(style.Cap)
	gss.SetLineJoin(style.Join)
	if style.Join == LineJoinMiter {
		gss.SetMiterLimit(style.MiterLimit)
	}
	gss.SetDashPattern(style.Dash)
	gss.SetStrokeColor(style.Color)

	return gss
}

// SetFlatness sets the flatness tolerance (i operator).
func (gss *GraphicsStateStack) SetFlatness(
	flatness float64,
) *GraphicsStateStack {
	if gss.current.Flatness != flatness {
		gss.operators.WriteString(
			fmt.Sprintf(
				"%s i\n",
				formatFloat(flatness),
			),
		)
		gss.current.Flatness = flatness
	}

	return gss
}

// String returns the accumulated PDF operators.
func (gss *GraphicsStateStack) String() string {
	return gss.operators.String()
}

// Reset clears the operators and resets to initial state.
func (gss *GraphicsStateStack) Reset() *GraphicsStateStack {
	gss.stack = gss.stack[:0]
	gss.current = NewFullGraphicsState()
	gss.operators.Reset()
	gss.extGStateRefs = make(map[string]bool)

	return gss
}

// WriteOperator writes a raw operator to the content stream.
func (gss *GraphicsStateStack) WriteOperator(
	op string,
) *GraphicsStateStack {
	gss.operators.WriteString(op)
	gss.operators.WriteByte('\n')

	return gss
}

// WriteString writes arbitrary content to the stream.
func (gss *GraphicsStateStack) WriteString(
	content string,
) *GraphicsStateStack {
	gss.operators.WriteString(content)

	return gss
}

// ExtGStateEntry represents an entry in the extended graphics state dictionary.
// These correspond to entries in a PDF ExtGState resource dictionary.
type ExtGStateEntry struct {
	// Name is the resource name for this ExtGState (e.g., "GS1").
	Name string
	// StrokeAlpha is the stroke alpha (CA entry), nil means not set.
	StrokeAlpha *float64
	// FillAlpha is the fill alpha (ca entry), nil means not set.
	FillAlpha *float64
	// BlendMode is the blend mode (BM entry), empty means not set.
	BlendMode BlendMode
	// Overprint controls overprinting for stroking (OP entry).
	Overprint *bool
	// OverprintFill controls overprinting for filling (op entry).
	OverprintFill *bool
	// OverprintMode is the overprint mode (OPM entry).
	OverprintMode *int
	// SmoothnessTolerance is for smooth shading (SM entry).
	SmoothnessTolerance *float64
	// StrokeAdjustment controls stroke adjustment (SA entry).
	StrokeAdjustment *bool
	// AlphaIsShape controls alpha interpretation (AIS entry).
	AlphaIsShape *bool
	// TextKnockout controls text knockout (TK entry).
	TextKnockout *bool
}

// NewExtGStateEntry creates a new ExtGState entry with the given name.
func NewExtGStateEntry(
	name string,
) *ExtGStateEntry {
	return &ExtGStateEntry{
		Name: name,
	}
}

// WithStrokeAlpha sets the stroke alpha.
func (e *ExtGStateEntry) WithStrokeAlpha(
	alpha float64,
) *ExtGStateEntry {
	e.StrokeAlpha = &alpha

	return e
}

// WithFillAlpha sets the fill alpha.
func (e *ExtGStateEntry) WithFillAlpha(
	alpha float64,
) *ExtGStateEntry {
	e.FillAlpha = &alpha

	return e
}

// WithAlpha sets both stroke and fill alpha to the same value.
func (e *ExtGStateEntry) WithAlpha(
	alpha float64,
) *ExtGStateEntry {
	e.StrokeAlpha = &alpha
	e.FillAlpha = &alpha

	return e
}

// WithBlendMode sets the blend mode.
func (e *ExtGStateEntry) WithBlendMode(
	mode BlendMode,
) *ExtGStateEntry {
	e.BlendMode = mode

	return e
}

// WithOverprint sets overprint for stroking.
func (e *ExtGStateEntry) WithOverprint(
	on bool,
) *ExtGStateEntry {
	e.Overprint = &on

	return e
}

// WithOverprintFill sets overprint for filling.
func (e *ExtGStateEntry) WithOverprintFill(
	on bool,
) *ExtGStateEntry {
	e.OverprintFill = &on

	return e
}

// WithOverprintMode sets the overprint mode.
func (e *ExtGStateEntry) WithOverprintMode(
	mode int,
) *ExtGStateEntry {
	e.OverprintMode = &mode

	return e
}

// ToDictEntries returns the dictionary entries for this ExtGState.
// The returned map can be used to build a PDF dictionary object.
func (e *ExtGStateEntry) ToDictEntries() map[string]interface{} {
	dict := make(map[string]interface{})
	dict["Type"] = "/ExtGState"

	if e.StrokeAlpha != nil {
		dict["CA"] = *e.StrokeAlpha
	}
	if e.FillAlpha != nil {
		dict["ca"] = *e.FillAlpha
	}
	if e.BlendMode != "" &&
		e.BlendMode != BlendModeNormal {
		dict["BM"] = "/" + string(e.BlendMode)
	}
	if e.Overprint != nil {
		dict["OP"] = *e.Overprint
	}
	if e.OverprintFill != nil {
		dict["op"] = *e.OverprintFill
	}
	if e.OverprintMode != nil {
		dict["OPM"] = *e.OverprintMode
	}
	if e.SmoothnessTolerance != nil {
		dict["SM"] = *e.SmoothnessTolerance
	}
	if e.StrokeAdjustment != nil {
		dict["SA"] = *e.StrokeAdjustment
	}
	if e.AlphaIsShape != nil {
		dict["AIS"] = *e.AlphaIsShape
	}
	if e.TextKnockout != nil {
		dict["TK"] = *e.TextKnockout
	}

	return dict
}

// ApplyOperator returns the gs operator to apply this ExtGState.
func (e *ExtGStateEntry) ApplyOperator() string {
	return fmt.Sprintf("/%s gs", e.Name)
}

// ExtGStateManager manages extended graphics state dictionaries.
// It tracks created ExtGState resources and generates unique names.
type ExtGStateManager struct {
	entries   []*ExtGStateEntry
	nameIndex int
}

// NewExtGStateManager creates a new ExtGState manager.
func NewExtGStateManager() *ExtGStateManager {
	return &ExtGStateManager{}
}

// CreateAlpha creates an ExtGState for alpha transparency.
func (m *ExtGStateManager) CreateAlpha(
	strokeAlpha, fillAlpha float64,
) *ExtGStateEntry {
	m.nameIndex++
	entry := NewExtGStateEntry(
		fmt.Sprintf("GS%d", m.nameIndex),
	)
	if strokeAlpha < 1.0 {
		entry.WithStrokeAlpha(strokeAlpha)
	}
	if fillAlpha < 1.0 {
		entry.WithFillAlpha(fillAlpha)
	}
	m.entries = append(m.entries, entry)

	return entry
}

// CreateBlendMode creates an ExtGState for a blend mode.
func (m *ExtGStateManager) CreateBlendMode(
	mode BlendMode,
) *ExtGStateEntry {
	m.nameIndex++
	entry := NewExtGStateEntry(
		fmt.Sprintf("GS%d", m.nameIndex),
	).WithBlendMode(mode)
	m.entries = append(m.entries, entry)

	return entry
}

// Create creates a custom ExtGState entry.
func (m *ExtGStateManager) Create() *ExtGStateEntry {
	m.nameIndex++
	entry := NewExtGStateEntry(
		fmt.Sprintf("GS%d", m.nameIndex),
	)
	m.entries = append(m.entries, entry)

	return entry
}

// Entries returns all created ExtGState entries.
func (m *ExtGStateManager) Entries() []*ExtGStateEntry {
	return m.entries
}

// ResourceDict returns a map suitable for the ExtGState subdictionary
// in the page resources.
func (m *ExtGStateManager) ResourceDict() map[string]map[string]interface{} {
	if len(m.entries) == 0 {
		return nil
	}
	dict := make(
		map[string]map[string]interface{},
	)
	for _, entry := range m.entries {
		dict[entry.Name] = entry.ToDictEntries()
	}

	return dict
}

// SetExtGState applies an extended graphics state (gs operator).
func (gss *GraphicsStateStack) SetExtGState(
	entry *ExtGStateEntry,
) *GraphicsStateStack {
	gss.operators.WriteString(
		entry.ApplyOperator(),
	)
	gss.operators.WriteByte('\n')

	// Update internal state tracking
	if entry.StrokeAlpha != nil {
		gss.current.Transparency.StrokeAlpha = *entry.StrokeAlpha
	}
	if entry.FillAlpha != nil {
		gss.current.Transparency.FillAlpha = *entry.FillAlpha
	}
	if entry.BlendMode != "" {
		gss.current.Transparency.BlendMode = entry.BlendMode
	}
	if entry.AlphaIsShape != nil {
		gss.current.Transparency.AlphaIsShape = *entry.AlphaIsShape
	}

	gss.extGStateRefs[entry.Name] = true

	return gss
}

// SetAlpha sets stroke and fill alpha using an ExtGState.
// Returns the ExtGState entry that was created.
func (gss *GraphicsStateStack) SetAlpha(
	manager *ExtGStateManager,
	strokeAlpha, fillAlpha float64,
) *ExtGStateEntry {
	entry := manager.CreateAlpha(
		strokeAlpha,
		fillAlpha,
	)
	gss.SetExtGState(entry)

	return entry
}

// SetBlendMode sets the blend mode using an ExtGState.
// Returns the ExtGState entry that was created.
func (gss *GraphicsStateStack) SetBlendMode(
	manager *ExtGStateManager,
	mode BlendMode,
) *ExtGStateEntry {
	entry := manager.CreateBlendMode(mode)
	gss.SetExtGState(entry)

	return entry
}

// UsedExtGStates returns the names of ExtGState resources that have been used.
func (gss *GraphicsStateStack) UsedExtGStates() []string {
	result := make(
		[]string,
		0,
		len(gss.extGStateRefs),
	)
	for name := range gss.extGStateRefs {
		result = append(result, name)
	}

	return result
}

// WithState executes a function with a saved graphics state.
// The state is automatically restored after the function returns.
func (gss *GraphicsStateStack) WithState(
	fn func(*GraphicsStateStack),
) *GraphicsStateStack {
	gss.Save()
	fn(gss)
	gss.Restore()

	return gss
}

// WithTransform executes a function with a transform applied.
// The state is automatically restored after the function returns.
func (gss *GraphicsStateStack) WithTransform(
	t Transform,
	fn func(*GraphicsStateStack),
) *GraphicsStateStack {
	gss.Save()
	gss.ApplyTransform(t)
	fn(gss)
	gss.Restore()

	return gss
}

// WithTranslation executes a function with a translation applied.
func (gss *GraphicsStateStack) WithTranslation(
	tx, ty float64,
	fn func(*GraphicsStateStack),
) *GraphicsStateStack {
	return gss.WithTransform(
		Identity().Translate(tx, ty),
		fn,
	)
}

// WithScale executes a function with scaling applied.
func (gss *GraphicsStateStack) WithScale(
	sx, sy float64,
	fn func(*GraphicsStateStack),
) *GraphicsStateStack {
	return gss.WithTransform(
		Identity().Scale(sx, sy),
		fn,
	)
}

// WithRotation executes a function with rotation applied.
func (gss *GraphicsStateStack) WithRotation(
	degrees float64,
	fn func(*GraphicsStateStack),
) *GraphicsStateStack {
	return gss.WithTransform(
		Identity().Rotate(degrees),
		fn,
	)
}

// WithClip executes a function with a clipping path applied.
// The state is automatically restored after the function returns.
func (gss *GraphicsStateStack) WithClip(
	path *PathBuilder,
	rule ClipRule,
	fn func(*GraphicsStateStack),
) *GraphicsStateStack {
	gss.Save()

	// Write path operators
	gss.operators.WriteString(path.String())

	// Apply clipping
	gss.operators.WriteString(rule.String())
	gss.operators.WriteString(" n\n")

	gss.current.Clipping.Depth++

	fn(gss)
	gss.Restore()

	return gss
}

// StateScope provides RAII-style automatic state restoration.
// Use defer scope.Close() to ensure state is restored.
type StateScope struct {
	stack  *GraphicsStateStack
	closed bool
}

// NewStateScope creates a new state scope, saving the current state.
func NewStateScope(
	stack *GraphicsStateStack,
) *StateScope {
	stack.Save()

	return &StateScope{
		stack:  stack,
		closed: false,
	}
}

// Close restores the state. Safe to call multiple times.
func (s *StateScope) Close() {
	if !s.closed {
		s.stack.Restore()
		s.closed = true
	}
}

// TransformScope provides RAII-style state restoration with an initial transform.
type TransformScope struct {
	scope *StateScope
}

// NewTransformScope creates a new scope with a transform applied.
func NewTransformScope(
	stack *GraphicsStateStack,
	t Transform,
) *TransformScope {
	stack.Save()
	stack.ApplyTransform(t)

	return &TransformScope{
		scope: &StateScope{
			stack:  stack,
			closed: false,
		},
	}
}

// Close restores the state. Safe to call multiple times.
func (s *TransformScope) Close() {
	s.scope.Close()
}

// ClippingScope provides RAII-style state restoration with clipping.
// This differs from ClipScope in clip.go by working with GraphicsStateStack.
type ClippingScope struct {
	scope *StateScope
}

// NewClippingScope creates a new scope with a clipping path applied.
func NewClippingScope(
	stack *GraphicsStateStack,
	path *PathBuilder,
	rule ClipRule,
) *ClippingScope {
	stack.Save()

	// Write path operators
	stack.operators.WriteString(path.String())

	// Apply clipping
	stack.operators.WriteString(rule.String())
	stack.operators.WriteString(" n\n")

	stack.current.Clipping.Depth++

	return &ClippingScope{
		scope: &StateScope{
			stack:  stack,
			closed: false,
		},
	}
}

// Close restores the state. Safe to call multiple times.
func (s *ClippingScope) Close() {
	s.scope.Close()
}

// Common ExtGState presets for convenience.

// ExtGStateOpacity50 creates an ExtGState with 50% opacity.
func ExtGStateOpacity50(
	manager *ExtGStateManager,
) *ExtGStateEntry {
	return manager.CreateAlpha(0.5, 0.5)
}

// ExtGStateOpacity25 creates an ExtGState with 25% opacity.
func ExtGStateOpacity25(
	manager *ExtGStateManager,
) *ExtGStateEntry {
	return manager.CreateAlpha(0.25, 0.25)
}

// ExtGStateOpacity75 creates an ExtGState with 75% opacity.
func ExtGStateOpacity75(
	manager *ExtGStateManager,
) *ExtGStateEntry {
	return manager.CreateAlpha(0.75, 0.75)
}

// ExtGStateMultiply creates an ExtGState with multiply blend mode.
func ExtGStateMultiply(
	manager *ExtGStateManager,
) *ExtGStateEntry {
	return manager.CreateBlendMode(
		BlendModeMultiply,
	)
}

// ExtGStateScreen creates an ExtGState with screen blend mode.
func ExtGStateScreen(
	manager *ExtGStateManager,
) *ExtGStateEntry {
	return manager.CreateBlendMode(
		BlendModeScreen,
	)
}

// ExtGStateOverlay creates an ExtGState with overlay blend mode.
func ExtGStateOverlay(
	manager *ExtGStateManager,
) *ExtGStateEntry {
	return manager.CreateBlendMode(
		BlendModeOverlay,
	)
}
