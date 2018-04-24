package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"fmt"
	"encoding/json"
)

const (
	BitcoinTradeName xrate.ExchangeName = "BitcoinTrade"
)

type BTTicker struct {
	High float32 `json:"high"`
	Low  float32 `json:"low"`
	Vol  float32 `json:"volume"`
	Last float32 `json:"last"`
	Buy  float32 `json:"buy"`
	Sell float32 `json:"sell"`
}

type BitcoinTradeResponseBody struct {
	Data BTTicker `json:"data"`
}

type bitcoinTrade struct {
	xrate.ExchangeParams
}

func BitcoinTrade() xrate.Exchange {
	return &bitcoinTrade{
		xrate.ExchangeParams{
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

	return xrate.SliceContainsCurrency(bt.FiatCurrencies, f)
}

func (bt *bitcoinTrade) SupportsCryptoCurrency(c currency.Currency) bool {

	return xrate.SliceContainsCurrency(bt.CryptoCurrencies, c)
}

func (bt *bitcoinTrade) GetName() xrate.ExchangeName {

	return bt.Name
}

func (bt *bitcoinTrade) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*xrate.CrawlerResponse, error) {

	var res BitcoinTradeResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &xrate.CrawlerResponse{
		Exchange:           bt.ExchangeParams,
		CryptoCurrency:     cc,
		FiatCurrency:       fc,
		Last:               res.Data.Last,
		High24h:            res.Data.High,
		Low24h:             res.Data.Low,
		Volume24h:          res.Data.Vol,
		VolumeFiat24h:      xrate.UnsupportedField,
		MostRecentBidOrder: res.Data.Buy,
		MostRecentAskOrder: res.Data.Sell,
	}, nil
}
