package learning

import (
	"fmt"
)

// testing the main variable function
// variabel bisa menjadi global scope bila dideklarasikan di luar fungsi
func Variable() {
	var name string = "John Doe" // tidak wajib menuliskan tipe data dikarenkana golang sudah langsung mengenali tipe data dari nilai yang diberikan

	umur := 30 // tidak wajib menuliskan var tapi dengan syarat langsung diisi nilai dan menggunakan :=
	fmt.Println("Name = ", name)
	fmt.Println("Age = ", umur)

	umur = 40 // mengganti nilai variabel umur
	fmt.Println("Age = ", umur)

	// multiple variable declaration
	var (
		firstName = "Fajar"
		lastName  = "Aziz"
	)

	fmt.Println("First Name = ", firstName)
	fmt.Println("Last Name = ", lastName)

	// constant variable
	const pi = 3.14 // signle declaration
	const (
		phi = 1.61 // multiple declaration
	)

	// pi = 3.15 // error karena nilai konstanta tidak bisa diubah
	fmt.Println("Pi = ", pi)
	fmt.Println("Pi = ", phi)
}
