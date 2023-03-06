package controllers


import (
    "os"
    "fmt"
    "net/http"
    "encoding/json"

    "go.mongodb.org/mongo-driver/bson"
    "github.com/joho/godotenv"

    // "github.com/gorilla/mux"
    "github.com/imirjar/mongo-golang/mongo"
    "github.com/imirjar/mongo-golang/models"
)


func OrganizationHandler(w http.ResponseWriter, r *http.Request) {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }

    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "organization", bson.D{}, nil)
    if err != nil {
        panic(err)
    }
    
    var organizations []models.Organization
     
    if err := cursor.All(ctx, &organizations); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(organizations[0]) 
}

func ManagersHandler(w http.ResponseWriter, r *http.Request) {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }

    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
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
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }

    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "systems", bson.D{}, nil)
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
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }
    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
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

func PartnersHandler(w http.ResponseWriter, r *http.Request) {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }

    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "partners", bson.D{}, nil)
    if err != nil {
        panic(err)
    }
    
    var partners []models.Partners
     
    if err := cursor.All(ctx, &partners); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(partners) 
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