package statemament

import (
	"fmt"
)

// testing the switch statement function
func SwitchStatement() {

	namaDepan := "Fajar"
	switch namaDepan {
	case "Fajar":
		fmt.Println("Nama Depan = ", namaDepan)
	case "Kurniawan":
		fmt.Println("Nama Depan = ", namaDepan)
	default:
		fmt.Println("Nama Depan bukan Fajar")
	}

	// Bisa membuat ekspresi juga disetiap casenya
	switch lengthNamDepan := len(namaDepan); {
	case lengthNamDepan >= 5:
		fmt.Println("Nama Depan Besar =", namaDepan)
	case lengthNamDepan < 5:
		fmt.Println("Nama Depan Kecil =", namaDepan)
	case lengthNamDepan >= 10:
		fmt.Println("Nama Depan Sangat Besar =", namaDepan)
	}

	// bisa juga tanpa suatu variabel
	namaTengah := "Aziz"
	switch lengthNamTeng := len(namaTengah); {
	case namaTengah == "Fajar":
		fmt.Println("Nama Tengah =", namaTengah)
	case namaTengah == "Aziz" && lengthNamTeng >= 4:
		fmt.Println("Nama Tengah =", namaTengah)
	case namaTengah == "Kurniawan":
		fmt.Println("Nama Tengah =", namaTengah)
	default:
		fmt.Println("Nama Tengah bukan Fajar")
	}

}
