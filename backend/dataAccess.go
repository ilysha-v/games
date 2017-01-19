package backend

import "gopkg.in/mgo.v2"

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

func openConnection() (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	collection := session.DB("games").C("games")
	return session, collection
}
