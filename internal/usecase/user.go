package usecase

import (
	"ms-golang-fiber/internal/model"
	"ms-golang-fiber/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(user *model.RegisterUserRequest) error
	Login(email, password string) (model.User, error)
	FindAll() ([]model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) Register(user *model.RegisterUserRequest) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	return u.userRepo.Create(user)
}

func (u *userUsecase) Login(email, password string) (model.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}

func (u *userUsecase) FindAll() ([]model.User, error) {
	return u.userRepo.FindAll() // panggil repository
}
