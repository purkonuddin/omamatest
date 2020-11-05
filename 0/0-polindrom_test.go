package main

import (
	"testing"
)

// data input: 1 10
// output : 9

// deret angka antara 1 dan 10 adalah
var barisData = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// for i := 1; i <= 10; i++ {
// 	barisData = append(barisData, i)
// }

func TestIsPolindrom(t *testing.T) {
	var v []int
	v = IsPolindrom(barisData)
	output := len(v)
	if output != 9 {
		t.Error("Expected 1.5, got ", v)
	}
}

// go test
/*
  	PASS
	ok      omamatest/0     0.030s
*/
