package create

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

func CreateActCat(act types.Activities) {
	//activity, activityCategory, participants, price, link, activityKey, accessibility := api.GetFromAPI()
	//fmt.Println(activity, activityCategory, participants, price, link, activityKey, accessibility)
	conn := db.Conn()
	query := `create table if not exists categories(
	id serial PRIMARY KEY,
	category varchar(50),
	created_at TIMESTAMP
 	)`

	_, err := conn.Exec(query)
	if err != nil {
		lg.Errl.Fatalf("error CREATE ActCat TABLE: %v", err)
	}
	//Insert(act.Type)
}
