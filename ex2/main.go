package ex2

import (
	"fmt"
	"net/http"
)

func restHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes := []byte(`{"a": 1, "b": 2, "c": 3}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	//json.NewEncoder(w).Encode(jsonBytes)
}

func main() {
	http.HandleFunc("/demo", restHandler)
	fmt.Println("http server running listening & serving on '0.0.0.0:9999'")
	http.ListenAndServe("0.0.0.0:9999", nil)
	fmt.Println("Terminating the http server...")
}
