package poly

import (
	"fmt"
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
	fmt.Printf("draw: (%d, %d)\n", x, y)
}

func (d *Device) DrawMesh(mesh Mesh, c Color) {
	scale := float64(d.Width) / 2
	cx, cy := d.Width/2, d.Height/2
	tm := Translate(mesh.Position)
	rm := RotateX(mesh.Rotation.X).Mul(RotateY(mesh.Rotation.Y)).Mul(RotateZ(mesh.Rotation.Z))
	sm := Scale(mesh.Scale)
	modelMatrix := tm.Mul(rm).Mul(sm)
	for _, v := range mesh.Vertices {
		v.WorldCoordinates = TransformCoordinate(v.Coordinates, modelMatrix)
		x := v.WorldCoordinates.X*scale + float64(cx)
		y := v.WorldCoordinates.Y*scale + float64(cy)
		d.PutPixel(int(x), int(y), c)
	}
}
