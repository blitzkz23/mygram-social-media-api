package user_repository

import "mygram-social-media-api/entity"

type UserRepository interface {
	GetUserByID(userId uint) (*entity.User, error)
	Login(user *entity.User) (*entity.User, error)
	Register(user *entity.User) error
}
