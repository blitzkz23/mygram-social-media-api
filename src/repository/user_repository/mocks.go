package user_repository

import (
	"mygram-social-media-api/src/entity"
	"mygram-social-media-api/src/pkg/errs"
)

// * MockUserRepository is a mock of UserRepository interface
var (
	Login               func(user *entity.User) (*entity.User, errs.MessageErr)
	Register            func(user *entity.User) (*entity.User, errs.MessageErr)
	GetUserByIDAndEmail func(user *entity.User) (*entity.User, errs.MessageErr)
	UpdateUserData      func(userId uint, user *entity.User) (*entity.User, errs.MessageErr)
	DeleteUser          func(userId uint) errs.MessageErr
)

type mockUserRepository struct{}

func NewMockUserRepository() UserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Login(user *entity.User) (*entity.User, errs.MessageErr) {
	return Login(user)
}

func (m *mockUserRepository) Register(user *entity.User) (*entity.User, errs.MessageErr) {
	return Register(user)
}

func (m *mockUserRepository) GetUserByIDAndEmail(user *entity.User) (*entity.User, errs.MessageErr) {
	return nil, nil
}

func (m *mockUserRepository) UpdateUserData(userId uint, user *entity.User) (*entity.User, errs.MessageErr) {
	return UpdateUserData(userId, user)
}

func (m *mockUserRepository) DeleteUser(userId uint) errs.MessageErr {
	return DeleteUser(userId)
}
