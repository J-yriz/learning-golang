package struct_learning

import (
	"fmt"
)

// Interface any function
// mudahnya any itu seperti any di typescript
func Tampilkan(data any) {
	fmt.Println(data)
}

func PanggilTampil() {
	Tampilkan(1)
	Tampilkan("Fajar")
	Tampilkan(true)
	Tampilkan(3.14)
}
