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

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("You must specify the `server` or `worker` subcommand.")
	}

	cmd := args[0]

	webString := env("WEB_STRING", "Default header")
	workerString := env("WORKER_STRING", "Default header")
	hostname, _ := os.Hostname()

	log.Printf("[%s] Starting %s process.\n", hostname, cmd)

	switch cmd {
	case "server":
		log.Printf("Starting on %s", *port)
		log.Fatal(http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			log.Printf("[%s] %s: %s - %s", hostname, webString, r.Method, r.URL)
			w.WriteHeader(200)
			io.WriteString(w, fmt.Sprintf("[%s %s] %s: Ok\n", now, hostname, webString))
		})))
	case "worker":
		for {
			<-time.After(1 * time.Second)
			log.Printf("[%s] %s: Hard work %d...\n", hostname, workerString, rand.Int())
		}
	default:
		log.Fatalf("Unknown subcommand: %s", cmd)
	}
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
