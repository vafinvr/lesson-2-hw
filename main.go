package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	listenAddr = ":8000"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "OK"}`)
}

func main() {
	http.HandleFunc("/health/", health)

	log.Println("Listening on port ", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
