package request

type MovieRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	MovieTypeId uint   `json:"movie_type_id"`
	Status      int    `json:"status"`
}
