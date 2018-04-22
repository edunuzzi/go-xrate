package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency/real"
	"fmt"
)

const (
	MercadoBitcoinName xrate.ExchangeName = "Mercado Bitcoin"
)

type mercadoBitcoin struct {
	xrate.ExchangeParams
}

func MercadoBitcoin() xrate.Exchange {
	return &mercadoBitcoin{
		xrate.ExchangeParams{
			Name: MercadoBitcoinName,
			CryptoCurrencies: []currency.Currency{
				bitcoin.Currency(),
			},
			FiatCurrencies: []currency.Currency{
				real.Currency(),
			},
			BaseApiURL: "https://www.mercadobitcoin.net/api",
		},
	}
}

func (m *mercadoBitcoin) GetTickerURL(c currency.Currency) (string, error) {

	if !m.supportsCurrency(c) {
		return "", fmt.Errorf("exchange 'Mercado Bitcoin' does not support %s", c.Name)
	}

	return m.BaseApiURL+"/"+string(c.Acronym)+"/ticker", nil
}

func (m *mercadoBitcoin) supportsCurrency(c currency.Currency) bool {

	return xrate.SliceContainsCurrency(m.CryptoCurrencies, c)
}

func (m *mercadoBitcoin) getName() xrate.ExchangeName {
	return m.Name
}