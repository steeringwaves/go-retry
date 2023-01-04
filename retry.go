package retry

// inspired by https://upgear.io/blog/simple-golang-retry-function/

import (
	"context"
	"time"
)

type stop struct {
	error
}

func Do(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts < 0 { //infinite
			time.Sleep(sleep)
			return Do(attempts, sleep, fn)
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Do(attempts, sleep, fn)
		}
		return err
	}
	return nil
}

func DoWithContext(ctx context.Context, attempts int, sleep time.Duration, fn func() error) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts < 0 { //infinite
			time.Sleep(sleep)
			return DoWithContext(ctx, attempts, sleep, fn)
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return DoWithContext(ctx, attempts, sleep, fn)
		}
		return err
	}
	return nil
}

//TODO add backoff, jitter functions
