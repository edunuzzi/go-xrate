package currency

const (
	BTC_MAINNET Network = "btc_mainnet"
	BTC_TESTNET Network = "btc_testnet3"
)

const (
	BitcoinName          Name = "Bitcoin"
	BitcoinAcronym            = "BTC"
	BitcoinSymbol             = "BTC"
	BitcoinDecimalPlaces uint = 8
)

func Bitcoin() Currency {
	return Currency{
		Name:          BitcoinName,
		Symbol:        BitcoinAcronym,
		Acronym:       BitcoinSymbol,
		DecimalPlaces: BitcoinDecimalPlaces,
	}
}

func BitcoinMainnet() Chain {
	return Chain{
		Network:  BTC_MAINNET,
		Currency: Bitcoin(),
	}
}

func BitcoinTestnet() Chain {
	return Chain{
		Network:  BTC_TESTNET,
		Currency: Bitcoin(),
	}
}

func IsValidBitcoinChain(c Chain) bool {
	return c.Currency.Name == BitcoinName && IsValidBitcoinNetwork(c.Network)
}

func IsValidBitcoinNetwork(n Network) bool {
	return n == BTC_MAINNET || n == BTC_TESTNET
}
