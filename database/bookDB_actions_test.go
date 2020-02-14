package database

import (
	"github.com/RezaZahedi/Go-Gin/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var inBook1 = &model.Book{
	ID:          model.ID{BackField: 2},
	Name:        "Where the Wild Things Are",
	Description: "Des2",
}
var inBook2 = &model.Book{
ID:          model.ID{BackField: 3},
Name:        "One Flew Over the Cuckoo's Nest",
Description: "des2sec",
}

func testCorrectnessBook(t *testing.T, err error, outBook *model.Book, inBook *model.Book) {
	assert.Nil(t, err)
	assert.Equal(t, inBook, outBook, "input and output books")
}

func TestBookDB_CreateFindAll(t *testing.T) {
	bookDB := NewBookDB()

	outBook1, err := bookDB.Create(&inBook1.ID, inBook1)
	testCorrectnessBook(t, err, outBook1, inBook1)
	outBook2, err := bookDB.Create(&inBook2.ID, inBook2)
	testCorrectnessBook(t, err, outBook2, inBook2)

	books, err := bookDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(books), "length of books array after creating a new book")
	testCorrectnessBook(t, err, books[0], inBook1)
	testCorrectnessBook(t, err, books[1], inBook2)

}

func TestBookDB_FindByID(t *testing.T) {
	bookDB := NewBookDB()

	bookDB.Create(&inBook1.ID, inBook1)
	outBook, err := bookDB.FindByID(&model.ID{BackField: 2})
	testCorrectnessBook(t, err, outBook, inBook1)
}

func TestBookDB_Delete(t *testing.T) {
	bookDb := NewBookDB()

	bookDb.Create(&inBook1.ID, inBook1)
	books, _ := bookDb.FindAll()
	assert.Equal(t, 1, len(books), "length of books array before deleting a book")
	err := bookDb.Delete(&model.ID{BackField: 2})
	assert.Nil(t, err)
	books, _ = bookDb.FindAll()
	assert.Equal(t, 0, len(books), "length of books array after deleting a book")
}

func TestBookDB_Update(t *testing.T) {
	bookDB := NewBookDB()

	bookDB.Create(&inBook1.ID, inBook1)


	inBook2IDTemp := inBook2.BackField

	inBook2.BackField = inBook1.BackField
	outBook, err := bookDB.Update(&model.ID{BackField: 2}, inBook2)
	testCorrectnessBook(t, err, outBook, inBook2)
	outBook, _ = bookDB.FindByID(&model.ID{2})
	assert.Equal(t, inBook2, outBook)

	inBook2.BackField = inBook2IDTemp
}
