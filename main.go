package main

import (
	"fmt"

	// "github.com/J-yriz/learning-golang/learning"
	// "github.com/J-yriz/learning-golang/learning/array"
	// "github.com/J-yriz/learning-golang/learning/statemament"
	// "github.com/J-yriz/learning-golang/learning/looping"
	// "github.com/J-yriz/learning-golang/learning/function"
	"github.com/J-yriz/learning-golang/learning/dpr"
)

// testing the main function
func main() {
	fmt.Println("Hello, this is the main function!")

	// call the data type function from learning package
	// learning.DataType()

	// call the variable function from learning package
	// learning.Variable()

	// call the convertion function from learning package
	// learning.Convertion()

	// call the type declare function from learning package
	// learning.TypeDeclare()

	// call the perbandingan function from learning package
	// learning.Perbandingan()

	// call the array function from learning package
	// learning.Array()

	// call the slice type function from learning package
	// array.SliceType()

	// call the map type function from learning package
	// array.MapType()

	// call the if statement function from statemament package
	// statemament.IfStatement()

	// call the else statement function from statemament package
	// statemament.ElseStatement()

	// call the else if statement function from statemament package
	// statemament.ElseIfStatement()

	// call the if short function from statemament package
	// statemament.IfShortStatement()

	// call the switch statement function from statemament package
	// statemament.SwitchStatement()

	// call the switch short function from statemament package
	// statemament.SwitchShortStatement()

	// call the for looping function from looping package
	// looping.ForLooping()

	// call the for looping range function from looping package
	// looping.ForloopingRange()

	// call the function parameter function from function package
	// function.FunctionSayHello("Fajar")

	// call the function return function from function package
	// fmt.Println(function.FunctionReturn1("Fajar"))

	// call the multiple return function from function package
	// fmt.Println(function.FunctionReturn2("Fajar", 19))
	// dengan catatan jika return valuenya lebih dari satu, pada saat di simpan menggunakan variabel maka variabel harus lebih dari satu
	// hasil, umur := function.FunctionReturn2("Fajar", 19) // akan tetapi jika ingin dihiraukan bisa menggunakan _ (underscore)
	// hasil, _ := function.FunctionReturn2("Fajar", 19)
	// fmt.Println(hasil, umur)

	// call the return array, slice or map function from function package
	// fmt.Println(function.FunctionReturn3())

	// call the return menggunakan variabel function from function package
	// masih sama jika ingin mengabaikan salah satu return value bisa menggunakan _ (underscore)
	// nama, umur := function.FunctionReturn4()
	// fmt.Println(nama, umur)

	// call the function with variadic function from function package
	// function.FunctionVariadicAddData("Fajar", "Nur", "Hidayat", "19", "Jakarta", "Indonesia")
	// dataNumber := []int{1, 2, 3, 4, 5}
	// // jika parameter tersebut itu variadic dan data yang kita punya itu berupa slice kita bisa menggunakan ... untuk mengubah slice menjadi variadic
	// fmt.Println(function.SumNumber(dataNumber...))
	// fmt.Println(function.SumNumber(1, 2, 3, 4, 5))
	// fmt.Println(function.SumNumber2(dataNumber))

	// call the function with value function from function package
	// function.FunctionValue("Fajar")

	// call the function with function parameter function from function package
	// functions := func(s1, s2 string) string { return fmt.Sprintf("Halo %s %s", s1, s2) }
	// variable function, sifatnya sama seperti func pada umumnya (jika variabel harus dari atas ke bawah)
	// function biasa selain function variable tidak bisa dibuat di dalam function
	// hasil := function.FunctionParameter("Fajar", "Aziz", functions)
	// fmt.Println(hasil)

	// call the function recursive function from function package
	// fmt.Println(function.FunctionRecursiveLoop(5)) // menggunakan loop biasa
	// fmt.Println(function.FunctionRecursive(5))     // menggunakan recursive

	// call the function dpr function from dpr package
	// dpr.DeferFunction()

	// call the function panic function from dpr package
	// dpr.PanicFunction(true)

	// call the function recovery function from dpr package
	dpr.RecoveryFunction(false)
}
