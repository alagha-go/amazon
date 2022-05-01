package main

import (
	"amazon/lib/crawler"
	"amazon/lib/handler"
	"amazon/lib/types"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	port = ":6000"
	mongodb = "mongodb://127.0.0.1:27017/"
)

func main() {
	go crawler.Main()
	var err error
	fmt.Println("starting server...")

	clientOptions := options.Client().ApplyURI(mongodb)
	ctx := context.Background()
	types.Client, err = mongo.Connect(ctx, clientOptions)
	HandleError(err)

	go handler.Main()

	err = http.ListenAndServe(port, nil)
	HandleError(err)
}


func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}