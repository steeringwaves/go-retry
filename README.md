# go-retry

![workflow](https://github.com/steeringwaves/go-retry/actions/workflows/test.yml/badge.svg)

Retry is a package that aims to make it simpler to retry a function. Inspired by https://upgear.io/blog/simple-golang-retry-function/

## roadmap

- [x] Specify a maximum number of retries (0 means infinite)
- [x] Specify a context that is obeyed
- [x] Specify a delay between retries
- [x] Specify a backoff amount that is added after each failure (this value is added to the specified delay time)
- [ ] Specify a jitter amount to randomize delay

The full API documentation is available here: http://godoc.org/github.com/steeringwaves/go-retry.

## usage

```go
import (
	"time"
	"context"
	"github.com/steeringwaves/go-retry"
)

ctx := context.TODO()

count := 0
err := retry.DoWithOptions(retry.Options{
	Context: ctx,
	Attempt: 3,
	Delay: 50 * time.Microsecond,
	Backoff: 100 * time.Microsecond,
	func() error {
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

### helper functions

```go
import (
	"time"
	"github.com/steeringwaves/go-retry"
)

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
