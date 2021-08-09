package dto

type SearchBannerDTO struct {
	BannerId    uint   `json:"banner_id"`
	MovieId     uint   `json:"movie_id"`
	OriginName  string `json:"origin_name"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ReleaseDate string `json:"release_date"`
	Image       string `json:"image"`
	Status      int    `json:"status"`
}
