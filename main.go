package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello my friend")
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server running")
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
