package main

import (
	"log"
	"os"
	"fmt"
	"wishlist/controllers"
	"wishlist/models"

	"github.com/joho/godotenv"
	"github.com/mileusna/crontab"
)

var server = controllers.Server{}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	ctab := crontab.New() // create cron table
    ctab.MustAddJob("* * * * *", UpdateStatus) // every minute

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(os.Getenv("SERVER_PORT"))
}

func UpdateStatus() {
	var id string
	row := server.DB.Model(&models.WhistBm{}).Debug().Where("is_deleted != ?", true).Select("id").Row()
	row.Scan(&id)

	fmt.Println(id)
}