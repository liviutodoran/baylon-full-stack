package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func handlerLiviu(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Fom the Other Liviu +Emilio .."
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func main_old() {


	/***** MIERDA del otro main
	/*cnn, err := sql.Open("mysql", "quitSmokeDev:Proyecto098.@tcp(82.223.0.98:3306)/quite_smoke_dev")
	if err != nil {
			log.Fatal(err)
	}

	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
			log.Fatal(err)
	}

	fmt.Println(id, name)*/
	//TIME TO 20
	//ch := time.After(20 * time.Second)
	//defer (func() { fmt.Println("waiting"); <-ch; fmt.Println("waited") })()
	
	fmt.Println("Starting web server on port 8080")
	
	*/
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handlerLiviu)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Configure Logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}