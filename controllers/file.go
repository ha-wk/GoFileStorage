package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wobot-file-storage/config"
	"wobot-file-storage/models"

	"github.com/gin-gonic/gin"
)

var files = make(map[string][]models.FileMetadata)

func UploadFileHandler(c *gin.Context) {
	username := c.GetString("username")
	user := users[username]

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
                                                                                 
	if user.Used+file.Size > user.Quota {                                         //check incming file size with size available
		c.JSON(http.StatusForbidden, gin.H{"error": "Storage limit exceeded"})
		return
	}

	userDir := filepath.Join(config.BaseStoragePath, username)                   // Save file to user directory
	os.MkdirAll(userDir, os.ModePerm)
	filePath := filepath.Join(userDir, file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	user.Used += file.Size
	users[username] = user

	metadata := models.FileMetadata{
		Filename:    file.Filename,
		Size:        file.Size,
		UploadedAt:  time.Now(), 
		OriginalName: file.Filename,
	}
	files[username] = append(files[username], metadata)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func GetStorageHandler(c *gin.Context) {
	username := c.GetString("username")
	user := users[username]

	c.JSON(http.StatusOK, gin.H{                 //calculation logic
		"total":    user.Quota,
		"remaining": user.Quota - user.Used,
	})
}

func ListFilesHandler(c *gin.Context) {
	username := c.GetString("username")              //fetch by usrrname
	c.JSON(http.StatusOK, files[username])
}
