package database

func Init()(struct{userDB *UserDB; bookDB *BookDB}, error){
	userDB := NewUserDB()
	bookDB := NewBookDB()
	if err := _init(userDB, bookDB); err != nil {
		return struct {
			userDB *UserDB
			bookDB *BookDB
		}{nil, nil} , err
	}
	return struct {
		userDB *UserDB
		bookDB *BookDB
	}{userDB: userDB, bookDB: bookDB}, nil
}

func _init(userdb *UserDB,bookdb *BookDB) error {
	userdb = NewUserDB()
	bookdb = NewBookDB()
	if err := initializeBookDB(bookdb); err != nil {
		return err
	}
	if err := initializeUserDB(userdb); err != nil {
		return err
	}
	return nil
}

func initializeBookDB(bookdb *BookDB) error{
	var err error
	for _, val := range books {
		if _, err = bookdb.Create(&val); err != nil {
			return err
		}
	}
	return nil
}
func initializeUserDB(userdb *UserDB) error {
	var err error
	for _, val := range users {
		if _, err = userdb.Create(&val); err != nil {
			return err
		}
	}
	return nil
}
