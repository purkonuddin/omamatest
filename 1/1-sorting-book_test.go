package main

import (
	"strings"
	"testing"
)

// data input: 3A13 5X19 9Y20
// output : 9Y20 5X19 3A13
var allBooks = []book{
	{categoryID: 3, bookName: "A", bookHeight: 13, bookID: "3A13"},
	{categoryID: 5, bookName: "X", bookHeight: 19, bookID: "5X19"},
	{categoryID: 9, bookName: "Y", bookHeight: 20, bookID: "9Y20"},
}

func TestSortAllBooks(t *testing.T) {
	var v []string
	v = SortAllBooks(allBooks)
	output := strings.Join(v, " ")
	if output != "9Y20 5X19 3A13" {
		t.Error("Expected 1.5, got ", v)
	}
}

// go test
/*
  	PASS
	ok      omamatest/1     0.052s
*/
