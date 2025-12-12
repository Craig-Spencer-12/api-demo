package dto

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
