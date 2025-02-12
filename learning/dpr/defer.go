package dpr

import (
	"fmt"
)

// DeferFunction is a function to show how to create a defer function in golang
// digunakan untuk menunda eksekusi function sampai function yang di dalamnya selesai di eksekusi
func DeferFunction() {

	log := func() {
		fmt.Println("Selesai eksekusi")
	}

	// pendifinisian defer sendiri akan selalu dilakukan walaupun function yang di dalamnya error
	defer log() // akan dieksekui pada saat print dibawah ini dieksekusi
	fmt.Println("Sedang eksekusi")

}
