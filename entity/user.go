package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const secret_key = "RAHASIA NEGARA"

type User struct {
	GormModel
	Username string `gorm:"not null;unique;type:varchar(191)" form:"username" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;unique;type:varchar(191)" form:"email" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `gorm:"not null;type:varchar(191)" form:"password" json:"password" valid:"required~Password is required, minstringlength(6)~Password must be at least 6 characters"`
	Age      uint8  `gorm:"not null" form:"age" json:"age" valid:"required~Age is required, age~Age must be between 8 and 100"`
}

// ! TODO: Validator Age not trigerred, also hash pashing not working in hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *User) HashPass() error {
	salt := 8
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return errors.New("interal server error")
	}

	u.Password = string(hash)

	return nil
}

// * Verify encrypted password with bcrypt
func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

// * Generate token with jwt
func (u *User) GenerateToken() string {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secret_key))

	return signedToken
}
