package br

import (
	"go-xrate"
)

type Ticker struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
}

type MercadoBitcoinResponseBody struct {
	Ticker `json:"ticker"`
}

type MercadoBitcoinCrawler struct {
	*xrate.ExchangeCrawler
}

func NewMercadoBitcoinCrawler() MercadoBitcoinCrawler {
	return MercadoBitcoinCrawler{
		&xrate.ExchangeCrawler{BaseUrl: "https://www.mercadobitcoin.net/api"},
	}
}

func (m MercadoBitcoinCrawler) BTC() xrate.CryptoCurrencyTicker {
	var t MercadoBitcoinResponseBody

	xrate.BaseGet(
		m.BaseUrl+"/"+string(xrate.BTC)+"/ticker",
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

func (m MercadoBitcoinCrawler) BCH() xrate.CryptoCurrencyTicker {
	var t MercadoBitcoinResponseBody

	xrate.BaseGet(
		m.BaseUrl+"/"+string(xrate.BCH)+"/ticker",
		&t,
	)

	return xrate.CryptoCurrencyTicker{
		Acronym:             xrate.BCH,
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

func (m MercadoBitcoinCrawler) LTC() xrate.CryptoCurrencyTicker {
	var t MercadoBitcoinResponseBody

	xrate.BaseGet(
		m.BaseUrl+"/"+string(xrate.LTC)+"/ticker",
		&t,
	)

	return xrate.CryptoCurrencyTicker{
		Acronym:             xrate.LTC,
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
