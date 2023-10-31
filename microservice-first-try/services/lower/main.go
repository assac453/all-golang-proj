package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/lower", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from lower")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
