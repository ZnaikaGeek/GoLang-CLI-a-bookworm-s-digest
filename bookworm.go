package main

// Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
    Name  string `json:"name"`
    Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
    Author string `json:"author"`
    Title  string `json:"title"`
}
