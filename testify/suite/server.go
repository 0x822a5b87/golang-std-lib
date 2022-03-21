package suite

import (
	"fmt"
	"log"
	"net/http"
)

var server *http.Server

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome, %s", r.URL.Query().Get("name"))
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/greeting", greeting)

	server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
