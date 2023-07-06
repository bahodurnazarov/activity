package handler

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	// Parse the JSON request body
	var user types.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode": 400,
			"error":     "Invalid request body",
		})
		return
	}

	// Check if the username already exists
	conn := db.Conn()
	var count int
	err1 := conn.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&count)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorCode": 500,
			"error":     err.Error(),
		})
		return
	}
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorCode": 400,
			"error":     "email already exists",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode": 500,
			"error":     "Failed to hash password",
		})
		return
	}

	// Insert the user into the database
	currentTime := time.Now()
	time_stamp := currentTime.Format("2006.01.02 15:04")
	// Insert the user into the database
	_, err = conn.Exec("INSERT INTO users (name, username, email, password, created_at) VALUES ($1, $2, $3, $4, $5)", user.Name, user.Username, user.Email, hashedPassword, time_stamp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorCode": 500,
			"error":     err.Error(),
		})
		return
	}

	//// Create a new session and return the session ID to the client
	//sessionID, err := createSession(user.ID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
	//	return
	//}
	//
	//// Generate a new access token and refresh token and return them to the client
	//accessToken, err := generateAccessToken(user.ID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
	//	return
	//}
	//refreshToken, err := generateRefreshToken(user.ID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"errorCode": 200,
		"message":   "User registered successfully",
		//"session_id":    sessionID,
		//"access_token":  accessToken,
		//"refresh_token": refreshToken,
	})
}
