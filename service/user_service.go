package service

import (
	"fmt"
	"mygram-social-media-api/repository/user_repository"
)

type UserService interface {
	Login()
	Register()
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Login() {
	fmt.Println("Login")
}

func (u *userService) Register() {
	fmt.Println("Register")
}
