package poly

import (
	"image/color"
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

func (c Color) NRGBA() color.NRGBA {
	cr := uint8(c.R * 255)
	cg := uint8(c.G * 255)
	cb := uint8(c.B * 255)
	ca := uint8(c.A * 255)
	return color.NRGBA{R: cr, G: cg, B: cb, A: ca}
}
