package dto

type MovieRequest struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	MovieType uint   `json:"movie_type"`
	Status    int    `json:"status"`
}
