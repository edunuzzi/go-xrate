package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringToFloat32(t *testing.T) {
	base := func(str string, expectedReturn float32, shouldErr bool) {
		f, err := StringToFloat32(str)

		if shouldErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, expectedReturn, f)
	}

	base("102", 102, false)
	base("102.3", 102.3, false)
	base("102476.32767", 102476.32767, false)
	base("-102476.32767", -102476.32767, false)
	base("wrong", 0, true)
}