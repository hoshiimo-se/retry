package retry

import (
	"errors"
	"testing"
)

// Mock operation for Retry
func mockOperation(isSuccess bool) Operation {
	return func() error {
		if isSuccess {
			return nil
		}
		return errors.New("mock error")
	}
}

// Mock operation for RetryOneResult
func mockOperationOneResult(isSuccess bool) OperationOneResult[int] {
	return func() (int, error) {
		if isSuccess {
			return 1, nil
		}
		return 0, errors.New("mock error")
	}
}

// Mock operation for RetryTwoResult
func mockOperationTwoResult(isSuccess bool) OperationTwoResult[int, string] {
	return func() (int, string, error) {
		if isSuccess {
			return 1, "success", nil
		}
		return 0, "", errors.New("mock error")
	}
}

func TestRetry(t *testing.T) {
	// Test successful operation
	err := Retry(mockOperation(true), WithMaxRetries(3))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test failing operation
	err = Retry(mockOperation(false), WithMaxRetries(3))
	if err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

func TestRetryOneResult(t *testing.T) {
	// Test successful operation
	result, err := RetryOneResult(mockOperationOneResult(true), WithMaxRetries(3))
	if err != nil || *result != 1 {
		t.Errorf("Expected 1, got %v with error %v", *result, err)
	}

	// Test failing operation
	result, err = RetryOneResult(mockOperationOneResult(false), WithMaxRetries(3))
	if err == nil || result != nil {
		t.Errorf("Expected error, got %v with result %v", err, result)
	}
}

func TestRetryTwoResult(t *testing.T) {
	// Test successful operation
	result1, result2, err := RetryTwoResult(mockOperationTwoResult(true), WithMaxRetries(3))
	if err != nil || *result1 != 1 || *result2 != "success" {
		t.Errorf("Expected 1 and success, got %v, %v with error %v", *result1, *result2, err)
	}

	// Test failing operation
	result1, result2, err = RetryTwoResult(mockOperationTwoResult(false), WithMaxRetries(3))
	if err == nil || result1 != nil || result2 != nil {
		t.Errorf("Expected error, got %v with result %v, %v", err, result1, result2)
	}
}
