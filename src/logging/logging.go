package logger

import (
	"github.com/juju/loggo"
)

func InitLogger() {
	loggo.ConfigureLoggers("INFO")
}
