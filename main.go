package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		port = flag.String("port", env("PORT", "8080"), "The port")
	)
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("You must specify the `server` or `worker` subcommand.")
	}

	cmd := os.Args[1]
	switch cmd {
	case "server":
		log.Printf("Starting on %s", *port)
		log.Fatal(http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s - %s", r.Method, r.URL)
			w.WriteHeader(200)
			io.WriteString(w, "Ok\n")
		})))
	case "worker":
		for {
			<-time.After(1 * time.Second)
			fmt.Printf("Hard work %d...\n", rand.Int())
		}
	default:
		log.Fatalf("Unkown subcommand: %s", cmd)
	}
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
