package struct_learning

import (
	"fmt"
)

// jika ingin method tersebut langsung ada fungsinya
// tapi hanya bisa digunakan diluar function
type PersonMethod2 struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *PersonMethod2) UbahUmur(umur int) { // dengan begini function tersebut akan menjadi method
	p.Age++
}

// struct method
func StructMethod() {

	// diatas ini secara resmi, bisa juga seperti ini
	type PersonMethod struct {
		FirstName string
		LastName  string
		Age       int
	}

	human1 := PersonMethod{FirstName: "Fajar", LastName: "Aziz", Age: 19}
	tambahUmur := func(p *PersonMethod) {
		p.Age++
	}

	tambahUmur(&human1) 
	// penjelasan dari & adalah pointer https://www.notion.so/Golang-Method-Struct-with-Pointer-198e4a17501980af9e08c8fcf0c9e5eb?pvs=4
	fmt.Println(human1)
}

// contoh pembuatan lainnya

type DataPerson struct {
	FirstName string
	LastName  string
	Age       int
}

// apa perbedaan menggunakan * dan tidak menggunakan?
// singkat * adalah sebagai penanda jika kita mendifinikan perubahan itu akan berpengaruh ke struct aslinya
// contoh yang tidak menggunakan *
func (DP DataPerson) sayHello(name string) {
	fmt.Printf("Hello %s, my name is %s", name, DP.FirstName)
	DP.Age++ // bisa dilihat ada wanring disini yang menandakan bahwa perubahan tidak akan berpengaruh ke struct aslinya
	// jadi hanya cocok untuk menampilkan data saja
}

func StructMethod2() {
	data := DataPerson{
		FirstName: "Fajar",
		LastName:  "Aziz",
		Age:       19,
	}

	data.sayHello("Kurniawan")
	fmt.Println(data.Age) // output 19
}

// contoh yang menggunakan *
func (DP *DataPerson) sayHello2(name string) {
	fmt.Printf("Hello %s, my name is %s", name, DP.FirstName)
	DP.Age++
}

func StructMethod3() {
	data := DataPerson{
		FirstName: "Fajar",
		LastName:  "Aziz",
		Age:       19,
	}

	data.sayHello2("Kurniawan")
	fmt.Println(data.Age) // output 20
}
