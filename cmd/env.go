package main

import (
	"io"
	"log"
	"os"
)

// Check if directory exists
func dirExists(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

// Create Directory if needed
func checkCreateDirectory(path string) {
	if !dirExists(path) {
		os.Mkdir(path, 0750)
	}
}

// prepare layout
func prepareFS() {
	checkCreateDirectory(Cfg.rootFS.folder)
	checkCreateDirectory(Cfg.rootFS.scanFolder)
	checkCreateDirectory(Cfg.rootFS.outFolder)
	checkCreateDirectory(Cfg.rootFS.sdFolder)
	checkCreateDirectory(Cfg.rootFS.hdFolder)
	checkCreateDirectory(Cfg.rootFS.fhdFolder)
}

// Generate log name
func prepareLogs() {
	logFile, err := os.OpenFile(Cfg.logFolder.LogName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	// log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
}
