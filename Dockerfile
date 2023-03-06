# Базовый образ
FROM golang:1.14.4-buster
# Папка приложения
ARG APP_DIR=app
# Копирование файлов
COPY . /go/tmp/src/${APP_NAME}
# Рабочая директория
WORKDIR /go/tmp/src/${APP_NAME}
# Библиотеки
RUN go get "github.com/imirjar/mongo-golang/router"
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/joho/godotenv"
RUN go get "go.mongodb.org/mongo-driver/bson"
RUN go get "go.mongodb.org/mongo-driver/mongo"
RUN go get "go.mongodb.org/mongo-driver/mongo/options"
RUN go get "go.mongodb.org/mongo-driver/mongo/readpref"