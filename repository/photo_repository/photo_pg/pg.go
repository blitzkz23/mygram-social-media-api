package photo_pg

import (
	"fmt"
	"mygram-social-media-api/dto"
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

func (p *photoPG) GetAllPhotos() ([]*dto.GetPhotoResponse, errs.MessageErr) {
	photos := []*dto.GetPhotoResponse{}
	photoWithUser := photo_repository.PhotoWithUser{}

	rows, err := p.db.Debug().Model(photoWithUser).Table("photos").Select("photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.email, users.username").Joins("JOIN users ON photos.user_id = users.id").Rows()
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}
	for rows.Next() {
		err = p.db.ScanRows(rows, &photoWithUser)
		if err != nil {
			return nil, errs.NewInternalServerErrorr("Something went wrong")
		}

		dto := photoWithUser.ToGetPhotoResponseDTO()
		photos = append(photos, &dto)
	}

	return photos, nil
}

func (p *photoPG) GetPhotoByID(photoID uint) (*entity.Photo, errs.MessageErr) {
	photo := entity.Photo{}

	fmt.Println("APAKAH ADA ID DISINI", photoID)
	err := p.db.Debug().Model(photo).Where("id = ?", photoID).First(&photo).Error
	fmt.Println("APAKAH ADA ID DISINI 2", photoID, &photo)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) EditPhotoData(photoID uint, photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	photo := entity.Photo{}
	fmt.Println("APAKAH ADA ID DISINI 3", photoID)

	err := p.db.Debug().Model(photo).Where("id = ?", photoID).Updates(&photoPayload).Take(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}
	fmt.Println("Melihat photo", &photo)

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
