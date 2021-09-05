package poly

import (
	"image/color"
	"math"

	. "github.com/arata-nvm/poly/vecmath"
)

type Color struct {
	R, G, B, A float64
}

var (
	WHITE = Color{1, 1, 1, 1}
	BLACK = Color{0, 0, 0, 1}
)

func NewColor(r, g, b, a float64) Color {
	return Color{r, g, b, a}
}

func NewColorFromVec(v Vector3) Color {
	return NewColor(v.X, v.Y, v.Z, 1)
}

func (c1 Color) Add(c2 Color) Color {
	return Color{c1.R + c2.R, c1.G + c2.G, c1.B + c2.B, c1.A + c2.A}
}

func (c1 Color) Mul(c2 Color) Color {
	return Color{c1.R * c2.R, c1.G * c2.G, c1.B * c2.B, c1.A * c2.A}
}

func (c Color) MulScalar(f float64) Color {
	return Color{c.R * f, c.G * f, c.B * f, c.A}
}

func (c Color) Min(min Color) Color {
	return Color{
		math.Min(c.R, min.R),
		math.Min(c.G, min.G),
		math.Min(c.B, min.B),
		math.Min(c.A, min.A),
	}
}

func (c Color) NRGBA() color.NRGBA {
	cr := uint8(c.R * 255)
	cg := uint8(c.G * 255)
	cb := uint8(c.B * 255)
	ca := uint8(c.A * 255)
	return color.NRGBA{R: cr, G: cg, B: cb, A: ca}
}
