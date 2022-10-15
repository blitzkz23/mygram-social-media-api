package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null;type:varchar(191)" form:"title" json:"title" valid:"required~Title is required"`
	Caption  string `form:"caption" json:"caption" valid:"required~Caption is required"`
	PhotoURL string `gorm:"not null;type:varchar(191)" form:"photo_url" json:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
