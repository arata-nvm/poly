package poly

import (
	. "github.com/arata-nvm/poly/vecmath"
	"image"
)

type Device struct {
	Width  int
	Height int

	camera      Camera
	colorBuffer *image.NRGBA

	viewMatrix       Matrix4
	projectionMatrix Matrix4
}

func NewDevice(width, height int) *Device {
	return &Device{
		Width:       width,
		Height:      height,
		colorBuffer: image.NewNRGBA(image.Rect(0, 0, width, height)),
	}
}

func (d *Device) ClearColorBuffer(c Color) {
	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			d.colorBuffer.SetNRGBA(x, y, c.NRGBA())
		}
	}
}

func (d *Device) Image() image.Image {
	return d.colorBuffer
}

func (d *Device) SetCamera(c Camera) {
	d.camera = c
	d.viewMatrix = LookAt(c.Position, c.Target, c.Up)
}

func (d *Device) Perspective(fovy, aspect, near, far float64) {
	d.projectionMatrix = Perspective(fovy, aspect, near, far)
}

func (d *Device) PutPixel(x, y int, z float64,  c Color) {
	d.colorBuffer.Set(x, y, c.NRGBA())
}

func (d *Device) DrawPoint(v Vector3, c Color) {
	d.PutPixel(int(v.X), int(v.Y), v.Y, c)
}

// TODO v1 > v2
func (d *Device) DrawLine(v1, v2 Vector3, c Color) {
	x1 := int(v1.X)
	y1 := int(v1.Y)
	x2 := int(v2.X)
	y2 := int(v2.Y)

	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := sign(x2 - x1)
	sy := sign(x2 - x1)
	err := dx - dy

	gz := 1.0
	if dx > dy {
		gz /= float64(dx)
	} else {
		gz /= float64(dy)
	}

	for {
		var g float64
		if dx > dy {
			g = gz * float64(x2 - x1)
		} else {
			g = gz * float64(y2 - y1)
		}

		z := interpolate(v1.Z, v2.Z, 1 - g)

		d.PutPixel(x1, y1, z, c)

		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dx {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func (d *Device) DrawMesh(mesh *Mesh, c Color) {
	cx, cy := d.Width/2, d.Height/2
	tm := Translate(mesh.Position)
	rm := RotateX(mesh.Rotation.X).Mul(RotateY(mesh.Rotation.Y)).Mul(RotateZ(mesh.Rotation.Z))
	sm := Scale(mesh.Scale)
	modelMatrix := tm.Mul(rm).Mul(sm)

	transformMatrix := d.projectionMatrix.Mul(d.viewMatrix).Mul(modelMatrix)

	scale := float64(d.Width) / 2
	for i := range mesh.Vertices {
		mesh.Vertices[i].WorldCoordinates = TransformCoordinate(mesh.Vertices[i].Coordinates, transformMatrix)
		mesh.Vertices[i].WorldCoordinates.X = mesh.Vertices[i].WorldCoordinates.X*scale + float64(cx)
		mesh.Vertices[i].WorldCoordinates.Y = mesh.Vertices[i].WorldCoordinates.Y*scale + float64(cy)
	}

	for _, f := range mesh.Faces {
		v1 := mesh.Vertices[f.V1].WorldCoordinates
		v2 := mesh.Vertices[f.V2].WorldCoordinates
		v3 := mesh.Vertices[f.V3].WorldCoordinates
		d.DrawWiredTriangle(v1, v2, v3, c)
	}
}

func (d *Device) DrawWiredTriangle(v1, v2, v3 Vector3, c Color) {
	d.DrawLine(v1, v2, c)
	d.DrawLine(v2, v3, c)
	d.DrawLine(v3, v1, c)
}

// TODO y1 = y2 = y3
func (d *Device) DrawTriangle(v1, v2, v3 Vector3, c Color) {
	v1, v2, v3 = sortVectorsWithY(v1, v2, v3)
	// top
	for y := int(v1.Y); y < int(v2.Y); y++ {
		yf := float64(y)
		g1 := (yf-v1.Y)/(v3.Y-v1.Y)
		x1 := interpolate(v1.X, v3.X, g1)
		z1 := interpolate(v1.Z, v3.Z, g1)

		g2 := (yf-v1.Y)/(v2.Y-v1.Y)
		x2 := interpolate(v1.X, v2.X, g2)
		z2 := interpolate(v1.Z, v2.Z, g2)

		vd1 := NewVector3(x1, yf, z1)
		vd2 := NewVector3(x2, yf, z2)
		d.DrawLine(vd1, vd2, c)
	}

	// bottom
	for y := int(v2.Y); y < int(v3.Y); y++ {
		yf := float64(y)
		g1 := (yf-v1.Y)/(v3.Y-v1.Y)
		x1 := interpolate(v1.X, v3.X, g1)
		z1 := interpolate(v1.Z, v3.Z, g1)

		g2 := (yf-v2.Y)/(v3.Y-v2.Y)
		x2 := interpolate(v2.X, v3.X, g2)
		z2 := interpolate(v2.Z, v3.Z, g2)

		vd1 := NewVector3(x1, yf, z1)
		vd2 := NewVector3(x2, yf, z2)
		d.DrawLine(vd1, vd2, c)
	}
}

func sortVectorsWithY(v1, v2, v3 Vector3) (Vector3, Vector3, Vector3) {
	if v1.Y > v2.Y {
		v1, v2 = v2, v1
	}
	if v2.Y > v3.Y {
		v2, v3 = v3, v2
	}
	if v1.Y > v2.Y {
		v1, v2 = v2, v1
	}
	return v1, v2, v3
}
