package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpySleeperWriter struct {
	operations []string
}

func (s *SpySleeperWriter) Sleep() {
	s.operations = append(s.operations, sleep)
}

func (s *SpySleeperWriter) Write(p []byte) (n int, err error) {
	s.operations = append(s.operations, write)
	return 0, nil
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("test output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeperWriter{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		assert.Equal(t, want, got)

	})

	t.Run("test order of operation", func(t *testing.T) {
		spySleeperWriter := &SpySleeperWriter{}
		Countdown(spySleeperWriter, spySleeperWriter)

		expectedOperations := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assert.Equal(t, expectedOperations, spySleeperWriter.operations)
	})

}

type SpySleep struct {
	durationSlept time.Duration
}

func (s *SpySleep) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spySleep := &SpySleep{}
	sleeper := ConfigurableSleeper{sleepTime, spySleep.Sleep}

	sleeper.Sleep()

	assert.Equal(t, sleepTime, spySleep.durationSlept)
}
