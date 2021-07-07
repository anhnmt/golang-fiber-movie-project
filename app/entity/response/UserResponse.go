package response

type UserResponse struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Gender   int    `json:"gender"`
}
