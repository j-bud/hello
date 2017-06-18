package main

import "net/http"

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/james", hellojames)
    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!"))
}

func hellojames(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello James!"))
}