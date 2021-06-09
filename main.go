package main

import (
	"fmt"
	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/database"
	"log"
	"os"
)

func main() {
	cfg := config.ReadYaml("")
	fmt.Println("Xin chao")

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	// run migrations; update tables
	database.Migrate(db)
}
