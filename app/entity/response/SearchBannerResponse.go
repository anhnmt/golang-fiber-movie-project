package response

type SearchBannerResponse struct {
	BannerId int64               `json:"banner_id"`
	Image    string              `json:"image"`
	Status   int                 `json:"status"`
	Movie    SearchMovieResponse `json:"movie"`
}
