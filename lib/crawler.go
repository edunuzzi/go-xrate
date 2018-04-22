package xrate

import (
	"github.com/Swipecoin/go-currency/currency"
)

type CrawlerResponse struct {
	CryptoCurrency      currency.Currency
	FiatCurrency        currency.Currency
	Last                float32
	High24h             float32
	Low24h              float32
	Volume24h           float32
	VolumeFiat24h       float32
	MostRecentBuyOrder  float32
	MostRecentSellOrder float32
}

type crawler struct {
	currency currency.Currency
	exchanges []Exchange
}

func (c *crawler) Rates() ([]CrawlerResponse, error) {
}
