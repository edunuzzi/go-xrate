package util

import (
	"fmt"
	"strconv"
)

func StringToFloat32(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return float32(value)
}