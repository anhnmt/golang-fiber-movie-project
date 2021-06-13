package database

import (
	"fmt"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	err error
	DB  *gorm.DB
)

func init() {
	dbConfig := config.GetDatabase()

	DB, err = newDatabase(dbConfig)
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
	return DB
}

// newDatabase : creates a new Database object
func newDatabase(dbConfig *config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	fmt.Println(dsn)

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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// migrate : updates the db with new columns, and tables
func migrate() {
	if err := DB.AutoMigrate(
		model.User{},
		model.Role{},
		model.Permission{},
		model.UserRole{},
		model.UserPermission{},
		model.RolePermission{},

		model.Tag{},
		model.Genre{},
		model.Country{},
		model.People{},

		model.Movie{},
		model.MovieTag{},
		model.MovieGenre{},
		model.MovieCountry{},
		model.MoviePeople{},
	); err != nil {
		log.Println(err)
	}
}
