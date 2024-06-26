package http

import (
	"errors"
	"mime/multipart"
)

const (
	TYPE_JPEG = "image/jpeg"
	TYPE_PNG  = "image/png"
	TYPE_WEBP = "image/webp"
	MAX_SIZE  = 5 << 20 //5 MB
)

func checkInputImage(file *multipart.FileHeader) error {
	if file.Size > MAX_SIZE {
		return errors.New("file size exceeds the maximum allowed size = 5 MB")
	}

	fileType := file.Header.Get("Content-Type")
	switch fileType {
	case TYPE_JPEG, TYPE_PNG, TYPE_WEBP:
	default:
		return errors.New("file extension must be JPEG, PNG or WEBP")
	}

	return nil
}
