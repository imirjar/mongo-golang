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
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
