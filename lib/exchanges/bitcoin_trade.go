package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"time"
)

const (
	BitcoinTradeName ExchangeName = "BitcoinTrade"
)

type BTTicker struct {
	High float64 `json:"high"`
	Low  float64 `json:"low"`
	Vol  float64 `json:"volume"`
	Last float64 `json:"last"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
}

type BitcoinTradeResponseBody struct {
	Data BTTicker `json:"data"`
}

type bitcoinTrade struct {
	ExchangeParams
}

func BitcoinTrade() Exchange {
	return &bitcoinTrade{
		ExchangeParams{
			Name: BitcoinTradeName,
			CryptoCurrencies: []currency.Currency{
				currency.Bitcoin(),
			},
			FiatCurrencies: []currency.Currency{
				currency.Real(),
			},
			BaseApiURL: "https://api.bitcointrade.com.br/v1/public",
		},
	}
}

func (bt *bitcoinTrade) GetTickerURL(c currency.Currency, _ currency.Currency) (string, error) {

	if !bt.SupportsCryptoCurrency(c) {
		return "", fmt.Errorf("exchange 'Bitcoin Trade' does not support %s", c.Name)
	}

	return bt.BaseApiURL + "/" + string(c.Acronym) + "/ticker", nil
}

func (bt *bitcoinTrade) SupportsFiatCurrency(f currency.Currency) bool {

	return SliceContainsCurrency(bt.FiatCurrencies, f)
}

func (bt *bitcoinTrade) SupportsCryptoCurrency(c currency.Currency) bool {

	return SliceContainsCurrency(bt.CryptoCurrencies, c)
}

func (bt *bitcoinTrade) GetName() ExchangeName {

	return bt.Name
}

func (bt *bitcoinTrade) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*CrawlerResponse, error) {

	var res BitcoinTradeResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &CrawlerResponse{
		Exchange:           bt.ExchangeParams,
		CryptoCurrency:     cc,
		FiatCurrency:       fc,
		Last:               res.Data.Last,
		High24h:            res.Data.High,
		Low24h:             res.Data.Low,
		Volume24h:          res.Data.Vol,
		VolumeFiat24h:      UnsupportedField,
		MostRecentBidOrder: res.Data.Buy,
		MostRecentAskOrder: res.Data.Sell,
		CreatedAt:          time.Now(),
	}, nil
}
