package main

import (
	"fmt"
	. "github.com/arata-nvm/poly/poly"
	. "github.com/arata-nvm/poly/vecmath"
	"math"
)

func main() {
	m := LoadObj("examples/box3.obj")
	m.Rotation = Unit().MulScalar(10 * math.Pi / 180)
	m.Scale = Unit().MulScalar(0.5)

	size := 256

	d := NewDevice(size, size)

	c := NewCamera(NewVector3(2, 0, 10), Zero(), UnitY())
	d.SetCamera(c)

	d.Perspective(10, float64(size/size), 1, 10)

	cl := NewColor(0.5, 1, 0.6, 1)
	light := NewVector3(1, 1, 1)
	d.SetShader(NewFlatShader(cl, light))

	d.ClearColorBuffer(BLACK)
	d.DrawMesh(m)

	Save("out.png", d.Image())
}

func angleTest() {
	for a := 0; a < 90; a++ {
		m := LoadObj("examples/sphere.obj")
		m.Rotation = Unit().MulScalar(float64(a) * math.Pi / 180)
		m.Scale = Unit().MulScalar(0.5)

		size := 256

		d := NewDevice(size, size)

		c := NewCamera(NewVector3(2, 0, 10), Zero(), UnitY())
		d.SetCamera(c)

		d.Perspective(10, float64(size/size), 1, 10)

		cl := NewColor(0.5, 1, 0.6, 1)
		light := NewVector3(1, 1, 1)
		d.SetShader(NewFlatShader(cl, light))

		d.ClearColorBuffer(BLACK)
		d.DrawMesh(m)

		Save(fmt.Sprintf("out/out_%d.png", a), d.Image())
	}
}
