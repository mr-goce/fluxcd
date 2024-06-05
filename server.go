package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Mock server responding to %s\n", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting mock server on :8080")
    http.ListenAndServe(":8080", nil)
}