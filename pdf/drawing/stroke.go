package drawing

import (
	"fmt"
	"strings"
)

// LineCap specifies the shape to be used at the ends of open subpaths when they are stroked.
// This corresponds to the PDF J (setlinecap) operator.
type LineCap int

const (
	// LineCapButt cuts off the stroke at the endpoint of the path.
	// This is the default line cap style.
	LineCapButt LineCap = 0
	// LineCapRound ends the stroke with a semicircular arc.
	LineCapRound LineCap = 1
	// LineCapSquare extends the stroke beyond the endpoint by half the line width.
	LineCapSquare LineCap = 2
)

// String returns a human-readable name for the line cap style.
func (lc LineCap) String() string {
	switch lc {
	case LineCapButt:
		return "butt"
	case LineCapRound:
		return "round"
	case LineCapSquare:
		return "square"
	default:
		return fmt.Sprintf("LineCap(%d)", lc)
	}
}

// LineJoin specifies the shape to be used at the corners of stroked paths.
// This corresponds to the PDF j (setlinejoin) operator.
type LineJoin int

const (
	// LineJoinMiter extends the outer edges of the stroke until they meet.
	// This is the default line join style.
	LineJoinMiter LineJoin = 0
	// LineJoinRound joins the path segments with a circular arc.
	LineJoinRound LineJoin = 1
	// LineJoinBevel joins the path segments with a straight line.
	LineJoinBevel LineJoin = 2
)

// String returns a human-readable name for the line join style.
func (lj LineJoin) String() string {
	switch lj {
	case LineJoinMiter:
		return "miter"
	case LineJoinRound:
		return "round"
	case LineJoinBevel:
		return "bevel"
	default:
		return fmt.Sprintf("LineJoin(%d)", lj)
	}
}

// DashPattern represents a line dash pattern for stroking paths.
// The array specifies alternating dash and gap lengths.
// The phase specifies the distance into the dash pattern to start.
type DashPattern struct {
	// Array contains alternating dash and gap lengths.
	// An empty array indicates a solid line (no dashing).
	Array []float64
	// Phase is the distance into the dash pattern at which to start.
	Phase float64
}

// NewDashPattern creates a new dash pattern with the given array and phase.
func NewDashPattern(
	array []float64,
	phase float64,
) DashPattern {
	return DashPattern{
		Array: array,
		Phase: phase,
	}
}

// NewDashPatternSimple creates a new dash pattern with equal dash and gap lengths.
func NewDashPatternSimple(
	dashLength, gapLength float64,
) DashPattern {
	return DashPattern{
		Array: []float64{dashLength, gapLength},
		Phase: 0,
	}
}

// IsSolid returns true if this represents a solid line (no dashing).
func (d DashPattern) IsSolid() bool {
	return len(d.Array) == 0
}

// ContentStream returns the PDF d (setdash) operator for this dash pattern.
func (d DashPattern) ContentStream() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range d.Array {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(formatFloat(v))
	}
	sb.WriteString("] ")
	sb.WriteString(formatFloat(d.Phase))
	sb.WriteString(" d")

	return sb.String()
}

// Preset dash patterns matching common DrawingML/Office line styles.
var (
	// DashSolid is a solid line (no dashing).
	DashSolid = DashPattern{}

	// DashDot is a dotted line pattern.
	// DrawingML: prstDash val="dot"
	DashDot = DashPattern{
		Array: []float64{1, 2},
		Phase: 0,
	}

	// DashDash is a dashed line pattern.
	// DrawingML: prstDash val="dash"
	DashDash = DashPattern{
		Array: []float64{4, 3},
		Phase: 0,
	}

	// DashLongDash is a long dashed line pattern.
	// DrawingML: prstDash val="lgDash"
	DashLongDash = DashPattern{
		Array: []float64{8, 3},
		Phase: 0,
	}

	// DashDashDot is a dash-dot pattern.
	// DrawingML: prstDash val="dashDot"
	DashDashDot = DashPattern{
		Array: []float64{4, 3, 1, 3},
		Phase: 0,
	}

	// DashLongDashDot is a long dash-dot pattern.
	// DrawingML: prstDash val="lgDashDot"
	DashLongDashDot = DashPattern{
		Array: []float64{8, 3, 1, 3},
		Phase: 0,
	}

	// DashLongDashDotDot is a long dash-dot-dot pattern.
	// DrawingML: prstDash val="lgDashDotDot"
	DashLongDashDotDot = DashPattern{
		Array: []float64{8, 3, 1, 3, 1, 3},
		Phase: 0,
	}

	// DashSystemDot is a system-defined dotted pattern (small dots).
	// DrawingML: prstDash val="sysDot"
	DashSystemDot = DashPattern{
		Array: []float64{1, 1},
		Phase: 0,
	}

	// DashSystemDash is a system-defined dashed pattern.
	// DrawingML: prstDash val="sysDash"
	DashSystemDash = DashPattern{
		Array: []float64{3, 1},
		Phase: 0,
	}

	// DashSystemDashDot is a system-defined dash-dot pattern.
	// DrawingML: prstDash val="sysDashDot"
	DashSystemDashDot = DashPattern{
		Array: []float64{3, 1, 1, 1},
		Phase: 0,
	}

	// DashSystemDashDotDot is a system-defined dash-dot-dot pattern.
	// DrawingML: prstDash val="sysDashDotDot"
	DashSystemDashDotDot = DashPattern{
		Array: []float64{3, 1, 1, 1, 1, 1},
		Phase: 0,
	}
)

// DashPatternFromPreset returns a DashPattern for the given DrawingML preset name.
// If the preset is not recognized, a solid line pattern is returned.
func DashPatternFromPreset(
	preset string,
) DashPattern {
	switch preset {
	case "solid":
		return DashSolid
	case "dot":
		return DashDot
	case "dash":
		return DashDash
	case "lgDash":
		return DashLongDash
	case "dashDot":
		return DashDashDot
	case "lgDashDot":
		return DashLongDashDot
	case "lgDashDotDot":
		return DashLongDashDotDot
	case "sysDot":
		return DashSystemDot
	case "sysDash":
		return DashSystemDash
	case "sysDashDot":
		return DashSystemDashDot
	case "sysDashDotDot":
		return DashSystemDashDotDot
	default:
		return DashSolid
	}
}

// StrokeStyle combines all stroke properties into a single type.
// This matches the combined effect of DrawingML line properties.
type StrokeStyle struct {
	// Color is the stroke color.
	Color Color
	// Width is the line width in points.
	Width float64
	// Cap is the line cap style.
	Cap LineCap
	// Join is the line join style.
	Join LineJoin
	// MiterLimit is the miter limit for miter joins.
	// When the ratio of miter length to line width exceeds this limit,
	// a bevel join is used instead. Default is 10.
	MiterLimit float64
	// Dash is the dash pattern for the stroke.
	Dash DashPattern
}

// DefaultLineWidth is the default line width in PDF (1 user unit).
const DefaultLineWidth = 1.0

// DefaultMiterLimit is the default miter limit in PDF.
const DefaultMiterLimit = 10.0

// NewStrokeStyle creates a new stroke style with default settings.
// The default is a 1-point solid black line with butt caps and miter joins.
func NewStrokeStyle() *StrokeStyle {
	return &StrokeStyle{
		Color:      Black,
		Width:      DefaultLineWidth,
		Cap:        LineCapButt,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashSolid,
	}
}

// NewStrokeStyleWithColor creates a new stroke style with the specified color.
func NewStrokeStyleWithColor(
	color Color,
) *StrokeStyle {
	style := NewStrokeStyle()
	style.Color = color

	return style
}

// NewStrokeStyleWithWidth creates a new stroke style with the specified width.
func NewStrokeStyleWithWidth(
	width float64,
) *StrokeStyle {
	style := NewStrokeStyle()
	style.Width = width

	return style
}

// SetColor sets the stroke color.
func (s *StrokeStyle) SetColor(
	color Color,
) *StrokeStyle {
	s.Color = color

	return s
}

// SetWidth sets the line width.
func (s *StrokeStyle) SetWidth(
	width float64,
) *StrokeStyle {
	s.Width = width

	return s
}

// SetCap sets the line cap style.
func (s *StrokeStyle) SetCap(
	lineCap LineCap,
) *StrokeStyle {
	s.Cap = lineCap

	return s
}

// SetJoin sets the line join style.
func (s *StrokeStyle) SetJoin(
	join LineJoin,
) *StrokeStyle {
	s.Join = join

	return s
}

// SetMiterLimit sets the miter limit.
func (s *StrokeStyle) SetMiterLimit(
	limit float64,
) *StrokeStyle {
	s.MiterLimit = limit

	return s
}

// SetDash sets the dash pattern.
func (s *StrokeStyle) SetDash(
	dash DashPattern,
) *StrokeStyle {
	s.Dash = dash

	return s
}

// SetDashFromPreset sets the dash pattern from a DrawingML preset name.
func (s *StrokeStyle) SetDashFromPreset(
	preset string,
) *StrokeStyle {
	s.Dash = DashPatternFromPreset(preset)

	return s
}

// IsDefault returns true if this stroke style uses all default values.
func (s *StrokeStyle) IsDefault() bool {
	return s.Width == DefaultLineWidth &&
		s.Cap == LineCapButt &&
		s.Join == LineJoinMiter &&
		s.MiterLimit == DefaultMiterLimit &&
		s.Dash.IsSolid()
}

// ContentStream returns the PDF operators to apply this stroke style.
// This includes operators for line width, cap, join, miter limit, dash pattern, and color.
func (s *StrokeStyle) ContentStream() string {
	var sb strings.Builder

	// Set line width (w operator)
	sb.WriteString(SetLineWidth(s.Width))
	sb.WriteString("\n")

	// Set line cap (J operator)
	sb.WriteString(SetLineCap(s.Cap))
	sb.WriteString("\n")

	// Set line join (j operator)
	sb.WriteString(SetLineJoin(s.Join))
	sb.WriteString("\n")

	// Set miter limit (M operator) - only if using miter join
	if s.Join == LineJoinMiter &&
		s.MiterLimit != DefaultMiterLimit {
		sb.WriteString(
			SetMiterLimit(s.MiterLimit),
		)
		sb.WriteString("\n")
	}

	// Set dash pattern (d operator)
	if !s.Dash.IsSolid() {
		sb.WriteString(s.Dash.ContentStream())
		sb.WriteString("\n")
	}

	// Set stroke color (RG operator for RGB)
	sb.WriteString(s.Color.SetStrokeRGB())

	return sb.String()
}

// ContentStreamMinimal returns only the non-default stroke operators.
// Use this when you want to minimize content stream size.
func (s *StrokeStyle) ContentStreamMinimal() string {
	var parts []string

	// Only output non-default values
	if s.Width != DefaultLineWidth {
		parts = append(
			parts,
			SetLineWidth(s.Width),
		)
	}

	if s.Cap != LineCapButt {
		parts = append(parts, SetLineCap(s.Cap))
	}

	if s.Join != LineJoinMiter {
		parts = append(parts, SetLineJoin(s.Join))
	}

	if s.Join == LineJoinMiter &&
		s.MiterLimit != DefaultMiterLimit {
		parts = append(
			parts,
			SetMiterLimit(s.MiterLimit),
		)
	}

	if !s.Dash.IsSolid() {
		parts = append(
			parts,
			s.Dash.ContentStream(),
		)
	}

	// Always include color
	parts = append(parts, s.Color.SetStrokeRGB())

	return strings.Join(parts, "\n")
}

// Preset stroke styles for common use cases.
var (
	// StrokeBlackThin is a thin (0.5pt) black solid stroke.
	StrokeBlackThin = &StrokeStyle{
		Color:      Black,
		Width:      0.5,
		Cap:        LineCapButt,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashSolid,
	}

	// StrokeBlackMedium is a medium (1pt) black solid stroke.
	StrokeBlackMedium = &StrokeStyle{
		Color:      Black,
		Width:      1.0,
		Cap:        LineCapButt,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashSolid,
	}

	// StrokeBlackThick is a thick (2pt) black solid stroke.
	StrokeBlackThick = &StrokeStyle{
		Color:      Black,
		Width:      2.0,
		Cap:        LineCapButt,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashSolid,
	}

	// StrokeBlackDashed is a dashed black stroke.
	StrokeBlackDashed = &StrokeStyle{
		Color:      Black,
		Width:      1.0,
		Cap:        LineCapButt,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashDash,
	}

	// StrokeBlackDotted is a dotted black stroke.
	StrokeBlackDotted = &StrokeStyle{
		Color:      Black,
		Width:      1.0,
		Cap:        LineCapRound,
		Join:       LineJoinMiter,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashDot,
	}

	// StrokeBlackRounded is a stroke with round caps and joins.
	StrokeBlackRounded = &StrokeStyle{
		Color:      Black,
		Width:      1.0,
		Cap:        LineCapRound,
		Join:       LineJoinRound,
		MiterLimit: DefaultMiterLimit,
		Dash:       DashSolid,
	}
)

// PDF stroke operator functions

// SetLineWidth returns the PDF w (setlinewidth) operator.
func SetLineWidth(width float64) string {
	return fmt.Sprintf("%s w", formatFloat(width))
}

// SetLineCap returns the PDF J (setlinecap) operator.
func SetLineCap(lineCap LineCap) string {
	return fmt.Sprintf("%d J", lineCap)
}

// SetLineJoin returns the PDF j (setlinejoin) operator.
func SetLineJoin(join LineJoin) string {
	return fmt.Sprintf("%d j", join)
}

// SetMiterLimit returns the PDF M (setmiterlimit) operator.
func SetMiterLimit(limit float64) string {
	return fmt.Sprintf("%s M", formatFloat(limit))
}

// SetDashPattern returns the PDF d (setdash) operator for the given array and phase.
func SetDashPattern(
	array []float64,
	phase float64,
) string {
	return NewDashPattern(
		array,
		phase,
	).ContentStream()
}

// StrokePath applies a stroke style to a path and returns the combined content stream.
// This sets the stroke properties and appends the stroke operator.
func StrokePath(
	path *PathBuilder,
	style *StrokeStyle,
) string {
	if path == nil || style == nil {
		return ""
	}

	var sb strings.Builder

	// Set stroke style
	sb.WriteString(style.ContentStream())
	sb.WriteString("\n")

	// Add path and stroke
	sb.WriteString(path.Stroke())

	return sb.String()
}

// StrokePathMinimal applies only non-default stroke properties before stroking.
func StrokePathMinimal(
	path *PathBuilder,
	style *StrokeStyle,
) string {
	if path == nil || style == nil {
		return ""
	}

	var sb strings.Builder

	// Set only non-default stroke properties
	minimal := style.ContentStreamMinimal()
	if minimal != "" {
		sb.WriteString(minimal)
		sb.WriteString("\n")
	}

	// Add path and stroke
	sb.WriteString(path.Stroke())

	return sb.String()
}

// NoStroke represents no stroke (used for shape outlines).
type NoStroke struct{}

// ContentStream returns an empty string.
func (n NoStroke) ContentStream() string {
	return ""
}

// ScaledDashPattern returns a dash pattern scaled by the given factor.
// This is useful when line width affects how the dash should appear.
func (d DashPattern) ScaledDashPattern(
	scale float64,
) DashPattern {
	if scale == 1.0 || d.IsSolid() {
		return d
	}

	scaled := make([]float64, len(d.Array))
	for i, v := range d.Array {
		scaled[i] = v * scale
	}

	return DashPattern{
		Array: scaled,
		Phase: d.Phase * scale,
	}
}

// Clone returns a copy of the stroke style.
func (s *StrokeStyle) Clone() *StrokeStyle {
	return &StrokeStyle{
		Color:      s.Color,
		Width:      s.Width,
		Cap:        s.Cap,
		Join:       s.Join,
		MiterLimit: s.MiterLimit,
		Dash: DashPattern{
			Array: append(
				[]float64(nil),
				s.Dash.Array...),
			Phase: s.Dash.Phase,
		},
	}
}

// String returns a human-readable description of the stroke style.
func (s *StrokeStyle) String() string {
	dashStr := "solid"
	if !s.Dash.IsSolid() {
		dashStr = fmt.Sprintf(
			"dash[%v]",
			s.Dash.Array,
		)
	}

	return fmt.Sprintf(
		"StrokeStyle{color:%s, width:%.2f, cap:%s, join:%s, miter:%.1f, %s}",
		s.Color.String(),
		s.Width,
		s.Cap.String(),
		s.Join.String(),
		s.MiterLimit,
		dashStr,
	)
}
