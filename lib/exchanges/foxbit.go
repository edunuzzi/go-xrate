package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency/real"
	"fmt"
	"encoding/json"
)

const (
	FoxbitName xrate.ExchangeName = "foxbit"
)

type FoxbitTickerResponseBody struct {
	High    float32 `json:"high"`
	Vol     float32 `json:"vol"`
	Buy     float32 `json:"buy"`
	Last    float32 `json:"last"`
	Low     float32 `json:"low"`
	Sell    float32 `json:"sell"`
	Vol_brl float32 `json:"vol_brl"`
}

type foxbit struct {
	xrate.ExchangeParams
}

func Foxbit() xrate.Exchange {
	return &foxbit{
		xrate.ExchangeParams{
			Name: FoxbitName,
			CryptoCurrencies: []currency.Currency{
				bitcoin.Currency(),
			},
			FiatCurrencies: []currency.Currency{
				real.Currency(),
			},
			BaseApiURL: "https://api.blinktrade.com/api/v1/BRL",
		},
	}
}

func (f *foxbit) GetTickerURL(c currency.Currency) (string, error) {

	if !f.SupportsCryptoCurrency(c) {
		return "", fmt.Errorf("exchange 'Foxbit' does not support %s", c.Name)
	}

	return f.BaseApiURL + "/ticker?crypto_currency=" + string(c.Acronym), nil
}

func (f *foxbit) SupportsFiatCurrency(fc currency.Currency) bool {

	return xrate.SliceContainsCurrency(f.FiatCurrencies, fc)
}

func (f *foxbit) SupportsCryptoCurrency(c currency.Currency) bool {

	return xrate.SliceContainsCurrency(f.CryptoCurrencies, c)
}

func (f *foxbit) GetName() xrate.ExchangeName {

	return f.Name
}

func (f *foxbit) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*xrate.CrawlerResponse, error) {

	var res FoxbitTickerResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &xrate.CrawlerResponse{
		Exchange:            f.ExchangeParams,
		CryptoCurrency:      cc,
		FiatCurrency:        fc,
		Last:                res.Last,
		High24h:             res.High,
		Low24h:              res.Low,
		Volume24h:           res.Vol,
		VolumeFiat24h:       res.Vol_brl,
		MostRecentBuyOrder:  res.Buy,
		MostRecentSellOrder: res.Sell,
	}, nil
}
