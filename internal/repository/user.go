package repository

import (
	"ms-golang-fiber/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.RegisterUserRequest) error
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.RegisterUserRequest) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
