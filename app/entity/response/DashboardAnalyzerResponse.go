package response

type DashboardAnalyzerResponse struct {
	Movies       int64                 `json:"movies,omitempty"`
	Genres       int64                 `json:"genres,omitempty"`
	Countries    int64                 `json:"countries,omitempty"`
	Banners      int64                 `json:"banners,omitempty"`
	LatestMovies []SearchMovieResponse `json:"latest_movies,omitempty"`
}
