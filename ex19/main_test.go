package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	_ "net/http/httptest"
	"testing"
)

var tests = []struct {
	req Request
} {
	{
		req: Request{Name: "Abdul Kalam", Country: "India", Id: "0001", Age: 100,},
    },
}

func BenchmarkHandlerWithSyncPool(b *testing.B) {
	for _, test := range tests {
		data, _ := json.Marshal(test)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				http.Post("http://127.0.0.1:9080/withpool", "application/json", bytes.NewBuffer(data))
			}
		})
	}
}

func BenchmarkHandlerWithoutSyncPool(b *testing.B) {
	for _,test := range tests {
		data, _ := json.Marshal(test)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				http.Post("http://127.0.0.1:9080/withoutpool", "application/json", bytes.NewBuffer(data))
			}
		})
	}
}