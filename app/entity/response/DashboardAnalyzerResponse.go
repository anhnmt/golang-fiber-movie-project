package response

import "github.com/xdorro/golang-fiber-movie-project/app/entity/dto"

type DashboardAnalyzerResponse struct {
	Movies       int64                 `json:"movies,omitempty"`
	Genres       int64                 `json:"genres,omitempty"`
	Countries    int64                 `json:"countries,omitempty"`
	Banners      int64                 `json:"banners,omitempty"`
	LatestMovies *[]dto.SearchMovieDTO `json:"latest_movies,omitempty"`
}
