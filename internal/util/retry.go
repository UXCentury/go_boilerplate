package util

import (
	"fmt"
	"time"
)

// Retry is a custom retry function with exponential backoff and a maximum number of attempts
func Retry(fn func() error) error {
	sugar := Logger.Sugar()
	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}

		sugar.Errorln("Error (attempt %d): %v", attempt, err)

		// Calculate the next retry interval using exponential backoff
		backoff := time.Duration(attempt * attempt)
		time.Sleep(backoff * time.Second)
	}

	return fmt.Errorf("exhausted all attempts")
}
