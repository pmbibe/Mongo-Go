package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo(db string, collection string, url string) (*mongo.Collection, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	quickstartDatabase := client.Database(db)
	pCollection := quickstartDatabase.Collection(collection)

	return pCollection, ctx
}

func updateMongo(pCollection *mongo.Collection, ctx context.Context, newURL string, oldURL string, sentence Sentence, key string) {
	result, err := pCollection.UpdateMany(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: bson.D{{Key: key, Value: fmt.Sprintf(strings.Replace(sentence.Record.Audio, oldURL, newURL, 1))}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.MatchedCount)
}
