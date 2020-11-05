package main

import (
	"fmt"
)

/*
	Application of Functions as Parameters
The way to create a function parameter is to directly write the function scheme as a data type
The callback parameter is a closure declared of type func (string) bool.
The closure is called on each loop in the filter () function.
Filter () function for filtering array data (which data is obtained from the first parameter)
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

// IsPolindrom for the unit test and selecting palindrome numbers
func IsPolindrom(barisData []int) []int {
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
	return polindrom
}

func main() {
	var a, n int

	/*
		input two numbers a and n separated by a space character (a n) resulting arithmetic sequence between a and n (barisData)
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

	dataPolindrom := IsPolindrom(barisData)

	/*
		palindrome => callback closure contains a statement to detect the palindrome number on each element. If any element is polyndrome, it means that element passed the filter.
	*/

	// var polindrom = filter(barisData, func(each int) bool {
	// 	var remainder, temp int
	// 	var reverse int = 0

	// 	temp = each

	// 	for {
	// 		remainder = each % 10
	// 		reverse = reverse*10 + remainder
	// 		each /= 10

	// 		if each == 0 {
	// 			break
	// 		}
	// 	}
	// 	return temp == reverse
	// })

	// fmt.Println("baris data polindrom \"1\"\t:", polindrom)
	fmt.Printf("Output : \n %v \n", len(dataPolindrom))

}
