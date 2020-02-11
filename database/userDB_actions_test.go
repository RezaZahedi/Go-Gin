package database

import (
	"github.com/RezaZahedi/Go-Gin/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var inUser = model.User{
	ID:       model.ID{BackField: 2},
	Username: "reza",
	Password: "rezareza",
}

func TestUserDB_CreateFindAll(t *testing.T) {
	userDB := NewUserDB()
	outUser, err := userDB.Create(&inUser.ID, &inUser)
	assert.Nil(t, err)
	assert.Equal(t, &inUser, outUser, "input and output users")
	users, err := userDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users), "length of users array after creating a new user")
}

func TestUserDB_FindByID(t *testing.T) {
	userDB := NewUserDB()
	userDB.Create(&inUser.ID, &inUser)
	outUser, err := userDB.FindByID(&model.ID{BackField: 2})
	assert.Nil(t, err)
	assert.Equal(t, &inUser, outUser)
}

func TestUserDB_Delete(t *testing.T) {
	userDb := NewUserDB()
	userDb.Create(&inUser.ID, &inUser)
	users, _ := userDb.FindAll()
	assert.Equal(t, 1, len(users), "length of users array before deleting a user")
	err := userDb.Delete(&model.ID{BackField: 2})
	assert.Nil(t, err)
	users, _ = userDb.FindAll()
	assert.Equal(t, 0, len(users), "length of users array after deleting a user")
}

func TestUserDB_Update(t *testing.T) {
	userDB := NewUserDB()
	userDB.Create(&inUser.ID, &inUser)
	secuser := model.User{
		ID:       model.ID{BackField: 2},
		Username: "Rez",
		Password: "RezRez",
	}
	outUser, err := userDB.Update(&model.ID{BackField: 2}, &secuser)
	assert.Nil(t, err)
	assert.Equal(t, &secuser, outUser)
	outUser, _ = userDB.FindByID(&model.ID{2})
	assert.Equal(t, &secuser, outUser)
}
