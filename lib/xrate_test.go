package xrate

import (
	"testing"
	"github.com/Swipecoin/go-xrate/lib/currency"
	"github.com/Swipecoin/go-xrate/lib/exchanges"
	"github.com/stretchr/testify/assert"
)

func TestNewBTCCrawler(t *testing.T) {

	t.Run("should return error if exchange does not support fiat currency", func(t *testing.T) {
		c, err := NewBTCCrawler(currency.Real(), exchanges.Binance())

		assert.NotNil(t, err)
		assert.Empty(t, c)

		c, err = NewBTCCrawler(currency.Real(), exchanges.Foxbit(), exchanges.Binance())

		assert.NotNil(t, err)
		assert.Empty(t, c)
	})

	t.Run("should successfully return a new crawler without error", func(t *testing.T) {
		c, err := NewBTCCrawler(currency.Tether(), exchanges.Binance())

		assert.Nil(t, err)
		assert.NotEmpty(t, c)

		c, err = NewBTCCrawler(currency.Real(), exchanges.Foxbit(), exchanges.BitcoinTrade())

		assert.Nil(t, err)
		assert.NotEmpty(t, c)
	})
}
