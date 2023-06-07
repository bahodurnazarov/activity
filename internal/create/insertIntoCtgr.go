package handler

import (
	"fmt"
	"github.com/bahodurnazarov/activity/pkg/db"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
	"time"
)

func Insert(activityCategory string) {
	fmt.Println("sd", activityCategory)
	conn := db.Conn()
	a := time.Now()
	query := `insert into act_category
			(ac_type, created_at)
			values ($1, $2)`
	resp, err := conn.Query(
		query,
		activityCategory,
		a)
	if err != nil {
		lg.Errl.Println("error INSERT: ", err)
	}
	fmt.Printf("%+v\n", resp)
}
