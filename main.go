package main

import (
	"image"
	"img/effects"
	"img/imagemanipulation"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

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

func check(err error) {
	if err != nil {
		print("Error Occured: ")
		panic(err)
	}
}

// The work flow of loading an image in, modifying it, and saving the new image.
func start(readName string, writeName string) {
	loadedImage, err := loadImage(readName)
	check(err)

	newImage := imagemanipulation.PixelLoop(loadedImage, effects.Greyscale)

	newFile, err := os.Create(writeName)
	check(err)

	q := jpeg.Options{Quality: 75}
	err = jpeg.Encode(newFile, newImage, &q)
	check(err)
}
