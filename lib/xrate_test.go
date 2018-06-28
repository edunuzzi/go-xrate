package xrate

import (
	"fmt"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"github.com/Swipecoin/go-xrate/lib/exchanges"
	"testing"
	"time"
)

// FIXME
func TestNewBTCCrawler(t *testing.T) {

	for {
		crawler, _ := NewBTCCrawler(currency.Real(), exchanges.Foxbit(), exchanges.BitcoinTrade())

		resps := crawler.Rates(0)

		for _, resp := range resps {
			fmt.Println(resp)
		}

		time.Sleep(2 * time.Second)
	}
}
