package main

import (
	"babylon-stack/api/handlers"
	"babylon-stack/api/models"
	"babylon-stack/utilstools"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	utilstools.GetDataXLX()
	router := mux.NewRouter()
	var countries models.Country
	var wage models.Minimumwage
	var languages models.Languages

	router.HandleFunc("/countries", handlers.GetAll(countries)).Methods("GET")
	router.HandleFunc("/wage", handlers.GetAll(wage)).Methods("GET")
	router.HandleFunc("/languages", handlers.GetAll(languages)).Methods("GET")

	router.HandleFunc("/country/{id}", handlers.GetItem(countries)).Methods("GET")
	router.HandleFunc("/wage/{id}", handlers.GetItem(wage)).Methods("GET")
	router.HandleFunc("/languages/{id}", handlers.GetItem(languages)).Methods("GET")

	router.HandleFunc("/country/{id}", handlers.UpdateItem(countries)).Methods("PUT")
	router.HandleFunc("/wage/{id}", handlers.UpdateItem(wage)).Methods("PUT")
	router.HandleFunc("/languages/{id}", handlers.UpdateItem(languages)).Methods("PUT")

	router.HandleFunc("/country", handlers.AddItem(countries)).Methods("POST")
	router.HandleFunc("/wage", handlers.AddItem(wage)).Methods("POST")
	router.HandleFunc("/languages", handlers.AddItem(languages)).Methods("POST")

	router.HandleFunc("/country", handlers.DeleteItem(countries)).Methods("DELETE")
	router.HandleFunc("/wage", handlers.DeleteItem(wage)).Methods("DELETE")
	router.HandleFunc("/languages", handlers.DeleteItem(languages)).Methods("DELETE")

	router.HandleFunc("/languages", handlers.DeleteItem(languages)).Methods("DELETE")

	router.HandleFunc("/currency/{item1}/{item2}", handlers.GetCurrency).Methods("GET")

	fmt.Println("Starting server on port 8020...")
	log.Fatal(http.ListenAndServe(":8020", router))

}
