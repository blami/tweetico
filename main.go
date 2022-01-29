package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	"image/draw"
	"image/png"

	"github.com/nfnt/resize"
	"github.com/ur65/go-ico"
)

func loadIco(path string, size int) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	imgs, err := ico.Decode(f)
	if err != nil {
		return nil, err
	}

	for _, img := range imgs {
		if img.Bounds().Max.X == size {
			return img, nil
		}
	}

	return nil, fmt.Errorf("file does not contain icon of size %d", size)
}

func main() {
	l := log.New(os.Stderr, "", 0)

	// Process command line arguments
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	icoPath := flag.Arg(0)

	// Register decoder with image library

	// Load icon
	ico, err := loadIco(icoPath, 32)
	if err != nil {
		l.Fatalln(err)
	}

	// Make new image
	out := image.NewRGBA(image.Rect(0, 0, 360, 200))

	// 32x32
	r := image.Rect(47, 37, 47+32, 37+32)
	draw.Draw(out, r, ico, image.Point{0, 0}, draw.Src)

	// 64x64
	ico2x := resize.Resize(64, 0, ico, resize.NearestNeighbor)
	r = image.Rect(31, 102, 31+64, 102+64)
	draw.Draw(out, r, ico2x, image.Point{0, 0}, draw.Src)

	// 128x128
	ico4x := resize.Resize(128, 0, ico, resize.NearestNeighbor)
	r = image.Rect(200, 39, 200+128, 39+128)
	draw.Draw(out, r, ico4x, image.Point{0, 0}, draw.Src)

	// Save image
	f, err := os.Create("out.png")
	if err != nil {
		l.Fatalln(err)
	}
	defer f.Close()

	err = png.Encode(f, out)
	if err != nil {
		l.Fatalln(err)
	}

}
