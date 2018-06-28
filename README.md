# Go-xrate
[![Go Report Card](https://goreportcard.com/badge/github.com/Swipecoin/go-xrate)](https://goreportcard.com/report/github.com/Swipecoin/go-xrate)

Small Golang implementation for fetching cryptocurrency rates from several Exchanges around the world.

Released under the terms of the [MIT LICENSE](LICENSE).

## Installation

To install, just run 
```bash
go get github.com/swipecoin/go-xrate
```

It doesn't rely on any external lib :D

## Usage

### BTC -> USDT - Binance
```golang
crawler, _ := xrate.NewBTCCrawler(
    currency.Real(), 
    exchanges.Binance(),
)
	
// Here we are fetching the rates with a 5 second timeout for each request 
r := crawler.Rates(time.Second * 5)

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
	
// Here we are fetching the rates with a 5 second timeout for each request 
r := crawler.Rates(time.Second * 5)

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
    r[0].Error,
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
    r[1].Error,
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
    r[2].Error,
)
```

## API

### `NewBTCCrawler(currency.Currency, ...exchanges.Exchange) (exchanges.Crawler, error)`
This is used to create new crypto crawler. It expects a fiatCurrency and the exchanges you want to fetch.

PS: It will return a error if you pass it a exchange that does not support Bitcoin or the given fiat currency.

### `(exchanges.Crawler) Rates(time.Duration) ([]CrawlerResponse)` 
This is the method used to fetch the rates for the exchanges passed on the previous method. 
It receives a timeout and returns a list of responses, one for each exchange. 

The library uses goroutines for every api call. So the total duration is equal to the slowest Exchange API response time at that moment.

PS: If a specific exchange does not support some field, it will return -1 instead.

## Supported Exchanges

As of now, only a few exchanges are supported. But the code is prepared to work with multiple countries, fiat currencies and cryptocurrencies. 

Here's the currently supported list:

### [Foxbit](https://foxbit.exchange)

#### Fiat currencies:
- Real

#### Cryptocurrencies:
- BTC

---

### [Mercado Bitcoin](https://mercadobitcoin.com.br)

#### Fiat currencies:
- Real

#### Cryptocurrencies:
- BTC

---

### [BitcoinTrade](https://bitcointrade.com.br)

#### Fiat currencies:
- Real

#### Cryptocurrencies:
- BTC

---

### [Bitcoin To You](https://bitcointoyou.com)

#### Fiat currencies:
- Real

#### Cryptocurrencies:
- BTC

---

### [Binance](https://binance.com)

#### Fiat currencies:
- Tether

#### Cryptocurrencies:
- BTC

## TODO
- [X] Report badge
- [ ] Unit tests
- [ ] Support more Exchanges
- [ ] Better error handling
- [ ] Other useful statistics from a specific cryptocurrency and/or exchange

## Contribution
Please feel free to contribute with both suggestions and pull requests :D