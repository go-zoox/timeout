package timeout

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	fn := func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("fn done")
		return nil
	}

	err := Timeout(fn, 50*time.Millisecond)
	if err == nil {
		t.Errorf("expected timeout error, got nil")
	}

	err = Timeout(fn, 300*time.Millisecond)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
