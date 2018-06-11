package currency

const (
	DollarName          Name = "Dollar"
	DollarAcronym            = "USD"
	DollarSymbol             = "$"
	DollarDecimalPlaces uint = 2
)

func Dollar() Currency {
	return Currency{
		Name:          DollarName,
		Symbol:        DollarSymbol,
		Acronym:       DollarAcronym,
		DecimalPlaces: DollarDecimalPlaces,
	}
}
