package main

import (
	"image"
	"image/png"
	"os"
)

func Save(filename string, img image.Image) {
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
}
