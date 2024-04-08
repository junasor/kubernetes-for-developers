package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// register getTime function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", getTime)

	// use PORT environment variable, or default to 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Printf("Quien cohones soy.")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "The time is  %s, UTC\n", time.Now().UTC().Format("3:04 PM"))
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
