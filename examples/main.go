package main

import (
	. "github.com/arata-nvm/poly/poly"
)

func main() {
	d := NewDevice(256, 256)
	d.ClearColorBuffer(BLACK)

	m := LoadObj("examples/box3.obj")

	d.DrawMesh(m, WHITE)

	Save("out.png", d.Image())
}
