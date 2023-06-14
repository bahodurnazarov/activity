package handler

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// JWTClaims is a struct for storing JWT claims
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	// Parse the JSON request body
	var user types.Login
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Query the database for the user
	conn := db.Conn()
	row := conn.QueryRow("SELECT id, password FROM users WHERE email = $1", user.Email)

	// Check if the user exists
	var userID int
	var hashedPassword string
	err = row.Scan(&userID, &hashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := JWTClaims{
		Username: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(types.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	currentTime := time.Now()
	time_stamp := currentTime.Format("2006.01.02 15:04")
	_, err = conn.Exec("INSERT INTO users_token (token, created_at, user_id) VALUES ($1, $2, $3)", tokenString, time_stamp, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the JWT token to the client
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
