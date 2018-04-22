package xrate

import (
	"github.com/Swipecoin/go-currency/currency"
)

type ExchangeName string

type ExchangeParams struct {
	Name ExchangeName
	CryptoCurrencies []currency.Currency
	FiatCurrencies []currency.Currency
	BaseApiURL string
}

type Exchange interface {
	getTickerURL(c currency.Currency) (string, error)
	supportsCurrency(c currency.Currency) bool
	getName() ExchangeName
}

func SliceContainsCurrency(currencies []currency.Currency, c currency.Currency) bool {
	isSupported := false

	for _, curr := range currencies {

		if c.Name == curr.Name {
			isSupported = true
		}
	}

	return isSupported
}