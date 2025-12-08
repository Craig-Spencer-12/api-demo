package dto

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
}
