package main

import (
	"fmt"
)

/*
	Penerapan Fungsi Sebagai Parameter
	Cara membuat parameter fungsi adalah dengan langsung menuliskan skema fungsi nya sebagai tipe data
	Parameter callback merupakan sebuah closure yang dideklarasikan bertipe func(string) bool.
	Closure tersebut dipanggil di tiap perulangan dalam fungsi filter().
	Fungsi filter() untuk filtering data array (yang datanya didapat dari parameter pertama)
*/

func filter(data []int, callback func(int) bool) []int {
	var result []int
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var a, n int

	/*
		input dua bilangan a & n yang dipisahkan dengan karakter spasi (a n) sehingga menghasilkan barisan aritmatika antara a dan n (barisData)
	*/

	fmt.Println("Input: ")
	fmt.Scan(&a)
	fmt.Scan(&n)

	var barisData []int

	for i := a; i <= n; i++ {
		// fmt.Printf("%v ", i)
		barisData = append(barisData, i)
	}

	fmt.Println()

	// fmt.Printf("Output : \n %v %v \n", input1, input2)
	// fmt.Println("baris data \t\t:", barisData)

	/*
		polindrom => closure callback berisikan statement untuk deteksi polindrom tiap elemen.
		Jika ada elemen yang polindrom, berarti elemen tersebut lolos filter.
	*/

	var polindrom = filter(barisData, func(each int) bool {
		var remainder, temp int
		var reverse int = 0

		temp = each

		for {
			remainder = each % 10
			reverse = reverse*10 + remainder
			each /= 10

			if each == 0 {
				break
			}
		}
		return temp == reverse
	})

	// fmt.Println("baris data polindrom \"1\"\t:", polindrom)
	fmt.Printf("Output : \n %v \n", len(polindrom))

}
