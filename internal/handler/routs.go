package handler

import (
	"github.com/bahodurnazarov/activity/pkg/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	router.GET("/", Home)
	router.NoRoute(func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"code": "works!", "message": "not route!"})
	})
	router.Run(":8888")
}
