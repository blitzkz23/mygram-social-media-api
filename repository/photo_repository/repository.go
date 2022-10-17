package photo_repository

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
)

type PhotoRepository interface {
	PostPhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhotos() ([]*dto.GetPhotoResponse, errs.MessageErr)
	EditPhotoData(photoID uint, photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletePhoto(photoID uint) errs.MessageErr
}
