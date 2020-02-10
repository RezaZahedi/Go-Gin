package database

type UserDB struct {
	DB
}

// NewUserDB returns a User database with a mapDB implementation
func NewUserDB() *UserDB {
	return &UserDB{DB: &mapDB{data: make(map[interface{}]interface{})}}
}

func (db *UserDB) FindAll() ([]*User, error)  {
	res, err := db.DB.FindAll()
	if res == nil || err != nil {
		return nil, err
	}
	output := make([]*User, 0, len(res))
	for _, r := range res {
		output = append(output, r.(*User))
	}
	return output, nil
}

func (db *UserDB) FindByID(id *ID) (*User, error) {
	res, err := db.DB.FindByID(id.BackField)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*User), nil
}

func (db *UserDB) Create(input *User) (*User, error) {
	res, err := db.DB.Create(input.BackField, input)
	return res.(*User), err
}

func (db *UserDB) Update(id *ID, input *User) (*User, error) {
	res, err := db.DB.Update(id.BackField, input)
	return res.(*User), err
}

func (db *UserDB) Delete(id *ID) error {
	return db.DB.Delete(id.BackField)
}

