package router

import (
	"github.com/imirjar/mongo-golang/controllers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/organization", controllers.OrganizationHandler)
	router.HandleFunc("/managers", controllers.ManagersHandler)
	router.HandleFunc("/systems", controllers.SystemsHandler)
	router.HandleFunc("/news", controllers.NewsHandler)
	router.HandleFunc("/partners", controllers.PartnersHandler)
	return router
}