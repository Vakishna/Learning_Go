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
type Book struct {
  ID  string 'json:"id"'
  Isbn  string 'json:"isbn"'
  Title  string 'json:"title"'
  // The * implies own type (Author) **see struct below**
  Author  *Author 'json:"author"'
}


// Author Struct (Model)
type Author struct {
  FirstName string 'json:"firstname"'
  LastName  string 'json:"lastname"'
}


// Get ALL Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}


func main() {

  // Initialize the Mux Router
  // := is the type inference for GoLang
  r := mux.NewRouter()

  // Route Handlers / Endpoints
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/book", createBook).Methods("POST")
  r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")


  // Run the server
  http.ListenAndServe(":8080", r)






}
