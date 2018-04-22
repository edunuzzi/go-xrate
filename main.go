package main

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-xrate/lib/exchanges"
)

func main() {
	crawler, _ := xrate.NewBTCCrawler(
		exchanges.MercadoBitcoin(),
		exchanges.BitcoinToYou(),
	)

	crawler.Rates()
}