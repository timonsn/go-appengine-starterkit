package main

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
)

func init() {
	http.Handle("/", SetupRouting(context.Background()))
}

func main() {
	appengine.Main()
}
