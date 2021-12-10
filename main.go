package main

import (
	"book-service/book"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var books map[string]book.Book

func main() {
	books = make(map[string]book.Book)

	r := mux.NewRouter()
	r.HandleFunc("/book", createBook).Methods(http.MethodPost)
	r.HandleFunc("/book/{isbn}", getBook).Methods(http.MethodGet)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	srv.ListenAndServe()
}

func getBook(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	e := json.NewEncoder(w)

	if b, ok := books[v["isbn"]]; ok {
		e.Encode(b)
		return
	}

	je := JsonError{
		Msg:  "No book could be found with ISBN: " + v["isbn"],
		Code: 20002,
	}

	e.Encode(je)
	w.WriteHeader(http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	b := book.Book{}

	json.NewDecoder(r.Body).Decode(&b)

	books[b.ISBN] = b

	w.WriteHeader(http.StatusCreated)
}

type JsonError struct {
	Msg  string `json:"message"`
	Code int    `json:"error_code"`
}
