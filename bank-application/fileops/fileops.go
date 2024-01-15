package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatValueToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatValueFromFile(fileName string) (float64, error) {
	data, error := os.ReadFile(fileName)
	if error != nil {
		//panic(error)
		error := fmt.Sprintf("Error while reading the file %v", fileName)
		return 1000.0, errors.New(error)
	}
	floatText := string(data)
	floatValue, err := strconv.ParseFloat(floatText, 64)
	if err != nil {
		//panic(err)
		return 1000.0, errors.New("failed to parse stored balance value")
	}
	return floatValue, nil
}
