package postgres

import (
	"github.com/atauov/image-converter/internal/models"
	"gorm.io/gorm"
	"time"
)

type ImagesPostgres struct {
	db *gorm.DB
}

func NewImagesPostgres(db *gorm.DB) *ImagesPostgres {
	return &ImagesPostgres{
		db: db,
	}
}

func (r *ImagesPostgres) CreateImage(imageItem models.Image) error {
	if err := r.db.Create(&imageItem).Error; err != nil {
		return err
	}

	return nil
}

func (r *ImagesPostgres) GetAllImages() ([]models.Image, error) {
	var images []models.Image
	if err := r.db.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (r *ImagesPostgres) UpdateImage(imageID int, imageItem models.Image) error {
	updateData := map[string]interface{}{
		"filename": imageItem.Filename,
		"is_done":  imageItem.IsDone,
	}
	return r.db.Model(&models.Image{}).Where("id = ?", imageID).Updates(updateData).Error
}

func (r *ImagesPostgres) GetImagesByUserID(userID int) ([]models.Image, error) {
	var images []models.Image

	err := r.db.Where("user_id = ?", userID).Find(&images).Error

	return images, err
}

func (r *ImagesPostgres) DeleteImageByURL(URL string) error {
	return r.db.Model(&models.Image{}).Where("original_url = ?", URL).Update("deleted_at", time.Now()).Error
}

func (r *ImagesPostgres) DeleteImagesByUserID(userID int) error {
	return r.db.Model(&models.Image{}).Where("user_id = ?", userID).Update("deleted_at", time.Now()).Error
}
