package retry

import (
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	// Initialize variables for testing
	var counter int
	alwaysFail := func() (interface{}, error) {
		return nil, errors.New("Always fail")
	}

	succeedOnThirdTry := func() (interface{}, error) {
		counter++
		if counter == 3 {
			return "Success", nil
		}
		return nil, errors.New("Failed")
	}

	// Test when operation always fails
	t.Run("AlwaysFail", func(t *testing.T) {
		_, err := Retry(alwaysFail, WithMaxRetries(3), WithoutJitter())
		if err == nil {
			t.Error("Expected error, but got no error")
		}
	})

	// Test when operation succeeds on the third try
	t.Run("SucceedOnThirdTry", func(t *testing.T) {
		counter = 0
		result, err := Retry(succeedOnThirdTry, WithMaxRetries(3), WithoutJitter())
		if err != nil {
			t.Errorf("Expected no error, but got error: %v", err)
		}
		if result != "Success" {
			t.Errorf("Expected result 'Success', but got '%v'", result)
		}
	})

	// Test when operation fails due to max retries
	t.Run("FailDueToMaxRetries", func(t *testing.T) {
		counter = 0
		_, err := Retry(succeedOnThirdTry, WithMaxRetries(2), WithoutJitter())
		if err == nil || err.Error() != "Maximum number of retries reached" {
			t.Errorf("Expected error 'Maximum number of retries reached', but got: %v", err)
		}
	})

	// Test custom initial delay and max delay
	t.Run("CustomDelays", func(t *testing.T) {
		counter = 0
		start := time.Now()
		_, err := Retry(succeedOnThirdTry, WithMaxRetries(3), WithInitialDelay(500*time.Millisecond), WithMaxDelay(2*time.Second), WithoutJitter())
		elapsed := time.Since(start)

		if err != nil {
			t.Errorf("Expected no error, but got error: %v", err)
		}
		if int64(elapsed) < int64(2*time.Second) || int64(elapsed) >= int64(4*time.Second) {
			t.Errorf("Expected elapsed time to be between 2 and 4 seconds, but got: %v", elapsed)
		}
	})
}
