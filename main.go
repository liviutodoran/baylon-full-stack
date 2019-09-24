package main

import( 
	
	"fmt"
	"log"	    
    "net/http"
    "github.com/gorilla/mux"
    "babylon-stack/api/handlers"    
    )


func main() {
    
    router :=mux.NewRouter()
    router.HandleFunc("/countries",handlers.GetAllCountriesEndPoint).Methods("GET")
    fmt.Println("Starting server on port 8020...")
	log.Fatal(http.ListenAndServe(":8020", router))
}