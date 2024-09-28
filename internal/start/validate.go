package start

import "errors"

func ValidateInput(input string) (int, error) {
	if len(input) != 4 {
		return 0, errors.New("input lengh must be 4")
	}

	index, maxInd := 0, 8
	for i := 0; i < 4; i++ {
		if input[i] == byte('1') {
			index += maxInd
		} else if input[i] != byte('0') {
			return 0, errors.New("invalid number")
		}
		maxInd /= 2
	}

	return index, nil
}
