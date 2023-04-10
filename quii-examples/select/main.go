package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecTimeout = 10 * time.Millisecond

func Racer(a,b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecTimeout)
}

func ConfigurableRacer(a,b string, timeout time.Duration) (string, error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b,nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

/*
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/