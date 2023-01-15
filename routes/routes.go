package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoute(r)
	AuthRoutes(r)
	ArticleRoutes(r)
	ConsultationRoutes(r)
}
