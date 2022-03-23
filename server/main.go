package main

import (
	"context"
	"fmt"
	"log"
	"time"
	// "where-am-i-server/Models"
	"where-am-i-server/Routes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, context.Context) {
    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

    if err != nil {
        log.Fatal(err)
    }

    return client, ctx
}

func main() {
    DB, ctx := ConnectDB()

    var result []bson.D
    filter := bson.D {}
    option := bson.D {}

    cursor, _ := DB.Database("WhereAmI").Collection("Students").Find(ctx, filter, options.Find().SetProjection(option))

    if err := cursor.All(ctx, &result); err != nil {
        log.Fatal(err)
    }

    for _, doc := range result {
        fmt.Println(doc)
    }
    // fmt.Printf("%+v\n", result)

    router := Routes.SetupRouter()
    router.Run("localhost:8080")
}
