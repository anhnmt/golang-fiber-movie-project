package controller

import (
	"github.com/xdorro/golang-fiber-base-project/app/repository"
)

var (
	genreController = &GenreController{
		genreRepository: repository.NewTagRepository(),
	}

	tagController = &TagController{
		tagRepository: repository.NewTagRepository(),
	}

	episodeController = &EpisodeController{
		episodeRepository:       repository.NewEpisodeRepository(),
		episodeDetailRepository: repository.NewEpisodeDetailRepository(),
	}
)
