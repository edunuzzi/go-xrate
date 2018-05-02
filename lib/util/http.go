package util

import (
	"io/ioutil"
	"net/http"
	"time"
)

func BaseGet(url string, timeout time.Duration) (body []byte, err error) {
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)

	return
}
