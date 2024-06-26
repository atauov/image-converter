package service

import (
	"github.com/atauov/image-converter/internal/models"
	"path/filepath"
	"strings"
)

const (
	link           = "http://dwah87zivx7wf.cloudfront.net/" //CDN domain
	EXTENSION      = ".webp"
	POSTFIX_LARGE  = "_800x800"
	POSTFIX_MEDIUM = "_400x400"
	POSTFIX_SMALL  = "_200x200"
)

func fillerLink(image *models.Image) {
	ext := filepath.Ext(image.Filename)
	clearFileName := strings.TrimSuffix(image.Filename, ext)

	image.OriginalUrl = link + clearFileName + EXTENSION
	image.SizeLarge = link + clearFileName + POSTFIX_LARGE + EXTENSION
	image.SizeMedium = link + clearFileName + POSTFIX_MEDIUM + EXTENSION
	image.SizeSmall = link + clearFileName + POSTFIX_SMALL + EXTENSION
}
