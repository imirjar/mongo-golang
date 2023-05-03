package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Modeller interface{}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func GetData(collectionName string, model Modeller, filter primitive.M) Modeller {

	client, ctx, cancel, err := connect(os.Getenv("MONGODB_URL"))
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(nil))

	// cursor, err := mongo.Query(client, ctx, db, collectionName, filter, nil)
	// if err != nil {
	//     panic(err)
	// }

	if err := cursor.All(ctx, &model); err != nil {
		fmt.Println(err)
	}

	return model
}

func SetData(collectionName string, filter primitive.M, update primitive.M) Modeller {

	//подключаемся к Mongodb по переменной подключения из env файла
	client, ctx, cancel, err := connect(os.Getenv("MONGODB_URL"))
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	coll := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collectionName)

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return result
}
