package repository

import (
	"github.com/atauov/image-converter/internal/models"
	"github.com/atauov/image-converter/internal/repository/postgres"
	"gorm.io/gorm"
)

type Images interface {
	CreateImage(imageItem models.Image) error
	GetAllImages() ([]models.Image, error)
	UpdateImage(imageID int, imageItem models.Image) error
	GetImagesByUserID(userID int) ([]models.Image, error)
	DeleteImageByURL(URL string) error
	DeleteImagesByUserID(userID int) error
}

type Repository struct {
	Images
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Images: postgres.NewImagesPostgres(db),
	}
}

func CloseRepository(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
