package database

import "errors"

var ErrNotExist error = errors.New("entry does not exist")
var ErrBadInput error = errors.New("bad entry data")

// DB defines database actions, to ease the change of implementation, e.g.: to SQL.
type DB interface {
	FindAll() ([]interface{}, error)
	FindByID(key interface{}) (interface{}, error)
	Create(id interface{}, input interface{}) (interface{}, error)
	Update(id interface{}, input interface{}) (interface{}, error)
	Delete(id interface{}) error
}

// mapDB implements DB
// it uses a map as the database implementation
type mapDB struct {
	data map[interface{}] interface{}
}


func (db *mapDB) FindAll() ([]interface{}, error) {
	output := make([]interface{}, 0, len(db.data))
	for _, val := range db.data {
		output = append(output, val)
	}
	return output, nil
}

func (db *mapDB) FindByID(id interface{}) (interface{}, error) {
	data, ok := db.data[id]
	if !ok {
		return nil, ErrNotExist
	}
	return data, nil
}

func (db *mapDB) Create(id interface{}, input interface{}) (interface{}, error) {
	db.data[id] = input
	return input, nil
}

func (db *mapDB) Update(id interface{}, input interface{}) (interface{}, error) {
	if _, ok := db.data[id]; !ok {
		return nil, ErrNotExist
	}
	db.data[id] = input
	return input, nil
}

func (db *mapDB) Delete(id interface{}) error {
	if _, ok := db.data[id]; !ok {
		return ErrNotExist
	}
	delete(db.data, id)
	return nil
}