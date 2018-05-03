package exchanges

import (
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-xrate/lib/util"
	"time"
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
}

type Crawler struct {
	FiatCurrency   currency.Currency
	CryptoCurrency currency.Currency
	Exchanges      []Exchange
}

func (c *Crawler) Rates(timeout time.Duration) ([]CrawlerResponse, []error) {

	fetch := func(resChan chan *CrawlerResponse, errChan chan error, e Exchange) {

		tickerUrl, err := e.GetTickerURL(c.CryptoCurrency, c.FiatCurrency)

		if err != nil {
			errChan <- err
			resChan <- nil
			return
		}

		body, err := util.BaseGet(tickerUrl, timeout)

		if err != nil {
			errChan <- err
			resChan <- nil
			return
		}

		res, err := e.ConvertToResponse(c.CryptoCurrency, c.FiatCurrency, body)

		if err != nil {
			errChan <- err
			resChan <- nil
			return
		}

		resChan <- res
		errChan <- nil
	}

	var resList []CrawlerResponse
	var errList []error

	numberOfExchanges := len(c.Exchanges)

	resChan := make(chan *CrawlerResponse, numberOfExchanges)
	errChan := make(chan error, numberOfExchanges)

	for _, e := range c.Exchanges {
		go fetch(resChan, errChan, e)
	}

	for range c.Exchanges {
		res := <-resChan
		err := <-errChan

		if res != nil {
			resList = append(resList, *res)
		}

		if err != nil {
			errList = append(errList, err)
		}
	}

	return resList, errList
}
