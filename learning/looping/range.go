package looping

import (
	"fmt"
)

// Forlooping range is a function to show the looping in Golang
func ForloopingRange() {

	// for yang biasa yang digunakan untuk array, slice, map, string
	// basic untuk mengakses data yang ada di dalam array, slice, map, string
	dataMahasiswa := []string{"Budi", "Andi", "Joko", "Rudi", "Dodi"}
	for i := 0; i < len(dataMahasiswa); i++ {
		fmt.Println(dataMahasiswa[i])
	}

	// menggunakan for range
	for index, data := range dataMahasiswa {
		fmt.Println(`Index ke`, index+1, `adalah`, data)
	}

	// jika ada data yang tidak di gunakan maka bisa menggunakan _ (underscore)
	for _, data := range dataMahasiswa {
		fmt.Println(data)
	}

}
