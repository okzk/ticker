# ticker

A thin wrapper for Go’s ticker.

## Installation

    go get github.com/okzk/ticker

## Example

```go
package main

import (
	"fmt"
	"github/okzk/ticker"
	"time"
)

func main() {
	ticker := ticker.New(10*time.Millisecond, func(t time.Time) {
		fmt.Println("tick at", t)
	})

	time.Sleep(35 * time.Millisecond)
	ticker.Stop()
}
```

## License

MIT
