package main

import (
	"github.com/Swipecoin/go-xrate/lib"
	"fmt"
	"time"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"github.com/Swipecoin/go-xrate/lib/exchanges"
)

func main() {
	c, _ := xrate.NewBTCCrawler(
		currency.Real(),
		exchanges.BitcoinToYou(),
	)

	r := c.Rates(time.Second * 5)

	fmt.Println(r)
}