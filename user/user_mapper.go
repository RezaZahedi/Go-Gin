package user

import (
	"errors"
	"github.com/RezaZahedi/Go-Gin/model"
)

func ToUser(userDTO UserDTO) (User, error) {
	if userDTO.ID == nil {
		return nil, errors.New("ID is required")
	}
	return User(&model.User{
		ID:       model.ID{BackField: int(*userDTO.ID)},
		Username: userDTO.Username,
		Password: userDTO.Password,
	}), nil
}

func ToUserDTO(user User) UserDTO {
	id := uint(user.ID.BackField)
	return UserDTO{
		ID:       &id,
		Username: user.Username,
		Password: user.Password,
	}
}

func ToUserDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}
