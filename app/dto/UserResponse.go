package dto

type UserResponse struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Gender   int8   `json:"gender"`
}
