package controllers

import "wishlist/middleware"

func (server *Server) initializeRoutes() {
	route := server.Router.HandleFunc
	setMiddleJSON := middleware.SetMiddlewareJSON
	
	// Auth
	route("/login", setMiddleJSON(server.Login)).Methods("POST")

	// Users
	route("/users/add", setMiddleJSON(server.CreateUsers)).Methods("POST")
	route("/users/get", setMiddleJSON(server.GetAllUsers)).Methods("GET")
	route("/users/get/{id}", setMiddleJSON(server.GetUsersByID)).Methods("GET")
	route("/users/update/{id}", setMiddleJSON(server.UpdateUsers)).Methods("PUT")
	route("/users/delete/{id}", setMiddleJSON(server.DeleteUsers)).Methods("DELETE")

	//BM
	route("/bm/add", setMiddleJSON(server.CreateBm)).Methods("POST")
	route("/bm/get", setMiddleJSON(server.GetAllBm)).Methods("GET")
	route("/bm/get/{id}", setMiddleJSON(server.GetBmByID)).Methods("GET")
	route("/bm/update/{id}", setMiddleJSON(server.UpdateBm)).Methods("PUT")
}