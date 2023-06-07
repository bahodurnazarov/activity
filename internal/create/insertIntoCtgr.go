package create

import (
	"fmt"
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
	"time"
)

func Insert(act types.Activities) {
	conn := db.Conn()
	currentTime := time.Now()
	time_stamp := currentTime.Format("2006.01.02 15:04:05")
	query := `insert into categories
			(category, created_at)
			values ($1, $2)`
	resp, err := conn.Query(
		query,
		act.Type,
		time_stamp)
	if err != nil {
		lg.Errl.Println("error INSERT: ", err)
	}
	fmt.Printf("%+v\n", resp)
}
