package main

import (
	"log"
	"os"
	"path"
)

var (
	lwarn, linfo, ldebug, lerror *log.Logger
)

func initLogging() {
	logFilePath := path.Join(notesLocation, "log")
	f, err := os.Create(logFilePath)
	if err != nil {
		log.Fatalf("Failed to create log file %s : %s", logFilePath, err)
	}

	lwarn = log.New(f, "[warn] ", log.Default().Flags())
	linfo = log.New(f, "[info] ", log.Default().Flags())
	ldebug = log.New(f, "[debug] ", log.Default().Flags())
	lerror = log.New(f, "[error] ", log.Default().Flags())

}
