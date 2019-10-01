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

func GetAll(data interface{}) interface{} {

	var collection = ""
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
	var collection = ""

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

func UpdateItem(data interface{}, itemID string) interface{} {
	objID, err := primitive.ObjectIDFromHex(itemID)
	filter := bson.D{{"_id", objID}}
	var collection = ""
	var update = bson.D{}
	var elem = ""

	switch data.(type) {
	case models.Country:
		var elem models.Country
		collection = "countries"
		update := bson.D{
			{"$set", bson.D{
				{"languages", elem.Languages},
				{"country", elem.Country},
				{"country_id", elem.Country_id},
				{"Capital", elem.Capital},
				{"currency_name", elem.Currency_name},
				{"currency_symbol", elem.Currency_symbol},
				{"currency_code", elem.Currency_code},
				{"iso", elem.Iso},
			}},
		}
		fmt.Println(update)
	case models.Minimumwage:
		collection = "wage"
	case models.Languages:
		collection = "languages"
	}

	updateResult, err := db.Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateResult)

	return elem
}

func AddCountry(country models.Country) {
	_, err := db.Collection(COLLCOUNTRIES).InsertOne(context.Background(), country)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteCountry(country models.Country) {
	deleteResult, err := db.Collection(COLLCOUNTRIES).DeleteOne(context.Background(), country, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the country collection\n", deleteResult.DeletedCount)
}

func GetAllWage() interface{} {
	cur, err := db.Collection(COLLWAGE).Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Minimumwage

	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		var elem models.Minimumwage
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

func GetWage(wage models.Minimumwage, wageID string) models.Minimumwage {

	objID, _ := primitive.ObjectIDFromHex(wageID)
	filter := bson.D{{"_id", objID}}
	value := db.Collection(COLLWAGE).FindOne(context.Background(), filter).Decode(&wage)
	if value != nil {
		log.Fatal(value)
	}

	return wage
}

func DeleteWage(wage models.Minimumwage) {
	deleteResult, err := db.Collection(COLLWAGE).DeleteOne(context.Background(), wage, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the wage collection\n", deleteResult.DeletedCount)
}

func UpdateWage(wage models.Minimumwage, wageID string) models.Minimumwage {
	objID, err := primitive.ObjectIDFromHex(wageID)
	filter := bson.D{{"_id", objID}}
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

	return wage
}

func AddWage(wage models.Minimumwage) {
	_, err := db.Collection(COLLWAGE).InsertOne(context.Background(), wage)
	if err != nil {
		log.Fatal(err)
	}
}
