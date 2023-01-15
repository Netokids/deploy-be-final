package routes

import (
	"BE-finaltask/handlers"
	"BE-finaltask/pkg/middleware"
	"BE-finaltask/pkg/mysql"
	"BE-finaltask/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {

	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUser).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", middleware.UploadFile(h.UpdateUser)).Methods("PATCH")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}
