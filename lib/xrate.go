package xrate

import (
	"fmt"
	"github.com/Swipecoin/go-currency/currency"

	"github.com/Swipecoin/go-xrate/lib/exchanges"
	"time"
)
//FIXME test concurrency access (mutex)
var cache exchanges.CrawlerResponse

func GetRate() exchanges.CrawlerResponse {
	return cache
}

func CacheTheLowestPrice(responses []exchanges.CrawlerResponse) {

	lowestResponse := exchanges.CrawlerResponse{}
	for _, resp1 := range responses {
		lowestResponse = resp1
		for _, resp2 := range responses {
			if resp2.Last < resp1.Last {
				lowestResponse = resp2
			}
		}
	}

	cache = lowestResponse
}

func StartBTCCrawler() {

	for {

		crawler, _ := NewBTCCrawler(currency.Real(), exchanges.Foxbit(), exchanges.BitcoinTrade())

		responses, _ := crawler.Rates(0)

		CacheTheLowestPrice(responses)

		time.Sleep(10 * time.Minute)
	}
}



func NewBTCCrawler(fiatCurrency currency.Currency, exs ...exchanges.Exchange) (*exchanges.Crawler, error) {

	crawler := &exchanges.Crawler{
		FiatCurrency:   fiatCurrency,
		CryptoCurrency: currency.Bitcoin(),
		Exchanges:      exs,
	}

	for _, e := range crawler.Exchanges {

		if !e.SupportsCryptoCurrency(crawler.CryptoCurrency) {
			return nil, fmt.Errorf("exchange '%s' does not support currency %s", e.GetName(), crawler.CryptoCurrency.Name)
		}

		if !e.SupportsFiatCurrency(crawler.FiatCurrency) {
			return nil, fmt.Errorf("exchange '%s' does not support currency %s", e.GetName(), crawler.FiatCurrency.Name)
		}
	}

	return crawler, nil
}
