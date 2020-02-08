package main

import (
	. "github.com/arata-nvm/poly/poly"
)

func main() {
	d := NewDevice(256, 256)
	d.ClearColorBuffer(BLACK)

	d.PutPixel(128, 128, WHITE)

	Save("out.png", d.Image())
}
