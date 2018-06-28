# go-xrate

Small lib for getting cryptocurrency rates from several Exchanges around the world.

Released under the terms of the [MIT LICENSE](LICENSE).

## Installation

To install, just run 
```bash
go get github.com/swipecoin/go-xrate
```

It doesn't rely on any external lib :D
  
## Supported Exchanges

As of now, only a few exchanges from Brazil and Binance are supported. But the code is prepared to work with multiple countries, fiat currencies and cryptocurrencies. 

Here's the supported list (We are constantly adding more and more):
- [FoxBit](https://foxbit.exchange)
- [Mercado Bitcoin](https://mercadobitcoin.com.br) 
- [BitcoinTrade](https://bitcointrade.com.br)
- [Bitcoin To You](https://bitcointoyou.com)
- [Binance](https://binance.com)

## Supported Cryptocurrencies

- BTC: Bitcoin

## Usage

### BTC -> USDT - Binance
```golang
crawler, _ := xrate.NewBTCCrawler(
    currency.Real(), 
    exchanges.Binance(),
)
	
// Here we are fetching the rates with a 10 second timeout for each request 
r, _ := c.Rates(time.Second * 10)

// r[0] -> response from Binance
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
```

### BTC -> BRL - Foxbit + Mercado Bitcoin
```golang
crawler, _ := xrate.NewBTCCrawler(
    currency.Real(), 
    exchanges.BitcoinTrade(),
    exchanges.Foxbit(),
    exchanges.MercadoBitcoin(),
)
	
// Here we are fetching the rates with a 10 second timeout for each request 
r, _ := c.Rates(time.Second * 10)

// r[0] -> response from BitcoinTrade
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

// r[1] -> response from Foxbit
fmt.Println(
    r[1].Exchange,
    r[1].CryptoCurrency,
    r[1].FiatCurrency,
    r[1].Last,
    r[1].High24h,
    r[1].Low24h,
    r[1].Volume24h,
    r[1].VolumeFiat24h,
    r[1].MostRecentBidOrder,
    r[1].MostRecentAskOrder,
    r[1].CreatedAt,
)

// r[2] -> response from Mercado Bitcoin
fmt.Println(
    r[2].Exchange,
    r[2].CryptoCurrency,
    r[2].FiatCurrency,
    r[2].Last,
    r[2].High24h,
    r[2].Low24h,
    r[2].Volume24h,
    r[2].VolumeFiat24h,
    r[2].MostRecentBidOrder,
    r[2].MostRecentAskOrder,
    r[2].CreatedAt,
)
```

## TODO
- [ ] Report badge
- [ ] Unit tests
- [ ] Support more Exchanges
- [ ] Better error handling
- [ ] Other useful statistics from a specific cryptocurrency and/or exchange

## Contribution
Please feel free to contribute with both suggestions and pull requests :D