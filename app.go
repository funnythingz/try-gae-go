package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)
	goji.Get("/", index)
	goji.Get("/hello/:name", hello)
	goji.Get("/unko", unko)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func unko(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "unko")
}
