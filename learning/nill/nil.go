package nill

import (
	"fmt"
)

// testing the nill function
func Nill() {

	// nill itu seperti null di javascript
	// nill itu tidak bisa di isi dengan data apapun
	// nill itu hanya bisa di isi dengan data yang sama dengan tipe data yang di tentukan
	// nill itu biasanya digunakan untuk mengecek apakah data tersebut ada atau tidak
	// nill itu bisa di isi pada pointer, interface, function, slice, map, channel, dan pointer, dan ekslusif return error

	// contoh
	membuatMap := func(dataMap string) map[string]string {
		if dataMap == "" {
			return nil // bisa return nil dikarenakan ini dapat disi di interface, function, slice, map, channel, dan pointer. Selain itu tidak bisa
		} else {
			return map[string]string{
				"data": dataMap,
			}
		}
	}

	dataBaru := membuatMap("Fajar")
	if dataBaru == nil {
		fmt.Println("Data Map Kosong")
	} else {
		fmt.Println("Data Map =", dataBaru["data"])
	}
}
