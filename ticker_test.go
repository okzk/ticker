package ticker

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestNumGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		New(time.Millisecond, func(_ time.Time) {}).Stop()
	}

	time.Sleep(10 * time.Millisecond)

	if runtime.NumGoroutine() >= 1000 {
		t.Fatal("goroutine leaked!!!")
	}
}

func ExampleTicker() {
	t := New(10*time.Millisecond, func(_ time.Time) {
		fmt.Println("ticked")
	})

	time.Sleep(35 * time.Millisecond)
	t.Stop()
	time.Sleep(15 * time.Millisecond)

	// Output:
	// ticked
	// ticked
	// ticked
}
