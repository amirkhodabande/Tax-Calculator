package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

func ReadFloatFromFile(path string) ([]float64, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []float64

	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			file.Close()
			return nil, err
		}

		lines = append(lines, floatPrice)
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, err
	}

	return lines, nil
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode data")
	}

	file.Close()
	return nil
}
