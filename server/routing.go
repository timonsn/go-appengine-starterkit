package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

func ExampleHandler(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "starterkit, you've hit %s\n", r.URL.Path)
	fmt.Fprintf(w, "context %s\n", c.Value("test"))
}

func AddContext(ctx context.Context, next func(context.Context, http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(ctx, w, r)
	})
}

func SetupRouting(ctx context.Context) http.Handler {
	ctx = context.WithValue(ctx, "test", "Test Value")
	mux := http.NewServeMux()
	mux.Handle("/", AddContext(ctx, ExampleHandler))

	return mux
}
