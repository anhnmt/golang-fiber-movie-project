package request

type PermissionRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
