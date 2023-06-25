[![](https://pkg.go.dev/badge/github.com/hoshiimo-se/retry)](https://pkg.go.dev/github.com/hoshiimo-se/retry)
[![](https://img.shields.io/github/license/hoshiimo-se/retry)](https://github.com/hoshiimo-se/retry/blob/master/license)
[![](https://img.shields.io/github/languages/code-size/hoshiimo-se/retry)](https://github.com/hoshiimo-se/retry)
[![](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Ftwitter.com%2Fhoshiimo_se)](https://twitter.com/hoshiimo_se)

# retry
Retries using the exponential backoff algorithm can be easily realized.

```mermaid
%%{init:{'theme':'neutral'}}%%
sequenceDiagram
    participant User
    participant Retry Function
    participant Operation

    User->>Retry Function: Retry(op, opts)
    loop Until success or max retries
        Retry Function->>Operation: Execute operation (op)
        Operation-->>Retry Function: If error
        Retry Function->>Retry Function: Sleep for delay
    end
    Retry Function->>User: Return result or error
```

# Usage
## No return value
```go
op := func() error {
    return errors.New("Error!!")
}

err := retry.Retry(op, retry.WithInitialDelay(2*time.Second), retry.WithMaxRetries(3))
```

## One return value
```go
op := func() (string, error) {
    return "", errors.New("Error!!")
}

result, err := retry.RetryOneResult(op, retry.WithInitialDelay(2*time.Second), retry.WithMaxRetries(3))
```

## Two return values
```go
	op := func() (string, bool, error) {
		return "", false, errors.New("Error!!")
	}

	result1, result2, err := retry.RetryTwoResult(op, retry.WithInitialDelay(2*time.Second), retry.WithMaxRetries(3))
```

# License
MIT

# Author
hoshiimo
