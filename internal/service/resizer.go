package service

import (
	"golang.org/x/image/draw"
	"image"
	"os"
)

const (
	LARGE_SIZE  = 800
	MEDIUM_SIZE = 400
	SMALL_SIZE  = 200
)

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func resizeImage(img image.Image, size int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.BiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return dst
}
