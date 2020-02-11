package user

import (
	"errors"
	"github.com/RezaZahedi/Go-Gin/database"
	"github.com/RezaZahedi/Go-Gin/model"
	"github.com/RezaZahedi/Go-Gin/next_user_id"
)

type UserRepository struct {
	UserDB *database.UserDB
	IDGen  func() int
}

func ProvideUserRepository(db *database.UserDB) UserRepository {
	return UserRepository{UserDB: db, IDGen: next_user_id.NewIncrementalIDWithNoMemory()}
}

func (u *UserRepository) FindAll() ([]User, error) {
	users, err := u.UserDB.FindAll()
	_users := make([]User, 0, len(users))
	for _, user := range users {
		_users = append(_users, User(user))
	}
	return _users, err
}

func (u *UserRepository) Create(id uint, user User) (User, error) {
	return u.UserDB.Create(&model.ID{BackField: int(id)}, user)
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
	id := u.IDGen()
	user := model.User{
		ID:       model.ID{BackField: id},
		Username: username,
		Password: password,
	}
	return u.Create(uint(id), &user)
}

func (u *UserRepository) IsUserNameAvailable(username string) bool {
	users, err := u.FindAll()
	if err != nil {
		return false
	}
	for _, user := range users {
		if user.Username == username {
			return false
		}
	}
	return true
}
