package timeout

import (
	"errors"
	"fmt"
	"time"
)

// Timeout call the function with timeout
func Timeout(fn func() error, timeout time.Duration, message ...string) (err error) {
	messageX := "timeout"
	if len(message) != 0 {
		messageX = message[0]
	}

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

		err = fn()
		c <- struct{}{}
	}()

	select {
	case <-c:
		return nil
	case <-time.After(timeout):
		return errors.New(messageX)
	}
}
