package http

import (
	"CleanArchitecture/internal/handlers"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, userHandler *handlers.UserHandler) {

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUser(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}
