package struct_learning

import (
	"fmt"
)

// jika struct berada diluar otomatis menjadi public
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Struct is a function to show how to create a struct in golang
func Struct() {

	// jika berada didalam otomatis menjadi private
	type Human struct {
		FirstName string
		LastName  string
		Age       int
	}

	human1 := Human{FirstName: "Fajar", LastName: "Aziz", Age: 19}
	// bisa juga seperti ini
	human2 := Human{"Fajar", "Kurniawan", 25}
	fmt.Println(human1, human2)

}
