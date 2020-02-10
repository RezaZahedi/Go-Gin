package database

type BookDB struct {
	DB
}

// NewBookDB returns a Book database with a mapDB implementation
func NewBookDB() *BookDB {
	return &BookDB{DB: &mapDB{data: make(map[interface{}]interface{})}}
}

func (db *BookDB) FindAll() ([]*Book, error)  {
	res, err := db.DB.FindAll()
	if res == nil || err != nil {
		return nil, err
	}
	output := make([]*Book, 0, len(res))
	for _, r := range res {
		output = append(output, r.(*Book))
	}
	return output, nil
}

func (db *BookDB) FindByID(id *ID) (*Book, error) {
	res, err := db.DB.FindByID(id.BackField)
	if res == nil || err != nil {
		return nil, err
	}
	return res.(*Book), nil
}

func (db *BookDB) Create(input *Book) (*Book, error) {
	res, err := db.DB.Create(input.BackField, input)
	return res.(*Book), err
}

func (db *BookDB) Update(id *ID, input *Book) (*Book, error) {
	res, err := db.DB.Update(id.BackField, input)
	return res.(*Book), err
}

func (db *BookDB) Delete(id *ID) error {
	return db.DB.Delete(id.BackField)
}

