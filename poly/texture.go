package poly

import (
	"image"
	"os"
)

type Texture struct {
	Width  int
	Height int

	Image image.Image
}

func NewTexture(filename string) (*Texture, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	return &Texture{
		Width:  rect.Max.X,
		Height: rect.Max.Y,
		Image:  img,
	}, nil
}

func (t *Texture) Map(u, v float64) Color {
	v = 1 - v
	x := int(u * float64(t.Width))
	y := int(v * float64(t.Height))
	r, g, b, a := t.Image.At(x, y).RGBA()
	f := float64(0xffff)
	return NewColor(float64(r)/f, float64(g)/f, float64(b)/f, float64(a)/f)
}
