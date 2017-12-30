# go-xrate

go-xrate is a small lib for getting cryptocurrency rates from several Exchanges around the world.

### Install

To install, just run 

`go get github.com/SwipeCoin/go-xrate`

It doesn't rely on any external lib :D
  
###Supported Exchanges

As of now, only exchanges from Brazil are supported. But the code is prepared to work with multiple countries, fiat currencies and cryptocurrencies. 

Here's the supported list (We are constantly adding more and more):
- [FoxBit](https://foxbit.exchange)
- [Mercado Bitcoin](https://wwww.mercadobitcoin.com.br) 
- [BitcoinTrade](https://bitcointrade.com.br)

### Usage
It's really simple to get started.

##### First, you import the package:
```
// Here we are importing from the brazilian package of exchanges
import (
    xrbr "go-xrate/exchanges/br"
)
```

#####And finally, instantiate a new ExchangeCrawler and call the specific method for the cryptocurrency you want:
```
// FoxBit - BTC
xrbr.NewFoxBitCrawler().BTC()
  
// OR 
// Mercado Bitcoin - BCH
xrbr.NewMercadoBitcoinCrawler().BCH()
 
// OR 
// BitcoinTrade - BTC
xrbr.NewBitcoinTradeCrawler().BTC()
```

### API
#### NewFoxBitCrawler() FoxBitCrawler
Creates and returns a new FoxBitCrawler

###### Methods:
- BTC() xrate.CryptoCurrencyTicker

#### NewMercadoBitcoinCrawler() MercadoBitcoinCrawler
Creates and returns a new MercadoBitcoinCrawler

###### Methods:
- BTC() xrate.CryptoCurrencyTicker
- LTC() xrate.CryptoCurrencyTicker
- BCH() xrate.CryptoCurrencyTicker

#### NewBitcoinTradeCrawler() BitcoinTradeCrawler
Creates and returns a new BitcoinTradeCrawler

###### Methods:
- BTC() xrate.CryptoCurrencyTicker

All methods return a xrate.CryptoCurrencyTicker:
```
type CryptoCurrencyTicker struct {
	Acronym             CryptoCurrencyAcronym
	FiatCurrencyAcronym FiatCurrencyAcronym
	Last                float32
	High24h             float32
	Low24h              float32
	Volume24h           float32
	VolumeFiat24h       float32
	RecentBuyOrder      float32
	RecentSellOrder     float32
}
```

In case when a exchange does not returns a specific field, it is set to it's 'falsy' value. (E.g. float32 -> 0.0)

### Support
Please feel free to give suggestions for improvements and implement any specific exchange code you may need. We are open to well written pull requests :P