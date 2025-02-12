package function

import (
	"fmt"
)

// FunctionVariadic is a function to show how to create a function with variadic in golang
// note: variadic adalah parameter yang bisa di isi lebih dari satu dan paramenter harus diketik di akhir
func FunctionVariadicAddData(firstName, lastName string, data ...string) {

	fmt.Println("Hello", firstName, lastName)
	fmt.Println("Data yang di inputkan adalah:")
	for _, val := range data {
		fmt.Println(val)
	}

}

func SumNumber(number ...int) int {

	result := 0
	for _, val := range number {
		result += val
	}
	return result

}

// bisa juga parameter menggunakan slice
// akan tetapi melakukan input parameter harus juga menggunakan slice []int{1, 2, 3, 4, 5}
func SumNumber2(number []int) int {

	result := 0
	for _, val := range number {
		result += val
	}
	return result

}
