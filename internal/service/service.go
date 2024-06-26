package service

import (
	"github.com/atauov/image-converter/internal/models"
	"github.com/atauov/image-converter/internal/repository"
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Images: NewImagesService(repos.Images),
	}
}
