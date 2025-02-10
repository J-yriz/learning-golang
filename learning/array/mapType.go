package array

import (
	"fmt"
)

// testing the map type function
func MapType() {

	// pada map type, key harus unik dan value bisa sama (seperti membuat new Map pada javascript)
	// map type tidak bisa menggunakan append
	dataMap := map[string]string{}
	dataMap["namaDepan"] = "Fajar"

	fmt.Println("Data Map =", dataMap)
	// jika key-nya sama maka value akan di replace
	dataMap["namaDepan"] = "Aziz"
	fmt.Println("Data Map =", dataMap)

	// bisa juga membuat seperti ini
	dataMap1 := map[string]string{
		"namaDepan":    "Fajar",
		"namaBelakang": "Aziz",
		"umur":         "30",
	}
	fmt.Println("Data Map 1 =", dataMap1["namaDepan"])
	// jika kita memanggil key yang tidak ada maka yang di tampilkan adalah default value (default value sesuai dengan tipe data)

	// Function untuk map 
	// kurang lebih sama dengan slice dan array akan tetapi ada berapa tambahan
	// delete

	// membuat map baru dengan make
	dataMap2 := make(map[string]string)
	dataMap2["namaDepan"] = "Fajar"
	dataMap2["namaBelakang"] = "Aziz"
	dataMap2["umur"] = "30"
	fmt.Println("Data Map 2 =", dataMap2)

	// delete data pada map
	delete(dataMap2, "umur")
	fmt.Println("Data Map 2 =", dataMap2)
}