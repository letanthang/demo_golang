package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width  int
	height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}
func (i Image) At(x, y int) color.Color {
	v := uint8(x) ^ uint8(y)
	return color.RGBA{v, v, 255, 255}
}
func main() {
	m := Image{256, 64}
	pic.ShowImage(m)
}
