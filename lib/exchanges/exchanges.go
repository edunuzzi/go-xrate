package exchanges

import (
	"fmt"
)

func GetExchangeByName(name ExchangeName) (Exchange, error) {

	switch name {
	case BinanceName:
		return Binance(), nil

	case FoxbitName:
		return Foxbit(), nil

	case BitcoinToYouName:
		return BitcoinToYou(), nil

	case BitcoinTradeName:
		return BitcoinTrade(), nil

	case MercadoBitcoinName:
		return MercadoBitcoin(), nil

	default:
		return nil, fmt.Errorf("Invalid exchange name: %s", name)
	}
}
