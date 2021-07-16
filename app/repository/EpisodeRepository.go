package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllEpisodesByMovieIdAndStatusNot(id string, status int) (*[]model.Episode, error) {
	movieId := util.ParseStringToUInt(id)
	episodes := make([]model.Episode, 0)

	if err := db.Find(&episodes, "movie_id = ? AND status <> ?", movieId, status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &episodes, nil
}
