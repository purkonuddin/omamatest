package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func deretAngka(angka []int) float64 {

	totalData := hitung(angka)
	tempTotal := jumlahSn(angka)
	hilang := tempTotal - totalData

	return hilang
}

func index(w http.ResponseWriter, r *http.Request) {
	deret := []int{23, 24, 25, 26, 27, 28, 30}

	angka := func(deret []int) []string {
		result := []string{}
		for _, a := range deret {
			result = append(result, strconv.Itoa(a))
		}
		return result
	}

	fmt.Fprintln(w, "Input: ", strings.Join(angka(deret), ""))
	fmt.Fprintln(w, "Nilai yang hilang: ", deretAngka(deret))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	http.HandleFunc("/index", index)
	fmt.Println("memulai web server pada localhost:8080")
	http.ListenAndServe(":8080", nil)
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

func hitung(angka []int) float64 {
	var total float64 = 0
	for _, angka := range angka {
		total += float64(angka)
	}
	return float64(total)
}

func jumlahSn(deretAngka []int) float64 {
	var channel2 = make(chan int)
	go nilaimaksimal(deretAngka, channel2)
	maksimal := <-channel2

	n := len(deretAngka) + 1
	a := deretAngka[0]
	un := maksimal
	sn := (1 * n / 2) * (a + un)
	return float64(sn)
}
