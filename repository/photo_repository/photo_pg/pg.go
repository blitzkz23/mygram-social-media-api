package photo_pg

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/photo_repository"

	"github.com/jinzhu/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPG{db: db}
}

func (p *photoPG) PostPhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr) {
	return nil, nil
}

func (p *photoPG) GetAllPhotos() ([]entity.Photo, errs.MessageErr) {
	return nil, nil
}

func (p *photoPG) EditPhotoData(photo *entity.Photo) (*entity.Photo, errs.MessageErr) {
	return nil, nil
}

func (p *photoPG) DeletePhoto(photoId uint) errs.MessageErr {
	return nil
}
