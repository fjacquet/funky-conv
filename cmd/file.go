package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// main
func checkFiles() {

	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(Cfg.rootFS.scanFolder); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
} // func validateFile(file *multipart.FileHeader) error {
// 	// Check the file type
// 	if file.Header.Get("Content-Type") != "video/mp4" {
// 		return errors.New("Invalid file type")
// 	}

// 	// Check the file size
// 	if file.Size > maxFileSize {
// 		return errors.New("File size exceeds the limit")
// 	}

// 	return nil
// }

// func saveFile() {
// 	tempPath := "/tmp/" + file.Filename
// 	if err := c.SaveUploadedFile(file, tempPath); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
// 		return
// 	}
// }
