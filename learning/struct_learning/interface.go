package struct_learning

import (
	"fmt"
)

// Interface
// interface disini berfungsi sebagai kontrak yang mendefinisikan method yang harus diimplementasikan oleh sebuah struct
// jadi mudahnya interface adalah aturan yang dibuat untuk suatu variabel yang dimana variabel tersebut akan di isi oleh struct
// dan struct tersebut harus mengimplementasikan method yang ada di interface tersebut
type Speaker interface {
	SayHello(string)
	// ChangeName(string)
}

// Struct yang mengimplementasikan interface
type PersonEa struct {
	Name string
}

func (p PersonEa) SayHello(target string) {
	fmt.Printf("Hello, my name is %s", p.Name)
}

// Struct lain dengan implementasi berbeda
type Robot struct {
	Model string
}

func (r Robot) SayHello(target string) {
	fmt.Printf("Beep boop! I am %s", r.Model)
}

func Interface() {
	// Bisa menyimpan berbagai tipe yang mengimplementasikan Speaker
	var s Speaker

	s = PersonEa{Name: "Alice"}
	s.SayHello("Fajar") // Output: Hello, my name is Alice

	s = Robot{Model: "T-800"}
	s.SayHello("Aziz") // Output: Beep boop! I am T-800
}

// karna interface juga tipe data, kita bisa melakukan hal seperti ini
// dikarenakan yang disebut interface itu adalah tipe data kita bisa menggunakan any type yang menerima semua tipe data
type DataUser interface {
	GetData() string
}

func SayHello(du DataUser) {
	fmt.Printf("Hallo %s\n", du.GetData())
}

type User struct {
	Name string
}

// implementasi method GetData
// jika di hapus akan terjadi error karna User tidak mengimplementasikan method GetData
func (u User) GetData() string {
	return u.Name
}

func Interface1() {

	user := User{Name: "Fajar"}
	SayHello(user) // diisi oleh struct User yang harus mengimplementasikan method GetData

}