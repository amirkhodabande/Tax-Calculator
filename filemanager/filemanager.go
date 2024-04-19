package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadFloatFromFile() ([]float64, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []float64

	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		lines = append(lines, floatPrice)
	}

	err = scanner.Err()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to encode data")
	}

	return nil
}
