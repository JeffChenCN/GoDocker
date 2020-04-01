package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello Gopher from Docker")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go running in Docker Container")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
