package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{}
		resp, err := client.Get("http://127.0.0.1:8081/ping")
		if err != nil {
			fmt.Fprintf(w, "error occure %v", err)
			return
		}
		var pong = struct {
			TimeNow string `json:"time"`
		}{}

		json.NewDecoder(resp.Body).Decode(&pong)

		log.Printf(pong.TimeNow)

		fmt.Fprintf(w, "%v", pong.TimeNow)
	})

	r.Post("/upper", func(w http.ResponseWriter, r *http.Request) {

		type data struct {
			Upper string `json:"upper"`
		}

		var first_layer_data_upr data
		json.NewDecoder(r.Body).Decode(&first_layer_data_upr)

		if first_layer_data_upr.Upper == "" {
			fmt.Fprintf(w, "no data")
		}

		client := http.Client{}
		resp, err := client.Post("http://127.0.0.1:8082/upper", "application/json", bytes.NewBuffer([]byte(first_layer_data_upr.Upper)))
		if err != nil {
			fmt.Fprintf(w, "error occure %v", err)
			return
		}
		var upr = data{}

		json.NewDecoder(resp.Body).Decode(&upr)

		log.Printf(upr.Upper)

		fmt.Fprintf(w, "%v", upr.Upper)

	})
	s := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
