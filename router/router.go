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
	router.HandleFunc("/manager/{id}", controllers.ManagerHandler)
	router.HandleFunc("/systems", controllers.SystemsHandler)
	router.HandleFunc("/system/{id}", controllers.SystemHandler)
	router.HandleFunc("/news", controllers.NewsHandler)
	router.HandleFunc("/news/{id}", controllers.ArticleHandler)
	router.HandleFunc("/partners", controllers.PartnersHandler)
	router.HandleFunc("/document/{id}", controllers.DocumentHandler)
	router.HandleFunc("/documents", controllers.DocumentsHandler)
	// router.HandleFunc("/upload", controllers.UploadFile)
	return router
}