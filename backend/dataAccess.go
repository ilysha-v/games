package backend

import "gopkg.in/mgo.v2"

func getGames() []Game {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	collection := session.DB("games").C("games")
	defer session.Close()

	var results []Game
	err = collection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}

	return results
}
