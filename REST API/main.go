package main

import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

// Book Struct (Model)






func main() {

  // Initialize the Mux Router
  // := is the type inference for GoLang
  r := mux.NewRouter()

  // Route Handlers / Endpoints
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")


  // Run the server
  http.ListenAndServe(":8080", r)






}
