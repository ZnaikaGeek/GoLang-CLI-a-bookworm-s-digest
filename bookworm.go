// bookworm.go
/* Key Concepts used:
os.Open: This opens the file in 
read-only mode.
defer: This keyword ensures that 
f.Close() is called automatically 
when the function exits, 
which prevents system resource leaks.
json.NewDecoder: This is efficient 
for reading from a file stream 
(an io.Reader) rather than loading 
the entire file into memory first.
*/

package main

import (
    "encoding/json"
    "fmt"
    "os"
    "sort"
)

/*Bookworm contains the list of books 
on a bookworm's shelf.*/
type Bookworm struct {
    Name  string `json:"name"`
    Books []Book `json:"books"`
}

/*Book describes a book on
a bookworm's shelf.*/
type Book struct {
    Author string `json:"author"`
    Title  string `json:"title"`
}

/* byAuthor is a list of Book.
Defining a custom type to implement
sort.Interface. */
type byAuthor []Book

/* Len implements sort.Interface 
by returning the length of the 
 collection. */
func (b byAuthor) Len() int { 
    return len(b) 
}

/* Swap implements sort.Interface and 
swaps two books in place. */
func (b byAuthor) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

/* Less implements sort.Interface and 
defines the sorting criteria.
We sort by Author first, and then by 
Title if the authors are the same. */
func (b byAuthor) Less(i, j int) bool {
    if b[i].Author != b[j].Author {
        return b[i].Author < b[j].Author
    }
    return b[i].Title < b[j].Title
}

/*load bookWorms reads the file 
and returns the list of bookworms,
and their beloved books, found therein.
*/

func loadBookworms(filePath string) ([]Bookworm, error) {
    // 1. Open the file for reading
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }

    /* 2. Ensure the file is closed when
    when the function finishes */
    defer f.Close()

    /* 3. Initialize the slice where the 
    file will be decoded
    */
    var bookworms []Bookworm

    /* 4. Create JSON decoder and decode
    the file content into the slice */
    err = json.NewDecoder(f).Decode(&bookworms)
    if err != nil {
        return nil, err
    }
    return bookworms, nil
}

/* booksCount registers all the books
and their occurrences from the 
bookworm's shelves
*/
func booksCount(bookworms []Bookworm) map[Book]uint {
    count := make(map[Book]uint)

    for _, bookworm := range bookworms {
        for _, book := range bookworm.Books {
/* if the book is not in the map, 
count[book] returns the zero value (0)
*/
        count[book]++
       }
    }
  return count
}

func findCommonBooks(bookworms []Bookworm) []Book {
/* function returns books that are 
on more than one bookworm's shelf
*/
    booksOnShelves := booksCount(bookworms)

    var commonBooks []Book

    for book, count := range booksOnShelves {
        if count > 1 {
            commonBooks = append(commonBooks, book)
        }
    }
    return sortBooks(commonBooks)
}

/* sortBooks sorts the books by Author
and then Title 
*/
func sortBooks(books []Book) []Book {
    sort.Slice(books, func(i, j int) bool {
    if books[i].Author != books[j].Author {
        return books[i].Author < books[j].Author
    }
    return books[i].Title < books[j].Title
    })
    return books
}

/* function to print out the titles 
and the authors of a list of books
*/
func displayBooks(books []Book) {
    for _, book := range books {
        fmt.Println("-", book.Title, "by", book.Author)
    }
}
