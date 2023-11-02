package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, this is the landing page of my personal website!")
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, this is a page listing my projects!")
}

func thoughtsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, this is a page for putting my thoughts on a page!")
}

func main() {
    http.HandleFunc("/", homePageHandler)
    http.HandleFunc("/projects", projectsHandler)
    http.HandleFunc("/thoughts", thoughtsHandler)
    log.Printf("Server starting on port: 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
