package poly

import (
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
