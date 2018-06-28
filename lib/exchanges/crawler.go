package exchanges

import (
	"github.com/Swipecoin/go-xrate/lib/util"
	"time"
	"github.com/Swipecoin/go-xrate/lib/currency"
)

const UnsupportedField float32 = -1

type CrawlerResponse struct {
	Exchange           ExchangeParams    `json:"exchange,omitempty"`
	CryptoCurrency     currency.Currency `json:"crypto_currency,omitempty"`
	FiatCurrency       currency.Currency `json:"fiat_currency,omitempty"`
	Last               float32           `json:"last,omitempty"`
	High24h            float32           `json:"high_24h,omitempty"`
	Low24h             float32           `json:"low_24h,omitempty"`
	Volume24h          float32           `json:"volume_24h,omitempty"`
	VolumeFiat24h      float32           `json:"volume_fiat_24h,omitempty"`
	MostRecentBidOrder float32           `json:"most_recent_bid_order,omitempty"`
	MostRecentAskOrder float32           `json:"most_recent_ask_order,omitempty"`
	CreatedAt          time.Time         `json:"created_at"`
	Error              error             `json:"error"`
}

type Crawler struct {
	FiatCurrency   currency.Currency
	CryptoCurrency currency.Currency
	Exchanges      []Exchange
}

func (c *Crawler) Rates(timeout time.Duration) ([]CrawlerResponse) {

	fetch := func(resChan chan CrawlerResponse, e Exchange) {

		tickerUrl, err := e.GetTickerURL(c.CryptoCurrency, c.FiatCurrency)

		if err != nil {
			resChan <- nil
			return
		}

		body, err := util.BaseGet(tickerUrl, timeout)

		if err != nil {
			resChan <- CrawlerResponse{Error: err}
			return
		}

		res, err := e.ConvertToResponse(c.CryptoCurrency, c.FiatCurrency, body)
		res.Error = err

		resChan <- *res
	}

	var resList []CrawlerResponse

	numberOfExchanges := len(c.Exchanges)

	resChan := make(chan CrawlerResponse, numberOfExchanges)

	for _, e := range c.Exchanges {
		go fetch(resChan, e)
	}

	for range c.Exchanges {
		resList = append(resList, <-resChan)
	}

	return resList
}
