package sync

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("test sync counter", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})

	t.Run("test async counter", func(t *testing.T) {
		counter := NewCounter()

		expectedCount := 1234

		wg := sync.WaitGroup{}

		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assert.Equal(t, expectedCount, counter.Value())
	})
}
