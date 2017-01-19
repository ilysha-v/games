package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router is application router
var Router *mux.Router

// Init is initializing method for all controllers in a serice
func Init() {
	Router = mux.NewRouter()
	Router.StrictSlash(true)
	Router.HandleFunc("/", Index)

	http.ListenAndServe(":8080", Router)
}

// Index in main page handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, web service!")
}
