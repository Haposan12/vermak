package controllers

import (
	"encoding/json"
	"net/http"
	"vermak2/constants"
	"vermak2/models"
	u "vermak2/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil{
		u.Respond(w, u.Message(constants.FalseStatus, "Invalid request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil{
		u.Respond(w, u.Message(constants.FalseStatus,"Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

