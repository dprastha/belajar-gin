package repository

import (
	"belajar-gin/server/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	Register(user *model.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetUsers() (*[]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) Register(user *model.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
