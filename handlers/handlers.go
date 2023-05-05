package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
	"github.com/imirjar/mongo-golang/cmd"
	"github.com/imirjar/mongo-golang/models"
	"github.com/imirjar/mongo-golang/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		obj := mongo.SetData("news", filter, update)
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
		obj := mongo.SetData("systems", filter, update)
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
		obj := mongo.SetData("managers", filter, update)
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
		obj := mongo.SetData("documents", filter, update)
		json.NewEncoder(w).Encode(obj)
	case "PATCH":
		var file models.File

		err := json.NewDecoder(r.Body).Decode(&file)
		if err != nil {
			fmt.Println(err)
		}

		filter := bson.M{"_id": documentId}
		update := bson.M{"$push": bson.M{"documents": file}}

		// fmt.Println(document, filter, update)
		obj := mongo.SetData("documents", filter, update)
		json.NewEncoder(w).Encode(obj)
	case "DELETE":
		var document models.Document
		var file models.File

		err := json.NewDecoder(r.Body).Decode(&file)
		if err != nil {
			fmt.Println(err)
		}

		filter := bson.M{"_id": documentId}
		update := bson.M{"$pop": bson.M{"documents": file}}

		fmt.Println(document, filter, update)
		// obj := mongo.SetData("documents", document, filter, update)
		// json.NewEncoder(w).Encode(obj)
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
	vars := mux.Vars(r)
	docType := vars["type"]
	var documents []models.Document
	filter := bson.M{"type": docType}
	obj := mongo.GetData("documents", documents, filter)
	json.NewEncoder(w).Encode(obj)
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
		obj := mongo.SetData("organization", filter, update)
		json.NewEncoder(w).Encode(obj)
	}
}

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		file := cmd.SaveUploadedFileToStorage(r)

		collectionName := r.FormValue("collectionName")
		collectionId, err := primitive.ObjectIDFromHex(r.FormValue("collectionId"))
		if err != nil {
			fmt.Println(err)
		}

		filter := bson.M{"_id": collectionId}
		update := bson.M{"$push": bson.M{"files": file}}

		mongo.SetData(collectionName, filter, update)
		json.NewEncoder(w).Encode(file)

	case "DELETE":
		documentId, err := primitive.ObjectIDFromHex(r.FormValue("documentId"))
		if err != nil {
			fmt.Println("Не удалось получить ID документа коллекции", err)
		}
		fileId, err := primitive.ObjectIDFromHex(r.FormValue("fileId"))
		if err != nil {
			fmt.Println("Не удалось получить ID файла", err)
		}
		fileLink := r.FormValue("fileLink")
		collectionName := r.FormValue("collectionName")

		//remove file from collection

		filter := bson.M{
			"_id": documentId,
			"files": bson.M{
				"$elemMatch": bson.M{
					"_id": fileId,
				},
			},
		}
		update := bson.M{
			"$pull": bson.M{
				"files": bson.M{
					"_id": fileId,
				},
			},
		}

		obj := mongo.SetData(collectionName, filter, update)

		//remove file from storage
		err = os.Remove(fileLink)
		if err != nil {
			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(obj)
	}
}
