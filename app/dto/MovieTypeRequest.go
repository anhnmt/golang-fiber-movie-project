package dto

type MovieTypeRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status int    `json:"status"`
}
