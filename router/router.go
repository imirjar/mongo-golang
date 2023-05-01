package router

import (
	"github.com/gorilla/mux"
	"github.com/imirjar/mongo-golang/handlers"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/organization", handlers.OrganizationHandler)
	router.HandleFunc("/managers", handlers.ManagersHandler)
	router.HandleFunc("/manager/{id}", handlers.ManagerHandler)
	router.HandleFunc("/systems", handlers.SystemsHandler)
	router.HandleFunc("/system/{id}", handlers.SystemHandler)
	router.HandleFunc("/news", handlers.NewsHandler)
	router.HandleFunc("/news/{id}", handlers.ArticleHandler)
	router.HandleFunc("/partners", handlers.PartnersHandler)
	router.HandleFunc("/document/{id}", handlers.DocumentHandler)
	router.HandleFunc("/documents", handlers.DocumentsHandler)
	router.HandleFunc("/documents/{type}", handlers.DocumentsByCategoryHandler)
	router.HandleFunc("/upload", handlers.UploadFile)
	router.HandleFunc("/delete/{collection}/{elementId}", handlers.DeleteFile)
	router.HandleFunc("/files", handlers.FilesHandler)
	return router
}
