package handler

import (
	"github.com/bahodurnazarov/activity/internal/api"
	"github.com/bahodurnazarov/activity/internal/create"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	act := api.GetFromAPI()
	create.Insert(*act)
	create.InsertCategory(*act)
	create.GetAllRecord()
}
