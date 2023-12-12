package main

import (
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

// test if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
