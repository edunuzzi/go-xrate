package exchanges

import (
	"fmt"
	"github.com/Swipecoin/go-xrate/lib"
)

func GetExchangeByName(name xrate.ExchangeName) (xrate.Exchange, error) {

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