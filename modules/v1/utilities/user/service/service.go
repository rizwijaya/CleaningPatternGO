package service

import (
	"ClearningPatternGO/modules/v1/utilities/user/models"
	"ClearningPatternGO/modules/v1/utilities/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input models.RegisterUserInput) (models.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input models.RegisterUserInput) (models.User, error) {
	user := models.User{}
	user.Nama = input.Nama
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Id_role = input.Role
	user.Avatar = "Images/avatar.png"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
