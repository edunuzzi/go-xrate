package util

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func BaseGet(url string, res interface{}) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer response.Body.Close()

	body, bodyErr := ioutil.ReadAll(response.Body)

	if bodyErr != nil {
		fmt.Println(bodyErr)
		panic(bodyErr)
	}

	jsonErr := json.Unmarshal(body, &res)

	if jsonErr != nil {
		fmt.Println(jsonErr)
		panic(jsonErr)
	}
}