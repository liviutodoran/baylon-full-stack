package main

import( 
	
	//"fmt"
	"log"	
	//"net/http"
	"context"
    Country "github.com/babylon-stack/country"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    )
    
   
func GetClient() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    return client
}

func ReturnAllCountries(client *mongo.Client, filter bson.M) []*Country {
    var countries []*Country
    collection := client.Database("babylon").Collection("countries")
    cur, err := collection.Find(context.TODO(), filter)
    if err != nil {
        log.Fatal("Error on Finding all the documents", err)
    }
    for cur.Next(context.TODO()) {
        var country Country
        err = cur.Decode(&country)
        if err != nil {
            log.Fatal("Error on Decoding the document", err)
        }
        countries = append(countries, &country)
    }
    return countries
}

func main() {
    Country := Country.New() 	
	c := GetClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
        log.Fatal("Couldn't connect to the database", err)
    } else {
        log.Println("Connected!")
    }
	
	countries := ReturnAllCountries(c, bson.M{})
	for _, hero := range countries {
		log.Println(countries.Name, countries.Capital)
	}
}