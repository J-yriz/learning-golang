package learning

import (
	"fmt"
)

// testing the main type declare function
func TypeDeclare() {

	type DataDiri struct {
		NamaDepan    string
		NamaBelakang string
		Umur         int
		Alamat       string
	}

	data := DataDiri{NamaDepan: "Fajar", NamaBelakang: "Aziz", Umur: 30, Alamat: "Jakarta"}
	fmt.Println("Nama Depan = ", data)
	fmt.Println(DataDiri{NamaDepan: "Fajar", NamaBelakang: "Kurniawan", Umur: 19, Alamat: "Semarang"})

	// conversion to type
	type buatSendiri string
	typeSendiri := "Fajar"
	typeSendiri1 := buatSendiri(typeSendiri)
	fmt.Println("Type Sendiri = ", typeSendiri1)

}
