package retry

import (
	"time"

	"github.com/go-zoox/safe"
)

// Retry trys to run fn with times and interval
func Retry(fn func() error, times int, interval time.Duration) error {
	safeFn := safe.Func(fn)

	var err error
	for i := 0; i < times; i++ {
		if err = safeFn(); err == nil {
			return nil
		}

		time.Sleep(interval)
	}

	return err
}
