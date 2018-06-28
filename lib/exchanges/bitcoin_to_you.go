package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"github.com/Swipecoin/go-xrate/lib/util"
	"time"
)

const (
	BitcoinToYouName ExchangeName = "BitcoinToYou"
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
	ExchangeParams
}

func BitcoinToYou() Exchange {
	return &bitcoinToYou{
		ExchangeParams{
			Name: BitcoinToYouName,
			CryptoCurrencies: []currency.Currency{
				currency.Bitcoin(),
			},
			FiatCurrencies: []currency.Currency{
				currency.Real(),
			},
			BaseApiURL: "https://www.bitcointoyou.com/api",
		},
	}
}

func (bty *bitcoinToYou) GetTickerURL(cc currency.Currency, _ currency.Currency) (string, error) {

	if !bty.SupportsCryptoCurrency(cc) {
		return "", fmt.Errorf("exchange 'Bitcointoyou' does not support %s", cc.Name)
	}

	return bty.BaseApiURL + "/ticker.aspx", nil
}

func (bty *bitcoinToYou) SupportsFiatCurrency(fc currency.Currency) bool {

	return SliceContainsCurrency(bty.FiatCurrencies, fc)
}

func (bty *bitcoinToYou) SupportsCryptoCurrency(cc currency.Currency) bool {

	return SliceContainsCurrency(bty.CryptoCurrencies, cc)
}

func (bty *bitcoinToYou) GetName() ExchangeName {

	return bty.Name
}

func (bty *bitcoinToYou) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*CrawlerResponse, error) {

	var res BitcoinToYouResponseBody

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	last, err := util.StringToFloat64(res.Ticker.Last)
	if err != nil {
		return nil, err
	}

	high, err := util.StringToFloat64(res.Ticker.High)
	if err != nil {
		return nil, err
	}

	low, err := util.StringToFloat64(res.Ticker.Low)
	if err != nil {
		return nil, err
	}

	vol, err := util.StringToFloat64(res.Ticker.Vol)
	if err != nil {
		return nil, err
	}

	bid, err := util.StringToFloat64(res.Ticker.Buy)
	if err != nil {
		return nil, err
	}

	ask, err := util.StringToFloat64(res.Ticker.Sell)
	if err != nil {
		return nil, err
	}

	return &CrawlerResponse{
		Exchange:           bty.ExchangeParams,
		CryptoCurrency:     cc,
		FiatCurrency:       fc,
		Last:               last,
		High24h:            high,
		Low24h:             low,
		Volume24h:          vol,
		VolumeFiat24h:      UnsupportedField,
		MostRecentBidOrder: bid,
		MostRecentAskOrder: ask,
		CreatedAt:          time.Now(),
	}, nil
}
