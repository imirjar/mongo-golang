package controllers


import (
    "encoding/json"
    "fmt"
    "net/http"

    "go.mongodb.org/mongo-driver/bson"

    // "github.com/gorilla/mux"
    "github.com/imirjar/mongo-golang/mongo"
    "github.com/imirjar/mongo-golang/models"
)

func ManagersHandler(w http.ResponseWriter, r *http.Request) {

    client, ctx, cancel, err := mongo.Connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "managers", bson.D{}, nil)
    if err != nil {
        panic(err)
    }
    
    var managers []models.Manager
     
    if err := cursor.All(ctx, &managers); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(managers) 
}


func SystemsHandler(w http.ResponseWriter, r *http.Request) {
    client, ctx, cancel, err := mongo.Connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "gises", bson.D{}, nil)
    if err != nil {
        panic(err)
    }
    
    var systems []models.System
     
    if err := cursor.All(ctx, &systems); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(systems) 

}


func NewsHandler(w http.ResponseWriter, r *http.Request) {
    client, ctx, cancel, err := mongo.Connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "news", bson.D{}, nil)
    if err != nil {
        panic(err)
    }
    
    var news []models.News
     
    if err := cursor.All(ctx, &news); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(news) 
}




    // vars := mux.Vars(r)
    // w.WriteHeader(http.StatusOK)
    // fmt.Fprintf(w, "Category: %v\n", vars["category"])

    // var manager models.Manager
    // err = json.NewDecoder(r.Body).Decode(&manager)

    // if err != nil {
    //     log.Fatalf("Unable to decode the request body.  %v", err)
    // }

    // //send response
    // res := models.Response{
    //     Message: "Notification/mails added to redis query",
    // }

    // json.NewEncoder(w).Encode(res)