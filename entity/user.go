package entity

import (
	"mygram-social-media-api/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique;type:varchar(191)" form:"username" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;unique;type:varchar(191)" form:"email" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `gorm:"not null;type:varchar(191)" form:"password" json:"password" valid:"required~Password is required, minstringlength(6)~Password must be at least 6 characters"`
	Age      uint8  `gorm:"not null" form:"age" json:"age" valid:"required~Age is required, age~Age must be between 8 and 100"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	// Hash password
	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}
