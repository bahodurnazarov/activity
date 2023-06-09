package create

import (
	"fmt"
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
	"time"
)

func InsertCategory(act types.Activities) {
	conn := db.Conn()
	currentTime := time.Now()
	time_stamp := currentTime.Format("2006.01.02 15:04:05")
	category := act.Type
	getID := fmt.Sprintf(`select c.id from categories as c 
            		where c.category='%s'`, category)
	rows, err := conn.Query(getID)
	if err != nil {
		lg.Errl.Println("Panic Query getID")
	}
	var categoryID int
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&categoryID)
		if err != nil {
			lg.Errl.Println("Panic Scan rows")
		}
		//fmt.Println("ID ID ID", categoryID)
	}
	query := `insert into category_group
			(activity, category, participants, price, link, ac_key, accessibility, created_at, category_id)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	resp, err := conn.Query(
		query,
		act.Activity,
		category,
		act.Participants,
		act.Price,
		act.Link,
		act.Key,
		act.Accessibility,
		time_stamp,
		categoryID)
	if err != nil {
		lg.Errl.Println("error INSERT CATEGORY: ", err)
	}
	lg.Server.Printf("%+v\n", resp)
}
