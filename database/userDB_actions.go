package database

import "github.com/RezaZahedi/Go-Gin/model"

type UserDB struct {
	back db
}

// NewUserDB returns a User database with a mapDB implementation
func NewUserDB() *UserDB {
	return &UserDB{back: &mapDB{data: make(map[interface{}]interface{})}}
}

func (db *UserDB) FindAll() ([]*model.User, error) {
	res, err := db.back.FindAll()
	if res == nil || err != nil {
		return nil, err
	}
	output := make([]*model.User, 0, len(res))
	for _, r := range res {
		output = append(output, r.(*model.User))
	}
	return output, nil
}

func (db *UserDB) FindByID(id *model.ID) (*model.User, error) {
	res, err := db.back.FindByID(id.BackField)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.User), nil
}

func (db *UserDB) Create(id *model.ID, input *model.User) (*model.User, error) {
	res, err := db.back.Create(id.BackField, input)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.User), nil
}

func (db *UserDB) Update(id *model.ID, input *model.User) (*model.User, error) {
	res, err := db.back.Update(id.BackField, input)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.User), nil
}

func (db *UserDB) Delete(id *model.ID) error {
	return db.back.Delete(id.BackField)
}
