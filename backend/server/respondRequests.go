package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var cost = 10

type response struct {
	ID      string "json:id"
	Message string "json:text"
}

var defaultResponse = []response{
	{ID: "abhi", Message: "get hoes"},
}

func getDefaultResponse(c *gin.Context) {
	c.JSON(http.StatusOK, defaultResponse)
}

func createNewUser(c *gin.Context, db *pgx.Conn) {
	userName, exists := c.GetPostForm("Username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must input username"})
		return
	}

	containsUsername, err := checkUserNameExists(db, userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "DB Query Error"})
	}

	if containsUsername {
		c.JSON(http.StatusOK, gin.H{"Message": "Username already taken"})
	}

	passWord, exists := c.GetPostForm("Password")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must input password"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passWord), cost)

	addUserToDB(db, userName, hashedPassword)

}
