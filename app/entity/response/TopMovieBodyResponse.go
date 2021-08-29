package response

type TopMovieBodyResponse struct {
	Cinemas  []SearchMovieResponse `json:"cinemas"`
	Movies   []SearchMovieResponse `json:"movies"`
	Series   []SearchMovieResponse `json:"series"`
	Cartoons []SearchMovieResponse `json:"cartoons"`
}
