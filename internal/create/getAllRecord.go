package create

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

func GetAllRecord() *types.Response {
	conn := db.Conn()
	query := `SELECT e.activity, x.category, e.participants, e.price, e.link, e.ac_key, e.accessibility
				FROM categories AS x, category_group AS e WHERE x.id = e.category_id`
	rows, err := conn.Query(query)
	if err != nil {
		lg.Errl.Println("Panic Query query")
	}
	var resp types.Response
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&resp.Activity, &resp.Category, &resp.Participants, &resp.Price, &resp.Link, &resp.Key, &resp.Accessibility)
		if err != nil {
			lg.Errl.Println("Panic Scan rows GetAllRecord")
		}
	}
	return &resp
}
