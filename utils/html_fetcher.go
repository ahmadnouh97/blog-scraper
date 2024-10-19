package utils

import (
	"fmt"
	"io"
	"net/http"
)

func FetchHTMLPage(url string) (string, error) {
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching the page: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Return the response body as a string
	return string(body), nil
}
