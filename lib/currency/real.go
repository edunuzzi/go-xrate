package currency

const (
	RealName          Name = "Real"
	RealAcronym            = "BRL"
	RealSymbol             = "R$"
	RealDecimalPlaces uint = 2
)

func Real() Currency {
	return Currency{
		Name:          RealName,
		Symbol:        RealSymbol,
		Acronym:       RealAcronym,
		DecimalPlaces: RealDecimalPlaces,
	}
}
