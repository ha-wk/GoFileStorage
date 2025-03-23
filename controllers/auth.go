package controllers

import (
	"net/http"
	"wobot-file-storage/models"
	"wobot-file-storage/utils"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]models.User)

func RegisterHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})    //checking for marshal/unmarshal
		return
	}

	if _, exists := users[user.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})   //checking if user alrady in my system
		return
	}

	hash, _ := utils.HashPassword(user.Password)                             //if not we will hash his password and assign storage quota and strore user 
	user.Password = hash
	user.Quota = 10 * 1024 * 1024
	users[user.Username] = user

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {
	var req models.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, exists := users[req.Username]                                            //check if user exists or creds are valid??
	if !exists || utils.ComparePassword(user.Password, req.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.Username)                             //if all fine,create jwt for particular user
	c.JSON(http.StatusOK, gin.H{"token": token})
}
