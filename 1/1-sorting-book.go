package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type kategori struct {
	kID   int
	kName string
}

type book struct {
	categoryID int
	bookName   string
	bookHeight int
	bookID     string
}

// Kombinasi Slice & Struct untuk menampung data kategory
var allCategories = []kategori{
	{kID: 6, kName: "Applied Sciences (600)"},
	{kID: 7, kName: "Arts (700)"},
	{kID: 0, kName: "General (000)"},
	{kID: 9, kName: "Geography, History (900)"},
	{kID: 4, kName: "Language (400)"},
	{kID: 8, kName: "Literature (800)"},
	{kID: 1, kName: "Literature (100)"},
	{kID: 2, kName: "Literature (200)"},
	{kID: 5, kName: "Literature (500)"},
	{kID: 3, kName: "Literature (300)"},
}

// appendtostruct, Kombinasi Slice & Struct untuk menampung data indut kedalam struct buku
func appendtostruct(s []string) []book {
	var allBooks = []book{}
	for _, bookID := range s {
		categoryID, _ := strconv.Atoi(bookID[0:1])
		bookName := bookID[1:2]
		bookHeight, _ := strconv.Atoi(bookID[2:4])

		b := book{categoryID, bookName, bookHeight, bookID}
		allBooks = append(allBooks, b)
	}
	return allBooks
}

// SortAllBooks fungsi ini digunakan untuk test
func SortAllBooks(allBooks []book) []string {
	//make chanel to sorting allBooks by category
	var channel1 = make(chan []string)
	//kirim channel1
	go sortByCategory(allBooks, allCategories, channel1)
	// tmpBooks untuk tampung data buku yang sudah di sorting
	tmpBooks := <-channel1
	return tmpBooks

}

func main() {
	fmt.Print("input : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	s := strings.Split(input, " ")
	// allBooks: untuk menampung data buku
	allBooks := appendtostruct(s)

	// fungsi untuk mengurutkan data allBooks
	tmpBooks := SortAllBooks(allBooks)

	output := strings.Join(tmpBooks, " ")
	fmt.Println("Output: ", output)
}

func sortByCategory(allBooks []book, allCategories []kategori, ch chan []string) {
	var tmpBooks []string

	for i := 0; i < len(allCategories); i++ {
		var num []int
		// fmt.Println(allCategories[i].kID)
		for j := 0; j < len(allBooks); j++ {
			// fmt.Println(allBooks[j].categoryID)
			if allCategories[i].kID == allBooks[j].categoryID {
				num = append(num, allBooks[j].bookHeight)
			}
		}

		sort.Slice(num, func(a, b int) bool {
			return num[a] > num[b]
		})

		for _, v := range num {
			for a := 0; a < len(allBooks); a++ {
				if allCategories[i].kID == allBooks[a].categoryID {
					if v == allBooks[a].bookHeight {
						tmpBooks = append(tmpBooks, allBooks[a].bookID)
					}
				}
			}
		}
	}
	ch <- tmpBooks
}

// input: 3A13 5X19 9Y20 2c18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14
// output: 0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2c18 5X19 5G14 3N21 3N20 3N14 3A13
