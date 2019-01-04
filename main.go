package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world 3")
	})

	http.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		log.Println(w, "hello world 3")
	})

	http.ListenAndServe(":1234", nil)
}
