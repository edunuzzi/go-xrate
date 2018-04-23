package xrate

import (
	"github.com/Swipecoin/go-currency/currency"
	"github.com/Swipecoin/go-xrate/lib/util"
	"time"
)

const UnsupportedField float32 = -1

type CrawlerResponse struct {
	Exchange            ExchangeParams    `json:"exchange,omitempty"`
	CryptoCurrency      currency.Currency `json:"crypto_currency,omitempty"`
	FiatCurrency        currency.Currency `json:"fiat_currency,omitempty"`
	Last                float32           `json:"last,omitempty"`
	High24h             float32           `json:"high_24h,omitempty"`
	Low24h              float32           `json:"low_24h,omitempty"`
	Volume24h           float32           `json:"volume_24h,omitempty"`
	VolumeFiat24h       float32           `json:"volume_fiat_24h,omitempty"`
	MostRecentBuyOrder  float32           `json:"most_recent_buy_order,omitempty"`
	MostRecentSellOrder float32           `json:"most_recent_sell_order,omitempty"`
}

type crawler struct {
	fiatCurrency   currency.Currency
	cryptoCurrency currency.Currency
	exchanges      []Exchange
}

func (c *crawler) Rates(timeout time.Duration) ([]CrawlerResponse, []error) {

	fetch := func(resChan chan *CrawlerResponse, errChan chan error, e Exchange) {

		tickerUrl, err := e.GetTickerURL(c.cryptoCurrency)

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

		res, err := e.ConvertToResponse(c.cryptoCurrency, c.fiatCurrency, body)

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

	numberOfExchanges := len(c.exchanges)

	resChan := make(chan *CrawlerResponse, numberOfExchanges)
	errChan := make(chan error, numberOfExchanges)

	for _, e := range c.exchanges {
		go fetch(resChan, errChan, e)
	}

	for range c.exchanges {
		res := <- resChan
		err := <- errChan

		if res != nil {
			resList = append(resList, *res)
		}

		if err != nil {
			errList = append(errList, err)
		}
	}

	return resList, errList
}


