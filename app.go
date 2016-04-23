package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func init() {
	http.Handle("/", goji.DefaultMux)
	goji.Get("/", index)
	goji.Get("/hello/:name", hello)
	goji.Get("/unko", unko)
}

func index(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}

func hello(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", pat.Param(c, "name"))
}

func unko(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "unko")
}
