package main

import (
	"encoding/csv"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

const boundX int = 20
const boundY int = 20

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No input file")
	} else if len(os.Args) > 2 {
		log.Fatal("Too many arguments")
	}

	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	img, err := jpeg.Decode(src)
	if err != nil {
		log.Fatal(err)
	}

	minX, maxX, minY, maxY := 19, 0, 19, 0

	// decode original image to 0 and 1
	b := img.Bounds()
	pixel := make([]string, 0)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			oldColor := img.At(x, y)
			cr, _, _, _ := oldColor.RGBA()
			if cr == 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
				pixel = append(pixel, "1")
			} else {
				pixel = append(pixel, "0")
			}
		}
	}

	// split original image
	splitImage := make([]string, 0)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			splitImage = append(splitImage, pixel[y*boundX+x])
		}
	}

	// zoom original image
	newBoundX := maxX - minX + 1
	newBoundY := maxY - minY + 1
	ratioX := float64(boundX) / float64(newBoundX)
	ratioY := float64(boundY) / float64(newBoundY)

	newImage := make([]string, 0)
	for y := 0; y < boundY; y++ {
		for x := 0; x < boundX; x++ {
			nX := int(float64(x) / ratioX)
			nY := int(float64(y) / ratioY)
			newImage = append(newImage, splitImage[nY*newBoundX+nX])
		}
	}

	// output Image
	//outputImage("out.jpg", newImage)

	// add target to csv
	path := strings.Split(os.Args[1], "/")
	file := path[len(path)-1]
	newImage = append(newImage, string(file[0]))
	// output csv
	writer := csv.NewWriter(os.Stdout)
	writer.Write(newImage)
	writer.Flush()

}

func outputImage(name string, newImage []string) {
	out, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	rgba := image.NewRGBA(image.Rect(0, 0, boundX, boundY))

	for y := 0; y < boundY; y++ {
		for x := 0; x < boundX; x++ {
			if newImage[y*boundX+x] == "1" {
				rgba.Set(x, y, color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)})
			} else {
				rgba.Set(x, y, color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)})
			}
		}
	}

	jpeg.Encode(out, rgba, nil)
}
