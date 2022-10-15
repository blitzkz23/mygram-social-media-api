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
	user := entity.User{}

	err := u.db.Debug().Model(user).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userPG) Login(userPayload *entity.User) (*entity.User, error) {
	err := u.db.Debug().Where("email = ?", userPayload.Email).Take(&userPayload).Error
	if err != nil {
		return nil, err
	}

	return userPayload, nil
}

func (u *userPG) Register(userPayload *entity.User) error {
	err := u.db.Debug().Create(userPayload).Error

	if err != nil {
		return err
	}

	return nil
}
