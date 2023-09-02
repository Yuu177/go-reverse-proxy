package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	helloProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8888",
	})

	worldProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:7777",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloProxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		worldProxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
