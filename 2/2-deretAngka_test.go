package main

import "testing"

func TestSearchLessNumber(t *testing.T) {
	deretAngka := []int{23, 24, 25, 26, 27, 28, 30}

	var v float64
	v = SearchLessNumber(deretAngka)
	if v != 29 {
		t.Error("Expected 1.5, got ", v)
	}
}

/*
comand:
omamtest/2> go test

output:
PASS
ok      omamatest/2     0.029s
*/
