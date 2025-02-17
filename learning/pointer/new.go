package pointer

import (
	"fmt"
)

// testing the new function
// new adalah function yang digunakan untuk membuat pointer dari tipe data yang diinginkan
func NewPointer() {

	// deklarasi variable
	type data struct {
		name string
		age  int
	}
	user := data{name: "Jyriz", age: 25}

	// deklarasi pointer
	// new digunakan untuk membuat pointer dari tipe data yang diinginkan
	// jika menggunakan new maka akan membuat sebuah data kosong
	user1 := new(data)
	user1.age = 26 // jika kita mengisi data seperti ini maka mudahnya seperti ini user1 := &data{age: 26}

	// menampilkan data awal
	fmt.Println("Data dari awal:", user)

	// menampilkan data dari pointer
	fmt.Println("Data dari pointer:", user1)

}