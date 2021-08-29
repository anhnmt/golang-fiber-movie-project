package request

type EpisodeTypeRequest struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
