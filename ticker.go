package ticker

import (
	"time"
)

type Ticker struct {
	t *time.Ticker
	c chan bool
}

func NewTicker(duration time.Duration, callback func(time.Time)) *Ticker {
	t := time.NewTicker(duration)
	c := make(chan bool)
	go func() {
		for {
			select {
			case arg := <-t.C:
				callback(arg)
			case <-c:
				return
			}
		}
	}()

	return &Ticker{t: t, c: c}
}

func (t *Ticker) Stop() {
	t.t.Stop()
	close(t.c)
}
