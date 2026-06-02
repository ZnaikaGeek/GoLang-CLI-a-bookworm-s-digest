package main

import (
    "fmt"
    "os"
)

func main() {
/* Call the function with the path
to the sample data
*/
    bookworms, err := loadBookworms("testdata/bookworms.json")

    if err != nil {
        /* If there's an ERROR,print it
        to standart Error and exit
        */
        fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)

        os.Exit(1)
    }
/* Print the raw data to verify it was
loaded correctly
*/
    fmt.Println(bookworms)
}
