package database

import (
	"errors"
	"sync"
)

var ErrNotExist error = errors.New("entry does not exist")

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
	sync.RWMutex
	data map[interface{}]interface{}
}

func (db *mapDB) FindAll() ([]interface{}, error) {
	db.RLock()
	defer db.RUnlock()
	output := make([]interface{}, 0, len(db.data))
	for _, val := range db.data {
		val := val
		output = append(output, val)
	}
	return output, nil
}

func (db *mapDB) FindByID(id interface{}) (interface{}, error) {
	db.RLock()
	defer db.RUnlock()
	data, ok := db.data[id]
	if !ok {
		return nil, ErrNotExist
	}
	return data, nil
}

func (db *mapDB) Create(id interface{}, input interface{}) (interface{}, error) {
	if _, err := db.FindByID(id); err != ErrNotExist {
		if err != nil {
			return nil, errors.New("id already exists" + err.Error())
		}
		return nil, errors.New("id already exists")
	}
	db.Lock()
	defer db.Unlock()
	db.data[id] = input
	return input, nil
}

func (db *mapDB) Update(id interface{}, input interface{}) (interface{}, error) {
	if _, err := db.FindByID(id); err != nil {
		return nil, err
	}
	db.Lock()
	defer db.Unlock()
	db.data[id] = input
	return input, nil
}

func (db *mapDB) Delete(id interface{}) error {
	if _, err := db.FindByID(id); err != nil {
		return err
	}
	db.Lock()
	defer db.Unlock()
	delete(db.data, id)
	return nil
}

