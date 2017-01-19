package main

import (
	"logging"
	"routing"
)

func main() {
	logging.InitLogger()
	routing.Init()
	logging.Log.Infof("Service started")
}
