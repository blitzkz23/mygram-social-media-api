package user_repository

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
)

type UserRepository interface {
	GetUserByID(userId uint) (*entity.User, errs.MessageErr)
	Login(user *entity.User) (*entity.User, errs.MessageErr)
	Register(user *entity.User) errs.MessageErr
	GetUserByIDAndEmail(user *entity.User) (*entity.User, errs.MessageErr)
	UpdateUserData(userId uint, user *entity.User) (*entity.User, errs.MessageErr)
	DeleteUser(userId uint) errs.MessageErr
}
