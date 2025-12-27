package drawing

import (
	"math"
)

// applyTransformations applies a sequence of color transformations to a base color.
// Transformations are applied in the order they appear in the slice.
func applyTransformations(
	base Color,
	transforms []ColorTransformation,
) Color {
	result := base

	for _, t := range transforms {
		result = applyTransformation(result, t)
	}

	return result
}

// applyTransformation applies a single color transformation.
func applyTransformation(
	color Color,
	trans ColorTransformation,
) Color {
	switch trans.Type {
	case "tint":
		return ApplyTint(color, trans.Value)
	case "shade":
		return ApplyShade(color, trans.Value)
	case "alpha":
		return ApplyAlpha(color, trans.Value)
	case "alphaOff":
		return ApplyAlphaOffset(
			color,
			trans.Value,
		)
	case "alphaMod":
		return ApplyAlphaModulation(
			color,
			trans.Value,
		)
	case "satMod":
		return ApplySaturationModulation(
			color,
			trans.Value,
		)
	case "satOff":
		return ApplySaturationOffset(
			color,
			trans.Value,
		)
	case "sat":
		return ApplySaturation(color, trans.Value)
	case "lumMod":
		return ApplyLuminanceModulation(
			color,
			trans.Value,
		)
	case "lumOff":
		return ApplyLuminanceOffset(
			color,
			trans.Value,
		)
	case "lum":
		return ApplyLuminance(color, trans.Value)
	case "hueMod":
		return ApplyHueModulation(
			color,
			trans.Value,
		)
	case "hueOff":
		return ApplyHueOffset(color, trans.Value)
	case "hue":
		return ApplyHue(color, trans.Value)
	case "red":
		return ApplyRed(color, trans.Value)
	case "redOff":
		return ApplyRedOffset(color, trans.Value)
	case "redMod":
		return ApplyRedModulation(
			color,
			trans.Value,
		)
	case "green":
		return ApplyGreen(color, trans.Value)
	case "greenOff":
		return ApplyGreenOffset(
			color,
			trans.Value,
		)
	case "greenMod":
		return ApplyGreenModulation(
			color,
			trans.Value,
		)
	case "blue":
		return ApplyBlue(color, trans.Value)
	case "blueOff":
		return ApplyBlueOffset(color, trans.Value)
	case "blueMod":
		return ApplyBlueModulation(
			color,
			trans.Value,
		)
	case "comp":
		return ApplyComplement(color)
	case "inv":
		return ApplyInverse(color)
	case "gray":
		return ApplyGrayscale(color)
	case "gamma":
		return ApplyGamma(color)
	case "invGamma":
		return ApplyInverseGamma(color)
	default:
		return color
	}
}

// ApplyTint applies a tint transformation (add white) to the color.
// value is in 1000ths of a percent (0-100000).
// At value=0, no change; at value=100000, the color becomes white.
func ApplyTint(color Color, value int) Color {
	// Convert value to 0-1 range.
	amount := float64(value) / 100000.0

	// Tint formula: newColor = color + (1 - color) * amount
	// This moves each channel toward 1 (white).
	return NewColor(
		color.R+(1-color.R)*amount,
		color.G+(1-color.G)*amount,
		color.B+(1-color.B)*amount,
		color.A,
	)
}

// ApplyShade applies a shade transformation (add black) to the color.
// value is in 1000ths of a percent (0-100000).
// At value=100000, no change; at value=0, the color becomes black.
func ApplyShade(color Color, value int) Color {
	// Convert value to 0-1 range.
	// Shade at 100% means original color, 0% means black.
	factor := float64(value) / 100000.0

	// Shade formula: newColor = color * factor
	// This moves each channel toward 0 (black).
	return NewColor(
		color.R*factor,
		color.G*factor,
		color.B*factor,
		color.A,
	)
}

// ApplyAlpha sets the alpha (transparency) of the color.
// value is in 1000ths of a percent (0-100000).
// 0 = fully transparent, 100000 = fully opaque.
func ApplyAlpha(color Color, value int) Color {
	alpha := float64(value) / 100000.0

	return color.WithAlpha(alpha)
}

// ApplyAlphaOffset offsets the alpha by a percentage.
// value is in 1000ths of a percent (-100000 to 100000).
func ApplyAlphaOffset(
	color Color,
	value int,
) Color {
	offset := float64(value) / 100000.0
	newAlpha := clamp01(color.A + offset)

	return color.WithAlpha(newAlpha)
}

// ApplyAlphaModulation modulates the alpha by a percentage.
// value is in 1000ths of a percent (0-100000+).
// 100000 = no change, 50000 = half alpha, 200000 = double alpha.
func ApplyAlphaModulation(
	color Color,
	value int,
) Color {
	factor := float64(value) / 100000.0
	newAlpha := clamp01(color.A * factor)

	return color.WithAlpha(newAlpha)
}

// ApplySaturationModulation modulates the saturation by a percentage.
// value is in 1000ths of a percent (0-100000+).
// 100000 = no change, 50000 = half saturation, 200000 = double saturation.
func ApplySaturationModulation(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	factor := float64(value) / 100000.0
	hsl.S = clamp01(hsl.S * factor)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplySaturationOffset offsets the saturation by a percentage.
// value is in 1000ths of a percent (-100000 to 100000).
func ApplySaturationOffset(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	offset := float64(value) / 100000.0
	hsl.S = clamp01(hsl.S + offset)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplySaturation sets the saturation to a specific value.
// value is in 1000ths of a percent (0-100000).
func ApplySaturation(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	hsl.S = clamp01(float64(value) / 100000.0)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyLuminanceModulation modulates the luminance by a percentage.
// value is in 1000ths of a percent (0-100000+).
// 100000 = no change, 50000 = half luminance, 200000 = double luminance.
func ApplyLuminanceModulation(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	factor := float64(value) / 100000.0
	hsl.L = clamp01(hsl.L * factor)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyLuminanceOffset offsets the luminance by a percentage.
// value is in 1000ths of a percent (-100000 to 100000).
func ApplyLuminanceOffset(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	offset := float64(value) / 100000.0
	hsl.L = clamp01(hsl.L + offset)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyLuminance sets the luminance to a specific value.
// value is in 1000ths of a percent (0-100000).
func ApplyLuminance(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	hsl.L = clamp01(float64(value) / 100000.0)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyHueModulation modulates the hue by a percentage.
// value is in 1000ths of a percent.
// 100000 = no change.
func ApplyHueModulation(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	factor := float64(value) / 100000.0
	hsl.H = math.Mod(hsl.H*factor, 360)
	if hsl.H < 0 {
		hsl.H += 360
	}
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyHueOffset offsets the hue by an angle.
// value is in 60000ths of a degree.
func ApplyHueOffset(
	color Color,
	value int,
) Color {
	hsl := RGBToHSL(color)
	offset := float64(
		value,
	) / 60000.0 // Convert to degrees
	hsl.H = math.Mod(hsl.H+offset, 360)
	if hsl.H < 0 {
		hsl.H += 360
	}
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyHue sets the hue to a specific value.
// value is in 60000ths of a degree (0-21600000).
func ApplyHue(color Color, value int) Color {
	hsl := RGBToHSL(color)
	hsl.H = math.Mod(
		float64(value)/60000.0,
		360,
	) // Convert to degrees
	if hsl.H < 0 {
		hsl.H += 360
	}
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyRed sets the red component to a specific value.
// value is in 1000ths of a percent (0-100000).
func ApplyRed(color Color, value int) Color {
	r := clamp01(float64(value) / 100000.0)

	return NewColor(r, color.G, color.B, color.A)
}

// ApplyRedOffset offsets the red component.
// value is in 1000ths of a percent (-100000 to 100000).
func ApplyRedOffset(
	color Color,
	value int,
) Color {
	offset := float64(value) / 100000.0
	r := clamp01(color.R + offset)

	return NewColor(r, color.G, color.B, color.A)
}

// ApplyRedModulation modulates the red component.
// value is in 1000ths of a percent.
func ApplyRedModulation(
	color Color,
	value int,
) Color {
	factor := float64(value) / 100000.0
	r := clamp01(color.R * factor)

	return NewColor(r, color.G, color.B, color.A)
}

// ApplyGreen sets the green component to a specific value.
// value is in 1000ths of a percent (0-100000).
func ApplyGreen(color Color, value int) Color {
	g := clamp01(float64(value) / 100000.0)

	return NewColor(color.R, g, color.B, color.A)
}

// ApplyGreenOffset offsets the green component.
// value is in 1000ths of a percent (-100000 to 100000).
func ApplyGreenOffset(
	color Color,
	value int,
) Color {
	offset := float64(value) / 100000.0
	g := clamp01(color.G + offset)

	return NewColor(color.R, g, color.B, color.A)
}

// ApplyGreenModulation modulates the green component.
// value is in 1000ths of a percent.
func ApplyGreenModulation(
	color Color,
	value int,
) Color {
	factor := float64(value) / 100000.0
	g := clamp01(color.G * factor)

	return NewColor(color.R, g, color.B, color.A)
}

// ApplyBlue sets the blue component to a specific value.
// value is in 1000ths of a percent (0-100000).
func ApplyBlue(color Color, value int) Color {
	b := clamp01(float64(value) / 100000.0)

	return NewColor(color.R, color.G, b, color.A)
}

// ApplyBlueOffset offsets the blue component.
// value is in 1000ths of a percent (-100000 to 100000).
func ApplyBlueOffset(
	color Color,
	value int,
) Color {
	offset := float64(value) / 100000.0
	b := clamp01(color.B + offset)

	return NewColor(color.R, color.G, b, color.A)
}

// ApplyBlueModulation modulates the blue component.
// value is in 1000ths of a percent.
func ApplyBlueModulation(
	color Color,
	value int,
) Color {
	factor := float64(value) / 100000.0
	b := clamp01(color.B * factor)

	return NewColor(color.R, color.G, b, color.A)
}

// ApplyComplement returns the complementary color (opposite hue).
func ApplyComplement(color Color) Color {
	hsl := RGBToHSL(color)
	hsl.H = math.Mod(hsl.H+180, 360)
	result := hsl.ToRGB()

	return result.WithAlpha(color.A)
}

// ApplyInverse returns the inverse color (1 - each RGB component).
func ApplyInverse(color Color) Color {
	return NewColor(
		1-color.R,
		1-color.G,
		1-color.B,
		color.A,
	)
}

// ApplyGrayscale converts the color to grayscale.
// Uses the luminosity method (weighted average based on human perception).
func ApplyGrayscale(color Color) Color {
	// Use perceptual weights for grayscale conversion.
	gray := 0.299*color.R + 0.587*color.G + 0.114*color.B

	return NewColor(gray, gray, gray, color.A)
}

// ApplyGamma applies gamma correction to the color.
// Uses a standard gamma of 2.2.
func ApplyGamma(color Color) Color {
	const gamma = 2.2

	return NewColor(
		math.Pow(color.R, gamma),
		math.Pow(color.G, gamma),
		math.Pow(color.B, gamma),
		color.A,
	)
}

// ApplyInverseGamma applies inverse gamma correction to the color.
// Uses a standard gamma of 2.2.
func ApplyInverseGamma(color Color) Color {
	const invGamma = 1.0 / 2.2

	return NewColor(
		math.Pow(color.R, invGamma),
		math.Pow(color.G, invGamma),
		math.Pow(color.B, invGamma),
		color.A,
	)
}

// HSVColor represents a color in HSV (Hue, Saturation, Value) space.
// This is different from HSL and is sometimes used for color manipulations.
type HSVColor struct {
	H, S, V float64 // H: 0-360, S: 0-1, V: 0-1
}

// NewHSV creates a new HSV color.
// H is in degrees (0-360), S and V are in the range [0, 1].
func NewHSV(h, s, v float64) HSVColor {
	// Normalize hue to 0-360.
	for h < 0 {
		h += 360
	}
	for h >= 360 {
		h -= 360
	}

	return HSVColor{
		H: h,
		S: clamp01(s),
		V: clamp01(v),
	}
}

// ToRGB converts an HSV color to an RGB color.
func (c HSVColor) ToRGB() Color {
	if c.S == 0 {
		// Achromatic (gray).
		return NewRGB(c.V, c.V, c.V)
	}

	h := c.H / 60.0 // Sector 0-5.
	i := int(math.Floor(h))
	f := h - float64(i) // Fractional part.

	p := c.V * (1 - c.S)
	q := c.V * (1 - c.S*f)
	t := c.V * (1 - c.S*(1-f))

	var r, g, b float64
	switch i {
	case 0:
		r, g, b = c.V, t, p
	case 1:
		r, g, b = q, c.V, p
	case 2:
		r, g, b = p, c.V, t
	case 3:
		r, g, b = p, q, c.V
	case 4:
		r, g, b = t, p, c.V
	default: // case 5
		r, g, b = c.V, p, q
	}

	return NewRGB(r, g, b)
}

// RGBToHSV converts an RGB color to an HSV color.
func RGBToHSV(color Color) HSVColor {
	maxC := math.Max(
		color.R,
		math.Max(color.G, color.B),
	)
	minC := math.Min(
		color.R,
		math.Min(color.G, color.B),
	)
	delta := maxC - minC

	// Value is the maximum component.
	v := maxC

	// If max is 0, color is black.
	if maxC == 0 {
		return NewHSV(0, 0, 0)
	}

	// Saturation is delta / max.
	s := delta / maxC

	// If there's no saturation, hue is undefined.
	if delta == 0 {
		return NewHSV(0, 0, v)
	}

	// Calculate hue.
	var h float64
	switch maxC {
	case color.R:
		h = (color.G - color.B) / delta
		if color.G < color.B {
			h += 6
		}
	case color.G:
		h = (color.B-color.R)/delta + 2
	case color.B:
		h = (color.R-color.G)/delta + 4
	}
	h *= 60

	return NewHSV(h, s, v)
}

// Interpolate linearly interpolates between two colors.
// t should be in the range [0, 1], where 0 returns the first color
// and 1 returns the second color.
func Interpolate(c1, c2 Color, t float64) Color {
	t = clamp01(t)

	return NewColor(
		c1.R+(c2.R-c1.R)*t,
		c1.G+(c2.G-c1.G)*t,
		c1.B+(c2.B-c1.B)*t,
		c1.A+(c2.A-c1.A)*t,
	)
}

// InterpolateHSL linearly interpolates between two colors in HSL space.
// This often produces more visually pleasing gradients than RGB interpolation.
func InterpolateHSL(
	c1, c2 Color,
	t float64,
) Color {
	t = clamp01(t)

	hsl1 := RGBToHSL(c1)
	hsl2 := RGBToHSL(c2)

	// Handle hue wraparound (take the shorter path around the color wheel).
	dh := hsl2.H - hsl1.H
	if dh > 180 {
		dh -= 360
	} else if dh < -180 {
		dh += 360
	}

	h := hsl1.H + dh*t
	if h < 0 {
		h += 360
	} else if h >= 360 {
		h -= 360
	}

	s := hsl1.S + (hsl2.S-hsl1.S)*t
	l := hsl1.L + (hsl2.L-hsl1.L)*t
	a := c1.A + (c2.A-c1.A)*t

	result := NewHSL(h, s, l).ToRGB()

	return result.WithAlpha(a)
}
