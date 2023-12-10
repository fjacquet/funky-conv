package main

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func validateFile(file *multipart.FileHeader) error {
	// Check the file type
	if file.Header.Get("Content-Type") != "video/mp4" {
		return errors.New("Invalid file type")
	}

	// Check the file size
	if file.Size > maxFileSize {
		return errors.New("File size exceeds the limit")
	}

	return nil
}

func saveFile() {
	tempPath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
}
