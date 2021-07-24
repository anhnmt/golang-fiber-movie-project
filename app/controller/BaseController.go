package controller

import (
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"sync"
)

var (
	once sync.Once

	genreInstance *GenreController
	tagInstance   *TagController
)

func NewGenreController() *GenreController {
	once.Do(func() {
		genreInstance = &GenreController{
			genreRepository: repository.NewTagRepository(),
		}
	})

	return genreInstance
}

func NewTagController() *TagController {
	once.Do(func() {
		tagInstance = &TagController{
			tagRepository: repository.NewTagRepository(),
		}
	})

	return tagInstance
}
