package main

import (
	"fmt"
	"strconv"
	"strings"
)

// hitung fungsi untuk menghitung jumlah deret angka
func countData(angka []int) float64 {
	var total float64 = 0
	for _, angka := range angka {
		total += float64(angka)
	}
	return float64(total)
}

// jumlahSn fungsi untuk menghitung jumlah deret aritmatika dengan rumus Sn=1/2n(a+un)
func countSn(deretAngka []int) float64 {
	var channel2 = make(chan int)
	go nilaimaksimal(deretAngka, channel2)
	maksimal := <-channel2

	n := len(deretAngka) + 1
	a := deretAngka[0]
	un := maksimal
	sn := (1 * n / 2) * (a + un)
	return float64(sn)
}

// SearchLessNumber untuk test dan mencari angka yang hilang
func SearchLessNumber(deretAngka []int) float64 {
	countByDataReal := countData(deretAngka)
	countByRumusSn := countSn(deretAngka)
	// selisih kedua nya menghasilkan bilangan yang hilang
	hilang := countByRumusSn - countByDataReal

	return hilang
}

func main() {

	deretAngka := []int{23, 24, 25, 26, 27, 28, 30}

	// menampilkan deret angka tanpa spasi
	angka := func(deretAngka []int) []string {
		result := []string{}
		for _, a := range deretAngka {
			result = append(result, strconv.Itoa(a))
		}
		return result
	}

	fmt.Println("Input: ", strings.Join(angka(deretAngka), ""))

	// cari angka yang hilang
	hilang := SearchLessNumber(deretAngka)
	// output angka yang hilang
	fmt.Println("hilang: ", hilang)
}

func nilaimaksimal(number []int, ch chan int) {
	var max = number[0]
	for _, e := range number {
		if max < e {
			max = e
		}
	}
	ch <- max
}
