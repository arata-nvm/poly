package poly

import (
	. "github.com/arata-nvm/poly/vecmath"
	"image"
)

type Device struct {
	Width  int
	Height int

	ColorBuffer *image.NRGBA
}

func NewDevice(width, height int) *Device {
	return &Device{
		Width:       width,
		Height:      height,
		ColorBuffer: image.NewNRGBA(image.Rect(0, 0, width, height)),
	}
}

func (d *Device) ClearColorBuffer(c Color) {
	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			d.ColorBuffer.SetNRGBA(x, y, c.NRGBA())
		}
	}
}

func (d *Device) Image() image.Image {
	return d.ColorBuffer
}

func (d *Device) PutPixel(x, y int, c Color) {
	d.ColorBuffer.Set(x, y, c.NRGBA())
}

func (d *Device) DrawLine(v1, v2 Vector3, c Color) {
	x1 := int(v1.X)
	y1 := int(v1.Y)
	x2 := int(v2.X)
	y2 := int(v2.Y)

	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := sign(x2 - x1)
	sy := sign(y2 - y1)
	err := dx - dy

	for {
		d.PutPixel(x1, y1, c)

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

func (d *Device) DrawMesh(camera Camera, mesh Mesh, c Color) {
	viewMatrix := LookAt(camera.Position, camera.Target, camera.Up)

	cx, cy := d.Width/2, d.Height/2
	tm := Translate(mesh.Position)
	rm := RotateX(mesh.Rotation.X).Mul(RotateY(mesh.Rotation.Y)).Mul(RotateZ(mesh.Rotation.Z))
	sm := Scale(mesh.Scale)
	modelMatrix := tm.Mul(rm).Mul(sm)

	transformMatrix := viewMatrix.Mul(modelMatrix)

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
		d.DrawLine(v1, v2, c)
		d.DrawLine(v2, v3, c)
		d.DrawLine(v3, v1, c)
	}
}
