package service

import (
	"fmt"
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/pkg/helpers"

	"mygram-social-media-api/repository/user_repository"
)

type UserService interface {
	Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	UpdateUserData(userId uint, userPayload *dto.UpdateUserDataRequest) (*dto.UpdateUserDataResponse, errs.MessageErr)
	DeleteUser(userId uint) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Login(userPayload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Email: userPayload.Email,
	}

	user, err := u.userRepo.Login(payload)
	if err != nil {
		return nil, err
	}

	validPassword := user.VerifyPassword(userPayload.Password)
	if !validPassword {
		return nil, errs.NewNotAuthenticated("Invalid email or password")
	}

	token := user.GenerateToken()

	response := &dto.LoginResponse{
		AccessToken: token,
	}

	return response, nil
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
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

	response := &dto.RegisterResponse{
		Message: "User data has been successfully registered",
	}

	return response, nil
}

func (u *userService) UpdateUserData(userId uint, userPayload *dto.UpdateUserDataRequest) (*dto.UpdateUserDataResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}

	user, err = u.userRepo.UpdateUserData(userId, user)
	if err != nil {
		return nil, err
	}
	fmt.Println("Data user", user.Email)

	response := &dto.UpdateUserDataResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(userId uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	err := u.userRepo.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "User data has been successfully deleted",
	}

	return response, nil
}
