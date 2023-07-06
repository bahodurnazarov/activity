package handler

import (
	"fmt"
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Join(c *gin.Context) {
	conn := db.Conn()
	ID, err := strconv.Atoi(c.Param("id"))
	log.Println("rIDe ==>", ID)
	if err != nil {
		log.Println("Error in convert id", err.Error())
		return
	}
	query := fmt.Sprintf(`SELECT participants FROM category_group WHERE id='%d'`, ID)
	rows, err := conn.Query(query)
	if err != nil {
		log.Println("Panic Query getID")
	}
	var categoryID int
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&categoryID)
		if err != nil {
			log.Println("Panic Scan rows")
		}
		//fmt.Println("ID ID ID", categoryID)
	}
	result := categoryID + 1

	update := fmt.Sprintf(`update category_group set participants = $1 where  id = $2`)
	rows1, err := conn.Query(update, result, ID)
	if err != nil {
		log.Println("Panic Query getID")
	}
	defer rows1.Close()
}
