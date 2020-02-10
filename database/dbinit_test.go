package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	db, err := Init()
	assert.Nil(t, err)

	books, err := db.bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(books))
}
