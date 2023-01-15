package routes

import (
	"BE-finaltask/handlers"
	"BE-finaltask/pkg/middleware"
	"BE-finaltask/pkg/mysql"
	"BE-finaltask/repositories"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router) {

	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	r.HandleFunc("/articles", h.FindArticle).Methods("GET")
	r.HandleFunc("/articles/{id}", h.GetArticle).Methods("GET")
	r.HandleFunc("/articles", middleware.Auth(middleware.UploadFile(h.AddArticle))).Methods("POST")
	r.HandleFunc("/articles/{id}", h.DeleteArticle).Methods("DELETE")
}
