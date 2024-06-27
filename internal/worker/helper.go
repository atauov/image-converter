package worker

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

const (
	ORIGINAL_SIZE = 0 //dont resize
	LARGE_SIZE    = 400
	MEDIUM_SIZE   = 200
	SMALL_SIZE    = 100
	EXT           = ".webp"
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

func resizeImg(img image.Image, size uint) image.Image {
	return resize.Resize(size, size, img, resize.Lanczos3)
}

func encodeToWebp(img image.Image) ([]byte, error) {
	var buf bytes.Buffer

	err := webp.Encode(&buf, img, &webp.Options{Lossless: true})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func modifyFilename(filePath string, size uint) string {
	var newName string
	base := filepath.Base(filePath)

	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	switch size {
	case ORIGINAL_SIZE:
		newName = fmt.Sprintf("%s%s", name, EXT)
	default:
		newName = fmt.Sprintf("%s%s%d%s", name, "_", size, EXT)
	}

	return newName
}

func collectBuffs(img image.Image, imgMap map[uint][]byte) error {
	sizes := []uint{ORIGINAL_SIZE, LARGE_SIZE, MEDIUM_SIZE, SMALL_SIZE}
	for _, size := range sizes {
		if size != ORIGINAL_SIZE {
			img = resizeImg(img, size)
		}

		buff, err := encodeToWebp(img)
		if err != nil {
			return err
		}

		imgMap[size] = buff
	}

	return nil
}

func getFilesFromUrl(url string) [4]string {
	files := [4]string{}
	fileName := filepath.Base(url)
	files[0] = fileName

	name := strings.TrimSuffix(filepath.Base(fileName), EXT)

	files[1] = fmt.Sprintf("%s%s%d%s", name, "_", LARGE_SIZE, EXT)
	files[2] = fmt.Sprintf("%s%s%d%s", name, "_", MEDIUM_SIZE, EXT)
	files[3] = fmt.Sprintf("%s%s%d%s", name, "_", SMALL_SIZE, EXT)

	return files
}

func deleteLocalFile(fileName string) error {
	return os.Remove(fileName)
}
