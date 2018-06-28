package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitcoin(t *testing.T) {

	t.Run("should return a new Bitcoin", func(t *testing.T) {
		b := Bitcoin()

		assert.Equal(t, BitcoinName, b.Name)
		assert.Equal(t, BitcoinAcronym, b.Acronym)
		assert.Equal(t, BitcoinSymbol, b.Symbol)
		assert.Equal(t, BitcoinDecimalPlaces, b.DecimalPlaces)
	})
}

func TestBitcoinMainnet(t *testing.T) {

	t.Run("should return a new Bitcoin mainnet", func(t *testing.T) {

		bm := BitcoinMainnet()
		assert.Equal(t, Chain{BTC_MAINNET, Bitcoin()}, bm)
	})
}

func TestBitcoinTestnet(t *testing.T) {

	t.Run("should return a new Bitcoin testnet", func(t *testing.T) {

		bt := BitcoinTestnet()
		assert.Equal(t, Chain{BTC_TESTNET, Bitcoin()}, bt)
	})
}

func TestIsValidBitcoinChain(t *testing.T) {

	t.Run("should check if valid bitcoin chain", func(t *testing.T) {

		bt := BitcoinTestnet()
		bm := BitcoinMainnet()

		assert.True(t, IsValidBitcoinChain(bt))
		assert.True(t, IsValidBitcoinChain(bm))
		assert.False(t, IsValidBitcoinChain(EthereumMainnet()))
		assert.False(t, IsValidBitcoinChain(Chain{ETH_MAINNET, Ethereum()}))
		assert.False(t, IsValidBitcoinChain(Chain{BTC_MAINNET, Ethereum()}))
	})
}

func TestIsValidBitcoinNetwork(t *testing.T) {

	t.Run("should check if valid bitcoin network", func(t *testing.T) {

		assert.True(t, IsValidBitcoinNetwork(BTC_MAINNET))
		assert.True(t, IsValidBitcoinNetwork(BTC_TESTNET))
		assert.False(t, IsValidBitcoinNetwork("wrong network"))
	})
}
