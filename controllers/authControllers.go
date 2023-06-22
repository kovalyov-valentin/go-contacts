package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kovalyov-valentin/go-contacts/models"
	u "github.com/kovalyov-valentin/go-contacts/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // декодируем тело запроса в struct и завершаем работу с ошибкой, если она возникнет
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() // Создать аккаунт
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) // декодируем тело запроса в struct и завершаем работу с ошибкой, если она возникнет
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return 
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}