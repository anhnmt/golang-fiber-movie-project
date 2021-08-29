package dto

type SearchBannerDTO struct {
	BannerId    int64  `json:"banner_id"`
	MovieId     int64  `json:"movie_id"`
	OriginName  string `json:"origin_name"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ReleaseDate string `json:"release_date"`
	Image       string `json:"image"`
	Status      int    `json:"status"`
}
