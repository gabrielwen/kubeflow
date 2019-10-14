package main

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

var (
	VERSION = "0.0.1"
)

func init() {
	// Add filename as one of the fields of the structured log message.
	filenameHook := filename.NewHook()
	filenameHook.Field = "filename"
	log.AddHook(filenameHook)
}

func main() {
	log.Info("Hello world.")
}
