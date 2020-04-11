package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello Gopher from Docker")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go running in Docker Container")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Gopher Docker %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
