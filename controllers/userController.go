package controllers

import (
	"encoding/json"
	"net/http"
	"wishlist/helper"
	"wishlist/models"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {

	data := &models.WhistUser{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	err = data.Validate("create")
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.SaveUsers()
	if err != nil {
		format := helper.FormatError(err.Error())
		resp := helper.Message(http.StatusBadRequest, format.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusCreated, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusCreated, resp)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	result, err := models.SaveAllUsers()
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusOK, resp)
}

func GetSingleUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := vars["id"]

	result, err := models.SaveSingleUsers(uid)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	resp["data"] = result
	helper.Response(w, http.StatusOK, resp)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	uid := vars["id"]

	data := &models.WhistUser{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	_, err = data.SaveUpdateUsers(uid)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := vars["id"]

	data := &models.WhistUser{}
	_, err := data.SaveDeleteUsers(uid)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}
