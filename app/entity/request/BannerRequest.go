package request

type BannerRequest struct {
	Image  string `json:"image"`
	Url    uint   `json:"url"`
	Status int    `json:"status"`
}
