package handler

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/gin-gonic/gin"
)

func DeleteRecord(c *gin.Context) {
	conn := db.Conn()
	ID := c.Query("getID")
	conn.Exec("DELETE FROM category_group WHERE id=$1", ID)
}
