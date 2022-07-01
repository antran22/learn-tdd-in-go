package _select

import (
	"fmt"
	"net/http"
	"time"
)

func measureRequestDuration(url string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	duration := time.Since(start)
	return duration, nil
}

func makeRequest(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-makeRequest(a):
		return a, nil
	case <-makeRequest(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("request timeout")
	}
}
