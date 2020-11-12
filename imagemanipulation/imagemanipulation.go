package imagemanipulation

import (
	"image"
	"image/color"
	"sync"
)

// Applies the function to the color and returns a new color
func colorHelper(colors color.Color, fn func(uint8, uint8, uint8, uint8) color.Color) color.Color {
	r, g, b, a := colors.RGBA()

	r = r >> 8
	g = g >> 8
	b = b >> 8
	a = a >> 8

	newColor := fn(uint8(r), uint8(g), uint8(b), uint8(a))

	return newColor
}

// PixelLoop Loops Over every Pixel and apply the function then set the color
func PixelLoop(loadedImage image.Image, fn func(uint8, uint8, uint8, uint8) color.Color) image.Image {
	w, h := loadedImage.Bounds().Dx(), loadedImage.Bounds().Dy()
	rect := image.Rect(0, 0, w, h)
	newImage := image.NewRGBA(rect)

	var wg sync.WaitGroup

	for y := 0; y <= h; y++ {
		wg.Add(1)

		go func(y int) {
			for x := 0; x <= w; x++ {

				colors := loadedImage.At(x, y)
				c := colorHelper(colors, fn)

				newImage.Set(x, y, c)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()

	return newImage
}
