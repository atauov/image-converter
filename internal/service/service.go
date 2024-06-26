package service

import (
	"github.com/atauov/image-converter/internal/models"
	"github.com/atauov/image-converter/internal/repository"
	"github.com/minio/minio-go/v7"
)

type Service struct {
	Images
}

type Images interface {
	CreateImage(imageItem models.Image) error
	GetAllImages(limit, offset int) ([]models.Image, error)
	UpdateImage(imageID int, imageItem models.Image) error
	GetImagesByUserID(userID int) ([]models.Image, error)
	DeleteImageByURL(URL string) error
	DeleteImagesByUserID(userID int) error
}

type S3service interface {
	PutObject() error
	DeleteObject(filename string) error
}

func NewService(repos *repository.Repository, client *minio.Client) *Service {
	return &Service{
		Images: NewImagesService(repos.Images),
	}
}
