package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const takeCount int = 25

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
	variables := r.URL.Query()

	withPaging := false
	page := variables["page"]
	if len(page) == 1 {
		withPaging = true
	}
	var games []GameInfo
	if !withPaging {
		games = getGames()
	} else {
		pageNumber, err := strconv.Atoi(page[0])
		if err != nil {
			panic(err)
		}
		games = getGamesWithPaging(pageNumber, takeCount)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(games); err != nil {
		// todo this should response 500 and log full error
		panic(err)
	}
}
