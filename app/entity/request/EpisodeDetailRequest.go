package request

type EpisodeDetailRequest struct {
	Name          string `json:"name"`
	Link          string `json:"link"`
	EpisodeTypeId int64  `json:"episode_type_id"`
	Status        int    `json:"status"`
}
