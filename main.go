package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/imirjar/mongo-golang/router"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error while parsing .env file: %v\n", err)
	}

	r := router.Router()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://0.0.0.0", "http://localhost", "http://127.0.0.1", "http://192.168.198.103/", "http://nginx"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})
	handler := c.Handler(r)

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
