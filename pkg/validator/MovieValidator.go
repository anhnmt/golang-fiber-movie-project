package validator

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func ValidateMovieId(movieId string) (*model.Movie, error) {
	movieRepository := repository.NewMovieRepository()

	movie, err := movieRepository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return nil, util.ResponseBadRequest("MovieId không tồn tại", err)
	}

	return movie, nil
}
