package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/pkg/config"
)

var (
	lock = &sync.Mutex{}
	err  error
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		if db == nil {
			dbConfig := config.GetDatabase()

			db, err = connect(dbConfig)
			if err != nil {
				log.Printf("database err %s", err)
				os.Exit(1)
			}

			// run migrations; update tables
			if dbConfig.Migrate {
				migrate()
			}
		}
	}

	return db
}
