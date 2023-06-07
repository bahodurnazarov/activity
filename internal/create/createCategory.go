package create

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

func CreateCategory(act types.Activities) {
	conn := db.Conn()
	//category := act.Type
	query := `create table if not exists category_group(
    id serial primary key,
    activity varchar(250),
    category varchar(50),
	participants numeric,
	price numeric,
	link varchar(250),
    ac_key varchar(50),
	accessibility numeric,
	created_at timestamp,
	category_id integer references categories(id) 
	)`

	_, err := conn.Exec(query)
	if err != nil {
		lg.Errl.Fatalf("error CREATE Cat TABLE: %v", err)
	}

}
