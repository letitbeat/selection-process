package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/letitbeat/selection-process/pkg/task"
	"github.com/letitbeat/selection-process/pkg/user"
)

var (
	usersHandler user.Handler
	tasksHandler task.Handler
)

func init() {
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("error: a DB URI (ie: 'mongodb://mongo:27017') should be set.")
	}

	scoringServer := os.Getenv("SCORING_SERVER")
	if scoringServer == "" {
		log.Fatal("error: a scoring server (ie: 'http://scoring:8080') should be set.")
	}

	usersHandler = user.Handler{DBURI: dbURI}
	tasksHandler = task.Handler{DBURI: dbURI, ScoringServer: scoringServer}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", usersHandler.GetByID).Methods("GET")

	router.HandleFunc("/tasks/{id}", tasksHandler.GetByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
