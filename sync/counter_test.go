package sync_go

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("incrementing concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, want)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d want %d", counter.Value(), want)
	}
}
