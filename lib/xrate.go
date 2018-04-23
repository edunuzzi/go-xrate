package xrate

import (
	"fmt"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
	"github.com/Swipecoin/go-currency/currency"
)

func NewBTCCrawler(fiatCurrency currency.Currency, exchanges ...Exchange) (*crawler, error) {

	crawler := &crawler{
		fiatCurrency: fiatCurrency,
		cryptoCurrency: bitcoin.Currency(),
		exchanges: exchanges,
	}

	for _, e := range crawler.exchanges {

		if !e.SupportsCryptoCurrency(crawler.cryptoCurrency) {
			return nil, fmt.Errorf("exchange '%s' does not support currency %s", e.GetName(), crawler.cryptoCurrency.Name)
		}

		if !e.SupportsFiatCurrency(crawler.fiatCurrency) {
			return nil, fmt.Errorf("exchange '%s' does not support currency %s", e.GetName(), crawler.fiatCurrency.Name)
		}
	}

	return crawler, nil
}