package main

import (
	"fmt"
	. "github.com/arata-nvm/poly/poly"
	. "github.com/arata-nvm/poly/vecmath"
	"math"
)

func main() {
	size := 1024
	d := NewDevice(size, size)
	d.ClearColorBuffer(BLACK)

	eye := NewVector3(1,  1, 10)
	target := NewVector3(0, 0, -10)
	c := NewCamera(eye, target, UnitY())
	d.SetCamera(c)

	d.Perspective(10, float64(size/size), 1, 10)

	cl := NewColor(0.5, 1, 0.6, 1)
	light := NewVector3(1, 1 ,0.1)
	shader := NewPhongShader(light.Normalize(), eye.Normalize(), cl, 100)
	d.SetShader(shader)
	d.SetShader(NewFlatShader(cl, light))

	scene := Scene1()
	for i := range scene {
		shader.Color.R = 255 * float64(i) / float64(len(scene))
		d.DrawMesh(scene[i])
	}

	Save("out.png", d.Image())
}

func Scene1() []*Mesh{
	n := 10
	meshes := make([]*Mesh, n)

	// balls
	for i := 0; i < n; i++ {
		box := LoadObj("examples/sphere.obj")
		//box.SmoothNormals()
		box.Scale = Unit().MulScalar(0.4)
		box.Position.Z = -float64(i) * float64(i)

		meshes[i] = box
	}

	// floor
	floor := LoadObj("examples/plane.obj")
	floor.Position.Y = -0.5
	floor.Position.Z = -float64(n * n)
	floor.Scale.X = 11
	floor.Scale.Z = float64(n * n)
	meshes = append(meshes, floor)

	return meshes
}

func angleTest() {
	wall := LoadObj("examples/box3.obj")

	for a := 0; a < 90; a++ {
		m := NewPlane() //LoadObj("examples/plane.obj")
		m.Rotation = Unit().MulScalar(float64(a) * math.Pi / 180)
		m.Scale = Unit().MulScalar(0.5)

		wall.Position = NewVector3(0, 0, 9)

		size := 256

		d := NewDevice(size, size)

		eye := NewVector3(2, 0, 10)
		c := NewCamera(eye, Zero(), UnitY())
		d.SetCamera(c)

		d.Perspective(10, float64(size/size), 1, 10)

		cl := NewColor(0.5, 1, 0.6, 1)
		light := NewVector3(1, 1, 1)
		d.SetShader(NewPhongShader(light.Normalize(), eye.Normalize(), cl, 100))

		d.ClearColorBuffer(BLACK)
		d.DrawMesh(m)

		d.DrawMesh(wall)

		Save(fmt.Sprintf("out/out_%d.png", a), d.Image())
	}
}
