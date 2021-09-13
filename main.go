package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("You must specify the `server` or `worker` subcommand.")
	}

	cmd := args[0]

	workerString := env("WORKER_STRING", "Default header")
	hostname, _ := os.Hostname()

	log.Printf("[%s] Starting %s process.\n", hostname, cmd)

	switch cmd {
	case "server":
		os.Exit(1)
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
