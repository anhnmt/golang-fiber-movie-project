package dto

type MovieTypeDTO struct {
	MovieTypeId int64  `json:"movie_type_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
}
