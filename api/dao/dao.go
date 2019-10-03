package dao

import (
	"babylon-stack/api/models"
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "babylon"
const COLLWAGE = "wage"
const COLLCOUNTRIES = "countries"

var db *mongo.Database
var collection = ""

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

func GetAll(data interface{}) interface{} {

	var elements []interface{}

	switch data.(type) {
	case models.Country:
		collection = "countries"
	case models.Minimumwage:
		collection = "wage"
	case models.Languages:
		collection = "languages"
	}

	cur, err := db.Collection(collection).Find(context.Background(), bson.D{}, nil)

	if err != nil {
		log.Fatal(err)
	}

	types := reflect.TypeOf(data)

	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		elem := reflect.New(types).Interface()
		err := cur.Decode(elem)
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

func GetItem(data interface{}, itemID string) interface{} {

	objID, _ := primitive.ObjectIDFromHex(itemID)
	filter := bson.D{{"_id", objID}}

	switch data.(type) {
	case models.Country:
		collection = "countries"
	case models.Minimumwage:
		collection = "wage"
	case models.Languages:
		collection = "languages"
	}

	types := reflect.TypeOf(data)
	elem := reflect.New(types).Interface()

	value := db.Collection(collection).FindOne(context.Background(), filter).Decode(elem)
	if value != nil {
		log.Fatal(value)
	}

	return elem
}

func UpdateItem(elem interface{}, itemID string) interface{} {
	objID, err := primitive.ObjectIDFromHex(itemID)
	filter := bson.D{{"_id", objID}}
	var update = bson.D{}

	switch elem.(type) {
	case *models.Country:
		fmt.Println("Model Country")
		collection = "countries"
		var c models.Country

		out, err := bson.Marshal(elem)
		if err != nil {
			fmt.Println(err)
		}
		bdoc := bson.Unmarshal([]byte(out), &c)
		if bdoc != nil {
			fmt.Println(bdoc)
		}

		update = bson.D{
			{"$set", c},
		}

	case *models.Minimumwage:
		fmt.Println("Model Wage")
		collection = "wage"
		var w models.Minimumwage

		out, err := bson.Marshal(elem)
		if err != nil {
			fmt.Println(err)
		}
		bdoc := bson.Unmarshal([]byte(out), &w)
		if bdoc != nil {
			fmt.Println(bdoc)
		}

		update = bson.D{
			{"$set", w},
		}
	case *models.Languages:
		fmt.Println("Model Languages")
		collection = "languages"
		var l models.Languages

		out, err := bson.Marshal(elem)
		if err != nil {
			fmt.Println(err)
		}
		bdoc := bson.Unmarshal([]byte(out), &l)
		if bdoc != nil {
			fmt.Println(bdoc)
		}

		update = bson.D{
			{"$set", l},
		}
	}

	updateResult, err := db.Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateResult)

	return elem
}

func AddItem(data interface{}) {

	switch data.(type) {
	case models.Country:
		collection = "countries"
		fmt.Println("Model Countries")
	case models.Minimumwage:
		fmt.Println("Model Wage")
		collection = "wage"
	case models.Languages:
		fmt.Println("Model Languages")
		collection = "languages"
	}

	_, err := db.Collection(collection).InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteItem(data interface{}) {
	switch data.(type) {
	case *models.Country:
		fmt.Println("Model Countries")
		collection = "countries"
	case *models.Minimumwage:
		fmt.Println("Model Wage")
		collection = "wage"
	case *models.Languages:
		fmt.Println("Model Languages")
		collection = "languages"
	}

	fmt.Println("Delete Collection :", collection)
	fmt.Println("Data to delete  :", data)

	deleteResult, err := db.Collection(collection).DeleteOne(context.Background(), data, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the collection\n", deleteResult.DeletedCount)
}
