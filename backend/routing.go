package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilysha-v/games/backend/auth"
)

const takeCount int = 25

// Router is application router
var Router *mux.Router

// Init is initializing method for all controllers in a serice
func Init() {
	InitLogger()

	auth.SetupAuth()
	Log.Infof("Auth system initialized")

	Router = mux.NewRouter()
	Router.StrictSlash(true)
	Router.HandleFunc("/api/test", indexHandler)
	Router.HandleFunc("/api/games", gamesHandler)
	Router.HandleFunc("/api/gamedetail", gameDetailHandler)
	Router.HandleFunc("/api/whoami", whoAmI)

	authRouter := auth.Ab.NewRouter()
	Router.PathPrefix("/api/auth").Handler(authRouter)

	Log.Infof("Service started")

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
		Log.Infof("Error")
	}

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

func gameDetailHandler(w http.ResponseWriter, r *http.Request) {
	variables := r.URL.Query()

	gameId := variables["gameid"]
	var game []FullGameInfo
	game = getGameFullInfo(gameId[0])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(game); err != nil {
		// todo this should response 500 and log full error
		panic(err)
	}
}
