package xrate

import (
	"fmt"
	"github.com/Swipecoin/go-currency/currency/bitcoin"
)

func NewBTCCrawler(exchanges ...Exchange) (*crawler, error) {

	crawler := &crawler{bitcoin.Currency(), exchanges}

	for _, e := range crawler.exchanges {

		if !e.supportsCurrency(crawler.Currency) {
			return nil, fmt.Errorf("exchange '%s' does not support currency %s", e.getName(), crawler.Currency.Name)
		}
	}

	return crawler, nil
}