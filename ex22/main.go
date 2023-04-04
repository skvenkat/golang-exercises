// Source : https://medium.com/@cerebrovinny/golang-goroutines-powering-high-performance-applications-767742d961ce

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	reqsPerSecond = 1000
)

// processRequest simulates a complex workflow by sleeping for 10ms
func processRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Millisecond)
	io.WriteString(w, "Request Processed successfully!\n")
}

// rateLimitHandler limits the request rate to the specified requests per second
func rateLimitHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	ticker := time.NewTicker(time.Second/reqsPerSecond)
	throttle := make(chan time.Time, reqsPerSecond)
	
	go func() {
		for t := range ticker.C {
			throttle <- t
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		<-throttle
		handlerFunc(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/process", rateLimitHandler(processRequest))

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	fmt.Println("Starting server on 0.0.0.0:8080")
	log.Fatal(server.ListenAndServe())
}
