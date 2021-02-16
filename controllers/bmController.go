package controllers

import (
	"encoding/json"
	"net/http"
	"wishlist/helper"
	"wishlist/models"

	"github.com/gorilla/mux"
)

func CreateBm(w http.ResponseWriter, r *http.Request) {
	data := &models.WhistBm{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.SaveBm()
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusCreated, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusCreated, resp)
}

func GetAllBm(w http.ResponseWriter, r *http.Request) {

	result, err := models.SaveAllBm()
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusOK, resp)
}

func GetSingleBm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ID := vars["id"]

	result, err := models.SaveSingleBm(ID)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusOK, resp)
}

func UpdateBm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ID := vars["id"]

	data := &models.WhistBm{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	_, err = data.SaveUpdateBm(ID)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}
