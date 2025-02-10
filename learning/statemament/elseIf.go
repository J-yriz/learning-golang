package statemament

import (
	"fmt"
)

// testing the else if statement function
func ElseIfStatement() {

	namaDepan := "Aziz"

	if namaDepan == "Fajar" {
		fmt.Println("Nama Depan = ", namaDepan)
	} else if namaDepan == "Kurniawan" {
		fmt.Println("Nama Depan = ", namaDepan)
	} else {
		fmt.Println("Nama Depan bukan Fajar")
	}

}