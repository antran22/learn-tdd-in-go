package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeDelayTestServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayTestServer(20 * time.Millisecond)
	defer slowServer.Close()
	slowUrl := slowServer.URL

	fastServer := makeDelayTestServer(time.Millisecond)
	defer fastServer.Close()
	fastUrl := fastServer.URL

	want := fastUrl
	got, err := Racer(slowUrl, fastUrl)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestTimeoutRacer(t *testing.T) {
	serverA := makeDelayTestServer(10 * time.Millisecond)
	defer serverA.Close()
	serverB := makeDelayTestServer(12 * time.Millisecond)
	defer serverB.Close()

	_, err := ConfigurableRacer(serverA.URL, serverB.URL, 5*time.Millisecond)

	assert.Error(t, err)
}
