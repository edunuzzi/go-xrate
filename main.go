package main

import (
	"github.com/Swipecoin/go-xrate/lib"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"github.com/Swipecoin/go-xrate/lib/exchanges"
	"time"
	"fmt"
)

func main() {
	c, _ := xrate.NewBTCCrawler(
		currency.Tether(),
		exchanges.BitcoinTrade(),
		exchanges.Foxbit(),
		exchanges.MercadoBitcoin(),
	)

	r := c.Rates(time.Second * 10)

	fmt.Println(
		r[0].Exchange,
		r[0].CryptoCurrency,
		r[0].FiatCurrency,
		r[0].Last,
		r[0].High24h,
		r[0].Low24h,
		r[0].Volume24h,
		r[0].VolumeFiat24h,
		r[0].MostRecentBidOrder,
		r[0].MostRecentAskOrder,
		r[0].CreatedAt,
	)
}
