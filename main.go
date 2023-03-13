package main
 
import (
    "fmt"
    "log"
    "net/http"
    "github.com/rs/cors"
    "github.com/imirjar/mongo-golang/router"
)
 

func main(){
    r := router.Router()
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
    })
    handler := c.Handler(r)

    fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", handler))

}