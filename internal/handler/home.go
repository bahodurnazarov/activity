package handler

import (
	"github.com/bahodurnazarov/activity/internal/api"
	"github.com/bahodurnazarov/activity/internal/create"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	act := api.GetFromAPI()
	create.Insert(*act)
	create.InsertCategory(*act)
	resp := create.GetAllRecord()
	c.JSON(http.StatusOK, gin.H{
		"activity": resp.Activity,
	})
}
