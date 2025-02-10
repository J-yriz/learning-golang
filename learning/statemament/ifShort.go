package statemament

import (
	"fmt"
)

// testing the if short statement function
func IfShortStatement() {

	namaDepan := "Fajar"

	// Mudahnya ini seperti sebelum if mendeclarasikan variabel panjangNamDep tapi bedanya ini hanya dibisa diakses di dalam statement
	if panjangNamDep := len(namaDepan); panjangNamDep > 5 {
		fmt.Println("Nama Depan Panjang =", namaDepan)
	} else if panjangNamDep > 10 {
		fmt.Println("Nama Depan Sangat Panjang =", namaDepan)
	} else {
		fmt.Println("Nama Depan Kependekan")
	}
	
}
