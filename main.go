package main

import (
	"fmt"
	"github.com/anis028/first_api_project/book_operations"
	"github.com/anis028/first_api_project/library"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", book_operations.GetAllBooks).
		Methods("GET")
	r.HandleFunc("/books/{id}", book_operations.GetSingleBook).
		Methods("GET")
	r.HandleFunc("/login", library.Login).
		Methods("GET")
	//r.HandleFunc("/token", library.GenerateToken).
	//	Methods("GET")
	//r.HandleFunc("/books", book_operations.CreateBook).Methods("POST")
	r.Handle("/books", library.IsAuthorized(book_operations.CreateBook())).
		Methods("POST")
	//r.HandleFunc("/books/{id}", book_operations.UpdateBook).Methods("PUT")
	r.Handle("/books/{id}", library.IsAuthorized(book_operations.UpdateBook())).
		Methods("PUT")
	//r.HandleFunc("/books/{id}", book_operations.DeleteBook).Methods("DELETE")
	r.Handle("/books/{id}", library.IsAuthorized(book_operations.DeleteBook())).
		Methods("DELETE")
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Sprint()
	log.Fatal(srv.ListenAndServe())
}