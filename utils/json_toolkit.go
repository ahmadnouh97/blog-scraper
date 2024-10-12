package utils

import (
	"encoding/json"
	"fmt"
)

func ToJSON[T any](data T, pretty bool) (string, error) {
	var jsonData []byte
	var err error

	if pretty {
		jsonData, err = json.MarshalIndent(data, "", "  ")
	} else {
		jsonData, err = json.Marshal(data)
	}

	if err != nil {
		return "", fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	return string(jsonData), nil
}
