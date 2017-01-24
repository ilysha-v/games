package backend

import "github.com/ilysha-v/games/backend/auth"

// GameInfo is DTO model for game object
type GameInfo struct {
	Name      string
	Thumbnail string
	Platform  string
}

type ShortUserInfo struct {
	Id    int
	Name  string
	Email string
}

func MakeShortUser(fullUser *auth.User) *ShortUserInfo {
	shortUser := ShortUserInfo{
		Id:    fullUser.ID,
		Name:  fullUser.Name,
		Email: fullUser.Email,
	}

	return &shortUser
}
