package main

import (
	. "github.com/arata-nvm/poly/poly"
	. "github.com/arata-nvm/poly/vecmath"
	"math"
)

func main() {
	m := LoadObj("examples/box3.obj")
	m.Rotation = Unit().MulScalar(90 * math.Pi / 180)
	m.Scale = Unit().MulScalar(0.5)

	size := 256

	d := NewDevice(size, size)

	c := NewCamera(NewVector3(2, 0, 10), Zero(), UnitY())
	d.SetCamera(c)

	d.Perspective(10, float64(size/size), 1, 10)

	d.ClearColorBuffer(BLACK)
	cl := NewColor(0.5, 1, 0.6, 1)
	d.DrawMesh(m, cl)

	Save("out.png", d.Image())
}
