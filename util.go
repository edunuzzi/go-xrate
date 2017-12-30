package xrate

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type CryptoCurrencyAcronym string
const (
	BTC CryptoCurrencyAcronym = "BTC"
	ETH CryptoCurrencyAcronym = "ETH"
	BCH CryptoCurrencyAcronym = "BCH"
	LTC CryptoCurrencyAcronym = "LTC"
)

type FiatCurrencyAcronym string
const (
	BRL FiatCurrencyAcronym = "BRL"
	USD FiatCurrencyAcronym = "USD"
)

type CryptoCurrencyTicker struct {
	Acronym             CryptoCurrencyAcronym
	FiatCurrencyAcronym FiatCurrencyAcronym
	Last                float32
	High24h             float32
	Low24h              float32
	Volume24h           float32
	VolumeFiat24h       float32
	RecentBuyOrder      float32
	RecentSellOrder     float32
}

type ExchangeCrawler struct {
	BaseUrl string
}

func StringToFloat32(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return float32(value)
}

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
