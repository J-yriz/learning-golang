package learning

import (
	"fmt"
)

// testing the main convertion function
func Convertion() {
	nilaiInt16 := int16(1000)
	nilaiInt32 := int32(nilaiInt16)
	nilaiInt64 := int64(nilaiInt32)

	fmt.Println("Nilai Int16 = ", nilaiInt16)

	ubah64Ke16 := int16(nilaiInt64)
	fmt.Println("Ubah Int64 ke Int16 = ", ubah64Ke16)
}
