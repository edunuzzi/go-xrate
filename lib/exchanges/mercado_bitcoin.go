package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency/real"
	"fmt"
	"encoding/json"
	"github.com/Swipecoin/go-xrate/lib/util"
)

const (
	MercadoBitcoinName xrate.ExchangeName = "Mercado Bitcoin"
)

type MBTicker struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
}

type MercadoBitcoinResponseBody struct {
	Ticker MBTicker `json:"ticker"`
}

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

func (m *mercadoBitcoin) GetTickerURL(cc currency.Currency) (string, error) {

	if !m.SupportsCryptoCurrency(cc) {
		return "", fmt.Errorf("exchange 'Mercado Bitcoin' does not support %s", cc.Name)
	}

	return m.BaseApiURL + "/" + string(cc.Acronym) + "/ticker", nil
}

func (m *mercadoBitcoin) SupportsFiatCurrency(f currency.Currency) bool {

	return xrate.SliceContainsCurrency(m.FiatCurrencies, f)
}

func (m *mercadoBitcoin) SupportsCryptoCurrency(cc currency.Currency) bool {

	return xrate.SliceContainsCurrency(m.CryptoCurrencies, cc)
}

func (m *mercadoBitcoin) GetName() xrate.ExchangeName {

	return m.Name
}

func (m *mercadoBitcoin) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*xrate.CrawlerResponse, error) {

	var res MercadoBitcoinResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &xrate.CrawlerResponse{
		Exchange:            m.ExchangeParams,
		CryptoCurrency:      cc,
		FiatCurrency:        fc,
		Last:                util.StringToFloat32(res.Ticker.Last),
		High24h:             util.StringToFloat32(res.Ticker.High),
		Low24h:              util.StringToFloat32(res.Ticker.Low),
		Volume24h:           util.StringToFloat32(res.Ticker.Vol),
		VolumeFiat24h:       xrate.UnsupportedField,
		MostRecentBuyOrder:  util.StringToFloat32(res.Ticker.Buy),
		MostRecentSellOrder: util.StringToFloat32(res.Ticker.Sell),
	}, nil
}
