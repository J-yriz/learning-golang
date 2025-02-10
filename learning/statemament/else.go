package statemament

import (
	"fmt"
)

// testing the else statement function
func ElseStatement() {

	namaDepan := "Fajar"

	if namaDepan == "Fajar" {
		fmt.Println("Nama Depan = ", namaDepan)
	} else {
		fmt.Println("Nama Depan bukan Fajar")
	}

}
