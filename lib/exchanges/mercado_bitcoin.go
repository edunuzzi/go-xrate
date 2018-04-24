package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-currency/currency"
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
				currency.Bitcoin(),
			},
			FiatCurrencies: []currency.Currency{
				currency.Real(),
			},
			BaseApiURL: "https://www.mercadobitcoin.net/api",
		},
	}
}

func (m *mercadoBitcoin) GetTickerURL(cc currency.Currency, _ currency.Currency) (string, error) {

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

	last, err := util.StringToFloat32(res.Ticker.Last)
	if err != nil { return nil, err }

	high, err := util.StringToFloat32(res.Ticker.High)
	if err != nil { return nil, err }

	low, err := util.StringToFloat32(res.Ticker.Low)
	if err != nil { return nil, err }

	vol, err := util.StringToFloat32(res.Ticker.Vol)
	if err != nil { return nil, err }

	bid, err := util.StringToFloat32(res.Ticker.Buy)
	if err != nil { return nil, err }

	ask, err := util.StringToFloat32(res.Ticker.Sell)
	if err != nil { return nil, err }

	return &xrate.CrawlerResponse{
		Exchange:           m.ExchangeParams,
		CryptoCurrency:     cc,
		FiatCurrency:       fc,
		Last:               last,
		High24h:            high,
		Low24h:             low,
		Volume24h:          vol,
		VolumeFiat24h:      xrate.UnsupportedField,
		MostRecentBidOrder: bid,
		MostRecentAskOrder: ask,
	}, nil
}
