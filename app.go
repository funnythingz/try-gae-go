package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func main() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/"), index)
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)
	mux.HandleFuncC(pat.Get("/unko"), unko)

	http.ListenAndServe("localhost:8000", mux)
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
