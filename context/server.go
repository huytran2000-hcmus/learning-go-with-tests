package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		result, err := store.Fetch(ctx)
		if err != nil {
			// TODO: log error
			return
		}
		fmt.Fprint(w, result)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}
