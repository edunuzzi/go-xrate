package util

import (
	"strconv"
)

func StringToFloat32(str string) (float32, error) {
	value, err := strconv.ParseFloat(str, 32)

	if err != nil {
		return 0, err
	}

	return float32(value), nil
}
