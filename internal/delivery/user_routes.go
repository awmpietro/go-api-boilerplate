package delivery

import (
	"net/http"
	"teste-go/internal/middleware"
)

func RegisterUserRoutes(userHandler UserHandler, mux *http.ServeMux) {
	mux.HandleFunc("/users", middleware.JWTMiddleware(userHandler.GetUsers))
	mux.HandleFunc("/users/create", userHandler.CreateUser)
	mux.HandleFunc("/users/update", userHandler.UpdateUser)
	mux.HandleFunc("/users/delete", userHandler.DeleteUser)
	mux.HandleFunc("/users/login", userHandler.Login)
}
