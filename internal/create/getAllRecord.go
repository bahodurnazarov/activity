package create

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

func GetAllRecord() []types.Response {
	conn := db.Conn()
	query := `SELECT e.id, e.activity, x.category, e.participants, e.price, e.link, e.ac_key, e.accessibility
				FROM categories AS x, category_group AS e WHERE x.id = e.category_id`
	rows, err := conn.Query(query)
	if err != nil {
		lg.Errl.Println("Panic Query query")
	}
	defer rows.Close()

	var resp = types.Response{}
	var retResp []types.Response
	for rows.Next() {
		err = rows.Scan(
			&resp.ID,
			&resp.Activity,
			&resp.Category,
			&resp.Participants,
			&resp.Price,
			&resp.Link,
			&resp.Key,
			&resp.Accessibility)
		if err != nil {
			lg.Errl.Println("Panic Scan rows GetAllRecord")
		}
		var newStrct = types.Response{
			ID:            resp.ID,
			Activity:      resp.Activity,
			Category:      resp.Category,
			Participants:  resp.Participants,
			Price:         resp.Price,
			Link:          resp.Link,
			Key:           resp.Key,
			Accessibility: resp.Accessibility,
		}
		retResp = append(retResp, newStrct)
		//fmt.Println("rdDDDDD", retResp)
	}
	//fmt.Println("EEEEE", retResp)
	return retResp
}
