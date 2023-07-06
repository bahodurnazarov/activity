package handler

import (
	"fmt"
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type editText struct {
	Text string `json:"text"`
}

func EditText(c *gin.Context) {
	conn := db.Conn()
	ID, _ := strconv.Atoi(c.Param("id"))
	log.Println("rIDe ==>", ID)
	var text editText
	err := c.ShouldBindJSON(&text)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode": 400,
			"error":     "Invalid request body",
		})
		return
	}
	query := fmt.Sprintf(`update category_group set activity = $1 where  id = $2`)
	rows1, err := conn.Query(query, text.Text, ID)
	if err != nil {
		log.Println("Panic Query getID")
	}
	defer rows1.Close()
}
