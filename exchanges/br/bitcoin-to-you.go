package br

import (
	"go-xrate"
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

type BitcoinToYouCrawler struct {
	*xrate.ExchangeCrawler
}

func NewBitcoinToYouCrawler() BitcoinToYouCrawler {
	return BitcoinToYouCrawler{
		&xrate.ExchangeCrawler{BaseUrl: "https://www.bitcointoyou.com/api"},
	}
}

func (f BitcoinToYouCrawler) BTC() xrate.CryptoCurrencyTicker {
	var t BitcoinToYouResponseBody

	xrate.BaseGet(
		f.BaseUrl+"/ticker.aspx",
		&t,
	)

	return xrate.CryptoCurrencyTicker{
		Acronym:             xrate.BTC,
		FiatCurrencyAcronym: xrate.BRL,
		Last:                xrate.StringToFloat32(t.Ticker.Last),
		High24h:             xrate.StringToFloat32(t.Ticker.High),
		Low24h:              xrate.StringToFloat32(t.Ticker.Low),
		Volume24h:           xrate.StringToFloat32(t.Ticker.Vol),
		VolumeFiat24h:       0.0,
		RecentBuyOrder:      xrate.StringToFloat32(t.Ticker.Buy),
		RecentSellOrder:     xrate.StringToFloat32(t.Ticker.Sell),
	}
}
