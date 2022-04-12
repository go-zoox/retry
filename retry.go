package retry

import (
	"errors"
	"time"
)

// Retry trys to run fn with times and interval
func Retry(fn func(), times int, interval time.Duration) error {
	var err error
	for i := 0; i < times; i++ {
		func() {
			defer func() {
				if re := recover(); re != nil {
					switch v := re.(type) {
					case error:
						err = v
					case string:
						err = errors.New(v)
					default:
						err = errors.New("unexpected error")
					}
				}
			}()

			// reset error before run fn
			err = nil
			fn()
		}()

		if err == nil {
			return nil
		}

		time.Sleep(interval)
	}

	return err
}
