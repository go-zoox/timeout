package retry

import (
	"errors"
	"fmt"
	"time"
)

// Timeout call the function with timeout
func Timeout(fn func(), timeout time.Duration) (err error) {
	c := make(chan struct{})

	go func() {
		defer func() {
			if e := recover(); e != nil {
				switch v := e.(type) {
				case error:
					err = v
				case string:
					err = errors.New(v)
				default:
					err = fmt.Errorf("%v", v)
				}
			}
		}()

		fn()
		c <- struct{}{}
	}()

	select {
	case <-c:
		return nil
	case <-time.After(timeout):
		return errors.New("timeout")
	}
}
