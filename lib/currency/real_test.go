package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReal(t *testing.T) {

	t.Run("should return a new Real", func(t *testing.T) {

		b := Real()

		assert.Equal(t, RealName, b.Name)
		assert.Equal(t, RealAcronym, b.Acronym)
		assert.Equal(t, RealSymbol, b.Symbol)
		assert.Equal(t, RealDecimalPlaces, b.DecimalPlaces)
	})
}
