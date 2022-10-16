package user_pg

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/user_repository"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) GetUserByID(userId uint) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Debug().Model(user).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByIDAndEmail(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Debug().Where("email = ? AND id = ?", userPayload.Email, userPayload.ID).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) Login(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Debug().Where("email = ?", userPayload.Email).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) Register(userPayload *entity.User) errs.MessageErr {
	err := u.db.Debug().Create(userPayload).Error

	if err != nil {
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}

func (u *userPG) UpdateUserData(userId uint, userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Debug().Model(user).Where("id = ?", userId).Updates(userPayload).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) DeleteUser(userId uint) errs.MessageErr {
	user := entity.User{}

	err := u.db.Debug().Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
