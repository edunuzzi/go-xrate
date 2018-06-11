package currency

const (
	ETH_MAINNET         Network = "eth_mainnet"
	ETH_TESTNET_ROPSTEN Network = "eth_testnet_ropsten"
	ETH_TESTNET_KOVAN   Network = "eth_testnet_kovan"
	ETH_TESTNET_RINKEBY Network = "eth_testnet_rinkeby"
)

const (
	EthereumName          Name = "Ethereum"
	EthereumAcronym            = "ETH"
	EthereumSymbol             = "ETH"
	EthereumDecimalPlaces uint = 18
)

func Ethereum() Currency {
	return Currency{
		Name:          EthereumName,
		Symbol:        EthereumAcronym,
		Acronym:       EthereumSymbol,
		DecimalPlaces: EthereumDecimalPlaces,
	}
}

func EthereumMainnet() Chain {
	return Chain{
		Network:  ETH_MAINNET,
		Currency: Ethereum(),
	}
}

func EthereumTestnet() Chain {
	return Chain{
		Network:  ETH_TESTNET_ROPSTEN,
		Currency: Ethereum(),
	}
}

func EthereumTestnetRopsten() Chain {
	return Chain{
		Network:  ETH_TESTNET_ROPSTEN,
		Currency: Ethereum(),
	}
}

func EthereumTestnetKovan() Chain {
	return Chain{
		Network:  ETH_TESTNET_KOVAN,
		Currency: Ethereum(),
	}
}

func EthereumTestnetRinkeby() Chain {
	return Chain{
		Network:  ETH_TESTNET_RINKEBY,
		Currency: Ethereum(),
	}
}

func IsValidEthereumChain(c Chain) bool {
	return c.Currency.Name == EthereumName && IsValidEthereumNetwork(c.Network)
}

func IsValidEthereumNetwork(n Network) bool {
	return n == ETH_MAINNET || n == ETH_TESTNET_ROPSTEN || n == ETH_TESTNET_RINKEBY || n == ETH_TESTNET_KOVAN
}
