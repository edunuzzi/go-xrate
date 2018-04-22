package br

type BTTicker struct {
	High float32 `json:"high"`
	Low  float32 `json:"low"`
	Vol  float32 `json:"volume"`
	Last float32 `json:"last"`
	Buy  float32 `json:"buy"`
	Sell float32 `json:"sell"`
}

type BitcoinTradeResponseBody struct {
	Data BTTicker `json:"data"`
}

type BitcoinTradeCrawler struct {
	*xrate.ExchangeCrawler
}

func NewBitcoinTradeCrawler() BitcoinTradeCrawler {
	return BitcoinTradeCrawler{
		&xrate.ExchangeCrawler{BaseUrl: "https://api.bitcointrade.com.br/v1/public"},
	}
}

func (f BitcoinTradeCrawler) BTC() xrate.CryptoCurrencyTicker {
	var t BitcoinTradeResponseBody

	xrate.BaseGet(
		f.BaseUrl+"/"+string(xrate.BTC)+"/ticker",
		&t,
	)

	return xrate.CryptoCurrencyTicker{
		Acronym:             xrate.BTC,
		FiatCurrencyAcronym: xrate.BRL,
		Last:                t.Data.Last,
		High24h:             t.Data.High,
		Low24h:              t.Data.Low,
		Volume24h:           t.Data.Vol,
		VolumeFiat24h:       0.0,
		RecentBuyOrder:      t.Data.Buy,
		RecentSellOrder:     t.Data.Sell,
	}
}
