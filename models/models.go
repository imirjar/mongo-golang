package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoConnection struct {
	Database   string
	Collection string
	Filter     bson.M
	Update     bson.M
}

type Manager struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	First_Name   string             `json:"first_name"`
	Last_Name    string             `json:"last_name"`
	Patronymic   string             `json:"patronymic"`
	Function     string             `json:"function"`
	Phone_Number string             `json:"phone_number"`
	Email        string             `json:"email"`
	Picture      string             `json:"picture"`
	Priority     int32              `json:"priority"`
}

type File struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name"`
	Link    string             `json:"link"`
	Updated time.Time          `json:"updated"`
}

type News struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Title   string             `json:"title"`
	Text    string             `json:"text"`
	Updated time.Time          `json:"updated"`
	Files   []File             `json:"files"`
	Media   []File             `json:"media"`
}

type System struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Short_Name  string             `json:"short_name"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}

type Partner struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
	Link string             `json:"link"`
	Logo string             `json:"logo"`
}

type Social struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
	Link string             `json:"link"`
}

type Organization struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Info          string             `json:"info"`
	Name          string             `json:"name"`
	Full_name     string             `json:"full_name"`
	Contact_phone string             `json:"contact_phone"`
	Support_phone string             `json:"support_phone"`
	Email         string             `json:"email"`
	Files         []File             `json:"files"`
	Social        []Social           `json:"social"`
}

type Document struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Files       []File             `json:"files"`
}
