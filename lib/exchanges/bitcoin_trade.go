package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency/real"
	"fmt"
)

const (
	BitcoinTradeName xrate.ExchangeName = "bitcoinTrade"
)

type BTTicker struct {
	High float32 `json:"high"`
	Low  float32 `json:"low"`
	Vol  float32 `json:"volume"`
	Last float32 `json:"last"`
	Buy  float32 `json:"buy"`
	Sell float32 `json:"sell"`
}

type bitcoinTrade struct {
	xrate.ExchangeParams
}

func BitcoinTrade() xrate.Exchange {
	return &bitcoinTrade{
		xrate.ExchangeParams{
			Name: BitcoinTradeName,
			CryptoCurrencies: []currency.Currency{
				bitcoin.Currency(),
			},
			FiatCurrencies: []currency.Currency{
				real.Currency(),
			},
			BaseApiURL: "https://api.bitcointrade.com.br/v1/public",
		},
	}
}

func (m *bitcoinTrade) getTickerURL(c currency.Currency) (string, error) {

	if !m.supportsCurrency(c) {
		return "", fmt.Errorf("exchange 'Bitcoin Trade' does not support %s", c.Name)
	}

	return m.BaseApiURL + "/" + string(c.Acronym) + "/ticker", nil
}

func (m *bitcoinTrade) supportsCurrency(c currency.Currency) bool {

	return xrate.SliceContainsCurrency(m.CryptoCurrencies, c)
}

func (m *bitcoinTrade) getName() xrate.ExchangeName {

	return m.Name
}
