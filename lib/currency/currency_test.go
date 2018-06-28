package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrency(t *testing.T) {

	t.Run("should successfully convert to main unit", func(t *testing.T) {

		assert.Equal(t, float64(1), Bitcoin().ToMainUnit(100000000))
		assert.Equal(t, float64(8.03), Bitcoin().ToMainUnit(803000000))
		assert.Equal(t, float64(15), Real().ToMainUnit(1500))
		assert.Equal(t, float64(1350), Real().ToMainUnit(135000))
	})

	t.Run("should successfully convert to fractional unit", func(t *testing.T) {

		assert.Equal(t, float64(100000000), Bitcoin().ToFractionalUnit(1))
		assert.Equal(t, float64(750000000), Bitcoin().ToFractionalUnit(7.5))
		assert.Equal(t, float64(950), Dollar().ToFractionalUnit(9.5))
		assert.Equal(t, float64(100000), Dollar().ToFractionalUnit(1000))
	})
}

func TestGetCurrencyByAcronym(t *testing.T) {

	t.Run("should successfully return new Currency for each acronym", func(t *testing.T) {

		base := func(acronym string, constructor func() Currency, shouldSuccess bool) {
			c, err := GetCurrencyByAcronym(acronym)

			assert.Nil(t, err)

			if !shouldSuccess {
				assert.NotEqual(t, constructor(), c)
				return
			}

			assert.Equal(t, constructor(), c)
		}

		base(BitcoinAcronym, Bitcoin, true)
		base(EthereumAcronym, Bitcoin, false)

		base(EthereumAcronym, Ethereum, true)
		base(BitcoinAcronym, Ethereum, false)

		base(DollarAcronym, Dollar, true)
		base(BitcoinAcronym, Dollar, false)

		base(RealAcronym, Real, true)
		base(BitcoinAcronym, Real, false)

		base(TetherAcronym, Tether, true)
		base(BitcoinAcronym, Tether, false)

		c, err := GetCurrencyByAcronym("wrong")

		assert.Equal(t, Currency{}, c)
		assert.NotNil(t, err)
	})
}
