package routes

import (
	"BE-finaltask/handlers"
	"BE-finaltask/pkg/middleware"
	"BE-finaltask/pkg/mysql"
	"BE-finaltask/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/registerdoctor", h.RegisterDoctor).Methods("POST")
	r.HandleFunc("/Login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
