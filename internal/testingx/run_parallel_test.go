package testingx_test

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/joeycorry/ds/internal/syncx"
	"github.com/joeycorry/ds/internal/testingx"
)

func TestRunParallel(t *testing.T) {
	t.Parallel()

	t.Run("runs subtests in parallel", func(t *testing.T) {
		t.Parallel()

		const minParallelism = 2

		parallelism, err := strconv.Atoi(flag.Lookup("test.parallel").Value.String())
		if err != nil {
			t.Fatalf("requires -parallel argument to be an integer")
		}

		if parallelism < minParallelism {
			t.Skipf("requires -parallel >= %d, got %d", minParallelism, parallelism)
		}

		var (
			mu      sync.Mutex
			arrived int
		)
		release := make(chan struct{})

		for i := range parallelism {
			testingx.RunParallel(t, fmt.Sprintf("subtest #%d", i), func(t *testing.T) {
				syncx.Locked(&mu, func() {
					arrived++
					if arrived == parallelism {
						close(release)
					}
				})

				select {
				case <-release:
				case <-time.After(500 * time.Millisecond):
					t.Fatal("subtests did not run in parallel (timed out waiting for siblings)")
				}
			})
		}
	})
}
