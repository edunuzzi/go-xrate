package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-xrate/lib/util"
)

const (
	BinanceName xrate.ExchangeName = "Binance"
)

type BinanceTicker struct {
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	Count              uint32 `json:"count"`
}

type binance struct {
	xrate.ExchangeParams
}

func Binance() xrate.Exchange {
	return &binance{
		xrate.ExchangeParams{
			Name: BinanceName,
			CryptoCurrencies: []currency.Currency{
				currency.Bitcoin(),
			},
			FiatCurrencies: []currency.Currency{
				currency.Tether(),
			},
			BaseApiURL: "https://api.binance.com",
		},
	}
}

func (b *binance) GetTickerURL(cc currency.Currency, fc currency.Currency) (string, error) {

	if !b.SupportsCryptoCurrency(cc) {
		return "", fmt.Errorf("exchange 'Bitcointoyou' does not support %s", cc.Name)
	}

	if !b.SupportsFiatCurrency(fc) {
		return "", fmt.Errorf("exchange 'Bitcointoyou' does not support %s", fc.Name)
	}

	return b.BaseApiURL + "/api/v1/ticker/24hr" + "?symbol=" + cc.Acronym + fc.Acronym, nil
}

func (b *binance) SupportsFiatCurrency(fc currency.Currency) bool {

	return xrate.SliceContainsCurrency(b.FiatCurrencies, fc)
}

func (b *binance) SupportsCryptoCurrency(cc currency.Currency) bool {

	return xrate.SliceContainsCurrency(b.CryptoCurrencies, cc)
}

func (b *binance) GetName() xrate.ExchangeName {

	return b.Name
}

func (b *binance) ConvertToResponse(cc currency.Currency, fc currency.Currency, body []byte) (*xrate.CrawlerResponse, error) {

	var res BinanceTicker

	err := json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	last, err := util.StringToFloat32(res.LastPrice)
	if err != nil {
		return nil, err
	}

	high, err := util.StringToFloat32(res.HighPrice)
	if err != nil {
		return nil, err
	}

	low, err := util.StringToFloat32(res.LowPrice)
	if err != nil {
		return nil, err
	}

	vol, err := util.StringToFloat32(res.Volume)
	if err != nil {
		return nil, err
	}

	bid, err := util.StringToFloat32(res.BidPrice)
	if err != nil {
		return nil, err
	}

	ask, err := util.StringToFloat32(res.AskPrice)
	if err != nil {
		return nil, err
	}

	return &xrate.CrawlerResponse{
		Exchange:           b.ExchangeParams,
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
