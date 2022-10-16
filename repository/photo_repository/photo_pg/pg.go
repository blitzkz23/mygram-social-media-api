package photo_pg

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/photo_repository"

	"gorm.io/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPG{db: db}
}

func (p *photoPG) PostPhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	photo := entity.Photo{}

	err := p.db.Debug().Model(photo).Create(&photoPayload).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	err = p.db.Last(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) GetAllPhotos() ([]entity.Photo, errs.MessageErr) {
	photo := entity.Photo{}
	photos := []entity.Photo{}

	err := p.db.Debug().Model(photo).Find(&photos).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return photos, nil
}

func (p *photoPG) EditPhotoData(photoID uint, photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	photo := entity.Photo{}

	err := p.db.Debug().Model(photo).Where("id = ?", photoID).Updates(&photoPayload).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) DeletePhoto(photoID uint) errs.MessageErr {
	photo := entity.Photo{}

	err := p.db.Debug().Model(photo).Where("id = ?", photoID).Delete(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewNotFoundError("Photo not found")
		}
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
