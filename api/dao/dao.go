package dao

import (
	"context"	
	"log"
	"babylon-stack/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "babylon"

// COLLNAME Collection name
const COLLNAME = "countries"

var db *mongo.Database



func init() {
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
    client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}


func GetAllCountries() []models.Country {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Country
	var elem models.Country
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}