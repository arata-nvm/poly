package main

import (
	. "github.com/arata-nvm/poly/poly"
	. "github.com/arata-nvm/poly/vecmath"
)

func main() {
	size := 1024
	d := NewDevice(size, size)
	d.ClearColorBuffer(BLACK)

	eye := NewVector3(-1, 0, 1)
	target := NewVector3(0, 0.1, 0)
	c := NewCamera(eye, target, UnitY())
	d.SetCamera(c)

	d.Perspective(10, 1.0, 1, 10)

	cl := NewColor(0.5, 1, 0.6, 1)
	light := NewVector3(1, 0, 1)
	shader := NewPhongShader(light, eye, cl, 64)
	d.SetShader(shader)

	bunny := LoadPly("examples/bunny/reconstruction/bun_zipper.ply")
	bunny.CalcNormal()
	bunny.SmoothNormals()
	d.DrawMesh(bunny)

	Save("out.png", d.Image())
}
