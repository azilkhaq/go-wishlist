package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"wishlist/controllers"
	"wishlist/middleware"
	"wishlist/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mileusna/crontab"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)

	// Auth
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	// Users
	router.HandleFunc("/users/get", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/get/{id}", controllers.GetSingleUsers).Methods("GET")
	router.HandleFunc("/users/update/{id}", controllers.UpdateUsers).Methods("PUT")
	router.HandleFunc("/users/delete/{id}", controllers.DeleteUsers).Methods("DELETE")

	//BM
	router.HandleFunc("/bm/add", controllers.CreateBm).Methods("POST")
	router.HandleFunc("/bm/get", controllers.GetAllBm).Methods("GET")
	router.HandleFunc("/bm/get/{id}", controllers.GetSingleBm).Methods("GET")
	router.HandleFunc("/bm/update/{id}", controllers.UpdateBm).Methods("PUT")

	ctab := crontab.New()                        // create cron table
	ctab.MustAddJob("* * * * *", UpdateStatusBm) // every minute

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "PUT", "GET", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	port := os.Getenv("PORT")
	handler := c.Handler(router)
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

func UpdateStatusBm() {
	var dueDate string
	var id string
	t := time.Now()

	row, err := models.GetDB().Model(&models.WhistBm{}).Where("is_deleted != ?", true).Select("id, due_date").Rows()
	if err != nil {
		log.Println(err)
	}

	for row.Next() {
		row.Scan(&id, &dueDate)

		if dueDate < t.Format(time.RFC3339) {
			models.GetDB().Debug().Exec(`UPDATE whist_bms SET status = ? WHERE id = ?`, "missed", id)
		}
	}
}
