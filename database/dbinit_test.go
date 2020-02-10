package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	db, err := Init()
	assert.Nil(t, err)

	books, err := db.bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(books), "Length of books")

	users, err := db.userDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users), "Length of Users")
}

func Test_init(t *testing.T) {
	userDB := NewUserDB()
	bookDB := NewBookDB()
	assert.Nil(t, _init(userDB, bookDB))
	books, err := bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(books), "Length of books")

}

func TestInitializeBookDB(t *testing.T) {
	bookDB := NewBookDB()
	err := initializeBookDB(bookDB)
	assert.Nil(t, err)
	books, err := bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(books), "Length of books")
}
