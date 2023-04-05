package main

import (
	_ "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitHandler(t *testing.T) {
	// Create new request with /process endpoint
	request, err := http.NewRequest("GET", "/process", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to record the response
	recorder := httptest.NewRecorder()

	// Define a new handler function that sleeps for 100ms to simulate the complex workflow
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		time.Sleep(100 * time.Millisecond)
		respTime := time.Since(start)
		w.Header().Set("X-Response-Time", respTime.String())
		w.Write([]byte("Request processed Successfully!\n"))
	})

	// Wrap the handlerFunc with the rateLimitHandler
	rateLimitHandler := rateLimitHandler(handlerFunc)

	// Execute the rateLimitHandler for x times
	for i := 0; i < 10; i++ {
		rateLimitHandler(recorder, request)
	}

	// Verify the rate is limited to 1000 requests per second
	responseTime := recorder.Result().Header.Get("X-Response-Time")
	if responseTime != "100ms" {
		t.Errorf("Unexpected response time: %s", responseTime)
		t.Errorf("All headers : %v", recorder.Result().Header)
	}
}

func TestProcessRequest(t *testing.T) {
	// Create a new request with /process endpoint
	request, err := http.NewRequest("GET", "/process", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to save the response
	recorder := httptest.NewRecorder()

	// Call the processRequest function with recorder and request
	processRequest(recorder, request)

	// Verify that the response status code is 200 OK
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
	}

	// Verify that the response body contains the expected message
	expected := "Request Processed Successfully!\n"
	if response := recorder.Body.String(); response != expected {
		t.Errorf("Unexpected response body: got %v, want %v", response, expected)
	}
}
