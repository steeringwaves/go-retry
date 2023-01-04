package retry

import (
	"fmt"

	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type failing struct {
	count int
}

func (f *failing) fail() error {
	f.count++
	return fmt.Errorf("failed on attempt %d", f.count)
}

func TestDo(t *testing.T) {
	t.Run("Do", func(t *testing.T) {
		var f failing

		err := Do(3, 1*time.Microsecond, func() error {
			failErr := f.fail()
			if failErr != nil {
				return failErr
			}

			return nil
		})

		assert.Equal(t, errors.New("failed on attempt 3"), err)
	})
}
