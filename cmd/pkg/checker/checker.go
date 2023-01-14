package checker

import (
	"strconv"
)

func CheckInputs(x, y string) (float64, float64, error) {
	firstVal, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, 0, err
	}
	secondVal, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, 0, err
	}
	return firstVal, secondVal, nil
}
