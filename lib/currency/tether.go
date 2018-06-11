package currency

const (
	TetherName          Name = "Tether"
	TetherAcronym            = "USDT"
	TetherSymbol             = "USDT"
	TetherDecimalPlaces uint = 2
)

func Tether() Currency {
	return Currency{
		Name:          TetherName,
		Symbol:        TetherSymbol,
		Acronym:       TetherAcronym,
		DecimalPlaces: TetherDecimalPlaces,
	}
}
