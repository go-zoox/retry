package retry

import (
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	c := 0
	if err := Retry(
		func() error {
			c++
			if c < 3 {
				panic("panic error")
			}

			return nil
		},
		1,
		time.Millisecond,
	); err == nil {
		t.Error("should be error retry 1 times")
	}
	if c != 1 {
		t.Error("should be 1 times")
	}

	c = 0
	if err := Retry(
		func() error {
			c++
			if c < 3 {
				panic("panic error")
			}

			return nil
		},
		2,
		time.Millisecond,
	); err == nil {
		t.Error("should be error retry 2 times")
	}
	if c != 2 {
		t.Error("should be 2 times")
	}

	c = 0
	if err := Retry(
		func() error {
			c++
			if c < 3 {
				panic("panic error")
			}

			return nil
		},
		3,
		time.Millisecond,
	); err != nil {
		t.Error("should be success retry 3 times")
	}
	if c != 3 {
		t.Error("should be 3 times")
	}

	c = 0
	if err := Retry(
		func() error {
			c++
			if c < 3 {
				panic("panic error")
			}

			return nil
		},
		4,
		time.Millisecond,
	); err != nil {
		t.Error("should be success retry 4 times")
	}
	if c != 3 {
		t.Error("should be 3 times")
	}
}
