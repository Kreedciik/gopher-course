package service

import (
	"blogpost/models"
	"blogpost/pkg/helper"
	"blogpost/pkg/repository"
	"fmt"
	"time"
)

type User interface {
	CreateUser(models.UserCreateDTO) error
	GetUserByEmail(string) (models.User, error)
	GetUserById(userId string) (models.User, error)
	LoginUser(models.UserSignInDTO) (string, error)
	UpdateUser(models.UpdateUserDTO) error
	FollowUser(userId, followerId string) error
}

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) CreateUser(user models.UserCreateDTO) error {
	founded, err := u.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if founded.Email != "" {
		return fmt.Errorf("user with email %s is already exists", user.Email)
	}
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.repository.InsertUser(user)
}

func (u *UserService) GetUserByEmail(email string) (models.User, error) {
	return u.repository.FindUserByEmail(email)
}

func (u *UserService) GetUserById(userId string) (models.User, error) {
	return u.repository.FindUserById(userId)
}

func (u *UserService) LoginUser(credentials models.UserSignInDTO) (string, error) {
	founded, err := u.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", err
	}
	if founded.Email == "" {
		return "", fmt.Errorf("user with %s not found", credentials.Email)
	}

	if err := helper.VerifyPassword(credentials.Password, founded.Password); err != nil {
		return "", fmt.Errorf("invalid password")
	}
	return helper.GenerateAccessToken(founded, time.Hour)
}

func (u *UserService) UpdateUser(user models.UpdateUserDTO) error {
	return u.repository.UpdateUser(user)
}

func (u *UserService) FollowUser(userId, followerId string) error {
	return u.repository.FollowUser(userId, followerId)
}
