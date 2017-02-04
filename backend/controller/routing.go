package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilysha-v/games/backend"
	"github.com/ilysha-v/games/backend/auth"
)

const takeCount int = 25

// Router is application router
var Router *mux.Router

// Init is initializing method for all controllers in a serice
func Init() {
	backend.InitLogger()

	auth.SetupAuth()
	backend.Log.Infof("Auth system initialized")

	Router = mux.NewRouter()
	Router.StrictSlash(true)
	Router.HandleFunc("/api/test", indexHandler)
	Router.HandleFunc("/api/games", gamesHandler)
	Router.HandleFunc("/api/userinfo", userInfoHandler)
	Router.HandleFunc("/api/whoami", whoAmI)
	Router.HandleFunc("/api/errors", sessionErrorHandler)

	authRouter := auth.Ab.NewRouter()
	Router.PathPrefix("/api/auth").Handler(authRouter)

	backend.Log.Infof("Service started")

	http.ListenAndServe(":8080", Router)
}

// Index in main page handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, web service!")
}

func whoAmI(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Ab.CurrentUser(w, r)
	if user != nil && err == nil {
		currentUserName := user.(*auth.User).Email
		fmt.Fprintf(w, "Hey, %s!", currentUserName)
	} else {
		backend.Log.Infof("Error")
	}
}

func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Ab.CurrentUser(w, r)
	if user != nil && err == nil {
		typedUser := user.(*auth.User)
		shortUser := backend.MakeShortUser(typedUser)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(shortUser); err != nil {
			// todo this should response 500 and log full error
			panic(err)
		}
	} else {
		fmt.Fprintf(w, "{ }")
	}
}

func sessionErrorHandler(w http.ResponseWriter, r *http.Request) {
	ctx := auth.Ab.InitContext(w, r)
	value, success := ctx.SessionStorer.Get("flash_error")
	if success {
		fmt.Fprint(w, value)
	}
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	variables := r.URL.Query()

	withPaging := false
	page := variables["page"]
	if len(page) == 1 {
		withPaging = true
	}

	var games []backend.GameInfo
	if !withPaging {
		games = backend.GetGames()
	} else {
		pageNumber, err := strconv.Atoi(page[0])
		if err != nil {
			panic(err)
		}
		games = backend.GetGamesWithPaging(pageNumber, takeCount)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(games); err != nil {
		// todo this should response 500 and log full error
		panic(err)
	}
}
