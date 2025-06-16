package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var float []float64

	for _, stringValue := range strings {
		// Convert string to float64
		floatValue, err := strconv.ParseFloat(stringValue, 64)

		// Error handling for conversion
		if err != nil { 
			return nil, errors.New("failed to convert string to float64")
		}

		// Append the converted float64 value to the float slice
		float = append(float, floatValue)
	}

	return float, nil
}
