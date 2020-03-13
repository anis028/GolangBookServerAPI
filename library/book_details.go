package library

type Author struct {
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
}

// Book struct (Model)
type Book struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Author  Author  `json:"author"`
}

var AllBooks []Book

func init() {

	//initializing with dummy values
	AllBooks = append(AllBooks, Book{
		ID:    1,
		Title: "first book",
		Author: Author{
			Firstname: "Anisur",
			Lastname:  "Rahman",
		},
	})
	AllBooks = append(AllBooks, Book{
		ID:    2,
		Title: "second book",
		Author: Author{
			Firstname: "Habibur",
			Lastname:  "Rahman",
		},
	})

	// read from local file system

	// fetch from database
}