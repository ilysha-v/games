package backend

import "gopkg.in/mgo.v2/bson"

// GameInfo is DTO model for game object
type GameInfo struct {
	Name      string        `json:"name"`
	Thumbnail string        `json:"thumbnail"`
	Platform  string        `json:"platform"`
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

type FullGameInfo struct {
	Name      string   `json:"name"`
	Thumbnail string   `json:"thumbnail"`
	Platform  string   `json:"platform"`
	Genre     []string `json:"genre"`
	Rating    string   `json:"rating"`
	Summary   string   `json:"summary"`
}
