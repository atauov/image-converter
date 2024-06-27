package service

import (
	"fmt"
	"github.com/atauov/image-converter/internal/models"
	"path/filepath"
	"strings"
)

const (
	link           = "http://dwah87zivx7wf.cloudfront.net/" //CDN domain
	EXTENSION      = ".webp"
	POSTFIX_LARGE  = "_400"
	POSTFIX_MEDIUM = "_200"
	POSTFIX_SMALL  = "_100"
)

func fillerLink(image *models.Image) {
	base := filepath.Base(image.Filename)
	ext := filepath.Ext(image.Filename)
	clearFileName := strings.TrimSuffix(base, ext)

	image.OriginalUrl = fmt.Sprintf("%s%s%s", link, clearFileName, EXTENSION)
	image.SizeLarge = fmt.Sprintf("%s%s%s%s", link, clearFileName, POSTFIX_LARGE, EXTENSION)
	image.SizeMedium = fmt.Sprintf("%s%s%s%s", link, clearFileName, POSTFIX_MEDIUM, EXTENSION)
	image.SizeSmall = fmt.Sprintf("%s%s%s%s", link, clearFileName, POSTFIX_SMALL, EXTENSION)
}
