package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("get fastest one", func(t *testing.T) {
		t.Parallel()
		slowServer := makeDelayedTestServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedTestServer(0)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := FetchURLsRace(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q\n", got, want)
		}
	})

	t.Run("responses take too long", func(t *testing.T) {
		t.Parallel()
		slowServer := makeDelayedTestServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedTestServer(30 * time.Millisecond)
		defer fastServer.Close()

		_, err := fetchURLsTimeOutRace(slowServer.URL, fastServer.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedTestServer(delayedTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delayedTime)
		writer.WriteHeader(http.StatusOK)
	}))
}
