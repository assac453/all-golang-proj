package main

import "net/http"

func webhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
}

func main() {

}
