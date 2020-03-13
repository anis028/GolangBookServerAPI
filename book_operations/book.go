package book_operations

import (
	"encoding/json"
	"github.com/anis028/first_api_project/library"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


// returns all all_books in the list
func GetAllBooks(w http.ResponseWriter, r *http.Request){
	//fmt.Println(r.Header.Get("Content-Type"))
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(library.AllBooks)
}

// returns a single book according to id
func GetSingleBook(w http.ResponseWriter, r *http.Request){
	//json.Marshal(r.Body)
	w.Header().Set("Content-Type", "application/json")
	searchId, _ := strconv.Atoi(mux.Vars(r)["id"])
	// Gets parameters from the request body

	//Loop through all_books and find one with the id from the params
	for _, item := range library.AllBooks {
		if item.ID == searchId {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	log.Println("notFound")
	_ = json.NewEncoder(w).Encode("")
}


// Add new book
func CreateBook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var book library.Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = len(library.AllBooks) + 1
		library.AllBooks = append(library.AllBooks, book)
		json.NewEncoder(w).Encode(book)
	})
}

// Update book
func UpdateBook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		search_id, _ := strconv.Atoi(mux.Vars(r)["id"])
		for index, item := range library.AllBooks {
			if item.ID == search_id {
				library.AllBooks = append(library.AllBooks[:index], library.AllBooks[index+1:]...)
				var book library.Book
				_ = json.NewDecoder(r.Body).Decode(&book)
				book.ID = search_id
				library.AllBooks = append(library.AllBooks, book)
				_ = json.NewEncoder(w).Encode(book)
				return
			}
		}
	})
}

// Delete book
func DeleteBook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		searchId, _ := strconv.Atoi(mux.Vars(r)["id"])
		for index, item := range library.AllBooks {
			if item.ID == searchId {
				library.AllBooks = append(library.AllBooks[:index], library.AllBooks[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(library.AllBooks)
	})
}