package xrate

import (
	"fmt"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
)

func NewBTCCrawler(exchanges ...Exchange) (*crawler, error) {

	crawler := &crawler{bitcoin.Currency(), exchanges}

	for _, e := range crawler.exchanges {

		if !e.supportsCurrency(crawler.currency) {
			return nil, fmt.Errorf("exchange '%s' does not support currency %s", e.getName(), crawler.currency.Name)
		}
	}

	return crawler, nil
}