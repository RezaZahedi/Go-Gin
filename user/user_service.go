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

//func (p *UserService) FindAll() ([]User, error) {
//	return p.UserRepository.FindAll()
//}

//func (p *UserService) FindByID(id uint) (User, error) {
//	return p.UserRepository.FindByID(id)
//}
//
//func (p *UserService) Create(id uint, user User) (User, error) {
//	return p.UserRepository.Create(id, user)
//}
//
//func (p *UserService) Update(id uint, user User) (User, error) {
//	return p.UserRepository.Update(id, user)
//}
//
//func (p *UserService) Delete(user User) error {
//	return p.UserRepository.Delete(user)
//}

func (u *UserService) IsUserValid(username, password string) (bool, error) {
	return u.UserRepository.IsUserValid(username, password)
}

func (u *UserService) RegisterNewUser(username, password string) (User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	}
	return u.UserRepository.RegisterNewUser(username, password)
}