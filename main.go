package main

import (
    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2"
    "net/http"

    "github.com/imirjar/golang_45_killer_projects/tree/master/mongo-golang/controllers"
)


func main() {
    r := httprouter.New()
    uc := controllers.NewsUserController(getSession())
    r.GET("/uer/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.DeleteUser)
    http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
    s, err := mgo.Dial("mongodb://mongo:27107")
    if err != nil {
        panic(err)
    }
    return s
}