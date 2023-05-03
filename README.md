![](https://pkg.go.dev/badge/github.com/hoshiimo-se/retry.svg)
![](https://img.shields.io/github/license/hoshiimo-se/retry)
![](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Ftwitter.com%2Fhoshiimo_se)

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
