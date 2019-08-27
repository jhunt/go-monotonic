package monotonic_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/jhunt/go-monotonic"
)

func TestJustAMassiveNumberOfGoRoutines(t *testing.T) {
	var wg sync.WaitGroup
	c := monotonic.NewClock(0)

	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(c *monotonic.Clock) {
			nap := rand.Intn(8) * 100
			time.Sleep(time.Duration(nap) * time.Millisecond)
			c.Increment()
			wg.Done()
		}(c)
	}

	wg.Wait()
	if c.Value() != int64(n) {
		t.Errorf("spun %d goroutines, but clock only incremented to %d", n, c.Value())
	}
}
