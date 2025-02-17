package typeassertion

import (
	"fmt"
)

// testing the type assertion function
// mudahnya type assertion itu seperti casting di typescript const text = data as string
func TypeAssertion() {

	var data interface{} = 1

	// Type assertion ke string
	str, ok := data.(string)
	// kenapa diberi ok? karena jika data tidak sesuai dengan tipe data yang diinginkan maka akan mengembalikan false
	// jika tanpa ok maka akan terjadi error, karen ok digunakan untuk menangkap apakah data tersebut sesuai dengan tipe data yang diinginkan
	if ok {
		fmt.Println("Data adalah string:", str)
	} else {
		fmt.Println("Data bukan string")
	}

	// bisa juga menggunakan switch
	switch value := data.(type) {
	case string:
		fmt.Println("Data adalah string:", value)
	case int:
		fmt.Println("Data adalah int:", value)
	default:
		// unknown type untuk type jika tidak sesuai dengan tipe yang ada di case
		fmt.Printf("Data bukan string atau int data adalah %T\n", value)
	}
}
