package controllers


import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "go.mongodb.org/mongo-driver/bson"
    "github.com/joho/godotenv"

    "github.com/gorilla/mux"
    "github.com/imirjar/mongo-golang/mongo"
    "github.com/imirjar/mongo-golang/models"
    "go.mongodb.org/mongo-driver/bson/primitive" 
)


func UploadFile(w http.ResponseWriter, r *http.Request) {
    

    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()


//
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
    err = godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }

    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)


    document := models.Document{
        Name: handler.Filename,
        Link : "/document_storage/handler.Filename",
    }

    insertOneResult, err := mongo.InsertOne(client, ctx, "sspkSite", "documents", document)
    // insertManyResult, err := insertMany(client, ctx, "gfg", "marks", documents)
     
    // handle the error
    if err != nil {
        panic(err)
    }
     
    // print the insertion id of the document,
    // if it is inserted.
    fmt.Println("Result of InsertOne")
    fmt.Println(insertOneResult.InsertedID)

//



    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("document_storage", "*"+handler.Filename)
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}


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

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    propsId := vars["id"]

    artileId, err := primitive.ObjectIDFromHex(propsId)
    if err != nil {
        fmt.Printf("Can't make primirive %v\n", err)
    }

    err = godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }
    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, "sspkSite", "news", bson.M{"_id": artileId}, nil)
    if err != nil {
        panic(err)
    }
    
    var news []models.News
     
    if err := cursor.All(ctx, &news); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(news[0]) 
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
    
    var partners []models.Partner
     
    if err := cursor.All(ctx, &partners); err != nil {
         fmt.Println(err)
    }
    
    json.NewEncoder(w).Encode(partners) 
}


func DocumentsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        fmt.Println("I'm made of GET!",)
        err := godotenv.Load(".env")
        if err != nil {
            fmt.Printf("Error while parsing .env file: %v\n", err)
        }

        client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
        if err != nil {
            panic(err)
        }
        defer mongo.Close(client, ctx, cancel)

        cursor, err := mongo.Query(client, ctx, "sspkSite", "documents", bson.D{}, nil)
        if err != nil {
            panic(err)
        }
        
        var documents []models.DocumentsType
         
        if err := cursor.All(ctx, &documents); err != nil {
             fmt.Println(err)
        }
        
        json.NewEncoder(w).Encode(documents) 
    } else {}    
}
