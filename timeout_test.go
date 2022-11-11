package timeout

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	fn := func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("fn done")
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
