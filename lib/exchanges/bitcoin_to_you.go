package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency/real"
	"fmt"
)

const (
	BitcoinToYouName xrate.ExchangeName = "bitcoinToYou"
)

type BTYTicker struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
}

type bitcoinToYou struct {
	xrate.ExchangeParams
}

func BitcoinToYou() xrate.Exchange {
	return &bitcoinToYou{
		xrate.ExchangeParams{
			Name: BitcoinToYouName,
			CryptoCurrencies: []currency.Currency{
				bitcoin.Currency(),
			},
			FiatCurrencies: []currency.Currency{
				real.Currency(),
			},
			BaseApiURL: "https://www.bitcointoyou.com/api",
		},
	}
}

func (m *bitcoinToYou) getTickerURL(c currency.Currency) (string, error) {

	if !m.supportsCurrency(c) {
		return "", fmt.Errorf("exchange 'Bitcointoyou' does not support %s", c.Name)
	}

	return m.BaseApiURL + "/ticker.aspx", nil
}

func (m *bitcoinToYou) supportsCurrency(c currency.Currency) bool {

	return xrate.SliceContainsCurrency(m.CryptoCurrencies, c)
}

func (m *bitcoinToYou) getName() xrate.ExchangeName {

	return m.Name
}
