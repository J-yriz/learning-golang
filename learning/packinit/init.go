package packinit

import (
	"fmt"
)

// testing the init function
// didalam golang pada saat kita melakukan import dan menambahkan init function maka function tersebut akan dijalankan terlebih dahulu sebelum function main
// atau bisa dibilang init adalah function yang pertama kali dijalankan ketika program dijalankan atau di panggil

// INIT = INITIALIZATION = INISIALISASI

// namun jika kita hanya ingin menjalankan init function bisa menggunakan _ (underscore) sebelum import package
// dikarenakan init ini akan dijalankan ketika di import dan dijalan terlebih dahulu sebelum function main
// maka agar kita bisa menjalankan init function tapi tidak menggunakan funciton yang ada di package tersebut kita bisa menggunakan _ (underscore)
// contoh: import _ "nameproject/packinit"

type TestingInit struct {
	Data string
}

var DataTesting = new(TestingInit)

func init() {
	DataTesting.Data = "Hello, this is the init function!"
}

func GetTestingData() {
	fmt.Println(DataTesting.Data)
}
