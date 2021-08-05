package request

type BannerRequest struct {
	Image  string `json:"image"`
	Url    string `json:"url"`
	Status int    `json:"status"`
}
