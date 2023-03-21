package controllers


import (
    "os"
    "fmt"
    "context"
    // "strconv"
    // "reflect"
    "net/http"
    // "io/ioutil"
    "encoding/json"

    "go.mongodb.org/mongo-driver/bson"
    "github.com/joho/godotenv"

    "github.com/gorilla/mux"
    "github.com/imirjar/mongo-golang/mongo"
    "github.com/imirjar/mongo-golang/models"
    "go.mongodb.org/mongo-driver/bson/primitive" 
)


type Modeller interface {}


//connect to MongoDB and get data by params
func getData(db string, table string, obj Modeller, bsonV  primitive.M) Modeller {

    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }

    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)

    cursor, err := mongo.Query(client, ctx, db, table, bsonV, nil)
    if err != nil {
        panic(err)
    }
     
    if err := cursor.All(ctx, &obj); err != nil {
         fmt.Println(err)
    }
    
    return obj 
}

func putData(db string, table string, obj Modeller, id  primitive.ObjectID) Modeller {
    //подгружаем файл env   
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error while parsing .env file: %v\n", err)
    }
    //подключаемся к Mongodb по переменной подключения из env файла
    client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
    if err != nil {
        panic(err)
    }
    defer mongo.Close(client, ctx, cancel)


    coll := client.Database(db).Collection(table)
    update := bson.D{{"$set", obj}}
    result, err := coll.UpdateOne(context.TODO(), bson.D{{"_id", id}}, update)
    if err != nil {
        panic(err)
    }
    return result
}


// func UploadFile(w http.ResponseWriter, r *http.Request) {
    

//     fmt.Println("File Upload Endpoint Hit")

//     // Parse our multipart form, 10 << 20 specifies a maximum
//     // upload of 10 MB files.
//     r.ParseMultipartForm(10 << 20)
//     // FormFile returns the first file for the given key `myFile`
//     // it also returns the FileHeader so we can get the Filename,
//     // the Header and the size of the file
//     file, handler, err := r.FormFile("myFile")
//     if err != nil {
//         fmt.Println("Error Retrieving the File")
//         fmt.Println(err)
//         return
//     }
//     defer file.Close()

//     err = godotenv.Load(".env")
//     if err != nil {
//         fmt.Printf("Error while parsing .env file: %v\n", err)
//     }

//     client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
//     if err != nil {
//         panic(err)
//     }
//     defer mongo.Close(client, ctx, cancel)


//     file := models.File{
//         Name: handler.Filename,
//         Link : "/document_storage/handler.Filename",
//     }

//     insertOneResult, err := mongo.InsertOne(client, ctx, "sspkSite", "documents", file)
//     // insertManyResult, err := insertMany(client, ctx, "gfg", "marks", documents)
     
//     // handle the error
//     if err != nil {
//         panic(err)
//     }
     
//     // print the insertion id of the document,
//     // if it is inserted.
//     fmt.Println("Result of InsertOne")
//     fmt.Println(insertOneResult.InsertedID)

// //



//     fmt.Printf("Uploaded File: %+v\n", handler.Filename)
//     fmt.Printf("File Size: %+v\n", handler.Size)
//     fmt.Printf("MIME Header: %+v\n", handler.Header)

//     // Create a temporary file within our temp-images directory that follows
//     // a particular naming pattern
//     tempFile, err := ioutil.TempFile("document_storage", "*"+handler.Filename)
//     if err != nil {
//         fmt.Println(err)
//     }
//     defer tempFile.Close()

//     // read all of the contents of our uploaded file into a
//     // byte array
//     fileBytes, err := ioutil.ReadAll(file)
//     if err != nil {
//         fmt.Println(err)
//     }
//     // write this byte array to our temporary file
//     tempFile.Write(fileBytes)
//     // return that we have successfully uploaded our file!
//     fmt.Fprintf(w, "Successfully Uploaded File\n")
// }

func ArticleHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    propsId := vars["id"]

    artileId, err := primitive.ObjectIDFromHex(propsId)
    if err != nil {
        fmt.Printf("Can't make primitive %v\n", err)
    }

    switch r.Method {
    case "GET": 
        var news []models.News
        obj := getData("sspkSite", "news", news, bson.M{"_id": artileId})
        json.NewEncoder(w).Encode(obj)
    case "PUT": 
        var news models.News  
        err := json.NewDecoder(r.Body).Decode(&news)
                if err != nil {
            fmt.Println(err)
        }
        obj := putData("sspkSite", "news", news, news.Id)
        json.NewEncoder(w).Encode(obj)
    }
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET": 
        var news []models.News
        obj := getData("sspkSite", "news", news, nil)
        json.NewEncoder(w).Encode(obj)
    }
}

func SystemHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    propsId := vars["id"]

    systemId, err := primitive.ObjectIDFromHex(propsId)
    if err != nil {
        fmt.Printf("Can't make primitive %v\n", err)
    }

    switch r.Method {
    case "GET": 
        var systems []models.System
        obj := getData("sspkSite", "systems", systems, bson.M{"_id": systemId})
        json.NewEncoder(w).Encode(obj)
    case "PUT": 
        var system models.System  
        err := json.NewDecoder(r.Body).Decode(&system)
                if err != nil {
            fmt.Println(err)
        }
        obj := putData("sspkSite", "systems", system, system.Id)
        json.NewEncoder(w).Encode(obj)
    }
}

func SystemsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET": 
        var systems []models.System
        obj := getData("sspkSite", "systems", systems, nil)
        json.NewEncoder(w).Encode(obj)
    }
}

func PartnersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET": 
        var partners []models.Partner
        obj := getData("sspkSite", "partners", partners, nil)
        json.NewEncoder(w).Encode(obj)
    }
}

func ManagerHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    propsId := vars["id"]

    managerId, err := primitive.ObjectIDFromHex(propsId)
    if err != nil {
        fmt.Printf("Can't make primitive %v\n", err)
    }

    switch r.Method {
    case "GET": 
        var managers []models.Manager
        obj := getData("sspkSite", "managers", managers, bson.M{"_id": managerId})
        json.NewEncoder(w).Encode(obj)
    case "PUT": 
        var manager models.Manager  
        err := json.NewDecoder(r.Body).Decode(&manager)
        if err != nil {
            fmt.Println(err)
        }
        obj := putData("sspkSite", "managers", manager, manager.Id)
        json.NewEncoder(w).Encode(obj)
    }
}

func ManagersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET": 
        var managers []models.Manager
        obj := getData("sspkSite", "managers", managers, nil)
        json.NewEncoder(w).Encode(obj)
    }
}

func DocumentHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    propsId := vars["id"]

    documentId, err := primitive.ObjectIDFromHex(propsId)
    if err != nil {
        fmt.Printf("Can't make primitive %v\n", err)
    }

    switch r.Method {
    case "GET": 
        var documents []models.Document
        obj := getData("sspkSite", "documents", documents, bson.M{"_id": documentId})
        json.NewEncoder(w).Encode(obj)
    case "PUT": 
        var document models.Document  
        err := json.NewDecoder(r.Body).Decode(&document)
        if err != nil {
            fmt.Println(err)
        }
        obj := putData("sspkSite", "documents", document, document.Id)
        json.NewEncoder(w).Encode(obj)
    }
}

func DocumentsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET": 
        fmt.Println("GET documents")
        var documents []models.DocumentsType
        obj := getData("sspkSite", "documents", documents, nil)
        json.NewEncoder(w).Encode(obj)
    case "POST": 
        fmt.Println("POST documents",)
    }
}

func OrganizationHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET": 
        var organizations []models.Organization
        obj := getData("sspkSite", "organization", organizations, nil)
        json.NewEncoder(w).Encode(obj)
    case "PUT": 
        var organization models.Organization  
        err := json.NewDecoder(r.Body).Decode(&organization)
                if err != nil {
            fmt.Println(err)
        }
        obj := putData("sspkSite", "organization", organization, organization.Id)
        json.NewEncoder(w).Encode(obj)
    }
}