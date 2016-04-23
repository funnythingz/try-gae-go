package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func init() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/"), index)
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)
	mux.HandleFuncC(pat.Get("/unko"), unko)

	http.ListenAndServe(":8080", mux)
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
