package database

func Init(userdb *UserDB, bookdb *BookDB) error {
	if err := initializeBookDB(bookdb); err != nil {
		return err
	}
	if err := initializeUserDB(userdb); err != nil {
		return err
	}
	return nil
}

func initializeBookDB(bookdb *BookDB) error {
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
