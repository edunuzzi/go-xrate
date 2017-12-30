package br

import (
	"go-xrate"
)

type FoxBitResponseBody struct {
	High    float32 `json:"high"`
	Vol     float32 `json:"vol"`
	Buy     float32 `json:"buy"`
	Last    float32 `json:"last"`
	Low     float32 `json:"low"`
	Sell    float32 `json:"sell"`
	Vol_brl float32 `json:"vol_brl"`
}

type FoxBitCrawler struct {
	*xrate.ExchangeCrawler
}

func NewFoxBitCrawler() FoxBitCrawler {
	return FoxBitCrawler{
		&xrate.ExchangeCrawler{BaseUrl: "https://api.blinktrade.com/api/v1/BRL"},
	}
}

func (f FoxBitCrawler) BTC() xrate.CryptoCurrencyTicker {
	var t FoxBitResponseBody

	xrate.BaseGet(
		f.BaseUrl+"/ticker?crypto_currency="+string(xrate.BTC),
		&t,
	)

	return xrate.CryptoCurrencyTicker{
		Acronym:             xrate.BTC,
		FiatCurrencyAcronym: xrate.BRL,
		Last:                t.Last,
		High24h:             t.High,
		Low24h:              t.Low,
		Volume24h:           t.Vol,
		VolumeFiat24h:       t.Vol_brl,
		RecentBuyOrder:      t.Buy,
		RecentSellOrder:     t.Sell,
	}
}
