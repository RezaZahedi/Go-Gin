package user

import (
	"errors"
	"strings"
)

type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(p UserRepository) UserService {
	return UserService{UserRepository: p}
}

func (u *UserService) IsUserValid(username, password string) (bool, error) {
	return u.UserRepository.IsUserValid(username, password)
}

func (u *UserService) RegisterNewUser(username, password string) (User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	}
	return u.UserRepository.RegisterNewUser(username, password)
}