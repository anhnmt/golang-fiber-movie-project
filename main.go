package main

import (
	"fmt"
	config2 "github.com/xdorro/golang-fiber-base-project/pkg/config"
	database2 "github.com/xdorro/golang-fiber-base-project/platform/database"
	"log"
	"os"
)

func main() {
	cfg := config2.ReadYaml("")
	fmt.Println("Xin chao")

	db, err := database2.NewDatabase(cfg)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	// run migrations; update tables
	database2.Migrate(db)
}
