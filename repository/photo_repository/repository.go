package photo_repository

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
)

type PhotoRepository interface {
	PostPhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhotos() ([]entity.Photo, errs.MessageErr)
	EditPhotoData(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletePhoto(photoId uint) errs.MessageErr
}
