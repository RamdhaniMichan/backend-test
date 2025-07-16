package repository

import (
	"test-naga-exchange/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	UpdateToken(user *model.User) error
	FindByToken(token string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) UpdateToken(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) FindByToken(token string) (*model.User, error) {
	var user model.User
	err := r.db.Where("token = ?", token).First(&user).Error
	return &user, err
}
