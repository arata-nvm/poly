package poly

import (
	"fmt"
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

func (d *Device) DrawMesh(mesh Mesh, c Color) {
	scale := float64(d.Width) * 0.8 / 2
	cx, cy := d.Width / 2, d.Height / 2
	for _, v := range mesh.Vertices {
		x := v.X * scale + float64(cx)
		y := v.Y * scale + float64(cy)
		d.PutPixel(int(x), int(y), c)
	}
}

