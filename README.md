![](https://pkg.go.dev/badge/github.com/hoshiimo-se/retry.svg)
![](https://img.shields.io/github/license/hoshiimo-se/retry)
![](https://coveralls.io/repos/github/hoshiimo-se/retry/badge.svg?branch=master)

# retry
Retries using the exponential backoff algorithm can be easily realized.

# Usage
```go
op := func() (interface{}, error) {
  fmt.Printf("some form of processing...（%v）\n", time.Now())
  return nil, errors.New("Error!!")
}
result, err := retry.Retry(op)
```

# License
MIT

# Author
hoshiimo
