package request

type ProfileRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Gender   int    `json:"gender"`
}
