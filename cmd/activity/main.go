package main

import (
	"github.com/bahodurnazarov/activity/pkg/db"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

func main() {
	//a, b, c, d, e, i, f := api.GetFromAPI()
	a := db.Conn()
	_, err := a.Exec("INSERT into test VALUES ($1)", 2)
	if err != nil {
		lg.Errl.Fatalf("111An error occured while executing query: %v", err)
	}

}
