package request

type EpisodeDetailRequest struct {
	EpisodeDetailId uint   `json:"episode_detail_id"`
	Name            string `json:"name"`
	Link            string `json:"link"`
	EpisodeTypeId   uint   `json:"episode_type_id"`
	Status          int    `json:"status"`
}
