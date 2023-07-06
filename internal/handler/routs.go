package handler

import (
	"github.com/bahodurnazarov/activity/pkg/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/bahodurnazarov/activity/docs"
)

func Run() {
	r := gin.Default()
	r.Use(cors.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", Home)
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.POST("/editText/:id", EditText)
	r.POST("/join/:id", Join)
	r.DELETE("/delete/:id", DeleteRecord)
	r.NoRoute(func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"code": "works!", "message": "not route!"})
	})
	r.Run(":8888")
}
