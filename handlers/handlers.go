package handlers

import (
	"fmt"
	// "os"

	// "log"
	// "context"
	// "strconv"
	// "reflect"
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "path/filepath"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
	"github.com/imirjar/mongo-golang/models"
	"github.com/imirjar/mongo-golang/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// // Parse our multipart form, 10 << 20 specifies a maximum
	// // upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File", err)
		return
	}

	defer uploadedFile.Close()

	tempFile, err := ioutil.TempFile("storage", "*"+handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	json.NewEncoder(w).Encode(tempFile.Name())

	// client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer mongo.Close(client, ctx, cancel)

	// file := models.File{
	// 	Id:   primitive.NewObjectID(),
	// 	Name: handler.Filename,
	// 	Link: tempFile.Name(),
	// }

	// coll := client.Database("sspkSite").Collection(collection)
	// id, _ := primitive.ObjectIDFromHex(documentId)
	// filter := bson.D{{"_id", id}}
	// update := bson.D{{"$push", bson.D{{"documents", file}}}}
	// result, err := coll.UpdateOne(context.TODO(), filter, update)
	// // fmt.Println(result)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// // fmt.Printf("File Size: %+v\n", handler.Size)
	// // fmt.Printf("MIME Header: %+v\n", handler.Header)
	// json.NewEncoder(w).Encode(result)
	// // fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	// var file models.File
	// err := json.NewDecoder(r.Body).Decode(&file)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// vars := mux.Vars(r)
	// collection := vars["collection"]
	// elementId, err := primitive.ObjectIDFromHex(vars["elementId"])
	// if err != nil {
	// 	fmt.Printf("Can't make primitive %v\n", err)
	// }

	// client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer mongo.Close(client, ctx, cancel)

	// coll := client.Database("sspkSite").Collection(collection)

	// filter := bson.M{
	// 	"_id": elementId,
	// 	"documents": bson.M{
	// 		"$elemMatch": bson.M{
	// 			"_id": file.Id,
	// 		},
	// 	},
	// }

	// update := bson.M{
	// 	"$pull": bson.M{
	// 		"documents": bson.M{
	// 			"_id": file.Id,
	// 		},
	// 	},
	// }

	// result, err := coll.UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = os.Remove(file.Link)
	// if err != nil {
	// 	fmt.Println("Ну удалось удалить файл", err)
	// }

	// json.NewEncoder(w).Encode(result)

}

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
		obj := mongo.GetData("news", news, bson.M{"_id": artileId})
		json.NewEncoder(w).Encode(obj)
	case "PUT":
		var news models.News
		err := json.NewDecoder(r.Body).Decode(&news)
		if err != nil {
			fmt.Println(err)
		}
		update := bson.M{"$set": news}
		filter := bson.M{"_id": news.Id}
		obj := mongo.SetData("news", news, filter, update)
		json.NewEncoder(w).Encode(obj)
	}
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var news []models.News
		obj := mongo.GetData("news", news, nil)
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
		obj := mongo.GetData("systems", systems, bson.M{"_id": systemId})
		json.NewEncoder(w).Encode(obj)
	case "PUT":
		var system models.System
		err := json.NewDecoder(r.Body).Decode(&system)
		if err != nil {
			fmt.Println(err)
		}
		update := bson.M{"$set": system}
		filter := bson.M{"_id": system.Id}
		obj := mongo.SetData("systems", system, filter, update)
		json.NewEncoder(w).Encode(obj)
	}
}

func SystemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var systems []models.System
		obj := mongo.GetData("systems", systems, nil)
		json.NewEncoder(w).Encode(obj)
	}
}

func PartnersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var partners []models.Partner
		obj := mongo.GetData("partners", partners, nil)
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
		obj := mongo.GetData("managers", managers, bson.M{"_id": managerId})
		json.NewEncoder(w).Encode(obj)
	case "PUT":
		var manager models.Manager
		err := json.NewDecoder(r.Body).Decode(&manager)
		if err != nil {
			fmt.Println(err)
		}
		update := bson.M{"$set": manager}
		filter := bson.M{"_id": manager.Id}
		obj := mongo.SetData("managers", manager, filter, update)
		json.NewEncoder(w).Encode(obj)
	}
}

func ManagersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var managers []models.Manager
		obj := mongo.GetData("managers", managers, nil)
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
		obj := mongo.GetData("documents", documents, bson.M{"_id": documentId})
		json.NewEncoder(w).Encode(obj)
	case "PUT":
		var document models.Document
		err := json.NewDecoder(r.Body).Decode(&document)
		if err != nil {
			fmt.Println(err)
		}
		update := bson.M{"$set": document}
		filter := bson.M{"_id": document.Id}
		obj := mongo.SetData("documents", document, filter, update)
		json.NewEncoder(w).Encode(obj)
	}
}

func DocumentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var documents []models.Document
		obj := mongo.GetData("documents", documents, nil)
		json.NewEncoder(w).Encode(obj)
	case "POST":
		fmt.Println("POST documents")
	}
}

func DocumentsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// documentsType := vars["type"]
	// var filter bson.D
	// if documentsType != "" {
	// 	filter = bson.D{{"type", documentsType}}
	// } else {
	// 	filter = bson.D{{}}
	// }

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Printf("Error while parsing .env file: %v\n", err)
	// }

	// client, ctx, cancel, err := mongo.Connect(os.Getenv("MONGODB_URL"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer mongo.Close(client, ctx, cancel)

	// var documents []models.Document

	// coll := client.Database("sspkSite").Collection("documents")

	// cursor, err := coll.Find(context.TODO(), filter)
	// if err != nil {
	// 	panic(err)
	// }

	// if err = cursor.All(context.TODO(), &documents); err != nil {
	// 	panic(err)
	// }

	// json.NewEncoder(w).Encode(documents)
}

func OrganizationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var organizations []models.Organization
		obj := mongo.GetData("organization", organizations, nil)
		json.NewEncoder(w).Encode(obj)
	case "PUT":
		var organization models.Organization
		err := json.NewDecoder(r.Body).Decode(&organization)
		if err != nil {
			fmt.Println(err)
		}

		update := bson.M{"$set": organization}
		filter := bson.M{"_id": organization.Id}
		obj := mongo.SetData("organization", organization, filter, update)
		json.NewEncoder(w).Encode(obj)
	}
}

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var systems []models.System
		obj := mongo.GetData("systems", systems, nil)
		json.NewEncoder(w).Encode(obj)
	case "POST":
		var systems []models.System
		obj := mongo.GetData("systems", systems, nil)
		json.NewEncoder(w).Encode(obj)
	case "DELETE":
		var systems []models.System
		obj := mongo.GetData("systems", systems, nil)
		json.NewEncoder(w).Encode(obj)
	}
}