package dpr

import (
	"fmt"
)

// RecoveryFunction is a function to show how to create a recovery function in golang
// digunakan untuk menangkap panic yang terjadi pada function, dan membiarkan code tetap berjalan
// ada cara yang benar dalam menggunakan recovery

// cara yang benar
func RecoveryFunction(isError bool) {

	log := func() {
		fmt.Println("Selesai eksekusi")
		msgError := recover()
		if msgError != nil {
			fmt.Println("Pesan errornya adalah: ", msgError)
		}
	}

	defer log()
	if isError {
		fmt.Println("Mencoba ekseskusi sebelum error")
		panic("Error terjadi")
	}

}

// cara yang salah
// kenapa salah? di karenakan panic itu sendiri jika sudah terjadi akan menghentikan eksekusi program
// jika di berikan ke defer akan di eksekui karna sifatnya defer akan di eksekui di akhir function apapun yang terjadi
func RecoveryFunction2(isError bool) {

	log := func() {
		fmt.Println("Selesai eksekusi")
	}

	defer log()
	if isError {
		fmt.Println("Mencoba ekseskusi sebelum error")
		panic("Error terjadi")
	}

	msgError := recover()
	fmt.Println("Pesan errornya adalah: ", msgError)

}
