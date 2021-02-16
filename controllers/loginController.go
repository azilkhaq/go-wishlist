package controllers

import (
	"encoding/json"
	"net/http"
	"wishlist/helper"
	"wishlist/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	
	data := &models.WhistUser{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	token, err := data.SignIn()
	if err != nil {
		resp := helper.Message(http.StatusBadRequest, "email or password incorrect")
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.Message(http.StatusOK, "Successfully")
	resp["data"] = token
	helper.Response(w, http.StatusBadRequest, resp)
}
