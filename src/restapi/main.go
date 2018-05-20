package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Init books of slice of book
var books []Book

func main() {
	//Init new router
	router := mux.NewRouter()

	//Mock books data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "111111",
		Title: ".Net Core",
		Author: &Author{
			FirstName: "Casi",
			LastName:  "COMP",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "222222",
		Title: "Spring Boot",
		Author: &Author{
			FirstName: "Karim",
			LastName:  "SIDIBE",
		},
	})

	//Route handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

//Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get url params
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(Book{})
}
