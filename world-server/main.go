package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "World, From 7777")
	})

	log.Fatal(http.ListenAndServe(":7777", nil))
}
