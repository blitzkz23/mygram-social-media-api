package service

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/photo_repository"
)

type PhotoService interface {
	PostPhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhotos() ([]entity.Photo, errs.MessageErr)
	EditPhotoData(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletePhoto(photoId uint) errs.MessageErr
}

type photoService struct {
	photoRepository photo_repository.PhotoRepository
}

func NewPhotoService(photoRepository photo_repository.PhotoRepository) PhotoService {
	return &photoService{photoRepository: photoRepository}
}

func (p *photoService) PostPhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr) {
	return nil, nil
}

func (p *photoService) GetAllPhotos() ([]entity.Photo, errs.MessageErr) {
	return nil, nil
}

func (p *photoService) EditPhotoData(photo *entity.Photo) (*entity.Photo, errs.MessageErr) {
	return nil, nil
}

func (p *photoService) DeletePhoto(photoId uint) errs.MessageErr {
	return nil
}
