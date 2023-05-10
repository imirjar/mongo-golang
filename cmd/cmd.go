package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/imirjar/mongo-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveUploadedFileToStorage(r *http.Request) (models.File, error) {
	r.ParseMultipartForm(10 << 20)
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File", err)
	}
	if err != nil {
		fmt.Println(err)
	}

	defer uploadedFile.Close()

	tempFile, err := ioutil.TempFile("storage", "*"+handler.Filename)
	if err != nil {
		fmt.Println("Error Retrieving the File", err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		fmt.Println("Error Retrieving the File", err)
	}
	tempFile.Write(fileBytes)

	file := models.File{
		Id:   primitive.NewObjectID(),
		Name: handler.Filename,
		Link: tempFile.Name(),
	}

	return file, err
}

// func DeleteUploadedFile(*http.Request) error {}
