package main

import (
	"encoding/json"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// This function takes a filepath, opens the file and reads its contents
func loadBookworms(filepath string) ([]Bookworm, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func bookCount(bookworms []Bookworm) map[Book]int {
	bookInstances := map[Book]int{}

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			bookInstances[book] += 1
		}
	}

	return bookInstances
}

// findCommonBooks returns books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	bookCount := bookCount(bookworms)
	commonBooks := []Book{}
	for k, v := range bookCount {
		if v > 1 {
			commonBooks = append(commonBooks, k)
		}
	}

	return commonBooks
}

// sortBooks sorts the books by Author and then Title.
func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})

	return books
}
