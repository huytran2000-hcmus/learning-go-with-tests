package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	query     string
	cancelled bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	select {
	case result := <-Process(ctx, s.query):
		return result, nil
	case <-ctx.Done():
		s.Cancel()
		return "", ctx.Err()
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func Process(ctx context.Context, input string) chan string {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, v := range input {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(v)
			}
		}
		data <- result
	}()

	return data
}

type spyResponseWriter struct {
	written bool
}

func (s *spyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *spyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *spyResponseWriter) WriteHeader(int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", testDataReturn)
	t.Run("store cancel work if request is cancelled", testStoreCancel)
}

func testDataReturn(t *testing.T) {
	want := "Hello world!"
	store := &SpyStore{
		query: want,
	}
	server := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	got := response.Body.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if store.cancelled {
		t.Errorf("expected store is not cancelled")
	}
}

func testStoreCancel(t *testing.T) {
	want := "Hello World!"
	store := &SpyStore{
		query:     want,
		cancelled: false,
	}
	server := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(50*time.Millisecond, cancel)
	request = request.WithContext(ctx)

	response := &spyResponseWriter{}

	server.ServeHTTP(response, request)

	if !store.cancelled {
		t.Errorf("Expected store to cancel but it was not")
	}

	if response.written {
		t.Error("a response should not have been written")
	}
}
