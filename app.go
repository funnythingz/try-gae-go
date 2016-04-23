package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func init() {
	mux := goji.NewMux()

	mux.HandleFuncC(pat.Get("/"), index)
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)
	mux.HandleFuncC(pat.Get("/article"), article)

	http.Handle("/", mux)
}

func index(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}

func hello(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", pat.Param(c, "name"))
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func article(c context.Context, w http.ResponseWriter, r *http.Request) {
	article := Article{
		Title:   "title",
		Content: "hello world",
	}
	response, _ := json.Marshal(article)
	io.WriteString(w, string(response))
}
