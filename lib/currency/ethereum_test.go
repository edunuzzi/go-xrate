package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthereum(t *testing.T) {

	t.Run("should return a new Ethereum", func(t *testing.T) {
		b := Ethereum()

		assert.Equal(t, EthereumName, b.Name)
		assert.Equal(t, EthereumAcronym, b.Acronym)
		assert.Equal(t, EthereumSymbol, b.Symbol)
		assert.Equal(t, EthereumDecimalPlaces, b.DecimalPlaces)
	})
}

func TestEthereumMainnet(t *testing.T) {

	t.Run("should return a new Ethereum mainnet", func(t *testing.T) {

		bm := EthereumMainnet()
		assert.Equal(t, Chain{ETH_MAINNET, Ethereum()}, bm)
	})
}

func TestEthereumTestnet(t *testing.T) {

	t.Run("should return a new Ethereum testnet", func(t *testing.T) {

		bt := EthereumTestnet()
		assert.Equal(t, Chain{ETH_TESTNET_ROPSTEN, Ethereum()}, bt)
	})
}

func TestEthereumTestnetRopsten(t *testing.T) {

	t.Run("should return a new Ethereum testnet Ropsten", func(t *testing.T) {

		bt := EthereumTestnetRopsten()
		assert.Equal(t, Chain{ETH_TESTNET_ROPSTEN, Ethereum()}, bt)
	})
}

func TestEthereumTestnetKovan(t *testing.T) {

	t.Run("should return a new Ethereum testnet Kovam", func(t *testing.T) {

		bt := EthereumTestnetKovan()
		assert.Equal(t, Chain{ETH_TESTNET_KOVAN, Ethereum()}, bt)
	})
}

func TestEthereumTestnetRinkeby(t *testing.T) {

	t.Run("should return a new Ethereum testnet", func(t *testing.T) {

		bt := EthereumTestnetRinkeby()
		assert.Equal(t, Chain{ETH_TESTNET_RINKEBY, Ethereum()}, bt)
	})
}

func TestIsValidEthereumChain(t *testing.T) {

	t.Run("should check if valid Ethereum chain", func(t *testing.T) {

		bt := EthereumTestnet()
		bm := EthereumMainnet()

		assert.True(t, IsValidEthereumChain(bt))
		assert.True(t, IsValidEthereumChain(bm))
		assert.False(t, IsValidEthereumChain(BitcoinMainnet()))
		assert.False(t, IsValidEthereumChain(Chain{BTC_MAINNET, Tether()}))
		assert.False(t, IsValidEthereumChain(Chain{BTC_MAINNET, Bitcoin()}))
		assert.False(t, IsValidEthereumChain(Chain{ETH_MAINNET, Tether()}))
	})
}

func TestIsValidEthereumNetwork(t *testing.T) {

	t.Run("should check if valid Ethereum network", func(t *testing.T) {

		assert.True(t, IsValidEthereumNetwork(ETH_MAINNET))
		assert.True(t, IsValidEthereumNetwork(ETH_TESTNET_ROPSTEN))
		assert.False(t, IsValidEthereumNetwork("wrong network"))
	})
}
