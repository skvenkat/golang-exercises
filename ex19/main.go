package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
	"sync"
)

type Request struct {
 Name    string `json:"name"`
 Id      string `json:"id"`
 Country string `json:"country"`
 Age     int    `json:"age"`
}

type ObjectPool struct {
    ReqPool *sync.Pool
}

var objPool ObjectPool

func init() {
    // init object pool
    objPool.ReqPool = &sync.Pool{
        New: func() interface{} {
        return new(Request)
        },
    }
}

func HandlerWithoutSyncPool() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        req :=new(Request)
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
        err = json.NewEncoder(w).Encode(req)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    })
}

func HandlerWithSyncPool() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        req := objPool.ReqPool.Get().(*Request)
        
        //to make sure it return object to pool
        defer objPool.ReqPool.Put(req)
        
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
}

func main() {
    r := http.NewServeMux()

    r.Handle("/withpool", HandlerWithSyncPool())

    //Register pprof handlers
    r.HandleFunc("/debug/pprof", pprof.Index)
    r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    r.HandleFunc("/debug/pprof/profile", pprof.Profile)
    r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    r.HandleFunc("/debug/pprof/trace", pprof.Trace)

    err := http.ListenAndServe("0.0.0.0:9080", r)
    if err != nil {
        log.Fatal(err)
    }
}
