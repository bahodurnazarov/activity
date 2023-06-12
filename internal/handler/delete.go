package handler

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func DeleteRecord(c *gin.Context) {
	conn := db.Conn()
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Error in convert id", err.Error())
		return
	}

	conn.Exec("DELETE FROM category_group WHERE id=$1", ID)
}
