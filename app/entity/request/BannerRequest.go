package request

type BannerRequest struct {
	MovieId int64  `json:"movie_id"`
	Image   string `json:"image"`
	Status  int    `json:"status"`
}
