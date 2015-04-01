package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		port = flag.String("port", env("PORT", "8080"), "The port")
	)
	flag.Parse()

	log.Printf("Starting on %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s", r.Method, r.URL)
		w.WriteHeader(200)
		io.WriteString(w, "Ok\n")
	})))
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
