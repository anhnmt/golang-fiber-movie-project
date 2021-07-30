package request

type EpisodeRequest struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
