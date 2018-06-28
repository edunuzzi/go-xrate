package util

import (
	"strconv"
)

func StringToFloat64(str string) (float64, error) {
	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, err
	}

	return value, nil
}
