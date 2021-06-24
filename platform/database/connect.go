package database

import (
	"fmt"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// connect : creates a new Database object
func connect(dbConfig *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	var newLogger logger.Interface
	if dbConfig.Logger {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
