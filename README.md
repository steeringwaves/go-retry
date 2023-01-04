# go-retry

![workflow](https://github.com/github/docs/actions/workflows/test.yml/badge.svg)

## usage

```go
import "github.com/steeringwaves/go-retry"

count := 0
err := retry.Do(3, 250 * time.Microsecond, func() error {
	count++
	if count < 3 {
		return errors.New("still failing")
	}

	return nil
})

if err != nil {
	fmt.Printf("should always be nil in this case")
}
```
