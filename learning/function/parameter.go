package function

import (
	"fmt"
)

// Function is a function to show how to create a function in golang
func FunctionSayHello(firstName string) {

	fmt.Println("Hello", firstName)

}

// function yang menjadi parameter
// keyword func sendiri bisa menjadi tipe data dan di dalam kurung itu adalah parameter yang akan diisi dan yang di luar setelah kurung itu return value
func FunctionParameter(firstName, lastName string, fungsiParam func(s1, s2 string) string) string {
	return fungsiParam(firstName, lastName)
}

// dikanrekana func sendiri bisa menjadi tipe data, akan panjang jika tipe data tersebut tidak disimpan dalam custom type
// maka dari itu bisa menggunakan type alias
type FungsiParam func(s1, s2 string) string

func FungsiParameter1(firstName, lastName string, fungsiParam FungsiParam) string {
	return fungsiParam(firstName, lastName)
}
