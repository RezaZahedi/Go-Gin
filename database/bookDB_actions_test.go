package database

import (
	"github.com/RezaZahedi/Go-Gin/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var inBook = model.Book{
	ID:          model.ID{BackField: 2},
	Name:        "Where the Wild Things Are",
	Description: "Des2",
}

func TestBookDB_CreateFindAll(t *testing.T) {
	bookDB := NewBookDB()
	outBook, err := bookDB.Create(&inBook.ID, &inBook)
	assert.Nil(t, err)
	assert.Equal(t, &inBook, outBook, "input and output books")
	books, err := bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(books), "length of books array after creating a new book")
}

func TestBookDB_FindByID(t *testing.T) {
	bookDB := NewBookDB()
	bookDB.Create(&inBook.ID, &inBook)
	outBook, err := bookDB.FindByID(&model.ID{BackField: 2})
	assert.Nil(t, err)
	assert.Equal(t, &inBook, outBook)
}

func TestBookDB_Delete(t *testing.T) {
	bookDb := NewBookDB()
	bookDb.Create(&inBook.ID, &inBook)
	books, _ := bookDb.FindAll()
	assert.Equal(t, 1, len(books), "length of books array before deleting a book")
	err := bookDb.Delete(&model.ID{BackField: 2})
	assert.Nil(t, err)
	books, _ = bookDb.FindAll()
	assert.Equal(t, 0, len(books), "length of books array after deleting a book")
}

func TestBookDB_Update(t *testing.T) {
	bookDB := NewBookDB()
	bookDB.Create(&inBook.ID, &inBook)
	secbook := model.Book{
		ID:          model.ID{BackField: 2},
		Name:        "One Flew Over the Cuckoo's Nest",
		Description: "des2sec",
	}
	outBook, err := bookDB.Update(&model.ID{BackField: 2}, &secbook)
	assert.Nil(t, err)
	assert.Equal(t, &secbook, outBook)
	outBook, _ = bookDB.FindByID(&model.ID{2})
	assert.Equal(t, &secbook, outBook)
}
