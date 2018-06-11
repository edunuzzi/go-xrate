package exchanges

import (
	"reflect"
	"github.com/Swipecoin/go-xrate/lib/currency"
)

type ExchangeName string

type ExchangeParams struct {
	Name             ExchangeName
	CryptoCurrencies []currency.Currency
	FiatCurrencies   []currency.Currency
	BaseApiURL       string
}

type Exchange interface {
	GetTickerURL(cryptoCurrency currency.Currency, fiatCurrency currency.Currency) (string, error)
	SupportsFiatCurrency(f currency.Currency) bool
	SupportsCryptoCurrency(c currency.Currency) bool
	GetName() ExchangeName
	ConvertToResponse(cryptoCurrency currency.Currency, fiatCurrency currency.Currency, body []byte) (*CrawlerResponse, error)
}

func SliceContainsCurrency(currencies []currency.Currency, c currency.Currency) bool {

	isSupported := false

	for _, curr := range currencies {

		if reflect.DeepEqual(c, curr) {
			isSupported = true
		}
	}

	return isSupported
}
