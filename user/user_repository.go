package user

import (
	"errors"
	"github.com/RezaZahedi/Go-Gin/database"
	"github.com/RezaZahedi/Go-Gin/model"
)

type UserRepository struct {
	UserDB *database.UserDB
	
}

func ProvideUserRepository(db *database.UserDB) UserRepository {
	return UserRepository{UserDB: db}
}

func (u *UserRepository) FindAll() ([]User, error) {
	users, err := u.UserDB.FindAll()
	_users := make([]User, 0, len(users))
	for _, user := range users {
		_users = append(_users, User(user))
	}
	return _users, err
}

func (u *UserRepository) FindByID(id uint) (User, error) {
	return u.UserDB.FindByID(&model.ID{BackField: int(id)})
}

func (u *UserRepository) Create(id uint, user User) (User, error) {
	return u.UserDB.Create(&model.ID{BackField: int(id)}, user)
}

func (u *UserRepository) Update(id uint, user User) (User, error) {
	return u.UserDB.Update(&model.ID{BackField: int(id)}, user)
}

func (u *UserRepository) Delete(user User) error {
	return u.UserDB.Delete(&user.ID)
}

func (u *UserRepository) IsUserValid(username string, password string) (bool, error) {
	users, err := u.FindAll()
	if err != nil {
		return false, err
	}
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true, nil
		}
	}
	return false, nil
}

func (u *UserRepository) RegisterNewUser(username string, password string) (User, error) {
	if !u.IsUserNameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}
	user := model.User{
		ID:       model.ID{},
		Username: username,
		Password: password,
	}
}
