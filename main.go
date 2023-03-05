package main
 
import (
    
    "fmt"
    "github.com/rs/cors"
    "log"
    "net/http"

    "github.com/imirjar/mongo-golang/router"
)
 

func main(){


    r := router.Router()
 
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8000"},
        AllowCredentials: true,
    })

    handler := c.Handler(r)
    // documents = []interface{}{
    //     bson.D{
    //         {"rollNo", 153},
    //         {"maths", 65},
    //         {"science", 59},
    //         {"computer", 55},
    //     },
    //     bson.D{
    //         {"rollNo", 162},
    //         {"maths", 86},
    //         {"science", 80},
    //         {"computer", 69},
    //     },
    // }
     
    // insertOneResult, err := insertOne(client, ctx, "sspkSite", "managers", manager)
    // insertOneResult, err = insertOne(client, ctx, "sspkSite", "gises", gis)
    // insertManyResult, err := insertMany(client, ctx, "gfg", "marks", documents)
     
    // handle the error
    // if err != nil {
    //     panic(err)
    // }
     
    // print the insertion id of the document,
    // if it is inserted.
    // fmt.Println("Result of InsertOne")
    // fmt.Println(insertOneResult.InsertedID)
    fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", handler))
 

}