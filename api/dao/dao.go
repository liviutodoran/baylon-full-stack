package dao

import (
	"babylon-stack/api/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "babylon"

// COLLNAME Collection name
const COLLCOUNTRIES = "countries"
const COLLWAGE = "wage"

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
	cur, err := db.Collection(COLLCOUNTRIES).Find(context.Background(), bson.D{}, nil)
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

func GetWageMongo() []models.Minimumwage {
	cur, err := db.Collection(COLLWAGE).Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Minimumwage
	var elem models.Minimumwage
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

func DeleteWage(wage models.Minimumwage) {
	_, err := db.Collection(COLLWAGE).DeleteOne(context.Background(), wage, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateWage(wage models.Minimumwage, wageID string) {

	filter := bson.D{{"id", wageID}}
	fmt.Println(filter)
	update := bson.D{
		{"$set", bson.D{
			{"Country", wage.Country},
			{"Year", wage.Year},
			{"LocalAmount", wage.LocalAmount},
			{"USD", wage.USD},
		}},
	}

	updateResult, err := db.Collection(COLLWAGE).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateResult)
}
