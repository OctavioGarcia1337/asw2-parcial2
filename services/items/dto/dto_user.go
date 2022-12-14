package dto

type UserDto struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UsersDto []UserDto
