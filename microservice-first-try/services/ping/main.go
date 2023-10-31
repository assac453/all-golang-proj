package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi"
)

var schema = `
	CREATE TABLE ping(
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		time TEXT NOT NULL
	);
`
var dbname = "ping.db"
var db *sql.DB

func init() {
	os.Remove(dbname)
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatalf("error try open conn %v", err)
		os.Exit(-1)
	}
	db.Exec(schema)
	db.Close()
}

func main() {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		var pong = struct {
			TimeNow string `json:"time"`
		}{
			TimeNow: time.Now().String(),
		}

		_, err := db.Exec("INSERT INTO ping(time) VALUE($1)", pong.TimeNow)
		if err != nil {
			fmt.Fprintf(w, "error occure %v", err)
			return
		}

		log.Printf(pong.TimeNow)

		json.NewEncoder(w).Encode(&pong)
	})

	s := &http.Server{
		Handler:      r,
		Addr:         ":8081",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
