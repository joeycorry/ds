package syncx_test

import (
	"sync"
	"testing"
	"time"

	"github.com/joeycorry/ds/internal/syncx"
	"github.com/joeycorry/ds/internal/testingx"
)

func TestLocked(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "runs the function while holding the lock", func(t *testing.T) {
		var mu sync.Mutex
		acquired := make(chan struct{})

		syncx.Locked(&mu, func() {
			go func() {
				mu.Lock()
				defer mu.Unlock()

				close(acquired)
			}()

			select {
			case <-acquired:
				t.Fatal("lock should be held while the function runs")
			case <-time.After(500 * time.Millisecond):
			}
		})

		select {
		case <-acquired:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("lock was not acquired by contender after Locked returned")
		}
	})

	testingx.RunParallel(t, "releases the lock after the function returns", func(t *testing.T) {
		var mu sync.Mutex
		acquired := make(chan struct{})

		syncx.Locked(&mu, func() {})

		go func() {
			mu.Lock()
			defer mu.Unlock()

			close(acquired)
		}()

		select {
		case <-acquired:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("lock was not released after Locked returned")
		}
	})
}
