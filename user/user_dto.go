package user

type UserDTO struct {
	ID       *uint  `json:"id,string"`
	Username string `json:"username"`
	Password string `json:"password,string"`
}
