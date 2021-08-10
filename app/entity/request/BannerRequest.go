package request

type BannerRequest struct {
	MovieId uint   `json:"movie_id"`
	Image   string `json:"image"`
	Status  int    `json:"status"`
}
