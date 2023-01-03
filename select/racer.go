package main

import (
	"errors"
	"net/http"
	"time"
)

func FetchURLsRace(firstURL string, secondURL string) (winner string, err error) {
	return fetchURLsTimeOutRace(firstURL, secondURL, 10*time.Second)
}

func fetchURLsTimeOutRace(firstURL string, secondURL string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(firstURL):
		return firstURL, nil
	case <-ping(secondURL):
		return secondURL, nil
	case <-time.After(timeout):
		return "", errors.New("requests take too long")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		ch <- struct{}{}
		close(ch)
	}()

	return ch
}

func measureResponseTime(firstURL string) time.Duration {
	firstURLStartTime := time.Now()
	http.Get(firstURL)
	return time.Since(firstURLStartTime)
}
