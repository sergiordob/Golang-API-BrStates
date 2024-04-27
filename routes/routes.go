package routes

import (
	"net/http"
	"github.com/sergiordob/Golang-API-REST-01/controllers"
)

func LoadRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api", controllers.Index())
	mux.HandleFunc("GET /api/{id}", controllers.GetOneState())
	mux.HandleFunc("GET /api/total", controllers.GetAllStates())
}
