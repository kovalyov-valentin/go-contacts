package app

import (
	"net/http"
	u "github.com/kovalyov-valentin/go-contacts/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resourses was not found on our server"))
		next.ServeHTTP(w, r)
	}) 
}