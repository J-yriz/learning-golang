package errorLearn

import (
	"errors"
)

// testing the error function
// digolang kita bisa membuat error sendiri menggunakan package errors
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh mengandung angka 0")
	} else {
		return nilai / pembagi, nil
	}
}
