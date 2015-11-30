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

func ExampleWaitingForCallbackFinished() {
	t := New(10*time.Millisecond, func(_ time.Time) {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("ticked")
	})

	time.Sleep(15 * time.Millisecond)
	t.Stop()
	fmt.Println("stopped")

	// Output:
	// ticked
	// stopped
}
