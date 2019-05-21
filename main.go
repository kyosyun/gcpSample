package main

import (
    "net/http"
    "encoding/json"

    "google.golang.org/appengine"
)

type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/", handle)
    http.HandleFunc("/sample", handle2)
    appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Response{Status: "ok123", Message: "Hello world."})
}

func handle2(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Response{Status: "sample", Message: "Hello world."})
}

