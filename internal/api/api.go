package api

import (
	"encoding/json"
	"github.com/bahodurnazarov/activity/pkg/types"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
	"io/ioutil"
	"net/http"
)

func GetFromAPI() *types.Activities {
	url := "https://www.boredapi.com/api/activity"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		lg.Errl.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		lg.Errl.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		lg.Errl.Println(err)

	}

	activities := types.Activities{}
	jsonErr := json.Unmarshal(body, &activities)
	if jsonErr != nil {
		lg.Errl.Fatal(jsonErr)
	}
	//fmt.Println("activities.Activity", activities.Activity)
	//fmt.Println("activities.Type", activities.Type)
	//fmt.Println("activities.Participants", activities.Participants)
	//fmt.Println("activities.Price", activities.Price)
	//fmt.Println("activities.Link,", activities.Link)
	//fmt.Println("activities.Key", activities.Key)
	//fmt.Println("activities.Accessibility", activities.Accessibility)

	return &activities

}
