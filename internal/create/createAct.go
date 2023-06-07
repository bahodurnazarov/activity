package handler

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

func CreateAct() {
	//activity, activityCategory, participants, price, link, activityKey, accessibility := api.GetFromAPI()
	//fmt.Println(activity, activityCategory, participants, price, link, activityKey, accessibility)
	conn := db.Conn()

	query := `create table if not exists act_category(
	id serial primary key,
	ac_type varchar(50),
	created_at timestamp
	)`

	_, err := conn.Exec(query)
	if err != nil {
		lg.Errl.Fatalf("error CREATE TABLE: %v", err)
	}
	//fmt.Println("dfdsf", activityCategory)

}
