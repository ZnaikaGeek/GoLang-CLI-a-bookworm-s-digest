package main

import (
    "encoding/json"
    "fmt"
    "os"
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
