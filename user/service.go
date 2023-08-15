package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input InputRegistUser) (*Users, error)
	Login(input InputLoginUser) (Users, error)
}

type service struct {
	Repo UserRepository
}

func NewService(Repo UserRepository) *service {
	return &service{Repo}
}

func (s *service) Register(input InputRegistUser) (*Users, error) {
	// melakukan mapping tehadap struct Users
	user := Users{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errHash != nil {
		return nil, errHash
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	users, err := s.Repo.Save(user)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *service) Login(input InputLoginUser) (Users, error) {
	email := input.Email
	password := input.Password

	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("email failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
