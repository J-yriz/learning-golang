package learning

import (
	"fmt"
)

// testing the array function
func Array() {

	// array digolang tidak bisa dihapus untuk valuenya (tidak seperti ts)
	// array yang tidak ditentukan jumlahnya
	dataArray := []int{1, 2, 3, 4, 5}
	dataArray = append(dataArray, 6)
	fmt.Println("Data Array =", dataArray)

	// pada dasat digolang pada saat kita membuat array itu harus mendifinikan jumlah tetap
	jumlahTetap := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Jumlah Tetap =", jumlahTetap)

	// jika ada index yang ksoong maka akan di isi dengan nilai default
	indexKosongInt := [5]int{1, 2, 3}
	indexKosongStr := [5]string{"A", "B", "C"}
	fmt.Println("Index Kosong Int =", indexKosongInt) // nilai defaultnya 0
	fmt.Println("Index Kosong Str =", indexKosongStr) // nilai defaultnya ""

	fmt.Println("Panjang Data Array =", len(dataArray)) // len = length

	// jika enggan menuliskan totalnya bisa menggunakan ...
	// jika menggunakan ... maka golang akan menghitung jumlah data yang ada
	dataArray2 := [...]int{1, 2, 3, 4, 5} // tapi dengan syarat data harus langsung diisi
	fmt.Println("Data Array 2 =", dataArray2)
	
}
