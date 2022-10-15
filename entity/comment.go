package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null" form:"message" json:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
