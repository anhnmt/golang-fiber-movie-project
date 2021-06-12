package database

import (
	"fmt"
	config2 "github.com/xdorro/golang-fiber-base-project/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase creates a new Database object
func NewDatabase(config *config2.YamlConfig) (*gorm.DB, error) {
	var err error
	dbConfig := config.Database

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
