package main

import (
	"fmt"
)

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

func main() {

	deretAngka := []int{23, 24, 25, 26, 27, 28, 30}
	fmt.Println("Input: ", deretAngka)

	totalData := hitung(deretAngka)

	tempTotal := jumlahSn(deretAngka)

	hilang := tempTotal - totalData
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
