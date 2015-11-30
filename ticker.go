package ticker

import (
	"sync"
	"time"
)

type Ticker struct {
	t  *time.Ticker
	c  chan bool
	wg sync.WaitGroup
}

func New(duration time.Duration, callback func(time.Time)) *Ticker {
	ticker := &Ticker{
		t: time.NewTicker(duration),
		c: make(chan bool),
	}

	ticker.wg.Add(1)
	go func() {
		defer ticker.wg.Done()
		for {
			select {
			case t := <-ticker.t.C:
				callback(t)
			case <-ticker.c:
				return
			}
		}
	}()
	return ticker
}

func (t *Ticker) Stop() {
	t.t.Stop()
	close(t.c)
	t.wg.Wait()
}
