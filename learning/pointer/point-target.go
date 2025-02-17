package pointer

import (
	"fmt"
)

// testing the pointer function
// mudahnya pointer ini seperti reference di typescript
// jadi bisa dibilang pointer adalah kemampuan untuk mengakses referensi memory dari variable utama tanpa melakukan copy
func PointerDan() {

	// deklarasi variable
	type data struct {
		name string
		age  int
	}
	user := data{name: "Jyriz", age: 25}

	// deklarasi pointer
	// & digunakan untuk mengambil alamat memory dari variable, jika tanpa & itu akan melakukan copy saja
	// * jika berada di samping kanan adalah tipe data pointer
	// tanpa menggunakan pointer &
	user1 := user
	user1.age = 26

	// menggunakan pointer
	var user2 *data = &user
	user2.age = 27 // dengan begini user dan user2 akan memiliki value yang sama

	// menampilkan data awal
	fmt.Println("Data dari awal:", user)

	// menampilkan data dari pointer
	fmt.Println("Data dari bukan pointer:", user1)

	// menampilkan data dari pointer
	fmt.Println("Data dari pointer:", user2)

}
