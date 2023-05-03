package retry

import (
	"errors"
	"math/rand"
	"time"
)

type Operation func() (interface{}, error)

type RetryOptions struct {
	useJitter    bool
	initialDelay time.Duration
	maxDelay     time.Duration
	maxRetries   int
}

type RetryOption func(*RetryOptions)

func WithoutJitter() RetryOption {
	return func(options *RetryOptions) {
		options.useJitter = false
	}
}

func WithInitialDelay(delay time.Duration) RetryOption {
	return func(options *RetryOptions) {
		options.initialDelay = delay
	}
}

func WithMaxDelay(delay time.Duration) RetryOption {
	return func(options *RetryOptions) {
		options.maxDelay = delay
	}
}

func WithMaxRetries(maxRetries int) RetryOption {
	return func(options *RetryOptions) {
		options.maxRetries = maxRetries
	}
}

func Retry(op Operation, opts ...RetryOption) (interface{}, error) {
	const defaultInitialDelay = 100 * time.Millisecond
	const defaultMaxDelay = 10 * time.Second
	const defaultMaxRetries = 5

	options := &RetryOptions{
		useJitter:    true,
		initialDelay: defaultInitialDelay,
		maxDelay:     defaultMaxDelay,
		maxRetries:   defaultMaxRetries,
	}
	for _, opt := range opts {
		opt(options)
	}

	delay := options.initialDelay
	for i := 0; i < options.maxRetries; i++ {
		result, err := op()
		if err == nil {
			return result, nil
		}

		if options.useJitter {
			delay = time.Duration(float64(delay) * (1.5 + rand.Float64()))
		} else {
			delay = time.Duration(float64(delay) * 2)
		}
		if delay > options.maxDelay {
			delay = options.maxDelay
		}

		time.Sleep(delay)
	}

	return nil, errors.New("Maximum number of retries reached")
}
