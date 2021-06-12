package database

import (
	"fmt"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = newDatabase()
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	// run migrations; update tables
	Migrate(DB)
}

// newDatabase creates a new Database object
func newDatabase() (*gorm.DB, error) {
	var err error
	dbConfig := config.Config.Database

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return DB
}
