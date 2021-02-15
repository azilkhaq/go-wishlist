package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wishlist/controllers"
	"wishlist/middleware"
	// "wishlist/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	// "github.com/mileusna/crontab"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	// ctab := crontab.New()                      // create cron table
	// ctab.MustAddJob("* * * * *", UpdateStatus) // every minute

	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)
	router.HandleFunc("/users/add", controllers.CreateUsers).Methods("POST")

	// // Auth
	// route("/login", setMiddleJSON(server.Login)).Methods("POST")

	// Users
	// route("/users/get", setMiddleJSON(server.GetAllUsers)).Methods("GET")
	// route("/users/get/{id}", setMiddleJSON(server.GetUsersByID)).Methods("GET")
	// route("/users/update/{id}", setMiddleJSON(server.UpdateUsers)).Methods("PUT")
	// route("/users/delete/{id}", setMiddleJSON(server.DeleteUsers)).Methods("DELETE")

	// //BM
	// route("/bm/add", setMiddleJSON(server.CreateBm)).Methods("POST")
	// route("/bm/get", setMiddleJSON(server.GetAllBm)).Methods("GET")
	// route("/bm/get/{id}", setMiddleJSON(server.GetBmByID)).Methods("GET")
	// route("/bm/update/{id}", setMiddleJSON(server.UpdateBm)).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "PUT", "GET", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	handler := c.Handler(router)
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}

// func UpdateStatus() {
// 	var id string
// 	row := server.DB.Model(&models.WhistBm{}).Debug().Where("is_deleted != ?", true).Select("id").Row()
// 	row.Scan(&id)

// 	fmt.Println(id)
// }
