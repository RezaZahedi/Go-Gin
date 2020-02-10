package model

type ID struct {
	BackField int `json:"id"`
}

type Book struct {
	ID
	Name string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID
	Username string `json:"username"`
	Password string `json:"password"`
}
