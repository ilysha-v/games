package backend

import (
	"github.com/ilysha-v/games/backend/configuration"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getGames() []GameInfo {
	session, collection := openConnection()
	defer session.Close()

	var results []GameInfo
	err := collection.Find(nil).Sort("name").All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func getGamesWithPaging(pageNumber int, takeCount int) []GameInfo {
	session, collection := openConnection()
	defer session.Close()

	var results []GameInfo
	err := collection.Find(nil).Sort("name").Limit(takeCount).Skip(pageNumber * takeCount).All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func getGameFullInfo(gameId string) []FullGameInfo {
	session, collection := openConnection()
	defer session.Close()

	var results []FullGameInfo
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(gameId)}).All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func openConnection() (*mgo.Session, *mgo.Collection) {
	databaseHost := configuration.GetDatabaseHost()
	session, err := mgo.Dial(databaseHost)
	if err != nil {
		panic(err)
	}

	collection := session.DB("games").C("games")
	return session, collection
}
