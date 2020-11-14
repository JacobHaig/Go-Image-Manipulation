package effects

import (
	"image/color"
	"img/mathutil"
)

// Greyscale is a modification function. It simply chnages the values of the pixal so that it is greyscale
func Greyscale(r, g, b, a uint8) color.Color {

	var ave uint8 = uint8((int16(r) + int16(g) + int16(b)) / 3)

	rr := mathutil.Clamp(int16(ave), 0, 255)
	gg := mathutil.Clamp(int16(ave), 0, 255)
	bb := mathutil.Clamp(int16(ave), 0, 255)
	aa := mathutil.Clamp(int16(a), 0, 255)

	return color.RGBA{uint8(rr), uint8(gg), uint8(bb), uint8(aa)}
}

// Brighten is a modification function. It simply brightens the values of the colors
func Brighten(r, g, b, a uint8) color.Color {
	addAmount := int16(50)

	rr := mathutil.Clamp(int16(r)+addAmount, 0, 255)
	gg := mathutil.Clamp(int16(g)+addAmount, 0, 255)
	bb := mathutil.Clamp(int16(b)+addAmount, 0, 255)
	aa := mathutil.Clamp(int16(a), 0, 255)

	return color.RGBA{uint8(rr), uint8(gg), uint8(bb), uint8(aa)}
}
