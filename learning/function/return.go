package function

import (
	"fmt"
)

// FunctionReturn1 is a function to show how to create a function with return in golang
func FunctionReturn1(firstName string) string {

	fmt.Println("Hello", firstName)
	return fmt.Sprintf("Hello ini hasil return %s, dan berumur %d tahun", firstName, 19)

}

// multiple return
func FunctionReturn2(firstName string, umur int) (string, int) {

	fmt.Println("Hello", firstName)
	return fmt.Sprintf("Hello ini hasil return %s, dan berumur %d tahun", firstName, umur), umur

}

// return array, slice or map
func FunctionReturn3() ([]string, int) {

	data := []string{"Fajar", "Nur", "Hidayat"}
	return data, len(data)

}

// return menggunakan variabel atau nambed return
func FunctionReturn4() (nama string, umur int) { // jika semua tipe data sama semua maka bisa di tulis seperti ini (nama, umur string)

	nama = "Fajar"
	umur = 19
	return nama, umur

}
