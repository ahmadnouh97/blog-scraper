package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveJSON[T any](data T, fileName string) error {
	// Marshal the struct into JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	// Create or overwrite the file with the provided file name
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func LoadJSON[T any](fileName string) (T, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return *new(T), fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	var data T
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return *new(T), fmt.Errorf("failed to decode JSON: %v", err)
	}

	return data, nil
}
