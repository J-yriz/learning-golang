package learning

import (
	"fmt"
)

// testing the main convertion function
func Convertion() {
	nilaiInt32 := int32(32768)
	nilaiInt16 := int16(nilaiInt32)
	nilaiInt64 := int64(nilaiInt32)

	fmt.Println("Nilai Int16 = ", nilaiInt16) // kenapa bisa minus? karena nilai yang dihasilkan melebihi batas maksimal dari int16 makannya balik ke nilai awal
	fmt.Println("Nilai Int64 = ", nilaiInt64)

	ubah16ke32 := int16(nilaiInt32)
	fmt.Println("Ubah Int64 ke Int16 = ", ubah16ke32)

	// conversion to string
	keString := fmt.Sprintf("%d", nilaiInt16)
	fmt.Println("Ubah Int16 ke String = ", keString)

	// conversion to number / integer (tergantung dari size data)
	keInt := int16(nilaiInt16)
	fmt.Println("Ubah String ke Int16 = ", keInt)

	// conversion dari uint8 ke string
	namaLengkap := "Fajar Aziz"
	ambilChar := namaLengkap[0] // akan berubah menjadi byte
	keString = string(ambilChar) // harus diubah ke string (fungsi dari string adalah mengubah byte ke string)
	fmt.Println("Ambil Char ke String = ", keString)
}
