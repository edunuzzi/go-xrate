package currency

import (
	"fmt"
	"math"
)

type Name string

type Currency struct {
	Name          Name
	Symbol        string
	Acronym       string
	DecimalPlaces uint
}

func (c Currency) ToMainUnit(v float64) float64 {
	return v / math.Pow(10, float64(c.DecimalPlaces))
}

func (c Currency) ToFractionalUnit(v float64) float64 {
	return float64(
		v * math.Pow(10, float64(c.DecimalPlaces)),
	)
}

func GetCurrencyByAcronym(acronym string) (Currency, error) {

	switch acronym {
	case BitcoinAcronym:
		return Bitcoin(), nil

	case EthereumAcronym:
		return Ethereum(), nil

	case DollarAcronym:
		return Dollar(), nil

	case RealAcronym:
		return Real(), nil

	case TetherAcronym:
		return Tether(), nil

	default:
		return Currency{}, fmt.Errorf("invalid currency acronym: %s", acronym)
	}
}
