package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		buf := make([]byte, 100)
		for {
			buf = buf[:cap(buf)]
			n, err := r.Body.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			buf = buf[:n]
		}
		fmt.Println(string(buf))
	})
	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	s.ListenAndServe()
}
