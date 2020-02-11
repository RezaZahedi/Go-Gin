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
	for _, book := range books {
		if _, err = bookdb.Create(&book.ID, &book); err != nil {
			return err
		}
	}
	return nil
}
func initializeUserDB(userdb *UserDB) error {
	var err error
	for _, user := range users {
		if _, err = userdb.Create(&user.ID, &user); err != nil {
			return err
		}
	}
	return nil
}
