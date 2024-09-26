package main

import (
	"reflect"
	"testing"
)

type testcase struct {
	bookwormsFile string
	want          interface{}
	wantErr       bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestBookworm(t *testing.T) {
	tests := map[string]testcase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file does not exists": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)
			if testCase.wantErr && err == nil {
				t.Fatalf("expected an error %s, got none", err)
			}

			//NOTE: DeepEqual is not recommended for production code as it is slow performing
			if !reflect.DeepEqual(got, testCase.want) {
				t.Fatalf("expected: %v, got: %v", testCase.want, got)
			}
		})
	}

}

func TestBooksCount(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  map[Book]int
	}

	tt := map[string]testCase{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]int{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]int{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]int{handmaidsTale: 1, theBellJar: 1},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, handmaidsTale}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]int{handmaidsTale: 3, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := bookCount(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestCommonBooks(t *testing.T) {
	tests := map[string]testcase{
		"have one common book": {
			bookwormsFile: "testdata/bookworms.json",
			want:          []Book{handmaidsTale},
			wantErr:       false,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			bookworms, err := loadBookworms(testCase.bookwormsFile)
			if testCase.wantErr && err == nil {
				t.Fatalf("expected an error %s, got none", err)
			}

			got := findCommonBooks(bookworms)
			want := []Book{handmaidsTale}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("expected: %v, got: %v", testCase.want, got)
			}
		})

	}
}

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(t *testing.T, got, want map[Book]int) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	// Ranging over the want to retrieve all the keys.
	for book, targetCount := range want {
		// Verify the book in present in the map we check against.
		count, ok := got[book]
		// Book is not found or if found, counts are different.
		if !ok || targetCount != count {
			return false
		}
	}

	// Everything is equal!
	return true
}
