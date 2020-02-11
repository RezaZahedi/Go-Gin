package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	userDB := NewUserDB()
	bookDB := NewBookDB()
	assert.Nil(t, Init(userDB, bookDB))
	books, err := bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(books), "Length of books")
	users, err := userDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users), "Length of Users")
}

func TestInitializeBookDB(t *testing.T) {
	bookDB := NewBookDB()
	err := initializeBookDB(bookDB)
	assert.Nil(t, err)
	books, err := bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(books), "Length of books")
}
