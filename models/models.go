package models

import "time"


type Manager struct {
    First_Name   string `json:"first_name"`
    Last_Name    string `json:"last_name"`
    Patronymic   string `json:"patronymic"`
    Function     string `json:"function"`
    Phone_Number string `json:"phone_number"`
    Email        string `json:"email"`
    Picture      string `json:"picture"`
    Priority     int32  `json:"priority"`
}


type News struct {
    Title   string    `json:"title"`
    Text    string    `json:"text"`
    Updated time.Time `json:"updated"`
}

type System struct {
    Short_Name  string `json:"short_name"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Partners struct {
    Name string `json:"name"`
    Link string `json:"link"`
    Logo string `json:"logo"`
}

type Organization struct {
    Info string `json:"info"`
}

