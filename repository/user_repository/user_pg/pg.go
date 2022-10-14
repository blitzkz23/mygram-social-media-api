package user_pg

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/repository/user_repository"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) GetUserByID(userId uint) (*entity.User, error) {
	return nil, nil
}

func (u *userPG) Login(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *userPG) Register(user *entity.User) error {
	return nil
}
