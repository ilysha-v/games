package main

import (
	"github.com/ilysha-v/games/logging"
	"github.com/ilysha-v/games/routing"
)

func main() {
	logging.InitLogger()
	logging.Log.Infof("Service started")
	routing.Init()
}
