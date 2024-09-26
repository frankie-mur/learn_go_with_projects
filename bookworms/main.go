package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "testdata/bookworms.json", "file path to use")

	bookworms, err := loadBookworms(file)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
