package util

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func BaseGet(url string, res interface{}) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return err
	}
}