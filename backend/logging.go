package backend

import (
	"github.com/juju/loggo"
)

// Log - main service logger
var Log loggo.Logger

// InitLogger - loggint system initializations
func InitLogger() {
	loggo.ConfigureLoggers("INFO")
	Log = loggo.GetLogger("")
}
