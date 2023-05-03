# retry
Retries using the exponential backoff algorithm can be easily realized.

# Usage
```go
result, err := retry.Retry(func() (interface{}, error) {
  fmt.Printf("some form of processing...（%v）\n", time.Now())
  return nil, errors.New("Error!!")
})
if err != nil {
  fmt.Println("Fail:", err)
} else {
  fmt.Println("Succeed:", result)
}
```

# License
MIT

# Author
hoshiimo
