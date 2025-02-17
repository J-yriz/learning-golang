package pointer

import (
	"fmt"
)

// testing the art disk pointer
func ArteriskPointer() {

	type DataHuman struct {
		Name string
		Age  int
	}

	dataHuman1 := DataHuman{Name: "Jyriz", Age: 25}
	dataHuman2 := &dataHuman1
	// jika kita menggunakan & maka dataHuman2 akan memiliki alamat memory yang sama dengan dataHuman1
	// dan variable apapun yang kita ubah pada dataHuman2 maka dataHuman1 juga akan berubah
	// dikarenakan dataHuman2 adalah pointer dari dataHuman1 yang artinya dataHuman2 adalah referensi dari dataHuman1
	dataHuman2.Age = 26

	fmt.Println("Data Human 1:", dataHuman1)
	fmt.Println("Data Human 2:", dataHuman2)

	dataHuman3 := dataHuman2
	dataHuman3.Age = 27

	fmt.Println("Data Human 1:", dataHuman1)
	fmt.Println("Data Human 2:", dataHuman2)
	fmt.Println("Data Human 3:", dataHuman3)

	// kita juga bisa melakukan hal seperti ini, membuat DataHuman baru dan menambahkan pointer, jadi variabel yang menampung variable ini akan langsung meubah utamanya
	fajarData := &DataHuman{Name: "Fajar", Age: 19}
	fajarData.Age = 20

	fmt.Println("Fajar Data:", fajarData)

	emaData := fajarData
	emaData.Name = "Ema"

	fmt.Println("Fajar Data:", fajarData)
	fmt.Println("Ema Data:", emaData)

	// * jika berada di samping kanan yang diketahui itu adalah tipe data,
	// jika berada di samping kiri yang diketahui itu adalah pointer
	// penjelasan lengkapnya https://youtu.be/IO_vkyJnMas?t=17169
	// intinya jika * berada di samping kiri itu adalah pointer yang akan mengubah semua data yang ada di dalamnya, sehingga data yang lain juga akan berubah
	// mudahnya variable manapun yang mengacu pada data variable lama akan berubah
	*dataHuman2 = DataHuman{Name: "Jyriz", Age: 25} // ini akan mengubah dataHuman1 dan dataHuman3

	fmt.Println("Data Human 1:", dataHuman1)
	fmt.Println("Data Human 2:", dataHuman2)
	fmt.Println("Data Human 3:", dataHuman3)
}
