package postgres

import (
	"fmt"
	"github.com/atauov/image-converter/internal/config"
	"github.com/atauov/image-converter/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type StoragePostgres struct {
	db *gorm.DB
}

func NewPostgresDB(cfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func (s *StoragePostgres) CreateImage(imageItem models.Image) error {
	if err := s.db.Create(&imageItem).Error; err != nil {
		return err
	}

	return nil
}

func (s *StoragePostgres) GetAllImages() ([]models.Image, error) {
	var images []models.Image
	if err := s.db.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (s *StoragePostgres) UpdateImage(imageID int, imageItem models.Image) error {
	updateData := map[string]interface{}{
		"filename": imageItem.Filename,
		"is_done":  imageItem.IsDone,
	}
	return s.db.Model(&models.Image{}).Where("id = ?", imageID).Updates(updateData).Error
}

func (s *StoragePostgres) GetImagesByUserID(userID int) ([]models.Image, error) {
	var images []models.Image

	err := s.db.Where("user_id = ?", userID).Find(&images).Error

	return images, err
}

func (s *StoragePostgres) DeleteImageByURL(URL string) error {
	return s.db.Model(&models.Image{}).Where("original_url = ?", URL).Update("deleted_at", time.Now()).Error
}

func (s *StoragePostgres) DeleteImagesByUserID(userID int) error {
	return s.db.Model(&models.Image{}).Where("user_id = ?", userID).Update("deleted_at", time.Now()).Error
}
