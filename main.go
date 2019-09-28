package main

import (
	"babylon-stack/api/handlers"
	"babylon-stack/utilstools"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	utilstools.GetDataXLX()
	router := mux.NewRouter()
	router.HandleFunc("/countries", handlers.GetAllCountriesEndPoint).Methods("GET")
	router.HandleFunc("/country/{id}", handlers.GetCountryEndPoint).Methods("GET")
	router.HandleFunc("/country/{id}", handlers.UpdateCountryEndpoint).Methods("PUT")
	router.HandleFunc("/wage", handlers.GetAllMinWageEndPoint).Methods("GET")
	router.HandleFunc("/wage", handlers.DeleteWageEndpoint).Methods("DELETE")
	router.HandleFunc("/wage/{id}", handlers.UpdateWageEndpoint).Methods("PUT")
	fmt.Println("Starting server on port 8020...")
	log.Fatal(http.ListenAndServe(":8020", router))

}
