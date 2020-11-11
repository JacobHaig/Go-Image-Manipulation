package main

import (
	"image"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"image/color"
	"image/jpeg"
)

// Start of the program, general setup and reading of directories
func main() {
	infolder := `image\in\`
	outfolder := `image\out\`

	startTime := time.Now()

	files, err := ioutil.ReadDir(infolder)
	check(err)

	var wg sync.WaitGroup
	for index, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".jpg") {

			wg.Add(1)
			go func(index int, name, infolder, outfolder string) {
				startTIME := time.Now()
				start(infolder, outfolder)

				println(index, ":", name, ":", time.Since(startTIME).Milliseconds(), "ms")
				wg.Done()
			}(index, file.Name(), infolder+file.Name(), outfolder+file.Name())
		}
	}
	wg.Wait()

	println("Converted all Images :", time.Since(startTime).Milliseconds(), "ms")
}

func loadImage(imageName string) (image.Image, error) {
	file, err := os.Open(imageName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, err := jpeg.Decode(file)
	return image, nil
}

// Min function to return the smallest of the two values
func min(x, y uint16) uint16 {
	if x > y {
		return y
	}
	return x
}

func check(err error) {
	if err != nil {
		println(err)
	}
}

// This is the modification function. It simply brightens the values of the colors
func brighten(r, g, b, a uint8) color.Color {
	addAmount := uint16(50)

	r = uint8(min(uint16(r)+addAmount, uint16(math.MaxUint8)))
	g = uint8(min(uint16(g)+addAmount, uint16(math.MaxUint8)))
	b = uint8(min(uint16(b)+addAmount, uint16(math.MaxUint8)))
	//a = uint8(min(uint16(a)+addAmount, uint16(math.MaxUint8)))

	return color.RGBA{r, g, b, a}
}

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

// Loop Over every Pixel and apply the function then set the color
func pixelLoop(loadedImage image.Image) image.Image {
	w, h := loadedImage.Bounds().Dx(), loadedImage.Bounds().Dy()
	rect := image.Rect(0, 0, w, h)
	newImage := image.NewRGBA(rect)

	var wg sync.WaitGroup

	for y := 0; y <= h; y++ {
		wg.Add(1)

		go func(y int) {
			for x := 0; x <= w; x++ {

				colors := loadedImage.At(x, y)
				c := colorHelper(colors, brighten)

				newImage.Set(x, y, c)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()

	return newImage
}

// The work flow of loading an image in, modifying it, and saving the new image.
func start(readName string, writeName string) {
	loadedImage, err := loadImage(readName)
	check(err)

	newImage := pixelLoop(loadedImage)

	newFile, err := os.Create(writeName)
	check(err)

	q := jpeg.Options{Quality: 75}
	err = jpeg.Encode(newFile, newImage, &q)
	check(err)
}
