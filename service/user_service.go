package service

import (
	"errors"
	"fmt"
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/helpers"
	"mygram-social-media-api/repository/user_repository"
)

type UserService interface {
	Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, error) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Email: userPayload.Email,
	}

	user, err := u.userRepo.Login(payload)
	if err != nil {
		fmt.Println("Error get use by email")
		return nil, err
	}

	validPassword := user.VerifyPassword(userPayload.Password)
	if !validPassword {
		fmt.Println("Error compare password")
		return nil, errors.New("invalid password")
	}

	token := user.GenerateToken()

	response := &dto.LoginResponse{
		AccessToken: token,
	}

	return response, nil
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Password: userPayload.Password,
		Age:      userPayload.Age,
	}

	err = user.HashPass()
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	res := &dto.RegisterResponse{
		Message: "User data has been successfully registered",
	}

	return res, nil
}
