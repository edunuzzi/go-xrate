package currency

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDollar(t *testing.T) {

	t.Run("should return a new Dollar", func(t *testing.T) {

		b := Dollar()

		assert.Equal(t, DollarName, b.Name)
		assert.Equal(t, DollarAcronym, b.Acronym)
		assert.Equal(t, DollarSymbol, b.Symbol)
		assert.Equal(t, DollarDecimalPlaces, b.DecimalPlaces)
	})
}