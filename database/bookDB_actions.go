package database

import "github.com/RezaZahedi/Go-Gin/model"

type BookDB struct {
	back db
}

// NewBookDB returns a Book database with a mapDB implementation
func NewBookDB() *BookDB {
	return &BookDB{back: &mapDB{data: make(map[interface{}]interface{})}}
}

func (db *BookDB) FindAll() ([]*model.Book, error) {
	res, err := db.back.FindAll()
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
	res, err := db.back.FindByID(id.BackField)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.Book), nil
}

func (db *BookDB) Create(id *model.ID, input *model.Book) (*model.Book, error) {
	res, err := db.back.Create(id.BackField, input)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.Book), nil
}

func (db *BookDB) Update(id *model.ID, input *model.Book) (*model.Book, error) {
	res, err := db.back.Update(id.BackField, input)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.Book), nil
}

func (db *BookDB) Delete(id *model.ID) error {
	return db.back.Delete(id.BackField)
}
