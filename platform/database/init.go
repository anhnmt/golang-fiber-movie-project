package database

import (
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	err error
	db  *gorm.DB
)

func init() {
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

func GetDB() *gorm.DB {
	return db
}
