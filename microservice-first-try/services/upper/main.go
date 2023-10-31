package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upper", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from upper")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
