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

type BitcoinToYouResponseBody struct {
	Ticker BTYTicker `json:"ticker"`
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

func (bty *bitcoinToYou) GetTickerURL(cc currency.Currency) (string, error) {

	if !bty.SupportsCryptoCurrency(cc) {
		return "", fmt.Errorf("exchange 'Bitcointoyou' does not support %s", cc.Name)
	}

	return bty.BaseApiURL + "/ticker.aspx", nil
}

func (bty *bitcoinToYou) SupportsFiatCurrency(fc currency.Currency) bool {

	return xrate.SliceContainsCurrency(bty.FiatCurrencies, fc)
}

func (bty *bitcoinToYou) SupportsCryptoCurrency(cc currency.Currency) bool {

	return xrate.SliceContainsCurrency(bty.CryptoCurrencies, cc)
}

func (bty *bitcoinToYou) GetName() xrate.ExchangeName {

	return bty.Name
}

func (bty *bitcoinToYou) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*xrate.CrawlerResponse, error) {

	var res BitcoinToYouResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &xrate.CrawlerResponse{
		Exchange:            bty.ExchangeParams,
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
