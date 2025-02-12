package dpr

import (
	"fmt"
)

// PanicFunction is a function to show how to create a panic function in golang
// digunakan untuk menghentikan eksekusi program dan menampilkan pesan error
// akan tetapi jika ada defer didalamnya maka akan tetap di ekskusi
func PanicFunction(isError bool) {

	log := func() {
		fmt.Println("Selesai eksekusi")
	}
	
	defer log()
	if isError {
		fmt.Println("Mencoba ekseskusi sebelum error")
		panic("Error terjadi") // jika ingin menggunakan function harus mengabaikan return value
	}

}
