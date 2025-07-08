package main

import (
	"errors"
	"log/slog"

	"google.golang.org/genai"
)

func main() {
	// Example of using errors.As
	err := APIError()
	if err != nil {
		var apiErr genai.APIError
		if errors.As(err, &apiErr) {
			slog.Error("API error occurred", "error", apiErr)
		}
	}
}

func APIError() error {
	return genai.APIError{
		Message: "An example API error occurred",
		Code:    400,
	}
}
