package repository

import (
	"github.com/atauov/image-converter/internal/models"
	"gorm.io/gorm"
)

type Images interface {
	CreateImage(imageItem models.Image) error
	GetAllImages() ([]models.Image, error)
	UpdateImage(imageItem models.Image) error
	GetImagesByUserID(userID int) ([]models.Image, error)
	DeleteImageByURL(URL string) error
	DeleteImagesByUserID(userID int) error
}

type Repository struct {
	Images
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Images: repository.NewPostgresDB(db),
	}
}
