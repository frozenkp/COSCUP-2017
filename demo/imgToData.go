package main

import (
	"encoding/csv"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

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

	b := img.Bounds()
	pixel := make([]string, 0)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			oldColor := img.At(x, y)
			cr, _, _, _ := oldColor.RGBA()
			if cr != 65535 {
				pixel = append(pixel, "1")
			} else {
				pixel = append(pixel, "0")
			}
		}
	}

	path := strings.Split(os.Args[1], "/")
	file := path[len(path)-1]

	pixel = append(pixel, string(file[0]))

	writer := csv.NewWriter(os.Stdout)
	writer.Write(pixel)
	writer.Flush()
}
