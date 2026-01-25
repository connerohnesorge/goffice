package drawing

import (
	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// FillRenderer is an interface for rendering fills.
type FillRenderer interface {
	Apply(
		ctx *core.RenderingContext,
		path *PathBuilder,
	) error
}

// SolidFillRenderer renders solid color fills.
type SolidFillRenderer struct {
	fill *drawingml.SolidFill
}

// NewSolidFillRenderer creates a new solid fill renderer.
func NewSolidFillRenderer(
	fill *drawingml.SolidFill,
) *SolidFillRenderer {
	return &SolidFillRenderer{fill: fill}
}

// Apply applies the solid fill.
func (r *SolidFillRenderer) Apply(
	ctx *core.RenderingContext,
	path *PathBuilder,
) error {
	if r.fill == nil {
		return nil
	}

	// Get the color
	var color Color
	if rgb := r.fill.RgbColor(); rgb != nil {
		// Convert RGB color
		hexColor, found := rgb.GetAttribute(
			"val",
			"",
		)
		if found {
			color = ParseColor(
				hexColor.Value(),
			)
		}
	} else if schemeColor := r.fill.SchemeColor(); schemeColor != nil {
		// Use scheme color (default to black for now)
		color = NewColor(0, 0, 0, 1)
	}

	// Apply the fill
	ctx.Page.SetFillColor(
		color.R,
		color.G,
		color.B,
	)
	ctx.Page.WriteContent(path.Fill())

	return nil
}

// GradientFillRenderer renders gradient fills.
type GradientFillRenderer struct {
	fill *drawingml.GradientFill
}

// NewGradientFillRenderer creates a new gradient fill renderer.
func NewGradientFillRenderer(
	fill *drawingml.GradientFill,
) *GradientFillRenderer {
	return &GradientFillRenderer{fill: fill}
}

// Apply applies the gradient fill.
func (r *GradientFillRenderer) Apply(
	ctx *core.RenderingContext,
	path *PathBuilder,
) error {
	if r.fill == nil {
		return nil
	}

	// For now, approximate gradient with solid fill using first stop color
	// Full gradient support would require shading patterns in PDF

	// Get gradient stops
	gsLst := r.fill.GetElement(
		"gsLst",
		drawingml.NamespaceMain,
	)
	if gsLst == nil {
		// No stops, use default color
		ctx.Page.SetFillColor(0.5, 0.5, 0.5)
		ctx.Page.WriteContent(path.Fill())

		return nil
	}

	// Get first stop
	if composite, ok := gsLst.(openxml.CompositeElement); ok {
		for child := range composite.Children() {
			if child.LocalName() == "gs" &&
				child.NamespaceURI() == drawingml.NamespaceMain {
				// Get color from first stop
				if childComposite, ok := child.(openxml.CompositeElement); ok {
					for colorChild := range childComposite.Children() {
						if colorChild.LocalName() == "srgbClr" {
							hexAttr, found := colorChild.GetAttribute(
								"val",
								"",
							)
							if found {
								color := ParseColor(
									hexAttr.Value(),
								)
								ctx.Page.SetFillColor(
									color.R,
									color.G,
									color.B,
								)
								ctx.Page.WriteContent(
									path.Fill(),
								)

								return nil
							}
						}
					}
				}

				break
			}
		}
	}

	// Default fallback
	ctx.Page.SetFillColor(0.5, 0.5, 0.5)
	ctx.Page.WriteContent(path.Fill())

	return nil
}

// PatternFillRenderer renders pattern fills.
type PatternFillRenderer struct {
	fill *drawingml.PatternFill
}

// NewPatternFillRenderer creates a new pattern fill renderer.
func NewPatternFillRenderer(
	fill *drawingml.PatternFill,
) *PatternFillRenderer {
	return &PatternFillRenderer{fill: fill}
}

// Apply applies the pattern fill.
func (r *PatternFillRenderer) Apply(
	ctx *core.RenderingContext,
	path *PathBuilder,
) error {
	if r.fill == nil {
		return nil
	}

	// Pattern fills in PDF require tiling patterns
	// For basic implementation, use foreground color as solid fill

	fgClr := r.fill.GetElement(
		"fgClr",
		drawingml.NamespaceMain,
	)
	if fgClr != nil {
		if composite, ok := fgClr.(openxml.CompositeElement); ok {
			for child := range composite.Children() {
				if child.LocalName() == "srgbClr" {
					hexAttr, found := child.GetAttribute(
						"val",
						"",
					)
					if found {
						color := ParseColor(
							hexAttr.Value(),
						)
						ctx.Page.SetFillColor(
							color.R,
							color.G,
							color.B,
						)
						ctx.Page.WriteContent(
							path.Fill(),
						)

						return nil
					}
				}
			}
		}
	}

	// Default to gray
	ctx.Page.SetFillColor(0.75, 0.75, 0.75)
	ctx.Page.WriteContent(path.Fill())

	return nil
}

// PictureFillRenderer renders picture/image fills.
type PictureFillRenderer struct {
	fill *drawingml.BlipFill
}

// NewPictureFillRenderer creates a new picture fill renderer.
func NewPictureFillRenderer(
	fill *drawingml.BlipFill,
) *PictureFillRenderer {
	return &PictureFillRenderer{fill: fill}
}

// Apply applies the picture fill.
func (r *PictureFillRenderer) Apply(
	ctx *core.RenderingContext,
	path *PathBuilder,
) error {
	if r.fill == nil {
		return nil
	}

	if ctx == nil || ctx.Page == nil {
		return nil
	}

	bounds, ok := path.Bounds()
	if !ok || bounds.Width == 0 || bounds.Height == 0 {
		return nil
	}

	if ctx.ImageResolver == nil {
		renderPictureFillPlaceholder(ctx, path)
		return nil
	}

	imageRenderer := NewImageRenderer(ctx)
	err := imageRenderer.RenderPicture(
		r.fill,
		RenderBounds{
			X:      bounds.X,
			Y:      bounds.Y,
			Width:  bounds.Width,
			Height: bounds.Height,
		},
	)
	if err != nil {
		renderPictureFillPlaceholder(ctx, path)
	}

	return nil
}

func renderPictureFillPlaceholder(
	ctx *core.RenderingContext,
	path *PathBuilder,
) {
	ctx.Page.SetFillColor(0.9, 0.9, 1.0)
	ctx.Page.WriteContent(path.Fill())
}

// NoFillRenderer represents no fill (transparent).
type NoFillRenderer struct{}

// NewNoFillRenderer creates a new no-fill renderer.
func NewNoFillRenderer() *NoFillRenderer {
	return &NoFillRenderer{}
}

// Apply applies no fill (does nothing).
func (r *NoFillRenderer) Apply(
	ctx *core.RenderingContext,
	path *PathBuilder,
) error {
	// No fill - just return
	return nil
}

// CreateFillRenderer creates appropriate fill renderer from shape properties.
func CreateFillRenderer(
	sp *drawingml.ShapeProperties,
) FillRenderer {
	if sp == nil {
		return nil
	}

	// Check for no fill
	if sp.GetElement(
		"noFill",
		drawingml.NamespaceMain,
	) != nil {
		return NewNoFillRenderer()
	}

	// Check for solid fill
	if solidFillElem := sp.GetElement("solidFill", drawingml.NamespaceMain); solidFillElem != nil {
		if solidFill, ok := solidFillElem.(*drawingml.SolidFill); ok {
			return NewSolidFillRenderer(solidFill)
		}
	}

	// Check for gradient fill
	if gradFillElem := sp.GetElement("gradFill", drawingml.NamespaceMain); gradFillElem != nil {
		if gradFill, ok := gradFillElem.(*drawingml.GradientFill); ok {
			return NewGradientFillRenderer(
				gradFill,
			)
		}
	}

	// Check for pattern fill
	if pattFillElem := sp.GetElement("pattFill", drawingml.NamespaceMain); pattFillElem != nil {
		if pattFill, ok := pattFillElem.(*drawingml.PatternFill); ok {
			return NewPatternFillRenderer(
				pattFill,
			)
		}
	}

	// Check for blip fill (image/picture)
	if blipFillElem := sp.GetElement("blipFill", drawingml.NamespaceMain); blipFillElem != nil {
		if blipFill, ok := blipFillElem.(*drawingml.BlipFill); ok {
			return NewPictureFillRenderer(
				blipFill,
			)
		}
	}

	// Default: no fill
	return NewNoFillRenderer()
}
