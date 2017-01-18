package main

import (
	"logging"
)

func main() {
	logging.InitLogger()
	logging.Log.Infof("Service started")
}
