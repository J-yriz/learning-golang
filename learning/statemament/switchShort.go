package statemament

import (
	"fmt"
)

// testing the switch short statement function
func SwitchShortStatement() {

	nameDepan := "Fajar"

	switch lengthNamDepan := len(nameDepan); lengthNamDepan >= 5 {
	case true:
		fmt.Println("Nama Depan Besar =", nameDepan)
	case false:
		fmt.Println("Nama Depan Kecil =", nameDepan)
	}

}