package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

func ExampleHandler(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "starterkit, you've hit %s\n", r.URL.Path)
	fmt.Fprintf(w, "context %s\n", c.Value("test"))
}

func JsonHandler(c context.Context, w http.ResponseWriter, r *http.Request) {
	JSON(w, &struct {
		Hello string
		URLQuery url.Values
	}{
		Hello: "test",
		URLQuery: r.URL.Query(),
	})
}

func HtmlHandler(c context.Context, w http.ResponseWriter, r *http.Request) {
	HTML(w, "index.html", &struct {
		Title string
	}{
		Title: "Exampe Page",
	})
}

func AddContext(ctx context.Context, next func(context.Context, http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(ctx, w, r)
	})
}

func SetupRouting(ctx context.Context) http.Handler {
	ctx = context.WithValue(ctx, "test", "Test Value")
	mux := http.NewServeMux()
	mux.Handle("/", AddContext(ctx, HtmlHandler))
	mux.Handle("/json", AddContext(ctx, JsonHandler))

	return mux
}
