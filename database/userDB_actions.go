package database

import "github.com/RezaZahedi/Go-Gin/model"

type UserDB struct {
	DB

}

// NewUserDB returns a User database with a mapDB implementation
func NewUserDB() *UserDB {
	return &UserDB{DB: &mapDB{data: make(map[interface{}]interface{})}}
}

func (db *UserDB) FindAll() ([]*model.User, error) {
	res, err := db.DB.FindAll()
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
	res, err := db.DB.FindByID(id.BackField)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*model.User), nil
}

func (db *UserDB) Create(id *model.ID, input *model.User) (*model.User, error) {
	res, err := db.DB.Create(id.BackField, input)
	return res.(*model.User), err
}

func (db *UserDB) Update(id *model.ID, input *model.User) (*model.User, error) {
	res, err := db.DB.Update(id.BackField, input)
	return res.(*model.User), err
}

func (db *UserDB) Delete(id *model.ID) error {
	return db.DB.Delete(id.BackField)
}
