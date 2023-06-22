package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kovalyov-valentin/go-contacts/models"
	u "github.com/kovalyov-valentin/go-contacts/utils"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		noAuth := []string{"/api/user/new", "/api/user/login"} // эндпоинты для которых не требуется авторизация
		requestPath := r.URL.Path // текущий путь запроса

		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // Получение токена

		if tokenHeader == "" { // Если токен отсутствует, возвращаем 403 код
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") // Токен обычно поставляется в формате Bearer {token-body}, мы проверяем соответствует ли полученный токен этому требованию
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return 
		}

		tokenPart := splitted[1] // Часть токена, которая нам интересна
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil { // Неправильный токен возвращает 403 код
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return 
		}

		if !token.Valid { // токен недействителем, возможно не подписан на этом сервере
			response = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Все прошло хорошо, продолжаем выполнение запроса
		fmt.Sprintf("User %", tk.UserId) // полезно для мониторинга
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) // Передать управление следующему обработчику
	});
}