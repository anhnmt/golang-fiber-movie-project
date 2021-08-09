package dto

type MovieTypeDTO struct {
	MovieTypeId uint   `json:"movie_type_id,omitempty"`
	Name        string `json:"name,omitempty"`
}
