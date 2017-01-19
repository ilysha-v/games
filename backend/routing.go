package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router is application router
var Router *mux.Router

// Init is initializing method for all controllers in a serice
func Init() {
	InitLogger()
	Router = mux.NewRouter()
	Router.StrictSlash(true)
	Router.HandleFunc("/", indexHandler)
	Router.HandleFunc("/games", gamesHandler)

	Log.Infof("Service started")

	http.ListenAndServe(":8080", Router)
}

// Index in main page handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, web service!")
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	games := getGames()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(games); err != nil {
		// todo this should response 500 and log full error
		panic(err)
	}
}
