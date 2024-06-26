package service

import "github.com/atauov/image-converter/internal/models"

const link = "http://dwah87zivx7wf.cloudfront.net/" //CDN domain

func fillerLink(image *models.Image) {
	image.OriginalUrl = link + image.Filename
}
