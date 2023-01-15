package routes

import (
	"BE-finaltask/handlers"
	"BE-finaltask/pkg/middleware"
	"BE-finaltask/pkg/mysql"
	"BE-finaltask/repositories"

	"github.com/gorilla/mux"
)

func ConsultationRoutes(r *mux.Router) {

	consultationRepository := repositories.RepositoryConsultation(mysql.DB)
	h := handlers.HandlerConsultation(consultationRepository)

	r.HandleFunc("/consultations", h.FindConsultation).Methods("GET")
	r.HandleFunc("/consultations/{id}", h.GetConsultation).Methods("GET")
	r.HandleFunc("/consultations", middleware.Auth(h.CreateConsultation)).Methods("POST")
	r.HandleFunc("/consultations/{id}", h.UpdateConsultation).Methods("PATCH")
}
