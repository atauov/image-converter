package service

import (
	"github.com/atauov/image-converter/internal/models"
	"github.com/atauov/image-converter/internal/repository"
)

type ImagesService struct {
	repo repository.Images
}

func NewImagesService(repo repository.Images) *ImagesService {
	return &ImagesService{repo: repo}
}

func (s *ImagesService) CreateImage(imageItem models.Image) error {

	return s.repo.CreateImage(imageItem)
}

func (s *ImagesService) GetAllImages() ([]models.Image, error) {
	return s.repo.GetAllImages()
}

func (s *ImagesService) UpdateImage(imageID int, imageItem models.Image) error {
	return s.repo.UpdateImage(imageID, imageItem)
}

func (s *ImagesService) GetImagesByUserID(userID int) ([]models.Image, error) {
	return s.repo.GetImagesByUserID(userID)
}

func (s *ImagesService) DeleteImageByURL(URL string) error {
	return s.repo.DeleteImageByURL(URL)
}

func (s *ImagesService) DeleteImagesByUserID(userID int) error {
	return s.repo.DeleteImagesByUserID(userID)
}
