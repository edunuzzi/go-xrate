package currency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTether(t *testing.T) {

	t.Run("should return a new Tether", func(t *testing.T) {

		b := Tether()

		assert.Equal(t, TetherName, b.Name)
		assert.Equal(t, TetherAcronym, b.Acronym)
		assert.Equal(t, TetherSymbol, b.Symbol)
		assert.Equal(t, TetherDecimalPlaces, b.DecimalPlaces)
	})
}
