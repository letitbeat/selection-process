package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/letitbeat/selection-process/pkg/scoring"
)

var (
	scoringHandler scoring.Handler
)

func init() {
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("error: a DB URI (ie: 'mongodb://mongo:27017') should be set.")
	}
	scoringHandler = scoring.Handler{DBURI: dbURI}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/scoring/{taskID}", scoringHandler.Score).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
