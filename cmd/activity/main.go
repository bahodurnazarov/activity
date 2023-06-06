package main

import (
	"fmt"
	"guthub.com/bahodurnazarov/activity/internal/api"
	lg "guthub.com/bahodurnazarov/activity/pkg/utils"
)

func main() {
	a, b, c, d, e, i, f := api.GetFromAPI()
	fmt.Println(a, b, c, d, e, i, f)
	lg.Server.Println("hello")
	lg.Errl.Println("df")
}
