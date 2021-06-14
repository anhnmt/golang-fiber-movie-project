package dto

type SearchUserResponse struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Gender   int8   `json:"gender"`
}
