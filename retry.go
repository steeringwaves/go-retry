package retry

import (
	"context"
	"time"

	timer "github.com/steeringwaves/go-timer"
)

type stop struct {
	error
}

type Options struct {
	Context  context.Context
	Delay    time.Duration
	Backoff  time.Duration
	Attempts int
}

// DoWithOptions will retry a function until it returns nil.
// This will return the last error if the function does not return nil.
// This will cease if options Context is cancelled or the Attempts exceed the specified value.
func DoWithOptions(opts Options, fn func() error) error {
	if nil == opts.Context {
		opts.Context = context.Background()
	}

	var err error

	if err = opts.Context.Err(); err != nil {
		return err
	}

	if err = fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if opts.Attempts < 0 { //infinite
			t := timer.NewTimer(opts.Delay + opts.Backoff)
			defer t.Stop()

			select {
			case <-t.C:
			case <-opts.Context.Done():
				return opts.Context.Err()
			}

			return DoWithOptions(opts, fn)
		}

		if opts.Attempts--; opts.Attempts > 0 {
			t := timer.NewTimer(opts.Delay + opts.Backoff)
			defer t.Stop()

			select {
			case <-t.C:
			case <-opts.Context.Done():
				return opts.Context.Err()
			}

			// double backoff after each attempt
			opts.Backoff = opts.Backoff * 2
			return DoWithOptions(opts, fn)
		}
		return err
	}
	return nil
}

// Do provides a wrapper for DoWithOptions that only exposes the number of attempts and delay
func Do(attempts int, delay time.Duration, fn func() error) error {
	return DoWithOptions(Options{Attempts: attempts, Delay: delay}, fn)
}

// DoWithContext provides a wrapper for DoWithOptions that only exposes the context, number of attempts and delay
func DoWithContext(ctx context.Context, attempts int, delay time.Duration, fn func() error) error {
	return DoWithOptions(Options{Context: ctx, Attempts: attempts, Delay: delay}, fn)
}
