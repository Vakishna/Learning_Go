package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json: "_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json: "firstname,omitempty" bson: "firstname,omitempty"`
	LastName  string             `json: "lastname,omitempty" bson: "lastname,omitempty"`
}

var client *mongo.Client

func main() {
	fmt.Println("Starting Application:")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	http.ListenAndServe(":3302", router)

}
