package controllers

import (
	"encoding/json"
	"net/http"
	"wishlist/helper"
	"wishlist/middleware"
	"wishlist/models"

	"github.com/gorilla/mux"
)

func (server *Server) CreateBm(w http.ResponseWriter, r *http.Request) {
	data := &models.WhistBm{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.SaveBm(server.DB)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusCreated, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusCreated, resp)
}

func (server *Server) GetAllBm(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.ExtractTokenMetadata(r)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	data := models.WhistBm{}
	result, err := data.FindAllBm(server.DB, token)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusOK, resp)
}

func (server *Server) GetBmByID(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.ExtractTokenMetadata(r)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := models.WhistBm{}
	result, err := data.FindBmByID(server.DB, ID, token)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusOK, resp)
}

func (server *Server) UpdateBm(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.ExtractTokenMetadata(r)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := &models.WhistBm{}
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	_, err = data.SaveUpdateBm(server.DB, ID, token)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}