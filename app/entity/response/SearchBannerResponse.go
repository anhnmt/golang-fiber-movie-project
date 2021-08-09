package response

type SearchBannerResponse struct {
	BannerId uint                `json:"banner_id"`
	Image    string              `json:"image"`
	Status   int                 `json:"status"`
	Movie    SearchMovieResponse `json:"movie"`
}
