package main

import (
	"log"
	"net/http"
)

func rootController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("<h1> Welcome to my server </h1>"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootController)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
