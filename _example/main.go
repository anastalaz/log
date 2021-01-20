package main

import (
	"errors"
	"github.com/anastalaz/log"
)

func main() {
	log.Debugger = true

	log.Info("App initialized")
	log.Debug("Some", "info", "for you")

	err := errors.New("Something bad happened")
	if err != nil {
		log.Error(err)
	}
}
