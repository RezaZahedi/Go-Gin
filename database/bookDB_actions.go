package database

import "github.com/RezaZahedi/Go-Gin/model"

type BookDB struct {
	DB
}

// NewBookDB returns a Book database with a mapDB implementation
func NewBookDB() *BookDB {
	return &BookDB{DB: &mapDB{data: make(map[interface{}]interface{})}}
}

func (db *BookDB) FindAll() ([]*model.Book, error) {
	res, err := db.DB.FindAll()
	if res == nil || err != nil {
		return nil, err
	}
	output := make([]*model.Book, 0, len(res))
	for _, r := range res {
		output = append(output, r.(*model.Book))
	}
	return output, nil
}

func (db *BookDB) FindByID(id *model.ID) (*model.Book, error) {
	res, err := db.DB.FindByID(id.BackField)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.Book), nil
}

func (db *BookDB) Create(input *model.Book) (*model.Book, error) {
	res, err := db.DB.Create(input.BackField, input)
	return res.(*model.Book), err
}

func (db *BookDB) Update(id *model.ID, input *model.Book) (*model.Book, error) {
	res, err := db.DB.Update(id.BackField, input)
	return res.(*model.Book), err
}

func (db *BookDB) Delete(id *model.ID) error {
	return db.DB.Delete(id.BackField)
}
