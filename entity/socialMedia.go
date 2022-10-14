package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;type:varchar(191)" form:"name" json:"name" valid:"required~ame is required"`
	SocialMediaURL string `gorm:"not null" form:"social_media_url" json:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         int64
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
