package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency/real"
	"fmt"
)

const (
	FoxbitName xrate.ExchangeName = "foxbit"
)

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

func (m *foxbit) GetTickerURL(c currency.Currency) (string, error) {

	if !m.supportsCurrency(c) {
		return "", fmt.Errorf("exchange 'Foxbit' does not support %s", c.Name)
	}

	return m.BaseApiURL+"/ticker?crypto_currency="+string(c.Acronym), nil
}

func (m *foxbit) supportsCurrency(c currency.Currency) bool {

	return xrate.SliceContainsCurrency(m.CryptoCurrencies, c)
}

func (m *foxbit) getName() xrate.ExchangeName {
	return m.Name
}