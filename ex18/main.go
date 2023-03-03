package main

import (
    "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Service struct {
    logger chan string
}

type Request struct {
    Number int `json:"number"`
}

func main() {
    ctx := context.Background()
	svc := NewService(ctx)
	r := http.NewServeMux()
	r.HandleFunc("/", svc.Handler)
	err := http.ListenAndServe("0.0.0.0:9010", r)
	if err != nil {
	    log.Fatal(err)
	}
}

func NewService(ctx context.Context) *Service {
	svc := &Service{
		// buffered channell for async logging
		logger: make(chan string, 10)
	}
	
	// fire logInfo func to run in background
	go logInfo(ctx, svc.logger)
	
	return svc
}

func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {
    var r Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
	    // if any error log and proceed to response without any delay
		s.logger <- fmt.Sprintf("error occurred %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	
	s.logger <- fmt.Sprintf("received %d", req.Number)
	w.WriteHeader(http.StatusOk)
}

// async logger
func logInfo(ctx context.Context, logger chan string) {
    for {
	    select {
		    case <- ctx.Done():
			    return
			case x:= <- logger:
			    log.Println(x)
		}
	}
}
