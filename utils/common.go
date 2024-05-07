package utils

import (
	"context"
	"io"
	"log"
	"net/http"
)

// GetJSON - ...
func GetJSON(url string) []byte {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	FailOnError(err, "Failed to create new request")

	res, err := http.DefaultClient.Do(req)
	FailOnError(err, "Failed to get response")

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	FailOnError(err, "Failed to read response body")

	return body
}

// FailOnError - ...
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Error: %s - Message: %s", err, msg)
	}
}

// SetDefaultValue - ...
func SetDefaultValue(src *string, dst string) {
	*src = dst
}
