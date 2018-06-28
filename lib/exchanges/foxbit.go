package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"time"
)

const (
	FoxbitName ExchangeName = "Foxbit"
)

type FoxbitTickerResponseBody struct {
	High    float64 `json:"high"`
	Vol     float64 `json:"vol"`
	Buy     float64 `json:"buy"`
	Last    float64 `json:"last"`
	Low     float64 `json:"low"`
	Sell    float64 `json:"sell"`
	Vol_brl float64 `json:"vol_brl"`
}

type foxbit struct {
	ExchangeParams
}

func Foxbit() Exchange {
	return &foxbit{
		ExchangeParams{
			Name: FoxbitName,
			CryptoCurrencies: []currency.Currency{
				currency.Bitcoin(),
			},
			FiatCurrencies: []currency.Currency{
				currency.Real(),
			},
			BaseApiURL: "https://api.blinktrade.com/api/v1/BRL", //FIXME remove BRL
		},
	}
}

func (f *foxbit) GetTickerURL(c currency.Currency, _ currency.Currency) (string, error) {

	if !f.SupportsCryptoCurrency(c) {
		return "", fmt.Errorf("exchange 'Foxbit' does not support %s", c.Name)
	}

	return f.BaseApiURL + "/ticker?crypto_currency=" + string(c.Acronym), nil
}

func (f *foxbit) SupportsFiatCurrency(fc currency.Currency) bool {

	return SliceContainsCurrency(f.FiatCurrencies, fc)
}

func (f *foxbit) SupportsCryptoCurrency(c currency.Currency) bool {

	return SliceContainsCurrency(f.CryptoCurrencies, c)
}

func (f *foxbit) GetName() ExchangeName {

	return f.Name
}

func (f *foxbit) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*CrawlerResponse, error) {

	var res FoxbitTickerResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &CrawlerResponse{
		Exchange:           f.ExchangeParams,
		CryptoCurrency:     cc,
		FiatCurrency:       fc,
		Last:               res.Last,
		High24h:            res.High,
		Low24h:             res.Low,
		Volume24h:          res.Vol,
		VolumeFiat24h:      res.Vol_brl,
		MostRecentBidOrder: res.Buy,
		MostRecentAskOrder: res.Sell,
		CreatedAt:          time.Now(),
	}, nil
}
